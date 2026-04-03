# SystemData.Settings.Language.active

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | DetauntHelper, JunkDump, NerfedButtons, ZCurse_Profiler |
| Files seen in | CurseProfilerCompiled.lua, JunkDump.lua, JunkDumpOptions.lua, Source/Lang.lua, UI/Lang.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | DHGetActiveLanguageMapping, Initialize, NBSGetActiveLanguageMapping, lua_call, repeatnil |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 1 |
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

SystemData.SystemData.Settings.Language.active field accessed by 4 addons; commonly found in DHGetActiveLanguageMapping and Initialize, NBSGetActiveLanguageMapping, lua_call, repeatnil contexts.

## Seen In

- DetauntHelper
- JunkDump
- NerfedButtons
- ZCurse_Profiler

## Notes

- Observed in contexts: DHGetActiveLanguageMapping, Initialize, NBSGetActiveLanguageMapping, lua_call, repeatnil
