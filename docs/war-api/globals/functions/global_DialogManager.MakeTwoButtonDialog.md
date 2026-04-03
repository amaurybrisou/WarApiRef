# DialogManager.MakeTwoButtonDialog

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | AdvancedRenownTrainer, CM_ClosetGoblin, Enemy, RoR_SoR |
| Files seen in | `/workspace/data/raw/ClosetGoblin/ClosetGoblinCharacterWindow.lua:216`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLog.lua:63`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogStatsWindow.lua:623`, `/workspace/data/raw/Enemy/Code/Core/Main.lua:422`, `/workspace/data/raw/Enemy/Code/Core/Main.lua:432`, `/workspace/data/raw/Enemy/Code/Marks/Marks.lua:505`, `/workspace/data/raw/Enemy/Code/UnitFrames/EffectsIndicator.lua:918`, `/workspace/data/raw/Enemy/Code/UnitFrames/UnitFrames.lua:1468` |
| Namespaces detected | DialogManager |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize, CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnClickDeleteSetButton, Enemy: Enemy.CombatLogInitialize, Enemy: Enemy.CombatLogUI_StatsWindow_SessionDelete, Enemy: Enemy.MarksUI_EnemyMarkIcon_Delete, Enemy: Enemy.UI_ConfigDialog_Reset |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 21 |
| Global usage count | 21 |
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
DialogManager.MakeTwoButtonDialog(arg1, arg2, arg3, arg4)
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_CONFIRMATION,{MoneyFrame.FormatMoneyString(respecCost)}), L "Are you sure you want to reset SoR settings?", L "Claim Keep?" |
| arg2 | Observed as a text or wstring payload. | Observed values: GetString(StringTables.Default.LABEL_YES), L "Yes", confirmYes |
| arg3 | Observed as a function or method reference. | Observed values: AdvancedRenownTraining.RefundRenownPoints, ClosetGoblinCharacterWindow.OnConfirmDeleteSet, Enemy.CombatLogResetSettings |
| arg4 | Observed as a text or wstring payload. | Observed values: GetString(StringTables.Default.LABEL_NO), L "No", confirmNo |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AdvancedRenownTrainer
- CM_ClosetGoblin
- Enemy
- RoR_SoR

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize -> DialogManager.MakeTwoButtonDialog(GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_CONFIRMATION,{MoneyFrame.FormatMoneyString(respecCost)}), GetString(StringTables.Default.LABEL_YES), AdvancedRenownTraining.RefundRenownPoints, GetString(StringTables.Default.LABEL_NO))
- CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnClickDeleteSetButton -> DialogManager.MakeTwoButtonDialog(cgL["confirm_delete_set"]:gsub(L "#1#",set.name), confirmYes, ClosetGoblinCharacterWindow.OnConfirmDeleteSet, confirmNo, nil)
- Enemy: Enemy.CombatLogInitialize -> DialogManager.MakeTwoButtonDialog(L "Enemy addon\n\nCause of considerable changes in combat log statistics code it's recommended to reset combat log settings to default values. Only combat log settings will be affected.\n\nDo you want to reset it now (you will have to wait for game interface to reload)?\n\nRecommended - yes", L "Yes", Enemy.CombatLogResetSettings, L "No")
- Enemy: Enemy.CombatLogUI_StatsWindow_SessionDelete -> DialogManager.MakeTwoButtonDialog(L "Reset session '"..session.name..L "' ?", L "Yes", function()Enemy.CombatLog_ResetStatsCurrentSession()end, L "No")
- Enemy: Enemy.CombatLogUI_StatsWindow_SessionDelete -> DialogManager.MakeTwoButtonDialog(L "Delete session '"..session.name..L "' ?", L "Yes", function()Enemy.CombatLog_StatsSessionDelete(session.stats)end, L "No")
- Enemy: Enemy.MarksUI_EnemyMarkIcon_Delete -> DialogManager.MakeTwoButtonDialog(L "Delete mark template '"..g.templates[g.contextMenuTemplateIndex].name..L "' ?", L "Yes", Enemy.MarksUI_EnemyMarkIcon_DeleteConfirmed, L "No")

## Related APIs

- none

## Used With

- [DialogManager.MakeOneButtonDialog](global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining](../tables/table_EA_Window_InteractionRenownTraining.md) (HIGH 100/100) - Global Table
- [GameData.Player.GetRenownRefundCost](global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 80/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [EA_Window_InteractionRenownTraining](../tables/table_EA_Window_InteractionRenownTraining.md) (HIGH 100/100) - Global Table
- [GameData.Player.GetRenownRefundCost](global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
