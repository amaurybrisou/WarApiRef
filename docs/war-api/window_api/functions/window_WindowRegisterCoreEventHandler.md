# WindowRegisterCoreEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 11 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 153

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, Effigy, GCDsaver, LibWBToggler, Shinies, TidyChat, TidyQueue, TidyRoll |
| Files seen in | `/workspace/Ace/LibGUI.lua:285`, `/workspace/Effigy/LibGUI.lua:285`, `/workspace/Effigy/States/EffigyStateCastbar.lua:41`, `/workspace/Effigy/States/EffigyStatePlayer.lua:45`, `/workspace/GCDsaver/libs/LibGUI.lua:285`, `/workspace/LibWarBoardToggler/LibWBToggler.lua:32`, `/workspace/LibWarBoardToggler/libs/LibGUI.lua:285`, `/workspace/Shinies/Libraries/LibGUI.lua:285` |
| Namespaces detected | WindowRegisterCoreEventHandler |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:RegisterEvent, Effigy: Effigy.RegisterStateInfoForCastbar, Effigy: Effigy.RegisterStateInfoForPlayer, Effigy: LIBGUI_ELEMENT:RegisterEvent, GCDsaver: LIBGUI_ELEMENT:RegisterEvent, LibWBToggler: LIBGUI_ELEMENT:RegisterEvent |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 21 |
| Global usage count | 21 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
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
WindowRegisterCoreEventHandler(windowName, windowEvent, handlerName)
```

## Description

Observed binding On* window events directly to a named window.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "EA_Window_ScenarioLobby", "EffigySwingTimerWindow", "UFLayerTimerWindow" |
| windowEvent | Observed as an On* window event string. | Observed values: "OnHidden", "OnLButtonDown", "OnLButtonUp" |
| handlerName | Observed as a Lua handler reference. | Observed values: "LIBGUI_ELEMENT.events."..e.."."..self.name, "TidyChat.OnChannelButtonMButtonUp", "TidyChat.OnEntryBoxTextInputLBD" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- Ace
- Effigy
- GCDsaver
- LibWBToggler
- Shinies
- TidyChat
- TidyQueue
- TidyRoll
- WarTriage
- WoH-Reticle
- nRarity

## Examples

- Ace: LIBGUI_ELEMENT:RegisterEvent -> WindowRegisterCoreEventHandler(self.name, e, "LIBGUI_ELEMENT.events."..e.."."..self.name)
- Effigy: Effigy.RegisterStateInfoForCastbar -> WindowRegisterCoreEventHandler("UFLayerTimerWindow", "OnUpdate", Addon.Name..".CastbarUpdate")
- Effigy: Effigy.RegisterStateInfoForPlayer -> WindowRegisterCoreEventHandler("EffigySwingTimerWindow", "OnUpdate", Addon.Name..".SwingTimerUpdate")
- Effigy: LIBGUI_ELEMENT:RegisterEvent -> WindowRegisterCoreEventHandler(self.name, e, "LIBGUI_ELEMENT.events."..e.."."..self.name)
- GCDsaver: LIBGUI_ELEMENT:RegisterEvent -> WindowRegisterCoreEventHandler(self.name, e, "LIBGUI_ELEMENT.events."..e.."."..self.name)
- LibWBToggler: LIBGUI_ELEMENT:RegisterEvent -> WindowRegisterCoreEventHandler(self.name, e, "LIBGUI_ELEMENT.events."..e.."."..self.name)

## Related APIs

- none

## Used With

- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- [SystemData.Events.INTERACT_DONE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ABILITIES_LIST_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ABILITIES_LIST_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ACTIVE_TITLE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ACTIVE_TITLE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_AREA_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_AREA_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_BEGIN_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_RANK_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_RANK_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAST_TIMER_SETBACK](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAST_TIMER_SETBACK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_END_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_END_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EXP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EXP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_GROUP_LEADER_STATUS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_GROUP_LEADER_STATUS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_INFLUENCE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_INFLUENCE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_RANK_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_RANK_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RVR_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RVR_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_START_INTERACT_TIMER](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_START_INTERACT_TIMER.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
