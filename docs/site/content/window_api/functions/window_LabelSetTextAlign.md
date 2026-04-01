# LabelSetTextAlign

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 10 addons

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
| Addons seen in | Ace, BuffHead, Effigy, Enemy, GCDsaver, LibWBToggler, Shinies, WSCT |
| Files seen in | `/workspace/Ace/LibGUI.lua:473`, `/workspace/BuffHead/EffectFrame.lua:52`, `/workspace/Effigy/LibGUI.lua:473`, `/workspace/Enemy/Code/UnitFrames/UnitFramePart.lua:210`, `/workspace/GCDsaver/libs/LibGUI.lua:473`, `/workspace/LibWarBoardToggler/libs/LibGUI.lua:473`, `/workspace/Shinies/Libraries/LibGUI.lua:473`, `/workspace/WarTriage/libs/LibGUI.lua:473` |
| Namespaces detected | LabelSetTextAlign |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Label:Align, BuffHead: BuffHeadEffectFrame:SetLayout, Effigy: LIBGUI_Label:Align, Enemy: Enemy.UnitFramePart_OnUpdate_ProceedTextWindowInitialization, GCDsaver: LIBGUI_Label:Align, LibWBToggler: LIBGUI_Label:Align |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
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
| arg2 | Observed as a runtime window or control identifier. | Observed values: align, arrAlign[adat.align], data.align |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- BuffHead
- Effigy
- Enemy
- GCDsaver
- LibWBToggler
- Shinies
- WSCT
- WarTriage
- WoH-Reticle

## Examples

- Ace: LIBGUI_Label:Align -> LabelSetTextAlign(self.name, align)
- BuffHead: BuffHeadEffectFrame:SetLayout -> LabelSetTextAlign(frameName.."Time", layoutSettings.Duration.Alignment)
- BuffHead: BuffHeadEffectFrame:SetLayout -> LabelSetTextAlign(frameName.."Stacks", layoutSettings.StackCount.Alignment)
- BuffHead: BuffHeadEffectFrame:SetLayout -> LabelSetTextAlign(frameName.."Name", layoutSettings.Name.Alignment)
- Effigy: LIBGUI_Label:Align -> LabelSetTextAlign(self.name, align)
- Enemy: Enemy.UnitFramePart_OnUpdate_ProceedTextWindowInitialization -> LabelSetTextAlign(t.windowName, data.align)

## Related APIs

- none

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LabelSetWordWrap](window_LabelSetWordWrap.md) (HIGH 100/100) - Window Function
- [StatusBar](../../xml/element_types/element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [StatusBar](../../xml/element_types/element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
