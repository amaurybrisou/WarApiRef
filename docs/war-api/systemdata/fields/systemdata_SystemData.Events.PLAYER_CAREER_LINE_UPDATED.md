# SystemData.Events.PLAYER_CAREER_LINE_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 8 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 188

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, FixGit, GuardBot, NAMBLA, Squared, TexturedButtons, Twister, TwisterSet |
| Files seen in | Code/Core/Main.lua, FixGit.lua, GuardBot.lua, NAMBLA.lua, Squared.lua, SquaredPlayer.lua, TexturedButtons.lua, Twister.lua |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, lua_call |
| Example locations | Initialize, NAMBLA_Initialize, OnLoad, OnLoadComplete, OnUnload, Shutdown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
| Local definition count | 0 |
| Documentation references | 3 |
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

SystemData.SystemData.Events.PLAYER_CAREER_LINE_UPDATED field accessed by 8 addons; commonly found in Initialize and NAMBLA_Initialize, OnLoad, OnLoadComplete, OnUnload, Shutdown, SquaredPlayer.UpdateCareer, SystemData.Events.PLAYER_CAREER_LINE_UPDATED, Unload, _PlayerCareerLineUpdated, event_page, event_registration, lua_call contexts.

## Seen In

- Enemy
- FixGit
- GuardBot
- NAMBLA
- Squared
- TexturedButtons
- Twister
- TwisterSet

## Related APIs

- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Notes

- Observed in contexts: Initialize, NAMBLA_Initialize, OnLoad, OnLoadComplete, OnUnload, Shutdown
