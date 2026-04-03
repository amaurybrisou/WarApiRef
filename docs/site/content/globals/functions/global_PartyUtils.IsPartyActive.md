# PartyUtils.IsPartyActive

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

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
| Addons seen in | GuardLine, LibGuard, MapPin, PartyCast, Pure, QuickWarReport, RoR_SoR, Swift Assist |
| Files seen in | GuardLine.lua, PartyCast.lua, QWRComms.lua, RoR_SoR.lua, Source/LibGuard.lua, Source/PureGroupPet.lua, SwiftAssist.lua, TokenMachine.lua |
| Namespaces detected | PartyUtils |
| Source kinds | lua_calls |
| Example locations | GuardLine: update, LibGuard: GetIdFromName, LibGuard: UpdateStateMachine, MapPin: CreateMainContextMenu, MapPin: CreateRequestContextMenu, MapPin: LButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 20 |
| Global usage count | 20 |
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
PartyUtils.IsPartyActive()
```

## Description

Observed as a global function across 9 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- GuardLine
- LibGuard
- MapPin
- PartyCast
- Pure
- QuickWarReport
- RoR_SoR
- Swift Assist
- TokenMachine

## Examples

- GuardLine: update -> PartyUtils.IsPartyActive()
- LibGuard: GetIdFromName -> PartyUtils.IsPartyActive()
- LibGuard: UpdateStateMachine -> PartyUtils.IsPartyActive()
- MapPin: CreateMainContextMenu -> PartyUtils.IsPartyActive()
- MapPin: CreateRequestContextMenu -> PartyUtils.IsPartyActive()
- MapPin: LButtonUp -> PartyUtils.IsPartyActive()

## Used With

- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function

## Affects

- [GameData.Player.Pet.name](../../gamedata/fields/gamedata_GameData.Player.Pet.name.md) (HIGH 100/100) - GameData Field
- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
