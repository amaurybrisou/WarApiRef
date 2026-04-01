# PetWindow.SwitchToPassiveStance

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 88/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because referenced by generated docs or reference files, called globally with no local definition, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, DaemonAssist |
| Files seen in | `/workspace/AdvancedPetAssist/APACore.lua:269`, `/workspace/AdvancedPetAssist/APACore.lua:283`, `/workspace/DaemonAssist/DACore.lua:23` |
| Namespaces detected | PetWindow |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.ApplyDefaultStance, AdvancedPetAssist: AdvancedPetAssist.SetPassiveFollow, DaemonAssist: DaemonAssist.EnsurePassiveFollow |
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
PetWindow.SwitchToPassiveStance()
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AdvancedPetAssist
- DaemonAssist

## Examples

- AdvancedPetAssist: AdvancedPetAssist.ApplyDefaultStance -> PetWindow.SwitchToPassiveStance()
- AdvancedPetAssist: AdvancedPetAssist.SetPassiveFollow -> PetWindow.SwitchToPassiveStance()
- DaemonAssist: DaemonAssist.EnsurePassiveFollow -> PetWindow.SwitchToPassiveStance()

## Related APIs

- none

## Used With

- [GameData.PetCommand.FOLLOW](../../gamedata/fields/gamedata_GameData.PetCommand.FOLLOW.md) (HIGH 100/100) - GameData Field

## Triggered By

- none

## Affects

- [GameData.PetCommand.DEFENSIVE](../../gamedata/fields/gamedata_GameData.PetCommand.DEFENSIVE.md) (HIGH 100/100) - GameData Field
- [GameData.PetCommand.FOLLOW](../../gamedata/fields/gamedata_GameData.PetCommand.FOLLOW.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
