# WindowGetHandleInput

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 28 addons

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
| Addons seen in | Ace, CMap, Crusher, DAoCBuff, EZCraftX, EZGuard, Effigy, EveryBodyGuard |
| Files seen in | LibGUI.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, Settings.lua, Source/DAoCBuffSettings.lua, libs/LibGUI.lua |
| Namespaces detected | WindowGetHandleInput |
| Source kinds | lua_calls |
| Example locations | Ace: TakesInput, CMap: TakesInput, Crusher: TakesInput, DAoCBuff: Disable, DAoCBuff: OpenOptionswindow, DAoCBuff: Reactivate |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 34 |
| Global usage count | 34 |
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
| windowName | Observed as a target window name. | Observed values: "DAoCBuff_Settings", "EveryBodyGuard_Settings", "Vectors_Settings" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- CMap
- Crusher
- DAoCBuff
- EZCraftX
- EZGuard
- Effigy
- EveryBodyGuard
- GCDsaver
- Hopper
- InfoScroller
- LibWBToggler
- Map
- Motion
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- Shinies
- TargetRing
- Tokens
- Vectors
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo

## Examples

- Ace: TakesInput -> WindowGetHandleInput(self.name)
- CMap: TakesInput -> WindowGetHandleInput(self.name)
- Crusher: TakesInput -> WindowGetHandleInput(self.name)
- DAoCBuff: Disable -> WindowGetHandleInput("DAoCBuff_Settings")
- DAoCBuff: OpenOptionswindow -> WindowGetHandleInput("DAoCBuff_Settings")
- DAoCBuff: Reactivate -> WindowGetHandleInput("DAoCBuff_Settings")

## Notes

- none
