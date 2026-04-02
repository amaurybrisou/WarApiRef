# SystemData.Events.PLAYER_ZONE_CHANGED

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin, Enemy, GCDsaver, Killer, LibWBToggler, MapMonster, Shinies |
| Files seen in | `/workspace_addons/ClosetGoblin/ClosetGoblin.lua:1300`, `/workspace_addons/Enemy/Code/Core/Groups/Groups.lua:22`, `/workspace_addons/Enemy/Code/GroupIcons/GroupIcons.lua:62`, `/workspace_addons/Enemy/Code/TalismanAlerter/TalismanAlerter.lua:37`, `/workspace_addons/GCDsaver/GCDsaver.lua:253`, `/workspace_addons/Killer/KillerLifecycle.lua:4`, `/workspace_addons/LibWarBoardToggler/LibWBTogglerManager.lua:12`, `/workspace_addons/MapMonster/Source/MapMonster_Player.lua:72` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | CM_ClosetGoblin: ClosetGoblin.ZoneChangeInit, Enemy: Enemy.GroupsInitialize, Enemy: Enemy._GroupIconsEnabledChanged, Enemy: Enemy._TalismanAlerterEnabledChanged, GCDsaver: GCDsaver.RegisterEvents, Killer: Killer.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 4 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
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

Observed as a shared SystemData runtime event used by 7 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- CM_ClosetGoblin
- Enemy
- GCDsaver
- Killer
- LibWBToggler
- MapMonster
- Shinies

## Registrars And Handlers

- ClosetGoblin.PlayerZoneChanged
- Enemy.GroupIcons_OnPlayerZoneChanged
- Enemy.GroupsCheckTarget
- Enemy.TalismanAlerter_Update
- GCDsaver.PLAYER_ZONE_CHANGED
- Killer.OnZoneChanged
- LibWBTogglerManager.CheckMods
- MapMonster.OnPlayerZoneChanged
- RegisterEventHandler
- ShiniesAPI.Core_OnPlayerZoneChanged
- global

## Examples

- CM_ClosetGoblin: ClosetGoblin.ZoneChangeInit -> SystemData.Events.PLAYER_ZONE_CHANGED -> ClosetGoblin.PlayerZoneChanged
- Enemy: Enemy.GroupsInitialize -> SystemData.Events.PLAYER_ZONE_CHANGED -> Enemy.GroupsCheckTarget
- Enemy: Enemy._GroupIconsEnabledChanged -> SystemData.Events.PLAYER_ZONE_CHANGED -> Enemy.GroupIcons_OnPlayerZoneChanged
- Enemy: Enemy._TalismanAlerterEnabledChanged -> SystemData.Events.PLAYER_ZONE_CHANGED -> Enemy.TalismanAlerter_Update
- GCDsaver: GCDsaver.RegisterEvents -> SystemData.Events.PLAYER_ZONE_CHANGED -> GCDsaver.PLAYER_ZONE_CHANGED
- Killer: Killer.Initialize -> SystemData.Events.PLAYER_ZONE_CHANGED -> Killer.OnZoneChanged

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
