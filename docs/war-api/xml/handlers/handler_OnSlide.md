# OnSlide

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
| Addons seen in | Atlas, BuffHead, CDown, ChattyCathy, DetauntHelper, DuffTimer, EA_ScenarioGroupWindow, Enemy |
| Namespaces detected | OnSlide |
| Source kinds | bindings, xml_handlers |
| Example locations | Atlas: .OnSlide, BuffHead: .OnSlide, CDown: .OnSlide, ChattyCathy: .OnSlide, DetauntHelper: .OnSlide, DuffTimer: .OnSlide |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 191 |
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

XML handler event observed across 37 addons.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- none

## Seen In

- Atlas
- BuffHead
- CDown
- ChattyCathy
- DetauntHelper
- DuffTimer
- EA_ScenarioGroupWindow
- Enemy
- Group Icons SG
- GroupRange
- GroupSpotter
- HealGrid
- LibGroup
- MapMonster
- MapPin
- MoraleCircle
- Obsidian
- PotionBar
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_Manager
- RVMOD_SquaredDistances
- ReliquaryHunter
- RoR_SoR
- SNT_INFO
- SNT_PANEL
- SOR
- Shinies
- Statdoll Remix
- TacticSetNames
- TastyButtons
- TexturedButtons
- TurretRange
- WSCT
- WarBoard
- WarBoard_FPS
- XpStatus+G

## Examples

- Atlas: .OnSlide -> Atlas.Configuration.OnSlide
- BuffHead: .OnSlide -> BuffHead.Setup.SelectColor.OnSlideTint
- BuffHead: .OnSlide -> BuffHead.Setup.AdvancedContainersItem.Properties.OnSlideScale
- BuffHead: .OnSlide -> BuffHead.Setup.Container.OnSlidePaddingHorizontal
- BuffHead: .OnSlide -> BuffHead.Setup.Container.OnSlidePaddingVertical
- BuffHead: .OnSlide -> BuffHead.Setup.Container.OnSlideColumns

## Related APIs

- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [SliderBar](../element_types/element_SliderBar.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected binding arguments remain uncertain because contract artifacts capture symbol linkage, not full handler signatures.
