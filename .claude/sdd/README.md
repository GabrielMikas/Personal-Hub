# Protocolo SDD (Spec-Driven Development) do PersonalHub

Protocolo para implementar funcionalidades a partir de uma **especificação escrita antes do código**. A spec é a fonte da verdade para humanos e para agentes: o código implementa a spec, os testes provam a spec.

> Nível adotado: **spec-anchored** (Böckeler, martinfowler.com). A spec nasce antes do código (*spec-first*) e é **mantida** depois do merge; quando a funcionalidade mudar, a spec muda primeiro. Não adotamos *spec-as-source* (código gerado e nunca editado).

## Estrutura

```
.claude/sdd/
├── README.md               # este arquivo: visão geral e índice
├── constitution.md         # princípios imutáveis do projeto (gates do plano)
├── workflow.md             # fases, gates de aprovação, status e regras de cada fase
├── guides/
│   ├── ears.md             # como escrever critérios de aceite em EARS
│   └── traceability.md     # IDs e rastreabilidade requisito → tarefa → teste
├── templates/
│   ├── spec.md             # O QUÊ e POR QUÊ (sem tecnologia)
│   ├── plan.md             # COMO (arquitetura, dados, contrato, testes)
│   ├── tasks.md            # tarefas executáveis, ordenadas, test-first
│   ├── quick-spec.md       # trilha leve: bugfix e mudanças pequenas
│   └── checklists/
│       ├── spec-quality.md     # "testes unitários" da spec
│       └── definition-of-done.md
└── specs/
    └── NNN-slug/           # uma pasta por funcionalidade (spec, plan, tasks)
```

## Agentes

O processo é conduzido por agentes em `.claude/agents/SDD-Agents/`:

| Agente | Papel |
|---|---|
| `sdd-orchestrator` | Conduz a funcionalidade inteira, é o único que fala com o usuário, aplica os gates de aprovação e move o status da spec. Não escreve código. |
| `sdd-spec-writer` | Escreve e atualiza `spec.md` / `quick-spec.md`; devolve perguntas abertas. |
| `sdd-planner` | Escreve `plan.md`, `research.md` e `tasks.md`. |
| `sdd-reviewer` | Somente leitura: análise antes de implementar e verificação critério × teste depois. |
| `go-backend-developer` / `react-electron-frontend-developer` | Implementam as tarefas aprovadas, test-first. |

Como subagentes não conseguem perguntar ao usuário, o orquestrador precisa rodar **como sessão principal**:

```sh
claude --agent sdd-orchestrator     # sessão dedicada, com allowlist de agentes
```

ou, numa sessão normal, `/sdd <descrição>` (ou `/sdd NNN` para retomar), que faz a sessão principal seguir as instruções do orquestrador.

## Comandos

Além do `/sdd`, slash commands por fase ficam em `.claude/commands/` (prefixo `sdd-`), para rodar uma fase isolada manualmente:

| Comando | Fase | Produz |
|---|---|---|
| `/sdd-specify <descrição>` | 1. Especificar | `specs/NNN-slug/spec.md` |
| `/sdd-clarify [NNN]` | 2. Esclarecer | spec sem `[NEEDS CLARIFICATION]` |
| `/sdd-plan [NNN]` | 3. Planejar | `plan.md` (+ `research.md` se preciso) |
| `/sdd-tasks [NNN]` | 4. Quebrar em tarefas | `tasks.md` |
| `/sdd-analyze [NNN]` | 5. Analisar consistência | relatório (não edita arquivos) |
| `/sdd-implement [NNN] [T00X]` | 6. Implementar | código + testes, tarefas marcadas |
| `/sdd-verify [NNN]` | 7. Verificar e fechar | spec `Implemented`, memória atualizada |
| `/sdd-quick <descrição>` | trilha leve | `specs/NNN-slug/quick-spec.md` |

## Fluxo resumido

```
triagem ─┬─ trivial ───────────────────────────────► implementa direto (sem spec)
         ├─ pequena/bugfix ─► quick-spec ─[aprova]─► implementa ─► verifica
         └─ funcionalidade ─► spec ─► clarify ─[aprova]─► plan ─[aprova]─► tasks ─[aprova]
                               ─► analyze ─► implement (test-first) ─► verify
```

Detalhes, critérios de triagem e gates em [`workflow.md`](./workflow.md).

## Princípios que guiam o protocolo

1. **O quê antes do como.** `spec.md` não cita tecnologia; `plan.md` é onde entram Go, Gin, GORM, React.
2. **Marcar, não adivinhar.** Toda ambiguidade vira `[NEEDS CLARIFICATION: pergunta]`. Nada é planejado com marcadores pendentes.
3. **Critérios testáveis.** Critérios de aceite em EARS; cada um vira pelo menos um teste.
4. **Gates humanos.** O usuário aprova spec, plano e tarefas antes da fase seguinte. O agente nunca se autoaprova.
5. **Constituição.** O plano passa pelos gates de [`constitution.md`](./constitution.md); violações exigem justificativa registrada.
6. **Test-first.** Testes do critério de aceite são escritos e vistos falhando antes da implementação.
7. **Proporcionalidade.** O processo escala com o tamanho da mudança; não se escreve spec completa para um typo.
8. **Desconfiar do agente.** Ter spec não garante que o agente a seguiu: a verificação compara código e testes com cada critério, item a item.

## Fontes

O protocolo foi montado a partir de:

- GitHub Spec Kit — [`spec-driven.md`](https://github.com/github/spec-kit/blob/main/spec-driven.md) e [site do projeto](https://github.github.com/spec-kit/): fases specify → clarify → plan → tasks → analyze → implement, constituição com gates no plano, marcadores `[NEEDS CLARIFICATION]`, tarefas `[P]`, checklists como "testes unitários" da spec, ordem contrato → testes → código.
- Birgitta Böckeler, [Understanding Spec-Driven-Development: Kiro, spec-kit, and Tessl](https://www.martinfowler.com/articles/exploring-gen-ai/sdd-3-tools.html) (martinfowler.com, 2025): níveis spec-first / spec-anchored / spec-as-source, diferença entre spec e memory bank, críticas (overhead em mudanças pequenas, fardo de revisão, agente ignorando a spec).
- Thoughtworks Technology Radar — [Spec-driven development](https://www.thoughtworks.com/radar/techniques/spec-driven-development).
- Kiro — [Feature specs / Requirements-first](https://kiro.dev/docs/specs/feature-specs/requirements-first/): requirements → design → tasks com aprovação entre fases, critérios em EARS, trilha de bugfix.
- Mavin, Wilkinson, Harwood, Novak. *Easy Approach to Requirements Syntax (EARS)*. IEEE RE'09, pp. 317–322, doi:10.1109/RE.2009.9 ([resumo](https://en.wikipedia.org/wiki/Easy_Approach_to_Requirements_Syntax)).
- Microsoft Learn — [Examine the GitHub Spec Kit](https://learn.microsoft.com/en-us/training/modules/spec-driven-development-github-spec-kit-enterprise-developers/3-examine-github-spec-kit).
