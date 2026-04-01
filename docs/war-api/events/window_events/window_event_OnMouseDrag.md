# OnMouseDrag

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 106

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, bigger_MacroWindow |
| Files seen in | `/workspace/Enemy/Code/Assist/AssistConfiguration.xml:130`, `/workspace/Enemy/Code/Assist/AssistConfiguration.xml:141`, `/workspace/Enemy/Code/Core/ConfigurationWindow.xml:417`, `/workspace/Enemy/Code/ScenarioInfo/ScenarioInfoConfiguration.xml:33`, `/workspace/bigger_macrowindow/Source/MacroWindow.xml:9` |
| Namespaces detected | OnMouseDrag |
| Source kinds | event_page, xml_handlers |
| Example locations | Enemy: EnemyAssistConfigurationMacroMark.OnMouseDrag, Enemy: EnemyAssistConfigurationMacroTarget.OnMouseDrag, Enemy: EnemyConfigurationWindow_MacroTemplateButton.OnMouseDrag, Enemy: EnemyScenarioInfoConfigurationMacroToggle.OnMouseDrag, bigger_MacroWindow: EA_Window_MacroIconButton.OnMouseDrag |
| XML usage count | 5 |
| XML attribute usage count | 5 |
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

Observed as an engine-supplied UI event hook used by 2 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- Enemy
- bigger_MacroWindow

## Registrars And Handlers

- EA_Window_Macro.IconMouseDrag
- Enemy.AssistUI_ConfigDialog_OnMacroMarkMouseDrag
- Enemy.AssistUI_ConfigDialog_OnMacroTargetMouseDrag
- Enemy.ConfigurationWindow_OnMacroMouseDrag
- Enemy.ScenarioInfoUI_ConfigDialog_OnMacroToggleMouseDrag

## Examples

- Enemy: EnemyAssistConfigurationMacroMark -> EnemyAssistConfigurationMacroMark.OnMouseDrag -> Enemy.AssistUI_ConfigDialog_OnMacroMarkMouseDrag
- Enemy: EnemyAssistConfigurationMacroTarget -> EnemyAssistConfigurationMacroTarget.OnMouseDrag -> Enemy.AssistUI_ConfigDialog_OnMacroTargetMouseDrag
- Enemy: EnemyConfigurationWindow_MacroTemplateButton -> EnemyConfigurationWindow_MacroTemplateButton.OnMouseDrag -> Enemy.ConfigurationWindow_OnMacroMouseDrag
- Enemy: EnemyScenarioInfoConfigurationMacroToggle -> EnemyScenarioInfoConfigurationMacroToggle.OnMouseDrag -> Enemy.ScenarioInfoUI_ConfigDialog_OnMacroToggleMouseDrag
- bigger_MacroWindow: EA_Window_MacroIconButton -> EA_Window_MacroIconButton.OnMouseDrag -> EA_Window_Macro.IconMouseDrag

## Related APIs

- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](../../globals/functions/global_Cursor.PickUp.md) (HIGH 100/100) - Global Function

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](../../globals/functions/global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [OnMouseDrag](../../xml/handlers/handler_OnMouseDrag.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none

## Notes

- none
