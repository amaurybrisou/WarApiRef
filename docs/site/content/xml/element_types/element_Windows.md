# Windows

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Aura, AutoMark, BankArkel, BuffHead |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.xml:0`, `/workspace/data/raw/Aura/Source/AuraColorPicker.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTexture.xml:0` |
| Namespaces detected | Windows |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD, AdvancedPetAssist: APAInstantOnlyHUD, AdvancedPetAssist: APAKitingHUD, AdvancedPetAssist: APAOptions, AdvancedPetAssist: APAOptionsContent, AdvancedPetAssist: APAOptionsTabs |
| XML usage count | 530 |
| XML attribute usage count | 530 |
| Lua usage count | 0 |
| Global usage count | 0 |
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

Observed XML element type instantiated by 32 addons.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Window](element_Window.md) — 476× (HIGH)
- [ScrollWindow](element_ScrollWindow.md) — 24× (HIGH)
- [Button](element_Button.md) — 18× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 6× (MEDIUM)
- [AnimatedImage](element_AnimatedImage.md) — 2× (LOW)
- [CircleImage](element_CircleImage.md) — 1× (LOW)
- [FullResizeImage](element_FullResizeImage.md) — 1× (LOW)
- [Label](element_Label.md) — 1× (LOW)
- [LogDisplay](element_LogDisplay.md) — 1× (LOW)

## Common Named Child Elements

- [Label](element_Label.md) — 1243× (HIGH)
- [Window](element_Window.md) — 1001× (HIGH)
- [Button](element_Button.md) — 664× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 237× (HIGH)
- [ComboBox](element_ComboBox.md) — 233× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 156× (HIGH)
- [EditBox](element_EditBox.md) — 151× (HIGH)
- [SliderBar](element_SliderBar.md) — 83× (HIGH)
- [ListBox](element_ListBox.md) — 42× (HIGH)
- [ScrollWindow](element_ScrollWindow.md) — 26× (HIGH)
- [VerticalScrollbar](element_VerticalScrollbar.md) — 25× (HIGH)
- [CircleImage](element_CircleImage.md) — 23× (HIGH)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 19× (HIGH)
- [AnimatedImage](element_AnimatedImage.md) — 12× (HIGH)
- [StatusBar](element_StatusBar.md) — 9× (MEDIUM)
- [VerticalResizeImage](element_VerticalResizeImage.md) — 4× (MEDIUM)
- [ColorPicker](element_ColorPicker.md) — 1× (LOW)
- [CooldownDisplay](element_CooldownDisplay.md) — 1× (LOW)
- [LogDisplay](element_LogDisplay.md) — 1× (LOW)
- [MapDisplay](element_MapDisplay.md) — 1× (LOW)

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Aura
- AutoMark
- BankArkel
- BuffHead
- CM_ClosetGoblin
- CombatTextNames
- DAoCBuff
- Enemy
- GuardLine
- Killer
- LibGroup
- LibWBToggler
- MiracleGrowLight
- MoraleCircle
- PartyCast
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- Swift Assist
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow

## Examples

- AdvancedPetAssist: APAFollowTargetHUD -> Windows in Window APAFollowTargetHUD
- AdvancedPetAssist: APAInstantOnlyHUD -> Windows in Window APAInstantOnlyHUD
- AdvancedPetAssist: APAKitingHUD -> Windows in Window APAKitingHUD
- AdvancedPetAssist: APAOptions -> Windows in Window APAOptions
- AdvancedPetAssist: APAOptionsContent -> Windows in Window APAOptionsContent
- AdvancedPetAssist: APAOptionsTabs -> Windows in Window APAOptionsTabs

## Related APIs

- [AnimatedImage](element_AnimatedImage.md) (HIGH 100/100) - XML Element Type
- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [CircleImage](element_CircleImage.md) (HIGH 100/100) - XML Element Type
- [ColorPicker](element_ColorPicker.md) (HIGH 100/100) - XML Element Type
- [ComboBox](element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EditBox](element_EditBox.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [HorizontalResizeImage](element_HorizontalResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [ScrollWindow](element_ScrollWindow.md) (HIGH 100/100) - XML Element Type
- [SliderBar](element_SliderBar.md) (HIGH 100/100) - XML Element Type
- [StatusBar](element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [VerticalScrollbar](element_VerticalScrollbar.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [VerticalResizeImage](element_VerticalResizeImage.md) (HIGH 90/100) - XML Element Type
- [CooldownDisplay](element_CooldownDisplay.md) (MEDIUM 45/100) - XML Element Type
- [LogDisplay](element_LogDisplay.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none
