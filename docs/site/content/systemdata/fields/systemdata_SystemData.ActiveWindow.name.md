# SystemData.ActiveWindow.name

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 92 addons

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
| Addons seen in | AdjustTheTip, AdvancedRenownTrainer, AggroMeter, AuctionStats, Aura, AutoBand, BlackBook, BuffHead |
| Files seen in | AdjustTheTip.lua, AdvancedRenownTraining.lua, AggroMeter.lua, AutoBand.lua, BlackBookWindow.lua, Button.lua, CCTV.lua, CDownSettings.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | AbilityTooltip, ActivatorMouseOver, AddCraftingItems, AddRarityItems, BackpackList_EA_Window_Backpack_ListViewInventoryLButtonDown, BeginResize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 263 |
| Global usage count | 263 |
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

SystemData.SystemData.ActiveWindow.name field accessed by 92 addons; commonly found in AbilityTooltip and ActivatorMouseOver, AddCraftingItems, AddRarityItems, BackpackList_EA_Window_Backpack_ListViewInventoryLButtonDown, BeginResize, BuffsOnMouseOver, ChangeFont, ChangeGroupSpell, ChangeSorting, ClearSelectedButton, ClearSelectedGroup, ColorListOnLButtonUp, CombatLogUI_SnapshotWindow_ListRowMouseOver, CombatLogUI_StatsWindow_SortColumnRClick, ConfigSoundTestTip, ConfigurationWindow_OnButtonClick, ConfigurationWindow_OnChange, DK_OnClickSortButton, DK_OnMouseoverBookmark, DK_OnRollOverSortButton, DK_OnSelectProfil, DK_OnSelectProfilMini, DK_PageTopList, DestroyWindow, DisplayTooltip, EquipmentLButtonUp, EquipmentMouseOver, EquipmentRButtonUp, FilterRowOnLButtonUp, FixBankWindowTooltip, GetIconFromClick, GetSlotRowNumForActiveListRow, GetSlotRowNumForAutoRow, GetSortType, HandleDrag, HelpTip, HighScoreOnMouseOver, IconOnMouseOver, InteractionQuest_EA_Window_InteractionQuest_OnMouseOverChoiceReward, InteractionStore_EA_Window_InteractionStore_OnItemLButtonDown, InventoryLButtonDown, InventoryLButtonUp, InventoryRButtonUp, ItemSlotMouseOver, JoinGroup, KeepUnClaimDialog, LButtonUp, LayoutIdentify, LeaveScenario, LeftButtonClick, LeftListOnLButtonUp, LeftListOnMouseOver, LeftListOnRButtonUp, ListSelectRow, MenuItemSelected, MouseOver, MouseOverQuest, MouseOverRow, MouseoverRPBar, MouseoverXPBar, NewIcon, OnAbilityTT, OnAccept, OnButtonSelectListRUp, OnButtonSelected, OnChoiceChange, OnClickName, OnClickSetRow, OnClickZoneRow, OnContainerClick, OnContainerRClick, OnControlFrameLButtonDown, OnControlFrameRButtonDown, OnControlFrameRButtonUp, OnDecline, OnDropSlotMouseOver, OnEnableFilter, OnGroupIconRUp, OnGroupSelected, OnGroupSpellLClick, OnGroupSpellRClick, OnHyperLinkLButtonUp2, OnIconLButtonUp, OnIconMouseLeave, OnIconMouseOver, OnJoinGroupMouseOver, OnKeyEscape, OnLButtonDown, OnLButtonDownRow, OnLButtonDownStatusIcon, OnLButtonUp, OnLButtonUpEntryRow, OnLButtonUpTab, OnLButtonUp_Results_SortButton, OnLButtonUp_Searches_SortButton, OnLButtonUp_SortButton, OnLayersChanged, OnLayoutFrameButtonDown, OnLeaveGroupMouseOver, OnListRowLButtonUp, OnLootItemHook, OnMenuClickMakeMainAssist, OnMouseOver, OnMouseOverApplyButton, OnMouseOverCareerIcon, OnMouseOverCreateGroupStatsTooltip, OnMouseOverCreateTooltip, OnMouseOverEnd_ContextMenu_Recipe_Delete, OnMouseOverFAQButton, OnMouseOverFilterAll, OnMouseOverFilterEA, OnMouseOverFilterMenuButton, OnMouseOverFilterRV, OnMouseOverFilteremotesButton, OnMouseOverHover, OnMouseOverIcon, OnMouseOverInfluenceBar, OnMouseOverItem, OnMouseOverLabelEditBox, OnMouseOverLfgIconsLabel, OnMouseOverLowVP, OnMouseOverMagnifier, OnMouseOverMapPin, OnMouseOverMessageEditBox, OnMouseOverMessageEndLabel, OnMouseOverMessageStartLabel, OnMouseOverMessageTextColorLabel, OnMouseOverResetButton, OnMouseOverReward, OnMouseOverRewardEnd, OnMouseOverRowEnd, OnMouseOverSaveButton, OnMouseOverSettingsIcon, OnMouseOverSortBy, OnMouseOverStart, OnMouseOverStatusIcon, OnMouseOverSubmenuComboBox, OnMouseOverTypeComboBox, OnMouseOverVP, OnMouseOver_ContextMenu_Recipe_Delete, OnMouseOver_Sigil, OnMouseover, OnMouseoverBtn, OnMouseoverMailIcon, OnMouseoverScenarioGroupBtn, OnMouseoverScenarioQueue, OnMouseoverScenarioSummaryBtn, OnMouseoverVisibleButton, OnNavigateLUp, OnNavigateMouseOver, OnNubLBU, OnOtherLClick, OnOtherRClick, OnRButtonUp, OnRButtonUpIcon, OnRButtonUpTab, OnRButtonUp_Results_ListItem, OnRButtonUp_Searches, OnRaceTabSelected, OnRawDeviceInput, OnRefer, OnRowMouseOver, OnSelChanged, OnSelectCampaignMode, OnSetRowContextMenu, OnSetZoneRowContextMenu, OnSlide, OnSliderChange, OnSortButtonLUp, OnSortList, OnSpellLClick, OnSpellRClick, OnTabLBU, OnTabRight, OnTabSelect, OnTargetTypeLUp, OnTestButtonToolTip, OnTextChanged, OnToggleHidden, OnToolsUp, OnUpdateActionBar, OnUpdateHealthBar, OpenColorDialog, OpenTome, OptionsCHKclick, RButtonUp, Results_GetSlotRowNumForActiveListRow, RightListOnLButtonUp, RightListOnMouseOver, RightListOnRButtonUp, Searches_GetSlotRowNumForActiveListRow, SelectActive, SelectColor, SelectColorGroup, SelectExpire, SelectOption, SelectQueue, SelectTab, SetFont, SetItemDye, SetTactic, Show, ShowAbout, ShowAurasMenu, ShowContext, ShowDropdownMenu, ShowFontSelection, ShowOptions, ShowOverrideMenu, ShowStatus, ShowTooltip, SortChange, SwitchTab, TOGGLE_TYPE, TabColorChangeBackgroundColor, TabColorChangeSignColor, TacticRemoved, TalismanMakingWindow_OnSlotMouseOver, ToggleHover, ToggleVisibility, TomeClearNewEntriesMOver, TomeQuest_TomeWindow_OnLButtonDownChoiceReward, TomeQuest_TomeWindow_OnLButtonDownGivenReward, TomeWarJournal_TomeWindow_OnLButtonDownInfluenceRewardLevel1, TomeWarJournal_TomeWindow_OnLButtonDownInfluenceRewardLevel2, TomeWarJournal_TomeWindow_OnLButtonDownInfluenceRewardLevel3, Tooltip, UI_ChooseIconDialog_OnListRowLButtonUp, UI_ChooseIconDialog_OnListRowRButtonUp, UnitFramesUI_Anchor_OnLButtonDown, UnitFramesUI_Anchor_OnLButtonUp, UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionMouseOver, UpdateBagModeBag, UpdateBagModeOnOff, UpdateBankSlot, UpdateGoldReportOnOff, UpdateIgnoreProfessionRarityOnOff, UpdateItemReportOnOff, UpdateRarity, UpdateSellApothOnOff, UpdateSellCultivationOnOff, UpdateSellSalvagingOnOff, UpdateSellTalismanOnOff, WindowHiderShowClicked, createTooltip, deleteIt, getSelectedPriorityData, lua_call, nLL_nLootLinkGUI_itemMouseDownL, onClickNutrient, onClickSeed, onClickSoil, onClickWater, onHHover, onHover, onHoverNutrient, onHoverSeed, onHoverSoil, onHoverWater, onRepeat, setColumn, toChat contexts.

## Seen In

- AdjustTheTip
- AdvancedRenownTrainer
- AggroMeter
- AuctionStats
- Aura
- AutoBand
- BlackBook
- BuffHead
- CCTV
- CDown
- CM_ClosetGoblin
- CaVES
- Calling
- CastSequence
- Crusher
- DAoCBuff
- DPSMeter
- DammazKron
- DetauntHelper
- DuffTimer
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EZCraftX
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- FixGit
- Group Icons SG
- GroupSpotter
- HealGrid
- HideHiddenFrames
- Hopper
- JunkDump
- KeyBar
- KillTracker
- Kwestor
- LibAddonButton
- MacroIcons
- Map
- MapPin
- MarkBuff
- Minmap
- Miracle Grow Remix
- MiracleGrow
- Motion
- NPC Item Sale Price
- NerfedButtons
- Obsidian
- OneClickGTAOE
- PeaceOut
- Pocket Palette
- PotionBar
- Pure
- QuickTacticSwitch
- QuickWarReport
- RVMOD_Manager
- RVMOD_PlayerStatus
- RaidMeter
- RandomMount
- RealmStatus
- Refer
- RoR_SoR
- SOR
- SessionRPs
- Shinies
- Squared
- Statdoll
- TacticSetNames
- TastyButtons
- TexturedButtons
- TidyChat
- TidyRoll
- TokenMachine
- Tome Titan
- TomeTracker
- TurretRange
- TwisterSet
- Vectors
- WarBoard_Menu
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- XpStatus+G
- emotes
- followTheLeader
- minesweep
- nRarity
- scenarioInfo
- talisman-monitor
- wbLeadHelper
- xHUD
- zMailMod

## Related APIs

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetStayDownFlag](../../window_api/functions/window_ButtonSetStayDownFlag.md) (HIGH 100/100) - Window Function
- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCurrentBackpackType](../../globals/functions/global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetItemsFromBackpack](../../globals/functions/global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetSlotFromActionButtonGroup](../../globals/functions/global_EA_Window_Backpack.GetSlotFromActionButtonGroup.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddUserDefinedMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddUserDefinedMenuItem.md) (HIGH 100/100) - Global Function
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [WindowGetGameActionData](../../window_api/functions/window_WindowGetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetMovable](../../window_api/functions/window_WindowGetMovable.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [EA_ChatTabManager.GetTabName](../../globals/functions/global_EA_ChatTabManager.GetTabName.md) (HIGH 98/100) - Global Function
- [LayoutEditor.Hide](../../window_api/functions/window_LayoutEditor.Hide.md) (HIGH 90/100) - Window Function
- [LayoutEditor.Show](../../window_api/functions/window_LayoutEditor.Show.md) (HIGH 90/100) - Window Function
- [WindowSetResizing](../../window_api/functions/window_WindowSetResizing.md) (HIGH 90/100) - Window Function
- [DefaultColor.SetLabelColor](../../globals/functions/global_DefaultColor.SetLabelColor.md) (HIGH 88/100) - Global Function
- [OnKeyEscape](../../xml/handlers/handler_OnKeyEscape.md) (HIGH 88/100) - XML Event
- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnSelChanged](../../xml/handlers/handler_OnSelChanged.md) (HIGH 88/100) - XML Event
- [OnSlide](../../xml/handlers/handler_OnSlide.md) (HIGH 88/100) - XML Event
- [OnTextChanged](../../xml/handlers/handler_OnTextChanged.md) (HIGH 88/100) - XML Event
- [EA_ChatWindow.InsertQuestLink](../../globals/functions/global_EA_ChatWindow.InsertQuestLink.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionStore.OnItemLButtonDown](../../globals/functions/global_EA_Window_InteractionStore.OnItemLButtonDown.md) (HIGH 80/100) - Global Function
- [EA_Window_Loot.OnLootItem](../../globals/functions/global_EA_Window_Loot.OnLootItem.md) (HIGH 80/100) - Global Function
- [WindowGameAction](../../window_api/functions/window_WindowGameAction.md) (HIGH 80/100) - Window Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Used With

- [EA_BackpackUtilsMediator.GetCurrentBackpackType](../../globals/functions/global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetSlotFromActionButtonGroup](../../globals/functions/global_EA_Window_Backpack.GetSlotFromActionButtonGroup.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event

## Notes

- Observed in contexts: AbilityTooltip, ActivatorMouseOver, AddCraftingItems, AddRarityItems, BackpackList_EA_Window_Backpack_ListViewInventoryLButtonDown, BeginResize
