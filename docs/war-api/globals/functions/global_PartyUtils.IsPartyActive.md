# PartyUtils.IsPartyActive

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 11 addons

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
| Addons seen in | CMap, EA_OpenPartyWindow, GuardLine, LibGuard, MapPin, PartyCast, Pure, QuickWarReport |
| Files seen in | Cmap_TidyQueueInt.lua, GuardLine.lua, PartyCast.lua, QWRComms.lua, RoR_SoR.lua, Source/LibGuard.lua, Source/PureGroupPet.lua, SwiftAssist.lua |
| Namespaces detected | PartyUtils |
| Source kinds | lua_calls |
| Example locations | CMap: CmapLButton, CMap: TidyQueueLButton, EA_OpenPartyWindow: OnGroupUpdated, EA_OpenPartyWindow: OnOpenPartySettingUpdated, EA_OpenPartyWindow: ToggleOpenPartyInterestFlag, EA_OpenPartyWindow: UpdateManageTabDisable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 26 |
| Global usage count | 26 |
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

Observed as a global function across 11 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- CMap
- EA_OpenPartyWindow
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

- CMap: CmapLButton -> PartyUtils.IsPartyActive()
- CMap: TidyQueueLButton -> PartyUtils.IsPartyActive()
- EA_OpenPartyWindow: OnGroupUpdated -> PartyUtils.IsPartyActive()
- EA_OpenPartyWindow: OnOpenPartySettingUpdated -> PartyUtils.IsPartyActive()
- EA_OpenPartyWindow: ToggleOpenPartyInterestFlag -> PartyUtils.IsPartyActive()
- EA_OpenPartyWindow: UpdateManageTabDisable -> PartyUtils.IsPartyActive()

## Used With

- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function

## Affects

- [GameData.Player.Group.Settings.isPublic](../../gamedata/fields/gamedata_GameData.Player.Group.Settings.isPublic.md) (HIGH 100/100) - GameData Field
- [GameData.Player.isGroupLeader](../../gamedata/fields/gamedata_GameData.Player.isGroupLeader.md) (HIGH 100/100) - GameData Field
- [GameData.Player.isOpenPartyInterested](../../gamedata/fields/gamedata_GameData.Player.isOpenPartyInterested.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
