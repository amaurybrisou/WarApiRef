# EA_FullResizeImage_TanBorder

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Busted, EA_UiDebugTools, EA_UiModWindow, Effigy, MapMonster, Miracle Grow Remix, Moth, RVAPI_ColorDialog |
| Files seen in | `/workspace_addons/Busted/Busted.xml:84`, `/workspace_addons/Effigy/Effigy.xml:51`, `/workspace_addons/MGRemix/layout.xml:863`, `/workspace_addons/MapMonster/Source/MapMonster_PinTypeEditorWindow.xml:226`, `/workspace_addons/MapMonster/Source/MapMonster_PinTypeEditorWindow.xml:365`, `/workspace_addons/Moth/Moth.xml:75`, `/workspace_addons/RVAPI_ColorDialog/RVAPI_ColorDialog.xml:116`, `/workspace_addons/RVMOD_Manager/RVMOD_ManagerTemplates.xml:170` |
| Namespaces detected | EA_FullResizeImage_TanBorder |
| Source kinds | xml_attributes |
| Example locations | BustedGUICountBackground, EA_ScrollWindow_ModInfoTemplateScrollChildStatusFrame, FrameTanBack, MapMonster_PinTypeEditorWindowMapPinIconOptionsBorder, MapMonster_PinTypeEditorWindowOptionsBorder, MiracleGrow2LayoutTemplateVBar |
| XML usage count | 9 |
| XML attribute usage count | 9 |
| Lua usage count | 0 |
| Global usage count | 0 |
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

Observed engine XML template or inherited constant referenced by 9 addons.

## Seen In

- Busted
- EA_UiDebugTools
- EA_UiModWindow
- Effigy
- MapMonster
- Miracle Grow Remix
- Moth
- RVAPI_ColorDialog
- RVMOD_Manager

## Used By

- BustedGUICountBackground
- EA_ScrollWindow_ModInfoTemplateScrollChildStatusFrame
- FrameTanBack
- MapMonster_PinTypeEditorWindowMapPinIconOptionsBorder
- MapMonster_PinTypeEditorWindowOptionsBorder
- MiracleGrow2LayoutTemplateVBar
- MothBorderTan
- RVAPI_ColorDialogWindowBackgroundFrame
- RVMOD_ManagerModInfoTemplateScrollChildStatusFrame

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
