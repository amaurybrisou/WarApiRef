# NaturalLog Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- NaturalLog.BasicVars
- NaturalLog.LogOptions

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | NaturalLog.OnLoad | ambiguous | MEDIUM | ActionBarHide.OnLoad, AuraAddon.OnLoad, AuraShares.OnLoad, AuraTexture.OnLoad, CleanTargetWindow.OnLoad, CrusherConfig.OnLoad, Crusher.OnLoad, DammazKron.OnLoad, Effigy.OnLoad, HammerTime.OnLoad, PhantomLib.OnLoad, HopperConfig.OnLoad, Hopper.OnLoad, LibRuString.OnLoad, LibRange.OnLoad, MotionConfig.OnLoad, Motion.OnLoad, NaturalLog.OnLoad, lnConfig.OnLoad, NBSBCore.OnLoad, Phantom.OnLoad, PhantomLib.OnLoad, PureConfig.OnLoad, Pure.OnLoad, PureCareerBar.OnLoad, QuickNameActionsRessurected.OnLoad, Squared.OnLoad, SquaredDisabled.OnLoad, SquaredGroup.OnLoad, SquaredPlayer.OnLoad, SquaredScenario.OnLoad, SquaredTestmode.OnLoad, SquaredWarband.OnLoad, Statdoll.onLoad, TidyChat.OnLoad, TidyQueue.OnLoad, TidyRoll.OnLoad, TokenMachine.OnLoad, TTmodules.OnLoad, Twister.OnLoad, TwisterSet.OnLoad, WCDBConfig.OnLoad, WCDB.OnLoad, WCDPConfig.OnLoad, WCDP.OnLoad, fpsbox.OnLoad, hideInf.OnLoad, scnoload.OnLoad, wbLeadHelperConfigTab.OnLoad, wbLeadHelperMessagesTab.OnLoad |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | lnConfig.EventUpdate | exact | HIGH | lnConfig.EventUpdate |
| CallFunction | lnConfig.ScrollUpdate | ambiguous | MEDIUM | NaturalLog.ScrollUpdate, lnConfig.ScrollUpdate |
| CallFunction | lnConfig.GeneralUpdate | exact | HIGH | lnConfig.GeneralUpdate |
