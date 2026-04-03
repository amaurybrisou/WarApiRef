# OnMouseOverEnd

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
| Addons seen in | Atlas, BuffHead, CM_ClosetGoblin, CastSequence, CleanUnitFrames, DAoCBuff, Enemy, GroupSpotter |
| Namespaces detected | OnMouseOverEnd |
| Source kinds | bindings, xml_handlers |
| Example locations | Atlas: .OnMouseOverEnd, BuffHead: .OnMouseOverEnd, CM_ClosetGoblin: .OnMouseOverEnd, CastSequence: .OnMouseOverEnd, CleanUnitFrames: .OnMouseOverEnd, DAoCBuff: .OnMouseOverEnd |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 126 |
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

XML handler event observed across 38 addons.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- none

## Seen In

- Atlas
- BuffHead
- CM_ClosetGoblin
- CastSequence
- CleanUnitFrames
- DAoCBuff
- Enemy
- GroupSpotter
- HealGrid
- Kwestor
- LibAddonButton
- MapMonster
- MarkBuff
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Motion
- NaturalLog
- NerfedButtons
- NoOverheal
- Obsidian
- Pure
- QuickWarReport
- RVMOD_SquaredDistances
- RandomMount
- ReliquaryHunter
- SOR
- SessionRPs
- Shinies
- Squared
- TalismanGenie
- TexturedButtons
- TomeTracker
- TurretRange
- WARCommander
- WarBoard
- ZonePOP
- wbLeadHelper

## Examples

- Atlas: .OnMouseOverEnd -> AtlasMap.OnMouseOverEnd
- BuffHead: .OnMouseOverEnd -> BuffHead.Setup.SelectTexture.OnTextureRowMouseOut
- BuffHead: .OnMouseOverEnd -> BuffHead.Setup.AdvancedCompression.OnRowMouseOut
- BuffHead: .OnMouseOverEnd -> BuffHead.Setup.AdvancedCompressionItem.OnRowMouseOut
- BuffHead: .OnMouseOverEnd -> BuffHead.Setup.AdvancedContainers.OnRowMouseOut
- BuffHead: .OnMouseOverEnd -> BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOut

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](../element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](../element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [FullResizeImage](../element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [ListBox](../element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.MousePosition.x](../../systemdata/fields/systemdata_SystemData.MousePosition.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.MousePosition.y](../../systemdata/fields/systemdata_SystemData.MousePosition.y.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected binding arguments remain uncertain because contract artifacts capture symbol linkage, not full handler signatures.
