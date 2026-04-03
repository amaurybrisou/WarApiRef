# EA_EditBox_NoFrame

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | EA_UiDebugTools, SNT_PANEL, zMailMod |
| Files seen in | Source/DebugWindow.xml, snt_panel.xml, zMailModSend.xml |
| Namespaces detected | EA_EditBox_NoFrame |
| Source kinds | xml_attributes |
| Example locations | DevPadCopyLog, SNT_PANEL_SETUP_WINDOW_T1, SNT_PANEL_SETUP_WINDOW_T2, SNT_PANEL_SETUP_WINDOW_T3, SNT_PANEL_SETUP_WINDOW_T4, SNT_PANEL_SETUP_WINDOW_T5 |
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

Engine-supplied XML constant or template class referenced by 3 addons.

## Seen In

- EA_UiDebugTools
- SNT_PANEL
- zMailMod

## Used By

- DevPadCopyLog
- SNT_PANEL_SETUP_WINDOW_T1
- SNT_PANEL_SETUP_WINDOW_T2
- SNT_PANEL_SETUP_WINDOW_T3
- SNT_PANEL_SETUP_WINDOW_T4
- SNT_PANEL_SETUP_WINDOW_T5
- zMailModSendSubjectEditBox
- zMailModSendToAutoCompleteNew
- zMailModSendToEditBox

## Related APIs

- [EditBox](../../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type

## Notes

- none
