# WindowSetMovable

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
| Addons seen in | Ace, Enemy, LibWBToggler, PartyCast, PotionBar, Shinies, TexturedButtons, WoH-Reticle |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:104`, `/workspace/data/raw/Ace/LibGUI.lua:110`, `/workspace/data/raw/Enemy/Code/Core/Main.lua:41`, `/workspace/data/raw/Enemy/Code/Guard/Guard.lua:323`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:104`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:110`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:104`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:110` |
| Namespaces detected | WindowSetMovable |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:MakeFixed, Ace: LIBGUI_ELEMENT:MakeMovable, Enemy: Enemy.Guard_GuardIndicator_Update, Enemy: Enemy._Initialize, LibWBToggler: LIBGUI_ELEMENT:MakeFixed, LibWBToggler: LIBGUI_ELEMENT:MakeMovable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 15 |
| Global usage count | 15 |
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
WindowSetMovable(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "DeathWindow", "EA_Window_CityCaptureJoinPromptWindow", "EnemyGuardDistanceIndicator" |
| arg2 | Observed as a boolean toggle. | Observed values: false, g.settings.guardDistanceIndicatorMovable, isChecked |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Enemy
- LibWBToggler
- PartyCast
- PotionBar
- Shinies
- TexturedButtons
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:MakeFixed -> WindowSetMovable(self.name, false)
- Ace: LIBGUI_ELEMENT:MakeMovable -> WindowSetMovable(self.name, true)
- Enemy: Enemy.Guard_GuardIndicator_Update -> WindowSetMovable("EnemyGuardDistanceIndicator", g.settings.guardDistanceIndicatorMovable)
- Enemy: Enemy._Initialize -> WindowSetMovable("DeathWindow", true)
- Enemy: Enemy._Initialize -> WindowSetMovable("EA_Window_CityCaptureJoinPromptWindow", true)
- LibWBToggler: LIBGUI_ELEMENT:MakeFixed -> WindowSetMovable(self.name, false)

## Related APIs

- none

## Used With

- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function

## Triggered By

- [SystemData.Events.INTERFACE_RELOADED](../../events/game_events/game_event_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - Game Event

## Affects

- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
