# WindowSetAlpha

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 30 addons

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
| Addons seen in | Ace, Aura, AutoMark, BuffHead, Busted, DAoCBuff, EA_UiDebugTools, Effigy |
| Files seen in | `/workspace/Ace/LibGUI.lua:200`, `/workspace/Aura/Source/Aura.lua:505`, `/workspace/Aura/Source/AuraHelpers.lua:33`, `/workspace/Aura/Source/AuraHelpers.lua:55`, `/workspace/Aura/Source/AuraTooltip.lua:21`, `/workspace/AutoMark/Source/AutoMark.lua:124`, `/workspace/BuffHead/Setup/SelectColor.lua:15`, `/workspace/BuffHead/Setup/SetupLayout.lua:160` |
| Namespaces detected | WindowSetAlpha |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Alpha, Aura: Aura:UpdateTimerWindow, Aura: AuraHelpers.SetCircleImageTexture, Aura: AuraHelpers.SetDynamicImageTexture, Aura: AuraTooltip.OnInitialize, AutoMark: AutoMark.OnUpdate |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 165 |
| Global usage count | 165 |
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
WindowSetAlpha(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "BustedMiniGUIText", "DAoCBuff_SettingsGeneralButtonBackground", "DAoCBuff_SettingsListMngrButtonBackground" |
| arg2 | Observed as a function or method reference. | Observed values: (chatwindow_tabs_handle_input~=false and 1)or 0, (scrollbar_position==c_SCROLLBAR_POSITION_HIDDEN and 0)or scrollbar_alpha, .1 |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Aura
- AutoMark
- BuffHead
- Busted
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- GuardLine
- Killer
- LibWBToggler
- MapMonster
- Miracle Grow Remix
- Pocket Palette
- PotionBar
- RVAPI_ColorDialog
- RandomMount
- RetAlert
- RoR_SoR
- Shinies
- Targets
- TexturedButtons
- TidyChat
- TurretRange
- WSCT
- WarTriage
- WoH-Reticle
- wbLeadHelper

## Examples

- Ace: LIBGUI_ELEMENT:Alpha -> WindowSetAlpha(self.name, alpha)
- Aura: Aura:UpdateTimerWindow -> WindowSetAlpha(windowId, self:Get("timer-alpha"))
- Aura: AuraHelpers.SetCircleImageTexture -> WindowSetAlpha(window, a)
- Aura: AuraHelpers.SetDynamicImageTexture -> WindowSetAlpha(window, a)
- Aura: AuraTooltip.OnInitialize -> WindowSetAlpha(windowId.."BackgroundInner", .9)
- Aura: AuraTooltip.OnInitialize -> WindowSetAlpha(shareWindowId.."BackgroundInner", .9)

## Related APIs

- none

## Used With

- [DynamicImageSetRotation](window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
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
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAlpha](window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetMovable](window_WindowSetMovable.md) (HIGH 100/100) - Window Function
- [WindowSetParent](window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](window_WindowStopAlphaAnimation.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
