# Spec: <Nome da funcionalidade>

| Campo | Valor |
|---|---|
| ID | NNN-slug |
| Status | Draft |
| Criada em | AAAA-MM-DD |
| Atualizada em | AAAA-MM-DD |
| Apps afetados | back / front |
| Specs relacionadas | — |

<!--
Regras deste template (apague este comentário ao preencher):
- Descreva O QUÊ e POR QUÊ. Nada de linguagem, framework, tabela, struct, pasta ou código.
- Toda dúvida vira [NEEDS CLARIFICATION: pergunta específica]. Não adivinhe.
- Critérios de aceite em EARS (.claude/sdd/guides/ears.md).
- Remova seções que não se aplicam em vez de deixá-las vazias ("N/A" se quiser deixar explícito).
-->

## Contexto e problema

Qual problema existe hoje, para quem, e por que resolver agora. 2 a 5 frases.

## Objetivos

- O que muda para o usuário quando isto estiver pronto.

## Fora de escopo

- O que **não** será feito nesta spec (evita escopo implícito).

## Histórias de usuário

Ordenadas por prioridade (P1 = mínimo entregável). Cada história deve ser testável e entregável sozinha.

### US1 — <título> (P1)

**Como** <persona>, **quero** <ação>, **para** <benefício>.

**Critérios de aceite**

- **AC-1.1** QUANDO <gatilho>, a API DEVE <resposta>.
- **AC-1.2** SE <condição indesejada>, ENTÃO a API DEVE <resposta>.

### US2 — <título> (P2)

**Como** …, **quero** …, **para** ….

**Critérios de aceite**

- **AC-2.1** …

## Requisitos funcionais

Regras que valem para mais de uma história ou não cabem num critério.

- **REQ-001** O sistema DEVE …
- **REQ-002** … [NEEDS CLARIFICATION: …]

## Entidades e dados (visão de negócio)

Conceitos envolvidos, atributos relevantes para o usuário e relações. Sem tipos de banco.

| Entidade | Descrição | Atributos principais | Regras |
|---|---|---|---|
| | | | |

## Requisitos não funcionais

- **NFR-001** Desempenho: …
- **NFR-002** Segurança/privacidade: …

## Casos de borda

- O que acontece quando …?

## Suposições

Decisões tomadas sem confirmação explícita, por serem de baixo impacto. O usuário pode contestar.

- …

## Critérios de sucesso

Mensuráveis e independentes de tecnologia.

- …

## Clarificações

Registro das perguntas respondidas (preenchido em `/sdd-clarify`).

- AAAA-MM-DD — P: … → R: … (afeta: AC-1.2)

## Histórico de mudanças

- AAAA-MM-DD — criada.
