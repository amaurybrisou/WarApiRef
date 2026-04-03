# LibSlash.UnregisterSlashCmd

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 21 addons

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
| Addons seen in | ActionFraction, AuctionStats, CDown, Cram The Spam, DAoCBuff, DaemonAssist, DeepSleep, EveryBodyGuard |
| Files seen in | AuctionStats.lua, CDown.lua, Core.lua, CramTheSpam.lua, DALifecycle.lua, DeepSleep.lua, EveryBodyGuard.lua, KillerLifecycle.lua |
| Namespaces detected | LibSlash |
| Source kinds | lua_calls |
| Example locations | ActionFraction: Shutdown, AuctionStats: Shutdown, CDown: Shutdown, Cram The Spam: Shutdown, DAoCBuff: Shutdown, DaemonAssist: Shutdown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 29 |
| Global usage count | 29 |
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
LibSlash.UnregisterSlashCmd(slashName)
```

## Description

Observed wiring slash commands through a shared command-registration table.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| slashName | Observed as a slash command token. | Observed values: "ActionFraction", "CDown", "DeepSleep" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Removes a previously registered callback or command binding.

## Seen In

- ActionFraction
- AuctionStats
- CDown
- Cram The Spam
- DAoCBuff
- DaemonAssist
- DeepSleep
- EveryBodyGuard
- Killer
- Mech
- PotionBar
- QuickWarReport
- RandomMount
- ScenarioStats
- SessionRPs
- Shinies
- ThankTheResser
- Vectors
- WARRatingBuster
- WhoHealedMe
- zMailMod

## Examples

- ActionFraction: Shutdown -> LibSlash.UnregisterSlashCmd("af")
- ActionFraction: Shutdown -> LibSlash.UnregisterSlashCmd("ActionFraction")
- AuctionStats: Shutdown -> LibSlash.UnregisterSlashCmd("au")
- AuctionStats: Shutdown -> LibSlash.UnregisterSlashCmd("uc")
- AuctionStats: Shutdown -> LibSlash.UnregisterSlashCmd("undercut")
- CDown: Shutdown -> LibSlash.UnregisterSlashCmd("CDown")

## Affects

- [SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED](../../systemdata/fields/systemdata_SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_AGRO_MODE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_AGRO_MODE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HEALTH_FADE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HEALTH_FADE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
