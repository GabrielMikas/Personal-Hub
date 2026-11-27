# Memória do projeto PersonalHub

Contexto que deve sobreviver entre sessões do Claude. Lido no início de toda sessão (via `CLAUDE.md`).

## Como manter este arquivo

- Atualize ao fim de tarefas relevantes: decisões, pendências descobertas, mudanças de processo, estado de features.
- Seja conciso: um bullet por fato. Use datas absolutas (AAAA-MM-DD).
- Remova ou corrija itens que ficaram desatualizados ou foram resolvidos (mova pendências resolvidas para o histórico em uma linha).
- Não duplique o que já está no código, no `git log`, no `README.md` ou no `openapi.yaml`.
- **Nunca** registre segredos (tokens, senhas, DSNs reais).
- Se o arquivo passar de ~200 linhas, resuma o histórico antigo.

## Decisões

- 2026-11-26 — Commits e branches seguem Conventional Commits (`.claude/rules/`).
- 2026-11-26 — Backend deve ter ≥ 90% de cobertura de testes.
- 2026-11-26 — Frontend será um app desktop Electron + React + Vite + TypeScript.
- 2026-11-26 — Testes do backend usam SQLite em memória (`src/internal/testdb`, driver puro Go, sem CGO/Docker); falhas de escrita são simuladas com `testdb.FailWrites`.
- 2026-11-27 — Repositório tornado público. `main` protegida: exige o check `Backend tests and coverage` (workflow `.github/workflows/backend.yml`) com branch atualizada, vale também para admins, sem force push nem deleção. Sem exigência de review.
- 2026-11-27 — Proibida atribuição ao Claude/IA em commits, PRs e merges (`.claude/rules/conventional-commits.md`, `.claude/settings.json` com `attribution` vazio). O workflow `commit-policy.yml` (job `No AI attribution`) reprova PRs que violem.
- 2026-11-26 — Token da brapi fica em `apps/personal-hub-back/.env` (`BRAPI_KEY`), nunca no código.

## Pendências

- Bug do driver `glebarez/sqlite`: `Migrator().HasTable` não enxerga tabelas em schemas anexados (`hobbies.*`), então re-migrar o SQLite de teste falha. No Postgres não acontece.
- Respostas da API saem em PascalCase (schemas GORM sem tags `json`), enquanto requests usam camelCase. Avaliar padronizar em camelCase.
- Rotas `/finance/investments/*` estão registradas em `routes.go` sem handlers (ainda não implementadas).
- `apps/personal-hub-front`: só existe `moon.yml`; a task `build` roda `vite dev` e não há task `dev` (chamada por `npm run front-dev`).
- `.moon/workspace.yml` define `defaultBranch: "master"`, mas a branch principal é `main`.

## Ambiente

- Windows sem `winget`; o `choco` precisa de admin. O GitHub CLI foi instalado na versão portátil em `%LOCALAPPDATA%\Programs\gh\bin` (no PATH do usuário).

## Histórico de sessões

- 2026-11-26 — Criados os subagentes `go-backend-developer` e `react-electron-frontend-developer`. Token da brapi movido para o `.env`. PR #1 (APIs de hobbies) mergeado. Histórico da `main` reescrito para remover o token; branches `feature/*` apagadas, incluindo `feature/mikas-mq` (trabalho descartado). Criados `CLAUDE.md`, regras, README e `openapi.yaml` na branch `chore/repo-improvements`.
- 2026-11-26 — Branch `test/back-coverage`: cobertura do backend de 0% para 96,0%. `config.Init` passou a retornar erros; `brapi` virou um `Client` injetável (`GetQuote`); `router.New()` separado de `Run()`; toolchain fixado em `go1.25.14` (o 1.25.0 não tem `covdata`).
- 2026-11-27 — Repo tornado público e `main` protegida com o check de CI do backend (PR #4). Token da brapi rotacionado e o antigo (vazado nos refs do PR #1) revogado.
- 2026-11-27 — Proibida atribuição ao Claude (PR #5). Histórico da `main` reescrito para remover 18 trailers `Co-Authored-By: Claude` (hashes mudaram desde `83375c6`); rodapés removidos das descrições dos PRs #1–#4. A `main` agora exige os checks `Backend tests and coverage` e `No AI attribution`.
