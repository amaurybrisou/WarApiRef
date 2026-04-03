# wstring

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 78/100

## Confidence Assessment

- Level: HIGH

- Score: 78/100

- Rationale: Promoted as HIGH confidence because matches a known engine namespace, called globally with no local definition, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Moth, TidyChat |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:270`, `/workspace/data/raw/Moth/MothHelpers.lua:39`, `/workspace/data/raw/TidyChat/TidyChat.lua:1469`, `/workspace/data/raw/TidyChat/TidyChat.lua:852` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.UpdateLevel, Moth: MothHelpers.CapitalizeWString, TidyChat: TidyChatHooks.OnHyperLinkLButtonUpHook, TidyChat: TidyChatLogs.ProcessLootRollEntry |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 6 |
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

Observed shared global table or namespace surfaced in 2 addons.

## Functions

- wstring.find
- wstring.gsub
- wstring.lower
- wstring.match
- wstring.sub
- wstring.upper

## Observed Members

- none

## Seen In

- Moth
- TidyChat

## Examples

- Moth: Moth.UpdateLevel -> wstring.sub(unitTierDesc, 1, 1)
- Moth: MothHelpers.CapitalizeWString -> wstring.lower(wstring.sub(wstr,2))
- Moth: MothHelpers.CapitalizeWString -> wstring.sub(wstr, 1, 1)
- Moth: MothHelpers.CapitalizeWString -> wstring.sub(wstr, 2)
- Moth: MothHelpers.CapitalizeWString -> wstring.upper(wstring.sub(wstr,1,1))
- TidyChat: TidyChatHooks.OnHyperLinkLButtonUpHook -> wstring.gsub(linkData, TCROLL_TAG, ERASE)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
