# OnEnterPressed

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 76/100

## Confidence Assessment

- Level: HIGH

- Score: 76/100

- Rationale: Promoted as HIGH confidence because referenced by generated docs or reference files, seen in 2 to 3 addons, used in event registration or dispatch.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | FastFriends, MegaphonePlusPlus |
| Namespaces detected | OnEnterPressed |
| Source kinds | bindings, xml_handlers |
| Example locations | FastFriends: .OnEnterPressed, MegaphonePlusPlus: .OnEnterPressed |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
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

XML handler event observed across 2 addons.

## Expected Lua Binding

```lua
function()
```

## Element Types

- none

## Seen In

- FastFriends
- MegaphonePlusPlus

## Examples

- FastFriends: .OnEnterPressed -> FastFriends.UI.OnMinLevelCommit
- MegaphonePlusPlus: .OnEnterPressed -> Megaphone.SaveSettings

## Related APIs

- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
