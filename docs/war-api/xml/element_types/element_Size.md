# Size

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Aura, AutoMark, BagOMatic, BankArkel |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.xml:0`, `/workspace/data/raw/Aura/Source/AuraColorPicker.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTexture.xml:0` |
| Namespaces detected | Size |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD, AdvancedPetAssist: APAFollowTargetHUDLabel, AdvancedPetAssist: APAInstantOnlyHUD, AdvancedPetAssist: APAInstantOnlyHUDLabel, AdvancedPetAssist: APAKitingHUD, AdvancedPetAssist: APAKitingHUDLabel |
| XML usage count | 2635 |
| XML attribute usage count | 2635 |
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

Size is a structural XML sub-element. It commonly appears under AnimatedImage and Button. It is typically used to organize structural children such as AbsPoint.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Label](element_Label.md) — 1168× (HIGH)
- [Window](element_Window.md) — 624× (HIGH)
- [Button](element_Button.md) — 259× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 190× (HIGH)
- [EditBox](element_EditBox.md) — 127× (HIGH)
- [SliderBar](element_SliderBar.md) — 76× (HIGH)
- [ComboBox](element_ComboBox.md) — 52× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 35× (HIGH)
- [ListBox](element_ListBox.md) — 31× (HIGH)
- [VerticalScrollbar](element_VerticalScrollbar.md) — 23× (HIGH)
- [CircleImage](element_CircleImage.md) — 22× (HIGH)
- [ScrollWindow](element_ScrollWindow.md) — 12× (HIGH)
- [AnimatedImage](element_AnimatedImage.md) — 9× (MEDIUM)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 4× (MEDIUM)
- [ColorPicker](element_ColorPicker.md) — 1× (LOW)
- [CooldownDisplay](element_CooldownDisplay.md) — 1× (LOW)
- [MapDisplay](element_MapDisplay.md) — 1× (LOW)

## Common Structural Child Elements

- [AbsPoint](element_AbsPoint.md) — 2634× (HIGH)

## Structural Sub-Elements

### [AbsPoint](element_AbsPoint.md)

Observed 2634 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 120, 220, 550, 0 |
| `y` | **required** | 28, 850, 0, 32 |
## Recursive Hierarchy

- Root: [Size](element_Size.md)
- [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Aura
- AutoMark
- BagOMatic
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

- AdvancedPetAssist: APAFollowTargetHUD -> Size in Window APAFollowTargetHUD
- AdvancedPetAssist: APAFollowTargetHUDLabel -> Size in Label APAFollowTargetHUDLabel
- AdvancedPetAssist: APAInstantOnlyHUD -> Size in Window APAInstantOnlyHUD
- AdvancedPetAssist: APAInstantOnlyHUDLabel -> Size in Label APAInstantOnlyHUDLabel
- AdvancedPetAssist: APAKitingHUD -> Size in Window APAKitingHUD
- AdvancedPetAssist: APAKitingHUDLabel -> Size in Label APAKitingHUDLabel

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
- [VerticalScrollbar](element_VerticalScrollbar.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [AbsPoint](element_AbsPoint.md) (MEDIUM 45/100) - XML Element Type
- [CooldownDisplay](element_CooldownDisplay.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 55/100) - XML Element Type

## Triggered By

- none

## Affects

- none
