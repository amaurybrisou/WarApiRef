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
| Addons seen in | CMap, Enemy, GuardLine, GuardList, GuardRange, MapMonster, MapPin, MoraleCircle |
| Files seen in | `/workspace_addons/Enemy/Code/Assist/Assist.xml:57`, `/workspace_addons/Enemy/Code/Assist/Assist.xml:67`, `/workspace_addons/Enemy/Code/Assist/Assist.xml:77`, `/workspace_addons/GuardLine/GuardLine.xml:131`, `/workspace_addons/GuardLine/GuardLine.xml:223`, `/workspace_addons/GuardList/GuardList.xml:55`, `/workspace_addons/GuardRange/GuardRange.xml:55`, `/workspace_addons/MapMonster/Source/MapMonster_Templates.xml:27` |
| Namespaces detected | AnimatedImage |
| Source kinds | xml_frames |
| Example locations | CMap: CMapWindowMapRallyCallGlowAnim, CMap: CMapWindowMapScenarioQueueGlowAnim, Enemy: EnemyTargetFlash, Enemy: EnemyTargetGlow, Enemy: EnemyTargetSpark, GuardLine: GuardLineSelfWindowGlow |
| XML usage count | 19 |
| XML attribute usage count | 19 |
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

Observed XML element type instantiated by 10 addons.

## Common Attributes

- name
- layer
- handleinput
- texture
- fps
- inherits
- textureScale
- alpha
- sticky
- popable
- savesettings
- texturescale

## Common Inherits

- EA_MoraleButtonAnimation
- LoadingScreenWarSymbolAnimation

## Common Parent Elements

- [Window](element_Window.md)
- [Button](element_Button.md)


## Structural Sub-Elements

### [AnimFrame](element_AnimFrame.md)

- Observed in 5 parent frames
- Attributes: `id`, `x`, `y`
  - `id`: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`
  - `x`: `64`, `101`, `138`, `175`, `212`, `249`, `2`, `200`
  - `y`: `796`, `0`, `128`, `256`, `64`, `85`, `170`

### [AnimFrames](element_AnimFrames.md)

- Observed in 5 parent frames

## Typical XML Structure

```xml
<AnimatedImage name="..." texture="MM_AniArrows_Tintable" fps="8" sticky="false" handleinput="false" alpha="1" popable="true" savesettings="false">
  <AnimFrames>
    <AnimFrame id="1" x="0" y="0"/>
    <AnimFrame id="2" x="64" y="0"/>
    <AnimFrame id="3" x="128" y="0"/>
    <AnimFrame id="4" x="192" y="0"/>
    <AnimFrame id="5" x="256" y="0"/>
    <AnimFrame id="6" x="0" y="64"/>
    <AnimFrame id="7" x="64" y="64"/>
    <AnimFrame id="8" x="128" y="64"/>
    <AnimFrame id="9" x="192" y="64"/>
    <AnimFrame id="10" x="256" y="64"/>
  </AnimFrames>
</AnimatedImage>
```


## Attribute Reference

| Attribute | Role | Observed Values | Count |
|-----------|------|-----------------|-------|
| `name` | string | `CMapWindowMapRallyCallGlowAnim`, `CMapWindowMapScenarioQueueGlowAnim`, `EnemyTargetFlash`, `EnemyTargetGlow`, … | 19 |
| `layer` | string | `background`, `secondary`, `overlay`, `default` | 18 |
| `handleinput` | boolean | `false` | 16 |
| `texture` | string | `EA_HUD_01`, `anim_fury_round_1`, `anim_morale_flash`, `anim_morale_glow`, … | 16 |
| `fps` | number | `6`, `10`, `20`, `13`, … | 15 |
| `inherits` | frame-ref | `EA_MoraleButtonAnimation`, `LoadingScreenWarSymbolAnimation` | 14 |
| `textureScale` | number | `1.0`, `0.43`, `1.2`, `0.74`, … | 14 |
| `alpha` | number | `1`, `0.9` | 8 |
| `sticky` | boolean | `false` | 5 |
| `popable` | boolean | `true` | 1 |
| `savesettings` | boolean | `false` | 1 |
| `texturescale` | number | `1` | 1 |

## Seen In

- CMap
- Enemy
- GuardLine
- GuardList
- GuardRange
- MapMonster
- MapPin
- MoraleCircle
- Shinies
- TidyRoll

## Examples

- CMap: CMapWindowMapRallyCallGlowAnim -> AnimatedImage CMapWindowMapRallyCallGlowAnim
- CMap: CMapWindowMapScenarioQueueGlowAnim -> AnimatedImage CMapWindowMapScenarioQueueGlowAnim
- Enemy: EnemyTargetFlash -> AnimatedImage EnemyTargetFlash
- Enemy: EnemyTargetGlow -> AnimatedImage EnemyTargetGlow
- Enemy: EnemyTargetSpark -> AnimatedImage EnemyTargetSpark
- GuardLine: GuardLineSelfWindowGlow -> AnimatedImage GuardLineSelfWindowGlow

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
