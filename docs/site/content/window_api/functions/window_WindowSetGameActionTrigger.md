# WindowSetGameActionTrigger

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
| Addons seen in | Ace, Amethyst, BuddyBind, CMap, CastSequence, Crusher, DetauntHelper, EZCraftX |
| Files seen in | BuddyBind.lua, Core.lua, EveryBodyGuard.lua, GroupSpotter.lua, LibGUI.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, Modules/HealGridMouseClick.lua |
| Namespaces detected | WindowSetGameActionTrigger |
| Source kinds | lua_calls |
| Example locations | Ace: Trigger, Amethyst: Trigger, BuddyBind: init, CMap: Trigger, CastSequence: SetTrigger, Crusher: Trigger |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 45 |
| Global usage count | 45 |
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
WindowSetGameActionTrigger(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "BuddyBindWindowExecute", "EffigyTLMOAction", "GroupSpotter" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 0, GetActionIdFromName("ACTION_BAR_"..BuddyBind.Button), GetActionIdFromName("ACTION_BAR_"..Twister.Settings.button) |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Amethyst
- BuddyBind
- CMap
- CastSequence
- Crusher
- DetauntHelper
- EZCraftX
- EZGuard
- Effigy
- EveryBodyGuard
- GCDsaver
- GroupSpotter
- HealGrid
- Hopper
- InfoScroller
- LibWBToggler
- Map
- MarkBuff
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- Shinies
- SquaredClick
- TargetRing
- Tokens
- Twister
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD

## Examples

- Ace: Trigger -> WindowSetGameActionTrigger(self.name, action)
- Amethyst: Trigger -> WindowSetGameActionTrigger(self.name, action)
- BuddyBind: init -> WindowSetGameActionTrigger("BuddyBindWindowExecute", GetActionIdFromName("ACTION_BAR_"..BuddyBind.Button))
- CMap: Trigger -> WindowSetGameActionTrigger(self.name, action)
- CastSequence: SetTrigger -> WindowSetGameActionTrigger(self:GetName().."Bar"..index, actionId)
- Crusher: Trigger -> WindowSetGameActionTrigger(self.name, action)

## Related APIs

- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 88/100) - XML Event

## Affects

- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field

## Notes

- none
