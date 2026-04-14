# WAR API runtime facts

- Category: conventions
- Confidence: HIGH

## Description

Implementation-validated WAR engine API facts discovered through addon development and confirmed against the live engine.

## Involved APIs

- [Anchor](../xml/element_types/element_Anchor.md) (HIGH 100/100) - XML Element Type
- [Color](../xml/element_types/element_Color.md) (HIGH 100/100) - XML Element Type
- [Cursor](../globals/tables/table_Cursor.md) (HIGH 100/100) - Global Table
- [Cursor.Clear](../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Text](../xml/element_types/element_Text.md) (HIGH 100/100) - XML Element Type
- [wstring](../globals/tables/table_wstring.md) (HIGH 90/100) - Global Table
- [Icon](../xml/element_types/element_Icon.md) (MEDIUM 45/100) - XML Element Type
- [windowName](../events/game_events/game_event_windowName.md) (MEDIUM 43/100) - Game Event

## Flow Diagram

```text
Cursor.Clear <-> Cursor.IconOnCursor
```

## Example Code

```lua
`HIGH`: The standard pattern for a multi-line text-only tooltip is: Tooltips.CreateTextOnlyTooltip(windowName, nil) then Tooltips.SetTooltipText(row, col, wstring) for each line, Tooltips.SetTooltipColorDef(1, 1, Tooltips.COLOR_HEADING) for the heading row, then Tooltips.Finalize() and Tooltips.AnchorTooltip(Tooltips.ANCHOR_WINDOW_RIGHT). The MouseOverEnd handler should call Tooltips.ClearTooltip().
```

## Evidence

- `HIGH`: The standard pattern for a multi-line text-only tooltip is: Tooltips.CreateTextOnlyTooltip(windowName, nil) then Tooltips.SetTooltipText(row, col, wstring) for each line, Tooltips.SetTooltipColorDef(1, 1, Tooltips.COLOR_HEADING) for the heading row, then Tooltips.Finalize() and Tooltips.AnchorTooltip(Tooltips.ANCHOR_WINDOW_RIGHT). The MouseOverEnd handler should call Tooltips.ClearTooltip().
- Guidance: Use row=1, col=1 for the heading line with COLOR_HEADING applied. Use row=2+ for body lines. The row index is 1-based. Do NOT call Tooltips.AddLineToTooltip — that function does not exist and causes a nil-call runtime error.
- `HIGH`: Tooltips.AddLineToTooltip does not exist in the WAR engine. Calling it produces the runtime error: attempt to call field 'AddLineToTooltip' (a nil value).
- Guidance: Never call Tooltips.AddLineToTooltip. To add a separator or body line, call Tooltips.SetTooltipText with the next row index instead.
- `MEDIUM`: Tooltips.CreateAbilityTooltip(abilityData, windowName, anchor, hintWstring, nil) creates an ability tooltip anchored to a named window. The first argument is the table returned by GetAbilityData(id).
- Guidance: Pass the table from GetAbilityData(id) as the first argument. Use Tooltips.ANCHOR_WINDOW_RIGHT as anchor. The fifth argument appears unused (nil). Corroborate in additional addons before treating as a hard contract.
- `HIGH`: When an ability is dragged from the action bar: Cursor.IconOnCursor() returns true; Cursor.Data.Source equals Cursor.SOURCE_ACTION_LIST; Cursor.Data.ObjectId contains the ability ID as an integer. After handling the drop, Cursor.Clear() must be called to consume the cursor item.
- Guidance: Always check Cursor.IconOnCursor() first. Guard with Cursor.Data.Source == Cursor.SOURCE_ACTION_LIST before reading Cursor.Data.ObjectId. Call Cursor.Clear() after every drop handling path (success or rejection) to avoid leaving a stale cursor icon.
