# DynamicImageSetTextureScale

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 39 addons

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
| Addons seen in | Ace, Amethyst, CMap, CastSequence, Crusher, DammazKron, DetauntHelper, EZCraftX |
| Files seen in | CMap.lua, CooldownPulse.lua, Core/ToolTip/DK_Tooltip.lua, Elements/EffigyBar.lua, Gui.lua, InfoScroller.lua, LibGUI.lua, LibWBToggler.lua |
| Namespaces detected | DynamicImageSetTextureScale |
| Source kinds | lua_calls |
| Example locations | Ace: TexScale, Amethyst: TexScale, CMap: PopulateFilterCell, CMap: TexScale, CastSequence: UpdateButton, Crusher: TexScale |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 67 |
| Global usage count | 67 |
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
DynamicImageSetTextureScale(arg1, arg2)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "warwhispererMainInnerTextCareerButtonIconBase", MBuffGui.Elements.Images.SpellIcon, WICON |
| arg2 | Observed as a runtime window or control identifier. | Observed values: .125, 0.5, 0.7 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Amethyst
- CMap
- CastSequence
- Crusher
- DammazKron
- DetauntHelper
- EZCraftX
- EZGuard
- Effigy
- GCDsaver
- Hopper
- InfoScroller
- LibWBToggler
- Map
- MapMonster
- MarkBuff
- MiniMapMonster
- Miracle Grow Remix
- Moth
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- SNT_CASTBAR
- Shinies
- TargetRing
- TexturedButtons
- Tokens
- WarBoard_Menu
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD

## Examples

- Ace: TexScale -> DynamicImageSetTextureScale(self.name, scale)
- Amethyst: TexScale -> DynamicImageSetTextureScale(self.name, scale)
- CMap: PopulateFilterCell -> DynamicImageSetTextureScale(iconFrame, filterData.scale)
- CMap: TexScale -> DynamicImageSetTextureScale(self.name, scale)
- CastSequence: UpdateButton -> DynamicImageSetTextureScale(button.WarningIcon, 1)
- Crusher: TexScale -> DynamicImageSetTextureScale(self.name, scale)

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [LabelGetTextDimensions](window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function

## Notes

- none
