# Ges Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- Ges.Settings

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | Ges.Init | ambiguous | MEDIUM | ArsenalRank.init, AuctionAssist.Init, AutoBand.init, AutoFocus.Init, BagOMatic.init, BankArkel.Init, BuddyBind.init, Calling.Init, ChattyCathy.Init, DAoCTooltips.Init, tactic.init, EffigyBar.Init, FightFinder.init, GuardLine.init, ItemRack.Backpack.Init, LibRuString.Init, LibGuard.Init, LibSkillicon.Init, mech.init, MiniMapMonster.Init, MoraleBG.Init, MoraleCircle.init, Nisp.Init, lnHandler.Init, lnConfig.Init, PartyCast.Init, PeaceOut.Init, RvRStats.Init, RvRStatsRvRTab.Init, RvRStats.Init, RvRStatsRvRTab.Init, ShowHealth.Init, Targets.init, TTitan.Init, TTitan.UI.Init, WARCommanderConfig.Init, WARCommander.Init, ZonePOP.init, nLootLinkData.init, nLootLinkGUIIntegration.init, nLootLinkOptions.init, nLootLinkChat.init, nLootLink.init, nLootLinkGatherer.init, nLootLinkGUI.init, nLootLinkComm.init, nRarity.init |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | Ges.Shutdown | ambiguous | MEDIUM | AbilityAlert.Shutdown, ActionFractionWindow.Shutdown, AdjustTheTip.Shutdown, AdvancedPetAssist.Shutdown, AggroMeter.Shutdown, AnywhereTrainer.Shutdown, Atlas.Shutdown, Atlas.Configuration.Shutdown, AtlasMap.Shutdown, AuctionStats.Shutdown, AutoBand.Shutdown, AutoDismount.Shutdown, BlackBook.Shutdown, CCTV.Shutdown, CDown.Shutdown, CaVES.Shutdown, CallingSetup.Shutdown, CharacterScreenTabFix.Shutdown, ChatAlert.Shutdown, CleanTargetWindow.Shutdown, CraftValueTip.Shutdown, CramTheSpam.Shutdown, DAoCBuff.Shutdown, DaemonAssist.Shutdown, DeepSleep.Shutdown, DuffTimer.Shutdown, Enemy.Shutdown, EveryBodyGuard.Shutdown, Standard.Shutdown, _Squared.Shutdown, Greedy.Shutdown, GuardBot.Shutdown, HealGridSkin.Shutdown, HealGridTags.Shutdown, HealGridIcon.Shutdown, HealGrid.Shutdown, HealGridBuffTrack.Shutdown, HealGridMouseClick.Shutdown, HealGridUnitTooltip.Shutdown, HealGridGuiColorSelect.Shutdown, HealGridGuiSpellList.Shutdown, HealGridGuiTabGeneral.Shutdown, HealGridGuiTabHUD.Shutdown, HealGridGuiTabHUDBuffs.Shutdown, HealGridGuiTabGroup.Shutdown, HealGridGuiTabBattlegroup.Shutdown, HealGridGuiTabScenariogroup.Shutdown, HealGridGuiTabRangeScan.Shutdown, HealGridGuiTabSpellTrack.Shutdown, HealGridGuiTabMouseClick.Shutdown, HealGridGuiTabTooltip.Shutdown, HealGridGui.Shutdown, Killer.Shutdown, KwestorGuiTabGeneral.Shutdown, KwestorGuiTabDisplay.Shutdown, KwestorGuiTabArea.Shutdown, KwestorGui.Shutdown, Kwestor.Shutdown, KwestorTracker.Shutdown, LibRange.Shutdown, LPET.Shutdown, Map.Shutdown, MapMonster.Editor.Shutdown, MapMonster.PinTypeEditor.Shutdown, mech.shutdown, Minmap.Shutdown, NerfedButtons.Shutdown, PartyAd.Shutdown, PotionBar.Shutdown, QuickWarReport.Shutdown, ScrollingCombatText.Shutdown, RVAPI_ColorDialog.Shutdown, RVAPI_Range.Shutdown, RVMOD_3DPortrait.Shutdown, RVMOD_Manager.Shutdown, RVMOD_PlayerStatus.Shutdown, RVMOD_SquaredDistances.Shutdown, RVMOD_Targets.Shutdown, RaidMeter.Shutdown, RandomMount.Shutdown, RangeFadingSquared.Shutdown, ReferWindow.Shutdown, ReliquaryHunter.Shutdown, RezEmote.Shutdown, snt_info.tdps_module.Shutdown, SessionRPs.Shutdown, StopRes.Shutdown, TacticSetNames.Shutdown, TastyButtons.Shutdown, TastyButtonsOptions.Shutdown, ThankTheResser.Shutdown, ThinkOutLoud.Shutdown, TOLSettingsUI.Shutdown, TidyRoll.Shutdown, TTitan.UI.Shutdown, TortallDPSCore.Shutdown, Vectors.Shutdown, Vis.Shutdown, VerticalMorale.Shutdown, VerticalTactics.Shutdown, TacticsEditor.Shutdown, VerticalTactics.Shutdown, WARRatingBuster.Shutdown, WhoHealedMe.Shutdown, XpStatus.Shutdown, hideInf.Shutdown, zmm.Shutdown |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | Ges.OnUpdate | ambiguous | MEDIUM | ActionFractionWindow.OnUpdate, AdvancedPetAssist.OnUpdate, AggroMeter.OnUpdate, Amethyst.OnUpdate, AtlasMap.OnUpdate, AuctionStats.OnUpdate, AutoFocus.OnUpdate, AutoMark.OnUpdate, AutoSalvage.OnUpdate, BlackBookWindow.OnUpdate, BloodyMess.OnUpdate, BuffHead.OnUpdate, BuffHead.Setup.Layout.OnUpdate, BuffHead.Setup.Demo.OnUpdate, BuffHead.Setup.ContainerDemo.OnUpdate, BuffThrottle.OnUpdate, Calling.OnUpdate, CallingTests.OnUpdate, CastSequence.OnUpdate, CleanTargetWindow.OnUpdate, Compass3D.OnUpdate, CoolDownLine.OnUpdate, Countdown.OnUpdate, CramTheSpam.OnUpdate, Crusher.OnUpdate, DPSMeter.OnUpdate, DaemonAssist.OnUpdate, DetauntHelper.OnUpdate, DuffTimer.OnUpdate, ScenarioGroupWindow.OnUpdate, EZCraftX.OnUpdate, LibConfig.OnUpdate, EZGuard.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, FastFriends.OnUpdate, FlagCap.OnUpdate, GCDTracker.OnUpdate, LibConfig.OnUpdate, GCDsaver.OnUpdate, GroupIcons.OnUpdate, GroupIconsSG.OnUpdate, GroupRange.OnUpdate, HealGridProgressBar.OnUpdate, HealGrid.OnUpdate, HealGridUnitTooltip.OnUpdate, HealHoverAssist.OnUpdate, Hopper.OnUpdate, IHYTM.OnUpdate, InfoScroller.OnUpdate, LibConfig.OnUpdate, Killer.OnUpdate, Kwestor.OnUpdate, KwestorTracker.OnUpdate, LibAddonButton.OnUpdate, LibAddonButton.Manager.OnUpdate, LibCooldown.OnUpdate, LibGroup.OnUpdate, LibRange.OnUpdate, LibConfig.OnUpdate, LPET.OnUpdate, MapPin.OnUpdate, MBuff.OnUpdate, MassRefine.OnUpdate, MiracleGrow2.OnUpdate, MiracleGrow.OnUpdate, MiracleGrowLight.OnUpdate, Motion.OnUpdate, MouseHint.OnUpdate, NerfedButtons.OnUpdate, NerfedEngine.OnUpdate, NBSBCore.OnUpdate, NoOverheal.OnUpdate, Obsidian.OnUpdate, Obsidian.Setup.Demo.OnUpdate, LibConfig.OnUpdate, QueueQueuer.OnUpdate, QuickWarReport.OnUpdate, QuickWarReport.OnUpdate, ScrollingCombatText.OnUpdate, RVAPI_ColorDialog.OnUpdate, RVAPI_Range.OnUpdate, RVMOD_SquaredDistances.OnUpdate, RealmStatus.OnUpdate, RealmStatusClock.OnUpdate, ReferWindow.OnUpdate, ResHelp.OnUpdate, Rotation.OnUpdate, SORObjs.OnUpdate, SimpleCombatText.OnUpdate, Soloq.OnUpdate, Squared.OnUpdate, SquaredScenario.OnUpdate, SquaredWarband.OnUpdate, SquaredRangeFading.OnUpdate, SquaredClick.OnUpdate, HDI.OnUpdate, SHI.OnUpdate, Statdoll.Local.onUpdate, Statdoll.Getstats.onUpdate, StatdollWnd.onUpdate, StatdollWndLight.onUpdate, SwiftAssist.onUpdate, LibConfig.OnUpdate, TastyButtons.OnUpdate, TidyQueue.OnUpdate, TidyRoll.OnUpdate, TimeToDie.OnUpdate, TokenMachine.OnUpdate, TTitan.UI.OnUpdate, TortallDPSCore.OnUpdate, TurretRange.OnUpdate, Twister.OnUpdate, Vectors.OnUpdate, WARCommander.OnUpdate, WarBoard_FPS.OnUpdate, WarBoard_Session.OnUpdate, LibConfig.OnUpdate, WarTriage.OnUpdate, wargames.OnUpdate, WHMEvents.OnUpdate, WhoHealedMe.OnUpdate, WCDB.OnUpdate, LibConfig.OnUpdate, WoHReticle.OnUpdate, emotes.OnUpdate, nLootLinkData.onUpdate, nLootLinkComm.onUpdate, scenarioInfo.OnUpdate, xHUD.OnUpdate, PetWindow.OnUpdate, xHUD.OnUpdate, xPanels.OnUpdate, mmt.OnUpdate |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | LibSlash | unresolved |  | — |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | LibSlash | unresolved |  | — |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
