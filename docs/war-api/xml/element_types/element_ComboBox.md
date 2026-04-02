# ComboBox

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 186

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
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
| Addons seen in | TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChatTabTextEntry.xml:49`, `/workspace/data/raw/TidyChat/TidyChatTabTextEntry.xml:68`, `/workspace/data/raw/TidyChat/TidyChatTabWindows.xml:125`, `/workspace/data/raw/TidyChat/TidyChatTabWindows.xml:27`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.xml:56` |
| Namespaces detected | ComboBox |
| Source kinds | xml_frames, xml_handlers |
| Example locations | TidyChat: TChatTabTextEntryTemplateAnchorPointCombo, TidyChat: TChatTabTextEntryTemplateRelativeToCombo, TidyChat: TChatTabWindowsGroupTemplateScrollbarPositionCombo, TidyChat: TChatTabWindowsTemplateSelectWindowCombo, TidyRoll: AutoRollRowTemplateChoice |
| XML usage count | 5 |
| XML attribute usage count | 5 |
| Lua usage count | 1 |
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

Observed XML element type instantiated by 2 addons.

## Common Attributes

- inherits
- layer
- name

## Common Handlers

- [OnSelChanged](../handlers/handler_OnSelChanged.md)

## Common Handler Functions

- TidyChat.Options.UpdateGroupTabs
- TidyRoll.CustomAutoRoll.OnChoiceChange


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnSelChanged](../handlers/handler_OnSelChanged.md) | data | TidyChat.Options.UpdateGroupTabs, TidyRoll.CustomAutoRoll.OnChoiceChange | `index` | MEDIUM |

### Per-Event Lua API Calls

**OnSelChanged** handlers call: `ComboBoxGetSelectedMenuItem`, `LabelSetText`, `WindowGetId`

## Common Inherits

- EA_ComboBox_DefaultResizable
- EA_ComboBox_DefaultResizableSmall

## Common Parent Elements

- [Window](element_Window.md)

## Common Template Bases

- EA_ComboBox_DefaultResizable
- EA_ComboBox_DefaultResizableSmall


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 100% | EA_ComboBox_DefaultResizable, EA_ComboBox_DefaultResizableSmall |
| `layer` | **required** | 100% | popup |
## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `ComboBoxGetSelectedMenuItem` | ui | 2 | OnSelChanged |
| `LabelSetText` | ui | 1 | OnSelChanged |
| `WindowGetId` | ui | 1 | OnSelChanged |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnSelChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `index` | number | selected_index |

## Binding Resolution

- Total handler declarations: 2
- Resolved to Lua functions: 2 (100%)

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TChatTabTextEntryTemplateAnchorPointCombo -> ComboBox TChatTabTextEntryTemplateAnchorPointCombo
- TidyChat: TChatTabTextEntryTemplateRelativeToCombo -> ComboBox TChatTabTextEntryTemplateRelativeToCombo
- TidyChat: TChatTabWindowsGroupTemplateScrollbarPositionCombo -> ComboBox TChatTabWindowsGroupTemplateScrollbarPositionCombo
- TidyChat: TChatTabWindowsTemplateSelectWindowCombo -> ComboBox TChatTabWindowsTemplateSelectWindowCombo
- TidyRoll: AutoRollRowTemplateChoice -> ComboBox AutoRollRowTemplateChoice

## Related APIs

- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function

## Used With

- [OnSelChanged](../../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event
- [OnSelChanged](../handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none
