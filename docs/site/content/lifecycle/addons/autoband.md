# AutoBand Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- AutoBand.saved

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | AutoBand.init | ambiguous | MEDIUM | ArsenalRank.init, AuctionAssist.Init, AutoBand.init, AutoFocus.Init, BagOMatic.init, BankArkel.Init, BuddyBind.init, Calling.Init, ChattyCathy.Init, DAoCTooltips.Init, tactic.init, EffigyBar.Init, FightFinder.init, GuardLine.init, ItemRack.Backpack.Init, LibRuString.Init, LibGuard.Init, LibSkillicon.Init, mech.init, MiniMapMonster.Init, MoraleBG.Init, MoraleCircle.init, Nisp.Init, lnHandler.Init, lnConfig.Init, PartyCast.Init, PeaceOut.Init, RvRStats.Init, RvRStatsRvRTab.Init, RvRStats.Init, RvRStatsRvRTab.Init, ShowHealth.Init, Targets.init, TTitan.Init, TTitan.UI.Init, WARCommanderConfig.Init, WARCommander.Init, ZonePOP.init, nLootLinkData.init, nLootLinkGUIIntegration.init, nLootLinkOptions.init, nLootLinkChat.init, nLootLink.init, nLootLinkGatherer.init, nLootLinkGUI.init, nLootLinkComm.init, nRarity.init |
| CreateWindow | AutoBandMapIcon | exact | HIGH | AutoBandMapIcon |
| CreateWindow | AutoBandWindow | exact | HIGH | AutoBandWindow |
| CreateWindow | AutoBandWindowTemplateSave | exact | HIGH | AutoBandWindowTemplateSave |
| CreateWindow | AutoBandWindowTemplateDeleteConfirm | exact | HIGH | AutoBandWindowTemplateDeleteConfirm |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | AutoBand.Shutdown | ambiguous | MEDIUM | AbilityAlert.Shutdown, ActionFractionWindow.Shutdown, AdjustTheTip.Shutdown, AdvancedPetAssist.Shutdown, AggroMeter.Shutdown, AnywhereTrainer.Shutdown, Atlas.Shutdown, Atlas.Configuration.Shutdown, AtlasMap.Shutdown, AuctionStats.Shutdown, AutoBand.Shutdown, AutoDismount.Shutdown, BlackBook.Shutdown, CCTV.Shutdown, CDown.Shutdown, CaVES.Shutdown, CallingSetup.Shutdown, CharacterScreenTabFix.Shutdown, ChatAlert.Shutdown, CleanTargetWindow.Shutdown, CraftValueTip.Shutdown, CramTheSpam.Shutdown, DAoCBuff.Shutdown, DaemonAssist.Shutdown, DeepSleep.Shutdown, DuffTimer.Shutdown, Enemy.Shutdown, EveryBodyGuard.Shutdown, Standard.Shutdown, _Squared.Shutdown, Greedy.Shutdown, GuardBot.Shutdown, HealGridSkin.Shutdown, HealGridTags.Shutdown, HealGridIcon.Shutdown, HealGrid.Shutdown, HealGridBuffTrack.Shutdown, HealGridMouseClick.Shutdown, HealGridUnitTooltip.Shutdown, HealGridGuiColorSelect.Shutdown, HealGridGuiSpellList.Shutdown, HealGridGuiTabGeneral.Shutdown, HealGridGuiTabHUD.Shutdown, HealGridGuiTabHUDBuffs.Shutdown, HealGridGuiTabGroup.Shutdown, HealGridGuiTabBattlegroup.Shutdown, HealGridGuiTabScenariogroup.Shutdown, HealGridGuiTabRangeScan.Shutdown, HealGridGuiTabSpellTrack.Shutdown, HealGridGuiTabMouseClick.Shutdown, HealGridGuiTabTooltip.Shutdown, HealGridGui.Shutdown, Killer.Shutdown, KwestorGuiTabGeneral.Shutdown, KwestorGuiTabDisplay.Shutdown, KwestorGuiTabArea.Shutdown, KwestorGui.Shutdown, Kwestor.Shutdown, KwestorTracker.Shutdown, LibRange.Shutdown, LPET.Shutdown, Map.Shutdown, MapMonster.Editor.Shutdown, MapMonster.PinTypeEditor.Shutdown, mech.shutdown, Minmap.Shutdown, NerfedButtons.Shutdown, PartyAd.Shutdown, PotionBar.Shutdown, QuickWarReport.Shutdown, ScrollingCombatText.Shutdown, RVAPI_ColorDialog.Shutdown, RVAPI_Range.Shutdown, RVMOD_3DPortrait.Shutdown, RVMOD_Manager.Shutdown, RVMOD_PlayerStatus.Shutdown, RVMOD_SquaredDistances.Shutdown, RVMOD_Targets.Shutdown, RaidMeter.Shutdown, RandomMount.Shutdown, RangeFadingSquared.Shutdown, ReferWindow.Shutdown, ReliquaryHunter.Shutdown, RezEmote.Shutdown, snt_info.tdps_module.Shutdown, SessionRPs.Shutdown, StopRes.Shutdown, TacticSetNames.Shutdown, TastyButtons.Shutdown, TastyButtonsOptions.Shutdown, ThankTheResser.Shutdown, ThinkOutLoud.Shutdown, TOLSettingsUI.Shutdown, TidyRoll.Shutdown, TTitan.UI.Shutdown, TortallDPSCore.Shutdown, Vectors.Shutdown, Vis.Shutdown, VerticalMorale.Shutdown, VerticalTactics.Shutdown, TacticsEditor.Shutdown, VerticalTactics.Shutdown, WARRatingBuster.Shutdown, WhoHealedMe.Shutdown, XpStatus.Shutdown, hideInf.Shutdown, zmm.Shutdown |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | AutoBand.update | ambiguous | MEDIUM | MoraleButton.Update, ActionBarColor.Update, AggroMeter.Update, LayerTimerWindow.Update, Amethyst.Update, ArsenalRank.update, AutoBand.update, AutoSalvage.OptionsWindow.Update, BagOMatic.update, BlackBook.Update, BuddyBind.update, BuffHead.AdvancedContainers.Update, BuffHead.Setup.AdvancedContainers.Update, BuffHead.Setup.PriorityEffects.Update, BuffHead.Setup.PriorityEffectsItem.Update, BuffHead.Setup.AdvancedCompression.Update, BuffHead.Setup.AdvancedCompressionItem.Update, BuffHead.Setup.Filter.Update, CCTV.Update, CastSequence.Builder.Update, TargetUnitFrame.Update, CleanUnitFrames.Update, DPSMeter.local.Update, DPSMeter.Update, tactic.update, EZCraftX.Update, Enemy.Update, Fader.Update, EveryBodyGuard.Update, FastFriends.Update, FightFinder.update, GroupRange.Styles.Update, GroupRange.Styles.Update, GroupRange.Styles.Update, GroupRange.Styles.Update, GuardBot.Update, GuardLine.update, Info_Loot.Update, KeyBar.Update, LibAddonButton.Manager.Update, LootAlert.Update, mech.update, MoraleCircle.update, Moth.Update, MoraleButton.Update, NAMBLA.Update, NaturalLog.Update, LayerTimerWindow.Update, Obsidian.Update, PartyCast.Update, QueuePopTimer.Update, LayerTimerWindow.Update, snt_castbar.update, SOR.Update, SessionRPs.Update, Statdoll.Update, Statdoll.Update, Swinger.Update, Targets.update, MoraleButton.Update, TexturedButtons.Update, ThinkOutLoud.Update, TurretRange.Display.Update, TurretRange.Setup.Distances.Update, Vis.Update, Fader.Update, War_RU.Update, XpStatus.Update, CurseProfiler.Update, ZonePOP.Update, followTheLeader.Update |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Categories |  | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
| Dependency | EASystem_Strings | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Categories |  | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
| Dependency | EASystem_Strings | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
