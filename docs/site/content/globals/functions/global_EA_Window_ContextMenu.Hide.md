# EA_Window_ContextMenu.Hide

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
| Addons seen in | BuffHead, TexturedButtons, TurretRange |
| Files seen in | `/workspace/data/raw/BuffHead/Setup/SelectFont.lua:62`, `/workspace/data/raw/TexturedButtons/Setup/SelectFont.lua:64`, `/workspace/data/raw/TurrentRange/Setup/SetupDisplay.lua:366` |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | BuffHead: BuffHead.Setup.SelectFont.SetFont, TexturedButtons: TexturedButtons.Setup.SelectFont.SetFont, TurretRange: TurretRange.Setup.Distance.SetFont |
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
EA_Window_ContextMenu.Hide(arg1)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: menu |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- BuffHead
- TexturedButtons
- TurretRange

## Examples

- BuffHead: BuffHead.Setup.SelectFont.SetFont -> EA_Window_ContextMenu.Hide(menu)
- TexturedButtons: TexturedButtons.Setup.SelectFont.SetFont -> EA_Window_ContextMenu.Hide(menu)
- TurretRange: TurretRange.Setup.Distance.SetFont -> EA_Window_ContextMenu.Hide(menu)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
