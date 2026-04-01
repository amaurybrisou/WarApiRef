# PetWindow.HasPet

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 96/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Score: 96/100

- Rationale: Promoted as HIGH confidence because referenced by generated docs or reference files, called globally with no local definition, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, Aura, DaemonAssist |
| Files seen in | `/workspace/AdvancedPetAssist/APACore.lua:214`, `/workspace/Aura/Source/AuraEngine.lua:1142`, `/workspace/DaemonAssist/DACore.lua:7` |
| Namespaces detected | PetWindow |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.HasPet, Aura: AuraEngine.HandleTriggerType_PetStatus, DaemonAssist: DaemonAssist.HasPet |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
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
PetWindow.HasPet()
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AdvancedPetAssist
- Aura
- DaemonAssist

## Examples

- AdvancedPetAssist: AdvancedPetAssist.HasPet -> PetWindow.HasPet()
- Aura: AuraEngine.HandleTriggerType_PetStatus -> PetWindow.HasPet()
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

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
