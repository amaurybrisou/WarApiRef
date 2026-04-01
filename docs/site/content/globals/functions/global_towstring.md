# towstring

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 68 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, AutoBand, AutoChannel, BagOMatic |
| Files seen in | `/workspace/Ace/AceLocale-3.0.lua:57`, `/workspace/Ace/LibGUI.lua:1068`, `/workspace/Ace/LibGUI.lua:1124`, `/workspace/Ace/LibGUI.lua:426`, `/workspace/Ace/LibGUI.lua:533`, `/workspace/Ace/LibGUI.lua:656`, `/workspace/Ace/LibGUI.lua:711`, `/workspace/AdvancedPetAssist/APACommands.lua:207` |
| Namespaces detected | towstring |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:SetText, Ace: LIBGUI_Combobox:Add, Ace: LIBGUI_Combobox:Select, Ace: LIBGUI_Label:SetText, Ace: LIBGUI_MultiTextbox:SetText, Ace: LIBGUI_Textbox:SetText |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1528 |
| Global usage count | 1528 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
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
towstring(arg1)
```

## Description

Observed as a global function across 68 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: "    Posting COUNT auction(s) for ITEM with a price of PRICE per stack of SIZE.", " Missing roles\n "..stats.." \n", "" |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- AutoBand
- AutoChannel
- BagOMatic
- BankArkel
- BuffHead
- Busted
- CM_ClosetGoblin
- CombatTextNames
- Crafting Info Tooltip
- DAoCBuff
- DaemonAssist
- DeepSleep
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- GetStats
- GoldTracker
- GuardLine
- HealAll
- JunkDump
- Killer
- LibGroup
- LibGuard
- LibSlash
- LibWBToggler
- LoyalPet
- MapMonster
- MapPin
- MegaphonePlusPlus
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- NPC Item Sale Price
- PeaceOut
- Pocket Palette
- PotionBar
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RRCount
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RetAlert
- RoR_SoR
- ScenarioAlert
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyRoll
- TortallDPSCore
- TurretRange
- WSCT
- WarTriage
- WhoHealedMe
- WoH-Reticle
- followTheLeader
- wbLeadHelper

## Examples

- Ace: LIBGUI_Button:SetText -> towstring(text)
- Ace: LIBGUI_Combobox:Add -> towstring(itemText)
- Ace: LIBGUI_Combobox:Select -> towstring(value)
- Ace: LIBGUI_Label:SetText -> towstring(text)
- Ace: LIBGUI_MultiTextbox:SetText -> towstring(text)
- Ace: LIBGUI_Textbox:SetText -> towstring(text)

## Related APIs

- [AlertText](../tables/table_AlertText.md) (HIGH 100/100) - Global Table
- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.UnregisterSlashCmd](global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function

## Used With

- [AlertText](../tables/table_AlertText.md) (HIGH 100/100) - Global Table
- [AlertTextWindow.AddLine](global_AlertTextWindow.AddLine.md) (HIGH 100/100) - Global Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](../../window_api/functions/window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeTextEntryDialog](global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](../../window_api/functions/window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [EA_ChatTabManager.GetTabName](global_EA_ChatTabManager.GetTabName.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.GetCurrentChannel](global_EA_ChatWindow.GetCurrentChannel.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [EA_ChatWindowGroups](../tables/table_EA_ChatWindowGroups.md) (HIGH 100/100) - Global Table
- [EA_Window_ContextMenu.AddCascadingMenuItem](global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap](../tables/table_EA_Window_WorldMap.md) (HIGH 100/100) - Global Table
- [EA_Window_WorldMap.ShowZone](global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
- [EditBox](../../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [GameData.ChatData.name](../../gamedata/fields/gamedata_GameData.ChatData.name.md) (HIGH 100/100) - GameData Field
- [GameData.ChatData.text](../../gamedata/fields/gamedata_GameData.ChatData.text.md) (HIGH 100/100) - GameData Field
- [GameData.Player.zone](../../gamedata/fields/gamedata_GameData.Player.zone.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.QUEST_ABANDONED](../../gamedata/fields/gamedata_GameData.Sound.QUEST_ABANDONED.md) (HIGH 100/100) - GameData Field
- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetResolutionScale](global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [LabelGetTextDimensions](../../window_api/functions/window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetFont](../../window_api/functions/window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](../../window_api/functions/window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [OnTextChanged](../../xml/handlers/handler_OnTextChanged.md) (HIGH 100/100) - XML Handler
- [OnTextChanged](../../events/window_events/window_event_OnTextChanged.md) (HIGH 100/100) - Window Event
- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [ScrollWindowUpdateScrollRect](../../window_api/functions/window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [StatusBarSetCurrentValue](../../window_api/functions/window_StatusBarSetCurrentValue.md) (HIGH 100/100) - Window Function
- [StatusBarSetForegroundTint](../../window_api/functions/window_StatusBarSetForegroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](../../window_api/functions/window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.screenResolution.x](../../systemdata/fields/systemdata_SystemData.screenResolution.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.screenResolution.y](../../systemdata/fields/systemdata_SystemData.screenResolution.y.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](../../window_api/functions/window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetLayer](../../window_api/functions/window_WindowGetLayer.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](../../window_api/functions/window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetMovable](../../window_api/functions/window_WindowSetMovable.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](../../window_api/functions/window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [wstring.gsub](global_wstring.gsub.md) (HIGH 100/100) - Global Function
- [wstring.match](global_wstring.match.md) (HIGH 100/100) - Global Function
- [BroadcastEvent](global_BroadcastEvent.md) (HIGH 93/100) - Global Function
- [EA_Window_CurrentEvents.SetCurrentTier](global_EA_Window_CurrentEvents.SetCurrentTier.md) (HIGH 90/100) - Global Function
- [WindowStartScaleAnimation](../../window_api/functions/window_WindowStartScaleAnimation.md) (HIGH 90/100) - Window Function
- [AlertTextWindow.AddAlert](global_AlertTextWindow.AddAlert.md) (HIGH 88/100) - Global Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [DestroyWindow](global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [wstring.format](global_wstring.format.md) (HIGH 75/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function
- [wstring.find](global_wstring.find.md) (MEDIUM 63/100) - Global Function
- [wstring.reverse](global_wstring.reverse.md) (MEDIUM 45/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
