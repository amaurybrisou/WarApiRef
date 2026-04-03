# wstring.format

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 44 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | ActionFraction, AdjustTheTip, AggroMeter, BarText (Influence), CCTV, CaVES, Calling, CastSequence |
| Files seen in | AdjustTheTip.lua, AggroMeter.lua, BarText_Influence.lua, Bars/HealGridCastBar.lua, CCTV.lua, CallingList.lua, Castbar.lua, Clock.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | ActionFraction: Initialize, AdjustTheTip: GetAbilityCastTimeText, AdjustTheTip: UpdateInfluenceTooltip, AggroMeter: OnMouseOverStart, BarText (Influence): PlayerInfluenceUpdated, CCTV: Update |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 461 |
| Global usage count | 461 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
wstring.format(arg1, arg2)
```

## Description

Observed as a global function across 44 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: HealGridSettings[HealGrid.settingsIndex].trackedBuffs[idx].duration, HealGridSettings[HealGrid.settingsIndex].trackedBuffs[idx].warnExpire, L "%.01f" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: #activeReults, ((GameData.Player.Stats[GameData.Stats.WEAPONSKILL].baseValue/(GameData.Player.battleLevel*7.5+50)*.25)*100.0), (AggroMeter.AggroHolder[tostring(MobNumber)][tonumber(TimerNumber)].aggro/AggroMeter.MaxAggro[tostring(MobNumber)])*100 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- ActionFraction
- AdjustTheTip
- AggroMeter
- BarText (Influence)
- CCTV
- CaVES
- Calling
- CastSequence
- Clock
- CoolDownLine
- Crusher
- DAoCBuff
- DuffTimer
- Effigy
- Enemy
- GroupSpotter
- HealGrid
- Hopper
- MapPin
- Motion
- NAMBLA
- Obsidian
- PotionBar
- Pure
- QueuePopTimer
- RealmStatus
- ReliquaryHunter
- SNT_BUTTONS
- SNT_CASTBAR
- SOR
- Shinies
- Squared
- Statdoll
- Statdoll Light
- Statdoll Remix
- Swinger
- TexturedButtons
- Tokens
- VPBreakdown
- Vectors
- WSCT
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- XpStatus+G

## Examples

- ActionFraction: Initialize -> wstring.format(L "%.01f", ActionFraction.Version)
- AdjustTheTip: GetAbilityCastTimeText -> wstring.format(L "%.1f", castTime)
- AdjustTheTip: GetAbilityCastTimeText -> wstring.format(L "%d", castTime)
- AdjustTheTip: UpdateInfluenceTooltip -> wstring.format(L "/%d (", influenceData.rewardLevel[3].amountNeeded)
- AdjustTheTip: UpdateInfluenceTooltip -> wstring.format(L "%d%%,", percent)
- AdjustTheTip: UpdateInfluenceTooltip -> wstring.format(L "%d%%)", percent)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [EA_Window_PublicQuestTracker.GetLocalAreaInfluenceID](global_EA_Window_PublicQuestTracker.GetLocalAreaInfluenceID.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [wstring.gsub](global_wstring.gsub.md) (HIGH 100/100) - Global Function

## Affects

- [SystemData.ChatLogFilters.SHOUT](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
