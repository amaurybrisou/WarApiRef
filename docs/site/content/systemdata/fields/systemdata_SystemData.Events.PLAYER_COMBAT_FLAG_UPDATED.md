# SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 10 addons

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
| Addons seen in | Ace, Ace, AdvancedPetAssist, DAoCBuff, Enemy, PlanB, Shinies, WSCT, AdvancedPetAssist, Aura, DAoCBuff, Effigy, Enemy, PlanB |
| Files seen in | `/workspace_addons/AdvancedPetAssist/AdvancedPetAssist.lua:60`, `/workspace_addons/Aura/Source/AuraEngine.lua:312`, `/workspace_addons/DAoCBuff/Source/DAoCBuff.lua:219`, `/workspace_addons/DAoCBuff/Source/DAoCBuff.lua:25`, `/workspace_addons/Effigy/States/EffigyStateCombat.lua:4`, `/workspace_addons/Enemy/Code/Core/Groups/Groups.lua:22`, `/workspace_addons/PlanB/PlanB.lua:35`, `/workspace_addons/wsct/wsct.lua:117` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | AdvancedPetAssist.OnCombatFlagUpdated, AuraEngine.JumpStartEventBasedAuras, DAoCBuff.CombatUpdate, DAoCBuff.Initialize, DAoCBuff.Shutdown, Effigy.RegisterStateInfoForCombat |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 19 |
| Global usage count | 19 |
| Local definition count | 0 |
| Documentation references | 4 |
| Initialization flow references | 6 |
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

Observed SystemData field used by 10 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- Ace
- Ace, AdvancedPetAssist, DAoCBuff, Enemy, PlanB, Shinies, WSCT
- AdvancedPetAssist
- Aura
- DAoCBuff
- Effigy
- Enemy
- PlanB
- Shinies
- WSCT

## Related APIs

- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function

## Used With

- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.CUSTOM_UI_SCALE_CHANGED](systemdata_SystemData.Events.CUSTOM_UI_SCALE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_EFFECTS_UPDATED](systemdata_SystemData.Events.GROUP_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH_CLEARED](systemdata_SystemData.Events.PLAYER_DEATH_CLEARED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](systemdata_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
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

- Observed in contexts: AdvancedPetAssist.OnCombatFlagUpdated, AuraEngine.JumpStartEventBasedAuras, DAoCBuff.CombatUpdate, DAoCBuff.Initialize, DAoCBuff.Shutdown, Effigy.RegisterStateInfoForCombat
