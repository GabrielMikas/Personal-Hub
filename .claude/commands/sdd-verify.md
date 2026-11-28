---
description: SDD fase 7 — verifica critérios × testes, roda a Definition of Done e fecha a spec
argument-hint: [NNN]
---

Siga `.claude/sdd/workflow.md` (seção "7. Verificar"). Pasta alvo: `.claude/sdd/specs/$ARGUMENTS*/`.

1. Para **cada** critério de aceite (e REQ), localize o teste que o prova (arquivo + nome do caso) e confirme que ele verifica de fato o comportamento descrito. Critério sem teste convincente = pendente.
2. Rode `.claude/sdd/templates/checklists/definition-of-done.md`, executando os comandos (em `apps/personal-hub-back`: `gofmt -l ./src`, `go vet ./...`, `go build ./...`, testes com cobertura). Reporte a cobertura total.
3. Compare o `openapi.yaml` com as rotas implementadas.
4. Se tudo passar: status `Implemented`, matriz completa, índice em `.claude/sdd/specs/README.md` atualizado, `.claude/memory/MEMORY.md` atualizado e um rascunho de descrição de PR que aponta para a pasta da spec e lista os critérios atendidos.
5. Se algo falhar: liste as pendências e não mude o status.
