# XML list binding pattern

- Category: conventions
- Confidence: MEDIUM

## Description

ListBox rows are commonly bound through ListData-backed Lua tables, with ListColumns supplying text fields and Lua population callbacks handling extra row setup such as icons or reordered display.

## Involved APIs

- [ListBox](../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [Window](../globals/tables/table_Window.md) (HIGH 100/100) - Global Table

## Flow Diagram

```text
OnLButtonUp
  -> handlers: DebugWindow.ClearTextLog, InterfaceCore.ReloadUI
  -> ui: Button, ColorPicker, DynamicImage, FullResizeImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
QuickTacticSwitch: `ListData table="QTS.listDisplayData" populationfunction="QTS.PopulateDisplay"` binds a ListBox to Lua-backed row data.
```

## Evidence

- QuickTacticSwitch: `ListData table="QTS.listDisplayData" populationfunction="QTS.PopulateDisplay"` binds a ListBox to Lua-backed row data.
- QuickTacticSwitch: `ListColumns` binds `Name` and `Enemy`, while `QTS.PopulateDisplay` uses `QuickTacticSwitchWindowList.PopulatorIndices` to populate row icons.
- QuickTacticSwitch: `ListBoxSetDisplayOrder` and `ListBoxGetDataIndex` are used to manage visible ordering and row-to-data mapping.
- AggroMeter: `ListData table="AggroMeter.Listdata" populationfunction=""` suggests column-only text binding works without a custom population callback.
