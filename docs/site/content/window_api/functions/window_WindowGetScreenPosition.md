# WindowGetScreenPosition

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 37 addons

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
| Addons seen in | AdjustTheTip, Amethyst, Atlas, AuctionStats, AutoMark, BuffHead, Calling, CastSequence |
| Files seen in | AdjustTheTip.lua, Amethyst.lua, AuctionStats.lua, Button.lua, CallingTargetMarker.lua, ChattyCathy.lua, Code/Core/ObjectWindows.lua, Container.lua |
| Namespaces detected | WindowGetScreenPosition |
| Source kinds | lua_calls |
| Example locations | AdjustTheTip: UpdateCallback, Amethyst: SavePosition, Atlas: ShowCoordinatesOnMouseOver, AuctionStats: AddExtraWindow, AutoMark: OnUpdate, BuffHead: GetRelativeScreenPosition |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 111 |
| Global usage count | 111 |
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
| windowName | Observed as a target window name. | Observed values: "AbilityTooltip", "AbilityTooltipBackground", "AtlasFrameMapContainerMapDisplay" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdjustTheTip
- Amethyst
- Atlas
- AuctionStats
- AutoMark
- BuffHead
- Calling
- CastSequence
- ChattyCathy
- CraftingWillard
- Effigy
- Enemy
- Group Icons
- GroupRange
- GuardLine
- KeyBar
- LibAddonButton
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrowLight
- Obsidian
- Pocket Palette
- RVAPI_ColorDialog
- ResHelp
- RvRContribution
- SNT_PANEL
- Shinies
- TastyButtons
- TexturedButtons
- Tokens
- TurretRange
- Vectors
- WARCommander
- WARRatingBuster
- WSCT
- scenarioInfo

## Examples

- AdjustTheTip: UpdateCallback -> WindowGetScreenPosition("AbilityTooltipBackground")
- AdjustTheTip: UpdateCallback -> WindowGetScreenPosition("AbilityTooltip")
- Amethyst: SavePosition -> WindowGetScreenPosition(C.name)
- Atlas: ShowCoordinatesOnMouseOver -> WindowGetScreenPosition("AtlasFrameMapContainerMapDisplay")
- AuctionStats: AddExtraWindow -> WindowGetScreenPosition(Tooltips.curTooltipWindow)
- AuctionStats: AddExtraWindow -> WindowGetScreenPosition(Tooltips.curMouseOverWindow)

## Related APIs

- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../../xml/handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Used With

- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [SystemData.MousePosition.x](../../systemdata/fields/systemdata_SystemData.MousePosition.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.MousePosition.y](../../systemdata/fields/systemdata_SystemData.MousePosition.y.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Notes

- Advanced return analysis: No strong return evidence observed
