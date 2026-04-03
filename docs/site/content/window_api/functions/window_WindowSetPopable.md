# WindowSetPopable

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 30 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, CMap, Crusher, EZCraftX, EZGuard, Effigy |
| Files seen in | LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, config_Main.lua, libs/LibGUI.lua |
| Namespaces detected | WindowSetPopable |
| Source kinds | lua_calls |
| Example locations | Ace: Popable, ActionBarHide: Popable, Amethyst: Popable, CMap: Popable, Crusher: Popable, EZCraftX: Popable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 31 |
| Global usage count | 31 |
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
WindowSetPopable(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: PanelCache[panel].name, self.name |
| arg2 | Observed as a runtime window or control identifier. | Observed values: false, val |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- CMap
- Crusher
- EZCraftX
- EZGuard
- Effigy
- GCDsaver
- Hopper
- InfoScroller
- LibWBToggler
- Map
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- Shinies
- TargetRing
- Tokens
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: Popable -> WindowSetPopable(self.name, val)
- ActionBarHide: Popable -> WindowSetPopable(self.name, val)
- Amethyst: Popable -> WindowSetPopable(self.name, val)
- CMap: Popable -> WindowSetPopable(self.name, val)
- Crusher: Popable -> WindowSetPopable(self.name, val)
- EZCraftX: Popable -> WindowSetPopable(self.name, val)

## Used With

- [WindowGetPopable](window_WindowGetPopable.md) (HIGH 100/100) - Window Function

## Notes

- none
