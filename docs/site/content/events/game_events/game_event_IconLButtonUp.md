# IconLButtonUp

- Category: Game Event
- Confidence level: MEDIUM
- Confidence score: 43/100

## Confidence Assessment

- Level: MEDIUM

- Score: 43/100

- Rationale: Promoted as MEDIUM confidence because referenced by generated docs or reference files, called globally with no local definition, used in event registration or dispatch.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- -20 Only one weak usage site: Evidence is too shallow to trust as platform API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy |
| Files seen in | Code/Core/Events.lua |
| Namespaces detected | IconLButtonUp |
| Source kinds | event_page, lua_event_registration |
| Example locations | Enemy: AddEventHandler |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | no |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | yes |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Runtime event with 1 handler registrations observed across 1 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload not inferable from addon-level documentation alone.

## Seen In

- Enemy

## Registrars And Handlers

- AddEventHandler
- Enemy.IntercomUI_IntercomDialog_Open
- addon

## Examples

- Enemy: AddEventHandler -> IconLButtonUp -> Enemy.IntercomUI_IntercomDialog_Open
- Enemy: Enemy.IntercomUI_IntercomDialog_Open -> AddEventHandler(IconLButtonUp, Enemy.IntercomUI_IntercomDialog_Open)

## Notes

- Only one addon surfaced this event in the current addon-api corpus.
