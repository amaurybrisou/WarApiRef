# PlayerMenuWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 113

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy, Enemy |
| Files seen in | `/workspace_addons/Effigy/Effigy.lua:474`, `/workspace_addons/Enemy/Code/UnitFrames/UnitFrame.lua:413` |
| Namespaces detected | PlayerMenuWindow |
| Source kinds | lua_calls |
| Example locations | Effigy: Effigy.local.FriendlyRightClickMenu, Effigy: FriendlyRightClickMenu, Enemy: Enemy.UnitFramesUI_UnitFrame_OnRButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 2 |
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

Observed shared global table or namespace surfaced in 2 addons.

## Functions

- PlayerMenuWindow.NewCustomItem
- PlayerMenuWindow.ShowMenu

## Observed Members

- none

## Seen In

- Effigy
- Enemy

## Examples

- Effigy: Effigy.local.FriendlyRightClickMenu -> PlayerMenuWindow.ShowMenu(TargetInfo:UnitName("selffriendlytarget"), TargetInfo:UnitEntityId("selffriendlytarget"), nil)
- Effigy: FriendlyRightClickMenu -> PlayerMenuWindow.ShowMenu(TargetInfo:UnitName("selffriendlytarget"), TargetInfo:UnitEntityId("selffriendlytarget"), nil)
- Enemy: Enemy.UnitFramesUI_UnitFrame_OnRButtonUp -> PlayerMenuWindow.NewCustomItem(GetString(StringTables.Default.LABEL_GROUP_OPTIONS), EA_Window_OpenParty.OpenToManageTab, false)
- Enemy: Enemy.UnitFramesUI_UnitFrame_OnRButtonUp -> PlayerMenuWindow.NewCustomItem(L "Unit frames options", function()Enemy.UI_ConfigDialog_Open("UnitFrames")end, false)
- Enemy: Enemy.UnitFramesUI_UnitFrame_OnRButtonUp -> PlayerMenuWindow.NewCustomItem(L "Group icons options", function()Enemy.UI_ConfigDialog_Open("GroupIcons")end, false)
- Enemy: Enemy.UnitFramesUI_UnitFrame_OnRButtonUp -> PlayerMenuWindow.NewCustomItem(GetString(StringTables.Default.LABEL_PARTY_FORM_WARPARTY), GroupWindow.OnFormWarparty, Enemy.groups.groupType==Enemy.GroupTypes.Warband)

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
