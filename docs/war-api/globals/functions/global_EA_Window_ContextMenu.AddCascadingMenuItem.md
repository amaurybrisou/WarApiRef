# EA_Window_ContextMenu.AddCascadingMenuItem

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

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
| Addons seen in | MapPin, MoraleCircle, RoR_SoR |
| Files seen in | `/workspace/MapPin/source/MapPin.lua:2861`, `/workspace/MapPin/source/MapPin.lua:449`, `/workspace/MoraleCircle/MoraleCircle.lua:374`, `/workspace/RoR_SoR/RoR_SoR.lua:2400`, `/workspace/RoR_SoR/RoR_SoR.lua:2425` |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | MapPin: MapPin.CreateMainContextMenu, MapPin: MapPin.OptionsMenu, MoraleCircle: MoraleCircle.ColorChanger, RoR_SoR: RoR_SoR.BroadCastOption, RoR_SoR: RoR_SoR.POPOption |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 18 |
| Global usage count | 18 |
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
EA_Window_ContextMenu.AddCascadingMenuItem(arg1, arg2, arg3, arg4)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: L "<icon00075>Party (1)..", L "<icon00075>Party (2)..", L "<icon00075>Party (3).." |
| arg2 | Observed as a runtime window or control identifier. | Observed values: MakeCallBack2(LineWord1,LineWord2,LineWord3,LineWord4,LineWord5,1), MakeCallBack2(LineWord1,LineWord2,LineWord3,LineWord4,LineWord5,2), MakeCallBack2(LineWord1,LineWord2,LineWord3,LineWord4,LineWord5,3) |
| arg3 | Observed as a boolean toggle. | Observed values: false |
| arg4 | Observed as a numeric value. | Observed values: 1, 2, EA_Window_ContextMenu.CONTEXT_MENU_1 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- MapPin
- MoraleCircle
- RoR_SoR

## Examples

- MapPin: MapPin.CreateMainContextMenu -> EA_Window_ContextMenu.AddCascadingMenuItem(L "<icon00076>Warband All..", Warband_Menu(4), false, 1)
- MapPin: MapPin.CreateMainContextMenu -> EA_Window_ContextMenu.AddCascadingMenuItem(L "<icon00075>Party (1)..", Warband_Menu(7), false, 1)
- MapPin: MapPin.CreateMainContextMenu -> EA_Window_ContextMenu.AddCascadingMenuItem(L "<icon00075>Party (2)..", Warband_Menu(8), false, 1)
- MapPin: MapPin.CreateMainContextMenu -> EA_Window_ContextMenu.AddCascadingMenuItem(L "<icon00075>Party (3)..", Warband_Menu(9), false, 1)
- MapPin: MapPin.CreateMainContextMenu -> EA_Window_ContextMenu.AddCascadingMenuItem(L "<icon00075>Party (4)..", Warband_Menu(10), false, 1)
- MapPin: MapPin.CreateMainContextMenu -> EA_Window_ContextMenu.AddCascadingMenuItem(L "<icon00078>Party Requests..", Request_Menu(1), false, 1)

## Related APIs

- none

## Used With

- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](../../window_api/functions/window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowStartScaleAnimation](../../window_api/functions/window_WindowStartScaleAnimation.md) (HIGH 90/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
