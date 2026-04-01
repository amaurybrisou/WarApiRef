# OnRButtonDown

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 118

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, CM_ClosetGoblin, Enemy, Targets, TexturedButtons, TurretRange, bigger_MacroWindow |
| Files seen in | `/workspace/BuffHead/Display.xml:27`, `/workspace/BuffHead/Setup/General.xml:33`, `/workspace/BuffHead/Setup/SetupAdvancedCompression.xml:45`, `/workspace/BuffHead/Setup/SetupAdvancedCompressionItem.xml:44`, `/workspace/BuffHead/Setup/SetupAdvancedContainers.xml:45`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItem.xml:44`, `/workspace/BuffHead/Setup/SetupLayout.xml:127`, `/workspace/BuffHead/Setup/SetupLayout.xml:216` |
| Namespaces detected | OnRButtonDown |
| Source kinds | event_page, xml_handlers |
| Example locations | BuffHead: BuffHeadBuffTemplate.OnRButtonDown, BuffHead: BuffHeadLayoutControlFrameWindow.OnRButtonDown, BuffHead: BuffHeadLayoutFrameWindow.OnRButtonDown, BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate.OnRButtonDown, BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate.OnRButtonDown, BuffHead: BuffHeadSetupAdvancedContainersItemRowTemplate.OnRButtonDown |
| XML usage count | 22 |
| XML attribute usage count | 22 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
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

Observed as an engine-supplied UI event hook used by 7 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from API_Ref alone.

## Seen In

- BuffHead
- CM_ClosetGoblin
- Enemy
- Targets
- TexturedButtons
- TurretRange
- bigger_MacroWindow

## Registrars And Handlers

- BuffHead.Setup.AdvancedCompression.OnRowRDown
- BuffHead.Setup.AdvancedCompressionItem.OnRowRDown
- BuffHead.Setup.AdvancedContainers.OnRowRDown
- BuffHead.Setup.AdvancedContainersItem.OnRowRDown
- BuffHead.Setup.Layout.OnControlFrameRButtonDown
- BuffHead.Setup.Layout.OnLayoutWindowRButtonDown
- BuffHead.Setup.Layout.TrapClick
- BuffHead.Setup.OnRowRDown
- BuffHead.Setup.PriorityEffects.OnRowRDown
- BuffHead.Setup.PriorityEffectsItem.OnRowRDown
- BuffHead.Setup.SelectTexture.OnTextureRowRDown
- CG_ItemRack.EquipmentRButtonDown
- ClosetGoblinCharacterWindow.EquipmentRButtonDown
- EA_Window_Macro.DetailIconRButtonDown
- EA_Window_Macro.IconRButtonDown
- EA_Window_Macro.SelectionIconRButtonDown
- Enemy.UnitFramesUI_UnitFrame_OnRButtonDown
- FrameManager.OnRButtonDown
- TargetsUnitFrame.UnitRClick
- TexturedButtons.Setup.OnRowRDown
- TurretRange.Setup.Distances.OnRowRDown

## Examples

- BuffHead: BuffHeadBuffTemplate -> BuffHeadBuffTemplate.OnRButtonDown -> FrameManager.OnRButtonDown
- BuffHead: BuffHeadLayoutControlFrameWindow -> BuffHeadLayoutControlFrameWindow.OnRButtonDown -> BuffHead.Setup.Layout.OnControlFrameRButtonDown
- BuffHead: BuffHeadLayoutFrameWindow -> BuffHeadLayoutFrameWindow.OnRButtonDown -> BuffHead.Setup.Layout.OnLayoutWindowRButtonDown
- BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate -> BuffHeadSetupAdvancedCompressionItemRowTemplate.OnRButtonDown -> BuffHead.Setup.AdvancedCompressionItem.OnRowRDown
- BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate -> BuffHeadSetupAdvancedCompressionRowTemplate.OnRButtonDown -> BuffHead.Setup.AdvancedCompression.OnRowRDown
- BuffHead: BuffHeadSetupAdvancedContainersItemRowTemplate -> BuffHeadSetupAdvancedContainersItemRowTemplate.OnRButtonDown -> BuffHead.Setup.AdvancedContainersItem.OnRowRDown

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnRButtonDown](../../xml/handlers/handler_OnRButtonDown.md) (HIGH 100/100) - XML Handler
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- none
