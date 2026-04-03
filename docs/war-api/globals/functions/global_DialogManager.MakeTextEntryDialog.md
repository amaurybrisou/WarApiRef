# DialogManager.MakeTextEntryDialog

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin, Enemy, MapPin, Motion, Shinies, Vectors |
| Files seen in | ClosetGoblinCharacterWindow.lua, Code/CombatLog/CombatLogStatsWindow.lua, Modules/UI/Shinies-UI-Browse.lua, Settings.lua, Source/Motion.lua, source/MapPin.lua |
| Namespaces detected | DialogManager |
| Source kinds | lua_calls |
| Example locations | CM_ClosetGoblin: OnClickNewSetButton, CM_ClosetGoblin: OnRowMenuRenameSetClick, CM_ClosetGoblin: OnSubmitNewSetName, CM_ClosetGoblin: OnSubmitSetRename, Enemy: CombatLogUI_StatsWindow_SessionAdd, Enemy: CombatLogUI_StatsWindow_SessionRename |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 20 |
| Global usage count | 20 |
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
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
DialogManager.MakeTextEntryDialog(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
```

## Description

Observed as a global function across 6 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: L "MapPin Export", L "MapPin Filter", L "MapPin" |
| arg2 | Observed as a text or wstring payload. | Observed values: L "Current MapPin Data: ", L "Enter new Profile Name", L "Enter new session name" |
| arg3 | Observed as a runtime window or control identifier. | Observed values: AddLink, DataText, L "" |
| arg4 | Observed as a function or method reference. | Observed values: ClosetGoblinCharacterWindow.OnSubmitNewSetName, ClosetGoblinCharacterWindow.OnSubmitSetRename, InputSubmit |
| arg5 | Observed as a runtime window or control identifier. | Observed values: Motion.OnRecipeAddCancel, ShiniesBrowseUI.Searches_OnAddCancel, nil |
| arg6 | Observed as a numeric value. | Observed values: 1000, 20, 250 |
| arg7 | Observed as a boolean toggle. | Observed values: false, true |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- CM_ClosetGoblin
- Enemy
- MapPin
- Motion
- Shinies
- Vectors

## Examples

- CM_ClosetGoblin: OnClickNewSetButton -> DialogManager.MakeTextEntryDialog(cgL["Set_Name"], cgL["Enter_set_name"], L "", ClosetGoblinCharacterWindow.OnSubmitNewSetName, nil)
- CM_ClosetGoblin: OnRowMenuRenameSetClick -> DialogManager.MakeTextEntryDialog(cgL["New_set_name"], cgL["Enter_set_name"], set.name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil)
- CM_ClosetGoblin: OnSubmitNewSetName -> DialogManager.MakeTextEntryDialog(cgL["Set name"], L "Enter set name :", name, ClosetGoblinCharacterWindow.OnSubmitNewSetName, nil)
- CM_ClosetGoblin: OnSubmitSetRename -> DialogManager.MakeTextEntryDialog(cgL["New_set_name"], cgL["Enter_set_name"], name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil)
- Enemy: CombatLogUI_StatsWindow_SessionAdd -> DialogManager.MakeTextEntryDialog(L "New combat log session", L "Enter session name", L "Default", function(text)Enemy.CombatLog_StatsSessionAdd(Enemy.trim(text))end, nil, 250, false)
- Enemy: CombatLogUI_StatsWindow_SessionRename -> DialogManager.MakeTextEntryDialog(L "Rename combat log session", L "Enter new session name", session.stats.name, function(text)local index=ComboBoxGetSelectedMenuItem("EnemyCombatLogStatsWindowSession")session.stats.name=Enemy.trim(text)Enemy.CombatLogUI_StatsWindow_UpdateSessionsList()ComboBoxSetSelectedMenuItem("EnemyCombatLogStatsWindowSession",index)Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged()end, nil, 250, false)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Used With

- [EA_ChatWindow.InsertText](global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [EA_Window_WorldMap.ShowZone](global_EA_Window_WorldMap.ShowZone.md) (HIGH 90/100) - Global Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
