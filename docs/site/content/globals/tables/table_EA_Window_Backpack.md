# EA_Window_Backpack

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 130

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BagOMatic, BankArkel, BankWindowFix, JunkDump, Killer, Miracle Grow Remix, NPC Item Sale Price, Shinies |
| Files seen in | `/workspace/BankArkel/BankArkel.lua:371`, `/workspace/BankWindowFix/Source/BankWindowFix.lua:108`, `/workspace/JunkDump/JunkDump.lua:618`, `/workspace/Killer/KillerRenown.lua:108`, `/workspace/Killer/KillerRenown.lua:127`, `/workspace/MGRemix/MGRemix.lua:996`, `/workspace/Shinies/Modules/UI/Shinies-UI-Browse.lua:298`, `/workspace/nisp/Source/Nisp.lua:174` |
| Namespaces detected | EA_Window_Backpack |
| Source kinds | globals, lua_calls |
| Example locations | BankArkel: BankArkel.UpdateData, BankWindowFix: BankWindowFix.BagEquipmentRButtonUp, JunkDump: JunkDump.SetTestModeTint, Killer: Killer.CaptureInitialWarCrests, Killer: Killer.OnCurrencyUpdated, Miracle Grow Remix: MiracleGrow2.GetCursorForBackpack |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 9 |
| Local definition count | 8 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed shared global table or namespace surfaced in 9 addons.

## Functions

- EA_Window_Backpack.AutoAddCraftingItemIfPossible
- EA_Window_Backpack.ConfirmThenRefine
- EA_Window_Backpack.GetCursorForBackpack
- EA_Window_Backpack.GetItemsFromBackpack
- EA_Window_Backpack.GetPocketName
- EA_Window_Backpack.GetPocketNumberForSlot
- EA_Window_Backpack.GetSlotFromActionButtonGroup
- EA_Window_Backpack.IsRefinable
- EA_Window_Backpack.IsSlotLocked

## Observed Members

- none

## Seen In

- BagOMatic
- BankArkel
- BankWindowFix
- JunkDump
- Killer
- Miracle Grow Remix
- NPC Item Sale Price
- Shinies
- nRarity

## Examples

- BankArkel: BankArkel.UpdateData -> EA_Window_Backpack.GetItemsFromBackpack(Read[i])
- BankWindowFix: BankWindowFix.BagEquipmentRButtonUp -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.currentMode)
- BankWindowFix: BankWindowFix.BagEquipmentRButtonUp -> EA_Window_Backpack.GetCursorForBackpack(EA_Window_Backpack.currentMode)
- BankWindowFix: BankWindowFix.BagEquipmentRButtonUp -> EA_Window_Backpack.IsSlotLocked(slot, EA_Window_Backpack.currentMode)
- BankWindowFix: BankWindowFix.BagEquipmentRButtonUp -> EA_Window_Backpack.AutoAddCraftingItemIfPossible(slot)
- BankWindowFix: BankWindowFix.BagEquipmentRButtonUp -> EA_Window_Backpack.IsRefinable(itemData)

## Related APIs

- [DialogManager.MakeOneButtonDialog](../functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](../functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [LibSlash.IsSlashCmdRegistered](../functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](../functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [DestroyWindow](../functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Used With

- [BankWindow](table_BankWindow.md) (HIGH 100/100) - Global Table
- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeOneButtonDialog](../functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.OnKeyEnter](../functions/global_EA_ChatWindow.OnKeyEnter.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](../functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [LibSlash.IsSlashCmdRegistered](../functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](../functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](../functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Settings.Language.active](../../systemdata/fields/systemdata_SystemData.Settings.Language.active.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [BankWindow.Hide](../functions/global_BankWindow.Hide.md) (HIGH 88/100) - Global Function
- [DoesWindowExist](../functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none
