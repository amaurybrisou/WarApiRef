# BankWindow.Show

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 88/100
- Seen in: 3 addons

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
| Addons seen in | AnywhereTrainer, BagOMatic, BankArkel |
| Files seen in | `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.lua:271`, `/workspace/data/raw/BankArkel/BankArkel.lua:563`, `/workspace/data/raw/bagomatic/BagOMatic.lua:287` |
| Namespaces detected | BankWindow |
| Source kinds | globals, lua_calls |
| Example locations | AnywhereTrainer: AnywhereTrainer.OnLeftClickBank, BagOMatic: BagOMatic.cmd_showbank, BankArkel: BankArkel.ShowBank |
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
BankWindow.Show()
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

- AnywhereTrainer
- BagOMatic
- BankArkel

## Examples

- AnywhereTrainer: AnywhereTrainer.OnLeftClickBank -> BankWindow.Show()
- BagOMatic: BagOMatic.cmd_showbank -> BankWindow.Show()
- BankArkel: BankArkel.ShowBank -> BankWindow.Show()

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
- Advanced return analysis: No strong return evidence observed
