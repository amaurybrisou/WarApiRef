# WindowSetParent

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

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
| Addons seen in | Ace, BagOMatic, CM_ClosetGoblin, LibWBToggler, PartyCast, Shinies, TidyChat, TurretRange |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:132`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.lua:87`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:132`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:132`, `/workspace/data/raw/Shinies/Libraries/LibGUI.lua:132`, `/workspace/data/raw/Shinies/Source/Shinies.lua:417`, `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TurrentRange/Core.lua:335` |
| Namespaces detected | WindowSetParent |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Parent, BagOMatic: BagOMatic.init, CM_ClosetGoblin: ClosetGoblin.Initialize, LibWBToggler: LIBGUI_ELEMENT:Parent, PartyCast: LIBGUI_ELEMENT:Parent, Shinies: CreateUI |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 14 |
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
WindowSetParent(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "BagOMaticButton", "Map"..MapNumber, "TurretPortrait" |
| arg2 | Observed as a text or wstring payload. | Observed values: "CharacterWindow", "EA_Window_Backpack", "Root" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- BagOMatic
- CM_ClosetGoblin
- LibWBToggler
- PartyCast
- Shinies
- TidyChat
- TurretRange
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:Parent -> WindowSetParent(self.name, self.parent)
- BagOMatic: BagOMatic.init -> WindowSetParent("BagOMaticButton", "EA_Window_Backpack")
- CM_ClosetGoblin: ClosetGoblin.Initialize -> WindowSetParent("yClosetGoblinButton", "CharacterWindow")
- LibWBToggler: LIBGUI_ELEMENT:Parent -> WindowSetParent(self.name, self.parent)
- PartyCast: LIBGUI_ELEMENT:Parent -> WindowSetParent(self.name, self.parent)
- Shinies: CreateUI -> WindowSetParent(config.windowId, windowId)

## Related APIs

- none

## Used With

- none

## Triggered By

- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event

## Affects

- [EA_Window_Backpack](../../globals/tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
