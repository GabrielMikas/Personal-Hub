# Regra: backend com no mínimo 90% de cobertura de testes

O backend Go (`apps/personal-hub-back`) **deve** manter cobertura total de testes de **pelo menos 90%** (statements).

## Como aplicar

- Todo código novo ou alterado no backend vem acompanhado de testes no mesmo PR.
- Um PR não pode reduzir a cobertura total abaixo de 90%. Se a cobertura atual já estiver abaixo da meta, o PR não pode reduzi-la ainda mais e deve, sempre que possível, aumentá-la.
- Antes de concluir qualquer tarefa no backend, rode os testes e verifique a cobertura (comandos abaixo). Reporte o valor obtido.

## Comandos

Na pasta `apps/personal-hub-back`:

```sh
go test ./... -coverprofile=coverage.out -covermode=atomic
go tool cover -func=coverage.out | tail -n 1   # linha "total:" deve ser >= 90.0%
```

Ou via moon: `moon run personal-hub-back:test`.

`coverage.out` não é versionado.

## Como testar

- Testes com o pacote padrão `testing`, no estilo table-driven, arquivos `*_test.go` ao lado do código.
- **Controllers**: `httptest` + `gin.SetMode(gin.TestMode)`, cobrindo sucesso, 400 (bind inválido), 404 (`gorm.ErrRecordNotFound`) e 500.
- **Commands e repositories**: banco isolado para testes (ex.: Postgres via testcontainers ou um banco de teste configurado por variável de ambiente). Nunca use o banco de desenvolvimento/produção.
- **Services externos** (ex.: brapi): nunca chame a API real; use `httptest.Server` ou injete o cliente HTTP/URL.
- Se uma dependência global (ex.: `config.GetDB()`) dificultar o teste, prefira refatorar para injeção de dependência em vez de deixar o código sem teste.
- `main.go` e código puramente de bootstrap podem ficar fora da meta apenas se forem triviais; justifique no PR.
