# Plano: <Nome da funcionalidade>

| Campo | Valor |
|---|---|
| Spec | [spec.md](./spec.md) |
| Status | Draft / Approved (AAAA-MM-DD) |
| Branch | `feat/<slug>` |

<!--
Regras deste template (apague este comentário ao preencher):
- O plano diz COMO. Só implementa o que a spec pede.
- Antes de escrever, leia o código existente do domínio e liste o que será reutilizado.
- Alto nível: sem blocos longos de código. Detalhes indispensáveis em details/<tema>.md.
- Pesquisa de alternativas em research.md.
-->

## Resumo

Abordagem técnica em 3 a 6 frases.

## Contexto técnico

| Item | Valor |
|---|---|
| Backend | Go 1.25, Gin, GORM, Postgres (testes: SQLite em memória `src/internal/testdb`) |
| Frontend | Electron + React + Vite + TypeScript |
| Domínio(s) | ex.: `hobbies` |
| Serviços externos | ex.: brapi / nenhum |
| Dependências novas | nenhuma / <nome> (justificativa) |

## Verificação da constituição

Marque ✅, ou ❌ com justificativa em "Complexidade justificada". Ver [constitution.md](../../constitution.md).

| Artigo | Gate | Resultado |
|---|---|---|
| I — Spec antes do código | Spec `Approved`, sem `[NEEDS CLARIFICATION]` | |
| II — Test-first | Testes antes da implementação em cada história; cobertura ≥ 90% mantida | |
| III — Contrato primeiro | Mudanças no `openapi.yaml` listadas abaixo | |
| IV — Camadas | Cada componente mapeado em uma camada existente | |
| V — Simplicidade | Sem projeto novo, sem funcionalidade especulativa, dependências justificadas | |
| VI — Sem abstração desnecessária | Framework usado diretamente; um modelo por conceito | |
| VII — Integração realista | Nível de cada teste definido; sem chamadas externas reais | |
| VIII — Segurança | Segredos só no `.env`; erros 500 genéricos | |
| IX — Rastreabilidade | Todos os REQ/AC aparecem no plano | |

## Código existente reutilizado

| O quê | Onde | Como será usado |
|---|---|---|
| | | |

## Componentes por camada

| Camada | Arquivo (novo/alterado) | Responsabilidade | Atende |
|---|---|---|---|
| Schema | `src/schemas/<dominio>/<x>.go` | | REQ-001 |
| Repository | `src/repositories/<dominio>/<x>.go` | | |
| Command | `src/commands/<dominio>/<x>.go` | | AC-1.1 |
| Controller | `src/controllers/<dominio>/<x>.go` | | |
| Router | `src/router/routes.go` | | |
| Config | `src/config/postgres.go` (AutoMigrate) | | |
| Frontend | `apps/personal-hub-front/...` | | |

## Modelo de dados

Tabelas/colunas novas ou alteradas, schema Postgres, chaves, índices e relações. Valores monetários em inteiros.

## Contrato da API (`openapi.yaml`)

| Método | Rota | Request | Respostas | Atende |
|---|---|---|---|---|
| POST | `/...` | `{...}` | 201, 400, 500 | AC-1.1, AC-1.2 |

## Fluxos

Sequência dos casos principais e de erro (texto ou diagrama mermaid).

## Estratégia de testes

| Critério | Nível | Arquivo de teste | Observação |
|---|---|---|---|
| AC-1.1 | Controller (`httptest`) | | |
| AC-1.2 | Command (testdb) | | |

Erros de escrita simulados com `testdb.FailWrites`; APIs externas com `httptest.Server`.

## Decisões técnicas

- **D-01** <decisão> — motivo; alternativas descartadas (detalhes em `research.md`).

## Riscos e mitigação

| Risco | Impacto | Mitigação |
|---|---|---|
| | | |

## Complexidade justificada

Preencha só se algum gate da constituição falhou.

| Violação | Por que é necessária | Alternativa mais simples rejeitada porque |
|---|---|---|
| | | |
