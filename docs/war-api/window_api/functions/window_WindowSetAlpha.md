# WindowSetAlpha

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 19 addons

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
| Addons seen in | Ace, Aura, AutoMark, BuffHead, DAoCBuff, Enemy, GuardLine, Killer |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:200`, `/workspace/data/raw/Aura/Source/Aura.lua:505`, `/workspace/data/raw/Aura/Source/AuraHelpers.lua:33`, `/workspace/data/raw/Aura/Source/AuraHelpers.lua:55`, `/workspace/data/raw/Aura/Source/AuraTooltip.lua:21`, `/workspace/data/raw/AutoMark/Source/AutoMark.lua:124`, `/workspace/data/raw/BuffHead/Setup/SelectColor.lua:15`, `/workspace/data/raw/BuffHead/Setup/SetupLayout.lua:160` |
| Namespaces detected | WindowSetAlpha |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Alpha, Aura: Aura:UpdateTimerWindow, Aura: AuraHelpers.SetCircleImageTexture, Aura: AuraHelpers.SetDynamicImageTexture, Aura: AuraTooltip.OnInitialize, AutoMark: AutoMark.OnUpdate |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 135 |
| Global usage count | 135 |
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
| windowName | Observed as a target window name. | Observed values: "DAoCBuff_SettingsGeneralButtonBackground", "DAoCBuff_SettingsListMngrButtonBackground", "DyeWindowSelectedDyeColor" |
| arg2 | Observed as a function or method reference. | Observed values: (chatwindow_tabs_handle_input~=false and 1)or 0, (scrollbar_position==c_SCROLLBAR_POSITION_HIDDEN and 0)or scrollbar_alpha, .9 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Aura
- AutoMark
- BuffHead
- DAoCBuff
- Enemy
- GuardLine
- Killer
- LibWBToggler
- PartyCast
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TurretRange
- WSCT
- WoH-Reticle

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
- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
