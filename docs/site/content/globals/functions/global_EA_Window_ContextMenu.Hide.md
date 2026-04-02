# EA_Window_ContextMenu.Hide

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, MapMonster, Miracle Grow Remix, TexturedButtons, TurretRange |
| Files seen in | `/workspace_addons/BuffHead/Setup/SelectFont.lua:62`, `/workspace_addons/MGRemix/context.lua:105`, `/workspace_addons/MapMonster/Source/MapMonster_Pins.lua:947`, `/workspace_addons/TexturedButtons/Setup/SelectFont.lua:64`, `/workspace_addons/TurrentRange/Setup/SetupDisplay.lua:366` |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | BuffHead: BuffHead.Setup.SelectFont.SetFont, MapMonster: HideLastFilterSubMenus, MapMonster: MapMonster.local.HideLastFilterSubMenus, Miracle Grow Remix: MiracleGrow2.ContextItem, TexturedButtons: TexturedButtons.Setup.SelectFont.SetFont, TurretRange: TurretRange.Setup.Distance.SetFont |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
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
EA_Window_ContextMenu.Hide(arg1)
```

## Description

Observed as a global function across 5 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: 1, 2, Filter_SUB_MENU |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- BuffHead
- MapMonster
- Miracle Grow Remix
- TexturedButtons
- TurretRange

## Examples

- BuffHead: BuffHead.Setup.SelectFont.SetFont -> EA_Window_ContextMenu.Hide(menu)
- MapMonster: HideLastFilterSubMenus -> EA_Window_ContextMenu.Hide(Filter_SUB_MENU)
- MapMonster: MapMonster.local.HideLastFilterSubMenus -> EA_Window_ContextMenu.Hide(Filter_SUB_MENU)
- Miracle Grow Remix: MiracleGrow2.ContextItem -> EA_Window_ContextMenu.Hide(1)
- Miracle Grow Remix: MiracleGrow2.ContextItem -> EA_Window_ContextMenu.Hide(2)
- TexturedButtons: TexturedButtons.Setup.SelectFont.SetFont -> EA_Window_ContextMenu.Hide(menu)

## Related APIs

- none

## Used With

- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
