# PersonalHub — instruções para o Claude

## Início de toda sessão (obrigatório)

Antes de qualquer tarefa, leia e siga as regras e o contexto abaixo. Eles são importados automaticamente; se por algum motivo não estiverem no contexto, leia os arquivos manualmente.

- Regras de commit: @.claude/rules/conventional-commits.md
- Regras de branch: @.claude/rules/branch-naming.md
- Regras de testes do backend: @.claude/rules/backend-testing.md
- Memória do projeto (contexto de sessões anteriores): @.claude/memory/MEMORY.md

Essas regras valem para qualquer agente ou subagente que trabalhe no repositório.

## Sessão principal = `sdd-orchestrator` (obrigatório)

- Toda sessão neste repositório roda como o agente `sdd-orchestrator` (`"agent": "sdd-orchestrator"` em `.claude/settings.json`). Ao iniciar, ele lê o protocolo e pergunta se deve retomar uma spec ou começar uma nova.
- **Nenhum desenvolvimento fora dos agentes.** O hook `.claude/hooks/sdd-guard.mjs` (PreToolUse) bloqueia qualquer escrita em `apps/**` — por ferramentas de arquivo ou por comandos shell — que não venha do subagente dono do app: `apps/personal-hub-back/**` só pelo `go-backend-developer`, `apps/personal-hub-front/**` só pelo `react-electron-frontend-developer`. Vale até para mudanças triviais: o orquestrador delega.
- `.claude/settings*.json` e `.claude/hooks/` são protegidos: nenhum agente os edita; só o usuário, manualmente.
- Se um bloqueio acontecer, não tente contorná-lo: siga o fluxo SDD ou peça ao usuário. Desligar o guard é decisão exclusiva do usuário (iniciar o Claude Code com `SDD_GUARD=off`).

## Implementação de funcionalidades (SDD)

Funcionalidades seguem o protocolo de Spec-Driven Development em `.claude/sdd/` (leia `.claude/sdd/README.md` e `.claude/sdd/workflow.md` antes de implementar qualquer coisa não trivial):

- Ponto de entrada: `/sdd <descrição>` (ou `claude --agent sdd-orchestrator`). O orquestrador (`.claude/agents/SDD-Agents/`) faz a triagem, pede aprovação humana e delega a `sdd-spec-writer`, `sdd-planner`, `sdd-reviewer` e aos agentes de back/front.
- Fases avulsas: **leve** (`/sdd-quick`) ou **completa** (`/sdd-specify` → `/sdd-clarify` → `/sdd-plan` → `/sdd-tasks` → `/sdd-analyze` → `/sdd-implement` → `/sdd-verify`).
- Specs ficam em `.claude/sdd/specs/NNN-slug/` e são mantidas após o merge.
- Todo plano passa pelos gates de `.claude/sdd/constitution.md`. O usuário aprova spec, plano e tarefas; nunca se autoaprove.
- Subagentes que recebem tarefas de uma spec devem ler a pasta da spec e respeitar os critérios de aceite.

## Fim de sessão / marcos importantes

Ao concluir uma tarefa relevante (feature, decisão de arquitetura, pendência descoberta, mudança de processo), atualize `.claude/memory/MEMORY.md` conforme as instruções no próprio arquivo. Não registre segredos (tokens, senhas, DSNs reais).

## Visão geral do repositório

Monorepo gerenciado pelo [moon](https://moonrepo.dev) (`.moon/workspace.yml`):

| Projeto | Caminho | Stack |
|---|---|---|
| `personal-hub-back` | `apps/personal-hub-back` | Go 1.25, Gin, GORM, Postgres |
| `personal-hub-front` | `apps/personal-hub-front` | Electron + React + Vite (em construção) |

- Contrato da API: `apps/personal-hub-back/openapi.yaml` — mantenha-o atualizado sempre que rotas, requests ou responses mudarem.
- Variáveis de ambiente do backend: `apps/personal-hub-back/.env` (não versionado), modelo em `.env.example`.
- Subagentes especializados em `.claude/agents/`: `go-backend-developer` (backend) e `react-electron-frontend-developer` (frontend). Prefira delegar a eles tarefas do respectivo app.

## Comandos úteis

```sh
npm run back-dev                       # sobe a API (moon run personal-hub-back:dev) em :3000
moon run personal-hub-back:test        # testes + cobertura do backend
cd apps/personal-hub-back && go vet ./... && go build ./...
```
