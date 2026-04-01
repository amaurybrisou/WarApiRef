# LabelSetWordWrap

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
| Addons seen in | Ace, Effigy, Enemy, GCDsaver, LibWBToggler, Shinies, WarTriage, WoH-Reticle |
| Files seen in | `/workspace/Ace/LibGUI.lua:451`, `/workspace/Effigy/LibGUI.lua:451`, `/workspace/Enemy/Code/UnitFrames/UnitFramePart.lua:210`, `/workspace/GCDsaver/libs/LibGUI.lua:451`, `/workspace/LibWarBoardToggler/libs/LibGUI.lua:451`, `/workspace/Shinies/Libraries/LibGUI.lua:451`, `/workspace/WarTriage/libs/LibGUI.lua:451`, `/workspace/WoH-Reticle/libs/LibGUI.lua:451` |
| Namespaces detected | LabelSetWordWrap |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Label:WordWrap, Effigy: LIBGUI_Label:WordWrap, Enemy: Enemy.UnitFramePart_OnUpdate_ProceedTextWindowInitialization, GCDsaver: LIBGUI_Label:WordWrap, LibWBToggler: LIBGUI_Label:WordWrap, Shinies: LIBGUI_Label:WordWrap |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 15 |
| Global usage count | 15 |
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
LabelSetWordWrap(arg1, arg2)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: self.name, t.windowName |
| arg2 | Observed as a boolean toggle. | Observed values: data.wrap==true, false, true |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Effigy
- Enemy
- GCDsaver
- LibWBToggler
- Shinies
- WarTriage
- WoH-Reticle

## Examples

- Ace: LIBGUI_Label:WordWrap -> LabelSetWordWrap(self.name, true)
- Ace: LIBGUI_Label:WordWrap -> LabelSetWordWrap(self.name, false)
- Effigy: LIBGUI_Label:WordWrap -> LabelSetWordWrap(self.name, true)
- Effigy: LIBGUI_Label:WordWrap -> LabelSetWordWrap(self.name, false)
- Enemy: Enemy.UnitFramePart_OnUpdate_ProceedTextWindowInitialization -> LabelSetWordWrap(t.windowName, data.wrap==true)
- GCDsaver: LIBGUI_Label:WordWrap -> LabelSetWordWrap(self.name, true)

## Related APIs

- none

## Used With

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelGetWordWrap](window_LabelGetWordWrap.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
