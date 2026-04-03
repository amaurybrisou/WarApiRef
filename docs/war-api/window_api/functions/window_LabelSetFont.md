# LabelSetFont

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 14 addons

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
| Addons seen in | Ace, BuffHead, DAoCBuff, Enemy, Killer, LibWBToggler, PartyCast, PotionBar |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:444`, `/workspace/data/raw/BuffHead/EffectFrame.lua:52`, `/workspace/data/raw/BuffHead/Setup/SetupLayoutProperties.lua:223`, `/workspace/data/raw/BuffHead/Setup/SetupLayoutProperties.lua:538`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffFrames.lua:40`, `/workspace/data/raw/Enemy/Code/Marks/MarkTemplate.lua:85`, `/workspace/data/raw/Enemy/Code/UnitFrames/UnitFramePart.lua:210`, `/workspace/data/raw/Killer/KillerUiCache.lua:71` |
| Namespaces detected | LabelSetFont |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Label:Font, BuffHead: BuffHead.local.LoadFontSettings, BuffHead: BuffHead.local.SelectFont, BuffHead: BuffHeadEffectFrame:SetLayout, BuffHead: LoadFontSettings, BuffHead: SelectFont |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 36 |
| Global usage count | 36 |
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
| arg1 | Observed as a function or method reference. | Observed values: adat.textname, buttonLabel:GetName(), c_TROLL_BINDS_LABEL |
| arg2 | Observed as a function or method reference. | Observed values: "font_clear_large_bold", "font_clear_medium", "font_default_text" |
| arg3 | Observed as a function or method reference. | Observed values: 10, 20, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- BuffHead
- DAoCBuff
- Enemy
- Killer
- LibWBToggler
- PartyCast
- PotionBar
- Shinies
- TexturedButtons
- TidyRoll
- TurretRange
- WSCT
- WoH-Reticle

## Examples

- Ace: LIBGUI_Label:Font -> LabelSetFont(self.name, font, linespacing)
- BuffHead: BuffHead.local.LoadFontSettings -> LabelSetFont(windowName.."FontExampleLabel", settings.Font, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING)
- BuffHead: BuffHead.local.SelectFont -> LabelSetFont(label, font.Font, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING)
- BuffHead: BuffHeadEffectFrame:SetLayout -> LabelSetFont(frameName.."Time", layoutSettings.Duration.Font, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING)
- BuffHead: BuffHeadEffectFrame:SetLayout -> LabelSetFont(frameName.."Stacks", layoutSettings.StackCount.Font, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING)
- BuffHead: BuffHeadEffectFrame:SetLayout -> LabelSetFont(frameName.."Name", layoutSettings.Name.Font, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING)

## Related APIs

- none

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetTextAlign](window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
