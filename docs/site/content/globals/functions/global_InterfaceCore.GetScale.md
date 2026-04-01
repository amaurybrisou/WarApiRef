# InterfaceCore.GetScale

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 22 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 168

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, BuffHead, DAoCBuff, EA_UiDebugTools, Effigy, Enemy, GCDsaver, GuardLine |
| Files seen in | `/workspace/Ace/LibGUI.lua:236`, `/workspace/BuffHead/AdvancedContainers.lua:275`, `/workspace/BuffHead/Container.lua:758`, `/workspace/BuffHead/EffectFrame.lua:52`, `/workspace/BuffHead/Setup/LayoutControlFrame.lua:8`, `/workspace/BuffHead/Setup/LayoutFrame.lua:105`, `/workspace/BuffHead/Setup/LayoutFrame.lua:3`, `/workspace/BuffHead/Setup/LayoutFrame.lua:71` |
| Namespaces detected | InterfaceCore |
| Source kinds | globals, lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Scale, BuffHead: AutoSize, BuffHead: BuffHead.AdvancedContainers.OnLayoutEditorFinished, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.Show, BuffHead: BuffHead.Setup.Layout.Properties.Show, BuffHead: BuffHead.Setup.LayoutFrame:Create |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 57 |
| Global usage count | 57 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
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
InterfaceCore.GetScale()
```

## Description

Observed as a global function across 22 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- Ace
- BuffHead
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- GuardLine
- LibWBToggler
- MapPin
- Miracle Grow Remix
- PotionBar
- RVAPI_ColorDialog
- RVMOD_Manager
- RoR_SoR
- Shinies
- TexturedButtons
- TidyQueue
- TurretRange
- WarTriage
- WoH-Reticle
- wbLeadHelper

## Examples

- Ace: LIBGUI_ELEMENT:Scale -> InterfaceCore.GetScale()
- BuffHead: AutoSize -> InterfaceCore.GetScale()
- BuffHead: BuffHead.AdvancedContainers.OnLayoutEditorFinished -> InterfaceCore.GetScale()
- BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.Show -> InterfaceCore.GetScale()
- BuffHead: BuffHead.Setup.Layout.Properties.Show -> InterfaceCore.GetScale()
- BuffHead: BuffHead.Setup.LayoutFrame:Create -> InterfaceCore.GetScale()

## Related APIs

- [InterfaceCore.GetResolutionScale](global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [LabelSetTextAlign](../../window_api/functions/window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetWordWrap](../../window_api/functions/window_LabelSetWordWrap.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [WindowGetAnchor](../../window_api/functions/window_WindowGetAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function

## Used With

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetResolutionScale](global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [LabelSetFont](../../window_api/functions/window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](../../window_api/functions/window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LabelSetWordWrap](../../window_api/functions/window_LabelSetWordWrap.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](../../window_api/functions/window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [StatusBar](../../xml/element_types/element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [SystemData.screenResolution.x](../../systemdata/fields/systemdata_SystemData.screenResolution.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.screenResolution.y](../../systemdata/fields/systemdata_SystemData.screenResolution.y.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAnchor](../../window_api/functions/window_WindowGetAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetOffsetFromParent](../../window_api/functions/window_WindowGetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](../../window_api/functions/window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.format](global_wstring.format.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
