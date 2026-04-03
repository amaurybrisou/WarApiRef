# InterfaceCore.GetScale

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 64 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 143

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, ActionFraction, AdjustTheTip, Amethyst, BuffHead, CMap, CleanUnitFrames, Crusher |
| Files seen in | AdjustTheTip.lua, AdvancedContainers.lua, Amethyst.lua, Button.lua, CleanTargetWindow.lua, Code/UnitFrames/EffectsIndicator.lua, Code/UnitFrames/UnitFrame.lua, Code/UnitFrames/UnitFramePart.lua |
| Namespaces detected | InterfaceCore |
| Source kinds | lua_calls |
| Example locations | Ace: Scale, ActionFraction: ResetWindow, AdjustTheTip: UpdateCallback, Amethyst: Recreate, Amethyst: Scale, BuffHead: AnchorContainers |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 183 |
| Global usage count | 183 |
| Local definition count | 0 |
| Documentation references | 0 |
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

Observed as a global function across 64 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionFraction
- AdjustTheTip
- Amethyst
- BuffHead
- CMap
- CleanUnitFrames
- Crusher
- DAoCBuff
- EA_LoadingScreen
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- Enemy
- GCDsaver
- Group Icons
- GroupRange
- GuardLine
- Hopper
- InfoScroller
- KeyBar
- LibAddonButton
- LibWBToggler
- Map
- MapPin
- Mass Refine
- Miracle Grow Remix
- Motion
- Obsidian
- PartyCast
- PotionBar
- Preciousss
- Pure
- Pure Careerbar
- RO-Style Combat Text
- RVAPI_ColorDialog
- RVMOD_Manager
- RVMOD_SquaredDistances
- RealmStatus
- RoR_SoR
- SNT_PANEL
- Shinies
- Squared
- SquaredHDIndicator
- SquaredHotIndicators
- TargetRing
- TastyButtons
- TexturedButtons
- TidyQueue
- Tokens
- TurretRange
- Vectors
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Scale -> InterfaceCore.GetScale()
- ActionFraction: ResetWindow -> InterfaceCore.GetScale()
- AdjustTheTip: UpdateCallback -> InterfaceCore.GetScale()
- Amethyst: Recreate -> InterfaceCore.GetScale()
- Amethyst: Scale -> InterfaceCore.GetScale()
- BuffHead: AnchorContainers -> InterfaceCore.GetScale()

## Related APIs

- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [InterfaceCore.GetResolutionScale](global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](../../window_api/functions/window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](../../window_api/functions/window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
