# SystemData.Events.L_BUTTON_UP_PROCESSED

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

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
| Addons seen in | BuffHead, Miracle Grow Remix, RVAPI_ColorDialog, RVMOD_Manager, Shinies, TidyChat |
| Files seen in | `/workspace/BuffHead/Setup/SetupLayout.lua:234`, `/workspace/MGRemix/layout.lua:262`, `/workspace/RVAPI_ColorDialog/RVAPI_ColorDialog.lua:83`, `/workspace/RVMOD_Manager/RVMOD_Manager.lua:210`, `/workspace/Shinies/Modules/UI/Shinies-UI-Browse.lua:463`, `/workspace/TidyChat/TidyChat.lua:676` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | BuffHead: BuffHead.Setup.Layout.Show, Miracle Grow Remix: MiracleGrow2.InitLayout, RVAPI_ColorDialog: RVAPI_ColorDialog.Initialize, RVMOD_Manager: RVMOD_Manager.Initialize, Shinies: _G.Shinies:NewModule.OnSelChanged_Criteria_MultiSelCombo, TidyChat: TidyChatHooks.SetupHooks |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
| Local definition count | 0 |
| Documentation references | 2 |
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

Observed as a shared SystemData runtime event used by 6 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- BuffHead
- Miracle Grow Remix
- RVAPI_ColorDialog
- RVMOD_Manager
- Shinies
- TidyChat

## Registrars And Handlers

- BuffHead.Setup.Layout.OnLButtonUpProcessed
- MiracleGrow2.LayoutEndDrag
- RVAPI_ColorDialog.OnLButtonUpColor
- RVMOD_Manager.OnLButtonUpProcessed
- RegisterEventHandler
- ShiniesBrowseUI.Criteria_ReopenComboBox
- TidyChat.OnLBU
- WindowRegisterEventHandler
- global
- window

## Examples

- BuffHead: BuffHead.Setup.Layout.Show -> SystemData.Events.L_BUTTON_UP_PROCESSED -> BuffHead.Setup.Layout.OnLButtonUpProcessed
- Miracle Grow Remix: MiracleGrow2.InitLayout -> SystemData.Events.L_BUTTON_UP_PROCESSED -> MiracleGrow2.LayoutEndDrag
- RVAPI_ColorDialog: RVAPI_ColorDialog.Initialize -> SystemData.Events.L_BUTTON_UP_PROCESSED -> RVAPI_ColorDialog.OnLButtonUpColor
- RVMOD_Manager: RVMOD_Manager.Initialize -> SystemData.Events.L_BUTTON_UP_PROCESSED -> RVMOD_Manager.OnLButtonUpProcessed
- Shinies: _G.Shinies:NewModule.OnSelChanged_Criteria_MultiSelCombo -> SystemData.Events.L_BUTTON_UP_PROCESSED -> ShiniesBrowseUI.Criteria_ReopenComboBox
- TidyChat: TidyChatHooks.SetupHooks -> SystemData.Events.L_BUTTON_UP_PROCESSED -> TidyChat.OnLBU

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
