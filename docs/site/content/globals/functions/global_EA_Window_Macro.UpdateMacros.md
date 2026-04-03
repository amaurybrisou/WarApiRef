# EA_Window_Macro.UpdateMacros

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 8 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, Atlas, BuddyBind, Calling, RandomMount, Squared, Swift Assist, yAssistHelper |
| Files seen in | APACommands.lua, BuddyBind.lua, MacroUtilites.lua, RandomMount.lua, Shared/Utilities.lua, SquaredClickAssist.lua, SwiftAssist.lua, yAssist.lua |
| Namespaces detected | EA_Window_Macro |
| Source kinds | lua_calls |
| Example locations | AdvancedPetAssist: CreateOrUpdateMacros, Atlas: CreateMacro, BuddyBind: CreateMacro, Calling: CreateMacro, RandomMount: CreateMacro, Squared: CreateMacro |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
EA_Window_Macro.UpdateMacros()
```

## Description

Observed as a global function across 8 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedPetAssist
- Atlas
- BuddyBind
- Calling
- RandomMount
- Squared
- Swift Assist
- yAssistHelper

## Examples

- AdvancedPetAssist: CreateOrUpdateMacros -> EA_Window_Macro.UpdateMacros()
- Atlas: CreateMacro -> EA_Window_Macro.UpdateMacros()
- BuddyBind: CreateMacro -> EA_Window_Macro.UpdateMacros()
- Calling: CreateMacro -> EA_Window_Macro.UpdateMacros()
- RandomMount: CreateMacro -> EA_Window_Macro.UpdateMacros()
- Squared: CreateMacro -> EA_Window_Macro.UpdateMacros()

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
