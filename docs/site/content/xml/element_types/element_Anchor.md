# Anchor

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
| Addons seen in | DetauntHelper, Enemy, NerfedButtons, Pocket Palette, SOR, Sequencer, SocialWindow 2.0, WarBoard |
| Files seen in | Code/CombatLog/CombatLogConfiguration.xml, Code/GroupIcons/GroupIconsConfiguration.xml, Code/Guard/GuardConfiguration.xml, Code/KillSpam/KillSpamConfiguration.xml, Code/TalismanAlerter/TalismanAlerterConfiguration.xml, Code/Timer/TimerConfiguration.xml, Code/UnitFrames/ClickCastingDialog.xml, Code/UnitFrames/EffectsIndicatorDialog.xml |
| Namespaces detected | Anchor |
| Source kinds | xml_frames |
| Example locations | DetauntHelper: DTC_TARGETS_TemplateButtonBackground, DetauntHelper: DetauntHelperBarTargetDamage, DetauntHelper: DetauntHelperBarTargetName, Enemy: EnemyClickCastingDialogContentScrollChild, Enemy: EnemyCombatLogConfigurationContentScrollChild, Enemy: EnemyEffectsIndicatorDialogContentScrollChild |
| XML usage count | 22 |
| XML attribute usage count | 22 |
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

Anchor is a structural XML sub-element. It commonly appears under Anchor and Anchors. It is typically used to organize structural children such as AbsPoint, Anchor.

## Common Attributes

- point
- relativePoint
- relativeTo

## Common Inherits

- none

## Common Parent Elements

- [Anchors](element_Anchors.md) — 14161× (HIGH)
- [Anchor](element_Anchor.md) — 29× (HIGH)
- [Window](element_Window.md) — 16× (HIGH)
- [Button](element_Button.md) — 2× (LOW)
- [FullResizeImage](element_FullResizeImage.md) — 2× (LOW)
- [Label](element_Label.md) — 2× (LOW)
- [DynamicImage](element_DynamicImage.md) — 1× (LOW)

## Common Structural Child Elements

- [AbsPoint](element_AbsPoint.md) — 12549× (HIGH)
- [Anchor](element_Anchor.md) — 29× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `point` | **required** | 99% | center, topleft, topright, bottomright, ... |
| `relativePoint` | **required** | 99% | center, topleft, topright, bottomright, ... |
| `relativeTo` | **required** | 90% | Root, $parent, $parentSliderLabel, $parentSliderValueLabel, ... |
| `relativeto` | optional | 0% | $parentTitle, $parentPrint, $parentAlert, $parentHyperlink, ... |
| `layer` | optional | 0% | secondary, overlay |
| `parent` | optional | 0% | Root, $parent |
| `relateiveTo` | optional | 0% | $parentDevPadCode, $parentObjectEditBox |
| `textalign` | optional | 0% | center |
| `handleinput` | optional | 0% | false |
| `relative` | optional | 0% | $parent |
| `xOffset` | optional | 0% | 0 |
| `yOffset` | optional | 0% | 0 |
| `input` | optional | 0% | numbers |
| `realtivePoint` | optional | 0% | center |
| `realtiveTo` | optional | 0% | $parentBackground |
## Structural Sub-Elements

### [AbsPoint](element_AbsPoint.md)

Observed 12549 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 800, 200, 0, 100 |
| `y` | **required** | 50, -50, 100, -200 |
### [Anchor](element_Anchor.md)

Observed 29 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `point` | **required** | center, topleft, topright, bottomright |
| `relativePoint` | **required** | center, topleft, topright, bottomright |
| `relativeTo` | **required** | Root, $parent, $parentSliderLabel, $parentSliderValueLabel |
| `relativeto` | optional | $parentTitle, $parentPrint, $parentAlert, $parentHyperlink |
| `layer` | optional | secondary, overlay |
| `parent` | optional | Root, $parent |
| `relateiveTo` | optional | $parentDevPadCode, $parentObjectEditBox |
| `textalign` | optional | center |
| `handleinput` | optional | false |
| `relative` | optional | $parent |
| `xOffset` | optional | 0 |
| `yOffset` | optional | 0 |
| `input` | optional | numbers |
| `realtivePoint` | optional | center |
| `realtiveTo` | optional | $parentBackground |
## Recursive Hierarchy

- Root: [Anchor](element_Anchor.md)
- [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
- [Anchor](element_Anchor.md) (structural, 29×, HIGH)
  - (cycle)

## Seen In

- DetauntHelper
- Enemy
- NerfedButtons
- Pocket Palette
- SOR
- Sequencer
- SocialWindow 2.0
- WarBoard
- followTheLeader
- wbLeadHelper

## Examples

- DetauntHelper: DTC_TARGETS_TemplateButtonBackground -> Anchor in Window DTC_TARGETS_TemplateButtonBackground
- DetauntHelper: DetauntHelperBarTargetDamage -> Anchor in Label DetauntHelperBarTargetDamage
- DetauntHelper: DetauntHelperBarTargetName -> Anchor in Label DetauntHelperBarTargetName
- Enemy: EnemyClickCastingDialogContentScrollChild -> Anchor in Window EnemyClickCastingDialogContentScrollChild
- Enemy: EnemyCombatLogConfigurationContentScrollChild -> Anchor in Window EnemyCombatLogConfigurationContentScrollChild
- Enemy: EnemyEffectsIndicatorDialogContentScrollChild -> Anchor in Window EnemyEffectsIndicatorDialogContentScrollChild

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [AbsPoint](element_AbsPoint.md) (MEDIUM 45/100) - XML Element Type
