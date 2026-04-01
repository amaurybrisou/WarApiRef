# WindowGetHandleInput

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 8 addons

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
| Addons seen in | Ace, DAoCBuff, Effigy, GCDsaver, LibWBToggler, Shinies, WarTriage, WoH-Reticle |
| Files seen in | `/workspace/Ace/LibGUI.lua:99`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:112`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:161`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:169`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:91`, `/workspace/Effigy/LibGUI.lua:99`, `/workspace/GCDsaver/libs/LibGUI.lua:99`, `/workspace/LibWarBoardToggler/libs/LibGUI.lua:99` |
| Namespaces detected | WindowGetHandleInput |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:TakesInput, DAoCBuff: DAoCBuffSettings.Disable, DAoCBuff: DAoCBuffSettings.OpenOptionswindow, DAoCBuff: DAoCBuffSettings.Reactivate, DAoCBuff: DAoCBuffSettings.UC, Effigy: LIBGUI_ELEMENT:TakesInput |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
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
WindowGetHandleInput(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "DAoCBuff_Settings", self.name |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- Ace
- DAoCBuff
- Effigy
- GCDsaver
- LibWBToggler
- Shinies
- WarTriage
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:TakesInput -> WindowGetHandleInput(self.name)
- DAoCBuff: DAoCBuffSettings.Disable -> WindowGetHandleInput("DAoCBuff_Settings")
- DAoCBuff: DAoCBuffSettings.OpenOptionswindow -> WindowGetHandleInput("DAoCBuff_Settings")
- DAoCBuff: DAoCBuffSettings.Reactivate -> WindowGetHandleInput("DAoCBuff_Settings")
- DAoCBuff: DAoCBuffSettings.UC -> WindowGetHandleInput("DAoCBuff_Settings")
- Effigy: LIBGUI_ELEMENT:TakesInput -> WindowGetHandleInput(self.name)

## Related APIs

- none

## Used With

- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [WindowGetAlpha](window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [SystemData.Events.UPDATE_PROCESSED](../../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
