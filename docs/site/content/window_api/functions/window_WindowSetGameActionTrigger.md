# WindowSetGameActionTrigger

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

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
| Addons seen in | Ace, Effigy, GCDsaver, LibWBToggler, Shinies, WarTriage, WoH-Reticle |
| Files seen in | `/workspace/Ace/LibGUI.lua:579`, `/workspace/Effigy/LibGUI.lua:578`, `/workspace/Effigy/States/EffigyStateTargets.lua:325`, `/workspace/GCDsaver/libs/LibGUI.lua:578`, `/workspace/LibWarBoardToggler/libs/LibGUI.lua:578`, `/workspace/Shinies/Libraries/LibGUI.lua:578`, `/workspace/WarTriage/libs/LibGUI.lua:578`, `/workspace/WoH-Reticle/libs/LibGUI.lua:578` |
| Namespaces detected | WindowSetGameActionTrigger |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:Trigger, Effigy: Effigy.UpdateTarget, Effigy: LIBGUI_Button:Trigger, GCDsaver: LIBGUI_Button:Trigger, LibWBToggler: LIBGUI_Button:Trigger, Shinies: LIBGUI_Button:Trigger |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
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
WindowSetGameActionTrigger(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "EffigyTLMOAction", self.name |
| arg2 | Observed as a runtime window or control identifier. | Observed values: GetActionIdFromName("ACTION_BAR_120"), action |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Effigy
- GCDsaver
- LibWBToggler
- Shinies
- WarTriage
- WoH-Reticle

## Examples

- Ace: LIBGUI_Button:Trigger -> WindowSetGameActionTrigger(self.name, action)
- Effigy: Effigy.UpdateTarget -> WindowSetGameActionTrigger("EffigyTLMOAction", GetActionIdFromName("ACTION_BAR_120"))
- Effigy: LIBGUI_Button:Trigger -> WindowSetGameActionTrigger(self.name, action)
- GCDsaver: LIBGUI_Button:Trigger -> WindowSetGameActionTrigger(self.name, action)
- LibWBToggler: LIBGUI_Button:Trigger -> WindowSetGameActionTrigger(self.name, action)
- Shinies: LIBGUI_Button:Trigger -> WindowSetGameActionTrigger(self.name, action)

## Related APIs

- none

## Used With

- [GameData.Player.level](../../gamedata/fields/gamedata_GameData.Player.level.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [WindowSetGameActionData](window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [GameData.Player.level](../../gamedata/fields/gamedata_GameData.Player.level.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
