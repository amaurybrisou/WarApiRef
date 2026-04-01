# LibSlash.IsSlashCmdRegistered

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 143

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +10 Appears in slash command registration patterns: Observed in shared command registration flows.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AutoBand, BagOMatic, Effigy, JunkDump, PlanB, wbLeadHelper |
| Files seen in | `/workspace/Autoband/AutoBand.lua:30`, `/workspace/Effigy/EffigySlashCommands.lua:19`, `/workspace/JunkDump/JunkDump.lua:50`, `/workspace/PlanB/PlanB.lua:35`, `/workspace/bagomatic/BagOMatic.lua:15`, `/workspace/wbLeadHelper/wbLeadHelper.lua:34` |
| Namespaces detected | LibSlash |
| Source kinds | globals, lua_calls |
| Example locations | AutoBand: AutoBand.init, BagOMatic: BagOMatic.init, Effigy: Effigy.RegisterWithLibSlash, JunkDump: JunkDump.Initialize, PlanB: PlanB.Initialize, wbLeadHelper: wbLeadHelper.OnInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
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
LibSlash.IsSlashCmdRegistered(arg1)
```

## Description

Observed wiring slash commands through a shared command-registration table.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "ab", "bom", "effigy" |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- AutoBand
- BagOMatic
- Effigy
- JunkDump
- PlanB
- wbLeadHelper

## Examples

- AutoBand: AutoBand.init -> LibSlash.IsSlashCmdRegistered("ab")
- BagOMatic: BagOMatic.init -> LibSlash.IsSlashCmdRegistered("bom")
- Effigy: Effigy.RegisterWithLibSlash -> LibSlash.IsSlashCmdRegistered("effigy")
- Effigy: Effigy.RegisterWithLibSlash -> LibSlash.IsSlashCmdRegistered("unitframes")
- JunkDump: JunkDump.Initialize -> LibSlash.IsSlashCmdRegistered("jd")
- PlanB: PlanB.Initialize -> LibSlash.IsSlashCmdRegistered("planb")

## Related APIs

- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Settings.Language.active](../../systemdata/fields/systemdata_SystemData.Settings.Language.active.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Settings.Language.active](../../systemdata/fields/systemdata_SystemData.Settings.Language.active.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
