---
description: SDD fase 3 — cria o plano técnico (como) a partir da spec aprovada
argument-hint: [NNN]
---

Siga `.claude/sdd/workflow.md` (seção "3. Planejar") e `.claude/sdd/constitution.md`. Pasta alvo: `.claude/sdd/specs/$ARGUMENTS*/` (se vazio, a mais recente com status `Approved`).

1. Confirme que a spec está `Approved` e sem `[NEEDS CLARIFICATION]`. Se não, pare e indique `/sdd-clarify`.
2. Leia o código existente do domínio afetado (no backend, `hobbies` é a referência), `apps/personal-hub-back/openapi.yaml` e os subagentes em `.claude/agents/`. Liste o que será reutilizado.
3. Crie `plan.md` na pasta da spec a partir de `.claude/sdd/templates/plan.md`. Pesquisas e alternativas vão em `research.md`.
4. Preencha a verificação da constituição artigo por artigo. Violações só com justificativa em "Complexidade justificada".
5. Não acrescente nada que a spec não peça. Todo REQ/AC deve aparecer no plano.
6. Responda com o resumo do plano, os gates e as decisões, e peça aprovação (Gate 2). Com aprovação explícita, status da spec `Planned`.

Não escreva código de produção nesta fase.
