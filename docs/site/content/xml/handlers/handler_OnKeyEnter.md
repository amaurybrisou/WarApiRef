# OnKeyEnter

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
| Addons seen in | DuffTimer, EA_UiDebugTools, MapMonster, MapPin, Mass Refine, RandomMount, SNT_PANEL, Twister |
| Namespaces detected | OnKeyEnter |
| Source kinds | bindings, xml_handlers |
| Example locations | DuffTimer: .OnKeyEnter, EA_UiDebugTools: .OnKeyEnter, MapMonster: .OnKeyEnter, MapPin: .OnKeyEnter, Mass Refine: .OnKeyEnter, RandomMount: .OnKeyEnter |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 30 |
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

XML handler event observed across 11 addons.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- none

## Seen In

- DuffTimer
- EA_UiDebugTools
- MapMonster
- MapPin
- Mass Refine
- RandomMount
- SNT_PANEL
- Twister
- WarBoard_FPS
- XpStatus+G
- nLootLink

## Examples

- DuffTimer: .OnKeyEnter -> DuffTimer.Options.OnEditBoxChanged
- EA_UiDebugTools: .OnKeyEnter -> DevPadWindow.SaveFile
- EA_UiDebugTools: .OnKeyEnter -> DevPadWindow.CreateNewFile
- EA_UiDebugTools: .OnKeyEnter -> DevPadWindow.ConfirmRename
- EA_UiDebugTools: .OnKeyEnter -> ObjectInspector.InspectObject
- MapMonster: .OnKeyEnter -> MapMonster.Editor.OnPosEnterKey

## Related APIs

- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Expected binding arguments remain uncertain because contract artifacts capture symbol linkage, not full handler signatures.
