# DialogManager.MakeTextEntryDialog

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 98/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin, Enemy, Shinies |
| Files seen in | `/workspace/data/raw/ClosetGoblin/ClosetGoblinCharacterWindow.lua:147`, `/workspace/data/raw/ClosetGoblin/ClosetGoblinCharacterWindow.lua:154`, `/workspace/data/raw/ClosetGoblin/ClosetGoblinCharacterWindow.lua:212`, `/workspace/data/raw/ClosetGoblin/ClosetGoblinCharacterWindow.lua:234`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogStatsWindow.lua:612`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogStatsWindow.lua:645`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Browse.lua:1105` |
| Namespaces detected | DialogManager |
| Source kinds | lua_calls |
| Example locations | CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnClickNewSetButton, CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnRowMenuRenameSetClick, CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnSubmitNewSetName, CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnSubmitSetRename, Enemy: Enemy.CombatLogUI_StatsWindow_SessionAdd, Enemy: Enemy.CombatLogUI_StatsWindow_SessionRename |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
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
DialogManager.MakeTextEntryDialog(arg1, arg2, arg3, arg4, arg5)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: L "New combat log session", L "Rename combat log session", T["Shinies - Save Search"] |
| arg2 | Observed as a text or wstring payload. | Observed values: L "Enter new session name", L "Enter session name", L "Enter set name :" |
| arg3 | Observed as a runtime window or control identifier. | Observed values: L "", L "Default", T["New Search"] |
| arg4 | Observed as a function or method reference. | Observed values: ClosetGoblinCharacterWindow.OnSubmitNewSetName, ClosetGoblinCharacterWindow.OnSubmitSetRename, ShiniesBrowseUI.Searches_OnAddSuccess |
| arg5 | Observed as a runtime window or control identifier. | Observed values: ShiniesBrowseUI.Searches_OnAddCancel, nil |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- CM_ClosetGoblin
- Enemy
- Shinies

## Examples

- CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnClickNewSetButton -> DialogManager.MakeTextEntryDialog(cgL["Set_Name"], cgL["Enter_set_name"], L "", ClosetGoblinCharacterWindow.OnSubmitNewSetName, nil)
- CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnRowMenuRenameSetClick -> DialogManager.MakeTextEntryDialog(cgL["New_set_name"], cgL["Enter_set_name"], set.name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil)
- CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnSubmitNewSetName -> DialogManager.MakeTextEntryDialog(cgL["Set name"], L "Enter set name :", name, ClosetGoblinCharacterWindow.OnSubmitNewSetName, nil)
- CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnSubmitSetRename -> DialogManager.MakeTextEntryDialog(cgL["New_set_name"], cgL["Enter_set_name"], name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil)
- Enemy: Enemy.CombatLogUI_StatsWindow_SessionAdd -> DialogManager.MakeTextEntryDialog(L "New combat log session", L "Enter session name", L "Default", function(text)Enemy.CombatLog_StatsSessionAdd(Enemy.trim(text))end, nil, 250, false)
- Enemy: Enemy.CombatLogUI_StatsWindow_SessionRename -> DialogManager.MakeTextEntryDialog(L "Rename combat log session", L "Enter new session name", session.stats.name, function(text)local index=ComboBoxGetSelectedMenuItem("EnemyCombatLogStatsWindowSession")session.stats.name=Enemy.trim(text)Enemy.CombatLogUI_StatsWindow_UpdateSessionsList()ComboBoxSetSelectedMenuItem("EnemyCombatLogStatsWindowSession",index)Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged()end, nil, 250, false)

## Related APIs

- none

## Used With

- none

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
