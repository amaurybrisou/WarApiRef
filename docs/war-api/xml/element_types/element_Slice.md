# Slice

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
| Addons seen in | CMap, EA_LoadingScreen, EA_ThreePartBar, FlagCap, GDes, Ges, KillStreak, MapPin |
| Files seen in | Assets/Assets.xml, Textures/EA_VictoryPoints01_32b.xml, Textures/Textures.xml, source/MapPin.xml, source/nLootLinkTemplates.xml |
| Namespaces detected | Slice |
| Source kinds | xml_frames |
| Example locations | CMap: EA_Scenario01_32b, EA_LoadingScreen: EA_Texture_Racial_Filigree_Rules, EA_ThreePartBar: EA_VictoryPoints01_32b, FlagCap: FlagCapTexture, GDes: GDesBackTexture, GDes: GDesForeTexture |
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

Slice is a XML UI element. It commonly appears under Texture.

## Common Attributes

- height
- id
- left
- top
- width

## Common Inherits

- none

## Common Parent Elements

- [Texture](element_Texture.md) — 287× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `height` | **required** | 100% | 46, 10, 42, 19, ... |
| `id` | **required** | 100% | banner-base, banner-hover, banner-press, scenario-base, ... |
| `left` | **required** | 100% | 0, 46, 92, 10, ... |
| `top` | **required** | 100% | 46, 0, 10, 42, ... |
| `width` | **required** | 100% | 46, 10, 256, 11, ... |
## Seen In

- CMap
- EA_LoadingScreen
- EA_ThreePartBar
- FlagCap
- GDes
- Ges
- KillStreak
- MapPin
- Pure Careerbar
- RVMOD_3DPortrait
- RVMOD_SquaredDistances
- Res
- Shinies
- SpamBayes
- TexturedButtons
- TurretScrap
- nLootLink

## Examples

- CMap: EA_Scenario01_32b -> Slice in Texture EA_Scenario01_32b
- EA_LoadingScreen: EA_Texture_Racial_Filigree_Rules -> Slice in Texture EA_Texture_Racial_Filigree_Rules
- EA_ThreePartBar: EA_VictoryPoints01_32b -> Slice in Texture EA_VictoryPoints01_32b
- FlagCap: FlagCapTexture -> Slice in Texture FlagCapTexture
- GDes: GDesBackTexture -> Slice in Texture GDesBackTexture
- GDes: GDesForeTexture -> Slice in Texture GDesForeTexture

## Related APIs

- [Texture](element_Texture.md) (HIGH 100/100) - XML Element Type
