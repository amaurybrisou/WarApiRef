# DAoCBuff Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- DAoCBuffVar

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | DAoCBuff.Initialize | ambiguous | MEDIUM | AdvancedPetAssist.Initialize, AdvancedRenownTraining.Initialize, AggroMeter.Initialize, AnywhereTrainer.Initialize, AnywhereTrainerAdditions.Initialize, BuffHead.Initialize, BuffHead.Setup.Initialize, BuffHead.Setup.SelectFont.Initialize, BuffHead.Setup.SelectColor.Initialize, BuffHead.Setup.SelectTexture.Initialize, BuffHead.Setup.General.Initialize, BuffHead.Setup.Layout.Initialize, BuffHead.Setup.Layout.Manager.Initialize, BuffHead.Setup.Layout.Properties.Initialize, BuffHead.Setup.Demo.Initialize, BuffHead.Setup.AdvancedContainers.Initialize, BuffHead.Setup.AdvancedContainersItem.Initialize, BuffHead.Setup.AdvancedContainersItem.Properties.Initialize, BuffHead.Setup.ContainerDemo.Initialize, BuffHead.Setup.PriorityEffects.Initialize, BuffHead.Setup.PriorityEffectsItem.Initialize, BuffHead.Setup.AdvancedCompression.Initialize, BuffHead.Setup.AdvancedCompressionItem.Initialize, BuffHead.Setup.AdvancedCompressionItemEffect.Initialize, BuffHead.Setup.Performance.Initialize, BuffHead.Setup.Display.Initialize, BuffHead.Setup.Trackers.Initialize, BuffHead.Setup.Container.Initialize, BuffHead.Setup.Filter.Initialize, BuffHead.Setup.EffectCache.Initialize, BuffHead.Setup.EffectCacheTable.Initialize, ClosetGoblin.Initialize, CombatTextNames.Initialize, CombatTextNames.RateLimiter.Initialize, DAoCBuff.Initialize, Enemy.Initialize, HotbarMorale.Initialize, Killer.Initialize, LibGroup.Setup.Initialize, LibGroup.Initialize, LibSlash.Initialize, LibWBTogglerManager.Initialize, MiracleGrowLight.Initialize, PlanB.Initialize, PP.Initialize, PotionBar.Initialize, ShiniesConstants.Initialize, SwiftAssist.Initialize, TexturedButtons.Initialize, TexturedButtons.Setup.Initialize, TexturedButtons.Setup.SelectFont.Initialize, TexturedButtons.Setup.SelectColor.Initialize, TexturedButtons.Setup.Textures.Initialize, TexturedButtons.Setup.AdvancedTextures.Initialize, TexturedButtons.Setup.Cooldown.Initialize, TexturedButtons.Setup.Fonts.Initialize, TexturedButtons.Setup.Tint.Initialize, TexturedButtons.Setup.Misc.Initialize, TexturedButtons.Setup.Actionbar.Initialize, TidyChat.Initialize, TidyChatFrames.Initialize, TidyChatLogs.Initialize, TidyRoll.CustomAutoRoll.Initialize, TidyRoll.Initialize, TidyRollOptions.Initialize, TurretRange.Initialize, TurretRange.Display.Initialize, TurretRange.Setup.Initialize, TurretRange.Setup.Distances.Initialize, TurretRange.Setup.Distance.Initialize, TurretRange.Setup.Display.Initialize, TurretRange.Setup.General.Initialize, WarBoard.Initialize, WarBoard.Options.Initialize, WhoHealedMe.Initialize, WoHReticle.Initialize, followTheLeader.Initialize |
| CallFunction | DAoCTooltips.Init | ambiguous | MEDIUM | BagOMatic.init, BankArkel.Init, DAoCTooltips.Init, GuardLine.init, LibGuard.Init, MoraleCircle.init, PartyCast.Init |
| CallFunction | DAoCBuffSettings.CreateOptionswindow | exact | HIGH | DAoCBuffSettings.CreateOptionswindow |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | DAoCBuff.Shutdown | ambiguous | MEDIUM | AdvancedPetAssist.Shutdown, AggroMeter.Shutdown, AnywhereTrainer.Shutdown, DAoCBuff.Shutdown, Enemy.Shutdown, Killer.Shutdown, PotionBar.Shutdown, TidyRoll.Shutdown, WhoHealedMe.Shutdown |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | DAoCBuff.UpdateWindow | exact | HIGH | DAoCBuff.UpdateWindow |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EATemplate_DefaultWindowSkin | unresolved |  | — |
| Dependency | EA_SettingsWindow | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EATemplate_DefaultWindowSkin | unresolved |  | — |
| Dependency | EA_SettingsWindow | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
