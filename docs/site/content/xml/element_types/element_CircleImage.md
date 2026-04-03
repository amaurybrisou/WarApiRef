# CircleImage

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
| Addons seen in | AggroMeter, Aura, Enemy, GuardLine, MoraleCircle, PartyCast, RoR_SoR, TurretRange |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/Enemy/Code/Core/Common.xml:0`, `/workspace/data/raw/Enemy/Code/Timer/Timer.xml:0`, `/workspace/data/raw/GuardLine/GuardLine.xml:0`, `/workspace/data/raw/MoraleCircle/MoraleCircle.xml:0`, `/workspace/data/raw/PartyCast/PartyCast.xml:0`, `/workspace/data/raw/RoR_SoR/RoR_SoR.xml:0` |
| Namespaces detected | CircleImage |
| Source kinds | xml_frames |
| Example locations | AggroMeter: Aggro_Button_TemplateIcon, Aura: AuraFrameImageCircle, Enemy: EnemyCircleImageTemplate, Enemy: EnemyStopwatchImage, GuardLine: GuardLineOffGuardSelfWindowIcon, GuardLine: GuardLineOffGuardTargetWindowIcon |
| XML usage count | 23 |
| XML attribute usage count | 23 |
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

- name
- layer
- numsegments
- popable
- textureScale
- handleinput
- sticky
- texture
- filtering
- movable

## Common Inherits

- none

## Common Parent Elements

- [Windows](element_Windows.md) — 23× (HIGH)

## Common Structural Child Elements

- [Size](element_Size.md) — 22× (HIGH)
- [TexCoords](element_TexCoords.md) — 20× (HIGH)
- [TintColor](element_TintColor.md) — 6× (HIGH)
- [TexDims](element_TexDims.md) — 2× (MEDIUM)
- [Windows](element_Windows.md) — 1× (LOW)


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<CircleImage handleinput="false" movable="false" name="..." texture="StopwatchButton">
  <TexCoords x="32" y="32"/>
  <TexDims x="64" y="64"/>
  <Size/>
</CircleImage>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `layer` | **required** | 91% | overlay, background, default |
| `numsegments` | **required** | 86% | 32, 16 |
| `popable` | **required** | 86% | false |
| `textureScale` | **required** | 86% | 0.8, 0.48, 0.38, 1.25, ... |
| `handleinput` | optional | 78% | false |
| `sticky` | optional | 65% | true |
| `texture` | optional | 60% | StopwatchButton, Circle, EA_TintableImage |
| `filtering` | optional | 4% | true |
| `movable` | optional | 4% | false |
## Structural Sub-Elements

### [Size](element_Size.md)

Observed 22 times as an unnamed child.

### [TexCoords](element_TexCoords.md)

Observed 20 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 0, 133, 88, 32 |
| `y` | optional | 0, 163, 51, 32 |
### [TintColor](element_TintColor.md)

Observed 6 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 110, 130, 50 |
| `g` | **required** | 0, 200, 185, 130 |
| `r` | **required** | 255, 0, 180, 200 |
| `a` | optional | 30, 255, 0.5, 200 |
### [TexDims](element_TexDims.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 64, 256, 32, 430 |
| `y` | **required** | 64, 256, 32, 430 |
### [Windows](element_Windows.md)

Observed 1 times as an unnamed child.

## Lua Functions Manipulating This Type

- Enemy.EnemyEffectsIndicator:Update
- TurretRange.OnUpdate

## Seen In

- AggroMeter
- Aura
- Enemy
- GuardLine
- MoraleCircle
- PartyCast
- RoR_SoR
- TurretRange

## Examples

- AggroMeter: Aggro_Button_TemplateIcon -> CircleImage Aggro_Button_TemplateIcon
- Aura: AuraFrameImageCircle -> CircleImage AuraFrameImageCircle
- Enemy: EnemyCircleImageTemplate -> CircleImage EnemyCircleImageTemplate
- Enemy: EnemyStopwatchImage -> CircleImage EnemyStopwatchImage
- GuardLine: GuardLineOffGuardSelfWindowIcon -> CircleImage GuardLineOffGuardSelfWindowIcon
- GuardLine: GuardLineOffGuardTargetWindowIcon -> CircleImage GuardLineOffGuardTargetWindowIcon

## Related APIs

- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TexDims](element_TexDims.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Used With

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none
