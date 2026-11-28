# Checklist de qualidade da spec

"Testes unitários" da especificação. Rode ao fim de `/sdd-specify` e `/sdd-clarify`; a spec só vai para aprovação com todos os itens marcados (ou exceção explicada).

## Conteúdo

- [ ] Sem detalhes de implementação (linguagem, framework, tabela, struct, pasta, código)
- [ ] Focada em valor para o usuário e no problema de negócio
- [ ] "Fora de escopo" preenchido
- [ ] Sem funcionalidades especulativas ("pode ser útil depois")

## Completude

- [ ] Nenhum `[NEEDS CLARIFICATION]` restante
- [ ] Toda história tem prioridade e ao menos um critério de aceite
- [ ] Toda história tem caminho feliz **e** caminho de erro (entrada inválida, recurso inexistente, falha de dependência)
- [ ] Casos de borda listados
- [ ] Suposições explícitas

## Clareza e testabilidade

- [ ] Critérios em EARS, um comportamento por frase, sujeito = sistema
- [ ] Sem termos vagos (rápido, fácil, amigável, robusto, etc.) sem métrica
- [ ] Critérios de sucesso mensuráveis
- [ ] Termos usados de forma consistente (mesma entidade, mesmo nome)
- [ ] IDs únicos (US, AC, REQ, NFR)

## Contexto do projeto

- [ ] Não contradiz specs existentes (ou declara a que substitui)
- [ ] Compatível com a constituição
- [ ] Nenhum segredo ou dado pessoal real
