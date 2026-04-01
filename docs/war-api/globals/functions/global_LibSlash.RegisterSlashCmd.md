# LibSlash.RegisterSlashCmd

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 50 addons

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
| Addons seen in | AdvancedPetAssist, Aura, AutoBand, AutoChannel, AutoMark, BagOMatic, CM_ClosetGoblin, DAoCBuff |
| Files seen in | `/workspace/AdvancedPetAssist/AdvancedPetAssist.lua:193`, `/workspace/Aura/Source/AuraAddon.lua:70`, `/workspace/AutoMark/Source/AutoMark.lua:33`, `/workspace/Autoband/AutoBand.lua:30`, `/workspace/ClosetGoblin/ClosetGoblin.lua:83`, `/workspace/DAoCBuff/Source/DAoCBuff.lua:25`, `/workspace/DaemonAssist/DALifecycle.lua:3`, `/workspace/DeepSleep/DeepSleep.lua:8` |
| Namespaces detected | LibSlash |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.Initialize, Aura: AuraAddon.OnInitialize, AutoBand: AutoBand.init, AutoChannel: AutoChannel.Initialize, AutoMark: AutoMark.OnInitialize, BagOMatic: BagOMatic.init |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 80 |
| Global usage count | 80 |
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
| slashName | Observed as a slash command token. | Observed values: "DeepSleep", "FastInteract", "MiracleGrow" |
| handler | Observed as a command handler callback. | Observed values: AuraAddon.Slash, AutoChannel.sendChatBand, AutoChannel.sendChatBandSay |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- AdvancedPetAssist
- Aura
- AutoBand
- AutoChannel
- AutoMark
- BagOMatic
- CM_ClosetGoblin
- DAoCBuff
- DaemonAssist
- DeepSleep
- Effigy
- FastInteract
- GCDsaver
- GoldTracker
- GuardLine
- Killer
- LibGroup
- LibGuard
- LibSlash
- LibWBToggler
- MapMonster
- MapPin
- MegaphonePlusPlus
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- NPC Item Sale Price
- PeaceOut
- PlanB
- Pocket Palette
- PotionBar
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RRCount
- RVMOD_Manager
- RandomMount
- RetAlert
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- WSCT
- WarBoard
- WarTriage
- WhoHealedMe
- WoH-Reticle
- followTheLeader
- wbLeadHelper

## Examples

- AdvancedPetAssist: AdvancedPetAssist.Initialize -> LibSlash.RegisterSlashCmd("apa", function(input)APA.SlashHandler(input)end)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("aura", AuraAddon.Slash)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("auraconfig", AuraAddon.Slash)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("showaura", AuraAddon.Slash)
- AutoBand: AutoBand.init -> LibSlash.RegisterSlashCmd("autoband", function(msg)AutoBand.parse_cmd(msg)end)
- AutoBand: AutoBand.init -> LibSlash.RegisterSlashCmd("ab", function(msg)AutoBand.parse_cmd(msg)end)

## Related APIs

- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Used With

- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [InterfaceCore.GetResolutionScale](global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
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

- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
