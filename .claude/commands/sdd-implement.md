---
description: SDD fase 6 — implementa as tarefas aprovadas em modo test-first
argument-hint: [NNN] [T00X | USn]
---

Siga `.claude/sdd/workflow.md` (seção "6. Implementar"), `.claude/sdd/constitution.md` e as regras em `.claude/rules/`. Alvo: `$ARGUMENTS` (pasta `.claude/sdd/specs/NNN*/`; opcionalmente uma tarefa ou história; se omitido, todas as tarefas pendentes em ordem).

1. Confirme: tarefas aprovadas (status `Tasked` ou `In Progress`) e sem achados CRÍTICOS de análise. Se a análise não foi feita, faça-a antes (como em `/sdd-analyze`).
2. Garanta a branch correta (`.claude/rules/branch-naming.md`), criada da `main` atualizada. Nunca commite na `main`.
3. Status da spec `In Progress`. Execute as tarefas em ordem, respeitando dependências; `[P]` podem ir em paralelo.
4. Delegue backend ao `go-backend-developer` e frontend ao `react-electron-frontend-developer`, passando a pasta da spec, os IDs das tarefas e os critérios de aceite envolvidos.
5. Tarefas de teste: escreva, rode e **confirme a falha** pelo motivo esperado antes de implementar. O nome do caso começa com o ID do critério.
6. Marque `[x]` em `tasks.md` a cada tarefa concluída e preencha a coluna "Teste" da matriz.
7. Se a spec se mostrar errada ou incompleta, pare e consulte o usuário; registre a decisão na spec antes de seguir.
8. Ao fim de cada história (checkpoint), rode os testes. Ao fim de tudo, sugira `/sdd-verify`.

Commits só se o usuário pedir, em Conventional Commits, sem atribuição a IA.
