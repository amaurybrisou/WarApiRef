# wstring.find

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 55/100
- Seen in: 1 addons

## Confidence Assessment

- Level: MEDIUM

- Score: 55/100

- Rationale: Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, argument pattern is consistent.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:1469` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChatLogs.ProcessLootRollEntry |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
wstring.find(arg1, arg2)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: text |
| arg2 | Observed as a runtime window or control identifier. | Observed values: find_TEXT_NEEDGREED_ROLL_HEADER, find_TEXT_SELECT, find_TEXT_WINNER_HEADER |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- TidyChat

## Examples

- TidyChat: TidyChatLogs.ProcessLootRollEntry -> wstring.find(text, find_TEXT_SELECT)
- TidyChat: TidyChatLogs.ProcessLootRollEntry -> wstring.find(text, find_TEXT_WINNER_HEADER)
- TidyChat: TidyChatLogs.ProcessLootRollEntry -> wstring.find(text, find_TEXT_NEEDGREED_ROLL_HEADER)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
