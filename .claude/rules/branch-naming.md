# Regra: branches seguem o padrão de Conventional Commits

Toda branch nova **deve** ser nomeada com os mesmos tipos de [Conventional Commits](./conventional-commits.md):

```
<tipo>/<descricao-curta-em-kebab-case>
```

- **tipo**: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore` ou `revert`.
- **descrição**: minúsculas, palavras separadas por hífen, sem acentos, curta (até ~40 caracteres). Pode incluir o escopo no início, ex.: `feat/hobbies-binders`.

## Exemplos

```
feat/hobbies-binders
feat/front-electron-setup
fix/reimbursement-not-found
docs/openapi-investments
chore/repo-improvements
refactor/back-camelcase-responses
```

## Não use

- `feature/...` (use `feat/...`)
- Nomes sem tipo (`minha-branch`, `teste`)
- Maiúsculas, espaços, underscores ou acentos

## Fluxo

- Nunca commite direto na `main`; crie a branch a partir da `main` atualizada.
- Abra PR da branch para a `main`. O título do PR segue Conventional Commits.
- Se o usuário pedir uma branch com nome fora do padrão, sugira o nome equivalente no padrão antes de criá-la.
