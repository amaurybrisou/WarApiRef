# EA_FullResizeImage_DefaultFrame

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
| Addons seen in | Aura, CCTV, CDown, DAoCBuff, EA_ScenarioGroupWindow, HealGrid, Moth, PeaceOut |
| Files seen in | CCTV.xml, CDownSettings.xml, Gui/HealGridGuiTabMouseClick.xml, Gui/HealGridGuiTabSpellTrack.xml, Moth.xml, PeaceOut.xml, Source/DAoCBuffSettingsTabs.xml, Source/MacroWindow.xml |
| Namespaces detected | EA_FullResizeImage_DefaultFrame |
| Source kinds | xml_attributes |
| Example locations | AddFriendDescriptionFrame, AuraWindowBackground, CCTV_BG_TEMPLATEBorderDecorative, CDownSectionBackground, DAoCBuffFrameSettingsTab_ScrollChildFilterFrame, DAoCBuffFrameSettingsTab_ScrollChildFrameSettingsFrame |
| XML usage count | 30 |
| XML attribute usage count | 30 |
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

Engine-supplied XML constant or template class referenced by 11 addons.

## Seen In

- Aura
- CCTV
- CDown
- DAoCBuff
- EA_ScenarioGroupWindow
- HealGrid
- Moth
- PeaceOut
- SocialWindow 2.0
- XpStatus+G
- bigger_MacroWindow

## Used By

- AddFriendDescriptionFrame
- AuraWindowBackground
- CCTV_BG_TEMPLATEBorderDecorative
- CDownSectionBackground
- DAoCBuffFrameSettingsTab_ScrollChildFilterFrame
- DAoCBuffFrameSettingsTab_ScrollChildFrameSettingsFrame
- DAoCBuffFrameSettingsTab_ScrollChildLayoutFrame
- DAoCBuffFrameSettingsTab_ScrollChildSortFrame
- DAoCBuffListManagerTab_ScrollChildAddDeleteFrame
- DAoCBuffListManagerTab_ScrollChildManagerFrame
- DAoCBuff_Settings_FilterFrame_ScrollChildActionFrame
- DAoCBuff_Settings_FilterFrame_ScrollChildConditionFrame
- DAoCBuff_Settings_FilterFrame_ScrollChildEnvironmentFrame
- DAoCBuff_Settings_FilterFrame_ScrollChildMiscFrame
- EA_Window_MacroMacrosBackground
- HGG_TabMouseClickWindowTemplateEditBorder
- HGG_TabSpellTrackWindowTemplateEditBorder
- MothBorderDecorative
- PeaceOutFrame
- ScenarioGroupBottomBarBG
- ScenarioGroupMiddleBarBG
- SingleSGroupTemplateBG
- SingleSGroupTemplateNameBG
- SocialWindowTabFriendsSocketTopHalfBorder
- SocialWindowTabIgnoreSocketTopHalfBorder
- SocialWindowTabOptionsSocketAFKBackground
- SocialWindowTabOptionsSocketOptionsBackground
- SocialWindowTabSearchSocketOptionsBackground
- SocialWindowTabSearchSocketTopHalfBorder
- XpStatusQuotaWindowFrame

## Related APIs

- [FullResizeImage](../../xml/element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type

## Notes

- none
