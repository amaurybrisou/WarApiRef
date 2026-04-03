# SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 19 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 168

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AbilityNotifier, ActionBarHide, ArmorGraphicToggle, Atlas, Aura, CNC, Crusher, DAoCBuff |
| Files seen in | AbilityNotifier.lua, ActionBarHide.lua, ArmorGraphicToggle.lua, CNC.lua, Code/Core/Groups/Groups.lua, DPSMeter.lua, IHYTM.lua, Map.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | ArmorGraphicToggle_Disable, ArmorGraphicToggle_Enable, GroupsInitialize, Initialize, JumpStartEventBasedAuras, LoadUnitFrame |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 20 |
| Global usage count | 20 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
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

SystemData.SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED field accessed by 19 addons; commonly found in ArmorGraphicToggle_Disable and ArmorGraphicToggle_Enable, GroupsInitialize, Initialize, JumpStartEventBasedAuras, LoadUnitFrame, OnDisable, OnEnable, OnInit, OnInitialize, OnLoad, OnShutdown, OnSkavenStatusChange, RegisterEventHandlers, RegisterSelfEvents, RegisterStateInfoForCombat, Shutdown, UnloadUnitFrame, UnregisterSelfEvents, lua_call contexts.

## Seen In

- AbilityNotifier
- ActionBarHide
- ArmorGraphicToggle
- Atlas
- Aura
- CNC
- Crusher
- DAoCBuff
- DPSMeter
- Effigy
- Enemy
- I HATE YOU THIS MUCH
- Map
- MoraleSet
- Motion
- PlanB
- Pure
- Shinies
- WSCT

## Related APIs

- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Notes

- Observed in contexts: ArmorGraphicToggle_Disable, ArmorGraphicToggle_Enable, GroupsInitialize, Initialize, JumpStartEventBasedAuras, LoadUnitFrame
