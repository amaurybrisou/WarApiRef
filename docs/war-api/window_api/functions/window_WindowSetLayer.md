# WindowSetLayer

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 17 addons

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
| Addons seen in | Ace, BankArkel, BuffHead, Effigy, Enemy, GCDsaver, LibWBToggler, PotionBar |
| Files seen in | `/workspace/Ace/LibGUI.lua:178`, `/workspace/BankArkel/BankArkel.lua:95`, `/workspace/BuffHead/Container.lua:417`, `/workspace/BuffHead/EffectFrame.lua:36`, `/workspace/BuffHead/EffectFrame.lua:52`, `/workspace/BuffHead/Setup/SelectColor.lua:27`, `/workspace/BuffHead/Setup/SelectTexture.lua:59`, `/workspace/Effigy/Elements/EffigyBar.lua:135` |
| Namespaces detected | WindowSetLayer |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Layer, BankArkel: BankArkel.Init, BuffHead: BuffHead.Setup.SelectColor.Show, BuffHead: BuffHead.Setup.SelectTexture.Show, BuffHead: BuffHead.local.SetLayer, BuffHead: BuffHead.local.SetupCoreFrame |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 47 |
| Global usage count | 47 |
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
WindowSetLayer(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "BankArkelBackpack", "EA_Window_ContextMenu1", "PotionBar" |
| arg2 | Observed as a function or method reference. | Observed values: 0, 2, 3 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- BankArkel
- BuffHead
- Effigy
- Enemy
- GCDsaver
- LibWBToggler
- PotionBar
- RVAPI_ColorDialog
- RetAlert
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- WarTriage
- WoH-Reticle
- nRarity

## Examples

- Ace: LIBGUI_ELEMENT:Layer -> WindowSetLayer(self.name, layer)
- BankArkel: BankArkel.Init -> WindowSetLayer("BankArkelBackpack", 4)
- BuffHead: BuffHead.Setup.SelectColor.Show -> WindowSetLayer(windowName, WindowGetLayer(window)+1)
- BuffHead: BuffHead.Setup.SelectTexture.Show -> WindowSetLayer(windowName, WindowGetLayer(window)+1)
- BuffHead: BuffHead.local.SetLayer -> WindowSetLayer(frame:GetName(), frame.Settings.Layer)
- BuffHead: BuffHead.local.SetupCoreFrame -> WindowSetLayer(windowName, layoutSettings.Layer)

## Related APIs

- none

## Used With

- [BankWindow](../../globals/tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow.OnKeyEnter](../../globals/functions/global_EA_ChatWindow.OnKeyEnter.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](../../globals/functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack](../../globals/tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LabelSetWordWrap](window_LabelSetWordWrap.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [ScrollWindowSetOffset](window_ScrollWindowSetOffset.md) (HIGH 100/100) - Window Function
- [StatusBar](../../xml/element_types/element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetLayer](window_WindowGetLayer.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetParent](window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [BankWindow.Hide](../../globals/functions/global_BankWindow.Hide.md) (HIGH 88/100) - Global Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- [BankWindow](../../globals/tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [EA_ChatWindow.OnKeyEnter](../../globals/functions/global_EA_ChatWindow.OnKeyEnter.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack](../../globals/tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [BankWindow.Hide](../../globals/functions/global_BankWindow.Hide.md) (HIGH 88/100) - Global Function

## Notes

- none
