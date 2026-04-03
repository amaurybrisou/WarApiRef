# AnimatedImage

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
| Addons seen in | Enemy, GuardLine, MoraleCircle, Shinies, TidyRoll |
| Files seen in | `/workspace/data/raw/Enemy/Code/Assist/Assist.xml:0`, `/workspace/data/raw/GuardLine/GuardLine.xml:0`, `/workspace/data/raw/MoraleCircle/MoraleCircle.xml:0`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Auctions.xml:0`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Browse.xml:0`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Post.xml:0`, `/workspace/data/raw/TidyRoll/TidyRoll.xml:0` |
| Namespaces detected | AnimatedImage |
| Source kinds | xml_frames |
| Example locations | Enemy: EnemyTargetFlash, Enemy: EnemyTargetGlow, Enemy: EnemyTargetSpark, GuardLine: GuardLineSelfWindowGlow, GuardLine: GuardLineTargetWindowGlow, MoraleCircle: MoraleTemplateFlash |
| XML usage count | 12 |
| XML attribute usage count | 12 |
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

Observed XML element type instantiated by 5 addons.

## Common Attributes

- layer
- name
- inherits
- handleinput
- texture
- fps
- textureScale
- alpha
- sticky
- texturescale

## Common Inherits

- EA_MoraleButtonAnimation
- LoadingScreenWarSymbolAnimation

## Common Parent Elements

- [Windows](element_Windows.md) — 12× (HIGH)

## Common Structural Child Elements

- [Size](element_Size.md) — 9× (HIGH)
- [AnimFrames](element_AnimFrames.md) — 2× (MEDIUM)
- [Windows](element_Windows.md) — 2× (MEDIUM)

## Common Template Bases

- EA_MoraleButtonAnimation
- LoadingScreenWarSymbolAnimation


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<AnimatedImage fps="10" handleinput="false" layer="overlay" name="..." sticky="false" texture="recharge_flash_anim" texturescale="1">
  <Size/>
  <AnimFrames>
    <AnimFrame id="1" x="0" y="0"/>
    <AnimFrame id="2" x="85" y="0"/>
    <AnimFrame id="3" x="170" y="0"/>
    <AnimFrame id="4" x="0" y="85"/>
    <AnimFrame id="5" x="85" y="85"/>
    <AnimFrame id="6" x="170" y="85"/>
    <AnimFrame id="7" x="0" y="170"/>
    <AnimFrame id="8" x="85" y="170"/>
  </AnimFrames>
</AnimatedImage>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `layer` | **required** | 100% | secondary, background, overlay |
| `inherits` | **required** | 83% | EA_MoraleButtonAnimation, LoadingScreenWarSymbolAnimation |
| `handleinput` | optional | 75% | false |
| `texture` | optional | 75% | anim_morale_glow, anim_morale_flash, GuardEffect, EA_ActionBarAnim_Casting, ... |
| `fps` | optional | 66% | 20, 13, 10 |
| `textureScale` | optional | 66% | 0.74, 1.2, 1, 0.9444 |
| `alpha` | optional | 25% | 1 |
| `sticky` | optional | 16% | false |
| `texturescale` | optional | 8% | 1 |
## Structural Sub-Elements

### [Size](element_Size.md)

Observed 9 times as an unnamed child.

### [AnimFrames](element_AnimFrames.md)

Observed 2 times as an unnamed child.

### [Windows](element_Windows.md)

Observed 2 times as an unnamed child.

## Lua Functions Manipulating This Type

- Enemy.AssistUI_Target_Show
- GuardLine.update

## Seen In

- Enemy
- GuardLine
- MoraleCircle
- Shinies
- TidyRoll

## Examples

- Enemy: EnemyTargetFlash -> AnimatedImage EnemyTargetFlash
- Enemy: EnemyTargetGlow -> AnimatedImage EnemyTargetGlow
- Enemy: EnemyTargetSpark -> AnimatedImage EnemyTargetSpark
- GuardLine: GuardLineSelfWindowGlow -> AnimatedImage GuardLineSelfWindowGlow
- GuardLine: GuardLineTargetWindowGlow -> AnimatedImage GuardLineTargetWindowGlow
- MoraleCircle: MoraleTemplateFlash -> AnimatedImage MoraleTemplateFlash

## Related APIs

- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [AnimFrames](element_AnimFrames.md) (MEDIUM 45/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
