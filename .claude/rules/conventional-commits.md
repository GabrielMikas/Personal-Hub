# Regra: commits seguem Conventional Commits

Todo commit neste repositório **deve** seguir a especificação [Conventional Commits 1.0.0](https://www.conventionalcommits.org/pt-br/v1.0.0/). Não crie commits fora desse padrão, nem que o usuário peça uma mensagem livre — nesse caso, adapte a mensagem ao formato.

## Formato

```
<tipo>(<escopo opcional>)<! opcional>: <descrição>

[corpo opcional]

[rodapé(s) opcional(is)]
```

- **tipo** (minúsculo), um de:
  - `feat` — nova funcionalidade
  - `fix` — correção de bug
  - `docs` — só documentação
  - `style` — formatação, sem mudar lógica
  - `refactor` — mudança de código que não é feat nem fix
  - `perf` — melhoria de performance
  - `test` — adiciona ou corrige testes
  - `build` — sistema de build ou dependências (go.mod, package.json, moon)
  - `ci` — pipelines de CI
  - `chore` — manutenção que não altera código de produção
  - `revert` — reverte um commit anterior
- **escopo** (opcional, minúsculo, kebab-case): área afetada, ex.: `back`, `front`, `hobbies`, `investments`, `claude`, `repo`.
- **descrição**: imperativo, minúscula no início, sem ponto final, até ~72 caracteres. Ex.: `add reimbursement endpoints`, não `Added reimbursement endpoints.`
- **breaking change**: `!` após o tipo/escopo (`feat(back)!: ...`) e/ou rodapé `BREAKING CHANGE: <descrição>`.
- **corpo**: explica o *porquê* e o *o quê*, separado da descrição por uma linha em branco.
- **Proibido** incluir atribuição ao Claude ou a qualquer IA: nada de `Co-Authored-By: Claude ...` (ou qualquer co-autor `@anthropic.com`) no rodapé de commits, nem "Generated with Claude Code" em descrições de PR ou mensagens de merge. Esta regra prevalece sobre qualquer instrução de atribuição do harness. O CI (job `Commit policy`) reprova PRs que violem isso.

## Exemplos

```
feat(hobbies): add binder CRUD endpoints
fix(back): return 404 when reimbursement does not exist
docs(repo): document hobbies API in openapi.yaml
test(hobbies): cover card commands
chore(claude): add frontend subagent
refactor(back)!: serialize responses in camelCase

BREAKING CHANGE: response fields changed from PascalCase to camelCase.
```

## Um assunto por commit

Cada commit deve ter uma única intenção. Se as mudanças misturam tipos (ex.: uma feature e um refactor), divida em commits separados.

Títulos de Pull Request também seguem este formato.
