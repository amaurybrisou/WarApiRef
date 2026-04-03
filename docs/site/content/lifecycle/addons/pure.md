# Pure Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- PureDB

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | Pure.OnInitialize | ambiguous | MEDIUM | AddonButtonTogglers.OnInitialize, ArmorGraphicToggle.OnInitialize, AuctionTweaker.OnInitialize, AuraTooltip.OnInitialize, AuraAddon.OnInitialize, AutoMark.OnInitialize, AutoSalvage.OnInitialize, AutoSalvage.OptionsWindow.OnInitialize, BWMT.OnInitialize, Biguns.OnInitialize, BloodyMess.OnInitialize, CCM.OnInitialize, ClosetGoblin.OnInitialize, ClosetGoblinCharacterWindow.OnInitialize, ClosetGoblinZoneWindow.OnInitialize, ClosetGoblinOptionWindow.OnInitialize, CB.OnInitialize, Compass3D.Camera.OnInitialize, Compass3D.OnInitialize, Countdown.OnInitialize, GG.OnInitialize, Crusher.OnInitialize, GG.OnInitialize, DeMoNiCore.OnInitialize, DuelInvite.OnInitialize, DuffTimer.OnInitialize, DwarfTalk.OnInitialize, EZCraftX.OnInitialize, Emojii.OnInitialize, FastFriends.OnInitialize, FlagCap.OnInitialize, GCDTracker.OnInitialize, GetStats.OnInitialize, GroupIcons.OnInitialize, GroupIconsSG.OnInitialize, GroupRange.Styles.OnInitialize, HammerTime.OnInitialize, HealHoverAssist.OnInitialize, Hopper.OnInitialize, IHYTM.OnInitialize, InfoScroller.OnInitialize, Info_Alert.OnInitialize, Info_DeathBlow.OnInitialize, Info_Loot.OnInitialize, Info_Points.OnInitialize, KeyBar.OnInitialize, KeyBarSettings.OnInitialize, LibStub:NewLibrary.GlobalTable.OnInitialize, LootAlert.OnInitialize, MacroIcons.OnInitialize, MapPin.OnInitialize, Motion.OnInitialize, MouseHint.OnInitialize, NervAlert.OnInitialize, NoOverheal.OnInitialize, OilTimer.OnInitialize, OverheadFonts.OnInitialize, PlayEffectsOn.OnInitialize, Pure.OnInitialize, PureCareerBar.OnInitialize, QueueQueuer.OnInitialize, quickperf.OnInitialize, RandomMountUI.OnInitialize, RandomSayings.OnInitialize, ResHelp.OnInitialize, RoR_SoR.OnInitialize, Rotation.OnInitialize, RvRContribution.OnInitialize, ScenarioAlert.OnInitialize, Sequencer.OnInitialize, Shinies.OnInitialize, _G.Shinies.OnInitialize, SimpleCombatText.OnInitialize, Soloq.OnInitialize, SquaredClick.OnInitialize, HDI.OnInitialize, SHI.OnInitialize, Statdoll.onInitialize, Statdoll.Stats.onInitialize, Statdoll.Local.onInitialize, Statdoll.Getstats.onInitialize, Statdoll.Options.onInitialize, StatdollWnd.onInitialize, StatdollWndLight.onInitialize, Swinger.OnInitialize, ThankTheTank.OnInitialize, TTmodules.OnInitialize, Trakario.OnInitialize, WARRatingBuster.OnInitialize, wargames.OnInitialize, WCDB.OnInitialize, WCDP.OnInitialize, WindowMovers.OnInitialize, XpStatus.OnInitialize, XpStatusWindowNub.OnInitialize, minesweep.OnInitialize, rorAutoInviter.OnInitialize, wbLeadHelper.OnInitialize, whatsPugSc.OnInitialize |
| CreateWindow | PureConfig | exact | HIGH | PureConfig |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | AceDBLogoutHandler | ambiguous | MEDIUM | Ace.AceDBLogoutHandler, Crusher.AceDBLogoutHandler, Hopper.AceDBLogoutHandler, Pure.AceDBLogoutHandler, Pure Careerbar.AceDBLogoutHandler, Shinies.AceDBLogoutHandler |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | PureConfig_Player_OnUpdate | exact | HIGH | Pure.PureConfig_Player_OnUpdate |
| CallFunction | PureConfig_PlayerPet_OnUpdate | exact | HIGH | Pure.PureConfig_PlayerPet_OnUpdate |
| CallFunction | PureConfig_PlayerPetTarget_OnUpdate | exact | HIGH | Pure.PureConfig_PlayerPetTarget_OnUpdate |
| CallFunction | PureConfig_Group_OnUpdate | exact | HIGH | Pure.PureConfig_Group_OnUpdate |
| CallFunction | PureConfig_GroupPet_OnUpdate | exact | HIGH | Pure.PureConfig_GroupPet_OnUpdate |
| CallFunction | PureConfig_TargetFriendly_OnUpdate | exact | HIGH | Pure.PureConfig_TargetFriendly_OnUpdate |
| CallFunction | PureConfig_TargetFriendlyHUD_OnUpdate | exact | HIGH | Pure.PureConfig_TargetFriendlyHUD_OnUpdate |
| CallFunction | PureConfig_TargetHostile_OnUpdate | exact | HIGH | Pure.PureConfig_TargetHostile_OnUpdate |
| CallFunction | PureConfig_TargetHostileHUD_OnUpdate | exact | HIGH | Pure.PureConfig_TargetHostileHUD_OnUpdate |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_PlayerMenu | unresolved |  | — |
| Dependency | EASystem_Strings | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_TargetInfo | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EATemplate_DefaultWindowSkin | unresolved |  | — |
| Dependency | EATemplate_UnitFrames | unresolved |  | — |
| Dependency | SharedAssetsPure | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_PlayerMenu | unresolved |  | — |
| Dependency | EASystem_Strings | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_TargetInfo | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EATemplate_DefaultWindowSkin | unresolved |  | — |
| Dependency | EATemplate_UnitFrames | unresolved |  | — |
| Dependency | SharedAssetsPure | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
