# DefaultColor.LabelSetTextColor

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 55/100
- Seen in: 1 addons

## Confidence Assessment

- Level: MEDIUM

- Score: 55/100

- Rationale: Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, argument pattern is consistent.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | EA_OpenPartyWindow |
| Files seen in | source/openpartywindowtablootrolloptions.lua |
| Namespaces detected | DefaultColor |
| Source kinds | lua_calls |
| Example locations | EA_OpenPartyWindow: Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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
DefaultColor.LabelSetTextColor(arg1, arg2)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: PARENT_WINDOW.."RarityLabelArtifact", PARENT_WINDOW.."RarityLabelCommon", PARENT_WINDOW.."RarityLabelRare" |
| arg2 | Observed as a function or method reference. | Observed values: GameDefs.ItemRarity[SystemData.ItemRarity.ARTIFACT].color, GameDefs.ItemRarity[SystemData.ItemRarity.COMMON].color, GameDefs.ItemRarity[SystemData.ItemRarity.RARE].color |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- EA_OpenPartyWindow

## Examples

- EA_OpenPartyWindow: Initialize -> DefaultColor.LabelSetTextColor(PARENT_WINDOW.."RarityLabelCommon", GameDefs.ItemRarity[SystemData.ItemRarity.COMMON].color)
- EA_OpenPartyWindow: Initialize -> DefaultColor.LabelSetTextColor(PARENT_WINDOW.."RarityLabelUncommon", GameDefs.ItemRarity[SystemData.ItemRarity.UNCOMMON].color)
- EA_OpenPartyWindow: Initialize -> DefaultColor.LabelSetTextColor(PARENT_WINDOW.."RarityLabelRare", GameDefs.ItemRarity[SystemData.ItemRarity.RARE].color)
- EA_OpenPartyWindow: Initialize -> DefaultColor.LabelSetTextColor(PARENT_WINDOW.."RarityLabelVeryRare", GameDefs.ItemRarity[SystemData.ItemRarity.VERY_RARE].color)
- EA_OpenPartyWindow: Initialize -> DefaultColor.LabelSetTextColor(PARENT_WINDOW.."RarityLabelArtifact", GameDefs.ItemRarity[SystemData.ItemRarity.ARTIFACT].color)
- EA_OpenPartyWindow: Initialize -> DefaultColor.LabelSetTextColor(PARENT_WINDOW.."TrashTitle", GameDefs.ItemRarity[SystemData.ItemRarity.UTILITY].color)

## Affects

- [SystemData.Events.BATTLEGROUP_MEMBER_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.BATTLEGROUP_MEMBER_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.BATTLEGROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.BATTLEGROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_SETTINGS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_SETTINGS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ZONE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ZONE_CHANGED.md) (HIGH 100/100) - SystemData Field

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
