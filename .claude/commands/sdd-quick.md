---
description: SDD trilha leve — quick spec para bugfix ou mudança pequena
argument-hint: <descrição do bug ou da mudança>
---

Siga `.claude/sdd/workflow.md` (seção "Triagem") e `.claude/sdd/guides/ears.md`. Pedido: $ARGUMENTS

1. Confirme que cabe na trilha leve (sem rota/tabela nova, sem mudança de contrato, uma camada/app). Se não couber, recomende `/sdd-specify`.
2. Crie `.claude/sdd/specs/NNN-slug/quick-spec.md` a partir de `.claude/sdd/templates/quick-spec.md` (próximo `NNN`) e adicione ao índice em `.claude/sdd/specs/README.md`.
3. Bugfix: reproduza e investigue a causa raiz antes de propor a correção. Escreva critérios EARS e as regressões a proteger.
4. Peça aprovação. Com aprovação: branch no padrão; delegue ao agente do app (`go-backend-developer` ou `react-electron-frontend-developer`) o teste que falha primeiro e a correção — escrever em `apps/**` é bloqueado para a sessão principal; confira checks e cobertura (seções aplicáveis de `.claude/sdd/templates/checklists/definition-of-done.md`); status `Implemented`.
5. Se surgir complexidade fora da trilha leve, pare e proponha promover para spec completa.
