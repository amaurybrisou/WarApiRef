# Inference Rules

## Thresholds

- Score clamp: 0 to 100.

- HIGH: 70 to 100.

- MEDIUM: 40 to 69.

- LOW: 0 to 39.

## Promotion Rules

- HIGH symbols are promoted automatically.

- MEDIUM symbols are promoted only when they also have known namespace, default UI, event-binding, direct XML attribute, or XML-plus-Lua evidence.

- LOW symbols remain evidence only and do not enter canonical symbol pages.

## Special Rules

- Known namespace override: a known engine namespace is lifted to at least MEDIUM unless strong contradictory local-only evidence is present.

- XML privilege: direct XML handler attribute evidence is treated as a strong platform signal.

- Addon-prefix penalty: addon-specific naming lowers confidence unless cross-addon or namespace evidence outweighs it.

- Wrapper detection: thin local wrappers over a stronger primitive API are penalized and usually remain non-canonical.

- Cross-source reinforcement: symbols corroborated by multiple generated source types receive additional weight.

## Signal Weights

| Condition | Weight | Category |
| --- | --- | --- |
| Appears in 4 or more distinct addons | +30 | positive |
| Appears in 2 to 3 distinct addons | +18 | positive |
| Appears in default interface or base UI namespace | +35 | positive |
| Appears directly in XML handler attributes | +30 | positive |
| Appears as a global call with no local definition | +20 | positive |
| Matches a known engine namespace | +25 | positive |
| Used in event registration or dispatch patterns | +18 | positive |
| Referenced from initialization flow evidence | +10 | positive |
| Observed from both XML and Lua paths | +20 | positive |
| Consistent role across multiple addons | +15 | positive |
| Argument pattern is consistent across usages | +10 | positive |
| Return usage is consistent across usages | +8 | positive |
| Observed in slash command registration patterns | +10 | positive |
| Referenced by generated docs or reference files | +25 | positive |
| Reinforced across multiple generated source types | +20 | positive |
| Defined locally in only one addon | -35 | negative |
| Only appears in one addon with project-specific naming | -25 | negative |
| Only appears as a local helper | -30 | negative |
| Contains addon-specific prefix or suffix | -18 | negative |
| Only referenced by one internal module with no XML or global usage | -15 | negative |
| Inferred from one weak usage site | -20 | negative |
| Conflicting signatures across usages | -15 | negative |
| Conflicting semantic roles across addons | -20 | negative |
| Likely wrapper around another platform API | -12 | negative |
| Never appears outside addon-local flow | -25 | negative |
| Looks like a placeholder or computed expression | -100 | negative |

## Pipeline Heuristics

- Parsed only generated markdown under docs/addon-api; raw addon Lua and XML were not rescanned.
- Dropped call-site aliases when a qualified and unqualified symbol shared caller, line, arguments, and final segment.
- Filtered addon-local calls when the symbol root or final segment matched addon-owned namespaces or function names.
- Promoted symbols to platform API candidates when they appeared in multiple addons, in XML/window handlers, or under known WAR/UI namespaces.
- Marked high, medium, or low confidence strictly from observed addon spread and namespace signal, not from external API databases.
- Inferred signatures only from repeated argument positions and clear naming signal; uncertain return values remain explicitly marked unknown.
- Extracted SystemData and GameData members only from observed tokens in calls, event pages, flow evidence, and generated globals tables.
