# OnUpdate

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
| Addons seen in | BuffHead, CMap, EA_LoadingScreen, EA_OpenPartyWindow, EA_ScenarioGroupWindow, EZCraftX, Group Icons, HealGrid |
| Namespaces detected | OnUpdate |
| Source kinds | bindings, xml_handlers |
| Example locations | BuffHead: .OnUpdate, CMap: .OnUpdate, EA_LoadingScreen: .OnUpdate, EA_OpenPartyWindow: .OnUpdate, EA_ScenarioGroupWindow: .OnUpdate, EZCraftX: .OnUpdate |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 24 |
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

XML handler event observed across 14 addons.

## Expected Lua Binding

```lua
function(elapsed)
```

## Element Types

- none

## Seen In

- BuffHead
- CMap
- EA_LoadingScreen
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EZCraftX
- Group Icons
- HealGrid
- KeyBar
- LootAlert
- RealmStatus
- SNT_CASTBAR
- TidyQueue
- TidyRoll

## Examples

- BuffHead: .OnUpdate -> BuffHead.Setup.Layout.OnUpdate
- CMap: .OnUpdate -> CMapWindow.UpdateCoordinatesWMap
- EA_LoadingScreen: .OnUpdate -> EA_Window_LoadingScreen.OnUpdate
- EA_OpenPartyWindow: .OnUpdate -> EA_Window_OpenParty.OnUpdateForFlyOut
- EA_OpenPartyWindow: .OnUpdate -> EA_Window_OpenParty.OnUpdate
- EA_OpenPartyWindow: .OnUpdate -> EA_Window_OpenPartyNearby.OnUpdate

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](../../window_api/functions/window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](../../window_api/functions/window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow.Print](../../globals/functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetItemsFromBackpack](../../globals/functions/global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_ScenarioLobby.OnJoinInstanceWait](../../globals/functions/global_EA_Window_ScenarioLobby.OnJoinInstanceWait.md) (HIGH 100/100) - Global Function
- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [ListBox](../element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [Size](../element_types/element_Size.md) (HIGH 100/100) - XML Element Type
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAlpha](../../window_api/functions/window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetGameActionData](../../window_api/functions/window_WindowGetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](../../window_api/functions/window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [wstring.len](../../globals/functions/global_wstring.len.md) (HIGH 100/100) - Global Function
- [wstring.upper](../../globals/functions/global_wstring.upper.md) (HIGH 100/100) - Global Function
- [BroadcastEvent](../../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [EA_Window_OpenParty.SelectTab](../../globals/functions/global_EA_Window_OpenParty.SelectTab.md) (HIGH 80/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function
- [SettingsChanged](../../events/game_events/game_event_SettingsChanged.md) (MEDIUM 63/100) - Game Event
- [Start](../../events/game_events/game_event_Start.md) (MEDIUM 43/100) - Game Event
- [Stop](../../events/game_events/game_event_Stop.md) (MEDIUM 43/100) - Game Event

## Used With

- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function

## Affects

- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [GameData.Player.zone](../../gamedata/fields/gamedata_GameData.Player.zone.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.DO_ABILITY](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_ABILITY.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.NONE](../../gamedata/fields/gamedata_GameData.PlayerActions.NONE.md) (HIGH 100/100) - GameData Field
- [GameData.ScenarioData.timeLeft](../../gamedata/fields/gamedata_GameData.ScenarioData.timeLeft.md) (HIGH 100/100) - GameData Field
- [SystemData.Events.EXIT_GAME](../../systemdata/fields/systemdata_SystemData.Events.EXIT_GAME.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_SELECT_SCENARIO_QUEUE_LIST](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_SELECT_SCENARIO_QUEUE_LIST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_PET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_PET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SEND_CHAT_TEXT](../../systemdata/fields/systemdata_SystemData.Events.SEND_CHAT_TEXT.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MousePosition.x](../../systemdata/fields/systemdata_SystemData.MousePosition.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.MousePosition.y](../../systemdata/fields/systemdata_SystemData.MousePosition.y.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
