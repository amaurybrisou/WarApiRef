# EA_Window_ContextMenu.AddMenuItem

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 11 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AggroMeter, BuffHead, CM_ClosetGoblin, Enemy, Killer, PotionBar, RoR_SoR, Shinies |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.lua:251`, `/workspace/data/raw/AggroMeter/AggroMeter.lua:378`, `/workspace/data/raw/BuffHead/Setup/LayoutControlFrame.lua:72`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.lua:130`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.lua:217`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.lua:439`, `/workspace/data/raw/BuffHead/Setup/SetupEffectCache.lua:195`, `/workspace/data/raw/BuffHead/Setup/SetupFilter.lua:154` |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | AggroMeter: AggroMeter.OnTabRBU, AggroMeter: AggroMeter.PickedListMenu, BuffHead: BuffHead.Setup.AdvancedContainersItem.OnContainerRClick, BuffHead: BuffHead.Setup.EffectCache.CreateContextMenu, BuffHead: BuffHead.Setup.Filter.CreateContextMenu, BuffHead: BuffHead.Setup.LayoutControlFrame:CreateContextMenu |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 107 |
| Global usage count | 107 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
EA_Window_ContextMenu.AddMenuItem(arg1, arg2, arg3, arg4)
```

## Description

Observed as a global function across 11 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: GetString(StringTables.Default.LABEL_TO_LOCK), GetString(StringTables.Default.LABEL_TO_UNLOCK), L "   <icon00057> Champions" |
| arg2 | Observed as a function or method reference. | Observed values: AggroMeter.Close, AggroMeter.ToggeBar, AggroMeter.ToggeEnable |
| arg3 | Observed as a boolean toggle. | Observed values: (not GameData.Player.rvrZoneFlagged), (not IsWarBandActive()), (not PartyUtils.IsPartyActive()) |
| arg4 | Observed as a boolean toggle. | Observed values: true |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AggroMeter
- BuffHead
- CM_ClosetGoblin
- Enemy
- Killer
- PotionBar
- RoR_SoR
- Shinies
- TidyChat
- TurretRange
- WarBoard

## Examples

- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuItem(L "<icon00057> Enabled", AggroMeter.ToggeEnable, false, true)
- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuItem(L "<icon00058> Disabled", AggroMeter.ToggeEnable, false, true)
- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuItem(L "   <icon00057> Champions", MakeCallBack(1), not AggroMeter.Enabled, true)
- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuItem(L "   <icon00058> Champions", MakeCallBack(1), not AggroMeter.Enabled, true)
- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuItem(L "   <icon00057> Heroes", MakeCallBack(2), not AggroMeter.Enabled, true)
- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuItem(L "   <icon00058> Heroes", MakeCallBack(2), not AggroMeter.Enabled, true)

## Related APIs

- none

## Used With

- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Triggered By

- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
