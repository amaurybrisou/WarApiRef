# wstring.format

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 7 addons

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
| Addons seen in | AggroMeter, DAoCBuff, Enemy, PotionBar, Shinies, TexturedButtons, WSCT |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.lua:205`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffFrames.lua:121`, `/workspace/data/raw/Enemy/Code/Core/Utils.lua:727`, `/workspace/data/raw/Enemy/Code/Core/Utils.lua:820`, `/workspace/data/raw/PotionBar/source/Main.lua:189`, `/workspace/data/raw/Shinies/Modules/Config/Shinies-Config-General.lua:106`, `/workspace/data/raw/Shinies/Modules/Config/Shinies-Config-General.lua:89`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Auctions.lua:571` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | AggroMeter: AggroMeter.OnMouseOverStart, DAoCBuff: DAoCBuffFrame:UpdateI_hpte, Enemy: Enemy.DateTimeToString, Enemy: Enemy.FormatNumberShort, PotionBar: PotionBar.Initialize, Shinies: UpdateSummaryDisplay |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 19 |
| Global usage count | 19 |
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

Observed as a global function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: L "%.01f", L "%.2f K", L "%.2f M" |
| arg2 | Observed as a function or method reference. | Observed values: (AggroMeter.AggroHolder[tostring(MobNumber)][tonumber(TimerNumber)].aggro/AggroMeter.MaxAggro[tostring(MobNumber)])*100, (duration*10)/10, PotionBar.Settings.Version |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AggroMeter
- DAoCBuff
- Enemy
- PotionBar
- Shinies
- TexturedButtons
- WSCT

## Examples

- AggroMeter: AggroMeter.OnMouseOverStart -> wstring.format(L "%.01f", (AggroMeter.AggroHolder[tostring(MobNumber)][tonumber(TimerNumber)].aggro/AggroMeter.MaxAggro[tostring(MobNumber)])*100)
- DAoCBuff: DAoCBuffFrame:UpdateI_hpte -> wstring.format(s, (duration*10)/10)
- Enemy: Enemy.DateTimeToString -> wstring.format(L "%02d.%02d.%04d %02d:%02d:%02d", dt.day, dt.month, dt.year, dt.hours, dt.minutes, dt.seconds)
- Enemy: Enemy.FormatNumberShort -> wstring.format(L "%.2f M", number/1000000)
- Enemy: Enemy.FormatNumberShort -> wstring.format(L "%.2f K", number/1000)
- PotionBar: PotionBar.Initialize -> wstring.format(L "%.01f", PotionBar.Settings.Version)

## Related APIs

- none

## Used With

- none

## Triggered By

- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event

## Affects

- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [SystemData.ChatLogFilters.SHOUT](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_RANK_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_RANK_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
