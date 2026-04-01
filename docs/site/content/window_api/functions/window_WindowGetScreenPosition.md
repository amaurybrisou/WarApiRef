# WindowGetScreenPosition

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
| Addons seen in | AutoMark, BuffHead, Effigy, Enemy, GuardLine, MapMonster, MapPin, Miracle Grow Remix |
| Files seen in | `/workspace/AutoMark/Source/AutoMark.lua:124`, `/workspace/BuffHead/Container.lua:1129`, `/workspace/BuffHead/Setup/LayoutControlFrame.lua:8`, `/workspace/BuffHead/Setup/Setup.lua:141`, `/workspace/BuffHead/Setup/Setup.lua:171`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:375`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:520`, `/workspace/BuffHead/Setup/SetupLayoutProperties.lua:332` |
| Namespaces detected | WindowGetScreenPosition |
| Source kinds | lua_calls |
| Example locations | AutoMark: AutoMark.OnUpdate, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.Show, BuffHead: BuffHead.Setup.Layout.Properties.Show, BuffHead: BuffHead.Setup.SetActiveWindow, BuffHead: BuffHead.Setup.Show, BuffHead: BuffHead.local.GetRelativeScreenPosition |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 52 |
| Global usage count | 52 |
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
WindowGetScreenPosition(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "EA_Window_WorldMapZoneViewMapDisplay", "GuardLineOffGuardSelfWindow", "GuardLineOffGuardTargetWindow" |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- AutoMark
- BuffHead
- Effigy
- Enemy
- GuardLine
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrowLight
- Pocket Palette
- RVAPI_ColorDialog
- Shinies
- TexturedButtons
- TurretRange
- WSCT

## Examples

- AutoMark: AutoMark.OnUpdate -> WindowGetScreenPosition(WINDOW_NAME_TEST)
- BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.Show -> WindowGetScreenPosition(BuffHead.Setup.AdvancedContainersItem.WindowName)
- BuffHead: BuffHead.Setup.Layout.Properties.Show -> WindowGetScreenPosition(BuffHead.Setup.Layout.WindowName)
- BuffHead: BuffHead.Setup.SetActiveWindow -> WindowGetScreenPosition(windowName)
- BuffHead: BuffHead.Setup.Show -> WindowGetScreenPosition(windowName)
- BuffHead: BuffHead.local.GetRelativeScreenPosition -> WindowGetScreenPosition(containerWindow)

## Related APIs

- none

## Used With

- [FullResizeImage](../../xml/element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [SystemData.MousePosition.x](../../systemdata/fields/systemdata_SystemData.MousePosition.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.MousePosition.y](../../systemdata/fields/systemdata_SystemData.MousePosition.y.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAnchorCount](window_WindowGetAnchorCount.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [SystemData.screenResolution.x](../../systemdata/fields/systemdata_SystemData.screenResolution.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.screenResolution.y](../../systemdata/fields/systemdata_SystemData.screenResolution.y.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
