---
description: SDD fase 2 — resolve [NEEDS CLARIFICATION] da spec com o usuário
argument-hint: [NNN]
---

Siga `.claude/sdd/workflow.md` (seção "2. Esclarecer"). Spec alvo: `.claude/sdd/specs/$ARGUMENTS*/spec.md` (se vazio, a spec mais recente com status `Draft`).

1. Liste os `[NEEDS CLARIFICATION]` e procure outras ambiguidades (escopo, dados, erros, segurança, termos vagos).
2. Pergunte ao usuário, priorizando escopo > segurança/dados > experiência > detalhes. Use AskUserQuestion com até 4 perguntas por vez, opções concretas e uma recomendada.
3. Registre cada resposta em "Clarificações" (data, pergunta, resposta, itens afetados), atualize critérios e requisitos e remova os marcadores.
4. Rode `.claude/sdd/templates/checklists/spec-quality.md`. Se tudo passar, status `Clarified` e peça ao usuário a aprovação da spec (Gate 1). Só mude para `Approved` com aprovação explícita.

Nunca responda um marcador sozinho.
