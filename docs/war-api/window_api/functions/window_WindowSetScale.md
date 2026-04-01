# WindowSetScale

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 23 addons

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
| Addons seen in | Ace, Aura, BuffHead, DAoCBuff, Effigy, Enemy, GCDsaver, GuardLine |
| Files seen in | `/workspace/Ace/LibGUI.lua:236`, `/workspace/Aura/Source/Aura.lua:505`, `/workspace/Aura/Source/Aura.lua:534`, `/workspace/BuffHead/AdvancedContainers.lua:36`, `/workspace/BuffHead/Container.lua:758`, `/workspace/BuffHead/EffectFrame.lua:52`, `/workspace/DAoCBuff/Source/DAoCBuff.lua:480`, `/workspace/DAoCBuff/Source/DAoCBuffFrames.lua:346` |
| Namespaces detected | WindowSetScale |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Scale, Aura: Aura:UpdateTimerWindow, Aura: Aura:UpdateWindow, BuffHead: BuffHead.local.RegisterLayoutEditor, BuffHead: BuffHeadContainer:AnchorContainers, BuffHead: BuffHeadEffectFrame:SetLayout |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 104 |
| Global usage count | 104 |
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
- Effigy
- Enemy
- GCDsaver
- GuardLine
- LibWBToggler
- MapMonster
- MapPin
- PotionBar
- RVMOD_Manager
- RoR_SoR
- Shinies
- TidyQueue
- TimeInQueue
- TurretRange
- WSCT
- WarTriage
- WoH-Reticle
- nRarity
- wbLeadHelper

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

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LabelSetWordWrap](window_LabelSetWordWrap.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterEditCallback](window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [StatusBar](../../xml/element_types/element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetMovable](window_WindowSetMovable.md) (HIGH 100/100) - Window Function
- [WindowSetParent](window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowStartScaleAnimation](window_WindowStartScaleAnimation.md) (HIGH 90/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [StatusBar](../../xml/element_types/element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
