# EA_Label_ChatText

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
| Addons seen in | BlackBook, EA_OpenPartyWindow, FozAuction, HealGrid, Kwestor, SocialWindow 2.0 |
| Files seen in | BlackBook.xml, Gui/HealGridGui.xml, Gui/HealGridGuiSpellList.xml, Gui/HealGridGuiTabMouseClick.xml, Gui/HealGridGuiTabSpellTrack.xml, Gui/KwestorGuiTabArea.xml, Source/AuctionWindowSearchControls.xml, Source/OpenPartyWindowTabNearby.xml |
| Namespaces detected | EA_Label_ChatText |
| Source kinds | xml_attributes |
| Example locations | AuctionSearchControlsRankSeparator, AuctionWindowControlLabel, BlackBookWindow_EntryRowText, EA_Label_SocialListItem, EA_Template_OpenPartyGroupLineRatioText, EA_Template_OpenPartyGroupLineTimedDistanceText |
| XML usage count | 35 |
| XML attribute usage count | 35 |
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

Engine-supplied XML constant or template class referenced by 6 addons.

## Seen In

- BlackBook
- EA_OpenPartyWindow
- FozAuction
- HealGrid
- Kwestor
- SocialWindow 2.0

## Used By

- AuctionSearchControlsRankSeparator
- AuctionWindowControlLabel
- BlackBookWindow_EntryRowText
- EA_Label_SocialListItem
- EA_Template_OpenPartyGroupLineRatioText
- EA_Template_OpenPartyGroupLineTimedDistanceText
- EA_Template_OpenPartyWorldGroupLineLocationText
- EA_Template_OpenPartyWorldGroupLineRatioText
- HGG_HealGridMenuItemTemplateLabel
- HGG_SpellListRowTemplateName
- HGG_TabMouseClickListRowTemplateAction
- HGG_TabMouseClickListRowTemplateActionDetail
- HGG_TabMouseClickListRowTemplateFrameType
- HGG_TabMouseClickListRowTemplateModifierKey
- HGG_TabMouseClickListRowTemplateMouseButton
- HGG_TabMouseClickListRowTemplateTargetType
- HGG_TabSpellTrackListRowTemplateDuration
- HGG_TabSpellTrackListRowTemplateLabelIndex
- HGG_TabSpellTrackListRowTemplateName
- HGG_TabSpellTrackListRowTemplateSignOther
- HGG_TabSpellTrackListRowTemplateSignOtherMultiple
- HGG_TabSpellTrackListRowTemplateSignSelf
- HGG_TabSpellTrackListRowTemplateSignSelfMultiple
- HGG_TabSpellTrackListRowTemplateWarnExpire
- HGG_TabSpellTrackListRowTemplateWarnExpireSign
- KwestorGui_TabAreaListRowTemplateAreaName
- KwestorGui_TabAreaListRowTemplateRvRStatus
- KwestorGui_TabAreaListRowTemplateSeen
- KwestorGui_TabAreaListRowTemplateZoneName
- SocialWindowTabOptionsSocketAFKNoteLabel
- SocialWindowTabOptionsSocketAdvisorDescLabel
- SocialWindowTabOptionsSocketAnonymousDescLabel
- SocialWindowTabOptionsSocketDisableBuddyListDescLabel
- SocialWindowTabOptionsSocketHiddenDescLabel
- SocialWindowTabOptionsSocketPrivatePartyDescLabel

## Related APIs

- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type

## Notes

- none
