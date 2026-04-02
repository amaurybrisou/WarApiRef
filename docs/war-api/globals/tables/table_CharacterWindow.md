# CharacterWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 140

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AnywhereTrainer, BankWindowFix, CM_ClosetGoblin, Effigy |
| Files seen in | `/workspace_addons/BankWindowFix/Source/BankWindowFix.lua:108`, `/workspace_addons/ClosetGoblin/ClosetGoblinCharacterWindow.lua:695`, `/workspace_addons/ClosetGoblin/ClosetGoblinCharacterWindow.lua:700`, `/workspace_addons/ClosetGoblin/ClosetGoblinCharacterWindow.lua:725`, `/workspace_addons/Effigy/States/EffigyStatePlayer.lua:193`, `/workspace_addons/Effigy/States/EffigyStatePlayer.lua:216` |
| Namespaces detected | CharacterWindow |
| Source kinds | globals, lua_calls |
| Example locations | BankWindowFix: BankWindowFix.BagEquipmentRButtonUp, CM_ClosetGoblin: ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly, CM_ClosetGoblin: ClosetGoblinCharacterWindow.ShowShowCloakOnly, CM_ClosetGoblin: ClosetGoblinCharacterWindow.ShowShowHelmOnly, Effigy: Effigy.UpdatePlayerArmor, Effigy: Effigy.UpdatePlayerToughness |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 3 |
| Local definition count | 1 |
| Documentation references | 1 |
| Initialization flow references | 2 |
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

Observed shared global table or namespace surfaced in 4 addons.

## Functions

- CharacterWindow.AutoEquipItem
- CharacterWindow.CalculateValueWithBonus
- CharacterWindow.CreateTooltip

## Observed Members

- none

## Seen In

- AnywhereTrainer
- BankWindowFix
- CM_ClosetGoblin
- Effigy

## Examples

- BankWindowFix: BankWindowFix.BagEquipmentRButtonUp -> CharacterWindow.AutoEquipItem(slot)
- CM_ClosetGoblin: ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly -> CharacterWindow.CreateTooltip("ClosetGoblinCharacterWindowContentsEquipmentShowCloakHeraldry", GetString(StringTables.Default.TOOLTIP_SHOW_HERALDRY))
- CM_ClosetGoblin: ClosetGoblinCharacterWindow.ShowShowCloakOnly -> CharacterWindow.CreateTooltip("ClosetGoblinCharacterWindowContentsEquipmentShowCloak", GetString(StringTables.Default.TOOLTIP_SHOW_ITEM))
- CM_ClosetGoblin: ClosetGoblinCharacterWindow.ShowShowHelmOnly -> CharacterWindow.CreateTooltip("ClosetGoblinCharacterWindowContentsEquipmentShowHelm", GetString(StringTables.Default.TOOLTIP_SHOW_ITEM))
- Effigy: Effigy.UpdatePlayerArmor -> CharacterWindow.CalculateValueWithBonus(GameData.BonusTypes.EBONUS_ARMOR, GameData.Player.armorValue)
- Effigy: Effigy.UpdatePlayerToughness -> CharacterWindow.CalculateValueWithBonus(GameData.BonusTypes.EBONUS_TOUGHNESS, baseValue)

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
