# OnHyperLinkLButtonUp

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 118

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | GuardLine, GuardList, GuardRange, MapPin, RoR_SoR, wbLeadHelper |
| Files seen in | `/workspace/GuardLine/GuardLine.xml:209`, `/workspace/GuardList/GuardList.xml:19`, `/workspace/GuardRange/GuardRange.xml:19`, `/workspace/MapPin/source/MapPin.xml:195`, `/workspace/RoR_SoR/RoR_SoR.xml:409`, `/workspace/RoR_SoR/RoR_SoR.xml:429`, `/workspace/RoR_SoR/RoR_SoR.xml:555`, `/workspace/RoR_SoR/RoR_SoR.xml:568` |
| Namespaces detected | OnHyperLinkLButtonUp |
| Source kinds | event_page, xml_handlers |
| Example locations | GuardLine: GuardLineSelfWindowLabel.OnHyperLinkLButtonUp, GuardList: GuardList_Window0Label.OnHyperLinkLButtonUp, GuardRange: GuardRange_Window0Label.OnHyperLinkLButtonUp, MapPin: MapPinCallTemplateWindowTitle.OnHyperLinkLButtonUp, RoR_SoR: RoR_SoR_City_Status_TemplateLock_STATUS.OnHyperLinkLButtonUp, RoR_SoR: RoR_SoR_City_Status_Template_TIMER.OnHyperLinkLButtonUp |
| XML usage count | 10 |
| XML attribute usage count | 10 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
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

Observed as an engine-supplied UI event hook used by 6 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- GuardLine
- GuardList
- GuardRange
- MapPin
- RoR_SoR
- wbLeadHelper

## Registrars And Handlers

- EA_ChatWindow.OnHyperLinkLButtonUp

## Examples

- GuardLine: GuardLineSelfWindowLabel -> GuardLineSelfWindowLabel.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
- GuardList: GuardList_Window0Label -> GuardList_Window0Label.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
- GuardRange: GuardRange_Window0Label -> GuardRange_Window0Label.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
- MapPin: MapPinCallTemplateWindowTitle -> MapPinCallTemplateWindowTitle.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
- RoR_SoR: RoR_SoR_City_Status_TemplateLock_STATUS -> RoR_SoR_City_Status_TemplateLock_STATUS.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
- RoR_SoR: RoR_SoR_City_Status_Template_TIMER -> RoR_SoR_City_Status_Template_TIMER.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp

## Related APIs

- none

## Used With

- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [OnHyperLinkLButtonUp](../../xml/handlers/handler_OnHyperLinkLButtonUp.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none

## Notes

- none
