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
| Addons seen in | AggroMeter, Aura, Enemy, GuardLine, GuardList, GuardRange, MapPin, Miracle Grow Remix |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.xml:50`, `/workspace_addons/Aura/Source/Templates.xml:73`, `/workspace_addons/Enemy/Code/Core/Common.xml:25`, `/workspace_addons/Enemy/Code/Timer/Timer.xml:41`, `/workspace_addons/GuardLine/GuardLine.xml:120`, `/workspace_addons/GuardLine/GuardLine.xml:164`, `/workspace_addons/GuardLine/GuardLine.xml:196`, `/workspace_addons/GuardLine/GuardLine.xml:260` |
| Namespaces detected | CircleImage |
| Source kinds | xml_frames |
| Example locations | AggroMeter: Aggro_Button_TemplateIcon, Aura: AuraFrameImageCircle, Enemy: EnemyCircleImageTemplate, Enemy: EnemyStopwatchImage, GuardLine: GuardLineOffGuardSelfWindowIcon, GuardLine: GuardLineOffGuardTargetWindowIcon |
| XML usage count | 29 |
| XML attribute usage count | 29 |
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

Observed XML element type instantiated by 11 addons.

## Common Attributes

- name
- layer
- handleinput
- numsegments
- textureScale
- popable
- texture
- sticky
- alpha
- texturescale
- filtering
- movable

## Common Handlers

- none

## Common Inherits

- none

## Common Structural Child Elements

- TintColor

## Seen In

- AggroMeter
- Aura
- Enemy
- GuardLine
- GuardList
- GuardRange
- MapPin
- Miracle Grow Remix
- MoraleCircle
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

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
