# SystemData.Events.ENTER_WORLD

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 27 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 168

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | ActionFraction, AutoBand, BBars - Mechanic Only, BarText (Influence), CDown, CaVES, ClickSoundSuppressor, CoolDownLine |
| Files seen in | AutoBand.lua, BBars.lua, BarText_Influence.lua, CDown.lua, ClickSoundSuppressor.lua, CoolDownLine.lua, Core/DK_Core.lua, Ding.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | Init, Initialize, OnInit, OnInitialize, OnInterfaceLoaded, OnLoad |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 19 |
| Global usage count | 19 |
| Local definition count | 0 |
| Documentation references | 1 |
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

SystemData.SystemData.Events.ENTER_WORLD field accessed by 27 addons; commonly found in Init and Initialize, OnInit, OnInitialize, OnInterfaceLoaded, OnLoad, OnShutdown, ProcessAlert, RegisterEvents, RegisterSelfEvents, Shutdown, Start, Stop, UnregisterEvents, init, lua_call, newEA_Window_LoadingScreenOnLoadBegin, onInit, shutdown contexts.

## Seen In

- ActionFraction
- AutoBand
- BBars - Mechanic Only
- BarText (Influence)
- CDown
- CaVES
- ClickSoundSuppressor
- CoolDownLine
- DAoCBuff
- DammazKron
- DetauntHelper
- Ding
- DuffTimer
- EA_LoadingScreen
- GCDsaver
- GroupSpotter
- Mech
- PartyCast
- PetFixWindow
- Rangechecker
- ReliquaryHunter
- ResHelp
- RoR_SoR
- Sequencer
- TomeTracker
- VPBreakdown
- scnoload

## Related APIs

- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [LabelSetFont](../../window_api/functions/window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [Start](../../events/game_events/game_event_Start.md) (MEDIUM 43/100) - Game Event
- [Stop](../../events/game_events/game_event_Stop.md) (MEDIUM 43/100) - Game Event

## Used With

- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.ChatLogFilters.SHOUT](systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [Start](../../events/game_events/game_event_Start.md) (MEDIUM 43/100) - Game Event

## Notes

- Observed in contexts: Init, Initialize, OnInit, OnInitialize, OnInterfaceLoaded, OnLoad
