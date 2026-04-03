# PartyUtils.IsPlayerInWarband

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 108

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | EZGuard, Queue Queuer, Refer, ResHelp |
| Files seen in | EZGuard.lua, QueueQueuer.lua, Refer.lua, ResHelp.lua |
| Namespaces detected | PartyUtils |
| Source kinds | lua_calls |
| Example locations | EZGuard: UpdateGroup, Queue Queuer: UpdateQueuerList, Refer: ModifiedAddGroupMenuItems, ResHelp: AddList |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
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
PartyUtils.IsPlayerInWarband(arg1)
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: EZGuard.Player.name, PlayerMenuWindow.curPlayer.name, player |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- EZGuard
- Queue Queuer
- Refer
- ResHelp

## Examples

- EZGuard: UpdateGroup -> PartyUtils.IsPlayerInWarband(EZGuard.Player.name)
- Queue Queuer: UpdateQueuerList -> PartyUtils.IsPlayerInWarband(player)
- Queue Queuer: UpdateQueuerList -> PartyUtils.IsPlayerInWarband(playerName)
- Refer: ModifiedAddGroupMenuItems -> PartyUtils.IsPlayerInWarband(PlayerMenuWindow.curPlayer.name)
- ResHelp: AddList -> PartyUtils.IsPlayerInWarband(towstring(ResName))

## Affects

- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
