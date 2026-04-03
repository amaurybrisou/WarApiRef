# CharacterWindow

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
| Addons seen in | AnywhereTrainer, BankWindowFix, CM_ClosetGoblin, CaVES, Effigy, RvRStats, RvRStatsTab, nRarity |
| Files seen in | Source/BankWindowFix.lua, Source/RvRStatsRvRTab.lua, States/EffigyStatePlayer.lua, source/AnywhereTrainer.lua, source/CaVES.lua, source/containers/Character.lua |
| Namespaces detected | CharacterWindow |
| Source kinds | lua_calls |
| Example locations | AnywhereTrainer: CharacterWindow_UpdateMode, BankWindowFix: BagEquipmentRButtonUp, CM_ClosetGoblin: ShowShowCloakHeraldryOnly, CM_ClosetGoblin: ShowShowCloakOnly, CM_ClosetGoblin: ShowShowHelmOnly, CaVES: GetTotalBonusPower |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 39 |
| Global usage count | 10 |
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

Shared function table with 10 member functions; the primary API surface for 8 addons.

## Functions

- CharacterWindow.AutoEquipItem
- CharacterWindow.CalcArmorPenetration
- CharacterWindow.CalculateValueWithBonus
- CharacterWindow.CreateTooltip
- CharacterWindow.GetItem
- CharacterWindow.IsResistDiminished
- CharacterWindow.IsStatDiminished
- CharacterWindow.UpdateDefenseLabels
- CharacterWindow.UpdateMode
- CharacterWindow.UpdateStatsNew

## Observed Members

- none

## Seen In

- AnywhereTrainer
- BankWindowFix
- CM_ClosetGoblin
- CaVES
- Effigy
- RvRStats
- RvRStatsTab
- nRarity

## Examples

- AnywhereTrainer: CharacterWindow_UpdateMode -> CharacterWindow.UpdateMode(mode)
- BankWindowFix: BagEquipmentRButtonUp -> CharacterWindow.AutoEquipItem(slot)
- CM_ClosetGoblin: ShowShowCloakHeraldryOnly -> CharacterWindow.CreateTooltip("ClosetGoblinCharacterWindowContentsEquipmentShowCloakHeraldry", GetString(StringTables.Default.TOOLTIP_SHOW_HERALDRY))
- CM_ClosetGoblin: ShowShowCloakOnly -> CharacterWindow.CreateTooltip("ClosetGoblinCharacterWindowContentsEquipmentShowCloak", GetString(StringTables.Default.TOOLTIP_SHOW_ITEM))
- CM_ClosetGoblin: ShowShowHelmOnly -> CharacterWindow.CreateTooltip("ClosetGoblinCharacterWindowContentsEquipmentShowHelm", GetString(StringTables.Default.TOOLTIP_SHOW_ITEM))
- CaVES: GetTotalBonusPower -> CharacterWindow.CalculateValueWithBonus(bonusType, 0)

## Notes

- none
