# OnLButtonDown

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
| Addons seen in | AdvancedRenownTrainer, BuffHead, CM_ClosetGoblin, CMap, DAoCBuff, DaemonAssist, EA_UiDebugTools, Effigy |
| Files seen in | `/workspace_addons/BuffHead/Setup/General.xml:31`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompression.xml:43`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItem.xml:42`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainers.xml:43`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItem.xml:42`, `/workspace_addons/BuffHead/Setup/SetupEffectCache.xml:49`, `/workspace_addons/BuffHead/Setup/SetupLayout.xml:125`, `/workspace_addons/BuffHead/Setup/SetupLayout.xml:215` |
| Namespaces detected | OnLButtonDown |
| Source kinds | event_page, flows, lua_event_registration, xml_handlers |
| Example locations | AdvancedRenownTrainer: AbilityButtonTemplate.OnLButtonDown, BuffHead: BuffHeadLayoutControlFrameWindow.OnLButtonDown, BuffHead: BuffHeadLayoutFrameWindow.OnLButtonDown, BuffHead: BuffHeadLayoutResizeButton.OnLButtonDown, BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate.OnLButtonDown, BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate.OnLButtonDown |
| XML usage count | 66 |
| XML attribute usage count | 66 |
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

Observed as an engine-supplied UI event hook used by 23 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- AdvancedRenownTrainer
- BuffHead
- CM_ClosetGoblin
- CMap
- DAoCBuff
- DaemonAssist
- EA_UiDebugTools
- Effigy
- Enemy
- GetStats
- Miracle Grow Remix
- MoraleCircle
- PotionBar
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RoR_SoR
- Targets
- TexturedButtons
- TidyChat
- TurretRange
- WhoHealedMe
- bigger_MacroWindow

## Registrars And Handlers

- AdvancedRenownTraining.Select
- BuffHead.Setup.AdvancedCompression.OnRowLDown
- BuffHead.Setup.AdvancedCompressionItem.OnRowLDown
- BuffHead.Setup.AdvancedContainers.OnRowLDown
- BuffHead.Setup.AdvancedContainersItem.OnRowLDown
- BuffHead.Setup.EffectCache.OnRowLDown
- BuffHead.Setup.Layout.BeginResize
- BuffHead.Setup.Layout.ClearSelection
- BuffHead.Setup.Layout.OnControlFrameLButtonDown
- BuffHead.Setup.Layout.OnLayoutWindowLButtonDown
- BuffHead.Setup.OnRowLDown
- BuffHead.Setup.PriorityEffects.OnRowLDown
- BuffHead.Setup.PriorityEffectsItem.OnRowLDown
- BuffHead.Setup.SelectTexture.OnTextureRowLDown
- CMapWindow.OnLButtonDown
- CMapWindow.OnResizeBeginLO
- CMapWindow.OnResizeBeginLU
- CMapWindow.OnResizeBeginRO
- CMapWindow.OnResizeBeginRU
- ClosetGoblinCharacterWindow.EquipmentLButtonDown
- DaemonAssist.HideWindow
- DaemonAssist.OnToggleButtonDown
- DebugWindow.OnResizeBegin
- DevPad.OnResizeBegin
- EA_Window_Macro.DetailIconLButtonDown
- EA_Window_Macro.SelectionIconLButtonDown
- Effigy.LButtonDown
- Enemy.AssistUI_Target_OnLButtonDown
- Enemy.CombatLogUI_IDSAnchor_OnLButtonDown
- Enemy.CombatLogUI_IDS_OnRowLButtonDown
- Enemy.GroupIcons_GroupIcon_OnLButtonDown
- Enemy.KillSpamUI_KillSpamDialog_OnLButtonDown
- Enemy.MarkUI_EnemyMark_OnLButtonDown
- Enemy.MarksUI_EnemyMarksWindow_OnLButtonDown
- Enemy.UnitFramesUI_Anchor_OnLButtonDown
- Enemy.UnitFramesUI_UnitFrame_OnLButtonDown
- GetStats.RefreshStats
- MiracleGrow2.LayoutStartDrag
- MiracleGrow2.onHClick
- MoraleCircle.Click
- ObjectInspector.OnResizeBegin
- PotionBarFloating.ActivatorLButtonDown
- PotionBarFloating.LButtonDown
- QuickWarReport.OnConfirmDecline
- QuickWarReport.OnConfirmNoop
- QuickWarReport.OnTierButtonClick
- RVAPI_ColorDialog.OnLButtonDownColor
- RVMOD_Manager.OnLButtonDownRow
- RVMOD_Manager.OnLButtonDownSettingsIcon
- RVMOD_Manager.OnLButtonDownStatusIcon
- TargetsUnitFrame.UnitLClick
- TexturedButtons.Setup.OnRowLDown
- TidyChat.OnEntryBoxTextInputLBD
- TurretRange.Setup.Distances.OnRowLDown
- WHMGui.OnHealerNameMouseDown
- WindowRegisterCoreEventHandler
- WindowUtils.TrapClick
- core

## Examples

- TidyChat: TidyChatFrames.Initialize -> OnLButtonDown -> TidyChat.OnEntryBoxTextInputLBD
- AdvancedRenownTrainer: AbilityButtonTemplate -> AbilityButtonTemplate.OnLButtonDown -> AdvancedRenownTraining.Select
- BuffHead: BuffHeadLayoutControlFrameWindow -> BuffHeadLayoutControlFrameWindow.OnLButtonDown -> BuffHead.Setup.Layout.OnControlFrameLButtonDown
- BuffHead: BuffHeadLayoutFrameWindow -> BuffHeadLayoutFrameWindow.OnLButtonDown -> BuffHead.Setup.Layout.OnLayoutWindowLButtonDown
- BuffHead: BuffHeadLayoutResizeButton -> BuffHeadLayoutResizeButton.OnLButtonDown -> BuffHead.Setup.Layout.BeginResize
- BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate -> BuffHeadSetupAdvancedCompressionItemRowTemplate.OnLButtonDown -> BuffHead.Setup.AdvancedCompressionItem.OnRowLDown

## Related APIs

- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [BroadcastEvent](../../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- none
