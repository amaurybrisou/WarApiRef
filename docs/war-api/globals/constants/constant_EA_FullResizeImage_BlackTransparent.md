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
| Addons seen in | AggroMeter, CMap, MapMonster, RVAPI_ColorDialog, TidyChat, TidyRoll |
| Files seen in | `/workspace/AggroMeter/AggroMeter.xml:158`, `/workspace/AggroMeter/AggroMeter.xml:288`, `/workspace/AggroMeter/AggroMeter.xml:344`, `/workspace/MapMonster/Source/MapMonster_Calibrate.xml:56`, `/workspace/MapMonster/Source/MapMonster_Calibrate.xml:71`, `/workspace/RVAPI_ColorDialog/RVAPI_ColorDialog.xml:110`, `/workspace/TidyChat/TidyChat.xml:21`, `/workspace/TidyRoll/CustomAutoRoll.xml:107` |
| Namespaces detected | EA_FullResizeImage_BlackTransparent |
| Source kinds | xml_attributes |
| Example locations | AggroMeterWindowBorderCheck, AggroMeterWindow_AggroWindow3BorderCheck, AggroMeterWindow_AggroWindow5BorderCheck, CMapWindowPinFilterMenuBackground, MapMonster_CalibrateWindowBackground, MapMonster_CalibrateWindowClickLayer |
| XML usage count | 9 |
| XML attribute usage count | 9 |
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

Observed engine XML template or inherited constant referenced by 6 addons.

## Seen In

- AggroMeter
- CMap
- MapMonster
- RVAPI_ColorDialog
- TidyChat
- TidyRoll

## Used By

- AggroMeterWindowBorderCheck
- AggroMeterWindow_AggroWindow3BorderCheck
- AggroMeterWindow_AggroWindow5BorderCheck
- CMapWindowPinFilterMenuBackground
- MapMonster_CalibrateWindowBackground
- MapMonster_CalibrateWindowClickLayer
- RVAPI_ColorDialogWindowBackgroundBackground
- TRollAutoRollBackground
- TidyChatOptionsBackground

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
