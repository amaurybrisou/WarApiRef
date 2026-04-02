# WindowSetLayer

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 15 addons

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
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:178`, `/workspace_addons/BankArkel/BankArkel.lua:95`, `/workspace_addons/BuffHead/Container.lua:417`, `/workspace_addons/BuffHead/EffectFrame.lua:36`, `/workspace_addons/BuffHead/EffectFrame.lua:52`, `/workspace_addons/BuffHead/Setup/SelectColor.lua:27`, `/workspace_addons/BuffHead/Setup/SelectTexture.lua:59`, `/workspace_addons/Effigy/Elements/EffigyBar.lua:135` |
| Namespaces detected | WindowSetLayer |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Layer, BankArkel: BankArkel.Init, BuffHead: BuffHead.Setup.SelectColor.Show, BuffHead: BuffHead.Setup.SelectTexture.Show, BuffHead: BuffHead.local.SetLayer, BuffHead: BuffHead.local.SetupCoreFrame |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 44 |
| Global usage count | 44 |
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
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
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

- [WindowGetLayer](window_WindowGetLayer.md) (HIGH 100/100) - Window Function

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
