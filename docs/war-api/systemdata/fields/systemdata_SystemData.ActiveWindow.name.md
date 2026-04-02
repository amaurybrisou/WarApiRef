# SystemData.ActiveWindow.name

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 29 addons

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
| Addons seen in | AdvancedRenownTrainer, AggroMeter, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, EA_UiDebugTools, Effigy |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.lua:426`, `/workspace_addons/Aura/Source/AuraShares.lua:257`, `/workspace_addons/Aura/Source/AuraShares.lua:297`, `/workspace_addons/Aura/Source/AuraShares.lua:319`, `/workspace_addons/Aura/Source/AuraTexture.lua:147`, `/workspace_addons/Aura/Source/AuraTexture.lua:195`, `/workspace_addons/Aura/Source/AuraTexture.lua:65`, `/workspace_addons/BuffHead/EffectFrame.lua:381` |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | AdvancedRenownTraining.AbilityTooltip, AdvancedRenownTraining.OnLButtonUpTab, AggroMeter.OnTabLBU, AuraShares.ChangeSorting, AuraShares.DisplayTooltip, AuraShares.GetSlotRowNumForActiveListRow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 151 |
| Global usage count | 151 |
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

Observed SystemData field used by 29 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- AdvancedRenownTrainer
- AggroMeter
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- JunkDump
- MapPin
- Miracle Grow Remix
- MiracleGrow
- NPC Item Sale Price
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- QuickWarReport
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- followTheLeader
- nRarity
- wbLeadHelper

## Related APIs

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatTabManager.GetTabName](../../globals/functions/global_EA_ChatTabManager.GetTabName.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetSlotFromActionButtonGroup](../../globals/functions/global_EA_Window_Backpack.GetSlotFromActionButtonGroup.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Hide](../../globals/functions/global_EA_Window_ContextMenu.Hide.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionStore.OnItemLButtonDown](../../globals/functions/global_EA_Window_InteractionStore.OnItemLButtonDown.md) (HIGH 100/100) - Global Function
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](../../window_api/functions/window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [BroadcastEvent](../../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EA_ChatTabManager.GetTabName](../../globals/functions/global_EA_ChatTabManager.GetTabName.md) (HIGH 100/100) - Global Function
- [EA_ChatWindowGroups](../../globals/tables/table_EA_ChatWindowGroups.md) (HIGH 100/100) - Global Table
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Hide](../../globals/functions/global_EA_Window_ContextMenu.Hide.md) (HIGH 100/100) - Global Function
- [EditBox](../../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [GameData.PlayerActions.DO_MACRO](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_MACRO.md) (HIGH 100/100) - GameData Field
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [OnTextChanged](../../xml/handlers/handler_OnTextChanged.md) (HIGH 100/100) - XML Handler
- [OnTextChanged](../../events/window_events/window_event_OnTextChanged.md) (HIGH 100/100) - Window Event
- [SystemData.Events.TARGET_PET](systemdata_SystemData.Events.TARGET_PET.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](../../window_api/functions/window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.gsub](../../globals/functions/global_wstring.gsub.md) (HIGH 100/100) - Global Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function
- [BroadcastEvent](../../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: AdvancedRenownTraining.AbilityTooltip, AdvancedRenownTraining.OnLButtonUpTab, AggroMeter.OnTabLBU, AuraShares.ChangeSorting, AuraShares.DisplayTooltip, AuraShares.GetSlotRowNumForActiveListRow
