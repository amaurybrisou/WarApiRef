# wstring.format

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 9 addons

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
| Addons seen in | AggroMeter, DAoCBuff, Effigy, Enemy, MapPin, PotionBar, Shinies, TexturedButtons |
| Files seen in | `/workspace/AggroMeter/AggroMeter.lua:205`, `/workspace/DAoCBuff/Source/DAoCBuffFrames.lua:121`, `/workspace/Effigy/Effigy.lua:227`, `/workspace/Effigy/States/EffigyStatePlayer.lua:483`, `/workspace/Effigy/States/EffigyStatePlayer.lua:532`, `/workspace/Enemy/Code/Core/Utils.lua:727`, `/workspace/Enemy/Code/Core/Utils.lua:820`, `/workspace/MapPin/source/MapPin.lua:2465` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | AggroMeter: AggroMeter.OnMouseOverStart, DAoCBuff: DAoCBuffFrame:UpdateI_hpte, Effigy: Effigy.LoadLayout, Effigy: Effigy.SetRPLabelText, Effigy: Effigy.SetXPLabelText, Enemy: Enemy.DateTimeToString |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 27 |
| Global usage count | 27 |
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

Observed as a global function across 9 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: L "%.01f", L "%.02f", L "%.2f K" |
| arg2 | Observed as a function or method reference. | Observed values: (AggroMeter.AggroHolder[tostring(MobNumber)][tonumber(TimerNumber)].aggro/AggroMeter.MaxAggro[tostring(MobNumber)])*100, (duration*10)/10, PotionBar.Settings.Version |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- AggroMeter
- DAoCBuff
- Effigy
- Enemy
- MapPin
- PotionBar
- Shinies
- TexturedButtons
- WSCT

## Examples

- AggroMeter: AggroMeter.OnMouseOverStart -> wstring.format(L "%.01f", (AggroMeter.AggroHolder[tostring(MobNumber)][tonumber(TimerNumber)].aggro/AggroMeter.MaxAggro[tostring(MobNumber)])*100)
- DAoCBuff: DAoCBuffFrame:UpdateI_hpte -> wstring.format(s, (duration*10)/10)
- Effigy: Effigy.LoadLayout -> wstring.format(L "%0.2f", towstring(scalemath))
- Effigy: Effigy.SetRPLabelText -> wstring.format(L "%d", curPoints/maRpoints*100)
- Effigy: Effigy.SetXPLabelText -> wstring.format(L "%d", curPoints/maxPoints*100)
- Enemy: Enemy.DateTimeToString -> wstring.format(L "%02d.%02d.%04d %02d:%02d:%02d", dt.day, dt.month, dt.year, dt.hours, dt.minutes, dt.seconds)

## Related APIs

- none

## Used With

- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [InterfaceCore.GetResolutionScale](global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.screenResolution.x](../../systemdata/fields/systemdata_SystemData.screenResolution.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.screenResolution.y](../../systemdata/fields/systemdata_SystemData.screenResolution.y.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event

## Affects

- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [GameData.Player.Experience.curXpEarned](../../gamedata/fields/gamedata_GameData.Player.Experience.curXpEarned.md) (HIGH 100/100) - GameData Field
- [GameData.Player.Experience.curXpNeeded](../../gamedata/fields/gamedata_GameData.Player.Experience.curXpNeeded.md) (HIGH 100/100) - GameData Field
- [GameData.Player.Experience.restXp](../../gamedata/fields/gamedata_GameData.Player.Experience.restXp.md) (HIGH 100/100) - GameData Field
- [GameData.Player.Renown.curRenownEarned](../../gamedata/fields/gamedata_GameData.Player.Renown.curRenownEarned.md) (HIGH 100/100) - GameData Field
- [GameData.Player.Renown.curRenownNeeded](../../gamedata/fields/gamedata_GameData.Player.Renown.curRenownNeeded.md) (HIGH 100/100) - GameData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.screenResolution.x](../../systemdata/fields/systemdata_SystemData.screenResolution.x.md) (HIGH 100/100) - SystemData Field
- [SystemData.screenResolution.y](../../systemdata/fields/systemdata_SystemData.screenResolution.y.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
