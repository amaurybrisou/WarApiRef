# SystemData.Events.PLAYER_EFFECTS_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

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
| Addons seen in | BuffHead, BuffHead, DAoCBuff, Enemy, GCDsaver, WSCT, DAoCBuff, Effigy, Enemy, GCDsaver, WSCT |
| Files seen in | `/workspace_addons/BuffHead/Core.lua:207`, `/workspace_addons/DAoCBuff/Source/DAoCBuff.lua:219`, `/workspace_addons/DAoCBuff/Source/DAoCBuff.lua:25`, `/workspace_addons/Effigy/States/EffigyStatePlayer.lua:45`, `/workspace_addons/Enemy/Code/Core/Groups/Groups.lua:22`, `/workspace_addons/GCDsaver/GCDsaver.lua:253`, `/workspace_addons/GCDsaver/GCDsaver.lua:268`, `/workspace_addons/wsct/wsct.lua:117` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | BuffHead.OnGroupEffectsUpdated, BuffHead.SetSelfTracking, DAoCBuff.Initialize, DAoCBuff.OnEvent, DAoCBuff.Shutdown, Effigy.RegisterStateInfoForPlayer |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 20 |
| Global usage count | 20 |
| Local definition count | 0 |
| Documentation references | 4 |
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

Observed SystemData field used by 7 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- BuffHead
- BuffHead, DAoCBuff, Enemy, GCDsaver, WSCT
- DAoCBuff
- Effigy
- Enemy
- GCDsaver
- WSCT

## Related APIs

- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function

## Used With

- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.CUSTOM_UI_SCALE_CHANGED](systemdata_SystemData.Events.CUSTOM_UI_SCALE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_EFFECTS_UPDATED](systemdata_SystemData.Events.GROUP_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH_CLEARED](systemdata_SystemData.Events.PLAYER_DEATH_CLEARED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED](systemdata_SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELEASE_CORPSE](systemdata_SystemData.Events.RELEASE_CORPSE.md) (HIGH 100/100) - SystemData Field
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: BuffHead.OnGroupEffectsUpdated, BuffHead.SetSelfTracking, DAoCBuff.Initialize, DAoCBuff.OnEvent, DAoCBuff.Shutdown, Effigy.RegisterStateInfoForPlayer
