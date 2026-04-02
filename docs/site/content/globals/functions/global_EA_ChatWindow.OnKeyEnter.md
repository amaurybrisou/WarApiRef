# EA_ChatWindow.OnKeyEnter

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 148

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy, Swift Assist, WSCT |
| Files seen in | `/workspace_addons/Effigy/EffigySlashCommands.lua:32`, `/workspace_addons/swift-assist/SwiftAssist.lua:435`, `/workspace_addons/wsct/wsct.lua:196` |
| Namespaces detected | EA_ChatWindow |
| Source kinds | globals, lua_calls |
| Example locations | Effigy: Effigy.RegisterManually, Swift Assist: SwiftAssist.Initialize, WSCT: WSCT:HookChatWindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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
EA_ChatWindow.OnKeyEnter(arg1)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: ... |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Effigy
- Swift Assist
- WSCT

## Examples

- Effigy: Effigy.RegisterManually -> EA_ChatWindow.OnKeyEnter(...)
- Swift Assist: SwiftAssist.Initialize -> EA_ChatWindow.OnKeyEnter(...)
- WSCT: WSCT:HookChatWindow -> EA_ChatWindow.OnKeyEnter(...)

## Related APIs

- [DialogManager.MakeOneButtonDialog](global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function

## Used With

- [BankWindow](../tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [DialogManager.MakeOneButtonDialog](global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [EA_TextEntryGroupEntryBoxTextInput](../tables/table_EA_TextEntryGroupEntryBoxTextInput.md) (HIGH 100/100) - Global Table
- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [BankWindow.Hide](global_BankWindow.Hide.md) (HIGH 88/100) - Global Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- [EA_TextEntryGroupEntryBoxTextInput](../tables/table_EA_TextEntryGroupEntryBoxTextInput.md) (HIGH 100/100) - Global Table
- [SystemData.Events.MACROS_LOADED](../../systemdata/fields/systemdata_SystemData.Events.MACROS_LOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.MACRO_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.MACRO_UPDATED.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
