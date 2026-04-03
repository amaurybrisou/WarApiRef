# Sounds

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
| Addons seen in | CMap, CastSequence, FastInteract, Minmap, Miracle Grow Remix, MiracleGrow, NerfedButtons, SocialWindow 2.0 |
| Files seen in | Setup/Builder.xml, Source/SocialWindow.xml, UI/NBSBChecks.xml |
| Namespaces detected | Sounds |
| Source kinds | xml_frames |
| Example locations | CMap: CMapWindowMapRallyCall, CMap: CMapWindowMapRankedLeaderboard, CMap: CMapWindowMapScenarioQueue, CMap: CMapWindowMapWorldMapButton, CastSequence: CastSequenceBuilderWindowSettingsButton, FastInteract: FastInteractWindow |
| XML usage count | 13 |
| XML attribute usage count | 13 |
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

Sounds is a structural XML sub-element. It commonly appears under Button and DynamicImage. It is typically used to organize structural children such as Sound.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 9× (HIGH)
- [Window](element_Window.md) — 2× (MEDIUM)
- [DynamicImage](element_DynamicImage.md) — 1× (LOW)
- [Label](element_Label.md) — 1× (LOW)

## Common Structural Child Elements

- [Sound](element_Sound.md) — 20× (HIGH)

## Structural Sub-Elements

### [Sound](element_Sound.md)

Observed 20 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `event` | **required** | OnLButtonDown, OnMouseOver, OnShown, OnHidden |
| `script` | **required** | Sound.Play( Sound.BUTTON_OVER ), Sound.Play(Sound.BUTTON_CLICK), Sound.Play(Sound.BUTTON_OVER), Sound.Play( Sound.TOME_OPEN ) |
## Recursive Hierarchy

- Root: [Sounds](element_Sounds.md)
- [Sound](element_Sound.md) (structural, 20×, HIGH)

## Seen In

- CMap
- CastSequence
- FastInteract
- Minmap
- Miracle Grow Remix
- MiracleGrow
- NerfedButtons
- SocialWindow 2.0
- zMailMod

## Examples

- CMap: CMapWindowMapRallyCall -> Sounds in Button CMapWindowMapRallyCall
- CMap: CMapWindowMapRankedLeaderboard -> Sounds in Button CMapWindowMapRankedLeaderboard
- CMap: CMapWindowMapScenarioQueue -> Sounds in Button CMapWindowMapScenarioQueue
- CMap: CMapWindowMapWorldMapButton -> Sounds in Button CMapWindowMapWorldMapButton
- CastSequence: CastSequenceBuilderWindowSettingsButton -> Sounds in DynamicImage CastSequenceBuilderWindowSettingsButton
- FastInteract: FastInteractWindow -> Sounds in Window FastInteractWindow

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Sound](element_Sound.md) (MEDIUM 45/100) - XML Element Type
