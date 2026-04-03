# XML list binding pattern

- Category: conventions
- Confidence: MEDIUM

## What is this pattern

The XML list binding pattern describes how ListBox elements are connected to Lua-backed data. A [ListData](../xml/element_types/element_ListData.md) element specifies the data source table and population callback, while [ListColumns](../xml/element_types/element_ListColumns.md) declare which fields display in each row.

## Why it exists

ListBox binding allows addons to display dynamic, sortable, filterable lists without manual row management:
- **Data-driven**: Changes to the data table automatically reflect in the list
- **Custom layout**: Population callbacks add computed fields, icons, colors
- **Efficient**: Only visible rows are rendered

## When it appears

- **List UI components**: Option panels, inventory displays, player lists
- **Data display**: Whenever tabular data must be shown
- **Dynamic updates**: When list contents change at runtime

## Minimal example

```xml
<ListBox name="MyList">
  <ListData table="MyAddon.listData" populationfunction="MyAddon.PopulateListRow" />
  <ListColumns>
    <ListColumn field="name" width="200" />
    <ListColumn field="count" width="50" />
  </ListColumns>
</ListBox>
```

```lua
MyAddon.listData = {}

function MyAddon.PopulateListRow(window, rowDataIndex)
  local rowData = MyAddon.listData[rowDataIndex]
  local row = ListBoxGetDisplayElement(window, rowDataIndex)
  
  -- Extra setup beyond declared columns
  row.iconImage:SetImage(GetItemIcon(rowData.itemID))
end
```

## Annotated real example

From QuickTacticSwitch (observed):

```xml
<!-- Declare list with data source and population function -->
<ListBox>
  <ListData table="QTS.listDisplayData" populationfunction="QTS.PopulateDisplay" />
  <ListColumns>
    <ListColumn field="Name" width="150" />
    <ListColumn field="Enemy" width="100" />
  </ListColumns>
</ListBox>
```

```lua
-- Lua backing: data table and population callback
QTS.listDisplayData = {
  { Name = "Firewall", Enemy = "Chaos" },
  { Name = "Shield", Enemy = "Order" }
}

function QTS.PopulateDisplay(window, rowIndex)
  -- Called for each visible row
  local row = ListBoxGetDisplayElement(window, rowIndex)
  local data = QTS.listDisplayData[rowIndex]
  
  -- Textual fields auto-populated by ListColumns
  -- Additional customization
  if data.isFavorite then
    row.starIcon:SetImage("icon_star_filled")
  end
end
```

**Key observations**:
- `ListData` table can contain any Lua table with indexed rows
- `PopulateDisplay` callback runs for each visible row
- `ListColumns` only declare simple text fields; population callback adds complexity
- `ListBoxGetDataIndex` maps display row to data table index

## Detection signals / evidence

**Observe**:
- XML `<ListBox>` with `<ListData table=` and `populationfunction=`
- Lua table with matching name contains array of row data
- Population callback function referenced matches function definition
- `ListBoxGetDisplayElement`, `ListBoxGetDataIndex`, `ListBoxSetDisplayOrder` calls in population callback

**Confidence indicators**:
- Table name exactly matches `ListData table=` attribute
- Population function signature matches callback invocation
- Row data is indexed (0-based or 1-based, consistently)

## Related patterns

- [XML to Lua binding](./conventions_XML_to_Lua_binding_pattern.md) — general handler binding
- [UI creation pattern](./conventions_UI_creation_pattern.md) — where ListBox window is created

## Common pitfalls

1. **Data table indexing mismatch**: Lua tables 1-indexed, display expects 0-indexed or vice versa.
   ```lua
   -- ❌ Wrong: mismatch causes off-by-one in display
   MyList[0] = {name = "First"}
   
   -- ✓ Correct: consistent 1-indexing
   MyList[1] = {name = "First"}
   ```

2. **Column field missing**: Declaring `<ListColumn field=` for field that doesn't exist in data rows.

3. **Slow population callback**: Population callback called for every visible row; expensive operations cause freezes.

4. **Not updating list after data change**: Modifying `listData` table directly without calling `ListBoxUpdateDisplay` or `ListBoxUpdateRow`.

## Involved APIs

- [ListBox](../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [ListData](../xml/element_types/element_ListData.md) (HIGH 100/100) - XML Element Type
- [ListColumns](../xml/element_types/element_ListColumns.md) (MEDIUM 55/100) - XML Element Type
- [ListColumn](../xml/element_types/element_ListColumn.md) (MEDIUM 45/100) - XML Element Type
