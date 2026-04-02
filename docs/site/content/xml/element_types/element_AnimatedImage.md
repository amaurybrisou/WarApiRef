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

## Common Handlers

- none

## Common Inherits

- EA_MoraleButtonAnimation
- LoadingScreenWarSymbolAnimation

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
