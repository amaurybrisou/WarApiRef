# wstring

- Category: Global Table
- Confidence level: MEDIUM
- Confidence score: 45/100

## Confidence Assessment

- Level: MEDIUM

- Score: 45/100

- Rationale: Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Clock |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | Clock: OnUpdate |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
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

Shared function table with 1 member functions; the primary API surface for 1 addons.

## Functions

- wstring.format

## Observed Members

- none

## Seen In

- Clock

## Examples

- Clock: OnUpdate -> wstring.format(fmt, 12, Clock.min, Clock.sec)
- Clock: OnUpdate -> wstring.format(fmt, Clock.hour%ClockSettings.Hours, Clock.min, Clock.sec)

## Notes

- none
