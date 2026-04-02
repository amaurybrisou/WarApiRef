# EA_Window_ContextMenu.AddUserDefinedMenuItem

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Miracle Grow Remix, MoraleCircle |
| Files seen in | `/workspace_addons/MGRemix/context.lua:27`, `/workspace_addons/MoraleCircle/MoraleCircle.lua:391`, `/workspace_addons/MoraleCircle/MoraleCircle.lua:415`, `/workspace_addons/MoraleCircle/MoraleCircle.lua:436`, `/workspace_addons/MoraleCircle/MoraleCircle.lua:458` |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | Miracle Grow Remix: AddContextItem, Miracle Grow Remix: Miracle Grow Remix.local.AddContextItem, MoraleCircle: MoraleCircle.ColorChanger1, MoraleCircle: MoraleCircle.ColorChanger2, MoraleCircle: MoraleCircle.ColorChanger3, MoraleCircle: MoraleCircle.ColorChanger4 |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
EA_Window_ContextMenu.AddUserDefinedMenuItem(arg1, arg2)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "MoraleSliders", "MoraleSlidersEmpty", "MoraleSlidersFill" |
| arg2 | Observed as a numeric value. | Observed values: 3, iMenuNum |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Miracle Grow Remix
- MoraleCircle

## Examples

- Miracle Grow Remix: AddContextItem -> EA_Window_ContextMenu.AddUserDefinedMenuItem(sName, iMenuNum)
- Miracle Grow Remix: Miracle Grow Remix.local.AddContextItem -> EA_Window_ContextMenu.AddUserDefinedMenuItem(sName, iMenuNum)
- MoraleCircle: MoraleCircle.ColorChanger1 -> EA_Window_ContextMenu.AddUserDefinedMenuItem("MoraleSliders", 3)
- MoraleCircle: MoraleCircle.ColorChanger2 -> EA_Window_ContextMenu.AddUserDefinedMenuItem("MoraleSlidersFill", 3)
- MoraleCircle: MoraleCircle.ColorChanger3 -> EA_Window_ContextMenu.AddUserDefinedMenuItem("MoraleSlidersFull", 3)
- MoraleCircle: MoraleCircle.ColorChanger4 -> EA_Window_ContextMenu.AddUserDefinedMenuItem("MoraleSlidersEmpty", 3)

## Related APIs

- none

## Used With

- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
