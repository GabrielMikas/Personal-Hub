---
description: SDD fase 1 — cria a spec (o quê/por quê) de uma nova funcionalidade
argument-hint: <descrição da funcionalidade>
---

Siga o protocolo SDD em `.claude/sdd/`. Leia `.claude/sdd/workflow.md` (seções "Triagem" e "1. Especificar"), `.claude/sdd/constitution.md` e `.claude/sdd/guides/ears.md`.

Funcionalidade pedida: $ARGUMENTS

1. Faça a triagem. Se for **trivial**, diga isso e pare. Se for **leve**, recomende `/sdd-quick` e pare, salvo se o usuário insistir.
2. Descubra o próximo número em `.claude/sdd/specs/` (maior `NNN` + 1, começando em `001`) e escolha um slug kebab-case curto.
3. Crie `.claude/sdd/specs/NNN-slug/spec.md` a partir de `.claude/sdd/templates/spec.md`. Leia specs relacionadas e `.claude/memory/MEMORY.md` antes.
4. Escreva só o quê e o porquê. Critérios em EARS com IDs. Marque ambiguidades com `[NEEDS CLARIFICATION: ...]` (no máximo 5, as de maior impacto); o resto vira "Suposições".
5. Rode `.claude/sdd/templates/checklists/spec-quality.md` e adicione a linha da spec ao índice em `.claude/sdd/specs/README.md`.
6. Responda com: caminho da spec, resumo das histórias, marcadores pendentes e próximo passo (`/sdd-clarify NNN` ou pedido de aprovação).

Não escreva código, plano nem tarefas nesta fase.
