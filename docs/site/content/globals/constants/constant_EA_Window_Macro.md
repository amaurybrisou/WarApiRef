# EA_Window_Macro

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
| Addons seen in | bigger_MacroWindow |
| Files seen in | `/workspace/data/raw/bigger_macrowindow/Source/MacroWindow.xml:0` |
| Namespaces detected | EA_Window_Macro |
| Source kinds | flows, xml_attributes |
| Example locations | EA_Window_MacroBackground, EA_Window_MacroClose, EA_Window_MacroDetails, EA_Window_MacroMacros, EA_Window_MacroMacrosBackground, EA_Window_MacroTitleBar |
| XML usage count | 6 |
| XML attribute usage count | 6 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 9 |
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

- bigger_MacroWindow

## Used By

- EA_Window_MacroBackground
- EA_Window_MacroClose
- EA_Window_MacroDetails
- EA_Window_MacroMacros
- EA_Window_MacroMacrosBackground
- EA_Window_MacroTitleBar

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
