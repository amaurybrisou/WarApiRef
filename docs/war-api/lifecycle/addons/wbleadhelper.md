# wbLeadHelper Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- wbLeadHelper.Settings

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | wbLeadHelper.OnInitialize | ambiguous | MEDIUM | AddonButtonTogglers.OnInitialize, ArmorGraphicToggle.OnInitialize, Assist.OnInitialize, AuctionTweaker.OnInitialize, AuraAddon.OnInitialize, AuraTooltip.OnInitialize, AutoMark.OnInitialize, AutoSalvage.OnInitialize, AutoSalvage.OptionsWindow.OnInitialize, BWMT.OnInitialize, BarText_Influence.OnInitialize, Biguns.OnInitialize, BloodyMess.OnInitialize, CCM.OnInitialize, ClosetGoblin.OnInitialize, ClosetGoblinCharacterWindow.OnInitialize, ClosetGoblinOptionWindow.OnInitialize, ClosetGoblinZoneWindow.OnInitialize, CMap_config.OnInitialize, Cheeseboard.OnInitialize, CB.OnInitialize, Compass3D.Camera.OnInitialize, Compass3D.OnInitialize, Countdown.OnInitialize, GG.OnInitialize, Crusher.OnInitialize, GG.OnInitialize, DeMoNiCore.OnInitialize, DuelInvite.OnInitialize, DuffTimer.OnInitialize, DwarfTalk.OnInitialize, EZCraft.OnInitialize, EZCraftX.OnInitialize, Emojii.OnInitialize, FastFriends.OnInitialize, FlagCap.OnInitialize, GCDTracker.OnInitialize, GetStats.OnInitialize, GroupIcons.OnInitialize, GroupIconsSG.OnInitialize, GroupRange.Styles.OnInitialize, HammerTime.OnInitialize, HealHoverAssist.OnInitialize, Hopper.OnInitialize, IHYTM.OnInitialize, InfoScroller.OnInitialize, Info_Alert.OnInitialize, Info_DeathBlow.OnInitialize, Info_Loot.OnInitialize, Info_Points.OnInitialize, KeyBar.OnInitialize, KeyBarSettings.OnInitialize, LibStub:NewLibrary.GlobalTable.OnInitialize, LootAlert.OnInitialize, MacroIcons.OnInitialize, MapPin.OnInitialize, Motion.OnInitialize, MouseHint.OnInitialize, MyReasons.OnInitialize, NervAlert.OnInitialize, NoOverheal.OnInitialize, OilTimer.OnInitialize, OverheadFonts.OnInitialize, PTLeader.OnInitialize, PlayEffectsOn.OnInitialize, Pure.OnInitialize, PureCareerBar.OnInitialize, QueueQueuer.OnInitialize, quickperf.OnInitialize, RandomMountUI.OnInitialize, RandomSayings.OnInitialize, ResHelp.OnInitialize, RoR_SoR.OnInitialize, Rotation.OnInitialize, RvRContribution.OnInitialize, ScenarioAlert.OnInitialize, Sequencer.OnInitialize, Shinies.OnInitialize, _G.Shinies.OnInitialize, SimpleCombatText.OnInitialize, Soloq.OnInitialize, SquaredClick.OnInitialize, HDI.OnInitialize, SHI.OnInitialize, Statdoll.onInitialize, Statdoll.Getstats.onInitialize, Statdoll.Local.onInitialize, Statdoll.Options.onInitialize, Statdoll.Stats.onInitialize, StatdollWnd.onInitialize, StatdollWndLight.onInitialize, Swinger.OnInitialize, ThankTheTank.OnInitialize, TTmodules.OnInitialize, Trakario.OnInitialize, WARRatingBuster.OnInitialize, warwhisperer.OnInitialize, wargames.OnInitialize, WCDB.OnInitialize, WCDP.OnInitialize, WindowMovers.OnInitialize, XpStatus.OnInitialize, XpStatusWindowNub.OnInitialize, minesweep.OnInitialize, rorAutoInviter.OnInitialize, wbLeadHelper.OnInitialize, whatsPugSc.OnInitialize |
| CreateWindow | wbLeadHelperConfigWindow | exact | HIGH | wbLeadHelperConfigWindow |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | wbLeadHelper.OnShutdown | ambiguous | MEDIUM | ActionPointWatch.OnShutdown, AddonButtonTogglers.OnShutdown, ArmorGraphicToggle.OnShutdown, Assist.OnShutdown, AuraAddon.OnShutdown, AutoSalvage.OnShutdown, BBars.OnShutdown, BWMT.OnShutdown, Biguns.OnShutdown, BloodyMess.OnShutdown, CCM.OnShutdown, ClosetGoblin.OnShutdown, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, CB.OnShutdown, Compass3D.Camera.OnShutdown, Compass3D.OnShutdown, DeMoNiCore.OnShutdown, DwarfTalk.OnShutdown, ObjectInspector.OnShutdown, EZGuard.OnShutdown, Emojii.OnShutdown, FastFriends.OnShutdown, FastInteract.OnShutdown, FlagCap.OnShutdown, GCDsaver.OnShutdown, GroupIcons.OnShutdown, GroupIconsSG.OnShutdown, Hopper.OnShutdown, IHYTM.OnShutdown, InfoScroller.OnShutdown, Info_Alert.OnShutdown, Info_DeathBlow.OnShutdown, InfoScroller.OnShutdown, KeyBar.OnShutdown, KillTracker.OnShutdown, LibGuard.OnShutdown, Megaphone.OnShutdown, MyReasons.OnShutdown, NoOverheal.OnShutdown, OilTimer.OnShutdown, PagerWentPoof.OnShutdown, PTLeader.OnShutdown, PartyCast.OnShutdown, QueueQueuer.OnShutdown, QuickNameActionsRessurected.OnShutdown, ResHelp.OnShutdown, RoR_SoR.OnShutdown, SquaredClick.OnShutdown, Statdoll.onShutdown, Statdoll.Getstats.onShutdown, Statdoll.Local.onShutdown, Statdoll.Stats.onShutdown, StatdollWnd.onShutdown, StatdollWndLight.onShutdown, TidyChat.OnShutdown, TomeTracker.OnShutdown, Vectors.local.onshutdown, Vectors.onshutdown, warwhisperer.OnShutdown, WarTriage.OnShutdown, wargames.OnShutdown, rorAutoInviter.OnShutdown, wbLeadHelper.OnShutdown, whatsPugSc.OnShutdown |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | wbLeadHelper.timerUpdate | ambiguous | MEDIUM | OilTimer.timerUpdate, wbLeadHelper.timerUpdate |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | LibSlash | unresolved |  | — |
| Categories |  | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | LibSlash | unresolved |  | — |
| Categories |  | unresolved |  | — |
