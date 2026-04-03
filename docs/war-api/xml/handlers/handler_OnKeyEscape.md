# OnKeyEscape

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
| Addons seen in | Aura, Enemy, Pocket Palette, Shinies, bigger_MacroWindow |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogSnapshotWindow.xml:0`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogStatsWindow.xml:0`, `/workspace/data/raw/Enemy/Code/Core/ChooseIconDialog.xml:0`, `/workspace/data/raw/Enemy/Code/Core/Common.xml:0`, `/workspace/data/raw/Enemy/Code/Core/ConfigDialog.xml:0`, `/workspace/data/raw/Enemy/Code/Core/Groups/EffectFilterDialog.xml:0` |
| Namespaces detected | OnKeyEscape |
| Source kinds | bindings, xml_handlers |
| Example locations | Aura: AuraConfigActivationAlertTextText.OnKeyEscape, Aura: AuraConfigDeactivationAlertTextText.OnKeyEscape, Aura: AuraConfigGeneralName.OnKeyEscape, Aura: AuraConfigGeneralOffsetX.OnKeyEscape, Aura: AuraConfigGeneralOffsetY.OnKeyEscape, Aura: AuraConfigTimerOffsetX.OnKeyEscape |
| XML usage count | 35 |
| XML attribute usage count | 35 |
| Lua usage count | 13 |
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

Observed as an XML handler hook bound by 5 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- EditBox
- Window

## Seen In

- Aura
- Enemy
- Pocket Palette
- Shinies
- bigger_MacroWindow

## Examples

- Aura: AuraConfigActivationAlertTextText -> AuraConfigActivationAlertTextText.OnKeyEscape -> none
- Aura: AuraConfigDeactivationAlertTextText -> AuraConfigDeactivationAlertTextText.OnKeyEscape -> none
- Aura: AuraConfigGeneralName -> AuraConfigGeneralName.OnKeyEscape -> none
- Aura: AuraConfigGeneralOffsetX -> AuraConfigGeneralOffsetX.OnKeyEscape -> none
- Aura: AuraConfigGeneralOffsetY -> AuraConfigGeneralOffsetY.OnKeyEscape -> none
- Aura: AuraConfigTimerOffsetX -> AuraConfigTimerOffsetX.OnKeyEscape -> none

## Related APIs

- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Used With

- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
