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
| Addons seen in | Aura, BuffHead, EA_UiDebugTools, Enemy, Killer, MapMonster, MapPin, Miracle Grow Remix |
| Files seen in | `/workspace_addons/Aura/Source/AuraConfig.xml:1569`, `/workspace_addons/Aura/Source/AuraConfig.xml:1591`, `/workspace_addons/Aura/Source/AuraConfig.xml:177`, `/workspace_addons/Aura/Source/AuraConfig.xml:199`, `/workspace_addons/BuffHead/Setup/General.xml:156`, `/workspace_addons/BuffHead/Setup/General.xml:196`, `/workspace_addons/BuffHead/Setup/General.xml:235`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:128` |
| Namespaces detected | OnTextChanged |
| Source kinds | bindings, xml_handlers |
| Example locations | Aura: AuraConfigGeneralOffsetX.OnTextChanged, Aura: AuraConfigGeneralOffsetY.OnTextChanged, Aura: AuraConfigTimerOffsetX.OnTextChanged, Aura: AuraConfigTimerOffsetY.OnTextChanged, BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeColumnsEditBox.OnTextChanged, BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeRowsEditBox.OnTextChanged |
| XML usage count | 112 |
| XML attribute usage count | 112 |
| Lua usage count | 112 |
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

Observed as an XML handler hook bound by 14 addons through frame event handlers.

## Expected Lua Binding

```lua
function()
```

## Element Types

- EditBox

## Seen In

- Aura
- BuffHead
- EA_UiDebugTools
- Enemy
- Killer
- MapMonster
- MapPin
- Miracle Grow Remix
- Pocket Palette
- RVAPI_ColorDialog
- Shinies
- TexturedButtons
- TurretRange
- wbLeadHelper

## Examples

- Aura: AuraConfigGeneralOffsetX -> AuraConfigGeneralOffsetX.OnTextChanged -> AuraConfig.OnTextureOffsetXChanged
- Aura: AuraConfigGeneralOffsetY -> AuraConfigGeneralOffsetY.OnTextChanged -> AuraConfig.OnTextureOffsetYChanged
- Aura: AuraConfigTimerOffsetX -> AuraConfigTimerOffsetX.OnTextChanged -> AuraConfig.OnTimerOffsetXChanged
- Aura: AuraConfigTimerOffsetY -> AuraConfigTimerOffsetY.OnTextChanged -> AuraConfig.OnTimerOffsetYChanged
- BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeColumnsEditBox -> BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeColumnsEditBox.OnTextChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeColumnsChanged
- BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeRowsEditBox -> BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeRowsEditBox.OnTextChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeRowsChanged

## Related APIs

- [TextEditBoxSetTextColor](../../window_api/functions/window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function

## Used With

- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [OnTextChanged](../../events/window_events/window_event_OnTextChanged.md) (HIGH 100/100) - Window Event

## Triggered By

- none

## Affects

- none

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
