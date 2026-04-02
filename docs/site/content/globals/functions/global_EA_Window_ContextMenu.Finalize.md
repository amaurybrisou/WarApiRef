# EA_Window_ContextMenu.Finalize

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 20 addons

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
| Addons seen in | AggroMeter, BuffHead, CM_ClosetGoblin, EA_UiDebugTools, Effigy, Enemy, Killer, MapMonster |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.lua:251`, `/workspace_addons/AggroMeter/AggroMeter.lua:378`, `/workspace_addons/BuffHead/Setup/LayoutControlFrame.lua:72`, `/workspace_addons/BuffHead/Setup/SelectFont.lua:76`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompression.lua:130`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItem.lua:217`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItem.lua:439`, `/workspace_addons/BuffHead/Setup/SetupEffectCache.lua:195` |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | AggroMeter: AggroMeter.OnTabRBU, AggroMeter: AggroMeter.PickedListMenu, BuffHead: BuffHead.Setup.AdvancedContainersItem.OnContainerRClick, BuffHead: BuffHead.Setup.EffectCache.CreateContextMenu, BuffHead: BuffHead.Setup.Filter.CreateContextMenu, BuffHead: BuffHead.Setup.LayoutControlFrame:CreateContextMenu |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 66 |
| Global usage count | 66 |
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
EA_Window_ContextMenu.Finalize()
```

## Description

Observed as a global function across 20 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AggroMeter
- BuffHead
- CM_ClosetGoblin
- EA_UiDebugTools
- Effigy
- Enemy
- Killer
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MoraleCircle
- PotionBar
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TurretRange
- WarBoard

## Examples

- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.Finalize()
- AggroMeter: AggroMeter.PickedListMenu -> EA_Window_ContextMenu.Finalize()
- BuffHead: BuffHead.Setup.AdvancedContainersItem.OnContainerRClick -> EA_Window_ContextMenu.Finalize()
- BuffHead: BuffHead.Setup.EffectCache.CreateContextMenu -> EA_Window_ContextMenu.Finalize()
- BuffHead: BuffHead.Setup.Filter.CreateContextMenu -> EA_Window_ContextMenu.Finalize()
- BuffHead: BuffHead.Setup.LayoutControlFrame:CreateContextMenu -> EA_Window_ContextMenu.Finalize()

## Related APIs

- none

## Used With

- [EA_Window_ContextMenu.AddCascadingMenuItem](global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddUserDefinedMenuItem](global_EA_Window_ContextMenu.AddUserDefinedMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function

## Triggered By

- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [ListBox](../../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
