# BBars - Mechanic Only Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | BBars.OnInit | ambiguous | MEDIUM | BBars.OnInit, CNC.OnInit, ClickSoundSuppressor.OnInit, PetFixWindow.onInit |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | BBars.OnShutdown | ambiguous | MEDIUM | ActionPointWatch.OnShutdown, AddonButtonTogglers.OnShutdown, ArmorGraphicToggle.OnShutdown, AuraAddon.OnShutdown, AutoSalvage.OnShutdown, BBars.OnShutdown, BWMT.OnShutdown, Biguns.OnShutdown, BloodyMess.OnShutdown, CCM.OnShutdown, ClosetGoblin.OnShutdown, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, CB.OnShutdown, Compass3D.Camera.OnShutdown, Compass3D.OnShutdown, DeMoNiCore.OnShutdown, DwarfTalk.OnShutdown, ObjectInspector.OnShutdown, EZGuard.OnShutdown, Emojii.OnShutdown, FastFriends.OnShutdown, FastInteract.OnShutdown, FlagCap.OnShutdown, GCDsaver.OnShutdown, GroupIcons.OnShutdown, GroupIconsSG.OnShutdown, Hopper.OnShutdown, IHYTM.OnShutdown, InfoScroller.OnShutdown, Info_Alert.OnShutdown, Info_DeathBlow.OnShutdown, InfoScroller.OnShutdown, KeyBar.OnShutdown, KillTracker.OnShutdown, LibGuard.OnShutdown, Megaphone.OnShutdown, NoOverheal.OnShutdown, OilTimer.OnShutdown, PagerWentPoof.OnShutdown, PartyCast.OnShutdown, QueueQueuer.OnShutdown, QuickNameActionsRessurected.OnShutdown, ResHelp.OnShutdown, RoR_SoR.OnShutdown, SquaredClick.OnShutdown, Statdoll.onShutdown, Statdoll.Stats.onShutdown, Statdoll.Local.onShutdown, Statdoll.Getstats.onShutdown, StatdollWnd.onShutdown, StatdollWndLight.onShutdown, TidyChat.OnShutdown, TomeTracker.OnShutdown, Vectors.local.onshutdown, Vectors.onshutdown, WarTriage.OnShutdown, wargames.OnShutdown, rorAutoInviter.OnShutdown, wbLeadHelper.OnShutdown, whatsPugSc.OnShutdown |
