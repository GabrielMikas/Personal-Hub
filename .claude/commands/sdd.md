---
description: SDD — orquestra uma funcionalidade de ponta a ponta (triagem, spec, gates de aprovação, implementação pelos agentes de back/front, verificação)
argument-hint: [descrição da funcionalidade | NNN para retomar]
---

Atue, nesta sessão principal, como o **orquestrador SDD**. Leia `.claude/agents/SDD-Agents/sdd-orchestrator.md` e siga integralmente as instruções do corpo do arquivo (ignore o frontmatter), incluindo os gates de aprovação humana com AskUserQuestion.

Delegue somente aos agentes listados na tabela "Sua equipe" daquele arquivo (`sdd-spec-writer`, `sdd-planner`, `sdd-reviewer`, `go-backend-developer`, `react-electron-frontend-developer`), pela ferramenta Agent.

Pedido: $ARGUMENTS

- Se for uma descrição, comece pela triagem.
- Se for um número `NNN`, retome a spec `.claude/sdd/specs/NNN-*/` a partir do status atual.
- Se estiver vazio, mostre as specs em andamento do índice `.claude/sdd/specs/README.md` e pergunte o que fazer.
