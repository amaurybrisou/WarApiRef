# WindowSetLayer

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 56 addons

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
| Addons seen in | Ace, ActionBarHide, Amethyst, Asshat, BankArkel, BetterCC, BuffHead, CMap |
| Files seen in | Asshat.lua, BankArkel.lua, BetterCC.lua, Castbar.lua, Code/CombatLog/CombatLogStatsWindow.lua, Code/GroupIcons/GroupIcon.lua, Code/Marks/MarkTemplate.lua, Code/ScenarioInfo/ScenarioInfo.lua |
| Namespaces detected | WindowSetLayer |
| Source kinds | lua_calls |
| Example locations | Ace: Layer, ActionBarHide: Layer, Amethyst: Layer, Asshat: Init, BankArkel: Init, BetterCC: UpdatePulse |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 123 |
| Global usage count | 123 |
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
- Asshat
- BankArkel
- BetterCC
- BuffHead
- CMap
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
- MyReasons
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
- WarBoard_WarWhisperer
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
- Asshat: Init -> WindowSetLayer(hostileWindow, 3)
- Asshat: Init -> WindowSetLayer(friendlyWindow, 3)
- BankArkel: Init -> WindowSetLayer("BankArkelBackpack", 4)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [WindowGetLayer](window_WindowGetLayer.md) (HIGH 100/100) - Window Function

## Affects

- [GameData.Account.ServerName](../../gamedata/fields/gamedata_GameData.Account.ServerName.md) (HIGH 100/100) - GameData Field
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field

## Notes

- none
