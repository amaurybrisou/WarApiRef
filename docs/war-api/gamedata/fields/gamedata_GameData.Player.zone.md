# GameData.Player.zone

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 20 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CMap, Killer, Map, MapMonster, MapPin, Minmap, Queue Queuer, RealmStatus |
| Files seen in | CMap.lua, Core.lua, CurseProfilerCompiled.lua, KillerRenown.lua, KillerUtils.lua, KillerZoneHistory.lua, Map.lua, QueueQueuer.lua |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | AllowedToShowDetailWindow, Broadcast1, Broadcast2, BroadcastA, BroadcastG, BroadcastP |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 44 |
| Global usage count | 44 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
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

## Description

GameData.GameData.Player.zone field accessed by 20 addons; commonly found in AllowedToShowDetailWindow and Broadcast1, Broadcast2, BroadcastA, BroadcastG, BroadcastP, BroadcastW, CaptureInitialWarCrests, ChatHandler, GetActiveRvrZoneId, GetCurrentZoneHistoryEntry, GetCurrentZoneId, HandleSlashCmd, InitialEpicEntry, IsTier, LoadCurrentZoneFromHistory, OnAreaNameChange, OnCurrencyUpdated, OnLoadComplete, OnNameChange, OnPetUpdated, OnPlayerPositionUpdated, OnRenownUpdated, OnUpdate, QueuerCheckMessage, ResolveZoneIdFromName, ScanLore, ScanPerson, SendMessage, ShowWorldMapHooked, SlashCmd, TIMER_UPDATE, UpdateDetailLabels, UpdateGroup, UpdateMap, UpdateVP, UpsertZoneHistory, check, isInAllowedZone, lua_call, untilfunction, zonelistprint, zoomDOWN, zoomUP contexts.

## Seen In

- CMap
- Killer
- Map
- MapMonster
- MapPin
- Minmap
- Queue Queuer
- RealmStatus
- RoR_SoR
- RoR_debolster
- RvRContribution
- SOR
- Squared
- Tome Titan
- TomeTracker
- TurretRange
- VPBreakdown
- WarBoard_Loc
- ZCurse_Profiler
- wbLeadHelper

## Related APIs

- [EA_Window_Backpack.GetItemsFromBackpack](../../globals/functions/global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_ScenarioLobby.OnJoinInstanceWait](../../globals/functions/global_EA_Window_ScenarioLobby.OnJoinInstanceWait.md) (HIGH 100/100) - Global Function
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Notes

- Observed in contexts: AllowedToShowDetailWindow, Broadcast1, Broadcast2, BroadcastA, BroadcastG, BroadcastP
