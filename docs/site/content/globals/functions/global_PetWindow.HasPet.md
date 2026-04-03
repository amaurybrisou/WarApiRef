# PetWindow.HasPet

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 71/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Score: 71/100

- Rationale: Promoted as HIGH confidence because called globally with no local definition, seen in 2 to 3 addons, role is consistent across addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, Aura, DaemonAssist |
| Files seen in | APACore.lua, DACore.lua, Source/AuraEngine.lua |
| Namespaces detected | PetWindow |
| Source kinds | lua_calls |
| Example locations | AdvancedPetAssist: HasPet, Aura: HandleTriggerType_PetStatus, DaemonAssist: HasPet |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 0 |
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

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedPetAssist
- Aura
- DaemonAssist

## Examples

- AdvancedPetAssist: HasPet -> PetWindow.HasPet()
- Aura: HandleTriggerType_PetStatus -> PetWindow.HasPet()
- DaemonAssist: HasPet -> PetWindow.HasPet()

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
