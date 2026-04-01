# LibSlash.RegisterWSlashCmd

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | BuffHead, Crafting Info Tooltip, JunkDump, TidyChat, TidyRoll, TurretRange |
| Files seen in | `/workspace/BuffHead/Core.lua:79`, `/workspace/CraftValueTip/CraftValueTip.lua:33`, `/workspace/JunkDump/JunkDump.lua:50`, `/workspace/TidyChat/TidyChat.lua:189`, `/workspace/TidyRoll/TidyRoll.lua:265`, `/workspace/TurrentRange/Core.lua:41` |
| Namespaces detected | LibSlash |
| Source kinds | globals, lua_calls |
| Example locations | BuffHead: BuffHead.local.RegisterLibs, BuffHead: RegisterLibs, Crafting Info Tooltip: CraftValueTip.Initialize, JunkDump: JunkDump.Initialize, TidyChat: TidyChat.OnLoad, TidyRoll: TidyRoll.OnLoad |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
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
LibSlash.RegisterWSlashCmd(arg1, arg2)
```

## Description

Observed wiring slash commands through a shared command-registration table.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "buffhead", "craftvaluetip", "jd" |
| arg2 | Observed as a function or method reference. | Observed values: TidyChat.ToggleOptions, TidyRoll.ToggleOptions, function(args)BuffHead.SlashCommand(args)end |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- BuffHead
- Crafting Info Tooltip
- JunkDump
- TidyChat
- TidyRoll
- TurretRange

## Examples

- BuffHead: BuffHead.local.RegisterLibs -> LibSlash.RegisterWSlashCmd("buffhead", function(args)BuffHead.SlashCommand(args)end)
- BuffHead: RegisterLibs -> LibSlash.RegisterWSlashCmd("buffhead", function(args)BuffHead.SlashCommand(args)end)
- Crafting Info Tooltip: CraftValueTip.Initialize -> LibSlash.RegisterWSlashCmd("craftvaluetip", function(args)CraftValueTip.SlashCmd(args)end)
- JunkDump: JunkDump.Initialize -> LibSlash.RegisterWSlashCmd("jd", function(args)JunkDump.SlashHandler(args)end)
- JunkDump: JunkDump.Initialize -> LibSlash.RegisterWSlashCmd("junkdump", function(args)JunkDump.SlashHandler(args)end)
- TidyChat: TidyChat.OnLoad -> LibSlash.RegisterWSlashCmd("tchat", TidyChat.ToggleOptions)

## Related APIs

- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [SystemData.Settings.Language.active](../../systemdata/fields/systemdata_SystemData.Settings.Language.active.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event

## Affects

- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [SystemData.Settings.Language.active](../../systemdata/fields/systemdata_SystemData.Settings.Language.active.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
