# ListData

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AggroMeter, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy, Pocket Palette, Shinies |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTexture.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainers.xml:0` |
| Namespaces detected | ListData |
| Source kinds | xml_frames |
| Example locations | AggroMeter: AggroMeterGrayListBox, Aura: AuraSettingsAuraList, Aura: AuraSharesAuraList, Aura: AuraTextureIconsIcons, BuffHead: BuffHeadSetupAdvancedCompressionItemWindowList, BuffHead: BuffHeadSetupAdvancedCompressionWindowList |
| XML usage count | 42 |
| XML attribute usage count | 42 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
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

ListData is a structural list-binding sub-element used inside list controls to connect XML definitions to Lua-backed row data.

## Common Attributes

- populationfunction
- table

## Common Inherits

- none

## Common Parent Elements

- [ListBox](element_ListBox.md) — 42× (HIGH)

## Common Structural Child Elements

- [ListColumns](element_ListColumns.md) — 25× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `table` | **required** | 100% | AggroMeter.Listdata, AuraShares.listDisplayData, AuraTexture.listIconDisplayData, AuraSettings.listDisplayData, ... |
| `populationfunction` | **required** | 97% | AuraShares.PopulateDisplay, AuraTexture.PopulateIconsListDisplay, AuraSettings.PopulateDisplay, BuffHead.Setup.OnPopulate, ... |
## Structural Sub-Elements

### [ListColumns](element_ListColumns.md)

Observed 25 times as an unnamed child.

## Recursive Hierarchy

- Root: [ListData](element_ListData.md)
- [ListColumns](element_ListColumns.md) (structural, 25×, HIGH)
  - [ListColumn](element_ListColumn.md) (structural, 42×, HIGH)

## Seen In

- AggroMeter
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Pocket Palette
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange

## Examples

- AggroMeter: AggroMeterGrayListBox -> ListData in ListBox AggroMeterGrayListBox
- Aura: AuraSettingsAuraList -> ListData in ListBox AuraSettingsAuraList
- Aura: AuraSharesAuraList -> ListData in ListBox AuraSharesAuraList
- Aura: AuraTextureIconsIcons -> ListData in ListBox AuraTextureIconsIcons
- BuffHead: BuffHeadSetupAdvancedCompressionItemWindowList -> ListData in ListBox BuffHeadSetupAdvancedCompressionItemWindowList
- BuffHead: BuffHeadSetupAdvancedCompressionWindowList -> ListData in ListBox BuffHeadSetupAdvancedCompressionWindowList

## Related APIs

- [ListBox](element_ListBox.md) (HIGH 100/100) - XML Element Type
- [ListColumns](element_ListColumns.md) (MEDIUM 55/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
