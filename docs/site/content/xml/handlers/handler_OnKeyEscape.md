# OnKeyEscape

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
| Addons seen in | AutoBand, DuffTimer, EA_UiDebugTools, Emojii, Enemy, MapPin, Mass Refine, SocialWindow 2.0 |
| Namespaces detected | OnKeyEscape |
| Source kinds | bindings, xml_handlers |
| Example locations | AutoBand: .OnKeyEscape, DuffTimer: .OnKeyEscape, EA_UiDebugTools: .OnKeyEscape, Emojii: .OnKeyEscape, Enemy: .OnKeyEscape, MapPin: .OnKeyEscape |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 25 |
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

XML handler event observed across 10 addons.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- none

## Seen In

- AutoBand
- DuffTimer
- EA_UiDebugTools
- Emojii
- Enemy
- MapPin
- Mass Refine
- SocialWindow 2.0
- XpStatus+G
- wbLeadHelper

## Examples

- AutoBand: .OnKeyEscape -> AutoBand.HideCopyLink
- DuffTimer: .OnKeyEscape -> DuffTimer.Options.OnEditBoxChanged
- EA_UiDebugTools: .OnKeyEscape -> DevPadWindow.OnKeyEscape
- Emojii: .OnKeyEscape -> Emojii.UI_ChooseIconDialog_Hide
- Enemy: .OnKeyEscape -> Enemy.CombatLogUI_SnapshotWindow_Hide
- Enemy: .OnKeyEscape -> Enemy.CombatLogUI_StatsWindow_Hide

## Related APIs

- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected binding arguments remain uncertain because contract artifacts capture symbol linkage, not full handler signatures.
