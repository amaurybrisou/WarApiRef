# LibSlash.RegisterSlashCmd

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

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
| Addons seen in | InfoScroller, NPC Item Sale Price, PartyCast, Soloq, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:30`, `/workspace/data/raw/PartyCast/PartyCast.lua:51`, `/workspace/data/raw/Soloq/Soloq.lua:22`, `/workspace/data/raw/minesweep/minesweep.lua:7`, `/workspace/data/raw/nisp/Source/Nisp.lua:26` |
| Namespaces detected | LibSlash |
| Source kinds | globals, lua_calls |
| Example locations | InfoScroller: InfoScroller.OnInitialize, NPC Item Sale Price: Nisp.Init, PartyCast: PartyCast.Init, Soloq: Soloq.OnInitialize, minesweep: minesweep.OnInitialize |
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
| slashName | Observed as a slash command token. | Observed values: "info", "infoscroller", "nisp" |
| handler | Observed as a command handler callback. | Observed values: function(args)Nisp.SlashHandler(args)end, function(args)Soloq.SlashCmd(args)end, function(input)InfoScroller_config.Slash(input)end |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- InfoScroller
- NPC Item Sale Price
- PartyCast
- Soloq
- minesweep

## Examples

- InfoScroller: InfoScroller.OnInitialize -> LibSlash.RegisterSlashCmd("infoscroller", function(input)InfoScroller_config.Slash(input)end)
- InfoScroller: InfoScroller.OnInitialize -> LibSlash.RegisterSlashCmd("info", function(input)InfoScroller_config.Slash(input)end)
- NPC Item Sale Price: Nisp.Init -> LibSlash.RegisterSlashCmd("nisp", function(args)Nisp.SlashHandler(args)end)
- PartyCast: PartyCast.Init -> LibSlash.RegisterSlashCmd("pc", function(input)PartyCast.Command(input)end)
- PartyCast: PartyCast.Init -> LibSlash.RegisterSlashCmd("partycast", function(input)PartyCast.Command(input)end)
- Soloq: Soloq.OnInitialize -> LibSlash.RegisterSlashCmd("soloq", function(args)Soloq.SlashCmd(args)end)

## Related APIs

- none

## Used With

- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_BEGIN_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAST_TIMER_SETBACK](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAST_TIMER_SETBACK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_END_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_END_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_START_INTERACT_TIMER](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_START_INTERACT_TIMER.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_POST_MODE](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_POST_MODE.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
