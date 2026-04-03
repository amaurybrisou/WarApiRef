# SystemData.Events.UPDATE_PROCESSED

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy, LibWBToggler, PartyCast |
| Files seen in | `/workspace/data/raw/BuffHead/Setup/ContainerDemo.lua:197`, `/workspace/data/raw/BuffHead/Setup/LayoutDemo.lua:147`, `/workspace/data/raw/BuffHead/Setup/SetupEffectCache.lua:386`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.lua:73`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettings.lua:148`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettings2ndTier.lua:104`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettings2ndTier.lua:1253`, `/workspace/data/raw/TexturedButtons/TexturedButtons.lua:408` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | BuffHead: BuffHead.Setup.ContainerDemo.Enable, BuffHead: BuffHead.Setup.Demo.Enable, BuffHead: BuffHead.Setup.EffectCache.OnSearchChanged, CM_ClosetGoblin: ClosetGoblin.OnInitialize, DAoCBuff: DAoCBuffSettings.CloseOptionswindow, DAoCBuff: FilterSettings.Close |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 14 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 7 |
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

Observed as a shared SystemData runtime event used by 12 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- Ace
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- LibWBToggler
- PartyCast
- Shinies
- Swift Assist
- TexturedButtons
- WoH-Reticle

## Registrars And Handlers

- AceAddon_OnUpdate_DONOTTOUCH
- AceTimerOnUpdateDONOTTOUCH
- BuffHead.Setup.ContainerDemo.OnUpdate
- BuffHead.Setup.Demo.OnUpdate
- BuffHead.Setup.EffectCache.UpdateSearch
- ClosetGoblin.OnActivationWatchdog
- DAoCBuffSettings.FilterSettings.Cleanup
- DAoCBuffSettings.ImExport.Cleanup
- DAoCBuffSettings.UC
- LibConfig.OnUpdate
- RegisterEventHandler
- TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH
- TexturedButtons.OnRegisteredTexturedChanged
- TexturedButtons.RefreshActionBars
- TexturedButtons.RefreshSettings
- TexturedButtons.UpdateButtons
- global

## Examples

- BuffHead: BuffHead.Setup.ContainerDemo.Enable -> SystemData.Events.UPDATE_PROCESSED -> BuffHead.Setup.ContainerDemo.OnUpdate
- BuffHead: BuffHead.Setup.Demo.Enable -> SystemData.Events.UPDATE_PROCESSED -> BuffHead.Setup.Demo.OnUpdate
- BuffHead: BuffHead.Setup.EffectCache.OnSearchChanged -> SystemData.Events.UPDATE_PROCESSED -> BuffHead.Setup.EffectCache.UpdateSearch
- CM_ClosetGoblin: ClosetGoblin.OnInitialize -> SystemData.Events.UPDATE_PROCESSED -> ClosetGoblin.OnActivationWatchdog
- DAoCBuff: DAoCBuffSettings.CloseOptionswindow -> SystemData.Events.UPDATE_PROCESSED -> DAoCBuffSettings.UC
- DAoCBuff: FilterSettings.Close -> SystemData.Events.UPDATE_PROCESSED -> DAoCBuffSettings.FilterSettings.Cleanup

## Related APIs

- [WindowGetAlpha](../../window_api/functions/window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetHandleInput](../../window_api/functions/window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (HIGH 93/100) - Global Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
