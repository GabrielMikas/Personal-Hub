# Rastreabilidade

Objetivo: qualquer pessoa (ou agente) consegue ir de um requisito até o teste que o prova, e de uma linha de código até o requisito que a justifica.

## IDs

| Item | Formato | Onde nasce |
|---|---|---|
| História de usuário | `US1`, `US2`… | `spec.md` |
| Requisito funcional | `REQ-001`… | `spec.md` |
| Critério de aceite | `AC-<história>.<n>` (ex.: `AC-2.3`) | `spec.md` |
| Requisito não funcional | `NFR-001`… | `spec.md` |
| Decisão técnica | `D-01`… | `plan.md` / `research.md` |
| Tarefa | `T001`… | `tasks.md` |

IDs nunca são reaproveitados. Requisito removido fica riscado (`~~REQ-004~~ removido em AAAA-MM-DD: motivo`).

## Matriz

`tasks.md` termina com a matriz de rastreabilidade:

| Requisito / critério | Tarefas | Teste |
|---|---|---|
| AC-1.1 | T003, T006 | `src/controllers/hobbies/binders_test.go` → `AC-1.1 create binder` |

A coluna "Teste" é preenchida na implementação e conferida no `/sdd-verify`.

## No código e no git

- Nome do caso de teste começa com o ID do critério.
- Corpo do commit pode citar a tarefa (`Refs: T004`); o PR cita a pasta da spec.
- Não espalhe IDs em comentários de código de produção — eles envelhecem mal. A ligação fica nos testes e na matriz.
