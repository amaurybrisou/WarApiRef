# wstring.lower

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

## Confidence Assessment

- Level: HIGH

- Score: 100/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Cram The Spam, Effigy, HealGrid, Moth, Refer, UnrealDBAnnouncer |
| Files seen in | CramTheSpam.lua, Elements/EffigyLabel.lua, HealGridTags.lua, MothHelpers.lua, Refer.lua, UnrealDBAnnouncer.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | Cram The Spam: FuzzyFilterTest, Cram The Spam: GenericSlashCommand, Effigy: Render, HealGrid: ProcessTagLower, Moth: CapitalizeWString, Refer: FormatPlayerName |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
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
wstring.lower(arg1)
```

## Description

Observed as a global function across 6 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: nameFormatted, str, text |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Cram The Spam
- Effigy
- HealGrid
- Moth
- Refer
- UnrealDBAnnouncer

## Examples

- Cram The Spam: FuzzyFilterTest -> wstring.lower(wStringParam)
- Cram The Spam: GenericSlashCommand -> wstring.lower(wFirstWord)
- Effigy: Render -> wstring.lower(text)
- HealGrid: ProcessTagLower -> wstring.lower(str)
- Moth: CapitalizeWString -> wstring.lower(wstring.sub(wstr,2))
- Refer: FormatPlayerName -> wstring.lower(nameFormatted)

## Used With

- [wstring.gsub](global_wstring.gsub.md) (HIGH 100/100) - Global Function
- [wstring.upper](global_wstring.upper.md) (HIGH 100/100) - Global Function
- [wstring.find](global_wstring.find.md) (HIGH 75/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
