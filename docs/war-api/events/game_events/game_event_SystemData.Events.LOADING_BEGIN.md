# SystemData.Events.LOADING_BEGIN

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 186

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
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
| Addons seen in | CM_ClosetGoblin, MapMonster, wbLeadHelper |
| Files seen in | `/workspace/ClosetGoblin/ClosetGoblin.lua:1218`, `/workspace/MapMonster/Source/MapMonster_Player.lua:72`, `/workspace/wbLeadHelper/wbLeadHelper.lua:34` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | CM_ClosetGoblin: ClosetGoblin.ZoneChangeInit, MapMonster: MapMonster.InitializePlayer, wbLeadHelper: wbLeadHelper.OnInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 3 |
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

Observed as a shared SystemData runtime event used by 3 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- CM_ClosetGoblin
- MapMonster
- wbLeadHelper

## Registrars And Handlers

- ClosetGoblin.LoadingBegin
- MapMonster.OnLoadingBegin
- RegisterEventHandler
- global
- wbLeadHelper.LOADING_START

## Examples

- CM_ClosetGoblin: ClosetGoblin.ZoneChangeInit -> SystemData.Events.LOADING_BEGIN -> ClosetGoblin.LoadingBegin
- MapMonster: MapMonster.InitializePlayer -> SystemData.Events.LOADING_BEGIN -> MapMonster.OnLoadingBegin
- wbLeadHelper: wbLeadHelper.OnInitialize -> SystemData.Events.LOADING_BEGIN -> wbLeadHelper.LOADING_START
- CM_ClosetGoblin: ClosetGoblin.LoadingBegin -> RegisterEventHandler(SystemData.Events.LOADING_BEGIN, ClosetGoblin.LoadingBegin)
- MapMonster: MapMonster.OnLoadingBegin -> RegisterEventHandler(SystemData.Events.LOADING_BEGIN, MapMonster.OnLoadingBegin)
- wbLeadHelper: wbLeadHelper.LOADING_START -> RegisterEventHandler(SystemData.Events.LOADING_BEGIN, wbLeadHelper.LOADING_START)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
