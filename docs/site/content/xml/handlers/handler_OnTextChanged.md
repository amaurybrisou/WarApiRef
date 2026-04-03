# OnTextChanged

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
| Addons seen in | Aura, BuffHead, Enemy, Killer, Pocket Palette, Shinies, TexturedButtons, TurretRange |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupDisplay.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupEffectCache.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayoutProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupPriorityEffectsItem.xml:0`, `/workspace/data/raw/Enemy/Code/Core/ConfigurationWindow.xml:0` |
| Namespaces detected | OnTextChanged |
| Source kinds | bindings, xml_handlers |
| Example locations | Aura: AuraConfigGeneralOffsetX.OnTextChanged, Aura: AuraConfigGeneralOffsetY.OnTextChanged, Aura: AuraConfigTimerOffsetX.OnTextChanged, Aura: AuraConfigTimerOffsetY.OnTextChanged, BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeColumnsEditBox.OnTextChanged, BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeRowsEditBox.OnTextChanged |
| XML usage count | 83 |
| XML attribute usage count | 83 |
| Lua usage count | 83 |
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

Observed as an XML handler hook bound by 8 addons through frame event handlers.

## Expected Lua Binding

```lua
function()
```

## Element Types

- EditBox

## Seen In

- Aura
- BuffHead
- Enemy
- Killer
- Pocket Palette
- Shinies
- TexturedButtons
- TurretRange

## Examples

- Aura: AuraConfigGeneralOffsetX -> AuraConfigGeneralOffsetX.OnTextChanged -> AuraConfig.OnTextureOffsetXChanged
- Aura: AuraConfigGeneralOffsetY -> AuraConfigGeneralOffsetY.OnTextChanged -> AuraConfig.OnTextureOffsetYChanged
- Aura: AuraConfigTimerOffsetX -> AuraConfigTimerOffsetX.OnTextChanged -> AuraConfig.OnTimerOffsetXChanged
- Aura: AuraConfigTimerOffsetY -> AuraConfigTimerOffsetY.OnTextChanged -> AuraConfig.OnTimerOffsetYChanged
- BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeColumnsEditBox -> BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeColumnsEditBox.OnTextChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeColumnsChanged
- BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeRowsEditBox -> BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeRowsEditBox.OnTextChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeRowsChanged

## Related APIs

- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type

## Used With

- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
