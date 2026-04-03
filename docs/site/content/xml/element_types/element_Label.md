# Label

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, Moth, PartyCast, Soloq, TidyChat, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.xml:101`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:129`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:135`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:163`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:169`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:197`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:203`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:21` |
| Namespaces detected | Label |
| Source kinds | xml_frames, xml_handlers |
| Example locations | InfoScroller: InfoScrollerMainWindowLabel1, InfoScroller: InfoScrollerTemplateLabel1, InfoScroller: InfoScrollerTemplateLabel1BG, InfoScroller: InfoScrollerTemplateLabel2, InfoScroller: InfoScrollerTemplateLabel2BG, InfoScroller: InfoScrollerTemplateLabel3 |
| XML usage count | 117 |
| XML attribute usage count | 117 |
| Lua usage count | 2 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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

Observed XML element type instantiated by 7 addons.

## Common Attributes

- name
- font
- textalign
- layer
- handleinput
- autoresize
- wordwrap
- warnOnTextCropped
- popable
- maxchars
- inherits
- textAutoFitMinScale

## Common Handlers

- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md)
- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md)

## Common Handler Functions

- EA_ChatWindow.OnHyperLinkRButtonUp
- EA_ChatWindow.OnHyperLinkLButtonUp


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) | custom | EA_ChatWindow.OnHyperLinkLButtonUp | `function(...)` | LOW |
| [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) | custom | EA_ChatWindow.OnHyperLinkRButtonUp | `function(...)` | LOW |

## Common Inherits

- TChatLabel
- TRollLabel
- EA_Label_DefaultText

## Common Parent Elements

- [Window](element_Window.md)
- [Button](element_Button.md)

## Common Template Bases

- EA_Label_DefaultText
- TChatLabel
- TRollLabel


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `font` | **required** | 84% | font_clear_medium, font_clear_small_bold, font_chat_text, font_clear_tiny, ... |
| `textalign` | **required** | 80% | left, center, bottom, right |
| `layer` | optional | 72% | overlay, secondary, default, background, ... |
| `handleinput` | optional | 71% | false, true |
| `autoresize` | optional | 63% | true, false |
| `wordwrap` | optional | 38% | true, false |
| `warnOnTextCropped` | optional | 37% | false |
| `popable` | optional | 34% | false |
| `maxchars` | optional | 29% | 1024, 72, 100, 16, ... |
| `inherits` | optional | 20% | EA_Label_DefaultText, TChatLabel, TRollLabel |
| `textAutoFitMinScale` | optional | 18% | 0.3 |
| `autoresizewidth` | optional | 12% | true |
| `sticky` | optional | 9% | false, true |
| `autoresizeheight` | optional | 7% | true |
| `ellipsisOnCrop` | optional | 3% | true |
| `fontscale` | optional | 1% | 0.5 |
| `scale` | optional | 1% | 0.5 |
| `text` | optional | 0% |  |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnHyperLinkLButtonUp

Confidence: LOW

### OnHyperLinkRButtonUp

Confidence: LOW

## Lua Functions Manipulating This Type

- TidyRollOptions.Initialize


## Binding Resolution

- Total handler declarations: 37
- Resolved to Lua functions: 0 (0%)

## Seen In

- InfoScroller
- Moth
- PartyCast
- Soloq
- TidyChat
- TidyRoll
- minesweep

## Examples

- InfoScroller: InfoScrollerMainWindowLabel1 -> Label InfoScrollerMainWindowLabel1
- InfoScroller: InfoScrollerTemplateLabel1 -> Label InfoScrollerTemplateLabel1
- InfoScroller: InfoScrollerTemplateLabel1BG -> Label InfoScrollerTemplateLabel1BG
- InfoScroller: InfoScrollerTemplateLabel2 -> Label InfoScrollerTemplateLabel2
- InfoScroller: InfoScrollerTemplateLabel2BG -> Label InfoScrollerTemplateLabel2BG
- InfoScroller: InfoScrollerTemplateLabel3 -> Label InfoScrollerTemplateLabel3

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type

## Used With

- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) (HIGH 100/100) - XML Event
- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) (HIGH 100/100) - XML Event
- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Event

## Affects

- none
