# TexturedButtons Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- TexturedButtons.Settings

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | TexturedButtons.Initialize | ambiguous | MEDIUM | AdvancedPetAssist.Initialize, AdvancedRenownTraining.Initialize, AggroMeter.Initialize, AnywhereTrainer.Initialize, AnywhereTrainerAdditions.Initialize, BuffHead.Initialize, BuffHead.Setup.Initialize, BuffHead.Setup.SelectFont.Initialize, BuffHead.Setup.SelectColor.Initialize, BuffHead.Setup.SelectTexture.Initialize, BuffHead.Setup.General.Initialize, BuffHead.Setup.Layout.Initialize, BuffHead.Setup.Layout.Manager.Initialize, BuffHead.Setup.Layout.Properties.Initialize, BuffHead.Setup.Demo.Initialize, BuffHead.Setup.AdvancedContainers.Initialize, BuffHead.Setup.AdvancedContainersItem.Initialize, BuffHead.Setup.AdvancedContainersItem.Properties.Initialize, BuffHead.Setup.ContainerDemo.Initialize, BuffHead.Setup.PriorityEffects.Initialize, BuffHead.Setup.PriorityEffectsItem.Initialize, BuffHead.Setup.AdvancedCompression.Initialize, BuffHead.Setup.AdvancedCompressionItem.Initialize, BuffHead.Setup.AdvancedCompressionItemEffect.Initialize, BuffHead.Setup.Performance.Initialize, BuffHead.Setup.Display.Initialize, BuffHead.Setup.Trackers.Initialize, BuffHead.Setup.Container.Initialize, BuffHead.Setup.Filter.Initialize, BuffHead.Setup.EffectCache.Initialize, BuffHead.Setup.EffectCacheTable.Initialize, ClosetGoblin.Initialize, CombatTextNames.Initialize, CombatTextNames.RateLimiter.Initialize, DAoCBuff.Initialize, Enemy.Initialize, HotbarMorale.Initialize, Killer.Initialize, LibGroup.Setup.Initialize, LibGroup.Initialize, LibSlash.Initialize, LibWBTogglerManager.Initialize, MiracleGrowLight.Initialize, PlanB.Initialize, PP.Initialize, PotionBar.Initialize, ShiniesConstants.Initialize, SwiftAssist.Initialize, TexturedButtons.Initialize, TexturedButtons.Setup.Initialize, TexturedButtons.Setup.SelectFont.Initialize, TexturedButtons.Setup.SelectColor.Initialize, TexturedButtons.Setup.Textures.Initialize, TexturedButtons.Setup.AdvancedTextures.Initialize, TexturedButtons.Setup.Cooldown.Initialize, TexturedButtons.Setup.Fonts.Initialize, TexturedButtons.Setup.Tint.Initialize, TexturedButtons.Setup.Misc.Initialize, TexturedButtons.Setup.Actionbar.Initialize, TidyChat.Initialize, TidyChatFrames.Initialize, TidyChatLogs.Initialize, TidyRoll.CustomAutoRoll.Initialize, TidyRoll.Initialize, TidyRollOptions.Initialize, TurretRange.Initialize, TurretRange.Display.Initialize, TurretRange.Setup.Initialize, TurretRange.Setup.Distances.Initialize, TurretRange.Setup.Distance.Initialize, TurretRange.Setup.Display.Initialize, TurretRange.Setup.General.Initialize, WarBoard.Initialize, WarBoard.Options.Initialize, WhoHealedMe.Initialize, WoHReticle.Initialize, followTheLeader.Initialize |
| CreateWindow | TexturedButtonsSetupSelectColorWindow | exact | HIGH | TexturedButtonsSetupSelectColorWindow |
| CreateWindow | TexturedButtonsSetupMenuWindow | exact | HIGH | TexturedButtonsSetupMenuWindow |
| CreateWindow | TexturedButtonsSetupTexturesWindow | exact | HIGH | TexturedButtonsSetupTexturesWindow |
| CreateWindow | TexturedButtonsSetupAdvancedTexturesWindow | exact | HIGH | TexturedButtonsSetupAdvancedTexturesWindow |
| CreateWindow | TexturedButtonsSetupCooldownWindow | exact | HIGH | TexturedButtonsSetupCooldownWindow |
| CreateWindow | TexturedButtonsSetupFontsWindow | exact | HIGH | TexturedButtonsSetupFontsWindow |
| CreateWindow | TexturedButtonsSetupTintWindow | exact | HIGH | TexturedButtonsSetupTintWindow |
| CreateWindow | TexturedButtonsSetupMiscWindow | exact | HIGH | TexturedButtonsSetupMiscWindow |
| CreateWindow | TexturedButtonsSetupActionbarWindow | exact | HIGH | TexturedButtonsSetupActionbarWindow |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | TexturedButtons.Unload | ambiguous | MEDIUM | LibGroup.Unload, TexturedButtons.Unload, TurretRange.Unload, TurretRange.Display.Unload |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_ActionBars | unresolved |  | — |
| Dependency | EASystem_ActionBarClusterManager | unresolved |  | — |
| Dependency | EA_SettingsWindow | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_ActionBars | unresolved |  | — |
| Dependency | EASystem_ActionBarClusterManager | unresolved |  | — |
| Dependency | EA_SettingsWindow | unresolved |  | — |
