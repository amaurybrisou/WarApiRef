# ListColumns

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
| Addons seen in | AggroMeter, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, EA_UiModWindow, Enemy, Pocket Palette |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.xml:131`, `/workspace_addons/Aura/Source/AuraSettings.xml:189`, `/workspace_addons/Aura/Source/AuraShares.xml:224`, `/workspace_addons/BuffHead/Setup/General.xml:51`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompression.xml:106`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItem.xml:182`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainers.xml:106`, `/workspace_addons/BuffHead/Setup/SetupEffectCache.xml:163` |
| Namespaces detected | ListColumns |
| Source kinds | xml_frames |
| Example locations | AggroMeter: AggroMeterGrayListBox, Aura: AuraSettingsAuraList, Aura: AuraSharesAuraList, BuffHead: BuffHeadSetupAdvancedCompressionItemWindowList, BuffHead: BuffHeadSetupAdvancedCompressionWindowList, BuffHead: BuffHeadSetupAdvancedContainersWindowList |
| XML usage count | 28 |
| XML attribute usage count | 28 |
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

Structural XML container inside ListBox that groups one or more ListColumn declarations, mapping backing-table fields to row-template child windows.

## Common Attributes

- none

## Common Inherits

- none

## Seen In

- AggroMeter
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- EA_UiModWindow
- Enemy
- Pocket Palette
- QuickTacticSwitch
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- Tortall_DPS
- TurretRange

## Examples

- AggroMeter: AggroMeterGrayListBox -> ListColumns in ListBox AggroMeterGrayListBox
- Aura: AuraSettingsAuraList -> ListColumns in ListBox AuraSettingsAuraList
- Aura: AuraSharesAuraList -> ListColumns in ListBox AuraSharesAuraList
- BuffHead: BuffHeadSetupAdvancedCompressionItemWindowList -> ListColumns in ListBox BuffHeadSetupAdvancedCompressionItemWindowList
- BuffHead: BuffHeadSetupAdvancedCompressionWindowList -> ListColumns in ListBox BuffHeadSetupAdvancedCompressionWindowList
- BuffHead: BuffHeadSetupAdvancedContainersWindowList -> ListColumns in ListBox BuffHeadSetupAdvancedContainersWindowList

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
