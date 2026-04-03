# DialogManager.MakeOneButtonDialog

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 12 addons

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
| Addons seen in | ActionFraction, AdvancedRenownTrainer, Aura, BankArkel, Crusher, Enemy, Miracle Grow Remix, Shinies |
| Files seen in | AdvancedRenownTraining.lua, BankArkel.lua, Code/Core/Groups/EnemyEffectFilter.lua, Code/UnitFrames/ClickCasting.lua, Code/UnitFrames/EffectsIndicator.lua, Code/UnitFrames/UnitFramePart.lua, MGRemix.lua, Modules/UI/Shinies-UI-Auctions.lua |
| Namespaces detected | DialogManager |
| Source kinds | lua_calls |
| Example locations | ActionFraction: SetLocationActionPointBar, AdvancedRenownTrainer: Respecialize, Aura: OnImportExportOkButton, BankArkel: ConvertDB, BankArkel: Init, Crusher: performCrush |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 25 |
| Global usage count | 25 |
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
DialogManager.MakeOneButtonDialog(arg1, arg2, arg3)
```

## Description

Observed as a global function across 12 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_NOT_ENOUGH_MONEY,{MoneyFrame.FormatMoneyString(respecCost)}), L "Did you know that you can use the mod 'Busted' to help submit more useful bug reports? You can find it on Curse.", L "In Squared 2.4+ you should no longer need to set rangefading-ability." |
| arg2 | Observed as a text or wstring payload. | Observed values: GetPregameString(StringTables.Pregame.LABEL_CONTINUE), GetString(StringTables.Default.LABEL_OKAY), L "OK" |
| arg3 | Observed as a runtime window or control identifier. | Observed values: BankArkel.ResetDB, nil |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- ActionFraction
- AdvancedRenownTrainer
- Aura
- BankArkel
- Crusher
- Enemy
- Miracle Grow Remix
- Shinies
- Squared
- TidyChat
- Vectors
- wbLeadHelper

## Examples

- ActionFraction: SetLocationActionPointBar -> DialogManager.MakeOneButtonDialog(LOC_TEXT.PLAYERWINDOW_UNAVAILABLE, GetPregameString(StringTables.Pregame.LABEL_CONTINUE))
- AdvancedRenownTrainer: Respecialize -> DialogManager.MakeOneButtonDialog(GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_NOT_ENOUGH_MONEY,{MoneyFrame.FormatMoneyString(respecCost)}), GetString(StringTables.Default.LABEL_OKAY), nil)
- Aura: OnImportExportOkButton -> DialogManager.MakeOneButtonDialog(T["An error occurred attempting to load the Aura.  Verify you have entered the correct data."], L "Ok", nil)
- BankArkel: ConvertDB -> DialogManager.MakeOneButtonDialog(StringToWString(convTxt.."v.1 >>> v.2"), StringToWString(errBtn1), nil)
- BankArkel: Init -> DialogManager.MakeOneButtonDialog(StringToWString(errTxt1), StringToWString(errBtn1), BankArkel.ResetDB)
- Crusher: performCrush -> DialogManager.MakeOneButtonDialog(T["Unable to salvage item.  You broke it, you bought it!"], T["Ok"], nil)

## Used With

- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeTwoButtonDialog](global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [EA_Window_InteractionRenownTraining.GetPointsAlreadyPurchased](global_EA_Window_InteractionRenownTraining.GetPointsAlreadyPurchased.md) (HIGH 98/100) - Global Function
- [GameData.Player.GetRenownRefundCost](global_GameData.Player.GetRenownRefundCost.md) (HIGH 80/100) - Global Function

## Affects

- [GameData.SalvagingTypes.MAGICAL](../../gamedata/fields/gamedata_GameData.SalvagingTypes.MAGICAL.md) (HIGH 100/100) - GameData Field
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field
- [GameData.Player.GetRenownRefundCost](global_GameData.Player.GetRenownRefundCost.md) (HIGH 80/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
