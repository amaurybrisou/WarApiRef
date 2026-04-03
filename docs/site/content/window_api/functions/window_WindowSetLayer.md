# WindowSetLayer

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 52 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, BankArkel, BetterCC, BuffHead, CastSequence, CraftingWillard |
| Files seen in | BankArkel.lua, BetterCC.lua, Castbar.lua, Code/CombatLog/CombatLogStatsWindow.lua, Code/GroupIcons/GroupIcon.lua, Code/Marks/MarkTemplate.lua, Code/ScenarioInfo/ScenarioInfo.lua, Code/UnitFrames/EffectsIndicator.lua |
| Namespaces detected | WindowSetLayer |
| Source kinds | lua_calls |
| Example locations | Ace: Layer, ActionBarHide: Layer, Amethyst: Layer, BankArkel: Init, BetterCC: UpdatePulse, BuffHead: SetLayer |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 117 |
| Global usage count | 117 |
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
| windowName | Observed as a target window name. | Observed values: "BankArkelBackpack", "EA_Window_ContextMenu1", "Map" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 0, 1, 2 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- Amethyst
- BankArkel
- BetterCC
- BuffHead
- CastSequence
- CraftingWillard
- Crusher
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- Enemy
- GCDsaver
- GroupRange
- HealGrid
- Hopper
- InfoScroller
- LibWBToggler
- Map
- Motion
- NaturalLog
- Obsidian
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- RVAPI_ColorDialog
- RVMOD_SquaredDistances
- RealmStatus
- ResHelp
- RetAlert
- Shinies
- SquaredHotIndicators
- TargetRing
- TastyButtons
- TexturedButtons
- TidyChat
- TidyRoll
- Tokens
- Twister
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- compass
- nRarity
- scenarioInfo
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Layer -> WindowSetLayer(self.name, layer)
- ActionBarHide: Layer -> WindowSetLayer(self.name, layer)
- Amethyst: Layer -> WindowSetLayer(self.name, layer)
- BankArkel: Init -> WindowSetLayer("BankArkelBackpack", 4)
- BetterCC: UpdatePulse -> WindowSetLayer(windowName, 0)
- BetterCC: UpdatePulse -> WindowSetLayer(windowName, 4)

## Related APIs

- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetLayer](window_WindowGetLayer.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field

## Notes

- none
