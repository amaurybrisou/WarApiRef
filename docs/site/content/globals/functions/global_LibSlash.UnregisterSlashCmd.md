# LibSlash.UnregisterSlashCmd

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

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
| Addons seen in | DAoCBuff, DaemonAssist, DeepSleep, Killer, PotionBar, QuickWarReport, RandomMount, Shinies |
| Files seen in | `/workspace/DAoCBuff/Source/DAoCBuff.lua:219`, `/workspace/DaemonAssist/DALifecycle.lua:36`, `/workspace/DeepSleep/DeepSleep.lua:32`, `/workspace/Killer/KillerLifecycle.lua:100`, `/workspace/PotionBar/source/Main.lua:292`, `/workspace/QuickWarReport/QWRLifecycle.lua:90`, `/workspace/RandomMount/RandomMount.lua:64`, `/workspace/Shinies/Source/Shinies.lua:206` |
| Namespaces detected | LibSlash |
| Source kinds | globals, lua_calls |
| Example locations | DAoCBuff: DAoCBuff.Shutdown, DaemonAssist: DaemonAssist.Shutdown, DeepSleep: DeepSleep.Shutdown, Killer: Killer.Shutdown, PotionBar: PotionBar.Shutdown, QuickWarReport: QuickWarReport.Shutdown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
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
LibSlash.UnregisterSlashCmd(slashName)
```

## Description

Observed wiring slash commands through a shared command-registration table.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| slashName | Observed as a slash command token. | Observed values: "DeepSleep", "da", "daocbuff" |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Removes a previously registered callback or command binding.

## Seen In

- DAoCBuff
- DaemonAssist
- DeepSleep
- Killer
- PotionBar
- QuickWarReport
- RandomMount
- Shinies
- WhoHealedMe

## Examples

- DAoCBuff: DAoCBuff.Shutdown -> LibSlash.UnregisterSlashCmd("daocbuff")
- DAoCBuff: DAoCBuff.Shutdown -> LibSlash.UnregisterSlashCmd("resetdaocbuff")
- DaemonAssist: DaemonAssist.Shutdown -> LibSlash.UnregisterSlashCmd("da")
- DeepSleep: DeepSleep.Shutdown -> LibSlash.UnregisterSlashCmd("ds")
- DeepSleep: DeepSleep.Shutdown -> LibSlash.UnregisterSlashCmd("DeepSleep")
- Killer: Killer.Shutdown -> LibSlash.UnregisterSlashCmd("killer")

## Related APIs

- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Used With

- [LayoutEditor.UnregisterWindow](../../window_api/functions/window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [SystemData.Events.CAMPAIGN_ZONE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.CAMPAIGN_ZONE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CUSTOM_UI_SCALE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.CUSTOM_UI_SCALE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH_CLEARED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH_CLEARED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ZONE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ZONE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELEASE_CORPSE](../../systemdata/fields/systemdata_SystemData.Events.RELEASE_CORPSE.md) (HIGH 100/100) - SystemData Field
- [UnregisterEventHandler](global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- [SystemData.Events.CAMPAIGN_ZONE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.CAMPAIGN_ZONE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CUSTOM_UI_SCALE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.CUSTOM_UI_SCALE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH_CLEARED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH_CLEARED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ZONE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ZONE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELEASE_CORPSE](../../systemdata/fields/systemdata_SystemData.Events.RELEASE_CORPSE.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
