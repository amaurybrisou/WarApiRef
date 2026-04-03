# OnSelChanged

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 88/100

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, referenced by generated docs or reference files, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, BankArkel, BuffHead, Busted, CCTV, CDown, Calling |
| Namespaces detected | OnSelChanged |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedPetAssist: .OnSelChanged, AdvancedRenownTrainer: .OnSelChanged, BankArkel: .OnSelChanged, BuffHead: .OnSelChanged, Busted: .OnSelChanged, CCTV: .OnSelChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 484 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
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

XML handler event observed across 70 addons.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- none

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- BankArkel
- BuffHead
- Busted
- CCTV
- CDown
- Calling
- CastSequence
- ChattyCathy
- Cheeseboard
- Crusher
- DAoCBuff
- DPSMeter
- DaemonAssist
- DammazKron
- DetauntHelper
- DuffTimer
- Dye Preview
- EA_OpenPartyWindow
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- FozAuction
- GroupRange
- HealGrid
- Hopper
- KeysetMonsterPlay
- Killer
- LibAddonButton
- LoyalPet
- Map
- MapMonster
- MapPin
- MarkBuff
- MegaphonePlusPlus
- Miracle Grow Remix
- MoraleSet
- Motion
- NerfedButtons
- Obsidian
- Pocket Palette
- PotionBar
- Pure
- RVMOD_Manager
- RVMOD_SquaredDistances
- RVMOD_Targets
- RaidMeter
- ScenarioStats
- Shinies
- SocialWindow 2.0
- TalismanGenie
- TastyButtons
- TaxPayer
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyRoll
- TokenMachine
- TurretRange
- TwisterSet
- WARCommander
- WSCT
- WarBoard_FPS
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- ZonePOP
- nLootLink
- wbLeadHelper

## Examples

- AdvancedPetAssist: .OnSelChanged -> APAGui.OnComboChanged
- AdvancedRenownTrainer: .OnSelChanged -> AdvancedRenownTraining.SelectedItemChanged
- BankArkel: .OnSelChanged -> BankArkel.PackCombo
- BuffHead: .OnSelChanged -> BuffHead.Setup.AdvancedContainersItem.OnTargetTypeChanged
- BuffHead: .OnSelChanged -> BuffHead.Setup.AdvancedContainersItem.OnPositionChanged
- BuffHead: .OnSelChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnPositionPlacementChanged

## Related APIs

- [ComboBox](../element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function

## Affects

- [GameData.ItemLocs.EQUIPPED](../../gamedata/fields/gamedata_GameData.ItemLocs.EQUIPPED.md) (HIGH 100/100) - GameData Field
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected binding arguments remain uncertain because contract artifacts capture symbol linkage, not full handler signatures.
