# OnRButtonUp

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 188

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AggroMeter, AnywhereTrainer, AnywhereTrainerAdditions, Aura, BuffHead, Busted, CM_ClosetGoblin, CMap |
| Files seen in | `/workspace/AggroMeter/AggroMeter.xml:135`, `/workspace/AggroMeter/AggroMeter.xml:8`, `/workspace/AnywhereTrainer/source/AnywhereTrainer.xml:50`, `/workspace/AnywhereTrainerAdditions/AnywhereTrainerAdditions.xml:6`, `/workspace/Aura/Source/AuraSettings.xml:28`, `/workspace/Aura/Source/AuraShares.xml:81`, `/workspace/BuffHead/Setup/General.xml:34`, `/workspace/BuffHead/Setup/SetupAdvancedCompression.xml:46` |
| Namespaces detected | OnRButtonUp |
| Source kinds | event_page, flows, lua_event_registration, xml_handlers |
| Example locations | AggroMeter: AggroMeterGrayListBox.OnRButtonUp, AggroMeter: AggroMeter_Button.OnRButtonUp, AnywhereTrainer: AnywhereTrainerTabTemplate.OnRButtonUp, AnywhereTrainerAdditions: AnywhereTrainerAdditionsTabTemplate.OnRButtonUp, Aura: AuraSharesRow.OnRButtonUp, Aura: AuraWindowRow.OnRButtonUp |
| XML usage count | 150 |
| XML attribute usage count | 150 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 3 |
| Initialization flow references | 1 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
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

Observed as an engine-supplied UI event hook used by 36 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- AggroMeter
- AnywhereTrainer
- AnywhereTrainerAdditions
- Aura
- BuffHead
- Busted
- CM_ClosetGoblin
- CMap
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- JunkDump
- Killer
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrowLight
- MoraleCircle
- PeaceOut
- Pocket Palette
- PotionBar
- Queue Queuer
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyQueue
- TidyRoll
- Tortall_DPS
- TurretRange
- WarBoard
- followTheLeader
- wbLeadHelper

## Registrars And Handlers

- AggroMeter.OnTabRBU
- AggroMeter.PickedListMenu
- AnywhereTrainer.OnRButtonUp
- AnywhereTrainerAdditions.OnRButtonUp
- AuraSettings.OnRButtonUpAuraList
- AuraShares.OnRButtonUpAuraList
- BuffHead.Setup.AdvancedCompression.OnRowRUp
- BuffHead.Setup.AdvancedCompressionItem.OnRowRUp
- BuffHead.Setup.AdvancedContainers.OnRowRUp
- BuffHead.Setup.AdvancedContainersItem.OnContainerRClick
- BuffHead.Setup.AdvancedContainersItem.OnRowRUp
- BuffHead.Setup.EffectCache.OnRowRUp
- BuffHead.Setup.Filter.OnRowRUp
- BuffHead.Setup.Layout.OnControlFrameRButtonUp
- BuffHead.Setup.OnRowRUp
- BuffHead.Setup.PriorityEffects.OnRowRUp
- BuffHead.Setup.PriorityEffectsItem.OnRowRUp
- BuffHead.Setup.SelectTexture.OnTextureRowRUp
- BustedGUI.ClearAlertFlash
- CMapWindow.OnScenarioQueueRButtonUp
- ClosetGoblinCharacterWindow.EquipmentRButtonUp
- ClosetGoblinCharacterWindow.OnSetRowContextMenu
- ClosetGoblinOptionWindow.OnRButtonUp
- ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
- DAoCBuffFrame.OnRButtonUp
- Effigy.RButtonUp
- Enemy.CombatLogUI_EpsWindow_OnRButtonUp
- Enemy.CombatLogUI_StatsWindow_SortColumnRClick
- Enemy.CombatLogUI_TargetDefenseTotalWindow_OnRButtonUp
- Enemy.CombatLogUI_TargetDefenseWindow_OnRButtonUp
- Enemy.KillSpamUI_KillSpamDialog_OnRButtonUp
- Enemy.KillSpamUI_KillSpamDialog_OnRowRButtonUp
- Enemy.KillSpamUI_PlayerKDR_OnRButtonUp
- Enemy.MarkUI_EnemyMark_OnRButtonUp
- Enemy.MarksUI_EnemyMarkIcon_OnRButtonUp
- Enemy.MarksUI_EnemyMarksWindow_OnAddRButtonUp
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_SortColumnRClick
- Enemy.StopwatchReset
- Enemy.TimerUI_OnRButtonUp
- Enemy.UI_Debug_OnRButtonUp
- Enemy.UI_Icon_OnRButtonUp
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionRButtonUp
- Enemy.UnitFramesUI_UnitFrame_OnRButtonUp
- FrameManager.OnRButtonUp
- JunkDumpOptions.Show
- Killer.OnPersonalCounterRButtonUp
- Map.OnRClickMap
- MapMonster.OnMouseRightClickFilter
- MapMonster.OnMouseRightClickPin
- MapMonster.ToggleFiltersMainMenu
- MapPin.OptionsMenu
- MapPin.RButtonUp
- MapPin.UI_ChooseIconDialog_OnListRowRButtonUp
- MiracleGrow2.ToggleRClick
- MiracleGrow2.onRClick
- MiracleGrowLight.switchBackground
- MiracleGrowLight.switchMode
- MoraleCircle.RightClick
- PP.ItemSlotRMouse
- PeaceOut.OnRButtonUp
- PotionBarFloating.ActivatorRButtonUp
- PotionBarFloating.RButtonUp
- QueueQueuer_GUI.MapButton_OnRButtonUp
- RandomMountUI.OnRowRightClick
- RoR_SoR.BroadCastOption
- RoR_SoR.OnTabRBU
- RoR_SoR.POPOption
- ShiniesAuctionsUI.OnRButtonUp_Results_ListItem
- ShiniesBrowseUI.OnRButtonUp_Results_ListItem
- ShiniesBrowseUI.OnRButtonUp_Searches
- ShiniesPostUI.OnRButtonUp_Results_ListItem
- TOLSettingsUI.OnEventSelChanged
- TexturedButtons.Setup.OnRowRUp
- TidyChat.ToggleOptions
- TidyQueue.OnAutojoinLBU
- TidyQueue.OnCheckboxRBU
- TortallDPSDetail.ShowColumnMenu
- TortallDPSMeter.OnRButtonUp
- TurretRange.Setup.Distances.OnRowRUp
- WarBoard.OpenLayoutMenu
- WindowRegisterCoreEventHandler
- core
- followTheLeader.OnRButtonUp
- none
- wbLeadHelper.onRMB

## Examples

- TidyChat: TidyChatFrames.Initialize -> OnRButtonUp -> TidyChat.ToggleOptions
- AggroMeter: AggroMeterGrayListBox -> AggroMeterGrayListBox.OnRButtonUp -> AggroMeter.PickedListMenu
- AggroMeter: AggroMeter_Button -> AggroMeter_Button.OnRButtonUp -> AggroMeter.OnTabRBU
- AnywhereTrainer: AnywhereTrainerTabTemplate -> AnywhereTrainerTabTemplate.OnRButtonUp -> AnywhereTrainer.OnRButtonUp
- AnywhereTrainerAdditions: AnywhereTrainerAdditionsTabTemplate -> AnywhereTrainerAdditionsTabTemplate.OnRButtonUp -> AnywhereTrainerAdditions.OnRButtonUp
- Aura: AuraSharesRow -> AuraSharesRow.OnRButtonUp -> AuraShares.OnRButtonUpAuraList

## Related APIs

- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](../../window_api/functions/window_WindowStopAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowGetMovable](../../window_api/functions/window_WindowGetMovable.md) (HIGH 90/100) - Window Function
- [PlayerMenuWindow.ShowMenu](../../globals/functions/global_PlayerMenuWindow.ShowMenu.md) (HIGH 88/100) - Global Function

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](../../xml/element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](../../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../../xml/element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- none

## Notes

- none
