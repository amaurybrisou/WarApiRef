# wstring.match

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 16 addons

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
| Addons seen in | AdvancedRenownTrainer, Aura, CCTV, CombatTextNames, Cram The Spam, EA_UiDebugTools, Enemy, LibDateTime |
| Files seen in | AdvancedRenownTrainingImportExport.lua, CCTV.lua, Code/CombatLog/CombatLog.lua, Code/Core/Utils.lua, CramTheSpam.lua, LibDateTime.lua, LibPickle.lua, Libraries/LibPickle.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: ImportOkButtonPressed, Aura: cleanWString, Aura: unpickle, CCTV: Update, CombatTextNames: TruncateAbilityName, Cram The Spam: GetCleanName |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 95 |
| Global usage count | 95 |
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
wstring.match(arg1, arg2)
```

## Description

Observed as a global function across 16 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: CText, DevPad_Settings.Code, Enemy.toWString(str) |
| arg2 | Observed as a function or method reference. | Observed values: BuildPattern, L "%]%[.+%]:(.+)", L "(%d+)." |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedRenownTrainer
- Aura
- CCTV
- CombatTextNames
- Cram The Spam
- EA_UiDebugTools
- Enemy
- LibDateTime
- LibPickle
- MapPin
- QuickNameActions+
- RoR_SoR
- SessionRPs
- TidyChat
- Trakario
- WSCT

## Examples

- AdvancedRenownTrainer: ImportOkButtonPressed -> wstring.match(link, arsenalUrl..BuildPattern)
- AdvancedRenownTrainer: ImportOkButtonPressed -> wstring.match(link, BuildPattern)
- AdvancedRenownTrainer: ImportOkButtonPressed -> wstring.match(link, wardrobeUrl..BuildPattern)
- Aura: cleanWString -> wstring.match(wstr, L "([^\^]+).*")
- Aura: unpickle -> wstring.match(s, L "P1$")
- CCTV: Update -> wstring.match(towstring(Snare_Timer), L "(%d+).")

## Used With

- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.gsub](global_wstring.gsub.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
