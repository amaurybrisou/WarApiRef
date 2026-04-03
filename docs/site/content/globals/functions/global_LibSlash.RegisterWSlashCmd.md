# LibSlash.RegisterWSlashCmd

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 26 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +10 Appears in slash command registration patterns: Observed in shared command registration flows.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AbilityAlert, ActionFraction, ActionPointWatch, BarText (Influence), BuffHead, Calling, CastSequence, Crafting Info Tooltip |
| Files seen in | AbilityAlert.lua, ActionPointWatch.lua, BarText_Influence.lua, Calling.lua, Core.lua, CraftValueTip.lua, CraftingWillard.lua, CramTheSpam.lua |
| Namespaces detected | LibSlash |
| Source kinds | lua_calls |
| Example locations | AbilityAlert: Initialize, ActionFraction: Initialize, ActionPointWatch: Initialize, BarText (Influence): OnInitialize, BuffHead: RegisterLibs, Calling: TryRegisterSlashCommands |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 50 |
| Global usage count | 50 |
| Local definition count | 0 |
| Documentation references | 0 |
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
LibSlash.RegisterWSlashCmd(arg1, arg2)
```

## Description

Observed wiring slash commands through a shared command-registration table.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "/actionpointwatch", "ActionFraction", "CraftingWillard" |
| arg2 | Observed as a function or method reference. | Observed values: NerfedTalks.Command, RealmStatus.Show, RealmStatus.ToggleShowingHistoryWindow |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- AbilityAlert
- ActionFraction
- ActionPointWatch
- BarText (Influence)
- BuffHead
- Calling
- CastSequence
- Crafting Info Tooltip
- CraftingWillard
- Cram The Spam
- DetauntHelper
- GroupRange
- JunkDump
- Keyset
- KeysetMonsterPlay
- Lib RuString
- NerfedButtons
- Obsidian
- RealmStatus
- Statdoll Remix
- TidyChat
- TidyRoll
- TimeToDie
- Tokens
- TurretRange
- whom

## Examples

- AbilityAlert: Initialize -> LibSlash.RegisterWSlashCmd("aa", function(args)AbilityAlert.SlashHandler(args)end)
- AbilityAlert: Initialize -> LibSlash.RegisterWSlashCmd("abilityalert", function(args)AbilityAlert.SlashHandler(args)end)
- ActionFraction: Initialize -> LibSlash.RegisterWSlashCmd("af", slashHandler)
- ActionFraction: Initialize -> LibSlash.RegisterWSlashCmd("ActionFraction", slashHandler)
- ActionPointWatch: Initialize -> LibSlash.RegisterWSlashCmd("apw", function(args)ActionPointWatch.SlashHandler(args)end)
- ActionPointWatch: Initialize -> LibSlash.RegisterWSlashCmd("/actionpointwatch", function(args)ActionPointWatch.SlashHandler(args)end)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event

## Used With

- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [SystemData.ChatLogFilters.SHOUT](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field
- [RegisterEventHandler](global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Affects

- [SystemData.ChatLogFilters.SHOUT](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
