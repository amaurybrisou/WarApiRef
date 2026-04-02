# SystemData.MouseOverWindow.name

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 24 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, AggroMeter, BankArkel, BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy, JunkDump |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.lua:205`, `/workspace_addons/AggroMeter/AggroMeter.lua:242`, `/workspace_addons/AggroMeter/AggroMeter.lua:251`, `/workspace_addons/AggroMeter/AggroMeter.lua:378`, `/workspace_addons/BankArkel/BankArkel.lua:153`, `/workspace_addons/BankArkel/BankArkel.lua:160`, `/workspace_addons/BankArkel/BankArkel.lua:438`, `/workspace_addons/BuffHead/Setup/SelectTexture.lua:107` |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | AdvancedRenownTraining.AbilityTooltip, AggroMeter.OnMouseOverStart, AggroMeter.OnTabRBU, AggroMeter.PickedListMenu, AggroMeter.SelectChar, BankArkel.PackIconOver |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 100 |
| Global usage count | 100 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
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

Observed SystemData field used by 24 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- AdvancedRenownTrainer
- AggroMeter
- BankArkel
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- JunkDump
- Killer
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrowLight
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- Swift Assist
- TexturedButtons
- TidyRoll
- TurretRange
- nRarity
- wbLeadHelper

## Related APIs

- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Hide](../../globals/functions/global_EA_Window_ContextMenu.Hide.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap.ShowZone](../../globals/functions/global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
- [LabelGetTextDimensions](../../window_api/functions/window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [PartyUtils.GetWarbandLeader](../../globals/functions/global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Hide](../../globals/functions/global_EA_Window_ContextMenu.Hide.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap](../../globals/tables/table_EA_Window_WorldMap.md) (HIGH 100/100) - Global Table
- [EA_Window_WorldMap.ShowZone](../../globals/functions/global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](../../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [PartyUtils.GetWarbandLeader](../../globals/functions/global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [SystemData.ActiveWindow.name](systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: AdvancedRenownTraining.AbilityTooltip, AggroMeter.OnMouseOverStart, AggroMeter.OnTabRBU, AggroMeter.PickedListMenu, AggroMeter.SelectChar, BankArkel.PackIconOver
