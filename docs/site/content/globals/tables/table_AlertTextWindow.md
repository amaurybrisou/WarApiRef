# AlertTextWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 113

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Aura, PlanB, WSCT |
| Files seen in | `/workspace/data/raw/Aura/Source/Aura.lua:186`, `/workspace/data/raw/Aura/Source/Aura.lua:282`, `/workspace/data/raw/PlanB/PlanB.lua:28`, `/workspace/data/raw/wsct/wsct.lua:590` |
| Namespaces detected | AlertTextWindow |
| Source kinds | lua_calls |
| Example locations | Aura: Aura:Activate, Aura: Aura:Deactivate, PlanB: PlanB.local.display_alert, PlanB: display_alert, WSCT: WSCT:DisplayMessage |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 0 |
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

Observed shared global table or namespace surfaced in 3 addons.

## Functions

- AlertTextWindow.AddAlert
- AlertTextWindow.AddLine

## Observed Members

- none

## Seen In

- Aura
- PlanB
- WSCT

## Examples

- Aura: Aura:Activate -> AlertTextWindow.AddLine(AuraConstants.AlertText[self:Get("activation-alerttexttype")].type, self:Get("activation-alerttext"))
- Aura: Aura:Deactivate -> AlertTextWindow.AddLine(AuraConstants.AlertText[self:Get("deactivation-alerttexttype")].type, self:Get("deactivation-alerttext"))
- PlanB: PlanB.local.display_alert -> AlertTextWindow.AddAlert()
- PlanB: display_alert -> AlertTextWindow.AddAlert()
- WSCT: WSCT:DisplayMessage -> AlertTextWindow.AddLine(SystemData.AlertText.Types.DEFAULT, msg)

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
