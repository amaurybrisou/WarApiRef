# OnMButtonDown

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 106

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, MoraleCircle |
| Files seen in | `/workspace_addons/Enemy/Code/UnitFrames/UnitFrame.xml:90`, `/workspace_addons/MoraleCircle/MoraleCircle.xml:18` |
| Namespaces detected | OnMButtonDown |
| Source kinds | event_page, xml_handlers |
| Example locations | Enemy: EnemyUnitFrame.OnMButtonDown, MoraleCircle: MoraleTemplate.OnMButtonDown |
| XML usage count | 2 |
| XML attribute usage count | 2 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
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

Observed as an engine-supplied UI event hook used by 2 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- Enemy
- MoraleCircle

## Registrars And Handlers

- Enemy.UnitFramesUI_UnitFrame_OnMButtonDown
- MoraleCircle.Reset

## Examples

- Enemy: EnemyUnitFrame -> EnemyUnitFrame.OnMButtonDown -> Enemy.UnitFramesUI_UnitFrame_OnMButtonDown
- MoraleCircle: MoraleTemplate -> MoraleTemplate.OnMButtonDown -> MoraleCircle.Reset

## Related APIs

- none

## Used With

- [OnMButtonDown](../../xml/handlers/handler_OnMButtonDown.md) (HIGH 100/100) - XML Handler
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- none
