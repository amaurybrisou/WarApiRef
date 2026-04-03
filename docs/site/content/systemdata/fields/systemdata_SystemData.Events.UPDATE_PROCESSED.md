# SystemData.Events.UPDATE_PROCESSED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 30 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 188

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, Atlas, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, DPSMeter, DammazKron |
| Files seen in | AceAddon-3.0.lua, ClosetGoblin.lua, Code/Core/TargetInfoFix.lua, Core/ToolTip/DK_Tooltip.lua, DPSMeter.lua, DPSMeterWindow.lua, DuffTimer.lua, Fix/TargetInfoFix.lua |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, lua_call |
| Example locations | APPLY_TARGETINFO_FIX_DONOTTOUCH, AceAddon_OnUpdate_DONOTTOUCH, ActionBarsUpdated, ButtonsUpdated, Cleanup, Close |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 53 |
| Global usage count | 53 |
| Local definition count | 0 |
| Documentation references | 3 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

SystemData.SystemData.Events.UPDATE_PROCESSED field accessed by 30 addons; commonly found in APPLY_TARGETINFO_FIX_DONOTTOUCH and AceAddon_OnUpdate_DONOTTOUCH, ActionBarsUpdated, ButtonsUpdated, Cleanup, Close, CloseOptionswindow, DelayStart, Disable, Enable, Hide, Initialize, MB3FlagCarrier, MB3FlagDrop, OnChat, OnCombatFlagUpdated, OnHidden, OnInitialize, OnLoadingEnd, OnRegisteredTexturedChanged, OnSearchChanged, OnShown, OnShutdown, OnSkavenStatusChange, PO_Tick, QuitSync, RefreshActionBars, RefreshSettings, RegStart, Register, RegisterEventHandlers, RetAlert_Disable, RetAlert_Enable, SET_TARGETINFO_FIX_UPDATE_FLAG_DONOTTOUCH, SettingsUpdated, SetupAction, Show, Shutdown, StartBR, StartRKOTH, Statdoll.Local.delayedUpdate, Sync, SystemData.Events.UPDATE_PROCESSED, TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH, UC, UpdateButtons, UpdateSearch, _deferred_apply, delayedUpdate, event_page, event_registration, lua_call, queueUpdate contexts.

## Seen In

- Ace
- Atlas
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- DPSMeter
- DammazKron
- DammazKron, Squared, Statdoll Remix, Swift Assist
- DetauntHelper
- DuffTimer
- Enemy
- EveryBodyGuard
- GuardBot
- GuildWarden
- LoyalPet
- Obsidian
- PeaceOut
- Pure
- RetAlert
- SNT_BUTTONS
- Shinies
- Squared
- Statdoll Remix
- Swift Assist
- TexturedButtons
- TheSeeker
- TortallDPSCore
- Trakario
- scenarioInfo

## Related APIs

- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [WindowGetHandleInput](../../window_api/functions/window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Notes

- Observed in contexts: APPLY_TARGETINFO_FIX_DONOTTOUCH, AceAddon_OnUpdate_DONOTTOUCH, ActionBarsUpdated, ButtonsUpdated, Cleanup, Close
