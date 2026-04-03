# AnimFrames

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
| Addons seen in | CMap, CoolDownLine, EA_LoadingScreen, MapMonster, RVMOD_SquaredDistances, TastyButtons, TheSeeker, TidyRoll |
| Files seen in | CustomTextures/TastyButtonsButtonTemplate.xml, Source/GeneralLoadingScreenTemplates.xml, Source/MapMonster_Templates.xml |
| Namespaces detected | AnimFrames |
| Source kinds | xml_frames |
| Example locations | CMap: CMapWindowMapRallyCallGlowAnim, CMap: CMapWindowMapScenarioQueueGlowAnim, CoolDownLine: CoolDownLineICONTemplateAFlash, CoolDownLine: CoolDownLineWndFlash, EA_LoadingScreen: LoadingScreenWarSymbolAnimation, MapMonster: AnimatedHighlight |
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

AnimFrames is a structural XML sub-element. It commonly appears under AnimatedImage. It is typically used to organize structural children such as AnimFrame.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [AnimatedImage](element_AnimatedImage.md) — 19× (HIGH)

## Common Structural Child Elements

- [AnimFrame](element_AnimFrame.md) — 222× (HIGH)

## Structural Sub-Elements

### [AnimFrame](element_AnimFrame.md)

Observed 222 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | **required** | 1, 2, 3, 4 |
| `x` | **required** | 64, 101, 138, 175 |
| `y` | **required** | 796, 0, 128, 256 |
## Recursive Hierarchy

- Root: [AnimFrames](element_AnimFrames.md)
- [AnimFrame](element_AnimFrame.md) (structural, 222×, HIGH)

## Seen In

- CMap
- CoolDownLine
- EA_LoadingScreen
- MapMonster
- RVMOD_SquaredDistances
- TastyButtons
- TheSeeker
- TidyRoll

## Examples

- CMap: CMapWindowMapRallyCallGlowAnim -> AnimFrames in AnimatedImage CMapWindowMapRallyCallGlowAnim
- CMap: CMapWindowMapScenarioQueueGlowAnim -> AnimFrames in AnimatedImage CMapWindowMapScenarioQueueGlowAnim
- CoolDownLine: CoolDownLineICONTemplateAFlash -> AnimFrames in AnimatedImage CoolDownLineICONTemplateAFlash
- CoolDownLine: CoolDownLineWndFlash -> AnimFrames in AnimatedImage CoolDownLineWndFlash
- EA_LoadingScreen: LoadingScreenWarSymbolAnimation -> AnimFrames in AnimatedImage LoadingScreenWarSymbolAnimation
- MapMonster: AnimatedHighlight -> AnimFrames in AnimatedImage AnimatedHighlight

## Related APIs

- [AnimatedImage](element_AnimatedImage.md) (HIGH 100/100) - XML Element Type
- [AnimFrame](element_AnimFrame.md) (MEDIUM 45/100) - XML Element Type
