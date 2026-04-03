# OnMouseWheel

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 106

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | MoraleCircle, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/MoraleCircle/MoraleCircle.xml:0`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:0`, `/workspace/data/raw/TidyRoll/TidyRoll.xml:0` |
| Namespaces detected | OnMouseWheel |
| Source kinds | event_page, xml_handlers |
| Example locations | MoraleCircle: MoraleTemplate.OnMouseWheel, TidyChat: TidyChatCopy.OnMouseWheel, TidyRoll: TidyRollFrame.OnMouseWheel |
| XML usage count | 3 |
| XML attribute usage count | 3 |
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

Observed as an engine-supplied UI event hook used by 3 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- MoraleCircle
- TidyChat
- TidyRoll

## Registrars And Handlers

- FrameManager.OnMouseWheel
- MoraleCircle.OnMouseWheel
- TidyChat.Copy.OnMouseWheel

## Examples

- MoraleCircle: MoraleTemplate -> MoraleTemplate.OnMouseWheel -> MoraleCircle.OnMouseWheel
- TidyChat: TidyChatCopy -> TidyChatCopy.OnMouseWheel -> TidyChat.Copy.OnMouseWheel
- TidyRoll: TidyRollFrame -> TidyRollFrame.OnMouseWheel -> FrameManager.OnMouseWheel

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
