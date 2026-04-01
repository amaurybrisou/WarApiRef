# WindowUnregisterCoreEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 10 addons

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
| Addons seen in | Ace, Effigy, GCDsaver, LibWBToggler, Shinies, TidyQueue, TidyRoll, WarTriage |
| Files seen in | `/workspace/Ace/LibGUI.lua:294`, `/workspace/Effigy/LibGUI.lua:294`, `/workspace/GCDsaver/libs/LibGUI.lua:294`, `/workspace/LibWarBoardToggler/libs/LibGUI.lua:294`, `/workspace/Shinies/Libraries/LibGUI.lua:294`, `/workspace/TidyQueue/TidyQueue.lua:256`, `/workspace/TidyRoll/TidyRoll.lua:265`, `/workspace/WarTriage/libs/LibGUI.lua:294` |
| Namespaces detected | WindowUnregisterCoreEventHandler |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:UnregisterEvent, Effigy: LIBGUI_ELEMENT:UnregisterEvent, GCDsaver: LIBGUI_ELEMENT:UnregisterEvent, LibWBToggler: LIBGUI_ELEMENT:UnregisterEvent, Shinies: LIBGUI_ELEMENT:UnregisterEvent, TidyQueue: TidyQueue.Initialize |
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
WindowUnregisterCoreEventHandler(arg1, arg2)
```

## Description

Observed as a window function across 10 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "EA_Window_LootRoll", "EA_Window_ScenarioLobby", self.name |
| arg2 | Observed as a runtime window or control identifier. | Observed values: "OnHidden", "OnShown", e |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Removes a previously registered callback or command binding.

## Seen In

- Ace
- Effigy
- GCDsaver
- LibWBToggler
- Shinies
- TidyQueue
- TidyRoll
- WarTriage
- WoH-Reticle
- nRarity

## Examples

- Ace: LIBGUI_ELEMENT:UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- Effigy: LIBGUI_ELEMENT:UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- GCDsaver: LIBGUI_ELEMENT:UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- LibWBToggler: LIBGUI_ELEMENT:UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- Shinies: LIBGUI_ELEMENT:UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- TidyQueue: TidyQueue.Initialize -> WindowUnregisterCoreEventHandler("EA_Window_ScenarioLobby", "OnHidden")

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [SystemData.Events.INTERACT_SHOW_SCENARIO_QUEUE_LIST](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_SHOW_SCENARIO_QUEUE_LIST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_UPDATED_SCENARIO_QUEUE_LIST](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_UPDATED_SCENARIO_QUEUE_LIST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_ACTIVE_QUEUE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_ACTIVE_QUEUE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
