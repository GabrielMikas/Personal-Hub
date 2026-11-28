# Definition of Done

Rodado em `/sdd-verify`. Reporte o resultado de cada item ao usuário (com a saída/valor quando houver).

## Spec

- [ ] Cada critério de aceite tem teste identificado (arquivo + nome do caso) na matriz de rastreabilidade
- [ ] Todas as tarefas de `tasks.md` marcadas `[x]` (ou removidas com justificativa)
- [ ] Desvios da spec durante a implementação registrados em "Clarificações"/"Histórico de mudanças"
- [ ] Status da spec `Implemented`

## Backend (`apps/personal-hub-back`)

- [ ] `gofmt -l ./src` sem saída
- [ ] `go vet ./...` ok
- [ ] `go build ./...` ok
- [ ] `go test ./... -coverprofile=coverage.out -covermode=atomic` ok
- [ ] `go tool cover -func=coverage.out | tail -n 1` ≥ 90.0% (valor: ___)
- [ ] `openapi.yaml` reflete rotas, requests e responses implementados
- [ ] Schemas novos no `AutoMigrate`

## Frontend (`apps/personal-hub-front`), quando aplicável

- [ ] Build e lint/typecheck ok
- [ ] Consome a API conforme o `openapi.yaml`

## Repositório

- [ ] Branch e commits em Conventional Commits, sem atribuição a IA
- [ ] Nenhum segredo versionado
- [ ] `.claude/memory/MEMORY.md` atualizado com decisões/pendências relevantes
- [ ] Descrição do PR aponta para `.claude/sdd/specs/NNN-slug/` e lista os critérios atendidos
