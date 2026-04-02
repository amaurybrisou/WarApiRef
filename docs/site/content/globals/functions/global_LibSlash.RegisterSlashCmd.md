# LibSlash.RegisterSlashCmd

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 40 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +10 Appears in slash command registration patterns: Observed in shared command registration flows.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, Aura, AutoMark, BagOMatic, CM_ClosetGoblin, DAoCBuff, DaemonAssist, Effigy |
| Files seen in | `/workspace_addons/AdvancedPetAssist/AdvancedPetAssist.lua:193`, `/workspace_addons/Aura/Source/AuraAddon.lua:70`, `/workspace_addons/AutoMark/Source/AutoMark.lua:33`, `/workspace_addons/ClosetGoblin/ClosetGoblin.lua:86`, `/workspace_addons/DAoCBuff/Source/DAoCBuff.lua:25`, `/workspace_addons/DaemonAssist/DALifecycle.lua:3`, `/workspace_addons/Effigy/EffigySlashCommands.lua:19`, `/workspace_addons/FastInteract/FastInteract.lua:132` |
| Namespaces detected | LibSlash |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.Initialize, Aura: AuraAddon.OnInitialize, AutoMark: AutoMark.OnInitialize, BagOMatic: BagOMatic.init, CM_ClosetGoblin: ClosetGoblin.Initialize, DAoCBuff: DAoCBuff.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 62 |
| Global usage count | 62 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | no |
| Slash command presence | yes |
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
LibSlash.RegisterSlashCmd(slashName, handler)
```

## Description

Observed wiring slash commands through a shared command-registration table.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| slashName | Observed as a slash command token. | Observed values: "FastInteract", "MiracleGrow", "MiracleGrow2" |
| handler | Observed as a command handler callback. | Observed values: AuraAddon.Slash, AutoMark.OnSlashCommand, ClosetGoblin.OnSlashCommand |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- AdvancedPetAssist
- Aura
- AutoMark
- BagOMatic
- CM_ClosetGoblin
- DAoCBuff
- DaemonAssist
- Effigy
- FastInteract
- GCDsaver
- GuardLine
- Killer
- LibGroup
- LibGuard
- LibSlash
- LibWBToggler
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- NPC Item Sale Price
- PlanB
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- QuickWarReport
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- WSCT
- WarBoard
- WhoHealedMe
- WoH-Reticle
- followTheLeader
- wbLeadHelper

## Examples

- AdvancedPetAssist: AdvancedPetAssist.Initialize -> LibSlash.RegisterSlashCmd("apa", function(input)APA.SlashHandler(input)end)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("aura", AuraAddon.Slash)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("auraconfig", AuraAddon.Slash)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("showaura", AuraAddon.Slash)
- AutoMark: AutoMark.OnInitialize -> LibSlash.RegisterSlashCmd("automark", AutoMark.OnSlashCommand)
- BagOMatic: BagOMatic.init -> LibSlash.RegisterSlashCmd("bagomatic", function(msg)BagOMatic.parse_cmd(msg)end)

## Related APIs

- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Used With

- [EA_ChatWindow.OnKeyEnter](global_EA_ChatWindow.OnKeyEnter.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [InterfaceCore.GetResolutionScale](global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
