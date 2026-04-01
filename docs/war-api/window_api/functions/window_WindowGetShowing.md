# WindowGetShowing

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 45 addons

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
| Addons seen in | Ace, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, AnywhereTrainerAdditions, Aura, AutoMark, BankWindowFix |
| Files seen in | `/workspace/Ace/LibGUI.lua:84`, `/workspace/AggroMeter/AggroMeter.lua:374`, `/workspace/AnywhereTrainer/source/AnywhereTrainer.lua:227`, `/workspace/AnywhereTrainer/source/AnywhereTrainer.lua:235`, `/workspace/AnywhereTrainer/source/AnywhereTrainer.lua:243`, `/workspace/AnywhereTrainer/source/AnywhereTrainer.lua:261`, `/workspace/AnywhereTrainer/source/AnywhereTrainer.lua:271`, `/workspace/AnywhereTrainer/source/AnywhereTrainer.lua:280` |
| Namespaces detected | WindowGetShowing |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Showing, AdvancedRenownTrainer: AdvancedRenownTraining.TogglePresets, AggroMeter: AggroMeter.Close, AnywhereTrainer: AnywhereTrainer.OnLeftClickAuction, AnywhereTrainer: AnywhereTrainer.OnLeftClickBank, AnywhereTrainer: AnywhereTrainer.OnLeftClickCareer |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 305 |
| Global usage count | 305 |
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
WindowGetShowing(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "AbilitiesWindow", "AdvancedRenownTrainingWindow", "AggroMeterGrayWindow" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- AnywhereTrainerAdditions
- Aura
- AutoMark
- BankWindowFix
- BuffHead
- Busted
- CombatTextNames
- DAoCBuff
- DaemonAssist
- DeepSleep
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- GoldTracker
- GuardLine
- LibGroup
- LibWBToggler
- MapMonster
- MapPin
- Miracle Grow Remix
- PeaceOut
- Pocket Palette
- PotionBar
- Queue Queuer
- QuickTacticSwitch
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyQueue
- TidyRoll
- TurretRange
- WarBoard
- WarTriage
- WhoHealedMe
- WoH-Reticle
- nRarity
- wbLeadHelper

## Examples

- Ace: LIBGUI_ELEMENT:Showing -> WindowGetShowing(self.name)
- AdvancedRenownTrainer: AdvancedRenownTraining.TogglePresets -> WindowGetShowing(PresetWindowName)
- AggroMeter: AggroMeter.Close -> WindowGetShowing("AggroMeterGrayWindow")
- AnywhereTrainer: AnywhereTrainer.OnLeftClickAuction -> WindowGetShowing("ShiniesWindow")
- AnywhereTrainer: AnywhereTrainer.OnLeftClickAuction -> WindowGetShowing("AuctionWindow")
- AnywhereTrainer: AnywhereTrainer.OnLeftClickBank -> WindowGetShowing("BankWindow")

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetDisabledFlag](window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](../../globals/functions/global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](../../globals/functions/global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetCursorForBackpack](../../globals/functions/global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetItemsFromBackpack](../../globals/functions/global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsRefinable](../../globals/functions/global_EA_Window_Backpack.IsRefinable.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsSlotLocked](../../globals/functions/global_EA_Window_Backpack.IsSlotLocked.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining.Hide](../../globals/functions/global_EA_Window_InteractionRenownTraining.Hide.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining.Show](../../globals/functions/global_EA_Window_InteractionRenownTraining.Show.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionStore](../../globals/tables/table_EA_Window_InteractionStore.md) (HIGH 100/100) - Global Table
- [EA_Window_ScenarioLobby.OnJoinInstanceWait](../../globals/functions/global_EA_Window_ScenarioLobby.OnJoinInstanceWait.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap](../../globals/tables/table_EA_Window_WorldMap.md) (HIGH 100/100) - Global Table
- [EA_Window_WorldMap.ShowZone](../../globals/functions/global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [PartyUtils.GetWarbandLeader](../../globals/functions/global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetId](window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetLayer](window_WindowGetLayer.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [BankWindow.IsShowing](../../globals/functions/global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function
- [WindowStartScaleAnimation](window_WindowStartScaleAnimation.md) (HIGH 90/100) - Window Function
- [BankWindow.Hide](../../globals/functions/global_BankWindow.Hide.md) (HIGH 88/100) - Global Function
- [BankWindow.Show](../../globals/functions/global_BankWindow.Show.md) (HIGH 88/100) - Global Function
- [HelpTips.SetFocusOnWindow](../../globals/functions/global_HelpTips.SetFocusOnWindow.md) (HIGH 88/100) - Global Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [EA_Window_InteractionCoreTraining.Hide](../../globals/functions/global_EA_Window_InteractionCoreTraining.Hide.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionCoreTraining.Show](../../globals/functions/global_EA_Window_InteractionCoreTraining.Show.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem](../../globals/functions/global_EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenSellItem](../../globals/functions/global_EA_Window_InteractionLibrarianStore.ConfirmThenSellItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore](../../globals/functions/global_EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithRepairMan](../../globals/functions/global_EA_Window_InteractionLibrarianStore.InteractingWithRepairMan.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionSpecialtyTraining.Hide](../../globals/functions/global_EA_Window_InteractionSpecialtyTraining.Hide.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionSpecialtyTraining.Show](../../globals/functions/global_EA_Window_InteractionSpecialtyTraining.Show.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionTomeTraining.Hide](../../globals/functions/global_EA_Window_InteractionTomeTraining.Hide.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionTomeTraining.Show](../../globals/functions/global_EA_Window_InteractionTomeTraining.Show.md) (HIGH 80/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
