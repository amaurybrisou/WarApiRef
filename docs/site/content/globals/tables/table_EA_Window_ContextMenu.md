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
| Addons seen in | ActionFraction, AdjustTheTip, AggroMeter, AutoBand, BuffHead, CM_ClosetGoblin, CastSequence, Crusher |
| Files seen in | Code/CombatLog/CombatLogEpsWindow.lua, Code/CombatLog/CombatLogStatsWindow.lua, Code/CombatLog/CombatLogTargetDefenseWindow.lua, Code/Core/Main.lua, Code/Marks/Marks.lua, Code/ScenarioInfo/ScenarioInfo.lua, Manager/Advanced.lua, Manager/Manager.lua |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | ActionFraction: AutoHideOnMouseOver, ActionFraction: RightClick, ActionFraction: SetFontSelectionMenu, ActionFraction: SetPresetLocation, AdjustTheTip: ContextMenu_Finalize, AdjustTheTip: CreateContextMenu |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 685 |
| Global usage count | 11 |
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

Shared function table with 11 member functions; the primary API surface for 54 addons.

## Functions

- EA_Window_ContextMenu.AddCascadingMenuItem
- EA_Window_ContextMenu.AddMenuDivider
- EA_Window_ContextMenu.AddMenuItem
- EA_Window_ContextMenu.AddUserDefinedMenuItem
- EA_Window_ContextMenu.CreateContextMenu
- EA_Window_ContextMenu.CreateDefaultContextMenu
- EA_Window_ContextMenu.Finalize
- EA_Window_ContextMenu.GameActionData
- EA_Window_ContextMenu.Hide
- EA_Window_ContextMenu.HideAll
- EA_Window_ContextMenu.OnMouseOverDefaultMenuItem

## Observed Members

- none

## Seen In

- ActionFraction
- AdjustTheTip
- AggroMeter
- AutoBand
- BuffHead
- CM_ClosetGoblin
- CastSequence
- Crusher
- Dascore
- DuelInvite
- DuffTimer
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EZCraftX
- Effigy
- Enemy
- HealGrid
- Killer
- LibAddonButton
- Map
- MapMonster
- MapPin
- MarkBuff
- Miracle Grow Remix
- MiracleGrow
- MoraleCircle
- Motion
- NerfedButtons
- Obsidian
- PeaceOut
- PlayEffectsOn
- PotionBar
- Pure
- QuickNameActions+
- RandomMount
- Refer
- RoR_SoR
- SOR
- Shinies
- ShowMeTheBubbles
- TacticSetNames
- TastyButtons
- TexturedButtons
- TidyChat
- TidyQueue
- TurretRange
- Vectors
- WARCommander
- WarBoard
- WarBoard_Loc
- WarBoard_Menu
- XpStatus+G
- nLootLink
- scenarioInfo

## Examples

- ActionFraction: AutoHideOnMouseOver -> EA_Window_ContextMenu.Hide(EA_Window_ContextMenu.CONTEXT_MENU_3)
- ActionFraction: RightClick -> EA_Window_ContextMenu.CreateContextMenu(windowName, EA_Window_ContextMenu.CONTEXT_MENU_1)
- ActionFraction: RightClick -> EA_Window_ContextMenu.AddMenuItem(GetString(StringTables.Default.LABEL_TO_LOCK), ActionFractionWindow.OnLock, not movable, true, EA_Window_ContextMenu.CONTEXT_MENU_1)
- ActionFraction: RightClick -> EA_Window_ContextMenu.AddMenuItem(GetString(StringTables.Default.LABEL_TO_UNLOCK), ActionFractionWindow.OnUnlock, movable, true, EA_Window_ContextMenu.CONTEXT_MENU_1)
- ActionFraction: RightClick -> EA_Window_ContextMenu.AddMenuDivider(EA_Window_ContextMenu.CONTEXT_MENU_1)
- ActionFraction: RightClick -> EA_Window_ContextMenu.AddCascadingMenuItem(LOC_TEXT.PRESET_LOCATIONS, ActionFractionWindow.SetPresetLocation, false, EA_Window_ContextMenu.CONTEXT_MENU_1)

## Notes

- none
