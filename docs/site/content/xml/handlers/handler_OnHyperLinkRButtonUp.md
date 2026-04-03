# OnHyperLinkRButtonUp

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 126

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | GuardLine, PartyCast, RoR_SoR |
| Files seen in | `/workspace/data/raw/GuardLine/GuardLine.xml:0`, `/workspace/data/raw/PartyCast/PartyCast.xml:0`, `/workspace/data/raw/RoR_SoR/RoR_SoR.xml:0` |
| Namespaces detected | OnHyperLinkRButtonUp |
| Source kinds | bindings, xml_handlers |
| Example locations | GuardLine: GuardLineSelfWindowLabel.OnHyperLinkRButtonUp, PartyCast: PartyCastWindow_TemplateTargetWindowLabel.OnHyperLinkRButtonUp, PartyCast: PartyCastWindow_TemplateTargetWindowLabelBG.OnHyperLinkRButtonUp, PartyCast: PartyCastWindow_Template_LargeTargetWindowLabel.OnHyperLinkRButtonUp, PartyCast: PartyCastWindow_Template_LargeTargetWindowLabelBG.OnHyperLinkRButtonUp, PartyCast: PartyCastWindow_Template_PlainName.OnHyperLinkRButtonUp |
| XML usage count | 17 |
| XML attribute usage count | 17 |
| Lua usage count | 17 |
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

Observed as an XML handler hook bound by 3 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Label

## Seen In

- GuardLine
- PartyCast
- RoR_SoR

## Examples

- GuardLine: GuardLineSelfWindowLabel -> GuardLineSelfWindowLabel.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
- PartyCast: PartyCastWindow_TemplateTargetWindowLabel -> PartyCastWindow_TemplateTargetWindowLabel.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
- PartyCast: PartyCastWindow_TemplateTargetWindowLabelBG -> PartyCastWindow_TemplateTargetWindowLabelBG.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
- PartyCast: PartyCastWindow_Template_LargeTargetWindowLabel -> PartyCastWindow_Template_LargeTargetWindowLabel.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
- PartyCast: PartyCastWindow_Template_LargeTargetWindowLabelBG -> PartyCastWindow_Template_LargeTargetWindowLabelBG.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
- PartyCast: PartyCastWindow_Template_PlainName -> PartyCastWindow_Template_PlainName.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp

## Related APIs

- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type

## Used With

- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
