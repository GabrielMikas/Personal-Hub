# Quick spec: <título>

| Campo | Valor |
|---|---|
| ID | NNN-slug |
| Tipo | bugfix / mudança pequena |
| Status | Draft / Approved (AAAA-MM-DD) / Implemented |
| Branch | `fix/<slug>` ou `feat/<slug>` |

<!--
Trilha leve (ver workflow.md → Triagem). Se durante o trabalho surgir rota nova, tabela
nova, mudança de contrato ou impacto em mais de uma camada/app, promova para spec completa.
-->

## Problema

Bugfix: comportamento atual × esperado, passos para reproduzir, evidência (log, resposta HTTP).
Mudança: o que muda e por quê.

## Causa raiz (bugfix)

Preencher após investigar, antes de corrigir. [NEEDS CLARIFICATION: …] se não for possível confirmar.

## Critérios de aceite (EARS)

- **AC-1** QUANDO …, a API DEVE ….
- **AC-2** SE …, ENTÃO a API DEVE ….

## Comportamento que NÃO pode mudar

Regressões a proteger (cada uma com teste existente ou novo).

- …

## Plano e tarefas

- [ ] T001 Teste que reproduz o problema / cobre AC-1 (deve falhar) — `caminho_test.go`
- [ ] T002 Correção — `caminho.go`
- [ ] T003 Testes, cobertura ≥ 90%, `gofmt`/`vet`/`build`; `openapi.yaml` se aplicável

## Verificação

| Critério | Teste |
|---|---|
| AC-1 | |
