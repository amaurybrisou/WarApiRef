# DialogManager.MakeOneButtonDialog

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 8 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, Aura, BankArkel, Enemy, Miracle Grow Remix, Shinies, TidyChat, wbLeadHelper |
| Files seen in | `/workspace/Aura/Source/AuraShares.lua:419`, `/workspace/BankArkel/BankArkel.lua:85`, `/workspace/BankArkel/BankArkel.lua:95`, `/workspace/Enemy/Code/Core/Groups/EnemyEffectFilter.lua:673`, `/workspace/Enemy/Code/UnitFrames/ClickCasting.lua:441`, `/workspace/Enemy/Code/UnitFrames/EffectsIndicator.lua:1111`, `/workspace/Enemy/Code/UnitFrames/UnitFramePart.lua:1006`, `/workspace/MGRemix/MGRemix.lua:492` |
| Namespaces detected | DialogManager |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize, Aura: AuraShares.OnImportExportOkButton, BankArkel: BankArkel.ConvertDB, BankArkel: BankArkel.Init, Enemy: Enemy.GroupsUI_EffectFilterDialog_Ok, Enemy: Enemy.UnitFramesUI_EffectsIndicatorDialog_Ok |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 15 |
| Global usage count | 15 |
| Local definition count | 0 |
| Documentation references | 1 |
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
DialogManager.MakeOneButtonDialog(arg1, arg2)
```

## Description

Observed as a global function across 8 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_NOT_ENOUGH_MONEY,{MoneyFrame.FormatMoneyString(respecCost)}), L "Invalid name specified. Name should not contain any spaces and can not be 'and', 'or', 'not' keyword.", L "Log is empty\n" |
| arg2 | Observed as a text or wstring payload. | Observed values: GetString(StringTables.Default.LABEL_OKAY), L "Ok", StringToWString(errBtn1) |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- AdvancedRenownTrainer
- Aura
- BankArkel
- Enemy
- Miracle Grow Remix
- Shinies
- TidyChat
- wbLeadHelper

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize -> DialogManager.MakeOneButtonDialog(GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_NOT_ENOUGH_MONEY,{MoneyFrame.FormatMoneyString(respecCost)}), GetString(StringTables.Default.LABEL_OKAY), nil)
- Aura: AuraShares.OnImportExportOkButton -> DialogManager.MakeOneButtonDialog(T["An error occurred attempting to load the Aura.  Verify you have entered the correct data."], L "Ok", nil)
- BankArkel: BankArkel.ConvertDB -> DialogManager.MakeOneButtonDialog(StringToWString(convTxt.."v.1 >>> v.2"), StringToWString(errBtn1), nil)
- BankArkel: BankArkel.Init -> DialogManager.MakeOneButtonDialog(StringToWString(errTxt1), StringToWString(errBtn1), BankArkel.ResetDB)
- Enemy: Enemy.GroupsUI_EffectFilterDialog_Ok -> DialogManager.MakeOneButtonDialog(L "Invalid name specified. Name should not contain any spaces and can not be 'and', 'or', 'not' keyword.", L "Ok")
- Enemy: Enemy.UnitFramesUI_EffectsIndicatorDialog_Ok -> DialogManager.MakeOneButtonDialog(L "You should add at least one effect filter", L "Ok")

## Related APIs

- none

## Used With

- [BankWindow](../tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeTwoButtonDialog](global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.OnKeyEnter](global_EA_ChatWindow.OnKeyEnter.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [EA_Window_InteractionRenownTraining](../tables/table_EA_Window_InteractionRenownTraining.md) (HIGH 100/100) - Global Table
- [GameData.Player.GetRenownRefundCost](global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [BankWindow.Hide](global_BankWindow.Hide.md) (HIGH 88/100) - Global Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [BankWindow](../tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [EA_ChatWindow.OnKeyEnter](global_EA_ChatWindow.OnKeyEnter.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [EA_Window_InteractionRenownTraining](../tables/table_EA_Window_InteractionRenownTraining.md) (HIGH 100/100) - Global Table
- [GameData.Player.GetRenownRefundCost](global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field
- [BankWindow.Hide](global_BankWindow.Hide.md) (HIGH 88/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
