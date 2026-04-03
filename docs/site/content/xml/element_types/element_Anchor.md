# Anchor

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

Observed XML element type instantiated by 33 addons.

## Common Attributes

- point
- relativePoint
- relativeTo
- xOffset
- yOffset

## Common Inherits

- none

## Common Parent Elements

- [Anchors](element_Anchors.md) — 3889× (HIGH)
- [Anchor](element_Anchor.md) — 22× (HIGH)
- [Window](element_Window.md) — 10× (HIGH)
- [Button](element_Button.md) — 2× (LOW)
- [DynamicImage](element_DynamicImage.md) — 1× (LOW)

## Common Structural Child Elements

- [AbsPoint](element_AbsPoint.md) — 3500× (HIGH)
- [Anchor](element_Anchor.md) — 22× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `point` | **required** | 99% | center, topleft, bottomright, right, ... |
| `relativePoint` | **required** | 99% | center, topleft, bottomright, left, ... |
| `relativeTo` | **required** | 88% | Root, $parent, $parentGeneral, $parentFollowTarget, ... |
| `xOffset` | optional | 0% | 0 |
| `yOffset` | optional | 0% | 0 |
## Structural Sub-Elements

### [AbsPoint](element_AbsPoint.md)

Observed 3500 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 120, 220, 550, 0 |
| `y` | **required** | 28, 850, 0, 32 |
### [Anchor](element_Anchor.md)

Observed 22 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `point` | **required** | center, topleft, bottomright, right |
| `relativePoint` | **required** | center, topleft, bottomright, left |
| `relativeTo` | **required** | Root, $parent, $parentGeneral, $parentFollowTarget |
| `xOffset` | optional | 0 |
| `yOffset` | optional | 0 |
## Recursive Hierarchy

- Root: [Anchor](element_Anchor.md)
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
- followTheLeader

## Examples

- none

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [AbsPoint](element_AbsPoint.md) (MEDIUM 45/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
