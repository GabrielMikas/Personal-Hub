# Constituição do PersonalHub

Princípios que **todo** plano (`plan.md`) deve respeitar. Inspirada nos artigos da constituição do GitHub Spec Kit, adaptada a este monorepo. O plano tem uma seção "Verificação da constituição" que confere cada artigo; um artigo violado só passa com justificativa na seção "Complexidade justificada" e aprovação do usuário.

Versão: 1.0.0 — ratificada em 2026-11-28.

## Artigo I — Spec antes do código

Nenhuma funcionalidade não trivial é implementada sem `spec.md` aprovada (ou `quick-spec.md`, na trilha leve). Mudou o comportamento? Atualize a spec primeiro.

**Gate:** a spec existe, está com status `Approved` ou posterior e não tem `[NEEDS CLARIFICATION]`.

## Artigo II — Test-first (inegociável)

Para cada critério de aceite, o teste é escrito antes do código de produção e deve falhar pelo motivo esperado antes de ser feito passar.

- Backend: cobertura total ≥ 90% (`.claude/rules/backend-testing.md`). Nenhuma mudança reduz a cobertura.
- Testes nunca chamam serviços externos reais nem o banco de desenvolvimento/produção.

**Gate:** `tasks.md` coloca as tarefas de teste antes das tarefas de implementação de cada história.

## Artigo III — Contrato primeiro

Toda mudança em rota, request ou response da API é desenhada primeiro em `apps/personal-hub-back/openapi.yaml`, e só depois implementada. O frontend consome a API pelo contrato, não pelo código do backend.

**Gate:** o plano lista as mudanças no `openapi.yaml`; `tasks.md` tem uma tarefa de contrato antes das de controller.

## Artigo IV — Arquitetura em camadas

O backend segue schema → repository → command → controller → router (ver `.claude/agents/Go-Agents/go-backend-developer.md`). Dependências só fluem para baixo; controllers nunca acessam repositories.

**Gate:** o plano mapeia cada componente novo em uma camada existente.

## Artigo V — Simplicidade

- Implemente só o que a spec pede. Nada de "pode ser útil depois".
- Nenhum novo projeto no monorepo, serviço ou banco sem justificativa.
- Nenhuma dependência nova sem motivo claro e alternativa considerada.

**Gate:** sem projetos novos? sem funcionalidades especulativas? dependências novas justificadas?

## Artigo VI — Sem abstrações desnecessárias

Use Gin, GORM e React diretamente. Não crie wrappers, camadas genéricas ou "frameworks internos" para um único uso. Um modelo por conceito (sem DTO duplicado sem motivo).

**Gate:** usa o framework diretamente? um modelo por conceito?

## Artigo VII — Integração realista

Prefira testar contra implementações reais e isoladas: SQLite em memória (`src/internal/testdb`) para repositories/commands, `httptest` para controllers e para APIs externas. Mocks só quando não houver alternativa razoável.

**Gate:** a estratégia de testes do plano indica o nível de cada teste (unitário, integração com testdb, controller).

## Artigo VIII — Segurança e segredos

Segredos só em `.env` (não versionado). Nenhum token, senha ou DSN real em spec, plano, código, teste, commit ou memória. Erros 500 não vazam detalhes internos.

**Gate:** novas configurações sensíveis entram em `.env.example` com valor fictício.

## Artigo IX — Rastreabilidade e processo

- Cada requisito tem ID; cada tarefa referencia os requisitos que atende; cada critério de aceite tem teste (ver `guides/traceability.md`).
- Branches, commits e títulos de PR seguem `.claude/rules/` (Conventional Commits, sem atribuição a IA).
- O PR referencia a pasta da spec.

**Gate:** todos os requisitos aparecem em pelo menos uma tarefa.

## Emendas

Alterar esta constituição exige: pedido explícito do usuário, justificativa registrada abaixo, incremento de versão (SemVer: MAJOR remove/redefine artigo, MINOR adiciona, PATCH esclarece) e avaliação de impacto nas specs em andamento.

### Histórico

- 1.0.0 — 2026-11-28 — versão inicial.
