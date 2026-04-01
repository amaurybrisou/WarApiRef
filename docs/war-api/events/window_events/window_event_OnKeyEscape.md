# OnKeyEscape

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
| Addons seen in | Aura, EA_UiDebugTools, Enemy, MapPin, Pocket Palette, Shinies, bigger_MacroWindow, wbLeadHelper |
| Files seen in | `/workspace/Aura/Source/AuraConfig.xml:1107`, `/workspace/Aura/Source/AuraConfig.xml:1148`, `/workspace/Aura/Source/AuraConfig.xml:1213`, `/workspace/Aura/Source/AuraConfig.xml:1254`, `/workspace/Aura/Source/AuraConfig.xml:127`, `/workspace/Aura/Source/AuraConfig.xml:1358`, `/workspace/Aura/Source/AuraConfig.xml:1482`, `/workspace/Aura/Source/AuraConfig.xml:1571` |
| Namespaces detected | OnKeyEscape |
| Source kinds | event_page, xml_handlers |
| Example locations | Aura: AuraConfigActivationAlertTextText.OnKeyEscape, Aura: AuraConfigDeactivationAlertTextText.OnKeyEscape, Aura: AuraConfigGeneralName.OnKeyEscape, Aura: AuraConfigGeneralOffsetX.OnKeyEscape, Aura: AuraConfigGeneralOffsetY.OnKeyEscape, Aura: AuraConfigTimerOffsetX.OnKeyEscape |
| XML usage count | 50 |
| XML attribute usage count | 50 |
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

Observed as an engine-supplied UI event hook used by 8 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- Aura
- EA_UiDebugTools
- Enemy
- MapPin
- Pocket Palette
- Shinies
- bigger_MacroWindow
- wbLeadHelper

## Registrars And Handlers

- DebugWindow.OnKeyEscape
- DebugWindow.TextClear
- DevPadWindow.OnKeyEscape
- Enemy.CombatLogUI_SnapshotWindow_Hide
- Enemy.CombatLogUI_StatsWindow_Hide
- Enemy.GroupsUI_EffectFilterDialog_Hide
- Enemy.IntercomUI_ChooseChannelDialog_Hide
- Enemy.IntercomUI_IntercomDialog_Hide
- Enemy.IntercomUI_IntercomJoinDialog_Hide
- Enemy.MarksUI_MarkConfigDialog_Hide
- Enemy.UI_ChooseIconDialog_Hide
- Enemy.UI_ConfigDialog_Hide
- Enemy.UI_TextEntryDialog_Close
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Hide
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Hide
- Enemy.UnitFramesUI_UnitFramePartDialog_Hide
- MapPin.UI_ChooseIconDialog_Hide
- WbLeadHelperMessage.MessageDialogHide
- none
- wbLeadHelperChooseIcon.Hide

## Examples

- Aura: AuraConfigActivationAlertTextText -> AuraConfigActivationAlertTextText.OnKeyEscape -> none
- Aura: AuraConfigDeactivationAlertTextText -> AuraConfigDeactivationAlertTextText.OnKeyEscape -> none
- Aura: AuraConfigGeneralName -> AuraConfigGeneralName.OnKeyEscape -> none
- Aura: AuraConfigGeneralOffsetX -> AuraConfigGeneralOffsetX.OnKeyEscape -> none
- Aura: AuraConfigGeneralOffsetY -> AuraConfigGeneralOffsetY.OnKeyEscape -> none
- Aura: AuraConfigTimerOffsetX -> AuraConfigTimerOffsetX.OnKeyEscape -> none

## Related APIs

- none

## Used With

- [EditBox](../../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [OnKeyEscape](../../xml/handlers/handler_OnKeyEscape.md) (HIGH 100/100) - XML Handler
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- none
