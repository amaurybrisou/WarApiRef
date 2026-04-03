# LabelSetTextAlign

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 38 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, BuffHead, CCTV, CMap, Crusher, EZCraftX |
| Files seen in | CCTV.lua, Castbar.lua, Code/UnitFrames/UnitFramePart.lua, EffectFrame.lua, EffectTracker/EffectBar.lua, Entry.lua, LibGUI.lua, LibGui.lua |
| Namespaces detected | LabelSetTextAlign |
| Source kinds | lua_calls |
| Example locations | Ace: Align, ActionBarHide: Align, Amethyst: Align, BuffHead: SetLayout, CCTV: UpdateSettings, CMap: Align |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 55 |
| Global usage count | 55 |
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
LabelSetTextAlign(arg1, arg2)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "CCTV"..WindowName[i].."Window"..WindowName[i].."Name", "CCTV"..WindowName[i].."Window"..WindowName[i].."NameBG", "CCTV"..WindowName[i].."Window"..WindowName[i].."Text" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: "leftcenter", "right", "rightcenter" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- BuffHead
- CCTV
- CMap
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
- Obsidian
- PartyCast
- Pure
- Pure Careerbar
- RVMOD_SquaredDistances
- RealmStatus
- Shinies
- Squared
- TargetRing
- Tokens
- WSCT
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Align -> LabelSetTextAlign(self.name, align)
- ActionBarHide: Align -> LabelSetTextAlign(self.name, align)
- Amethyst: Align -> LabelSetTextAlign(self.name, align)
- BuffHead: SetLayout -> LabelSetTextAlign(frameName.."Time", layoutSettings.Duration.Alignment)
- BuffHead: SetLayout -> LabelSetTextAlign(frameName.."Stacks", layoutSettings.StackCount.Alignment)
- BuffHead: SetLayout -> LabelSetTextAlign(frameName.."Name", layoutSettings.Name.Alignment)

## Used With

- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function

## Notes

- none
