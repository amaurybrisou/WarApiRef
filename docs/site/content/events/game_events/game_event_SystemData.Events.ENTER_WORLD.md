# SystemData.Events.ENTER_WORLD

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
| Addons seen in | DAoCBuff, PartyCast, RoR_SoR |
| Files seen in | `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:25`, `/workspace/data/raw/PartyCast/PartyCast.lua:51`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:178` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | DAoCBuff: DAoCBuff.Initialize, PartyCast: PartyCast.Init, RoR_SoR: RoR_SoR:RegisterSelfEvents |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
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

- DAoCBuff
- PartyCast
- RoR_SoR

## Registrars And Handlers

- DAoCBuff.OnReload
- DAoCBuff.TargetsCleared
- PartyCast.GROUP_UPDATED
- RegisterEventHandler
- RoR_SoR.Enable
- global

## Examples

- DAoCBuff: DAoCBuff.Initialize -> SystemData.Events.ENTER_WORLD -> DAoCBuff.OnReload
- DAoCBuff: DAoCBuff.Initialize -> SystemData.Events.ENTER_WORLD -> DAoCBuff.TargetsCleared
- PartyCast: PartyCast.Init -> SystemData.Events.ENTER_WORLD -> PartyCast.GROUP_UPDATED
- RoR_SoR: RoR_SoR:RegisterSelfEvents -> SystemData.Events.ENTER_WORLD -> RoR_SoR.Enable
- DAoCBuff: DAoCBuff.OnReload -> RegisterEventHandler(SystemData.Events.ENTER_WORLD, DAoCBuff.OnReload)
- DAoCBuff: DAoCBuff.TargetsCleared -> RegisterEventHandler(SystemData.Events.ENTER_WORLD, DAoCBuff.TargetsCleared)

## Related APIs

- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 100/100) - Global Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
