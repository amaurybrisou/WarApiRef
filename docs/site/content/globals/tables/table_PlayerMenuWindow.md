# PlayerMenuWindow

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
| Addons seen in | CleanUnitFrames, DuelInvite, Effigy, Enemy, HealGrid, Pure, Refer, Squared |
| Files seen in | Code/UnitFrames/UnitFrame.lua, Source/PureGroup.lua, Source/PureTarget.lua, Units/HealGridUnitFriendlyTarget.lua |
| Namespaces detected | PlayerMenuWindow |
| Source kinds | lua_calls |
| Example locations | CleanUnitFrames: OnRButtonUp, DuelInvite: OnDuel, Effigy: FriendlyRightClickMenu, Enemy: UnitFramesUI_UnitFrame_OnRButtonUp, HealGrid: ShowDropdownMenu, Pure: OnRButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 38 |
| Global usage count | 3 |
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

Shared function table with 3 member functions; the primary API surface for 9 addons.

## Functions

- PlayerMenuWindow.Done
- PlayerMenuWindow.NewCustomItem
- PlayerMenuWindow.ShowMenu

## Observed Members

- none

## Seen In

- CleanUnitFrames
- DuelInvite
- Effigy
- Enemy
- HealGrid
- Pure
- Refer
- Squared
- xHUD

## Examples

- CleanUnitFrames: OnRButtonUp -> PlayerMenuWindow.ShowMenu(targetName)
- DuelInvite: OnDuel -> PlayerMenuWindow.Done()
- Effigy: FriendlyRightClickMenu -> PlayerMenuWindow.ShowMenu(TargetInfo:UnitName("selffriendlytarget"), TargetInfo:UnitEntityId("selffriendlytarget"), nil)
- Enemy: UnitFramesUI_UnitFrame_OnRButtonUp -> PlayerMenuWindow.NewCustomItem(GetString(StringTables.Default.LABEL_GROUP_OPTIONS), EA_Window_OpenParty.OpenToManageTab, false)
- Enemy: UnitFramesUI_UnitFrame_OnRButtonUp -> PlayerMenuWindow.NewCustomItem(L "Unit frames options", function()Enemy.UI_ConfigDialog_Open("UnitFrames")end, false)
- Enemy: UnitFramesUI_UnitFrame_OnRButtonUp -> PlayerMenuWindow.NewCustomItem(L "Group icons options", function()Enemy.UI_ConfigDialog_Open("GroupIcons")end, false)

## Notes

- none
