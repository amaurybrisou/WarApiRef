# EA_ChatWindow.ShowTabCycleButtons

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 105

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat |
| Files seen in | `/workspace/TidyChat/TidyChat.lua:305` |
| Namespaces detected | EA_ChatWindow |
| Source kinds | globals, lua_calls |
| Example locations | TidyChat: TidyChatCore.SetWindowTabs |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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

## Signature (inferred)

```lua
EA_ChatWindow.ShowTabCycleButtons(arg1, arg2)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: wndGroupName |
| arg2 | Observed as a boolean toggle. | Observed values: false |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- TidyChat

## Examples

- TidyChat: TidyChatCore.SetWindowTabs -> EA_ChatWindow.ShowTabCycleButtons(wndGroupName, false)

## Related APIs

- none

## Used With

- [EA_ChatWindow.UpdateTabScrollWindow](global_EA_ChatWindow.UpdateTabScrollWindow.md) (HIGH 100/100) - Global Function
- [EA_ChatWindowGroups](../tables/table_EA_ChatWindowGroups.md) (HIGH 100/100) - Global Table
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [EA_ChatWindowGroups](../tables/table_EA_ChatWindowGroups.md) (HIGH 100/100) - Global Table

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
