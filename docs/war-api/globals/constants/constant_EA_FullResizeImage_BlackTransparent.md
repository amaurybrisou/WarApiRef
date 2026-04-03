# EA_FullResizeImage_BlackTransparent

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
| Addons seen in | AggroMeter, CMap, DPSMeter, DuffTimer, MapMonster, Minmap, RVAPI_ColorDialog, RaidMeter |
| Files seen in | AggroMeter.xml, CMap.xml, CustomAutoRoll.xml, DPSMeterWindow.xml, DuffTimer.xml, RVAPI_ColorDialog.xml, RaidMeter.xml, ReferList.xml |
| Namespaces detected | EA_FullResizeImage_BlackTransparent |
| Source kinds | xml_attributes |
| Example locations | AddFriendDescriptionBackground, AggroMeterWindowBorderCheck, AggroMeterWindow_AggroWindow3BorderCheck, AggroMeterWindow_AggroWindow5BorderCheck, CMapWindowPinFilterMenuBackground, DPSMeterRowTemplateBackground |
| XML usage count | 24 |
| XML attribute usage count | 24 |
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

Engine-supplied XML constant or template class referenced by 15 addons.

## Seen In

- AggroMeter
- CMap
- DPSMeter
- DuffTimer
- MapMonster
- Minmap
- RVAPI_ColorDialog
- RaidMeter
- Refer
- ResHelp
- SocialWindow 2.0
- TidyChat
- TidyRoll
- XpStatus+G
- emotes

## Used By

- AddFriendDescriptionBackground
- AggroMeterWindowBorderCheck
- AggroMeterWindow_AggroWindow3BorderCheck
- AggroMeterWindow_AggroWindow5BorderCheck
- CMapWindowPinFilterMenuBackground
- DPSMeterRowTemplateBackground
- DuffTimer_Bar_StatusBarBG
- DuffTimer_IconBar_StatusBarBG
- DuffTimer_ReverseBar_StatusBarBG
- MapMonster_CalibrateWindowBackground
- MapMonster_CalibrateWindowClickLayer
- MinmapPinMenuBackground
- MinmapScenarioMenuBackground
- RVAPI_ColorDialogWindowBackgroundBackground
- RaidMeterWindowBorderCheck
- ReferPlayerBox_TemplateBG
- ResHelpWindowBorderCheck
- ResHelpWindowBorderCheck2
- SocialWindowListWindowDetailBackground
- TRollAutoRollBackground
- Template_RaidWindowBorderCheck
- TidyChatOptionsBackground
- XpStatusQuotaWindowBackground
- emotesWindowBackground

## Related APIs

- [FullResizeImage](../../xml/element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type

## Notes

- none
