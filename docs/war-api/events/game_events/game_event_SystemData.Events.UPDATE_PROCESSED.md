# SystemData.Events.UPDATE_PROCESSED

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 168

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | DammazKron, Squared, Statdoll Remix, Swift Assist |
| Files seen in | Core/ToolTip/DK_Tooltip.lua, Squared.lua, StatdollGetstats.lua, StatdollLocal.lua, SwiftAssist.lua |
| Namespaces detected | SystemData |
| Source kinds | event_page, lua_event_registration |
| Example locations | DammazKron: RegisterEventHandler, Squared: RegisterEventHandler, Statdoll Remix: registerEventHandler, Swift Assist: RegisterEventHandler |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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

Runtime event with 2 handler registrations observed across 4 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from contract artifacts alone; treat this as an engine event identifier.

## Seen In

- DammazKron
- Squared
- Statdoll Remix
- Swift Assist

## Registrars And Handlers

- RegisterEventHandler
- Statdoll.Local.delayedUpdate
- TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH
- global
- registerEventHandler

## Examples

- DammazKron: RegisterEventHandler -> SystemData.Events.UPDATE_PROCESSED -> TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH
- Squared: RegisterEventHandler -> SystemData.Events.UPDATE_PROCESSED -> TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH
- Statdoll Remix: registerEventHandler -> SystemData.Events.UPDATE_PROCESSED -> Statdoll.Local.delayedUpdate
- Swift Assist: RegisterEventHandler -> SystemData.Events.UPDATE_PROCESSED -> TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH
- DammazKron: TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH -> RegisterEventHandler(SystemData.Events.UPDATE_PROCESSED, TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH)
- Squared: TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH -> RegisterEventHandler(SystemData.Events.UPDATE_PROCESSED, TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH)

## Related APIs

- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Notes

- none
