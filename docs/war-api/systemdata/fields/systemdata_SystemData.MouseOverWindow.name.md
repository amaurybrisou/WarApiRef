# SystemData.MouseOverWindow.name

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 63 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, AggroMeter, Assist, Atlas, AuctionStats, BankArkel, BuffHead, CM_ClosetGoblin |
| Files seen in | AdvancedRenownTraining.lua, AggroMeter.lua, AuctionAssist.lua, BankArkel.lua, Cheeseboard.lua, Classes/Display.lua, ClosetGoblinCharacterWindow.lua, ClosetGoblinOptionWindow.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | AbilityClear, AbilityCursorSwap, AbilityTooltip, AuctionHouseWindowCreateSearchItemName_LButtonDown, Autohide, BankWindow_EquipmentMouseOver |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 157 |
| Global usage count | 157 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

SystemData.SystemData.MouseOverWindow.name field accessed by 63 addons; commonly found in AbilityClear and AbilityCursorSwap, AbilityTooltip, AuctionHouseWindowCreateSearchItemName_LButtonDown, Autohide, BankWindow_EquipmentMouseOver, Clicky, CombatLogUI_EpsWindow_OnMouseOver, CombatLogUI_SnapshotWindow_ListRowMouseOver, CombatLogUI_StatsWindow_ListRowMouseOver, CombatLogUI_StatsWindow_OnEpsMouseOver, CombatLogUI_TargetDefenseTotalWindow_OnMouseOver, CombatLogUI_TargetDefenseWindow_OnMouseOver, ConfigurationWindow_ShowTooltip, ContextFillAll, ContextHover, ContextItem, ContextSubMenu, CreateNewTokenView, DestroyWindow, GroupsUI_EffectFilterDialog_OnAbilityIdsMouseOver, Guard_GuardIndicator_OnMouseOver, HotbarPageDownProxy, HotbarPageUpProxy, InitializeWindow, LButtonUp, L_BUTTON_UP_PROCESSED, LabelTint, LabelTintRemove, MainTooltip, MarkUI_EnemyMark_OnLButtonDown, MarkUI_EnemyMark_OnRButtonUp, MarksUI_EnemyMarkIcon_OnLButtonUp, MarksUI_EnemyMarkIcon_OnMouseOver, MarksUI_EnemyMarkIcon_OnRButtonUp, Mouse, MouseOver, MouseOverCareerResourceWindow, MouseOverCraftingSkill, MouseOverEndCraftingSkill, MouseOverEndListLabel, MouseOverEndRarity, MouseOverListLabel, MouseOverRarity, MouseOverSpecialButtons, MoveDown, OnButtonUp, OnCheckboxLBU, OnCheckboxRBU, OnContainerRClick, OnDontSplitNameCheck, OnEditLButtonUp, OnEnableFilter, OnHover, OnHoverLoss, OnHoverWin, OnItemMouseOver, OnLButtonDownPlayerRow, OnLButtonUp, OnLButtonUpProcessed, OnLButtonUp_ContextMenu_Item, OnLButtonUp_ContextMenu_Recipe, OnLButtonUp_ContextMenu_Recipe_Delete, OnLButtonUp_Sigil, OnListLButtonUp, OnMonitorClicked, OnMouseOver, OnMouseOverListMember, OnMouseOverPlayerRow, OnMouseOverResultsIcon, OnMouseOverSearchButton, OnMouseOverStart, OnMouseOverStart2, OnMouseOverWarbandListBox, OnMouseOver_ContextMenu_Item, OnMouseOver_ContextMenu_Recipe, OnMouseOver_Icon, OnMouseOver_Results_ListItem, OnMouseOver_Sigil, OnMouseOver_Window_Hint, OnMouseoverStatHeaderLabelIcon, OnMouseoverStatHeaderLabelRecent, OnMouseoverStatHeaderLabelReference, OnMouseoverStatHeaderLabelStat, OnPinMouseOver, OnPinMouseOverEnd, OnRButtonUp, OnRButtonUpPlayerRow, OnRowLDown, OnRowLUp, OnRowMenuActivateSet, OnRowMenuAssoiacteClick, OnRowMenuLinkTactics, OnRowRDown, OnRowRUp, OnSlotMouseOver, OnTabRBU, OnTargetPlayer, OnTextureRowLUp, OnUpdate, OnUseCheck, OptionsToggle, POPOption, PackIconOver, PackTab, PackTabMover, PickedListMenu, RButtonUp, RemoveChar, ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip, SelectChar, SelectDye, SetAssist, SetCurrentHT, ShowFiltersSubMenu, ShowItemTooltip, ShowMouseOverHoverFrame, ShowPersonalStatsTooltip, ShowRowTooltip, ShowTooltip, ShowTopKillersTooltip, SpawnSelectionMenu, SplashTextLClick, StartTooltip, TalismanAlerter_TalismanAlerterIndicator_OnMouseOver, TestTooltip, TimerUI_OnMouseOver, UnitFramesUI_ConfigDialog_OnClickCastingsListLButtonUp, UnitFramesUI_ConfigDialog_OnEffectsIndicatorsListLButtonUp, UnitFramesUI_ConfigDialog_OnUnitFramePartsListLButtonUp, UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListLButtonUp, UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionMouseOver, UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig2CommandMouseOver, UnitFramesUI_UnitFrame_OnLButtonDown, UnitFramesUI_UnitFrame_OnLButtonUp, UnitFramesUI_UnitFrame_OnMButtonDown, UnitFramesUI_UnitFrame_OnMButtonUp, UnitFramesUI_UnitFrame_OnMouseOver, UnitFramesUI_UnitFrame_OnRButtonDown, UnitFramesUI_UnitFrame_OnRButtonUp, UnitLButtonDown, UnitLButtonUp, UnitMButtonDown, UnitMButtonUp, UnitRButtonDown, UnitRButtonUp, UpdateBagModeBag, UpdatePartyTooltip, UpdatePartyTooltipFromButton, UpdatePartyTooltipFromWindow, UpdateTooltip, contextMenuOnClick, initializeUI, lua_call, onHover, onMouseOut, onMouseOver contexts.

## Seen In

- AdvancedRenownTrainer
- AggroMeter
- Assist
- Atlas
- AuctionStats
- BankArkel
- BuffHead
- CM_ClosetGoblin
- CaVES
- CastSequence
- Cheeseboard
- CraftingWillard
- Crusher
- DAoCBuff
- Dascore
- DetauntHelper
- EA_OpenPartyWindow
- EZCraft
- EZCraftX
- Emojii
- Enemy
- FozAuction
- HealGrid
- JunkDump
- Killer
- LibAddonButton
- LibSurveyor
- LootAlert
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrowLight
- Motion
- NaturalLog
- Obsidian
- PlayEffectsOn
- Pocket Palette
- PotionBar
- Pure
- Refer
- ResHelp
- RoR_SoR
- RvRContribution
- Sequencer
- Shinies
- SocialWindow 2.0
- Squared
- Swift Assist
- TastyButtons
- TexturedButtons
- TidyQueue
- TidyRoll
- Tokens
- Tome Titan
- Trakario
- TurretRange
- Twister
- WARCommander
- WarBoard_WarWhisperer
- nRarity
- wbLeadHelper
- zMailMod

## Related APIs

- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.HideAll](../../globals/functions/global_EA_Window_ContextMenu.HideAll.md) (HIGH 100/100) - Global Function
- [Icons.GetCareerIconIDFromCareerNamesID](../../globals/functions/global_Icons.GetCareerIconIDFromCareerNamesID.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [PartyUtils.GetWarbandLeader](../../globals/functions/global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandMember](../../globals/functions/global_PartyUtils.GetWarbandMember.md) (HIGH 100/100) - Global Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [EA_Window_WorldMap.ShowZone](../../globals/functions/global_EA_Window_WorldMap.ShowZone.md) (HIGH 90/100) - Global Function
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [PartyUtils.MoveWarbandMember](../../globals/functions/global_PartyUtils.MoveWarbandMember.md) (HIGH 88/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [PartyUtils.SwapWarbandMembers](../../globals/functions/global_PartyUtils.SwapWarbandMembers.md) (MEDIUM 45/100) - Global Function

## Used With

- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event

## Notes

- Observed in contexts: AbilityClear, AbilityCursorSwap, AbilityTooltip, AuctionHouseWindowCreateSearchItemName_LButtonDown, Autohide, BankWindow_EquipmentMouseOver
