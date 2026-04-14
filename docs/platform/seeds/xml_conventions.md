# XML Conventions Seed

## Layout Discipline

- Use consistent anchor strategy and naming.
- Keep template usage explicit for repeated structures.

## Input/Event Discipline

- Treat input flags and handler bindings as contract-sensitive.
- Prefer minimal, intentional event wiring.
- `MEDIUM`: if a child click target looks correctly wired but remains inert, validate `handleinput` across the parent or template chain as well as on the child node.

## Observed Layout Caveats

- `MEDIUM`: early layout reads on nested scroll content can under-report the eventual usable region.
- When first-render list height looks collapsed, compute dimensions from the outer parent first, then resize child content and refresh the scroll rect.

## Observed List Binding Patterns

- `MEDIUM`: `ListData` appears to be the XML binding point that connects a `ListBox` row definition to a Lua backing table.
- Use the `table` attribute as the backing collection path and treat `populationfunction` as optional custom row-population logic rather than the primary text-binding mechanism.
- `MEDIUM`: `ListColumn` entries under `ListData` appear to map row-template child windows to fields on each Lua table entry.
- For list rows that need extra runtime state such as images, use `ListColumns` for text fields and a Lua population function for the remaining row setup.
- `MEDIUM`: list ordering and visible-row mapping are commonly managed from Lua with `ListBoxSetDisplayOrder`, `<ListBoxName>.PopulatorIndices`, and `ListBoxGetDataIndex`.

## Provenance Rule

- XML behavior claims must be backed by canonical docs or MCP evidence.
- Mark inferred behavior with confidence notes.

<!-- OBSERVATION:aura_xml_layout_anchor_and_spacing_patterns (promoted:2026-04-05T00:20:00Z) -->
> Source: Aura | Confidence: MEDIUM | Promoted: 2026-04-05

- `MEDIUM`: When footer controls or edit boxes drift under sibling-based anchoring, anchoring them directly to a stable parent container is more reliable than chaining them through neighboring controls.
  - Guidance: If a control renders in the wrong column or outside the window, clear its anchors and re-anchor it from `$parent` or a stable container such as `$parentButtonBackground` with explicit offsets.
- `MEDIUM`: Text-driven header or form rows are more stable when label dimensions are measured after LabelSetText and adjacent controls are positioned from the measured width.
  - Guidance: Use `LabelGetTextDimensions()` and then resize or re-anchor neighboring labels and edit boxes to avoid truncation, overlap, or misplaced values.
- `MEDIUM`: A ListBox row highlight or tint only spans the width of the row template; widening the list box alone does not widen the tinted row background.
  - Guidance: If alternating row tint stops short, increase the row template window and its child label widths instead of only changing the ListBox anchors.
- `MEDIUM`: Default resizable WAR buttons render more cleanly at their native full height and need adequate footer/container height to prevent stacked controls from overlapping.
  - Guidance: Prefer keeping button height around 39px and increase the containing button-background height before squeezing stacked rows closer together.

<!-- OBSERVATION:apa_xml_button_draganddrop_attribute (promoted:2026-04-14T05:15:34Z) -->
> Source: AdvancedPetAssist | Confidence: HIGH | Promoted: 2026-04-14

- `HIGH`: Adding draganddrop="true" to an XML Button element enables it to receive drag-and-drop events. The OnLButtonUp handler fires when the player releases a dragged item or ability over the button. Without this attribute the button does not receive drop events.
  - Guidance: Set draganddrop="true" on the Button element in XML. In the OnLButtonUp Lua handler, check Cursor.IconOnCursor() and guard on Cursor.Data.Source before reading the payload. Ensure handleinput="true" or handleinput is not explicitly disabled on the button and its parent chain.
