# WindowSetDimensions

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 32 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, AdvancedPetAssist, AnywhereTrainer, Aura, BankArkel, BuffHead, Crafting Info Tooltip, DAoCBuff |
| Files seen in | `/workspace/Ace/LibGUI.lua:1174`, `/workspace/Ace/LibGUI.lua:1261`, `/workspace/Ace/LibGUI.lua:367`, `/workspace/Ace/LibGUI.lua:419`, `/workspace/Ace/LibGUI.lua:527`, `/workspace/Ace/LibGUI.lua:650`, `/workspace/Ace/LibGUI.lua:704`, `/workspace/Ace/LibGUI.lua:983` |
| Namespaces detected | WindowSetDimensions |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:Resize, Ace: LIBGUI_Image:Resize, Ace: LIBGUI_Label:Resize, Ace: LIBGUI_MultiTextbox:Resize, Ace: LIBGUI_Scrollbar:Resize, Ace: LIBGUI_Statusbar:Resize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 212 |
| Global usage count | 212 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
WindowSetDimensions(windowName, arg2, arg3)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "APAOptions", "APAOptionsContent", "APAOptionsTitleText" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: (mg.vLayout.dimx)*nMult+5, 0, 1 |
| arg3 | Observed as a runtime window or control identifier. | Observed values: (c_TOP_HEIGHT+c_BOTTOM_HEIGHT+c_CHECKBOX_HEIGHT*count), (mg.vLayout.dimy)*nMult+5, 0 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedPetAssist
- AnywhereTrainer
- Aura
- BankArkel
- BuffHead
- Crafting Info Tooltip
- DAoCBuff
- DaemonAssist
- Effigy
- Enemy
- GCDsaver
- GetStats
- Killer
- LibWBToggler
- MapMonster
- Miracle Grow Remix
- Moth
- PotionBar
- RVMOD_Manager
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyQueue
- TidyRoll
- TurretRange
- WSCT
- WarTriage
- WhoHealedMe
- WoH-Reticle
- wbLeadHelper

## Examples

- Ace: LIBGUI_Button:Resize -> WindowSetDimensions(self.name, width, self.height)
- Ace: LIBGUI_Image:Resize -> WindowSetDimensions(self.name, width, height)
- Ace: LIBGUI_Label:Resize -> WindowSetDimensions(self.name, self.width, self.height)
- Ace: LIBGUI_MultiTextbox:Resize -> WindowSetDimensions(self.name, self.width, self.height)
- Ace: LIBGUI_Scrollbar:Resize -> WindowSetDimensions(self.name, self.width, self.height)
- Ace: LIBGUI_Statusbar:Resize -> WindowSetDimensions(self.name, width, self.height)

## Related APIs

- none

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelGetTextDimensions](window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LabelSetWordWrap](window_LabelSetWordWrap.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetParent](window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
