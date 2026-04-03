# EA_Window_QuestTracker

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 80/100

## Confidence Assessment

- Level: HIGH

- Score: 80/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Kwestor |
| Files seen in | Tracker/KwestorTracker.lua |
| Namespaces detected | EA_Window_QuestTracker |
| Source kinds | lua_calls |
| Example locations | Kwestor: HideWARQuestTracker, Kwestor: NubLButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 1 |
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

Shared function table with 1 member functions; the primary API surface for 1 addons.

## Functions

- EA_Window_QuestTracker.ToggleShowing

## Observed Members

- none

## Seen In

- Kwestor

## Examples

- Kwestor: HideWARQuestTracker -> EA_Window_QuestTracker.ToggleShowing()
- Kwestor: NubLButtonUp -> EA_Window_QuestTracker.ToggleShowing()

## Notes

- none
