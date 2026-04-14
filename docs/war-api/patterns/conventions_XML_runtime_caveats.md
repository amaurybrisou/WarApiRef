# XML runtime caveats

- Category: conventions
- Confidence: MEDIUM

## Description

Implementation-validated findings show that XML input, anchoring, and scroll layout behavior can depend on ancestor state, stable parent containers, and outer-window sizing even when child nodes appear correctly configured.

## Involved APIs

- [Anchor](../xml/element_types/element_Anchor.md) (HIGH 100/100) - XML Element Type
- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Cursor](../globals/tables/table_Cursor.md) (HIGH 100/100) - Global Table
- [Cursor.IconOnCursor](../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [EA_EditBox_DefaultFrame](../globals/constants/constant_EA_EditBox_DefaultFrame.md) (HIGH 100/100) - Constant
- [EA_Window_Default](../globals/constants/constant_EA_Window_Default.md) (HIGH 100/100) - Constant
- [EA_Window_DefaultButtonBottomFrame](../globals/constants/constant_EA_Window_DefaultButtonBottomFrame.md) (HIGH 100/100) - Constant
- [EditBox](../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [Label](../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [LabelGetText](../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [LabelGetTextDimensions](../window_api/functions/window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetText](../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [ListBox](../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [ScrollWindow](../xml/element_types/element_ScrollWindow.md) (HIGH 100/100) - XML Element Type
- [ScrollWindowUpdateScrollRect](../window_api/functions/window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [Size](../xml/element_types/element_Size.md) (HIGH 100/100) - XML Element Type
- [Text](../xml/element_types/element_Text.md) (HIGH 100/100) - XML Element Type
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../events/window_events/window_event_OnLButtonUp.md) (MEDIUM 53/100) - Window Event
- [Icon](../xml/element_types/element_Icon.md) (MEDIUM 45/100) - XML Element Type
- [Bottom](../xml/element_types/element_Bottom.md) (MEDIUM 30/100) - XML Element Type

## Flow Diagram

```text
OnLButtonUp
  -> handlers: AuctionWindow.Hide, EA_Window_OpenParty.ToggleFullWindow, ScenarioGroupWindow.LeaveGroup
  -> ui: Button
```

## Example Code

```lua
WhoHealedMe: a child `OnLButtonUp` target remained inert until the parent or template input chain was made input-enabled.
```

## Evidence

- WhoHealedMe: a child `OnLButtonUp` target remained inert until the parent or template input chain was made input-enabled.
- guidance: validate ancestor `handleinput` state across the clickable chain, not only on the child node.
- caveat: treat this as a reusable runtime warning, not a guaranteed engine contract.
- WhoHealedMe: nested scroll content dimensions initially under-reported usable space during early layout.
- guidance: compute size from the outer parent first, then resize child content and call `ScrollWindowUpdateScrollRect`.
- `MEDIUM`: early layout reads on nested scroll content can under-report the eventual usable region.
- When first-render list height looks collapsed, compute dimensions from the outer parent first, then resize child content and refresh the scroll rect.
- `MEDIUM`: When footer controls or edit boxes drift under sibling-based anchoring, anchoring them directly to a stable parent container is more reliable than chaining them through neighboring controls.
- Guidance: If a control renders in the wrong column or outside the window, clear its anchors and re-anchor it from `$parent` or a stable container such as `$parentButtonBackground` with explicit offsets.
- `MEDIUM`: Text-driven header or form rows are more stable when label dimensions are measured after LabelSetText and adjacent controls are positioned from the measured width.
- Guidance: Use `LabelGetTextDimensions()` and then resize or re-anchor neighboring labels and edit boxes to avoid truncation, overlap, or misplaced values.
- `MEDIUM`: A ListBox row highlight or tint only spans the width of the row template; widening the list box alone does not widen the tinted row background.
