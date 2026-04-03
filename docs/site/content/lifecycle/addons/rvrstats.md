# RvRStats Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | RvRStatsRvRTab.Init | ambiguous | MEDIUM | ArsenalRank.init, AuctionAssist.Init, AutoBand.init, AutoFocus.Init, BagOMatic.init, BankArkel.Init, BuddyBind.init, Calling.Init, ChattyCathy.Init, DAoCTooltips.Init, tactic.init, EffigyBar.Init, FightFinder.init, GuardLine.init, ItemRack.Backpack.Init, LibRuString.Init, LibGuard.Init, LibSkillicon.Init, mech.init, MiniMapMonster.Init, MoraleBG.Init, MoraleCircle.init, Nisp.Init, lnHandler.Init, lnConfig.Init, PartyCast.Init, PeaceOut.Init, RvRStats.Init, RvRStatsRvRTab.Init, RvRStats.Init, RvRStatsRvRTab.Init, ShowHealth.Init, Targets.init, TTitan.Init, TTitan.UI.Init, WARCommanderConfig.Init, WARCommander.Init, ZonePOP.init, nLootLinkData.init, nLootLinkGUIIntegration.init, nLootLinkOptions.init, nLootLinkChat.init, nLootLink.init, nLootLinkGatherer.init, nLootLinkGUI.init, nLootLinkComm.init, nRarity.init |
| CallFunction | RvRStats.Init | ambiguous | MEDIUM | ArsenalRank.init, AuctionAssist.Init, AutoBand.init, AutoFocus.Init, BagOMatic.init, BankArkel.Init, BuddyBind.init, Calling.Init, ChattyCathy.Init, DAoCTooltips.Init, tactic.init, EffigyBar.Init, FightFinder.init, GuardLine.init, ItemRack.Backpack.Init, LibRuString.Init, LibGuard.Init, LibSkillicon.Init, mech.init, MiniMapMonster.Init, MoraleBG.Init, MoraleCircle.init, Nisp.Init, lnHandler.Init, lnConfig.Init, PartyCast.Init, PeaceOut.Init, RvRStats.Init, RvRStatsRvRTab.Init, RvRStats.Init, RvRStatsRvRTab.Init, ShowHealth.Init, Targets.init, TTitan.Init, TTitan.UI.Init, WARCommanderConfig.Init, WARCommander.Init, ZonePOP.init, nLootLinkData.init, nLootLinkGUIIntegration.init, nLootLinkOptions.init, nLootLinkChat.init, nLootLink.init, nLootLinkGatherer.init, nLootLinkGUI.init, nLootLinkComm.init, nRarity.init |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | RvRStats.OnShutDown | ambiguous | MEDIUM | ActionPointWatch.OnShutdown, AddonButtonTogglers.OnShutdown, ArmorGraphicToggle.OnShutdown, AuraAddon.OnShutdown, AutoSalvage.OnShutdown, BBars.OnShutdown, BWMT.OnShutdown, Biguns.OnShutdown, BloodyMess.OnShutdown, CCM.OnShutdown, ClosetGoblin.OnShutdown, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, CB.OnShutdown, Compass3D.Camera.OnShutdown, Compass3D.OnShutdown, DeMoNiCore.OnShutdown, DwarfTalk.OnShutdown, ObjectInspector.OnShutdown, EZGuard.OnShutdown, Emojii.OnShutdown, FastFriends.OnShutdown, FastInteract.OnShutdown, FlagCap.OnShutdown, GCDsaver.OnShutdown, GroupIcons.OnShutdown, GroupIconsSG.OnShutdown, Hopper.OnShutdown, IHYTM.OnShutdown, InfoScroller.OnShutdown, Info_Alert.OnShutdown, Info_DeathBlow.OnShutdown, InfoScroller.OnShutdown, KeyBar.OnShutdown, KillTracker.OnShutdown, LibGuard.OnShutdown, Megaphone.OnShutdown, NoOverheal.OnShutdown, OilTimer.OnShutdown, PagerWentPoof.OnShutdown, PartyCast.OnShutdown, QueueQueuer.OnShutdown, QuickNameActionsRessurected.OnShutdown, ResHelp.OnShutdown, RoR_SoR.OnShutdown, SquaredClick.OnShutdown, Statdoll.onShutdown, Statdoll.Stats.onShutdown, Statdoll.Local.onShutdown, Statdoll.Getstats.onShutdown, StatdollWnd.onShutdown, StatdollWndLight.onShutdown, TidyChat.OnShutdown, TomeTracker.OnShutdown, Vectors.local.onshutdown, Vectors.onshutdown, WarTriage.OnShutdown, wargames.OnShutdown, rorAutoInviter.OnShutdown, wbLeadHelper.OnShutdown, whatsPugSc.OnShutdown |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EA_CharacterWindow | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EA_CharacterWindow | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
