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
| Addons seen in | Busted, CCTV, ChattyCathy, DuffTimer, EA_UiDebugTools, EA_UiModWindow, Effigy, MapMonster |
| Files seen in | Busted.xml, CCTV.xml, ChattyCathy.xml, DuffTimerOptionsDefn.xml, Effigy.xml, Moth.xml, RVAPI_ColorDialog.xml, RVMOD_ManagerTemplates.xml |
| Namespaces detected | EA_FullResizeImage_TanBorder |
| Source kinds | xml_attributes |
| Example locations | BustedGUICountBackground, CCTV_BG_TEMPLATEBorderTan, ChattyCathyOptEntrySetupBorder, ChattyCathyOptScrollSetupBorder, DuffTimerOptions_ColorSection, EA_ScrollWindow_ModInfoTemplateScrollChildStatusFrame |
| XML usage count | 17 |
| XML attribute usage count | 17 |
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

Engine-supplied XML constant or template class referenced by 14 addons.

## Seen In

- Busted
- CCTV
- ChattyCathy
- DuffTimer
- EA_UiDebugTools
- EA_UiModWindow
- Effigy
- MapMonster
- Miracle Grow Remix
- Moth
- RVAPI_ColorDialog
- RVMOD_Manager
- SNT_PANEL
- WarBoard_FPS

## Used By

- BustedGUICountBackground
- CCTV_BG_TEMPLATEBorderTan
- ChattyCathyOptEntrySetupBorder
- ChattyCathyOptScrollSetupBorder
- DuffTimerOptions_ColorSection
- EA_ScrollWindow_ModInfoTemplateScrollChildStatusFrame
- FrameTanBack
- MapMonster_PinTypeEditorWindowMapPinIconOptionsBorder
- MapMonster_PinTypeEditorWindowOptionsBorder
- MiracleGrow2LayoutTemplateVBar
- MothBorderTan
- RVAPI_ColorDialogWindowBackgroundFrame
- RVMOD_ManagerModInfoTemplateScrollChildStatusFrame
- WarBoard_FPSOptions_Tan1
- WarBoard_FPSOptions_Tan2
- WarBoard_FPSOptions_Tan3
- snt_panel_template_layer2

## Related APIs

- [FullResizeImage](../../xml/element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type

## Notes

- none
