# EA_Window_ContextMenu.AddMenuItem

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 18 addons

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
| Addons seen in | AggroMeter, BuffHead, CM_ClosetGoblin, Effigy, Enemy, Killer, MapMonster, MapPin |
| Files seen in | `/workspace/AggroMeter/AggroMeter.lua:251`, `/workspace/AggroMeter/AggroMeter.lua:378`, `/workspace/BuffHead/Setup/LayoutControlFrame.lua:72`, `/workspace/BuffHead/Setup/SetupAdvancedCompression.lua:130`, `/workspace/BuffHead/Setup/SetupAdvancedCompressionItem.lua:217`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItem.lua:439`, `/workspace/BuffHead/Setup/SetupEffectCache.lua:195`, `/workspace/BuffHead/Setup/SetupFilter.lua:154` |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | AggroMeter: AggroMeter.OnTabRBU, AggroMeter: AggroMeter.PickedListMenu, BuffHead: BuffHead.Setup.AdvancedContainersItem.OnContainerRClick, BuffHead: BuffHead.Setup.EffectCache.CreateContextMenu, BuffHead: BuffHead.Setup.Filter.CreateContextMenu, BuffHead: BuffHead.Setup.LayoutControlFrame:CreateContextMenu |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 194 |
| Global usage count | 194 |
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

Observed as a global function across 18 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: GetString(StringTables.Default.LABEL_TO_LOCK), GetString(StringTables.Default.LABEL_TO_UNLOCK), GetStringFormat(StringTables.Default.TEXT_LEAVE_SCENARIO,{L "All Scenarios"}) |
| arg2 | Observed as a function or method reference. | Observed values: AbilitiesWindow.ToggleShowing, AggroMeter.Close, AggroMeter.ToggeBar |
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
- Effigy
- Enemy
- Killer
- MapMonster
- MapPin
- Miracle Grow Remix
- PeaceOut
- PotionBar
- RandomMount
- RoR_SoR
- Shinies
- TidyChat
- TidyQueue
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

- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EA_Window_ContextMenu.AddCascadingMenuItem](global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](../../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](../../window_api/functions/window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowStartScaleAnimation](../../window_api/functions/window_WindowStartScaleAnimation.md) (HIGH 90/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
