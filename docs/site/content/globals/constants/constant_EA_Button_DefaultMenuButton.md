# EA_Button_DefaultMenuButton

- Category: Constant
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
| Addons seen in | AdjustTheTip, BuffHead, EA_OpenPartyWindow, FozAuction, LibAddonButton, MapMonster, MarkBuff, Miracle Grow Remix |
| Files seen in | AdjustTheTip.xml, Core.xml, Gui.xml, MGRemix.xml, PartyAdWindow.xml, ScenarioStats.xml, Setup/General.xml, Setup/SetupDisplay.xml |
| Namespaces detected | EA_Button_DefaultMenuButton |
| Source kinds | xml_attributes |
| Example locations | AdjustTheTipMenuItemSlider, AuctionWindowContextMenuItem, BuffHeadContextMenuItemFontSelection, EA_Template_OpenParty_ComboBoxMenuButton, LibAddonButtonMenuItemTemplate, MapFilterContextMenuChoice |
| XML usage count | 16 |
| XML attribute usage count | 16 |
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

Engine-supplied XML constant or template class referenced by 14 addons.

## Seen In

- AdjustTheTip
- BuffHead
- EA_OpenPartyWindow
- FozAuction
- LibAddonButton
- MapMonster
- MarkBuff
- Miracle Grow Remix
- Obsidian
- PartyAd
- ScenarioStats
- SocialWindow 2.0
- TexturedButtons
- TurretRange

## Used By

- AdjustTheTipMenuItemSlider
- AuctionWindowContextMenuItem
- BuffHeadContextMenuItemFontSelection
- EA_Template_OpenParty_ComboBoxMenuButton
- LibAddonButtonMenuItemTemplate
- MapFilterContextMenuChoice
- MarkBuffContextMenuItemBuffCheckBox
- MiracleGrow2ContextItem
- MiracleGrow2ContextItemCIT
- NewCenterAligned_CBMenubutton
- ObsidianContextMenuItemFontSelection
- PartyAd_PurposePresetComboBoxMenuButton
- PartyAd_SelfRoleComboBoxMenuButton
- SocialContextMenuItemWithMenu
- TexturedButtonsContextMenuItemFontSelection
- TurretRangeContextMenuItemFontSelection

## Related APIs

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- none
