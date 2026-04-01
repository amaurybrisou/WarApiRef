# PetWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 130

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, Aura, DaemonAssist, TexturedButtons |
| Files seen in | `/workspace/AdvancedPetAssist/APACore.lua:214`, `/workspace/AdvancedPetAssist/APACore.lua:269`, `/workspace/AdvancedPetAssist/APACore.lua:283`, `/workspace/Aura/Source/AuraEngine.lua:1142`, `/workspace/DaemonAssist/DACore.lua:23`, `/workspace/DaemonAssist/DACore.lua:7` |
| Namespaces detected | PetWindow |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.ApplyDefaultStance, AdvancedPetAssist: AdvancedPetAssist.HasPet, AdvancedPetAssist: AdvancedPetAssist.SetPassiveFollow, Aura: AuraEngine.HandleTriggerType_PetStatus, DaemonAssist: DaemonAssist.EnsurePassiveFollow, DaemonAssist: DaemonAssist.HasPet |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 2 |
| Local definition count | 2 |
| Documentation references | 1 |
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

Observed shared global table or namespace surfaced in 4 addons.

## Functions

- PetWindow.HasPet
- PetWindow.SwitchToPassiveStance

## Observed Members

- none

## Seen In

- AdvancedPetAssist
- Aura
- DaemonAssist
- TexturedButtons

## Examples

- AdvancedPetAssist: AdvancedPetAssist.ApplyDefaultStance -> PetWindow.SwitchToPassiveStance()
- AdvancedPetAssist: AdvancedPetAssist.HasPet -> PetWindow.HasPet()
- AdvancedPetAssist: AdvancedPetAssist.SetPassiveFollow -> PetWindow.SwitchToPassiveStance()
- Aura: AuraEngine.HandleTriggerType_PetStatus -> PetWindow.HasPet()
- DaemonAssist: DaemonAssist.EnsurePassiveFollow -> PetWindow.SwitchToPassiveStance()
- DaemonAssist: DaemonAssist.HasPet -> PetWindow.HasPet()

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
