# GameData.InteractData.GetLootList

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | ZCurse_Profiler |
| Files seen in | `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:480`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:602` |
| Namespaces detected | GameData |
| Source kinds | globals, lua_calls |
| Example locations | ZCurse_Profiler: CurseProfiler.fornil, ZCurse_Profiler: CurseProfiler.whilethen |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
| Consistent arguments | yes |
| Consistent returns | yes |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
GameData.InteractData.GetLootList()
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- ZCurse_Profiler

## Examples

- ZCurse_Profiler: CurseProfiler.fornil -> GameData.InteractData.GetLootList()
- ZCurse_Profiler: CurseProfiler.whilethen -> GameData.InteractData.GetLootList()

## Related APIs

- none

## Used With

- none

## Triggered By

- [SystemData.Events.INTERACT_SHOW_LOOT](../../events/game_events/game_event_SystemData.Events.INTERACT_SHOW_LOOT.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_SHOW_PQ_LOOT](../../events/game_events/game_event_SystemData.Events.INTERACT_SHOW_PQ_LOOT.md) (HIGH 100/100) - Game Event

## Affects

- [GameData.InteractData.GetLootList](../../gamedata/fields/gamedata_GameData.InteractData.GetLootList.md) (HIGH 100/100) - GameData Field

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
