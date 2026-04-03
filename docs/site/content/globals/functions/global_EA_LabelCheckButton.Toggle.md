# EA_LabelCheckButton.Toggle

- Category: Global Function
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
| Addons seen in | CaVES, ReliquaryHunter, nLootLink |
| Files seen in | source/CaVES.lua, source/ReliquaryHunter.lua, source/nLootLinkOptions.lua |
| Namespaces detected | EA_LabelCheckButton |
| Source kinds | lua_calls |
| Example locations | CaVES: OnToolTipOption1Changed, CaVES: ToggleIgnoreEventSlot, CaVES: ToggleIgnoreTalismans, ReliquaryHunter: OnEnableWorldMapWindowBackgroundOpacityChanged, ReliquaryHunter: OnEnableWorldMapWindowOpacityChanged, nLootLink: onToggleBagButton |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
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
EA_LabelCheckButton.Toggle()
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

- CaVES
- ReliquaryHunter
- nLootLink

## Examples

- CaVES: OnToolTipOption1Changed -> EA_LabelCheckButton.Toggle()
- CaVES: ToggleIgnoreEventSlot -> EA_LabelCheckButton.Toggle()
- CaVES: ToggleIgnoreTalismans -> EA_LabelCheckButton.Toggle()
- ReliquaryHunter: OnEnableWorldMapWindowBackgroundOpacityChanged -> EA_LabelCheckButton.Toggle()
- ReliquaryHunter: OnEnableWorldMapWindowOpacityChanged -> EA_LabelCheckButton.Toggle()
- nLootLink: onToggleBagButton -> EA_LabelCheckButton.Toggle()

## Used With

- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
