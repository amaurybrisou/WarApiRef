# OnTextChanged

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 88/100

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, referenced by generated docs or reference files, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, CastSequence, EA_UiDebugTools, Enemy, GroupRange, Killer, LibAddonButton, MapMonster |
| Namespaces detected | OnTextChanged |
| Source kinds | bindings, xml_handlers |
| Example locations | BuffHead: .OnTextChanged, CastSequence: .OnTextChanged, EA_UiDebugTools: .OnTextChanged, Enemy: .OnTextChanged, GroupRange: .OnTextChanged, Killer: .OnTextChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 155 |
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

XML handler event observed across 21 addons.

## Expected Lua Binding

```lua
function()
```

## Element Types

- none

## Seen In

- BuffHead
- CastSequence
- EA_UiDebugTools
- Enemy
- GroupRange
- Killer
- LibAddonButton
- MapMonster
- MapPin
- MegaphonePlusPlus
- Miracle Grow Remix
- NerfedButtons
- Obsidian
- Pocket Palette
- RVAPI_ColorDialog
- Shinies
- TexturedButtons
- TurretRange
- Twister
- XpStatus+G
- wbLeadHelper

## Examples

- BuffHead: .OnTextChanged -> BuffHead.Setup.SelectColor.OnTintChanged
- BuffHead: .OnTextChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnPlacementFixedOffsetXChanged
- BuffHead: .OnTextChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnPlacementFixedOffsetYChanged
- BuffHead: .OnTextChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerPaddingXChanged
- BuffHead: .OnTextChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerPaddingYChanged
- BuffHead: .OnTextChanged -> BuffHead.Setup.AdvancedContainersItem.Properties.OnIndicatorPaddingXChanged

## Related APIs

- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
