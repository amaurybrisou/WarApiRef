# SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED

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
| Addons seen in | CM_ClosetGoblin, Enemy, Pocket Palette, PotionBar, Shinies |
| Files seen in | `/workspace/data/raw/ClosetGoblin/ClosetGoblin.lua:73`, `/workspace/data/raw/Enemy/Code/TalismanAlerter/TalismanAlerter.lua:37`, `/workspace/data/raw/PocketPalette/PocketPalette.lua:105`, `/workspace/data/raw/PotionBar/source/Main.lua:189` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | CM_ClosetGoblin: ClosetGoblin.OnInitialize, Enemy: Enemy._TalismanAlerterEnabledChanged, Pocket Palette: PP.Initialize, PotionBar: PotionBar.Initialize |
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

Observed as a shared SystemData runtime event used by 5 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- CM_ClosetGoblin
- Enemy
- Pocket Palette
- PotionBar
- Shinies

## Registrars And Handlers

- ClosetGoblin.OnInventorySlotUpdated
- Enemy.TalismanAlerter_Update
- PP.UpdateItemSlots
- PotionBar.InventorySlotUpdated
- RegisterEventHandler
- ShiniesAPI.Core_OnPlayerInventorySlotUpdated
- ShiniesDataInventory.OnPlayerInventorySlotUpdated
- ShiniesPostUI.OnPlayerInventorySlotUpdated
- global

## Examples

- CM_ClosetGoblin: ClosetGoblin.OnInitialize -> SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED -> ClosetGoblin.OnInventorySlotUpdated
- Enemy: Enemy._TalismanAlerterEnabledChanged -> SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED -> Enemy.TalismanAlerter_Update
- Pocket Palette: PP.Initialize -> SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED -> PP.UpdateItemSlots
- PotionBar: PotionBar.Initialize -> SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED -> PotionBar.InventorySlotUpdated
- CM_ClosetGoblin: ClosetGoblin.OnInventorySlotUpdated -> RegisterEventHandler(SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED, ClosetGoblin.OnInventorySlotUpdated)
- Enemy: Enemy.TalismanAlerter_Update -> RegisterEventHandler(SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED, Enemy.TalismanAlerter_Update)

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
