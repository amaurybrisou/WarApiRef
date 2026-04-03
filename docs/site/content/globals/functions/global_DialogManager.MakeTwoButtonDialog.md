# DialogManager.MakeTwoButtonDialog

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 16 addons

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
| Addons seen in | AdvancedRenownTrainer, BankWindowFix, CM_ClosetGoblin, Crusher, Enemy, Hopper, Lib RuString, MapMonster |
| Files seen in | AdvancedRenownTraining.lua, ClosetGoblinCharacterWindow.lua, Code/CombatLog/CombatLog.lua, Code/CombatLog/CombatLogStatsWindow.lua, Code/Core/Main.lua, Code/Marks/Marks.lua, Code/UnitFrames/EffectsIndicator.lua, Code/UnitFrames/UnitFrames.lua |
| Namespaces detected | DialogManager |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: Respecialize, BankWindowFix: BagEquipmentRButtonUp, CM_ClosetGoblin: OnClickDeleteSetButton, Crusher: InitializeWindow, Crusher: OnLButtonUp, Enemy: CombatLogInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 55 |
| Global usage count | 55 |
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

Observed as a global function across 16 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_CONFIRMATION,{MoneyFrame.FormatMoneyString(respecCost)}), GetStringFromTable("CustomizeUiStrings",StringTables.CustomizeUi.TEXT_UI_MOD_SETTINGS_CHANGED_DIALOG), L "Answer Rally Call or Open WAR Report ?" |
| arg2 | Observed as a text or wstring payload. | Observed values: GetString(StringTables.Default.LABEL_YES), L "Aye", L "Hell Yeah!" |
| arg3 | Observed as a function or method reference. | Observed values: ARC, AdvancedRenownTraining.RefundRenownPoints, ClosetGoblinCharacterWindow.OnConfirmDeleteSet |
| arg4 | Observed as a text or wstring payload. | Observed values: GetString(StringTables.Default.LABEL_NO), L "Nah", L "No" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedRenownTrainer
- BankWindowFix
- CM_ClosetGoblin
- Crusher
- Enemy
- Hopper
- Lib RuString
- MapMonster
- Minmap
- Motion
- Pure
- RoR_SoR
- TastyButtons
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- wbLeadHelper

## Examples

- AdvancedRenownTrainer: Respecialize -> DialogManager.MakeTwoButtonDialog(GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_CONFIRMATION,{MoneyFrame.FormatMoneyString(respecCost)}), GetString(StringTables.Default.LABEL_YES), AdvancedRenownTraining.RefundRenownPoints, GetString(StringTables.Default.LABEL_NO))
- BankWindowFix: BagEquipmentRButtonUp -> DialogManager.MakeTwoButtonDialog(text, GetString(StringTables.Default.LABEL_YES), function()DestroyItem(cursorType,slot)end, GetString(StringTables.Default.LABEL_NO), nil, nil, nil, nil, nil, nil, DIALOGID_DESTROY_ITEM)
- CM_ClosetGoblin: OnClickDeleteSetButton -> DialogManager.MakeTwoButtonDialog(cgL["confirm_delete_set"]:gsub(L "#1#",set.name), confirmYes, ClosetGoblinCharacterWindow.OnConfirmDeleteSet, confirmNo, nil)
- Crusher: InitializeWindow -> DialogManager.MakeTwoButtonDialog(T["Are you sure?"], GetString(StringTables.Default.LABEL_YES), CrusherConfig_Profiles_OnResetProfile, GetString(StringTables.Default.LABEL_NO), nil, nil, nil, nil, nil, DialogManager.TYPE_MODAL)
- Crusher: OnLButtonUp -> DialogManager.MakeTwoButtonDialog(T["Are you sure?"], GetString(StringTables.Default.LABEL_YES), CrusherConfig_Profiles_OnResetProfile, GetString(StringTables.Default.LABEL_NO), nil, nil, nil, nil, nil, DialogManager.TYPE_MODAL)
- Enemy: CombatLogInitialize -> DialogManager.MakeTwoButtonDialog(L "Enemy addon\n\nCause of considerable changes in combat log statistics code it's recommended to reset combat log settings to default values. Only combat log settings will be affected.\n\nDo you want to reset it now (you will have to wait for game interface to reload)?\n\nRecommended - yes", L "Yes", Enemy.CombatLogResetSettings, L "No")

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Used With

- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeOneButtonDialog](global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetCursorForBackpack](global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [EA_Window_InteractionRenownTraining.GetPointsAlreadyPurchased](global_EA_Window_InteractionRenownTraining.GetPointsAlreadyPurchased.md) (HIGH 98/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 90/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 80/100) - Global Function
- [GameData.Player.GetRenownRefundCost](global_GameData.Player.GetRenownRefundCost.md) (HIGH 80/100) - Global Function
- [BankWindow.IsShowing](global_BankWindow.IsShowing.md) (HIGH 71/100) - Global Function

## Affects

- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [GameData.Player.GetRenownRefundCost](global_GameData.Player.GetRenownRefundCost.md) (HIGH 80/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
