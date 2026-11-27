# PersonalHub

Hub pessoal para organizar finanças e hobbies num só lugar: uma API em Go e um app desktop em Electron + React.

## Funcionalidades

### Hobbies (disponível na API)

| Recurso | O que faz | Rotas |
|---|---|---|
| **Cards** | Cadastro de cartas colecionáveis: nome, código, valor e data de compra, e o fichário (binder) onde estão guardadas | `/hobbies/cards` |
| **Orders** | Pedidos feitos em lojas: loja, ID do pedido na loja, data, valor e status de entrega | `/hobbies/orders` |
| **Reimbursements** | Pedidos de reembolso ligados a um pedido: loja, contato, valor solicitado e status | `/hobbies/reimbursements` |

Todos os recursos têm CRUD completo (listar, buscar, criar, atualizar parcialmente e remover). O contrato está em [`apps/personal-hub-back/openapi.yaml`](apps/personal-hub-back/openapi.yaml).

### Finanças / investimentos (planejado)

- Rotas `/finance/investments` reservadas para posições em ações e outros investimentos (ainda sem implementação).
- Integração com a [brapi](https://brapi.dev) para cotações de ações da B3 (cliente em `apps/personal-hub-back/src/services/brapi`).

### App desktop (em construção)

- Front-end em Electron + React + Vite + TypeScript, em `apps/personal-hub-front`.

## Estrutura

```
.
├── apps/
│   ├── personal-hub-back/    # API Go (Gin + GORM + Postgres)
│   │   ├── openapi.yaml      # contrato da API
│   │   └── src/
│   │       ├── commands/     # regras de negócio (um command por operação)
│   │       ├── config/       # inicialização e migração do banco
│   │       ├── controllers/  # handlers HTTP
│   │       ├── repositories/ # acesso a dados
│   │       ├── router/       # definição das rotas
│   │       ├── schemas/      # modelos GORM
│   │       └── services/     # integrações externas (brapi)
│   └── personal-hub-front/   # app Electron + React
├── .claude/                  # regras, memória e subagentes do Claude Code
├── .moon/                    # configuração do monorepo (moon)
└── CLAUDE.md
```

## Requisitos

- [Go](https://go.dev) 1.25+
- [Node.js](https://nodejs.org) 20+
- [moon](https://moonrepo.dev)
- PostgreSQL com os schemas `hobbies` criados

## Como rodar

1. Configure as variáveis de ambiente do backend:

   ```sh
   cp apps/personal-hub-back/.env.example apps/personal-hub-back/.env
   # preencha DB_DSN e BRAPI_KEY
   ```

2. Suba a API (porta `3000`). As tabelas são criadas automaticamente via `AutoMigrate`.

   ```sh
   npm run back-dev
   ```

3. Teste:

   ```sh
   curl http://localhost:3000/hobbies/cards/
   ```

### Testes do backend

```sh
moon run personal-hub-back:test
```

A meta de cobertura é de pelo menos **90%**.

Os testes não precisam de Postgres nem de Docker: commands, repositories e controllers usam um SQLite em memória (`src/internal/testdb`), e a brapi é simulada com `httptest.Server`. O `go.mod` fixa o toolchain em `go1.25.14` porque o `go1.25.0` não tem a ferramenta `covdata` e falha ao gerar cobertura.

## Contribuindo

- **Branches:** `<tipo>/<descricao>`, ex.: `feat/hobbies-binders` (veja [`.claude/rules/branch-naming.md`](.claude/rules/branch-naming.md)).
- **Commits:** [Conventional Commits](https://www.conventionalcommits.org/pt-br/v1.0.0/), ex.: `feat(hobbies): add binder endpoints` (veja [`.claude/rules/conventional-commits.md`](.claude/rules/conventional-commits.md)).
- **Testes:** backend com cobertura mínima de 90% (veja [`.claude/rules/backend-testing.md`](.claude/rules/backend-testing.md)).
- **API:** toda mudança de rota, request ou response deve atualizar o `openapi.yaml`.
