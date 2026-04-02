# SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 153

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | PotionBar |
| Files seen in | `/workspace_addons/PotionBar/source/Main.lua:189`, `/workspace_addons/PotionBar/source/Main.lua:292` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | PotionBar.Initialize, PotionBar.OnPlayerBattleLevelUpdated, PotionBar.Shutdown, SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED, event_page, event_registration |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
| Local definition count | 0 |
| Documentation references | 4 |
| Initialization flow references | 1 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | no |
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

Observed SystemData field used by 1 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- PotionBar

## Related APIs

- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowGetOffsetFromParent](../../window_api/functions/window_WindowGetOffsetFromParent.md) (HIGH 100/100) - Window Function

## Used With

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED](systemdata_SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_RANK_UPDATED](systemdata_SystemData.Events.PLAYER_RENOWN_RANK_UPDATED.md) (HIGH 100/100) - SystemData Field
- [WindowGetOffsetFromParent](../../window_api/functions/window_WindowGetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: PotionBar.Initialize, PotionBar.OnPlayerBattleLevelUpdated, PotionBar.Shutdown, SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED, event_page, event_registration
