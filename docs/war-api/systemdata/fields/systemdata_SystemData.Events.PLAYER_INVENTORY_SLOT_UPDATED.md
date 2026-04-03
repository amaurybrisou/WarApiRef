# SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

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
| Addons seen in | CM_ClosetGoblin, CM_ClosetGoblin, Enemy, Pocket Palette, PotionBar, Shinies, Enemy, Pocket Palette, PotionBar |
| Files seen in | `/workspace/data/raw/ClosetGoblin/ClosetGoblin.lua:122`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.lua:73`, `/workspace/data/raw/Enemy/Code/TalismanAlerter/TalismanAlerter.lua:37`, `/workspace/data/raw/PocketPalette/PocketPalette.lua:105`, `/workspace/data/raw/PotionBar/source/Main.lua:189`, `/workspace/data/raw/PotionBar/source/Main.lua:292` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | ClosetGoblin.OnInitialize, ClosetGoblin.OnInventorySlotUpdated, ClosetGoblin.OnShutdown, Enemy.TalismanAlerter_Update, Enemy._TalismanAlerterEnabledChanged, PP.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 16 |
| Global usage count | 16 |
| Local definition count | 0 |
| Documentation references | 4 |
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

Observed SystemData field used by 5 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- CM_ClosetGoblin
- CM_ClosetGoblin, Enemy, Pocket Palette, PotionBar, Shinies
- Enemy
- Pocket Palette
- PotionBar

## Related APIs

- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowGetOffsetFromParent](../../window_api/functions/window_WindowGetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: ClosetGoblin.OnInitialize, ClosetGoblin.OnInventorySlotUpdated, ClosetGoblin.OnShutdown, Enemy.TalismanAlerter_Update, Enemy._TalismanAlerterEnabledChanged, PP.Initialize
