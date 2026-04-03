# OnKeyEnter

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 93/100

## Confidence Assessment

- Level: HIGH

- Score: 93/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Aura, Pocket Palette, Shinies, bigger_MacroWindow |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/PocketPalette/PocketPalette.xml:0`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Auctions.xml:0`, `/workspace/data/raw/bigger_macrowindow/Source/MacroWindow.xml:0` |
| Namespaces detected | OnKeyEnter |
| Source kinds | xml_handlers |
| Example locations | Aura: AuraConfigActivationAlertTextText.OnKeyEnter, Aura: AuraConfigDeactivationAlertTextText.OnKeyEnter, Aura: AuraConfigGeneralName.OnKeyEnter, Aura: AuraConfigGeneralOffsetX.OnKeyEnter, Aura: AuraConfigGeneralOffsetY.OnKeyEnter, Aura: AuraConfigTimerOffsetX.OnKeyEnter |
| XML usage count | 22 |
| XML attribute usage count | 22 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
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

Observed as an XML handler hook bound by 4 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- EditBox

## Seen In

- Aura
- Pocket Palette
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

- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type

## Used With

- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
