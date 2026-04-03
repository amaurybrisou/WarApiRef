# LabelSetWordWrap

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 32 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, CDown, Crusher, EZCraftX, EZGuard, Effigy |
| Files seen in | CDownFrames.lua, Code/UnitFrames/UnitFramePart.lua, LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, Phantom.lua, libs/LibGUI.lua |
| Namespaces detected | LabelSetWordWrap |
| Source kinds | lua_calls |
| Example locations | Ace: WordWrap, ActionBarHide: WordWrap, Amethyst: WordWrap, CDown: Create, Crusher: WordWrap, EZCraftX: WordWrap |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 61 |
| Global usage count | 61 |
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
LabelSetWordWrap(arg1, arg2)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "PhantomSettingsInstructions2Label", "PhantomSettingsInstructionsLabel", modName.."Label" |
| arg2 | Observed as a boolean toggle. | Observed values: data.wrap==true, false, true |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- CDown
- Crusher
- EZCraftX
- EZGuard
- Effigy
- Enemy
- GCDsaver
- Hopper
- InfoScroller
- LibWBToggler
- Map
- Motion
- NaturalLog
- PartyCast
- Phantom
- Pure
- Pure Careerbar
- RealmStatus
- Shinies
- TargetRing
- Tokens
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- warboard_TogglerWarBuilder
- xHUD
- xPanels

## Examples

- Ace: WordWrap -> LabelSetWordWrap(self.name, true)
- Ace: WordWrap -> LabelSetWordWrap(self.name, false)
- ActionBarHide: WordWrap -> LabelSetWordWrap(self.name, true)
- ActionBarHide: WordWrap -> LabelSetWordWrap(self.name, false)
- Amethyst: WordWrap -> LabelSetWordWrap(self.name, true)
- Amethyst: WordWrap -> LabelSetWordWrap(self.name, false)

## Used With

- [LabelGetWordWrap](window_LabelGetWordWrap.md) (HIGH 100/100) - Window Function

## Notes

- none
