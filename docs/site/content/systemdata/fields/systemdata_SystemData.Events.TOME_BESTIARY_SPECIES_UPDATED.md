# SystemData.Events.TOME_BESTIARY_SPECIES_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | KillTracker |
| Files seen in | Source/KillTracker.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | OnBestiarySpeciesKillCountUpdated, OnBestiarySpeciesUpdated, lua_call |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | no |
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

SystemData.SystemData.Events.TOME_BESTIARY_SPECIES_UPDATED field accessed by 1 addons; commonly found in OnBestiarySpeciesKillCountUpdated and OnBestiarySpeciesUpdated, lua_call contexts.

## Seen In

- KillTracker

## Notes

- Observed in contexts: OnBestiarySpeciesKillCountUpdated, OnBestiarySpeciesUpdated, lua_call
