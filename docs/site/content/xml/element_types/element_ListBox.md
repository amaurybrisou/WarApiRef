# ListBox

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
| Files seen in | `/workspace/data/raw/TidyChat/TidyChatLootRoll.xml:108`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.xml:228` |
| Namespaces detected | ListBox |
| Source kinds | xml_frames, xml_handlers |
| Example locations | TidyChat: TidyChatLootRollList, TidyRoll: TRollAutoRollList |
| XML usage count | 2 |
| XML attribute usage count | 2 |
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

- name
- rowcount
- rowdef
- rowspacing
- scrollbar
- visiblerows
- draganddrop
- layer

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)

## Common Handler Functions

- TidyRoll.CustomAutoRoll.OnListLbuttonUp


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | TidyRoll.CustomAutoRoll.OnListLbuttonUp | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonUp** handlers call: `Cursor.Clear`, `Cursor.IconOnCursor`

## Common Inherits

- none

## Common Parent Elements

- [Window](element_Window.md)

## Common Structural Child Elements

- [ListData](element_ListData.md) — 2× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `rowcount` | **required** | 100% | 50, 500 |
| `rowdef` | **required** | 100% | TidyChatLootRollRowTemplate, AutoRollRowTemplate |
| `rowspacing` | **required** | 100% | 0 |
| `scrollbar` | **required** | 100% | EA_ScrollBar_DefaultVerticalChain |
| `visiblerows` | **required** | 100% | 10, 7 |
| `draganddrop` | optional | 50% | true |
| `layer` | optional | 50% | secondary |
## Structural Sub-Elements

### [ListData](element_ListData.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `table` | **required** | TidyChat.LootRoll.itemRollData, TidyRoll.CustomAutoRoll.autoRollList |
| `populationfunction` | optional | TidyRoll.CustomAutoRoll.PopulateAutoRollList |
## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `Cursor.Clear` | ui | 1 | OnLButtonUp |
| `Cursor.IconOnCursor` | data | 1 | OnLButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |

## Binding Resolution

- Total handler declarations: 1
- Resolved to Lua functions: 1 (100%)

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChatLootRollList -> ListBox TidyChatLootRollList
- TidyRoll: TRollAutoRollList -> ListBox TRollAutoRollList

## Related APIs

- [ListData](element_ListData.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Affects

- none
