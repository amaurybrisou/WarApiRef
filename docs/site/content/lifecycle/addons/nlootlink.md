# nLootLink Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- nLootLinkData.itemDB
- nLootLinkData.versionDB
- nLootLinkData.optionDB

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | nLootLinkData.init | ambiguous | MEDIUM | ArsenalRank.init, AuctionAssist.Init, AutoBand.init, AutoFocus.Init, BagOMatic.init, BankArkel.Init, BuddyBind.init, Calling.Init, ChattyCathy.Init, DAoCTooltips.Init, tactic.init, EffigyBar.Init, FightFinder.init, GuardLine.init, ItemRack.Backpack.Init, LibRuString.Init, LibGuard.Init, LibSkillicon.Init, mech.init, MiniMapMonster.Init, MoraleBG.Init, MoraleCircle.init, Nisp.Init, lnHandler.Init, lnConfig.Init, PartyCast.Init, PeaceOut.Init, RvRStats.Init, RvRStatsRvRTab.Init, RvRStats.Init, RvRStatsRvRTab.Init, ShowHealth.Init, Targets.init, TTitan.Init, TTitan.UI.Init, WARCommanderConfig.Init, WARCommander.Init, ZonePOP.init, nLootLinkData.init, nLootLinkGUIIntegration.init, nLootLinkOptions.init, nLootLinkChat.init, nLootLink.init, nLootLinkGatherer.init, nLootLinkGUI.init, nLootLinkComm.init, nRarity.init |
| CreateWindow | nLootLinkMenuBarWindow | exact | HIGH | nLootLinkMenuBarWindow |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | nLootLinkData.onUpdate | ambiguous | MEDIUM | ActionFractionWindow.OnUpdate, AdvancedPetAssist.OnUpdate, AggroMeter.OnUpdate, Amethyst.OnUpdate, AtlasMap.OnUpdate, AuctionStats.OnUpdate, AutoFocus.OnUpdate, AutoMark.OnUpdate, AutoSalvage.OnUpdate, BlackBookWindow.OnUpdate, BloodyMess.OnUpdate, BuffHead.OnUpdate, BuffHead.Setup.Layout.OnUpdate, BuffHead.Setup.Demo.OnUpdate, BuffHead.Setup.ContainerDemo.OnUpdate, BuffThrottle.OnUpdate, Calling.OnUpdate, CallingTests.OnUpdate, CastSequence.OnUpdate, CleanTargetWindow.OnUpdate, Compass3D.OnUpdate, CoolDownLine.OnUpdate, Countdown.OnUpdate, CramTheSpam.OnUpdate, Crusher.OnUpdate, DPSMeter.OnUpdate, DaemonAssist.OnUpdate, DetauntHelper.OnUpdate, DuffTimer.OnUpdate, ScenarioGroupWindow.OnUpdate, EZCraftX.OnUpdate, LibConfig.OnUpdate, EZGuard.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, FastFriends.OnUpdate, FlagCap.OnUpdate, GCDTracker.OnUpdate, LibConfig.OnUpdate, GCDsaver.OnUpdate, GroupIcons.OnUpdate, GroupIconsSG.OnUpdate, GroupRange.OnUpdate, HealGridProgressBar.OnUpdate, HealGrid.OnUpdate, HealGridUnitTooltip.OnUpdate, HealHoverAssist.OnUpdate, Hopper.OnUpdate, IHYTM.OnUpdate, InfoScroller.OnUpdate, LibConfig.OnUpdate, Killer.OnUpdate, Kwestor.OnUpdate, KwestorTracker.OnUpdate, LibAddonButton.OnUpdate, LibAddonButton.Manager.OnUpdate, LibCooldown.OnUpdate, LibGroup.OnUpdate, LibRange.OnUpdate, LibConfig.OnUpdate, LPET.OnUpdate, MapPin.OnUpdate, MBuff.OnUpdate, MassRefine.OnUpdate, MiracleGrow2.OnUpdate, MiracleGrow.OnUpdate, MiracleGrowLight.OnUpdate, Motion.OnUpdate, MouseHint.OnUpdate, NerfedButtons.OnUpdate, NerfedEngine.OnUpdate, NBSBCore.OnUpdate, NoOverheal.OnUpdate, Obsidian.OnUpdate, Obsidian.Setup.Demo.OnUpdate, LibConfig.OnUpdate, QueueQueuer.OnUpdate, QuickWarReport.OnUpdate, QuickWarReport.OnUpdate, ScrollingCombatText.OnUpdate, RVAPI_ColorDialog.OnUpdate, RVAPI_Range.OnUpdate, RVMOD_SquaredDistances.OnUpdate, RealmStatus.OnUpdate, RealmStatusClock.OnUpdate, ReferWindow.OnUpdate, ResHelp.OnUpdate, Rotation.OnUpdate, SORObjs.OnUpdate, SimpleCombatText.OnUpdate, Soloq.OnUpdate, Squared.OnUpdate, SquaredScenario.OnUpdate, SquaredWarband.OnUpdate, SquaredRangeFading.OnUpdate, SquaredClick.OnUpdate, HDI.OnUpdate, SHI.OnUpdate, Statdoll.Local.onUpdate, Statdoll.Getstats.onUpdate, StatdollWnd.onUpdate, StatdollWndLight.onUpdate, SwiftAssist.onUpdate, LibConfig.OnUpdate, TastyButtons.OnUpdate, TidyQueue.OnUpdate, TidyRoll.OnUpdate, TimeToDie.OnUpdate, TokenMachine.OnUpdate, TTitan.UI.OnUpdate, TortallDPSCore.OnUpdate, TurretRange.OnUpdate, Twister.OnUpdate, Vectors.OnUpdate, WARCommander.OnUpdate, WarBoard_FPS.OnUpdate, WarBoard_Session.OnUpdate, LibConfig.OnUpdate, WarTriage.OnUpdate, wargames.OnUpdate, WHMEvents.OnUpdate, WhoHealedMe.OnUpdate, WCDB.OnUpdate, LibConfig.OnUpdate, WoHReticle.OnUpdate, emotes.OnUpdate, nLootLinkData.onUpdate, nLootLinkComm.onUpdate, scenarioInfo.OnUpdate, xHUD.OnUpdate, PetWindow.OnUpdate, xHUD.OnUpdate, xPanels.OnUpdate, mmt.OnUpdate |
| CallFunction | nLootLinkComm.onUpdate | ambiguous | MEDIUM | ActionFractionWindow.OnUpdate, AdvancedPetAssist.OnUpdate, AggroMeter.OnUpdate, Amethyst.OnUpdate, AtlasMap.OnUpdate, AuctionStats.OnUpdate, AutoFocus.OnUpdate, AutoMark.OnUpdate, AutoSalvage.OnUpdate, BlackBookWindow.OnUpdate, BloodyMess.OnUpdate, BuffHead.OnUpdate, BuffHead.Setup.Layout.OnUpdate, BuffHead.Setup.Demo.OnUpdate, BuffHead.Setup.ContainerDemo.OnUpdate, BuffThrottle.OnUpdate, Calling.OnUpdate, CallingTests.OnUpdate, CastSequence.OnUpdate, CleanTargetWindow.OnUpdate, Compass3D.OnUpdate, CoolDownLine.OnUpdate, Countdown.OnUpdate, CramTheSpam.OnUpdate, Crusher.OnUpdate, DPSMeter.OnUpdate, DaemonAssist.OnUpdate, DetauntHelper.OnUpdate, DuffTimer.OnUpdate, ScenarioGroupWindow.OnUpdate, EZCraftX.OnUpdate, LibConfig.OnUpdate, EZGuard.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, FastFriends.OnUpdate, FlagCap.OnUpdate, GCDTracker.OnUpdate, LibConfig.OnUpdate, GCDsaver.OnUpdate, GroupIcons.OnUpdate, GroupIconsSG.OnUpdate, GroupRange.OnUpdate, HealGridProgressBar.OnUpdate, HealGrid.OnUpdate, HealGridUnitTooltip.OnUpdate, HealHoverAssist.OnUpdate, Hopper.OnUpdate, IHYTM.OnUpdate, InfoScroller.OnUpdate, LibConfig.OnUpdate, Killer.OnUpdate, Kwestor.OnUpdate, KwestorTracker.OnUpdate, LibAddonButton.OnUpdate, LibAddonButton.Manager.OnUpdate, LibCooldown.OnUpdate, LibGroup.OnUpdate, LibRange.OnUpdate, LibConfig.OnUpdate, LPET.OnUpdate, MapPin.OnUpdate, MBuff.OnUpdate, MassRefine.OnUpdate, MiracleGrow2.OnUpdate, MiracleGrow.OnUpdate, MiracleGrowLight.OnUpdate, Motion.OnUpdate, MouseHint.OnUpdate, NerfedButtons.OnUpdate, NerfedEngine.OnUpdate, NBSBCore.OnUpdate, NoOverheal.OnUpdate, Obsidian.OnUpdate, Obsidian.Setup.Demo.OnUpdate, LibConfig.OnUpdate, QueueQueuer.OnUpdate, QuickWarReport.OnUpdate, QuickWarReport.OnUpdate, ScrollingCombatText.OnUpdate, RVAPI_ColorDialog.OnUpdate, RVAPI_Range.OnUpdate, RVMOD_SquaredDistances.OnUpdate, RealmStatus.OnUpdate, RealmStatusClock.OnUpdate, ReferWindow.OnUpdate, ResHelp.OnUpdate, Rotation.OnUpdate, SORObjs.OnUpdate, SimpleCombatText.OnUpdate, Soloq.OnUpdate, Squared.OnUpdate, SquaredScenario.OnUpdate, SquaredWarband.OnUpdate, SquaredRangeFading.OnUpdate, SquaredClick.OnUpdate, HDI.OnUpdate, SHI.OnUpdate, Statdoll.Local.onUpdate, Statdoll.Getstats.onUpdate, StatdollWnd.onUpdate, StatdollWndLight.onUpdate, SwiftAssist.onUpdate, LibConfig.OnUpdate, TastyButtons.OnUpdate, TidyQueue.OnUpdate, TidyRoll.OnUpdate, TimeToDie.OnUpdate, TokenMachine.OnUpdate, TTitan.UI.OnUpdate, TortallDPSCore.OnUpdate, TurretRange.OnUpdate, Twister.OnUpdate, Vectors.OnUpdate, WARCommander.OnUpdate, WarBoard_FPS.OnUpdate, WarBoard_Session.OnUpdate, LibConfig.OnUpdate, WarTriage.OnUpdate, wargames.OnUpdate, WHMEvents.OnUpdate, WhoHealedMe.OnUpdate, WCDB.OnUpdate, LibConfig.OnUpdate, WoHReticle.OnUpdate, emotes.OnUpdate, nLootLinkData.onUpdate, nLootLinkComm.onUpdate, scenarioInfo.OnUpdate, xHUD.OnUpdate, PetWindow.OnUpdate, xHUD.OnUpdate, xPanels.OnUpdate, mmt.OnUpdate |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | LibItemCompress-1.1 | unresolved |  | — |
| Dependency | AceLocale-3.0 | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EA_MenuBarWindow | unresolved |  | — |
| Dependency | EA_BackpackWindow | unresolved |  | — |
| Dependency | EA_CharacterWindow | unresolved |  | — |
| Dependency | EA_CultivationWindow | unresolved |  | — |
| Dependency | EA_CraftingSystem | unresolved |  | — |
| Dependency | EA_TomeOfKnowledge | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | LibItemCompress-1.1 | unresolved |  | — |
| Dependency | AceLocale-3.0 | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EA_MenuBarWindow | unresolved |  | — |
| Dependency | EA_BackpackWindow | unresolved |  | — |
| Dependency | EA_CharacterWindow | unresolved |  | — |
| Dependency | EA_CultivationWindow | unresolved |  | — |
| Dependency | EA_CraftingSystem | unresolved |  | — |
| Dependency | EA_TomeOfKnowledge | unresolved |  | — |
