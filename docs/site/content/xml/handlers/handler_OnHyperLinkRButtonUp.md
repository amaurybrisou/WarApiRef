# OnHyperLinkRButtonUp

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 138

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | GuardLine, GuardList, GuardRange, MapPin, RoR_SoR, wbLeadHelper |
| Files seen in | `/workspace_addons/GuardLine/GuardLine.xml:210`, `/workspace_addons/GuardList/GuardList.xml:20`, `/workspace_addons/GuardRange/GuardRange.xml:20`, `/workspace_addons/MapPin/source/MapPin.xml:196`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:409`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:429`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:555`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:568` |
| Namespaces detected | OnHyperLinkRButtonUp |
| Source kinds | bindings, xml_handlers |
| Example locations | GuardLine: GuardLineSelfWindowLabel.OnHyperLinkRButtonUp, GuardList: GuardList_Window0Label.OnHyperLinkRButtonUp, GuardRange: GuardRange_Window0Label.OnHyperLinkRButtonUp, MapPin: MapPinCallTemplateWindowTitle.OnHyperLinkRButtonUp, RoR_SoR: RoR_SoR_City_Status_TemplateLock_STATUS.OnHyperLinkRButtonUp, RoR_SoR: RoR_SoR_City_Status_Template_TIMER.OnHyperLinkRButtonUp |
| XML usage count | 10 |
| XML attribute usage count | 10 |
| Lua usage count | 10 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
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

Observed as an XML handler hook bound by 6 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Label

## Seen In

- GuardLine
- GuardList
- GuardRange
- MapPin
- RoR_SoR
- wbLeadHelper

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

- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [OnHyperLinkRButtonUp](../../events/window_events/window_event_OnHyperLinkRButtonUp.md) (HIGH 100/100) - Window Event

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
