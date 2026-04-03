# EA_Window_Macro

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
| Addons seen in | AdvancedPetAssist, Atlas, Brizio's Crappy Computer Medic, BuddyBind, Calling, Enemy, RandomMount, Squared |
| Files seen in | Code/Assist/Assist.lua, Code/Core/ConfigurationWindow.lua, Code/Core/Utils.lua, Shared/Utilities.lua |
| Namespaces detected | EA_Window_Macro |
| Source kinds | lua_calls |
| Example locations | AdvancedPetAssist: CreateOrUpdateMacros, Atlas: CreateMacro, Brizio's Crappy Computer Medic: SetMacro, BuddyBind: CreateMacro, Calling: CreateMacro, Enemy: ConfigurationWindow_OnMacroMouseDrag |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
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

Shared function table with 2 member functions; the primary API surface for 10 addons.

## Functions

- EA_Window_Macro.UpdateDetails
- EA_Window_Macro.UpdateMacros

## Observed Members

- none

## Seen In

- AdvancedPetAssist
- Atlas
- Brizio's Crappy Computer Medic
- BuddyBind
- Calling
- Enemy
- RandomMount
- Squared
- Swift Assist
- yAssistHelper

## Examples

- AdvancedPetAssist: CreateOrUpdateMacros -> EA_Window_Macro.UpdateMacros()
- Atlas: CreateMacro -> EA_Window_Macro.UpdateMacros()
- Brizio's Crappy Computer Medic: SetMacro -> EA_Window_Macro.UpdateDetails(slot)
- BuddyBind: CreateMacro -> EA_Window_Macro.UpdateMacros()
- Calling: CreateMacro -> EA_Window_Macro.UpdateMacros()
- Enemy: ConfigurationWindow_OnMacroMouseDrag -> EA_Window_Macro.UpdateDetails(p.macroId)

## Notes

- none
