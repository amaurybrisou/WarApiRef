# WAR API Facts Seed

## Intent

Normalized snippets of high-value platform facts for tooling and MCP context.

## Fact Style

- Keep facts atomic and verifiable.
- Prefer canonical symbol names as seen in generated docs.
- Attach confidence labels when facts are inferred.

## Seeded Baselines

- Many runtime collections are sparse in practice; iterate defensively.
- UI behavior frequently depends on XML configuration plus Lua runtime calls.
- Symbol confidence should distinguish source confidence from semantic confidence.

## Confidence Tags

- `HIGH`: directly evidenced in canonical docs and repeated usage
- `MEDIUM`: partial evidence or relation-based inference
- `LOW`: weak evidence; treat as provisional

<!-- OBSERVATION:apa_tooltips_namespace_text_pattern (promoted:2026-04-14T05:15:34Z) -->
> Source: AdvancedPetAssist | Confidence: HIGH | Promoted: 2026-04-14

- `HIGH`: The standard pattern for a multi-line text-only tooltip is: Tooltips.CreateTextOnlyTooltip(windowName, nil) then Tooltips.SetTooltipText(row, col, wstring) for each line, Tooltips.SetTooltipColorDef(1, 1, Tooltips.COLOR_HEADING) for the heading row, then Tooltips.Finalize() and Tooltips.AnchorTooltip(Tooltips.ANCHOR_WINDOW_RIGHT). The MouseOverEnd handler should call Tooltips.ClearTooltip().
  - Guidance: Use row=1, col=1 for the heading line with COLOR_HEADING applied. Use row=2+ for body lines. The row index is 1-based. Do NOT call Tooltips.AddLineToTooltip — that function does not exist and causes a nil-call runtime error.
- `HIGH`: Tooltips.AddLineToTooltip does not exist in the WAR engine. Calling it produces the runtime error: attempt to call field 'AddLineToTooltip' (a nil value).
  - Guidance: Never call Tooltips.AddLineToTooltip. To add a separator or body line, call Tooltips.SetTooltipText with the next row index instead.
- `MEDIUM`: Tooltips.CreateAbilityTooltip(abilityData, windowName, anchor, hintWstring, nil) creates an ability tooltip anchored to a named window. The first argument is the table returned by GetAbilityData(id).
  - Guidance: Pass the table from GetAbilityData(id) as the first argument. Use Tooltips.ANCHOR_WINDOW_RIGHT as anchor. The fifth argument appears unused (nil). Corroborate in additional addons before treating as a hard contract.

<!-- OBSERVATION:apa_cursor_data_drag_drop_pattern (promoted:2026-04-14T05:15:34Z) -->
> Source: AdvancedPetAssist | Confidence: HIGH | Promoted: 2026-04-14

- `HIGH`: When an ability is dragged from the action bar: Cursor.IconOnCursor() returns true; Cursor.Data.Source equals Cursor.SOURCE_ACTION_LIST; Cursor.Data.ObjectId contains the ability ID as an integer. After handling the drop, Cursor.Clear() must be called to consume the cursor item.
  - Guidance: Always check Cursor.IconOnCursor() first. Guard with Cursor.Data.Source == Cursor.SOURCE_ACTION_LIST before reading Cursor.Data.ObjectId. Call Cursor.Clear() after every drop handling path (success or rejection) to avoid leaving a stale cursor icon.
