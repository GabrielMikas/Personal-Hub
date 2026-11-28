---
description: SDD fase 5 — análise somente leitura de consistência entre spec, plano, tarefas e constituição
argument-hint: [NNN]
---

Siga `.claude/sdd/workflow.md` (seção "5. Analisar"). Pasta alvo: `.claude/sdd/specs/$ARGUMENTS*/`.

**Não edite nenhum arquivo.** Leia `spec.md`, `plan.md`, `tasks.md`, `.claude/sdd/constitution.md`, o `openapi.yaml` e o código dos domínios afetados. Produza uma tabela:

| ID | Severidade | Categoria | Local | Problema | Correção sugerida |

- Categorias: cobertura (REQ/AC sem tarefa ou sem teste; tarefa sem requisito), inconsistência de termos, ambiguidade remanescente, violação da constituição, conflito com código/contrato existente, duplicação.
- Severidades: CRÍTICO (bloqueia implementação), ALTO, MÉDIO, BAIXO.

Termine com métricas (total de REQ/AC, % com tarefa, % com teste planejado) e a recomendação: seguir para `/sdd-implement` ou voltar a qual fase. Ofereça aplicar as correções, mas só aplique se o usuário pedir.
