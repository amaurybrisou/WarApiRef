# GuardLine Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- GuardLine.Settings

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | GuardLine.init | ambiguous | MEDIUM | BagOMatic.init, BankArkel.Init, DAoCTooltips.Init, GuardLine.init, LibGuard.Init, MoraleCircle.init, PartyCast.Init |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | GuardLine.update | ambiguous | MEDIUM | AggroMeter.Update, BagOMatic.update, BuffHead.AdvancedContainers.Update, BuffHead.Setup.AdvancedContainers.Update, BuffHead.Setup.PriorityEffects.Update, BuffHead.Setup.PriorityEffectsItem.Update, BuffHead.Setup.AdvancedCompression.Update, BuffHead.Setup.AdvancedCompressionItem.Update, BuffHead.Setup.Filter.Update, Enemy.Update, GuardLine.update, MoraleCircle.update, PartyCast.Update, MoraleButton.Update, TexturedButtons.Update, TurretRange.Display.Update, TurretRange.Setup.Distances.Update, followTheLeader.Update |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EATemplate_Icons | unresolved |  | — |
| Dependency | LibGuard | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EATemplate_Icons | unresolved |  | — |
| Dependency | LibGuard | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
