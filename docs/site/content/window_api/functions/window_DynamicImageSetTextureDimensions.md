# DynamicImageSetTextureDimensions

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 50 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, Amethyst, BBars - Mechanic Only, BuffHead, CMap, CastSequence, Crusher, DetauntHelper |
| Files seen in | BBarsPetHP.lua, BBarsPlayerMechanic.lua, Bar.lua, Castbar.lua, Code/UnitFrames/Parts/CareerIcon.lua, Code/UnitFrames/UnitFramePart.lua, CooldownPulse.lua, DuffTimer.lua |
| Namespaces detected | DynamicImageSetTextureDimensions |
| Source kinds | lua_calls |
| Example locations | Ace: TexDims, Amethyst: TexDims, BBars - Mechanic Only: MechDraw, BBars - Mechanic Only: PetHPDraw, BuffHead: SetTextureButton, CMap: TexDims |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 168 |
| Global usage count | 168 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
DynamicImageSetTextureDimensions(arg1, arg2, arg3)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BBarsPetHPBG", "BBarsPetHPBack", "BBarsPetHPFront" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: (tw or 64), -32, 0 |
| arg3 | Observed as a numeric value. | Observed values: (th or 64), -32, 0 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Amethyst
- BBars - Mechanic Only
- BuffHead
- CMap
- CastSequence
- Crusher
- DetauntHelper
- DuffTimer
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- Enemy
- GCDsaver
- GroupSpotter
- HealGrid
- Hopper
- InfoScroller
- LibWBToggler
- Map
- MarkBuff
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- Obsidian
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- RealmStatus
- SNT_BUTTONS
- SNT_PANEL
- Shinies
- Squared
- TargetInfoRing
- TargetRing
- TexturedButtons
- Tokens
- WSCT
- WarBoard_TogglerRankedLeaderboard
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD
- zMailMod

## Examples

- Ace: TexDims -> DynamicImageSetTextureDimensions(self.name, width, height)
- Amethyst: TexDims -> DynamicImageSetTextureDimensions(self.name, width, height)
- BBars - Mechanic Only: MechDraw -> DynamicImageSetTextureDimensions("BBarsPlayerMechanic3BG", UnitWidth, UnitHeight)
- BBars - Mechanic Only: MechDraw -> DynamicImageSetTextureDimensions("BBarsPlayerMechanic2BG", UnitWidth, UnitHeight)
- BBars - Mechanic Only: MechDraw -> DynamicImageSetTextureDimensions("BBarsPlayerMechanic1BG", UnitWidth, UnitHeight)
- BBars - Mechanic Only: MechDraw -> DynamicImageSetTextureDimensions("BBarsPlayerMechanic4BG", UnitWidth, UnitHeight)

## Related APIs

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [TexDims](../../xml/element_types/element_TexDims.md) (HIGH 100/100) - XML Element Type
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [TexDims](../../xml/element_types/element_TexDims.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Notes

- none
