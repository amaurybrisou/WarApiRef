# LabelSetTextAlign

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 8 addons

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
| Addons seen in | Ace, BuffHead, Enemy, LibWBToggler, PartyCast, Shinies, WSCT, WoH-Reticle |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:473`, `/workspace/data/raw/BuffHead/EffectFrame.lua:52`, `/workspace/data/raw/Enemy/Code/UnitFrames/UnitFramePart.lua:210`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:473`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:473`, `/workspace/data/raw/Shinies/Libraries/LibGUI.lua:473`, `/workspace/data/raw/WoH-Reticle/libs/LibGUI.lua:473`, `/workspace/data/raw/wsct/wsct_animation.lua:107` |
| Namespaces detected | LabelSetTextAlign |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Label:Align, BuffHead: BuffHeadEffectFrame:SetLayout, Enemy: Enemy.UnitFramePart_OnUpdate_ProceedTextWindowInitialization, LibWBToggler: LIBGUI_Label:Align, PartyCast: LIBGUI_Label:Align, Shinies: LIBGUI_Label:Align |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
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
| arg1 | Observed as a function or method reference. | Observed values: adat.textname, frameName.."Name", frameName.."Stacks" |
| arg2 | Observed as a function or method reference. | Observed values: align, arrAlign[adat.align], data.align |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- BuffHead
- Enemy
- LibWBToggler
- PartyCast
- Shinies
- WSCT
- WoH-Reticle

## Examples

- Ace: LIBGUI_Label:Align -> LabelSetTextAlign(self.name, align)
- BuffHead: BuffHeadEffectFrame:SetLayout -> LabelSetTextAlign(frameName.."Time", layoutSettings.Duration.Alignment)
- BuffHead: BuffHeadEffectFrame:SetLayout -> LabelSetTextAlign(frameName.."Stacks", layoutSettings.StackCount.Alignment)
- BuffHead: BuffHeadEffectFrame:SetLayout -> LabelSetTextAlign(frameName.."Name", layoutSettings.Name.Alignment)
- Enemy: Enemy.UnitFramePart_OnUpdate_ProceedTextWindowInitialization -> LabelSetTextAlign(t.windowName, data.align)
- LibWBToggler: LIBGUI_Label:Align -> LabelSetTextAlign(self.name, align)

## Related APIs

- none

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
