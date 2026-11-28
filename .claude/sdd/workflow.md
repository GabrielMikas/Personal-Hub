# Workflow SDD

Regras de cada fase. Vale para o agente principal e para subagentes (`go-backend-developer`, `react-electron-frontend-developer`).

## Papéis

| Fase | Responsável | Aprovação |
|---|---|---|
| 0. Triagem | `sdd-orchestrator` | usuário confirma a trilha |
| 1. Especificar | `sdd-spec-writer` | — |
| 2. Esclarecer | `sdd-orchestrator` pergunta; `sdd-spec-writer` incorpora | **Gate 1** (usuário) |
| 3. Planejar | `sdd-planner` | **Gate 2** (usuário) |
| 4. Tarefas | `sdd-planner` | **Gate 3** (usuário) |
| 5. Analisar | `sdd-reviewer` (somente leitura) | — |
| 6. Implementar | `go-backend-developer` / `react-electron-frontend-developer` | — |
| 7. Verificar | `sdd-reviewer`; orquestrador fecha a spec | commit/PR só com pedido do usuário |

Só o orquestrador muda o `Status` da spec e marca tarefas em `tasks.md`. Agentes de desenvolvimento reportam o progresso; não editam a spec. Sem orquestrador (comandos `/sdd-*` avulsos), a sessão principal assume esses papéis — exceto escrever em `apps/**`, que continua delegado.

## Enforcement

- `.claude/settings.json` define `"agent": "sdd-orchestrator"`: toda sessão começa com o orquestrador como sessão principal.
- O hook PreToolUse `.claude/hooks/sdd-guard.mjs` nega escrita em `apps/personal-hub-back/**` a quem não for `go-backend-developer`, em `apps/personal-hub-front/**` a quem não for `react-electron-frontend-developer` e em qualquer outro `apps/**` a todos. Vale para Write/Edit/MultiEdit/NotebookEdit e, por heurística, para comandos Bash/PowerShell que escrevem (redirecionamento, `sed -i`, `rm`, `gofmt -w`, `go mod tidy`, `npm install`, `git checkout --` etc.). Comandos de leitura e checagem (`go test`, `go vet`, `gofmt -l`, `npx tsc --noEmit`) passam.
- `.claude/settings*.json` e `.claude/hooks/` são protegidos contra qualquer agente.
- Desligar: só o usuário, iniciando o Claude Code com a variável de ambiente `SDD_GUARD=off`. Sessão sem o orquestrador: `claude --agent <outro>` (a flag prevalece sobre o setting); o guard continua ativo.

## 0. Triagem

Antes de começar, classifique a mudança e diga ao usuário qual trilha vai seguir:

| Trilha | Quando | Artefatos |
|---|---|---|
| **Trivial** | Typo, texto, formatação, bump de dependência sem mudança de comportamento, ajuste de config/docs | Nenhum. Commit direto em branch. |
| **Leve** | Bugfix, ou mudança pequena em uma camada/um app, sem rota nova nem tabela nova, critérios claros | `quick-spec.md` (1 aprovação) |
| **Completa** | Funcionalidade nova, rota/tabela/tela nova, mudança de contrato, impacto em mais de uma camada ou app, ou requisito ambíguo | `spec.md` → `plan.md` → `tasks.md` (3 aprovações) |

Na dúvida entre duas trilhas, escolha a mais completa ou pergunte. Uma trilha leve que revela complexidade (contrato muda, nova tabela) é **promovida** para a completa: a quick-spec vira base da `spec.md`.

## Pasta da spec

- `.claude/sdd/specs/NNN-slug/`, com `NNN` sequencial de 3 dígitos (próximo número = maior existente + 1) e `slug` em kebab-case, sem acentos. Ex.: `003-hobbies-binders`.
- Branch da implementação segue `.claude/rules/branch-naming.md` e reaproveita o slug: `feat/hobbies-binders`.
- A pasta **fica no repositório depois do merge** (spec-anchored). Mudanças futuras na funcionalidade editam a spec existente ou criam uma nova que a referencia em "Specs relacionadas".

## Status da spec

Campo `Status` no cabeçalho do `spec.md`, atualizado a cada transição:

```
Draft → Clarified → Approved → Planned → Tasked → In Progress → Implemented
                                                            ↘ Superseded (por NNN) / Abandoned
```

Só o usuário move `Clarified → Approved`, aprova o plano (`→ Planned`) e as tarefas (`→ Tasked`).

## 1. Especificar (`/sdd-specify`)

- Entrada: descrição em linguagem natural do usuário.
- Leia `constitution.md`, `.claude/memory/MEMORY.md` e specs relacionadas.
- Preencha `templates/spec.md`. Foque em **o quê** e **por quê**: usuários, histórias, critérios de aceite em EARS (`guides/ears.md`), requisitos não funcionais, fora de escopo.
- **Proibido** na spec: linguagem, framework, nome de tabela, estrutura de pastas, código. Se a decisão técnica for um requisito do usuário (ex.: "precisa funcionar offline"), escreva como restrição observável.
- Cada ambiguidade vira `[NEEDS CLARIFICATION: pergunta específica]`. Não invente respostas. Limite: os 5 pontos de maior impacto; o resto vira suposição explícita na seção "Suposições".
- Rode `templates/checklists/spec-quality.md` e informe o resultado.
- Status: `Draft`.

## 2. Esclarecer (`/sdd-clarify`)

- Pergunte ao usuário **uma pergunta por vez** (ou em lote pequeno com AskUserQuestion), priorizando escopo > segurança/dados > experiência > detalhes. Ofereça opções com uma recomendada.
- Registre cada resposta na seção "Clarificações" com data, e atualize os requisitos afetados. Remova o marcador.
- Quando não houver marcadores e o checklist de qualidade passar: status `Clarified`. Peça aprovação.
- **Gate 1:** usuário aprova a spec → `Approved`.

## 3. Planejar (`/sdd-plan`)

- Pré-requisito: spec `Approved`.
- Leia o código existente do mesmo domínio antes de propor (o domínio `hobbies` é a referência do backend). Registre o que será **reutilizado** — o agente tende a recriar o que já existe.
- Preencha `templates/plan.md`: contexto técnico, verificação da constituição, componentes por camada, modelo de dados, mudanças no `openapi.yaml`, estratégia de testes, riscos.
- Pesquisas (libs, APIs externas, alternativas) vão em `research.md` na mesma pasta: decisão, motivo, alternativas descartadas.
- O plano fica em alto nível. Trechos longos de código não entram; quando indispensável, use `details/<tema>.md`.
- Todo gate da constituição marcado ✅ ou com justificativa em "Complexidade justificada".
- **Gate 2:** usuário aprova o plano → `Planned`.

## 4. Tarefas (`/sdd-tasks`)

- Pré-requisito: plano aprovado.
- Preencha `templates/tasks.md`:
  - IDs `T001`, `T002`… em ordem de execução.
  - `[P]` = pode rodar em paralelo (arquivos diferentes, sem dependência pendente).
  - `[US1]` = história atendida; `→ REQ-002, AC-2.1` = requisitos cobertos.
  - Caminho de arquivo explícito em cada tarefa.
  - Ordem dentro de cada história: contrato → testes (devem falhar) → schema → repository → command → controller → rotas → frontend.
  - Agrupadas por história, em ordem de prioridade, para que cada história seja entregável e testável sozinha.
- **Gate 3:** usuário aprova as tarefas → `Tasked`.

## 5. Analisar (`/sdd-analyze`)

Somente leitura. Produz um relatório com severidade (CRÍTICO / ALTO / MÉDIO / BAIXO):

- Requisito ou critério sem tarefa; tarefa sem requisito.
- Critério de aceite sem teste planejado.
- Termos inconsistentes entre spec, plano e tarefas.
- Violações da constituição não justificadas.
- Conflitos com o código existente ou com o `openapi.yaml`.

Itens CRÍTICOS bloqueiam a implementação até serem corrigidos (volte à fase correspondente).

## 6. Implementar (`/sdd-implement`)

- Pré-requisitos: tarefas aprovadas, análise sem CRÍTICOS, branch criada a partir da `main` atualizada.
- Execute as tarefas em ordem; delegue a subagentes conforme o app (passe a eles o caminho da pasta da spec e os IDs das tarefas).
- Para cada tarefa de teste: escreva o teste, rode, **confirme que falha** pelo motivo certo.
- Marque `[x]` em `tasks.md` ao concluir cada tarefa. Status da spec: `In Progress`.
- Descobriu algo que a spec não previu? **Pare**, registre em "Clarificações" ou proponha emenda da spec ao usuário. Não altere o comportamento silenciosamente.
- Commits pequenos em Conventional Commits; um commit pode referenciar a tarefa no corpo (`Refs: T004`).

## 7. Verificar (`/sdd-verify`)

- Percorra **cada** critério de aceite e aponte o teste (arquivo e nome) que o prova. Critério sem teste = não concluído.
- Rode `templates/checklists/definition-of-done.md` (build, vet, gofmt, testes, cobertura ≥ 90%, `openapi.yaml`).
- Status: `Implemented`. Atualize `.claude/memory/MEMORY.md` (decisões e pendências relevantes).
- O PR referencia `.claude/sdd/specs/NNN-slug/` e lista os critérios atendidos.

## Mudanças depois do merge

1. Abra a spec existente, edite requisitos/critérios (marque o que mudou na seção "Histórico de mudanças").
2. Volte o status para `Draft` e siga o fluxo a partir da fase 1 só para o delta.
3. Se a mudança redefinir a funcionalidade, crie uma spec nova e marque a antiga como `Superseded (por NNN)`.

## Anti-padrões (evite)

- Spec que descreve a implementação ("criar tabela X com coluna Y").
- Plano com funcionalidades que não estão na spec.
- Responder sozinho um `[NEEDS CLARIFICATION]`.
- Pular gates porque "é óbvio".
- Spec completa para mudança trivial (overhead que o processo existe para evitar).
- Considerar a tarefa feita só porque o código compila: a prova é o teste do critério.
