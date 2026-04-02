# EA_Window_Macro

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 140

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, Enemy, RandomMount, Swift Assist |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APACommands.lua:61`, `/workspace_addons/RandomMount/RandomMount.lua:79`, `/workspace_addons/swift-assist/SwiftAssist.lua:247` |
| Namespaces detected | EA_Window_Macro |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.CreateOrUpdateMacros, RandomMount: RandomMount.CreateMacro, Swift Assist: CreateMacro, Swift Assist: Swift Assist.local.CreateMacro |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 1 |
| Local definition count | 1 |
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

Observed shared global table or namespace surfaced in 4 addons.

## Functions

- EA_Window_Macro.UpdateMacros

## Observed Members

- none

## Seen In

- AdvancedPetAssist
- Enemy
- RandomMount
- Swift Assist

## Examples

- AdvancedPetAssist: AdvancedPetAssist.CreateOrUpdateMacros -> EA_Window_Macro.UpdateMacros()
- RandomMount: RandomMount.CreateMacro -> EA_Window_Macro.UpdateMacros()
- Swift Assist: CreateMacro -> EA_Window_Macro.UpdateMacros()
- Swift Assist: Swift Assist.local.CreateMacro -> EA_Window_Macro.UpdateMacros()

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
