# OnLButtonDown

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 138

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, BuffHead, CM_ClosetGoblin, CMap, DAoCBuff, DaemonAssist, EA_UiDebugTools, Effigy |
| Files seen in | `/workspace/BuffHead/Setup/General.xml:31`, `/workspace/BuffHead/Setup/SetupAdvancedCompression.xml:43`, `/workspace/BuffHead/Setup/SetupAdvancedCompressionItem.xml:42`, `/workspace/BuffHead/Setup/SetupAdvancedContainers.xml:43`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItem.xml:42`, `/workspace/BuffHead/Setup/SetupEffectCache.xml:49`, `/workspace/BuffHead/Setup/SetupLayout.xml:125`, `/workspace/BuffHead/Setup/SetupLayout.xml:215` |
| Namespaces detected | OnLButtonDown |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedRenownTrainer: AbilityButtonTemplate.OnLButtonDown, BuffHead: BuffHeadLayoutControlFrameWindow.OnLButtonDown, BuffHead: BuffHeadLayoutFrameWindow.OnLButtonDown, BuffHead: BuffHeadLayoutResizeButton.OnLButtonDown, BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate.OnLButtonDown, BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate.OnLButtonDown |
| XML usage count | 65 |
| XML attribute usage count | 65 |
| Lua usage count | 65 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
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

Observed as an XML handler hook bound by 21 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Button
- DynamicImage
- Label
- Window

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
- TurretRange
- bigger_MacroWindow

## Examples

- AdvancedRenownTrainer: AbilityButtonTemplate -> AbilityButtonTemplate.OnLButtonDown -> AdvancedRenownTraining.Select
- BuffHead: BuffHeadLayoutControlFrameWindow -> BuffHeadLayoutControlFrameWindow.OnLButtonDown -> BuffHead.Setup.Layout.OnControlFrameLButtonDown
- BuffHead: BuffHeadLayoutFrameWindow -> BuffHeadLayoutFrameWindow.OnLButtonDown -> BuffHead.Setup.Layout.OnLayoutWindowLButtonDown
- BuffHead: BuffHeadLayoutResizeButton -> BuffHeadLayoutResizeButton.OnLButtonDown -> BuffHead.Setup.Layout.BeginResize
- BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate -> BuffHeadSetupAdvancedCompressionItemRowTemplate.OnLButtonDown -> BuffHead.Setup.AdvancedCompressionItem.OnRowLDown
- BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate -> BuffHeadSetupAdvancedCompressionRowTemplate.OnLButtonDown -> BuffHead.Setup.AdvancedCompression.OnRowLDown

## Related APIs

- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [BroadcastEvent](../../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function

## Used With

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](../element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [GameData.PlayerActions.DO_MACRO](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_MACRO.md) (HIGH 100/100) - GameData Field
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.TARGET_PET](../../systemdata/fields/systemdata_SystemData.Events.TARGET_PET.md) (HIGH 100/100) - SystemData Field
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [BroadcastEvent](../../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
