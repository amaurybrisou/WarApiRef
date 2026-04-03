# DialogManager.MakeTextEntryDialog

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 90/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin |
| Files seen in | ClosetGoblinCharacterWindow.lua |
| Namespaces detected | DialogManager |
| Source kinds | lua_calls |
| Example locations | CM_ClosetGoblin: OnClickNewSetButton, CM_ClosetGoblin: OnRowMenuRenameSetClick, CM_ClosetGoblin: OnSubmitNewSetName, CM_ClosetGoblin: OnSubmitSetRename |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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
DialogManager.MakeTextEntryDialog(arg1, arg2, arg3, arg4, arg5)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: cgL["New_set_name"], cgL["Set name"], cgL["Set_Name"] |
| arg2 | Observed as a runtime window or control identifier. | Observed values: L "Enter set name :", cgL["Enter_set_name"] |
| arg3 | Observed as a runtime window or control identifier. | Observed values: L "", name, set.name |
| arg4 | Observed as a function or method reference. | Observed values: ClosetGoblinCharacterWindow.OnSubmitNewSetName, ClosetGoblinCharacterWindow.OnSubmitSetRename |
| arg5 | Observed as a runtime window or control identifier. | Observed values: nil |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: OnClickNewSetButton -> DialogManager.MakeTextEntryDialog(cgL["Set_Name"], cgL["Enter_set_name"], L "", ClosetGoblinCharacterWindow.OnSubmitNewSetName, nil)
- CM_ClosetGoblin: OnRowMenuRenameSetClick -> DialogManager.MakeTextEntryDialog(cgL["New_set_name"], cgL["Enter_set_name"], set.name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil)
- CM_ClosetGoblin: OnSubmitNewSetName -> DialogManager.MakeTextEntryDialog(cgL["Set name"], L "Enter set name :", name, ClosetGoblinCharacterWindow.OnSubmitNewSetName, nil)
- CM_ClosetGoblin: OnSubmitSetRename -> DialogManager.MakeTextEntryDialog(cgL["New_set_name"], cgL["Enter_set_name"], name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil)

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
