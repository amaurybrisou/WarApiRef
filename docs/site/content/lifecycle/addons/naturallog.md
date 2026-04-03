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
| CallFunction | NaturalLog.OnLoad | ambiguous | MEDIUM | ActionBarHide.OnLoad, AuraShares.OnLoad, AuraTexture.OnLoad, AuraAddon.OnLoad, CleanTargetWindow.OnLoad, CrusherConfig.OnLoad, Crusher.OnLoad, DammazKron.OnLoad, Effigy.OnLoad, HammerTime.OnLoad, PhantomLib.OnLoad, HopperConfig.OnLoad, Hopper.OnLoad, LibRuString.OnLoad, LibRange.OnLoad, MotionConfig.OnLoad, Motion.OnLoad, NaturalLog.OnLoad, lnConfig.OnLoad, NBSBCore.OnLoad, PhantomLib.OnLoad, Phantom.OnLoad, PureConfig.OnLoad, Pure.OnLoad, PureCareerBar.OnLoad, QuickNameActionsRessurected.OnLoad, Squared.OnLoad, SquaredDisabled.OnLoad, SquaredPlayer.OnLoad, SquaredGroup.OnLoad, SquaredScenario.OnLoad, SquaredWarband.OnLoad, SquaredTestmode.OnLoad, Statdoll.onLoad, TidyChat.OnLoad, TidyQueue.OnLoad, TidyRoll.OnLoad, TokenMachine.OnLoad, TTmodules.OnLoad, Twister.OnLoad, TwisterSet.OnLoad, WCDBConfig.OnLoad, WCDB.OnLoad, WCDPConfig.OnLoad, WCDP.OnLoad, hideInf.OnLoad, scnoload.OnLoad, wbLeadHelperMessagesTab.OnLoad, wbLeadHelperConfigTab.OnLoad |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | lnConfig.EventUpdate | exact | HIGH | lnConfig.EventUpdate |
| CallFunction | lnConfig.ScrollUpdate | ambiguous | MEDIUM | NaturalLog.ScrollUpdate, lnConfig.ScrollUpdate |
| CallFunction | lnConfig.GeneralUpdate | exact | HIGH | lnConfig.GeneralUpdate |
