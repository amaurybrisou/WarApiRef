# ComboBoxGetDisabledFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 31 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, Crusher, DeepSleep, EZCraftX, EZGuard, Effigy |
| Files seen in | LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, Settings.lua, libs/LibGUI.lua |
| Namespaces detected | ComboBoxGetDisabledFlag |
| Source kinds | lua_calls |
| Example locations | Ace: Enabled, ActionBarHide: Enabled, Amethyst: Enabled, Crusher: Enabled, DeepSleep: ComboBoxChanged, EZCraftX: Enabled |
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
ComboBoxGetDisabledFlag(arg1)
```

## Description

Observed as a window function across 31 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: box, self.name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- Crusher
- DeepSleep
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
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- Shinies
- TargetRing
- Tokens
- Vectors
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: Enabled -> ComboBoxGetDisabledFlag(self.name)
- ActionBarHide: Enabled -> ComboBoxGetDisabledFlag(self.name)
- Amethyst: Enabled -> ComboBoxGetDisabledFlag(self.name)
- Crusher: Enabled -> ComboBoxGetDisabledFlag(self.name)
- DeepSleep: ComboBoxChanged -> ComboBoxGetDisabledFlag(box)
- EZCraftX: Enabled -> ComboBoxGetDisabledFlag(self.name)

## Used With

- [ButtonGetDisabledFlag](window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function

## Notes

- none
