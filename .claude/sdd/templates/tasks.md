# Tarefas: <Nome da funcionalidade>

| Campo | Valor |
|---|---|
| Spec | [spec.md](./spec.md) |
| Plano | [plan.md](./plan.md) |
| Status | Draft / Approved (AAAA-MM-DD) |

<!--
Formato de cada tarefa:
- [ ] T001 [P] [US1] Descrição imperativa — `caminho/do/arquivo` → AC-1.1, REQ-002

[P]   = pode rodar em paralelo (arquivos diferentes, sem dependência pendente)
[US1] = história atendida (omitir em Preparação e Acabamento)
→     = requisitos/critérios cobertos

Ordem dentro de cada história: contrato → testes (devem FALHAR) → schema → repository
→ command → controller → rotas → frontend → validação da história.
Marque [x] ao concluir. Não renumere tarefas; tarefa nova recebe o próximo ID livre.
-->

## Fase 1 — Preparação

- [ ] T001 Criar branch `feat/<slug>` a partir da `main` atualizada
- [ ] T002 Atualizar `apps/personal-hub-back/openapi.yaml` com as rotas da spec → REQ-…

## Fase 2 — Fundação (bloqueia as histórias)

Itens compartilhados por mais de uma história (ex.: schema + AutoMigrate).

- [ ] T003 …

## Fase 3 — US1: <título> (P1) 🎯 mínimo entregável

**Teste independente:** como verificar esta história isoladamente.

### Testes (escrever primeiro e confirmar que falham)

- [ ] T004 [P] [US1] Teste do controller para AC-1.1 e AC-1.2 — `src/controllers/<dominio>/<x>_test.go` → AC-1.1, AC-1.2
- [ ] T005 [P] [US1] Teste do command — `src/commands/<dominio>/<x>_test.go` → AC-1.1

### Implementação

- [ ] T006 [US1] … — `src/repositories/<dominio>/<x>.go` → REQ-001
- [ ] T007 [US1] … — `src/commands/<dominio>/<x>.go` → AC-1.1
- [ ] T008 [US1] … — `src/controllers/<dominio>/<x>.go` → AC-1.2
- [ ] T009 [US1] Registrar rota — `src/router/routes.go`

**Checkpoint:** testes da US1 passando.

## Fase 4 — US2: <título> (P2)

…

## Fase final — Acabamento

- [ ] T0XX Rodar `gofmt -l ./src`, `go vet ./...`, `go build ./...`
- [ ] T0XX Rodar testes com cobertura e confirmar ≥ 90%
- [ ] T0XX Conferir `openapi.yaml` contra a implementação
- [ ] T0XX `/sdd-verify`: critérios × testes, status `Implemented`, memória atualizada

## Dependências

- Fase 2 depende da Fase 1; histórias dependem da Fase 2; histórias são independentes entre si salvo indicação.
- …

## Matriz de rastreabilidade

| Requisito / critério | Tarefas | Teste (preenchido na implementação) |
|---|---|---|
| AC-1.1 | T004, T007 | |
| AC-1.2 | T004, T008 | |
| REQ-001 | T006 | |
