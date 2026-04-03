# LabelSetFont

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 55 addons

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
| Addons seen in | Ace, ActionBarHide, ActionFraction, Amethyst, AutoBand, BuffHead, CDown, Countdown |
| Files seen in | AutoBand.lua, Bars/HealGridProgressBar.lua, CDownFrames.lua, CDownSettings.lua, Castbar.lua, Code/Marks/MarkTemplate.lua, Code/UnitFrames/UnitFramePart.lua, Display.lua |
| Namespaces detected | LabelSetFont |
| Source kinds | lua_calls |
| Example locations | Ace: Font, ActionBarHide: Font, ActionFraction: Initialize, ActionFraction: SetFontLarge, ActionFraction: SetFontMedium, ActionFraction: SetFontSmall |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 157 |
| Global usage count | 157 |
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
LabelSetFont(arg1, arg2, arg3)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "DefaultTooltipRow1Col1Text", "DefaultTooltipRow1Col2Text", "DefaultTooltipRow1Col3Text" |
| arg2 | Observed as a function or method reference. | Observed values: "font_clear_large_bold", "font_clear_medium", "font_clear_medium_bold" |
| arg3 | Observed as a function or method reference. | Observed values: 0, 10, 20 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- Amethyst
- AutoBand
- BuffHead
- CDown
- Countdown
- Crusher
- DAoCBuff
- DetauntHelper
- DuffTimer
- EZCraftX
- EZGuard
- Effigy
- Enemy
- GCDsaver
- HealGrid
- Hopper
- InfoScroller
- Killer
- Kwestor
- LibWBToggler
- Map
- Mech
- Moth
- Motion
- NAMBLA
- NaturalLog
- Obsidian
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- RVAPI_ColorDialog
- RVMOD_SquaredDistances
- RealmStatus
- Shinies
- Squared
- TacticSetNames
- TargetRing
- TastyButtons
- TexturedButtons
- TidyRoll
- Tokens
- TurretRange
- WSCT
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Font -> LabelSetFont(self.name, font, linespacing)
- ActionBarHide: Font -> LabelSetFont(self.name, font, linespacing)
- ActionFraction: Initialize -> LabelSetFont(windowName.."LabelCurrentAP", ActionFraction.Settings.FontSize, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING)
- ActionFraction: Initialize -> LabelSetFont(windowName.."LabelMaximumAP", ActionFraction.Settings.FontSize, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING)
- ActionFraction: SetFontLarge -> LabelSetFont(windowName.."LabelCurrentAP", ActionFractionWindow.FontSizeLarge, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING)
- ActionFraction: SetFontLarge -> LabelSetFont(windowName.."LabelMaximumAP", ActionFractionWindow.FontSizeLarge, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event

## Used With

- [LabelSetTextAlign](window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_AGRO_MODE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_AGRO_MODE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HEALTH_FADE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HEALTH_FADE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field

## Notes

- none
