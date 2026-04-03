# TextOffset

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, Aura, CM_ClosetGoblin, Pocket Palette, PotionBar, Shinies, WSCT, WarBoard |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.xml:0`, `/workspace/data/raw/PocketPalette/PocketPalette.xml:0`, `/workspace/data/raw/PotionBar/source/Floating.xml:0`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Auctions.xml:0`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Auto.xml:0`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Post.xml:0` |
| Namespaces detected | TextOffset |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APA_ComboBoxButton, AdvancedPetAssist: APA_ComboBoxButtonWide, Aura: Aura_Button_DefaultMenuButton, Aura: Aura_Button_DefaultMenuButtonLarge, Aura: Aura_Button_DefaultMenuButtonTiny, Aura: Aura_Button_DefaultResizableTinyComboBoxSelected |
| XML usage count | 33 |
| XML attribute usage count | 33 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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

Observed XML element type instantiated by 8 addons.

## Common Attributes

- x
- y

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 18× (HIGH)
- [EditBox](element_EditBox.md) — 15× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | **required** | 100% | 5, 4, 10, 0 |
| `y` | **required** | 87% | 5, 2, 3, 0 |
## Seen In

- AdvancedPetAssist
- Aura
- CM_ClosetGoblin
- Pocket Palette
- PotionBar
- Shinies
- WSCT
- WarBoard

## Examples

- AdvancedPetAssist: APA_ComboBoxButton -> TextOffset in Button APA_ComboBoxButton
- AdvancedPetAssist: APA_ComboBoxButtonWide -> TextOffset in Button APA_ComboBoxButtonWide
- Aura: Aura_Button_DefaultMenuButton -> TextOffset in Button Aura_Button_DefaultMenuButton
- Aura: Aura_Button_DefaultMenuButtonLarge -> TextOffset in Button Aura_Button_DefaultMenuButtonLarge
- Aura: Aura_Button_DefaultMenuButtonTiny -> TextOffset in Button Aura_Button_DefaultMenuButtonTiny
- Aura: Aura_Button_DefaultResizableTinyComboBoxSelected -> TextOffset in Button Aura_Button_DefaultResizableTinyComboBoxSelected

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [EditBox](element_EditBox.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
