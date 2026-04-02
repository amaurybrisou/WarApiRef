# OnHyperLinkRButtonUp

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
| Files seen in | `/workspace_addons/GuardLine/GuardLine.xml:210`, `/workspace_addons/GuardList/GuardList.xml:20`, `/workspace_addons/GuardRange/GuardRange.xml:20`, `/workspace_addons/MapPin/source/MapPin.xml:196`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:409`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:429`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:555`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:568` |
| Namespaces detected | OnHyperLinkRButtonUp |
| Source kinds | event_page, xml_handlers |
| Example locations | GuardLine: GuardLineSelfWindowLabel.OnHyperLinkRButtonUp, GuardList: GuardList_Window0Label.OnHyperLinkRButtonUp, GuardRange: GuardRange_Window0Label.OnHyperLinkRButtonUp, MapPin: MapPinCallTemplateWindowTitle.OnHyperLinkRButtonUp, RoR_SoR: RoR_SoR_City_Status_TemplateLock_STATUS.OnHyperLinkRButtonUp, RoR_SoR: RoR_SoR_City_Status_Template_TIMER.OnHyperLinkRButtonUp |
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

- EA_ChatWindow.OnHyperLinkRButtonUp
- MapPin.RButtonUp

## Examples

- GuardLine: GuardLineSelfWindowLabel -> GuardLineSelfWindowLabel.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
- GuardList: GuardList_Window0Label -> GuardList_Window0Label.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
- GuardRange: GuardRange_Window0Label -> GuardRange_Window0Label.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
- MapPin: MapPinCallTemplateWindowTitle -> MapPinCallTemplateWindowTitle.OnHyperLinkRButtonUp -> MapPin.RButtonUp
- RoR_SoR: RoR_SoR_City_Status_TemplateLock_STATUS -> RoR_SoR_City_Status_TemplateLock_STATUS.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
- RoR_SoR: RoR_SoR_City_Status_Template_TIMER -> RoR_SoR_City_Status_Template_TIMER.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp

## Related APIs

- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function

## Used With

- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [OnHyperLinkRButtonUp](../../xml/handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none

## Notes

- none
