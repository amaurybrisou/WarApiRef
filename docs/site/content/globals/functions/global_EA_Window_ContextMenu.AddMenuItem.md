# EA_Window_ContextMenu.AddMenuItem

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 39 addons

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
| Addons seen in | ActionFraction, AggroMeter, AutoBand, BuffHead, CM_ClosetGoblin, CastSequence, Dascore, DuelInvite |
| Files seen in | AggroMeter.lua, AutoBand.lua, Button.lua, ClosetGoblinCharacterWindow.lua, ClosetGoblinZoneWindow.lua, Code/CombatLog/CombatLogEpsWindow.lua, Code/CombatLog/CombatLogTargetDefenseWindow.lua, Code/Marks/Marks.lua |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | ActionFraction: RightClick, ActionFraction: SetFontSelectionMenu, ActionFraction: SetPresetLocation, AggroMeter: OnTabRBU, AggroMeter: PickedListMenu, AutoBand: AutoBand_L_AddContextMenuItem |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 390 |
| Global usage count | 390 |
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

Observed as a global function across 39 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: Checkbox(mode,MODE_ALL)..GetStringFromTable("UserSettingsStrings",StringTables.UserSettings.PERFORMANCE_ABILITY_EFFECTS_ALL)..L " (*)", Checkbox(mode,MODE_NONE)..GetStringFromTable("UserSettingsStrings",StringTables.UserSettings.PERFORMANCE_ABILITY_EFFECTS_NONE)..L " (0)", Checkbox(mode,MODE_PARTY)..GetStringFromTable("UserSettingsStrings",StringTables.UserSettings.PERFORMANCE_ABILITY_EFFECTS_PARTY)..L " (6)" |
| arg2 | Observed as a function or method reference. | Observed values: AbilitiesWindow.ToggleShowing, ActionFractionWindow.OnLock, ActionFractionWindow.OnUnlock |
| arg3 | Observed as a boolean toggle. | Observed values: (activeItem.Entry.Index==#activeMenu.Subitems), (activeItem.Entry.Index==1), (activeRowItem.Index==#textureSettings.Animation.Frames) |
| arg4 | Observed as a boolean toggle. | Observed values: EA_Window_ContextMenu.CONTEXT_MENU_1, true |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- ActionFraction
- AggroMeter
- AutoBand
- BuffHead
- CM_ClosetGoblin
- CastSequence
- Dascore
- DuelInvite
- DuffTimer
- EA_ScenarioGroupWindow
- EZCraftX
- Effigy
- Enemy
- HealGrid
- Killer
- LibAddonButton
- MapMonster
- MapPin
- MarkBuff
- Miracle Grow Remix
- Motion
- NerfedButtons
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
- TidyChat
- TidyQueue
- TurretRange
- WarBoard
- WarBoard_Loc
- XpStatus+G

## Examples

- ActionFraction: RightClick -> EA_Window_ContextMenu.AddMenuItem(GetString(StringTables.Default.LABEL_TO_LOCK), ActionFractionWindow.OnLock, not movable, true, EA_Window_ContextMenu.CONTEXT_MENU_1)
- ActionFraction: RightClick -> EA_Window_ContextMenu.AddMenuItem(GetString(StringTables.Default.LABEL_TO_UNLOCK), ActionFractionWindow.OnUnlock, movable, true, EA_Window_ContextMenu.CONTEXT_MENU_1)
- ActionFraction: RightClick -> EA_Window_ContextMenu.AddMenuItem(GetString(StringTables.Default.LABEL_CHAT_OPTIONS_RESET), ActionFractionWindow.ResetWindow, false, EA_Window_ContextMenu.CONTEXT_MENU_1)
- ActionFraction: SetFontSelectionMenu -> EA_Window_ContextMenu.AddMenuItem(LOC_TEXT.FONT_SIZE_SMALL, ActionFractionWindow.SetFontSmall, false, true, EA_Window_ContextMenu.CONTEXT_MENU_3)
- ActionFraction: SetFontSelectionMenu -> EA_Window_ContextMenu.AddMenuItem(LOC_TEXT.FONT_SIZE_MEDIUM, ActionFractionWindow.SetFontMedium, false, true, EA_Window_ContextMenu.CONTEXT_MENU_3)
- ActionFraction: SetFontSelectionMenu -> EA_Window_ContextMenu.AddMenuItem(LOC_TEXT.FONT_SIZE_LARGE, ActionFractionWindow.SetFontLarge, false, true, EA_Window_ContextMenu.CONTEXT_MENU_3)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event

## Used With

- [EA_Window_ContextMenu.AddCascadingMenuItem](global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddUserDefinedMenuItem](global_EA_Window_ContextMenu.AddUserDefinedMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Hide](global_EA_Window_ContextMenu.Hide.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
