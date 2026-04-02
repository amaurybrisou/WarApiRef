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
