# Conventions

## Initialization pattern

- Confidence: HIGH

- Description: Addons consistently move from manifest loading into initialize hooks, then into XML creation and runtime event registration.

- Evidence:

- none

## Event registration pattern

- Confidence: HIGH

- Description: Runtime events are typically wired through RegisterEventHandler or window-scoped variants, with handler names staying module-qualified.

- Evidence:

- RegisterEventHandler: AbilityAlert: RegisterEventHandler(SystemData.Events.WORLD_OBJ_COMBAT_EVENT, "AbilityAlert.CombatEvent")
  - RegisterEventHandler: AbilityNotifier: RegisterEventHandler(SystemData.Events.WORLD_OBJ_COMBAT_EVENT, "AbHelp.CombatEvent")
  - RegisterEventHandler: AbilityNotifier: RegisterEventHandler(SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED, "AbHelp.CombatUpdated")
  - RegisterEventHandler: ActionBarHide: RegisterEventHandler(SystemData.Events.LOADING_END, "ActionBarHide.OnLoad")
  - RegisterEventHandler: ActionBarHide: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ActionBarHide.OnLoad")
  - RegisterEventHandler: ActionBarHide: RegisterEventHandler(SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED, "ActionBarHide.Combat")
  - UnregisterEventHandler: Ace: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
  - UnregisterEventHandler: ActionFraction: UnregisterEventHandler(SystemData.Events.ENTER_WORLD, "ActionFraction.Initialize")

## UI creation pattern

- Confidence: HIGH

- Description: UI is commonly created from XML, then positioned in Lua through CreateWindow or CreateWindowFromTemplate followed by anchor calls.

- Evidence:

- Window creation: AbilityAlert: CreateWindow("AAWindow", true)
  - Window creation: AbilityNotifier: CreateWindow("AbHelpWindow", true)
  - Window creation: ActionFraction: CreateWindow(windowName, true)
  - Window creation: ActionPoints: CreateWindow("ActionPointsWindow", true)
  - Window creation: AdvancedPetAssist: CreateWindow("APAOptions", true)
  - Window creation: AdvancedRenownTrainer: CreateWindow("AdvancedRenownTrainingPresetsWindow", false)
  - Template instantiation: Ace: CreateWindowFromTemplate(w.name, base, w.parent)
  - Template instantiation: Ace: CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)

## XML to Lua binding pattern

- Confidence: HIGH

- Description: XML handler names map directly to Lua functions and can be cross-checked through the bindings page.

- Evidence:

- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: .OnSelChanged -> APAGui.OnComboChanged

## XML runtime caveats

- Confidence: MEDIUM

- Description: Implementation-validated findings show that XML input, anchoring, and scroll layout behavior can depend on ancestor state, stable parent containers, and outer-window sizing even when child nodes appear correctly configured.

- Evidence:

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
  - Guidance: If alternating row tint stops short, increase the row template window and its child label widths instead of only changing the ListBox anchors.
  - `MEDIUM`: Default resizable WAR buttons render more cleanly at their native full height and need adequate footer/container height to prevent stacked controls from overlapping.
  - Guidance: Prefer keeping button height around 39px and increase the containing button-background height before squeezing stacked rows closer together.
  - `HIGH`: Adding draganddrop="true" to an XML Button element enables it to receive drag-and-drop events. The OnLButtonUp handler fires when the player releases a dragged item or ability over the button. Without this attribute the button does not receive drop events.
  - Guidance: Set draganddrop="true" on the Button element in XML. In the OnLButtonUp Lua handler, check Cursor.IconOnCursor() and guard on Cursor.Data.Source before reading the payload. Ensure handleinput="true" or handleinput is not explicitly disabled on the button and its parent chain.

## XML list binding pattern

- Confidence: MEDIUM

- Description: ListBox rows are commonly bound through ListData-backed Lua tables, with ListColumns supplying text fields and Lua population callbacks handling extra row setup such as icons or reordered display.

- Evidence:

- QuickTacticSwitch: `ListData table="QTS.listDisplayData" populationfunction="QTS.PopulateDisplay"` binds a ListBox to Lua-backed row data.
  - QuickTacticSwitch: `ListColumns` binds `Name` and `Enemy`, while `QTS.PopulateDisplay` uses `QuickTacticSwitchWindowList.PopulatorIndices` to populate row icons.
  - QuickTacticSwitch: `ListBoxSetDisplayOrder` and `ListBoxGetDataIndex` are used to manage visible ordering and row-to-data mapping.
  - AggroMeter: `ListData table="AggroMeter.Listdata" populationfunction=""` suggests column-only text binding works without a custom population callback.
  - `MEDIUM`: `ListData` appears to be the XML binding point that connects a `ListBox` row definition to a Lua backing table.
  - Use the `table` attribute as the backing collection path and treat `populationfunction` as optional custom row-population logic rather than the primary text-binding mechanism.
  - `MEDIUM`: `ListColumn` entries under `ListData` appear to map row-template child windows to fields on each Lua table entry.
  - For list rows that need extra runtime state such as images, use `ListColumns` for text fields and a Lua population function for the remaining row setup.
  - `MEDIUM`: list ordering and visible-row mapping are commonly managed from Lua with `ListBoxSetDisplayOrder`, `<ListBoxName>.PopulatorIndices`, and `ListBoxGetDataIndex`.

## Lua encoding caveats

- Confidence: HIGH

- Description: Implementation-validated encoding and tooling findings for WAR Lua addon files.

- Evidence:

- `HIGH`: The WAR engine Lua parser rejects .lua files that begin with a UTF-8 BOM (byte sequence 0xEF 0xBB 0xBF). The error message is: Error Reading Lua buffer <file>: [string ...]:1: '=' expected near '»'. The '»' character is the visual artefact of the BOM rendered in the error log.
  - Guidance: Always save Lua addon files as UTF-8 without BOM. In PowerShell 5.1, do NOT use Set-Content -Encoding UTF8 (it adds a BOM). Use [System.IO.File]::WriteAllText(path, content, [System.Text.UTF8Encoding]::new($false)) instead. To verify: read the first 3 bytes — they must NOT be 239 187 191.

## WAR API runtime facts

- Confidence: HIGH

- Description: Implementation-validated WAR engine API facts discovered through addon development and confirmed against the live engine.

- Evidence:

- `HIGH`: The standard pattern for a multi-line text-only tooltip is: Tooltips.CreateTextOnlyTooltip(windowName, nil) then Tooltips.SetTooltipText(row, col, wstring) for each line, Tooltips.SetTooltipColorDef(1, 1, Tooltips.COLOR_HEADING) for the heading row, then Tooltips.Finalize() and Tooltips.AnchorTooltip(Tooltips.ANCHOR_WINDOW_RIGHT). The MouseOverEnd handler should call Tooltips.ClearTooltip().
  - Guidance: Use row=1, col=1 for the heading line with COLOR_HEADING applied. Use row=2+ for body lines. The row index is 1-based. Do NOT call Tooltips.AddLineToTooltip — that function does not exist and causes a nil-call runtime error.
  - `HIGH`: Tooltips.AddLineToTooltip does not exist in the WAR engine. Calling it produces the runtime error: attempt to call field 'AddLineToTooltip' (a nil value).
  - Guidance: Never call Tooltips.AddLineToTooltip. To add a separator or body line, call Tooltips.SetTooltipText with the next row index instead.
  - `MEDIUM`: Tooltips.CreateAbilityTooltip(abilityData, windowName, anchor, hintWstring, nil) creates an ability tooltip anchored to a named window. The first argument is the table returned by GetAbilityData(id).
  - Guidance: Pass the table from GetAbilityData(id) as the first argument. Use Tooltips.ANCHOR_WINDOW_RIGHT as anchor. The fifth argument appears unused (nil). Corroborate in additional addons before treating as a hard contract.
  - `HIGH`: When an ability is dragged from the action bar: Cursor.IconOnCursor() returns true; Cursor.Data.Source equals Cursor.SOURCE_ACTION_LIST; Cursor.Data.ObjectId contains the ability ID as an integer. After handling the drop, Cursor.Clear() must be called to consume the cursor item.
  - Guidance: Always check Cursor.IconOnCursor() first. Guard with Cursor.Data.Source == Cursor.SOURCE_ACTION_LIST before reading Cursor.Data.ObjectId. Call Cursor.Clear() after every drop handling path (success or rejection) to avoid leaving a stale cursor icon.
