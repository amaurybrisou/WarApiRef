# DialogManager

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 80/100

## Confidence Assessment

- Level: HIGH

- Score: 80/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin |
| Namespaces detected | DialogManager |
| Source kinds | lua_calls |
| Example locations | CM_ClosetGoblin: OnClickDeleteSetButton, CM_ClosetGoblin: OnClickNewSetButton, CM_ClosetGoblin: OnRowMenuRenameSetClick, CM_ClosetGoblin: OnSubmitNewSetName, CM_ClosetGoblin: OnSubmitSetRename |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
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

## Description

Shared function table with 2 member functions; the primary API surface for 1 addons.

## Functions

- DialogManager.MakeTextEntryDialog
- DialogManager.MakeTwoButtonDialog

## Observed Members

- none

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: OnClickDeleteSetButton -> DialogManager.MakeTwoButtonDialog(cgL["confirm_delete_set"]:gsub(L "#1#",set.name), confirmYes, ClosetGoblinCharacterWindow.OnConfirmDeleteSet, confirmNo, nil)
- CM_ClosetGoblin: OnClickNewSetButton -> DialogManager.MakeTextEntryDialog(cgL["Set_Name"], cgL["Enter_set_name"], L "", ClosetGoblinCharacterWindow.OnSubmitNewSetName, nil)
- CM_ClosetGoblin: OnRowMenuRenameSetClick -> DialogManager.MakeTextEntryDialog(cgL["New_set_name"], cgL["Enter_set_name"], set.name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil)
- CM_ClosetGoblin: OnSubmitNewSetName -> DialogManager.MakeTextEntryDialog(cgL["Set name"], L "Enter set name :", name, ClosetGoblinCharacterWindow.OnSubmitNewSetName, nil)
- CM_ClosetGoblin: OnSubmitSetRename -> DialogManager.MakeTextEntryDialog(cgL["New_set_name"], cgL["Enter_set_name"], name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil)

## Notes

- none
