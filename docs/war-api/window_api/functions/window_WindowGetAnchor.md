# WindowGetAnchor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 49 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, BuffHead, CMap, Crusher, DammazKron, DetauntHelper |
| Files seen in | Core.lua, Core/DK_Core.lua, DuffTimer.lua, Enhancements.lua, LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua |
| Namespaces detected | WindowGetAnchor |
| Source kinds | lua_calls |
| Example locations | Ace: GetPosition, ActionBarHide: GetPosition, Amethyst: GetPosition, BuffHead: UpdatePosition, CMap: GetPosition, Crusher: GetPosition |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 78 |
| Global usage count | 78 |
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
WindowGetAnchor(windowName, arg2)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "DammazKronHTS", "EA_Window_WorldMapZoneViewCoordinates", "MothHealthBar" |
| arg2 | Observed as a numeric value. | Observed values: 1, index, n |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- BuffHead
- CMap
- Crusher
- DammazKron
- DetauntHelper
- DuffTimer
- EZCraftX
- EZGuard
- Effigy
- FixGit
- GCDsaver
- Hopper
- InfoScroller
- KillTracker
- LibWBToggler
- Map
- Moth
- Motion
- NaturalLog
- PartyCast
- Pocket Palette
- Preciousss
- Pure
- Pure Careerbar
- RVMOD_PlayerStatus
- RealmStatus
- SOR
- Shinies
- SimpleXY
- Squared
- TargetRing
- TastyButtons
- Tokens
- Vectors
- VerticalMorale
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WindowMovers
- WoH-Reticle
- emotes
- fpsbox
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: GetPosition -> WindowGetAnchor(self.name, 1)
- ActionBarHide: GetPosition -> WindowGetAnchor(self.name, 1)
- Amethyst: GetPosition -> WindowGetAnchor(self.name, 1)
- BuffHead: UpdatePosition -> WindowGetAnchor(self:GetName(), 1)
- CMap: GetPosition -> WindowGetAnchor(self.name, 1)
- Crusher: GetPosition -> WindowGetAnchor(self.name, 1)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Notes

- none
