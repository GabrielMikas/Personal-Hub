# Guia EARS — critérios de aceite

EARS (*Easy Approach to Requirements Syntax*) foi criado por Alistair Mavin e colegas na Rolls-Royce (IEEE RE'09, doi:10.1109/RE.2009.9). Restringe a escrita do requisito a poucos padrões, o que elimina ambiguidade e deixa cada frase diretamente testável. É o formato usado pelo Kiro para critérios de aceite.

## Estrutura geral

```
[ENQUANTO <pré-condição>,] [QUANDO <gatilho>,] o <sistema> DEVE <resposta>
```

Zero ou mais pré-condições, no máximo um gatilho, exatamente um sistema, uma ou mais respostas.

## Os padrões

Usamos as palavras-chave em português, em maiúsculas. Equivalência com o original em inglês entre parênteses.

| Padrão | Modelo | Exemplo |
|---|---|---|
| Ubíquo (*ubiquitous*) | O `<sistema>` DEVE `<resposta>` | A API DEVE gerar o identificador de cada fichário no servidor. |
| Por evento (*WHEN*) | QUANDO `<gatilho>`, o `<sistema>` DEVE `<resposta>` | QUANDO o usuário criar um fichário com nome válido, a API DEVE responder 201 com o fichário criado. |
| Por estado (*WHILE*) | ENQUANTO `<estado>`, o `<sistema>` DEVE `<resposta>` | ENQUANTO a cotação externa estiver indisponível, a API DEVE responder com o último valor salvo. |
| Comportamento indesejado (*IF/THEN*) | SE `<condição indesejada>`, ENTÃO o `<sistema>` DEVE `<resposta>` | SE o fichário solicitado não existir, ENTÃO a API DEVE responder 404. |
| Funcionalidade opcional (*WHERE*) | ONDE `<funcionalidade presente>`, o `<sistema>` DEVE `<resposta>` | ONDE a integração com a brapi estiver configurada, a API DEVE atualizar as cotações. |
| Complexo | ENQUANTO `<estado>`, QUANDO `<gatilho>`, o `<sistema>` DEVE `<resposta>` | ENQUANTO o fichário estiver cheio, QUANDO o usuário adicionar uma carta, a API DEVE responder 409. |

## Regras práticas

- Um requisito por frase. Se tem "e" ligando duas respostas diferentes, divida.
- Sujeito sempre o sistema (`a API`, `o app`), nunca o usuário.
- Use DEVE para obrigatório. Evite "deveria", "pode", "idealmente", "rápido", "amigável", "etc." — não são testáveis.
- Quantifique: "em até 500 ms no p95", "no máximo 100 itens por página".
- Cubra o caminho feliz (QUANDO) **e** os erros (SE/ENTÃO) de cada história: entrada inválida, recurso inexistente, falha de dependência.
- Status HTTP e mensagens visíveis ao usuário são comportamento observável e podem aparecer no critério; nomes de tabela, structs e pacotes não.
- Mais de três pré-condições? Use uma tabela de decisão na spec em vez de forçar o EARS.
- EARS se encaixa mal em requisitos não funcionais amplos; para eles, use a seção de RNF com métrica.

## Do critério ao teste

Cada critério vira pelo menos um caso de teste. O mapeamento é direto:

| EARS | Teste (table-driven) |
|---|---|
| ENQUANTO / pré-condição | setup / arrange |
| QUANDO / SE (gatilho) | ação / act |
| DEVE (resposta) | asserção / assert |

Use o ID do critério no nome do caso: `{name: "AC-2.1 binder not found returns 404", ...}`.
