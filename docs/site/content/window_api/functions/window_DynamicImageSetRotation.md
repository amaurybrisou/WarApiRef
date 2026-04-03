# DynamicImageSetRotation

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 29 addons

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
| Addons seen in | Ace, Amethyst, Aura, Crusher, EZCraftX, EZGuard, Effigy, GCDsaver |
| Files seen in | KeyBar.lua, LibGUI.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, Source/AuraHelpers.lua, VerticalMorale.lua, libs/LibGUI.lua |
| Namespaces detected | DynamicImageSetRotation |
| Source kinds | lua_calls |
| Example locations | Ace: Rotate, Amethyst: Rotate, Aura: SetDynamicImageTexture, Crusher: Rotate, EZCraftX: Rotate, EZGuard: Rotate |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 40 |
| Global usage count | 40 |
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
DynamicImageSetRotation(arg1, arg2)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "KeyBarWindowSelector", self.name, window |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 270, 90, SelectorRotation |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Amethyst
- Aura
- Crusher
- EZCraftX
- EZGuard
- Effigy
- GCDsaver
- Hopper
- InfoScroller
- KeyBar
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
- VerticalMorale
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD

## Examples

- Ace: Rotate -> DynamicImageSetRotation(self.name, angle)
- Amethyst: Rotate -> DynamicImageSetRotation(self.name, angle)
- Aura: SetDynamicImageTexture -> DynamicImageSetRotation(window, rotation)
- Crusher: Rotate -> DynamicImageSetRotation(self.name, angle)
- EZCraftX: Rotate -> DynamicImageSetRotation(self.name, angle)
- EZGuard: Rotate -> DynamicImageSetRotation(self.name, angle)

## Related APIs

- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function

## Notes

- Advanced return analysis: No strong return evidence observed
