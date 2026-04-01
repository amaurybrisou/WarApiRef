# WindowGetAlpha

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 11 addons

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
| Addons seen in | Ace, DAoCBuff, Effigy, GCDsaver, GuardLine, LibWBToggler, RetAlert, RoR_SoR |
| Files seen in | `/workspace/Ace/LibGUI.lua:200`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:112`, `/workspace/DAoCBuff/Source/DAoCBuffSettings2ndTier.lua:111`, `/workspace/DAoCBuff/Source/DAoCBuffSettings2ndTier.lua:1260`, `/workspace/Effigy/LibGUI.lua:200`, `/workspace/GCDsaver/libs/LibGUI.lua:200`, `/workspace/GuardLine/GuardLine.lua:197`, `/workspace/LibWarBoardToggler/libs/LibGUI.lua:200` |
| Namespaces detected | WindowGetAlpha |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Alpha, DAoCBuff: DAoCBuffSettings.UC, DAoCBuff: FilterSettings.Cleanup, DAoCBuff: ImExport.Cleanup, Effigy: LIBGUI_ELEMENT:Alpha, GCDsaver: LIBGUI_ELEMENT:Alpha |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 19 |
| Global usage count | 19 |
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
WindowGetAlpha(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "DAoCBuff_Settings", "GuardLineSelfWindow", "RoR_SoR_Main_Window" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- DAoCBuff
- Effigy
- GCDsaver
- GuardLine
- LibWBToggler
- RetAlert
- RoR_SoR
- Shinies
- WarTriage
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:Alpha -> WindowGetAlpha(self.name)
- DAoCBuff: DAoCBuffSettings.UC -> WindowGetAlpha("DAoCBuff_Settings")
- DAoCBuff: FilterSettings.Cleanup -> WindowGetAlpha(fwindow)
- DAoCBuff: ImExport.Cleanup -> WindowGetAlpha(iewindow)
- Effigy: LIBGUI_ELEMENT:Alpha -> WindowGetAlpha(self.name)
- GCDsaver: LIBGUI_ELEMENT:Alpha -> WindowGetAlpha(self.name)

## Related APIs

- none

## Used With

- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [WindowGetHandleInput](window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- [SystemData.Events.UPDATE_PROCESSED](../../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event

## Affects

- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
