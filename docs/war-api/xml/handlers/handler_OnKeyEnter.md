# OnKeyEnter

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
| Addons seen in | Aura, EA_UiDebugTools, MapMonster, MapPin, ObjectInspector, Pocket Palette, RandomMount, Shinies |
| Files seen in | `/workspace/Aura/Source/AuraConfig.xml:1106`, `/workspace/Aura/Source/AuraConfig.xml:1147`, `/workspace/Aura/Source/AuraConfig.xml:1212`, `/workspace/Aura/Source/AuraConfig.xml:1253`, `/workspace/Aura/Source/AuraConfig.xml:126`, `/workspace/Aura/Source/AuraConfig.xml:1357`, `/workspace/Aura/Source/AuraConfig.xml:1481`, `/workspace/Aura/Source/AuraConfig.xml:1570` |
| Namespaces detected | OnKeyEnter |
| Source kinds | bindings, xml_handlers |
| Example locations | Aura: AuraConfigActivationAlertTextText.OnKeyEnter, Aura: AuraConfigDeactivationAlertTextText.OnKeyEnter, Aura: AuraConfigGeneralName.OnKeyEnter, Aura: AuraConfigGeneralOffsetX.OnKeyEnter, Aura: AuraConfigGeneralOffsetY.OnKeyEnter, Aura: AuraConfigTimerOffsetX.OnKeyEnter |
| XML usage count | 34 |
| XML attribute usage count | 34 |
| Lua usage count | 12 |
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

Observed as an XML handler hook bound by 9 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- EditBox

## Seen In

- Aura
- EA_UiDebugTools
- MapMonster
- MapPin
- ObjectInspector
- Pocket Palette
- RandomMount
- Shinies
- bigger_MacroWindow

## Examples

- Aura: AuraConfigActivationAlertTextText -> AuraConfigActivationAlertTextText.OnKeyEnter -> none
- Aura: AuraConfigDeactivationAlertTextText -> AuraConfigDeactivationAlertTextText.OnKeyEnter -> none
- Aura: AuraConfigGeneralName -> AuraConfigGeneralName.OnKeyEnter -> none
- Aura: AuraConfigGeneralOffsetX -> AuraConfigGeneralOffsetX.OnKeyEnter -> none
- Aura: AuraConfigGeneralOffsetY -> AuraConfigGeneralOffsetY.OnKeyEnter -> none
- Aura: AuraConfigTimerOffsetX -> AuraConfigTimerOffsetX.OnKeyEnter -> none

## Related APIs

- none

## Used With

- [EA_TextEntryGroupEntryBoxTextInput](../../globals/tables/table_EA_TextEntryGroupEntryBoxTextInput.md) (HIGH 100/100) - Global Table
- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [OnKeyEnter](../../events/window_events/window_event_OnKeyEnter.md) (HIGH 100/100) - Window Event

## Triggered By

- none

## Affects

- [EA_TextEntryGroupEntryBoxTextInput](../../globals/tables/table_EA_TextEntryGroupEntryBoxTextInput.md) (HIGH 100/100) - Global Table

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
