# DynamicImageSetTextureSlice

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 48 addons

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
| Addons seen in | Ace, Amethyst, Aura, CMap, Crusher, EA_LoadingScreen, EA_ThreePartBar, EZCraftX |
| Files seen in | CMap.lua, FlagCap.lua, InfoScroller.lua, LibGUI.lua, LibWBToggler.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, MGRemix.lua |
| Namespaces detected | DynamicImageSetTextureSlice |
| Source kinds | lua_calls |
| Example locations | Ace: TexSlice, Amethyst: TexSlice, Aura: SetDynamicImageTexture, CMap: PopulateFilterCell, CMap: TexSlice, Crusher: TexSlice |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 126 |
| Global usage count | 126 |
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
DynamicImageSetTextureSlice(arg1, arg2)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "MapPinTooltipsIcon", "SoR_"..Window_Name.."BG", "SoR_"..Window_Name.."CITY_RANK" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: "Black-Slot", "BombDestruction", "Dirt" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Amethyst
- Aura
- CMap
- Crusher
- EA_LoadingScreen
- EA_ThreePartBar
- EZCraftX
- EZGuard
- Effigy
- FlagCap
- GCDsaver
- Hopper
- InfoScroller
- Kwestor
- LibWBToggler
- Map
- MapMonster
- MapPin
- MiniMapMonster
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Moth
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- RoR_SoR
- SOR
- Shinies
- TargetRing
- TexturedButtons
- Tokens
- Trakario
- WarBoard_Menu
- WarBoard_TogglerWARCommander
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- minesweep
- scenarioInfo
- xHUD
- zMailMod

## Examples

- Ace: TexSlice -> DynamicImageSetTextureSlice(self.name, slice)
- Amethyst: TexSlice -> DynamicImageSetTextureSlice(self.name, slice)
- Aura: SetDynamicImageTexture -> DynamicImageSetTextureSlice(window, slice)
- CMap: PopulateFilterCell -> DynamicImageSetTextureSlice(iconFrame, filterData.slice)
- CMap: TexSlice -> DynamicImageSetTextureSlice(self.name, slice)
- Crusher: TexSlice -> DynamicImageSetTextureSlice(self.name, slice)

## Related APIs

- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Used With

- [DynamicImageSetRotation](window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function

## Notes

- Advanced return analysis: No strong return evidence observed
