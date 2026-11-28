#!/usr/bin/env node
// SDD guard (PreToolUse hook).
//
// Enforces the PersonalHub SDD protocol (.claude/sdd/): application code can only be
// changed by the developer subagents the sdd-orchestrator delegates to.
//
//   apps/personal-hub-back/**   -> only agent_type "go-backend-developer"
//   apps/personal-hub-front/**  -> only agent_type "react-electron-frontend-developer"
//   apps/** (anything else)     -> nobody
//   .claude/settings*.json, .claude/hooks/** -> nobody (the guard cannot be disabled by an agent)
//
// File tools (Write, Edit, MultiEdit, NotebookEdit) are checked by path. Shell tools
// (Bash, PowerShell) are checked heuristically: a command that references a protected
// path and looks like a write is denied.
//
// Escape hatch for humans only: start Claude Code with the environment variable
// SDD_GUARD=off. Agents cannot set it for the running process.

import { readFileSync } from "node:fs";
import path from "node:path";

const RULES = [
  { prefix: "apps/personal-hub-back/", allowed: ["go-backend-developer"], owner: "go-backend-developer" },
  { prefix: "apps/personal-hub-front/", allowed: ["react-electron-frontend-developer"], owner: "react-electron-frontend-developer" },
  { prefix: "apps/", allowed: [], owner: null },
];

const PROTECTED_CONFIG = [/^\.claude\/settings(\.local)?\.json$/, /^\.claude\/hooks\//];

const FILE_TOOLS = new Set(["Write", "Edit", "MultiEdit", "NotebookEdit"]);
const SHELL_TOOLS = new Set(["Bash", "PowerShell"]);

// Commands/flags that modify files. Matched only when the command also references a protected path.
// Read-only checks (go test/vet/build, gofmt -l, npx tsc --noEmit, cat, grep...) are not listed.
const SHELL_WRITE = new RegExp(
  [
    // file-modifying commands, as whole words
    String.raw`(^|[\s;&|(])(sed\s+(-[a-zA-Z]*i|--in-place)|perl\s+-[a-zA-Z]*i|tee|rm|rmdir|mv|cp|touch|mkdir|ln|truncate|dd|patch|chmod|unlink` +
      String.raw`|gofmt\s+(-[a-z]+\s+)*-w|goimports\s+(-[a-z]+\s+)*-w|go\s+(fmt|get|mod\s+(tidy|edit|init)|generate|fix)` +
      String.raw`|npm\s+(i|install|ci|uninstall|remove|rm|add|update|init|pkg)|yarn\s+(add|remove|install|upgrade)|pnpm\s+(add|remove|install|i|update)` +
      String.raw`|git\s+(checkout|restore|reset|apply|stash|clean|rm|mv|am|cherry-pick|revert)` +
      String.raw`|Set-Content|Add-Content|Out-File|New-Item|Remove-Item|Move-Item|Copy-Item|Rename-Item|Clear-Content|\[IO\.File\]::Write\w*)(?=[\s;&|)]|$)`,
    // output redirection to a file (not 2>, 2>&1, >/dev/null, >$null)
    String.raw`(?<![0-9&=-])>{1,2}\s*(?!\/dev\/null|&|\$null)[^\s&|>]`,
  ].join("|"),
  "i",
);

function deny(reason) {
  process.stdout.write(
    JSON.stringify({
      hookSpecificOutput: {
        hookEventName: "PreToolUse",
        permissionDecision: "deny",
        permissionDecisionReason: reason,
      },
    }),
  );
  process.exit(0);
}

function toRepoPath(p, projectDir) {
  if (!p) return null;
  const abs = path.resolve(projectDir, p);
  let rel = path.relative(projectDir, abs).split(path.sep).join("/");
  if (rel.startsWith("..")) return null; // outside the repository: not our business
  return process.platform === "win32" ? rel.toLowerCase() : rel;
}

function ruleFor(rel) {
  return RULES.find((r) => rel.startsWith(r.prefix));
}

function explain(rule, agent) {
  const who = agent ? `o agente "${agent}"` : "a sessão principal";
  const owner = rule.owner
    ? `Só o subagente "${rule.owner}" pode alterar ${rule.prefix}`
    : `Nenhum agente está autorizado a alterar ${rule.prefix}`;
  return (
    `Bloqueado pelo protocolo SDD: ${who} não pode alterar código de aplicação. ${owner}. ` +
    `Siga .claude/sdd/workflow.md: a mudança precisa de spec aprovada (ou triagem "trivial" confirmada pelo usuário) ` +
    `e deve ser delegada pelo sdd-orchestrator ao agente de desenvolvimento do app.`
  );
}

const CONFIG_REASON =
  "Bloqueado pelo protocolo SDD: .claude/settings*.json e .claude/hooks/ são protegidos contra edição por agentes. " +
  "Peça ao usuário que faça essa alteração manualmente.";

function main() {
  if ((process.env.SDD_GUARD || "").toLowerCase() === "off") process.exit(0);

  let input;
  try {
    input = JSON.parse(readFileSync(0, "utf8"));
  } catch {
    deny("SDD guard: não foi possível ler a entrada do hook; operação bloqueada por segurança.");
  }

  const tool = input.tool_name;
  const agent = input.agent_type || null;
  const ti = input.tool_input || {};
  const projectDir = process.env.CLAUDE_PROJECT_DIR || input.cwd || process.cwd();

  if (FILE_TOOLS.has(tool)) {
    const rel = toRepoPath(ti.file_path || ti.notebook_path, projectDir);
    if (!rel) process.exit(0);
    if (PROTECTED_CONFIG.some((re) => re.test(rel))) deny(CONFIG_REASON);
    const rule = ruleFor(rel);
    if (rule && !rule.allowed.includes(agent)) deny(explain(rule, agent));
    process.exit(0);
  }

  if (SHELL_TOOLS.has(tool)) {
    const cmd = String(ti.command || "");
    const norm = cmd.replace(/\\/g, "/");
    if (!SHELL_WRITE.test(cmd)) process.exit(0);

    if (/\.claude\/(settings(\.local)?\.json|hooks\/)/i.test(norm)) deny(CONFIG_REASON);

    // Which app paths does the command touch? A cd into an app counts as touching it.
    const touched = RULES.filter((r) => r.owner).filter((r) => norm.toLowerCase().includes(r.prefix.slice(0, -1)));
    const touchesOtherApps = /(^|[\s"'=/])apps\/(?!personal-hub-back\b|personal-hub-front\b)[\w.-]+/i.test(norm);
    if (touchesOtherApps) deny(explain(RULES[2], agent));
    for (const rule of touched) {
      if (!rule.allowed.includes(agent)) deny(explain(rule, agent));
    }
    process.exit(0);
  }

  process.exit(0);
}

main();
