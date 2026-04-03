# ComboBoxExternalOpenMenu

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | RVMOD_Targets, Shinies, TokenMachine |
| Files seen in | Modules/UI/Shinies-UI-Browse.lua, RVMOD_Targets.lua, TokenMachine.lua |
| Namespaces detected | ComboBoxExternalOpenMenu |
| Source kinds | lua_calls |
| Example locations | RVMOD_Targets: ReopenComboBox, Shinies: Criteria_ReopenComboBox, TokenMachine: OpenOption |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
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
ComboBoxExternalOpenMenu(arg1)
```

## Description

Observed as a window function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: LastComboBoxName, lastComboSelected, setting.."Choice" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- RVMOD_Targets
- Shinies
- TokenMachine

## Examples

- RVMOD_Targets: ReopenComboBox -> ComboBoxExternalOpenMenu(LastComboBoxName)
- Shinies: Criteria_ReopenComboBox -> ComboBoxExternalOpenMenu(lastComboSelected)
- TokenMachine: OpenOption -> ComboBoxExternalOpenMenu(setting.."Choice")

## Affects

- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- none
