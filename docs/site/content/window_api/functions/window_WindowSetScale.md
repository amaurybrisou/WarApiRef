# WindowSetScale

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
| Addons seen in | Ace, Aura, BuffHead, DAoCBuff, Enemy, GuardLine, LibWBToggler, PartyCast |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:236`, `/workspace/data/raw/Aura/Source/Aura.lua:505`, `/workspace/data/raw/Aura/Source/Aura.lua:534`, `/workspace/data/raw/BuffHead/AdvancedContainers.lua:36`, `/workspace/data/raw/BuffHead/Container.lua:758`, `/workspace/data/raw/BuffHead/EffectFrame.lua:52`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:480`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffFrames.lua:346` |
| Namespaces detected | WindowSetScale |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Scale, Aura: Aura:UpdateTimerWindow, Aura: Aura:UpdateWindow, BuffHead: BuffHead.local.RegisterLayoutEditor, BuffHead: BuffHeadContainer:AnchorContainers, BuffHead: BuffHeadEffectFrame:SetLayout |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 90 |
| Global usage count | 90 |
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
WindowSetScale(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "DeathWindow", "EA_Window_CityCaptureJoinPromptWindow", "EnemyGuardDistanceIndicator" |
| arg2 | Observed as a function or method reference. | Observed values: .8*adat.uiscale*(critsize/WSCT_TEXTSIZE_BASE), .8*adat.uiscale*WSCT_CRIT_SIZE_PERCENT*(adat.textsize/WSCT_TEXTSIZE_BASE), 0.000001 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Aura
- BuffHead
- DAoCBuff
- Enemy
- GuardLine
- LibWBToggler
- PartyCast
- PotionBar
- RoR_SoR
- Shinies
- TurretRange
- WSCT
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:Scale -> WindowSetScale(self.name, scale)
- Aura: Aura:UpdateTimerWindow -> WindowSetScale(windowId, self:Get("timer-scale"))
- Aura: Aura:UpdateWindow -> WindowSetScale(windowId, self:Get("texture-scale"))
- BuffHead: BuffHead.local.RegisterLayoutEditor -> WindowSetScale(layoutWindow, WindowGetScale(container:GetName()))
- BuffHead: BuffHeadContainer:AnchorContainers -> WindowSetScale(self:GetName(), self.Settings.Scale*InterfaceCore.GetScale())
- BuffHead: BuffHeadEffectFrame:SetLayout -> WindowSetScale(frameName.."Icon", scale*layoutSettings.Icon.Scale)

## Related APIs

- none

## Used With

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
