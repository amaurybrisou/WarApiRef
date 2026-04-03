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
| Addons seen in | AdvancedRenownTrainer, BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy, MoraleCircle, PotionBar, RoR_SoR |
| Files seen in | `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainers.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupEffectCache.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayout.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupMenu.xml:0` |
| Namespaces detected | OnLButtonDown |
| Source kinds | event_page, flows, lua_event_registration, xml_handlers |
| Example locations | AdvancedRenownTrainer: AbilityButtonTemplate.OnLButtonDown, BuffHead: BuffHeadLayoutControlFrameWindow.OnLButtonDown, BuffHead: BuffHeadLayoutFrameWindow.OnLButtonDown, BuffHead: BuffHeadLayoutResizeButton.OnLButtonDown, BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate.OnLButtonDown, BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate.OnLButtonDown |
| XML usage count | 42 |
| XML attribute usage count | 42 |
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

Observed as an engine-supplied UI event hook used by 13 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- AdvancedRenownTrainer
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- MoraleCircle
- PotionBar
- RoR_SoR
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
- ClosetGoblinCharacterWindow.EquipmentLButtonDown
- EA_Window_Macro.DetailIconLButtonDown
- EA_Window_Macro.SelectionIconLButtonDown
- Enemy.AssistUI_Target_OnLButtonDown
- Enemy.CombatLogUI_IDSAnchor_OnLButtonDown
- Enemy.CombatLogUI_IDS_OnRowLButtonDown
- Enemy.GroupIcons_GroupIcon_OnLButtonDown
- Enemy.KillSpamUI_KillSpamDialog_OnLButtonDown
- Enemy.MarkUI_EnemyMark_OnLButtonDown
- Enemy.MarksUI_EnemyMarksWindow_OnLButtonDown
- Enemy.UnitFramesUI_Anchor_OnLButtonDown
- Enemy.UnitFramesUI_UnitFrame_OnLButtonDown
- MoraleCircle.Click
- PotionBarFloating.ActivatorLButtonDown
- PotionBarFloating.LButtonDown
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

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
