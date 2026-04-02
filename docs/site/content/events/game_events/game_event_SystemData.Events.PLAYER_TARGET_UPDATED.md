# SystemData.Events.PLAYER_TARGET_UPDATED

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
| Addons seen in | AdvancedPetAssist, Aura, AutoMark, BuffHead, DAoCBuff, Enemy, GCDsaver, LoyalPet |
| Files seen in | `/workspace_addons/AdvancedPetAssist/AdvancedPetAssist.lua:60`, `/workspace_addons/AutoMark/Source/AutoMark.lua:33`, `/workspace_addons/AutoMark/Source/AutoMark.lua:78`, `/workspace_addons/BuffHead/Core.lua:152`, `/workspace_addons/DAoCBuff/Source/DAoCBuff.lua:25`, `/workspace_addons/Enemy/Code/Core/Groups/Groups.lua:22`, `/workspace_addons/Enemy/Code/Core/Main.lua:41`, `/workspace_addons/GCDsaver/GCDsaver.lua:253` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.local.RegisterCoreEvents, AdvancedPetAssist: RegisterCoreEvents, AutoMark: AutoMark.OnInitialize, AutoMark: AutoMark.OnSlashCommand, BuffHead: BuffHead.Initialize, DAoCBuff: DAoCBuff.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 9 |
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

Observed as a shared SystemData runtime event used by 12 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- AdvancedPetAssist
- Aura
- AutoMark
- BuffHead
- DAoCBuff
- Enemy
- GCDsaver
- LoyalPet
- Moth
- Targets
- WoH-Reticle
- followTheLeader

## Registrars And Handlers

- AdvancedPetAssist.OnTargetUpdated
- AutoMark.OnPlayerTargetUpdated
- BuffHead.OnTargetUpdated
- DAoCBuff.OnEventHC
- Enemy.Groups_OnPlayerTargetUpdated
- Enemy.OnPlayerTargetUpdated
- GCDsaver.PLAYER_TARGET_UPDATED
- Moth.UpdateTarget
- RegisterEventHandler
- TargetInfoFix.SET_TARGETINFO_FIX_UPDATE_FLAG_DONOTTOUCH
- Targets.on_target_updated
- WindowRegisterEventHandler
- WoHReticle.UpdateTargets
- followTheLeader.OnTargetUpdated
- global
- window

## Examples

- AdvancedPetAssist: AdvancedPetAssist.local.RegisterCoreEvents -> SystemData.Events.PLAYER_TARGET_UPDATED -> AdvancedPetAssist.OnTargetUpdated
- AdvancedPetAssist: RegisterCoreEvents -> SystemData.Events.PLAYER_TARGET_UPDATED -> AdvancedPetAssist.OnTargetUpdated
- AutoMark: AutoMark.OnInitialize -> SystemData.Events.PLAYER_TARGET_UPDATED -> AutoMark.OnPlayerTargetUpdated
- AutoMark: AutoMark.OnSlashCommand -> SystemData.Events.PLAYER_TARGET_UPDATED -> AutoMark.OnPlayerTargetUpdated
- BuffHead: BuffHead.Initialize -> SystemData.Events.PLAYER_TARGET_UPDATED -> BuffHead.OnTargetUpdated
- DAoCBuff: DAoCBuff.Initialize -> SystemData.Events.PLAYER_TARGET_UPDATED -> DAoCBuff.OnEventHC

## Related APIs

- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- [LPET](game_event_LPET.md) (HIGH 93/100) - Game Event

## Affects

- none

## Notes

- none
