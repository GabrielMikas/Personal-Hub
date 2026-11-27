# PersonalHub — instruções para o Claude

## Início de toda sessão (obrigatório)

Antes de qualquer tarefa, leia e siga as regras e o contexto abaixo. Eles são importados automaticamente; se por algum motivo não estiverem no contexto, leia os arquivos manualmente.

- Regras de commit: @.claude/rules/conventional-commits.md
- Regras de branch: @.claude/rules/branch-naming.md
- Regras de testes do backend: @.claude/rules/backend-testing.md
- Memória do projeto (contexto de sessões anteriores): @.claude/memory/MEMORY.md

Essas regras valem para qualquer agente ou subagente que trabalhe no repositório.

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
