# CraftingSystem

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 125

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AnywhereTrainer, AnywhereTrainerAdditions, EZCraft, EZCraftX, JunkDump, Motion, nLootLink |
| Files seen in | Source/EZCraftX.lua, Source/Motion.lua, source/AnywhereTrainer.lua, source/nLootLinkData.lua, source/nLootLinkGUI.lua |
| Namespaces detected | CraftingSystem |
| Source kinds | lua_calls |
| Example locations | AnywhereTrainer: OnLeftClickCraft, AnywhereTrainerAdditions: OnLeftClickCraft, EZCraft: ItemIsAllowedInSlot, EZCraftX: item_allowed_in_slot, JunkDump: professionChecks, JunkDump: professionChecks2 |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 15 |
| Global usage count | 6 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
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

Shared function table with 6 member functions; the primary API surface for 7 addons.

## Functions

- CraftingSystem.Clear
- CraftingSystem.GetCraftingData
- CraftingSystem.GetCraftingDataStrings
- CraftingSystem.GetResourceTypeName
- CraftingSystem.IsCraftingItem
- CraftingSystem.ToggleShowing

## Observed Members

- none

## Seen In

- AnywhereTrainer
- AnywhereTrainerAdditions
- EZCraft
- EZCraftX
- JunkDump
- Motion
- nLootLink

## Examples

- AnywhereTrainer: OnLeftClickCraft -> CraftingSystem.ToggleShowing(GameData.TradeSkills.APOTHECARY)
- AnywhereTrainer: OnLeftClickCraft -> CraftingSystem.ToggleShowing(GameData.TradeSkills.TALISMAN)
- AnywhereTrainerAdditions: OnLeftClickCraft -> CraftingSystem.ToggleShowing(GameData.TradeSkills.APOTHECARY)
- AnywhereTrainerAdditions: OnLeftClickCraft -> CraftingSystem.ToggleShowing(GameData.TradeSkills.TALISMAN)
- EZCraft: ItemIsAllowedInSlot -> CraftingSystem.GetCraftingData(itemData)
- EZCraftX: item_allowed_in_slot -> CraftingSystem.GetCraftingData(itemData)

## Notes

- none
