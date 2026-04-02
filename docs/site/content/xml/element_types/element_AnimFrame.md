# AnimFrame

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 98/100

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CMap, MapMonster, TidyRoll |
| Files seen in | `/workspace_addons/MapMonster/Source/MapMonster_Templates.xml:27`, `/workspace_addons/TidyRoll/TidyRoll.xml:137`, `/workspace_addons/TidyRoll/TidyRoll.xml:156`, `/workspace_addons/cmap/CMap.xml:197`, `/workspace_addons/cmap/CMap.xml:315` |
| Namespaces detected | AnimFrame |
| Source kinds | xml_frames |
| Example locations | CMap: CMapWindowMapRallyCallGlowAnim, CMap: CMapWindowMapScenarioQueueGlowAnim, MapMonster: AnimatedHighlight, TidyRoll: TRollOverlayFlash, TidyRoll: TRollOverlayGlow |
| XML usage count | 5 |
| XML attribute usage count | 5 |
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

Observed XML element type instantiated by 3 addons.

## Common Attributes

- id
- x
- y

## Common Inherits

- none

## Common Parent Elements

- [AnimatedImage](element_AnimatedImage.md) — 5× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `id` | **required** | 100% |  |
| `x` | **required** | 100% |  |
| `y` | **required** | 100% |  |
## Seen In

- CMap
- MapMonster
- TidyRoll

## Examples

- CMap: CMapWindowMapRallyCallGlowAnim -> AnimFrame in AnimatedImage CMapWindowMapRallyCallGlowAnim
- CMap: CMapWindowMapScenarioQueueGlowAnim -> AnimFrame in AnimatedImage CMapWindowMapScenarioQueueGlowAnim
- MapMonster: AnimatedHighlight -> AnimFrame in AnimatedImage AnimatedHighlight
- TidyRoll: TRollOverlayFlash -> AnimFrame in AnimatedImage TRollOverlayFlash
- TidyRoll: TRollOverlayGlow -> AnimFrame in AnimatedImage TRollOverlayGlow

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
