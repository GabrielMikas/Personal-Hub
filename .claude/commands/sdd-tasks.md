---
description: SDD fase 4 — quebra o plano aprovado em tarefas ordenadas e rastreáveis
argument-hint: [NNN]
---

Siga `.claude/sdd/workflow.md` (seção "4. Tarefas") e `.claude/sdd/guides/traceability.md`. Pasta alvo: `.claude/sdd/specs/$ARGUMENTS*/` (se vazio, a mais recente com status `Planned`).

1. Confirme que o plano foi aprovado. Se não, pare.
2. Crie `tasks.md` a partir de `.claude/sdd/templates/tasks.md`:
   - IDs `T001…` em ordem de execução, `[P]` para paralelizáveis, `[USn]` e `→ AC/REQ` em cada tarefa, caminho de arquivo explícito.
   - Agrupe por história em ordem de prioridade; cada história entregável sozinha.
   - Ordem: contrato → testes (devem falhar) → schema → repository → command → controller → rotas → frontend.
3. Preencha a matriz de rastreabilidade: todo AC e REQ com ao menos uma tarefa.
4. Responda com a contagem de tarefas por fase e o que pode rodar em paralelo, e peça aprovação (Gate 3). Com aprovação explícita, status da spec `Tasked`.
