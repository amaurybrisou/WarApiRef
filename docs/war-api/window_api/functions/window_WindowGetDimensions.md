# WindowGetDimensions

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
| Addons seen in | Ace, AdvancedPetAssist, AnywhereTrainer, Aura, BuffHead, Crafting Info Tooltip, DAoCBuff, EA_UiDebugTools |
| Files seen in | `/workspace/Ace/LibGUI.lua:1046`, `/workspace/Ace/LibGUI.lua:1149`, `/workspace/Ace/LibGUI.lua:1237`, `/workspace/Ace/LibGUI.lua:506`, `/workspace/Ace/LibGUI.lua:620`, `/workspace/Ace/LibGUI.lua:680`, `/workspace/Ace/LibGUI.lua:735`, `/workspace/Ace/LibGUI.lua:805` |
| Namespaces detected | WindowGetDimensions |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:New, Ace: LIBGUI_Checkbox:New, Ace: LIBGUI_Combobox:New, Ace: LIBGUI_Image:New, Ace: LIBGUI_MultiTextbox:New, Ace: LIBGUI_Optionbutton:New |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 158 |
| Global usage count | 158 |
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
WindowGetDimensions(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "APAPetTargetHUD", "CharacterWindow", "DAoCBuffCondenseTooltipText" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- AdvancedPetAssist
- AnywhereTrainer
- Aura
- BuffHead
- Crafting Info Tooltip
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- Killer
- LibWBToggler
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrowLight
- Moth
- Pocket Palette
- PotionBar
- RVAPI_ColorDialog
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TurretRange
- WarTriage
- WhoHealedMe
- WoH-Reticle
- nRarity

## Examples

- Ace: LIBGUI_Button:New -> WindowGetDimensions(w.name)
- Ace: LIBGUI_Checkbox:New -> WindowGetDimensions(w.name)
- Ace: LIBGUI_Combobox:New -> WindowGetDimensions(w.name)
- Ace: LIBGUI_Image:New -> WindowGetDimensions(w.name)
- Ace: LIBGUI_MultiTextbox:New -> WindowGetDimensions(w.name)
- Ace: LIBGUI_Optionbutton:New -> WindowGetDimensions(w.name)

## Related APIs

- none

## Used With

- [ButtonSetCheckButtonFlag](window_ButtonSetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [StatusBarSetMaximumValue](window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAnchorCount](window_WindowGetAnchorCount.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowSetResizing](window_WindowSetResizing.md) (HIGH 90/100) - Window Function
- [WindowStartScaleAnimation](window_WindowStartScaleAnimation.md) (HIGH 90/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
