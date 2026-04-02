# WindowUnregisterCoreEventHandler

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
| Addons seen in | Ace, Effigy, GCDsaver, LibWBToggler, Shinies, TidyRoll, WoH-Reticle, nRarity |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:294`, `/workspace_addons/Effigy/LibGUI.lua:294`, `/workspace_addons/GCDsaver/libs/LibGUI.lua:294`, `/workspace_addons/LibWarBoardToggler/libs/LibGUI.lua:294`, `/workspace_addons/Shinies/Libraries/LibGUI.lua:294`, `/workspace_addons/TidyRoll/TidyRoll.lua:265`, `/workspace_addons/WoH-Reticle/libs/LibGUI.lua:294`, `/workspace_addons/nRarity/source/Border.lua:70` |
| Namespaces detected | WindowUnregisterCoreEventHandler |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:UnregisterEvent, Effigy: LIBGUI_ELEMENT:UnregisterEvent, GCDsaver: LIBGUI_ELEMENT:UnregisterEvent, LibWBToggler: LIBGUI_ELEMENT:UnregisterEvent, Shinies: LIBGUI_ELEMENT:UnregisterEvent, TidyRoll: TidyRoll.OnLoad |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
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

Observed as a window function across 8 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "EA_Window_LootRoll", self.name |
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
- TidyRoll
- WoH-Reticle
- nRarity

## Examples

- Ace: LIBGUI_ELEMENT:UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- Effigy: LIBGUI_ELEMENT:UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- GCDsaver: LIBGUI_ELEMENT:UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- LibWBToggler: LIBGUI_ELEMENT:UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- Shinies: LIBGUI_ELEMENT:UnregisterEvent -> WindowUnregisterCoreEventHandler(self.name, e)
- TidyRoll: TidyRoll.OnLoad -> WindowUnregisterCoreEventHandler("EA_Window_LootRoll", "OnShown")

## Related APIs

- none

## Used With

- none

## Triggered By

- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
