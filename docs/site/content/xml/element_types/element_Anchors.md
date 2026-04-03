# Anchors

- Category: XML Element Type
- Confidence level: MEDIUM
- Confidence score: 55/100

## Confidence Assessment

- Level: MEDIUM

- Score: 55/100

- Rationale: unknown

## Evidence Signals

- none

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | none |
| Source kinds | none |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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

Anchors is a structural XML sub-element. It commonly appears under AnimatedImage and Button. It is typically used to organize structural children such as AbsPoint, Anchor.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Label](element_Label.md) — 1238× (HIGH)
- [Window](element_Window.md) — 745× (HIGH)
- [Button](element_Button.md) — 556× (HIGH)
- [ComboBox](element_ComboBox.md) — 228× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 216× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 139× (HIGH)
- [EditBox](element_EditBox.md) — 126× (HIGH)
- [SliderBar](element_SliderBar.md) — 83× (HIGH)
- [ListBox](element_ListBox.md) — 42× (HIGH)
- [VerticalScrollbar](element_VerticalScrollbar.md) — 25× (HIGH)
- [CircleImage](element_CircleImage.md) — 22× (HIGH)
- [ScrollWindow](element_ScrollWindow.md) — 16× (HIGH)
- [AnimatedImage](element_AnimatedImage.md) — 12× (HIGH)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 11× (HIGH)
- [StatusBar](element_StatusBar.md) — 9× (MEDIUM)
- [ColorPicker](element_ColorPicker.md) — 1× (LOW)
- [CooldownDisplay](element_CooldownDisplay.md) — 1× (LOW)
- [LogDisplay](element_LogDisplay.md) — 1× (LOW)
- [MapDisplay](element_MapDisplay.md) — 1× (LOW)

## Common Structural Child Elements

- [Anchor](element_Anchor.md) — 3889× (HIGH)
- [AbsPoint](element_AbsPoint.md) — 3× (MEDIUM)

## Structural Sub-Elements

### [Anchor](element_Anchor.md)

Observed 3889 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `point` | **required** | center, topleft, bottomright, right |
| `relativePoint` | **required** | center, topleft, bottomright, left |
| `relativeTo` | **required** | Root, $parent, $parentGeneral, $parentFollowTarget |
| `xOffset` | optional | 0 |
| `yOffset` | optional | 0 |
### [AbsPoint](element_AbsPoint.md)

Observed 3 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 120, 220, 550, 0 |
| `y` | **required** | 28, 850, 0, 32 |
## Recursive Hierarchy

- Root: [Anchors](element_Anchors.md)
- [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
- [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
    - (cycle)

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

- none

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
- [Anchor](element_Anchor.md) (MEDIUM 55/100) - XML Element Type
- [AbsPoint](element_AbsPoint.md) (MEDIUM 45/100) - XML Element Type
- [CooldownDisplay](element_CooldownDisplay.md) (MEDIUM 45/100) - XML Element Type
- [LogDisplay](element_LogDisplay.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 55/100) - XML Element Type

## Triggered By

- none

## Affects

- none
