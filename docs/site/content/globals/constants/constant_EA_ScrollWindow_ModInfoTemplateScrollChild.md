# EA_ScrollWindow_ModInfoTemplateScrollChild

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 125

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | EA_UiModWindow |
| Files seen in | `/workspace/ea_uimodwindow/Source/UiModInfoTemplate.xml:134`, `/workspace/ea_uimodwindow/Source/UiModInfoTemplate.xml:141`, `/workspace/ea_uimodwindow/Source/UiModInfoTemplate.xml:150`, `/workspace/ea_uimodwindow/Source/UiModInfoTemplate.xml:157`, `/workspace/ea_uimodwindow/Source/UiModInfoTemplate.xml:166`, `/workspace/ea_uimodwindow/Source/UiModInfoTemplate.xml:173`, `/workspace/ea_uimodwindow/Source/UiModInfoTemplate.xml:183`, `/workspace/ea_uimodwindow/Source/UiModInfoTemplate.xml:190` |
| Namespaces detected | EA_ScrollWindow_ModInfoTemplateScrollChild |
| Source kinds | flows, xml_attributes |
| Example locations | EA_ScrollWindow_ModInfoTemplateScrollChildAuthorLabel, EA_ScrollWindow_ModInfoTemplateScrollChildAuthorText, EA_ScrollWindow_ModInfoTemplateScrollChildCareers, EA_ScrollWindow_ModInfoTemplateScrollChildCareersLabel, EA_ScrollWindow_ModInfoTemplateScrollChildCareersText, EA_ScrollWindow_ModInfoTemplateScrollChildCategoriesLabel |
| XML usage count | 16 |
| XML attribute usage count | 16 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 1 |
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

Observed engine XML template or inherited constant referenced by 1 addons.

## Seen In

- EA_UiModWindow

## Used By

- EA_ScrollWindow_ModInfoTemplateScrollChildAuthorLabel
- EA_ScrollWindow_ModInfoTemplateScrollChildAuthorText
- EA_ScrollWindow_ModInfoTemplateScrollChildCareers
- EA_ScrollWindow_ModInfoTemplateScrollChildCareersLabel
- EA_ScrollWindow_ModInfoTemplateScrollChildCareersText
- EA_ScrollWindow_ModInfoTemplateScrollChildCategoriesLabel
- EA_ScrollWindow_ModInfoTemplateScrollChildCategoriesText
- EA_ScrollWindow_ModInfoTemplateScrollChildDescriptionLabel
- EA_ScrollWindow_ModInfoTemplateScrollChildDescriptionText
- EA_ScrollWindow_ModInfoTemplateScrollChildGameVersionLabel
- EA_ScrollWindow_ModInfoTemplateScrollChildGameVersionText
- EA_ScrollWindow_ModInfoTemplateScrollChildName
- EA_ScrollWindow_ModInfoTemplateScrollChildNameBackground
- EA_ScrollWindow_ModInfoTemplateScrollChildStatus
- EA_ScrollWindow_ModInfoTemplateScrollChildVersionLabel
- EA_ScrollWindow_ModInfoTemplateScrollChildVersionText

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
