# SystemData.Events.L_BUTTON_UP_PROCESSED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, BuffHead, Miracle Grow Remix, RVAPI_ColorDialog, RVMOD_Manager, Shinies, TidyChat, Miracle Grow Remix, RVAPI_ColorDialog, RVMOD_Manager, Shinies, TidyChat |
| Files seen in | `/workspace/BuffHead/Setup/SetupLayout.lua:234`, `/workspace/BuffHead/Setup/SetupLayout.lua:258`, `/workspace/MGRemix/layout.lua:262`, `/workspace/RVAPI_ColorDialog/RVAPI_ColorDialog.lua:83`, `/workspace/RVAPI_ColorDialog/RVAPI_ColorDialog.lua:97`, `/workspace/RVMOD_Manager/RVMOD_Manager.lua:210`, `/workspace/RVMOD_Manager/RVMOD_Manager.lua:234`, `/workspace/Shinies/Modules/UI/Shinies-UI-Browse.lua:290` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | BuffHead.Setup.Layout.OnHidden, BuffHead.Setup.Layout.OnLButtonUpProcessed, BuffHead.Setup.Layout.Show, MiracleGrow2.InitLayout, MiracleGrow2.LayoutEndDrag, RVAPI_ColorDialog.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 22 |
| Global usage count | 22 |
| Local definition count | 0 |
| Documentation references | 4 |
| Initialization flow references | 6 |
| Known engine namespace | yes |
| Default UI presence | yes |
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

Observed SystemData field used by 7 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- BuffHead
- BuffHead, Miracle Grow Remix, RVAPI_ColorDialog, RVMOD_Manager, Shinies, TidyChat
- Miracle Grow Remix
- RVAPI_ColorDialog
- RVMOD_Manager
- Shinies
- TidyChat

## Related APIs

- [WindowUnregisterEventHandler](../../window_api/functions/window_WindowUnregisterEventHandler.md) (HIGH 100/100) - Window Function
- [ComboBoxExternalOpenMenu](../../window_api/functions/window_ComboBoxExternalOpenMenu.md) (HIGH 80/100) - Window Function

## Used With

- [WindowUnregisterEventHandler](../../window_api/functions/window_WindowUnregisterEventHandler.md) (HIGH 100/100) - Window Function
- [ComboBoxExternalOpenMenu](../../window_api/functions/window_ComboBoxExternalOpenMenu.md) (HIGH 80/100) - Window Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: BuffHead.Setup.Layout.OnHidden, BuffHead.Setup.Layout.OnLButtonUpProcessed, BuffHead.Setup.Layout.Show, MiracleGrow2.InitLayout, MiracleGrow2.LayoutEndDrag, RVAPI_ColorDialog.Initialize
