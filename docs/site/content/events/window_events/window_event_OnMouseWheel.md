# OnMouseWheel

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
| Addons seen in | CMap, MoraleCircle, TidyChat, TidyRoll |
| Files seen in | `/workspace/MoraleCircle/MoraleCircle.xml:21`, `/workspace/TidyChat/TidyChatCopy.xml:17`, `/workspace/TidyRoll/TidyRoll.xml:203`, `/workspace/cmap/CMap.xml:105`, `/workspace/cmap/CMap.xml:156`, `/workspace/cmap/CMap.xml:79` |
| Namespaces detected | OnMouseWheel |
| Source kinds | event_page, xml_handlers |
| Example locations | CMap: CMapWindowMapDisplay.OnMouseWheel, CMap: CMapWindowMapWorldMapButton.OnMouseWheel, CMap: CMapWindowWMap.OnMouseWheel, MoraleCircle: MoraleTemplate.OnMouseWheel, TidyChat: TidyChatCopy.OnMouseWheel, TidyRoll: TidyRollFrame.OnMouseWheel |
| XML usage count | 6 |
| XML attribute usage count | 6 |
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

Observed as an engine-supplied UI event hook used by 4 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from API_Ref alone.

## Seen In

- CMap
- MoraleCircle
- TidyChat
- TidyRoll

## Registrars And Handlers

- CMapWindow.MWheel
- CMapWindow.MWheelWholeZoom
- FrameManager.OnMouseWheel
- MoraleCircle.OnMouseWheel
- TidyChat.Copy.OnMouseWheel

## Examples

- CMap: CMapWindowMapDisplay -> CMapWindowMapDisplay.OnMouseWheel -> CMapWindow.MWheel
- CMap: CMapWindowMapWorldMapButton -> CMapWindowMapWorldMapButton.OnMouseWheel -> CMapWindow.MWheelWholeZoom
- CMap: CMapWindowWMap -> CMapWindowWMap.OnMouseWheel -> CMapWindow.MWheel
- MoraleCircle: MoraleTemplate -> MoraleTemplate.OnMouseWheel -> MoraleCircle.OnMouseWheel
- TidyChat: TidyChatCopy -> TidyChatCopy.OnMouseWheel -> TidyChat.Copy.OnMouseWheel
- TidyRoll: TidyRollFrame -> TidyRollFrame.OnMouseWheel -> FrameManager.OnMouseWheel

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../../xml/element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [OnMouseWheel](../../xml/handlers/handler_OnMouseWheel.md) (HIGH 100/100) - XML Handler
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- none
