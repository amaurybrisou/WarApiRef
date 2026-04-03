# DynamicImageSetTextureOrientation

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 33 addons

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
| Addons seen in | Ace, Amethyst, Aura, CDown, CMap, Crusher, EZCraftX, EZGuard |
| Files seen in | CDownFrames.lua, FlagCap.lua, InfoScroller.lua, LibGUI.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, RoR_SoR.lua, Source/AuraHelpers.lua |
| Namespaces detected | DynamicImageSetTextureOrientation |
| Source kinds | lua_calls |
| Example locations | Ace: TexFlip, Amethyst: TexFlip, Aura: SetDynamicImageTexture, CDown: Create, CMap: TexFlip, Crusher: TexFlip |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 39 |
| Global usage count | 39 |
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
DynamicImageSetTextureOrientation(arg1, arg2)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "SoR_"..Window_Name.."Flames", TimeBarEnd, WindowName.."Image" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: Window_Name=="162", flipped, mirror |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Amethyst
- Aura
- CDown
- CMap
- Crusher
- EZCraftX
- EZGuard
- Effigy
- FlagCap
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
- RoR_SoR
- Shinies
- TargetRing
- Tokens
- VerticalMorale
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD

## Examples

- Ace: TexFlip -> DynamicImageSetTextureOrientation(self.name, flipped)
- Amethyst: TexFlip -> DynamicImageSetTextureOrientation(self.name, flipped)
- Aura: SetDynamicImageTexture -> DynamicImageSetTextureOrientation(window, mirror)
- CDown: Create -> DynamicImageSetTextureOrientation(TimeBarEnd, true)
- CMap: TexFlip -> DynamicImageSetTextureOrientation(self.name, flipped)
- Crusher: TexFlip -> DynamicImageSetTextureOrientation(self.name, flipped)

## Used With

- [DynamicImageSetRotation](window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function

## Notes

- none
