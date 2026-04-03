# GroupWindow

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
| Addons seen in | CleanUnitFrames, EA_ScenarioGroupWindow, Effigy, HealGrid, NerfedButtons, NoOverheal, Phantom, Pure |
| Files seen in | PhantomLib/PhantomLib.lua, Source/PureGroup.lua, Source/ScenarioGroupWindow.lua, Units/HealGridUnitGroupMember.lua |
| Namespaces detected | GroupWindow |
| Source kinds | lua_calls |
| Example locations | CleanUnitFrames: OnRButtonUp, EA_ScenarioGroupWindow: ToggleMainGroupVisibility, Effigy: RButtonUp, HealGrid: EnforceWindowStates, HealGrid: ShowDropdownMenu, NerfedButtons: inMyParty |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 8 |
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

Shared function table with 8 member functions; the primary API surface for 10 addons.

## Functions

- GroupWindow.ConditionalShow
- GroupWindow.GetGroupMember
- GroupWindow.GetMainAssist
- GroupWindow.IsMemberValid
- GroupWindow.IsPlayerInGroup
- GroupWindow.OnEffectsUpdated
- GroupWindow.OnViewGroupOptions
- GroupWindow.ShowMenu

## Observed Members

- none

## Seen In

- CleanUnitFrames
- EA_ScenarioGroupWindow
- Effigy
- HealGrid
- NerfedButtons
- NoOverheal
- Phantom
- Pure
- Refer
- Squared

## Examples

- CleanUnitFrames: OnRButtonUp -> GroupWindow.ShowMenu(CleanGroupMemberName)
- EA_ScenarioGroupWindow: ToggleMainGroupVisibility -> GroupWindow.ConditionalShow()
- Effigy: RButtonUp -> GroupWindow.ShowMenu(member_name)
- HealGrid: EnforceWindowStates -> GroupWindow.OnEffectsUpdated(i, GetBuffs(i), true)
- HealGrid: ShowDropdownMenu -> GroupWindow.ShowMenu(self:GetUnitName())
- NerfedButtons: inMyParty -> GroupWindow.IsPlayerInGroup(targetName)

## Notes

- none
