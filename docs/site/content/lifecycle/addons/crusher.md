# Crusher Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnUpdate
- OnInitialize
- OnShutdown

## Saved Variables

- CrusherDB

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | Crusher.OnInitialize | ambiguous | MEDIUM | AddonButtonTogglers.OnInitialize, ArmorGraphicToggle.OnInitialize, Assist.OnInitialize, AuctionTweaker.OnInitialize, AuraAddon.OnInitialize, AuraTooltip.OnInitialize, AutoMark.OnInitialize, AutoSalvage.OnInitialize, AutoSalvage.OptionsWindow.OnInitialize, BWMT.OnInitialize, BarText_Influence.OnInitialize, Biguns.OnInitialize, BloodyMess.OnInitialize, CCM.OnInitialize, ClosetGoblin.OnInitialize, ClosetGoblinCharacterWindow.OnInitialize, ClosetGoblinOptionWindow.OnInitialize, ClosetGoblinZoneWindow.OnInitialize, CMap_config.OnInitialize, Cheeseboard.OnInitialize, CB.OnInitialize, Compass3D.Camera.OnInitialize, Compass3D.OnInitialize, Countdown.OnInitialize, GG.OnInitialize, Crusher.OnInitialize, GG.OnInitialize, DeMoNiCore.OnInitialize, DuelInvite.OnInitialize, DuffTimer.OnInitialize, DwarfTalk.OnInitialize, EZCraft.OnInitialize, EZCraftX.OnInitialize, Emojii.OnInitialize, FastFriends.OnInitialize, FlagCap.OnInitialize, GCDTracker.OnInitialize, GetStats.OnInitialize, GroupIcons.OnInitialize, GroupIconsSG.OnInitialize, GroupRange.Styles.OnInitialize, HammerTime.OnInitialize, HealHoverAssist.OnInitialize, Hopper.OnInitialize, IHYTM.OnInitialize, InfoScroller.OnInitialize, Info_Alert.OnInitialize, Info_DeathBlow.OnInitialize, Info_Loot.OnInitialize, Info_Points.OnInitialize, KeyBar.OnInitialize, KeyBarSettings.OnInitialize, LibStub:NewLibrary.GlobalTable.OnInitialize, LootAlert.OnInitialize, MacroIcons.OnInitialize, MapPin.OnInitialize, Motion.OnInitialize, MouseHint.OnInitialize, MyReasons.OnInitialize, NervAlert.OnInitialize, NoOverheal.OnInitialize, OilTimer.OnInitialize, OverheadFonts.OnInitialize, PTLeader.OnInitialize, PlayEffectsOn.OnInitialize, Pure.OnInitialize, PureCareerBar.OnInitialize, QueueQueuer.OnInitialize, quickperf.OnInitialize, RandomMountUI.OnInitialize, RandomSayings.OnInitialize, ResHelp.OnInitialize, RoR_SoR.OnInitialize, Rotation.OnInitialize, RvRContribution.OnInitialize, ScenarioAlert.OnInitialize, Sequencer.OnInitialize, Shinies.OnInitialize, _G.Shinies.OnInitialize, SimpleCombatText.OnInitialize, Soloq.OnInitialize, SquaredClick.OnInitialize, HDI.OnInitialize, SHI.OnInitialize, Statdoll.onInitialize, Statdoll.Getstats.onInitialize, Statdoll.Local.onInitialize, Statdoll.Options.onInitialize, Statdoll.Stats.onInitialize, StatdollWnd.onInitialize, StatdollWndLight.onInitialize, Swinger.OnInitialize, ThankTheTank.OnInitialize, TTmodules.OnInitialize, Trakario.OnInitialize, WARRatingBuster.OnInitialize, warwhisperer.OnInitialize, wargames.OnInitialize, WCDB.OnInitialize, WCDP.OnInitialize, WindowMovers.OnInitialize, XpStatus.OnInitialize, XpStatusWindowNub.OnInitialize, minesweep.OnInitialize, rorAutoInviter.OnInitialize, wbLeadHelper.OnInitialize, whatsPugSc.OnInitialize |
| CreateWindow | CrusherConfig | exact | HIGH | CrusherConfig |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | AceDBLogoutHandler | ambiguous | MEDIUM | Ace.AceDBLogoutHandler, Crusher.AceDBLogoutHandler, Hopper.AceDBLogoutHandler, Pure.AceDBLogoutHandler, Pure Careerbar.AceDBLogoutHandler, Shinies.AceDBLogoutHandler |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | Crusher.OnUpdate | ambiguous | MEDIUM | ActionFractionWindow.OnUpdate, AdvancedPetAssist.OnUpdate, AggroMeter.OnUpdate, Amethyst.OnUpdate, Asshat.OnUpdate, Assist.OnUpdate, AtlasMap.OnUpdate, AuctionStats.OnUpdate, AutoFocus.OnUpdate, AutoMark.OnUpdate, AutoSalvage.OnUpdate, BlackBookWindow.OnUpdate, BloodyMess.OnUpdate, BuffHead.OnUpdate, BuffHead.Setup.ContainerDemo.OnUpdate, BuffHead.Setup.Demo.OnUpdate, BuffHead.Setup.Layout.OnUpdate, BuffThrottle.OnUpdate, LibConfig.OnUpdate, Calling.OnUpdate, CallingTests.OnUpdate, CastSequence.OnUpdate, Cheeseboard.OnUpdate, CleanTargetWindow.OnUpdate, Clock.OnUpdate, Compass3D.OnUpdate, CoolDownLine.OnUpdate, Countdown.OnUpdate, CramTheSpam.OnUpdate, Crusher.OnUpdate, DPSMeter.OnUpdate, DaemonAssist.OnUpdate, DetauntHelper.OnUpdate, DuffTimer.OnUpdate, EA_Window_LoadingScreen.OnUpdate, EA_Window_OpenParty.OnUpdate, EA_Window_OpenPartyNearby.OnUpdate, EA_Window_OpenPartyWorld.OnUpdate, ScenarioGroupWindow.OnUpdate, EZCraft.OnUpdate, EZCraftX.OnUpdate, EZGuard.OnUpdate, LibConfig.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, Enemy.OnUpdate, FastFriends.OnUpdate, FlagCap.OnUpdate, GCDTracker.OnUpdate, GCDsaver.OnUpdate, LibConfig.OnUpdate, GroupIcons.OnUpdate, GroupIconsSG.OnUpdate, GroupRange.OnUpdate, HealGridProgressBar.OnUpdate, HealGrid.OnUpdate, HealGridUnitTooltip.OnUpdate, HealHoverAssist.OnUpdate, Hopper.OnUpdate, IHYTM.OnUpdate, InfoScroller.OnUpdate, LibConfig.OnUpdate, Killer.OnUpdate, Kwestor.OnUpdate, KwestorTracker.OnUpdate, LibAddonButton.OnUpdate, LibAddonButton.Manager.OnUpdate, LibCooldown.OnUpdate, LibGroup.OnUpdate, LibRange.OnUpdate, LibConfig.OnUpdate, LPET.OnUpdate, MapPin.OnUpdate, MBuff.OnUpdate, MassRefine.OnUpdate, MiracleGrow2.OnUpdate, MiracleGrow.OnUpdate, MiracleGrowLight.OnUpdate, Motion.OnUpdate, MouseHint.OnUpdate, MyReasons.OnUpdate, NerfedButtons.OnUpdate, NerfedEngine.OnUpdate, NBSBCore.OnUpdate, NoOverheal.OnUpdate, Obsidian.OnUpdate, Obsidian.Setup.Demo.OnUpdate, PTLeader.OnUpdate, LibConfig.OnUpdate, QueueQueuer.OnUpdate, QuickWarReport.OnUpdate, QuickWarReport.OnUpdate, ScrollingCombatText.OnUpdate, RVAPI_ColorDialog.OnUpdate, RVAPI_Range.OnUpdate, RVMOD_SquaredDistances.OnUpdate, RealmStatus.OnUpdate, RealmStatusClock.OnUpdate, ReferWindow.OnUpdate, ResHelp.OnUpdate, Rotation.OnUpdate, SORObjs.OnUpdate, SimpleCombatText.OnUpdate, Soloq.OnUpdate, Squared.OnUpdate, SquaredRangeFading.OnUpdate, SquaredScenario.OnUpdate, SquaredWarband.OnUpdate, SquaredClick.OnUpdate, HDI.OnUpdate, SHI.OnUpdate, Statdoll.Getstats.onUpdate, Statdoll.Local.onUpdate, StatdollWnd.onUpdate, StatdollWndLight.onUpdate, SwiftAssist.onUpdate, LibConfig.OnUpdate, TastyButtons.OnUpdate, TidyQueue.OnUpdate, TidyRoll.OnUpdate, TimeToDie.OnUpdate, TokenMachine.OnUpdate, TTitan.UI.OnUpdate, TortallDPSCore.OnUpdate, TurretRange.OnUpdate, Twister.OnUpdate, Vectors.OnUpdate, WARCommander.OnUpdate, WarBoard_FPS.OnUpdate, WarBoard_Session.OnUpdate, WarTriage.OnUpdate, LibConfig.OnUpdate, wargames.OnUpdate, WHMEvents.OnUpdate, WhoHealedMe.OnUpdate, WCDB.OnUpdate, WoHReticle.OnUpdate, LibConfig.OnUpdate, emotes.OnUpdate, fpsbox.OnUpdate, nLootLinkComm.onUpdate, nLootLinkData.onUpdate, scenarioInfo.OnUpdate, xHUD.OnUpdate, PetWindow.OnUpdate, xHUD.OnUpdate, xPanels.OnUpdate, mmt.OnUpdate |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_BackpackWindow | unresolved |  | — |
| Dependency | EA_CraftingSystem | unresolved |  | — |
| Dependency | EASystem_Strings | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_ResourceFrames | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EATemplate_DefaultWindowSkin | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_BackpackWindow | unresolved |  | — |
| Dependency | EA_CraftingSystem | unresolved |  | — |
| Dependency | EASystem_Strings | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_ResourceFrames | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EATemplate_DefaultWindowSkin | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
