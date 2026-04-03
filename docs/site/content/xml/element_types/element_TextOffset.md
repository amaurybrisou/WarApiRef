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
| Addons seen in | AdvancedPetAssist, Assist, Aura, AutoBand, CM_ClosetGoblin, Crusher, DaemonAssist, DetauntHelper |
| Files seen in | Configuration/Config.xml, Configuration/HopperConfig.xml, Configuration/WCDBConfig.xml, Configuration/WCDPConfig.xml, CustomTextures/TastyButtonsButtonTemplate.xml, Modules/UI/Shinies-UI-Auctions.xml, Modules/UI/Shinies-UI-Auto.xml, Modules/UI/Shinies-UI-Post.xml |
| Namespaces detected | TextOffset |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APA_ComboBoxButton, AdvancedPetAssist: APA_ComboBoxButtonWide, Assist: AssistWindowBTN, Aura: Aura_Button_DefaultMenuButton, Aura: Aura_Button_DefaultMenuButtonLarge, Aura: Aura_Button_DefaultMenuButtonTiny |
| XML usage count | 138 |
| XML attribute usage count | 138 |
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

TextOffset is a XML UI element. It commonly appears under Button and ComboBox.

## Common Attributes

- x
- y

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 75× (HIGH)
- [EditBox](element_EditBox.md) — 63× (HIGH)
- [ComboBox](element_ComboBox.md) — 1× (LOW)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | **required** | 100% | 5, 6, 4, 0, ... |
| `y` | **required** | 94% | 5, 12, 2, 3, ... |
## Seen In

- AdvancedPetAssist
- Assist
- Aura
- AutoBand
- CM_ClosetGoblin
- Crusher
- DaemonAssist
- DetauntHelper
- EA_OpenPartyWindow
- EA_UiDebugTools
- EZCraftX
- FozAuction
- GDes
- Ges
- Hopper
- ItemRack
- LoyalPet
- Miracle Grow Remix
- MoraleSet
- Motion
- NaturalLog
- ObjectInspector
- PartyAd
- PieTracker
- Pocket Palette
- PotionBar
- Pure
- Queue Queuer
- RVMOD_Manager
- RaidMeter
- Res
- SNT_PANEL
- Shinies
- TastyButtons
- TomeTracker
- TwisterSet
- Vectors
- WSCT
- WarBoard
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- ZonePOP
- wbLeadHelper
- zMailMod

## Examples

- AdvancedPetAssist: APA_ComboBoxButton -> TextOffset in Button APA_ComboBoxButton
- AdvancedPetAssist: APA_ComboBoxButtonWide -> TextOffset in Button APA_ComboBoxButtonWide
- Assist: AssistWindowBTN -> TextOffset in Button AssistWindowBTN
- Aura: Aura_Button_DefaultMenuButton -> TextOffset in Button Aura_Button_DefaultMenuButton
- Aura: Aura_Button_DefaultMenuButtonLarge -> TextOffset in Button Aura_Button_DefaultMenuButtonLarge
- Aura: Aura_Button_DefaultMenuButtonTiny -> TextOffset in Button Aura_Button_DefaultMenuButtonTiny

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [EditBox](element_EditBox.md) (HIGH 100/100) - XML Element Type
