# PartyUtils.GetWarbandLeader

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

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
| Addons seen in | MapPin, Queue Queuer, QuickWarReport, followTheLeader, rorAutoInviter |
| Files seen in | QWRComms.lua, QueueQueuer.lua, followTheLeader.lua, rorAutoInviter.lua, source/MapPin.lua |
| Namespaces detected | PartyUtils |
| Source kinds | lua_calls |
| Example locations | MapPin: CreateRequestContextMenu, MapPin: LButtonUp, MapPin: QuickRequest, MapPin: RequestCommand, MapPin: test, Queue Queuer: ChatHandler |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
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
PartyUtils.GetWarbandLeader()
```

## Description

Observed as a global function across 5 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- MapPin
- Queue Queuer
- QuickWarReport
- followTheLeader
- rorAutoInviter

## Examples

- MapPin: CreateRequestContextMenu -> PartyUtils.GetWarbandLeader()
- MapPin: LButtonUp -> PartyUtils.GetWarbandLeader()
- MapPin: QuickRequest -> PartyUtils.GetWarbandLeader()
- MapPin: RequestCommand -> PartyUtils.GetWarbandLeader()
- MapPin: test -> PartyUtils.GetWarbandLeader()
- Queue Queuer: ChatHandler -> PartyUtils.GetWarbandLeader()

## Used With

- [PartyUtils.IsPartyActive](global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function

## Affects

- [GameData.ChatData.name](../../gamedata/fields/gamedata_GameData.ChatData.name.md) (HIGH 100/100) - GameData Field
- [GameData.ChatData.text](../../gamedata/fields/gamedata_GameData.ChatData.text.md) (HIGH 100/100) - GameData Field
- [GameData.Player.zone](../../gamedata/fields/gamedata_GameData.Player.zone.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.QUEST_ABANDONED](../../gamedata/fields/gamedata_GameData.Sound.QUEST_ABANDONED.md) (HIGH 100/100) - GameData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
