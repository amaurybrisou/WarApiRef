# EA_Window_ContextMenu

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 125

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AggroMeter, BuffHead, CM_ClosetGoblin, EA_UiDebugTools, Effigy, Enemy, Killer, MapMonster |
| Files seen in | `/workspace/AggroMeter/AggroMeter.lua:251`, `/workspace/AggroMeter/AggroMeter.lua:378`, `/workspace/BuffHead/Setup/LayoutControlFrame.lua:72`, `/workspace/BuffHead/Setup/SelectFont.lua:62`, `/workspace/BuffHead/Setup/SelectFont.lua:76`, `/workspace/BuffHead/Setup/SetupAdvancedCompression.lua:130`, `/workspace/BuffHead/Setup/SetupAdvancedCompressionItem.lua:217`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItem.lua:439` |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | AggroMeter: AggroMeter.OnTabRBU, AggroMeter: AggroMeter.PickedListMenu, BuffHead: BuffHead.Setup.AdvancedContainersItem.OnContainerRClick, BuffHead: BuffHead.Setup.EffectCache.CreateContextMenu, BuffHead: BuffHead.Setup.Filter.CreateContextMenu, BuffHead: BuffHead.Setup.LayoutControlFrame:CreateContextMenu |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 377 |
| Global usage count | 8 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed shared global table or namespace surfaced in 22 addons.

## Functions

- EA_Window_ContextMenu.AddCascadingMenuItem
- EA_Window_ContextMenu.AddMenuDivider
- EA_Window_ContextMenu.AddMenuItem
- EA_Window_ContextMenu.AddUserDefinedMenuItem
- EA_Window_ContextMenu.CreateContextMenu
- EA_Window_ContextMenu.Finalize
- EA_Window_ContextMenu.Hide
- EA_Window_ContextMenu.HideAll

## Observed Members

- none

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
- PeaceOut
- PotionBar
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyQueue
- TurretRange
- WarBoard

## Examples

- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.CreateContextMenu(SystemData.MouseOverWindow.name, EA_Window_ContextMenu.CONTEXT_MENU_1, L "Options")
- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuDivider(EA_Window_ContextMenu.CONTEXT_MENU_1)
- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuItem(L "<icon00057> Enabled", AggroMeter.ToggeEnable, false, true)
- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuItem(L "<icon00058> Disabled", AggroMeter.ToggeEnable, false, true)
- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuItem(L "   <icon00057> Champions", MakeCallBack(1), not AggroMeter.Enabled, true)
- AggroMeter: AggroMeter.OnTabRBU -> EA_Window_ContextMenu.AddMenuItem(L "   <icon00058> Champions", MakeCallBack(1), not AggroMeter.Enabled, true)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
