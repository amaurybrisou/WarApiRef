# wstring.match

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

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
| Addons seen in | AdvancedRenownTrainer, Aura, CombatTextNames, Enemy, RoR_SoR, TidyChat, WSCT |
| Files seen in | `/workspace/data/raw/Aura/Libraries/LibPickle.lua:428`, `/workspace/data/raw/Aura/Source/AuraHelpers.lua:129`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLog.lua:700`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLog.lua:709`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLog.lua:773`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLog.lua:816`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLog.lua:863`, `/workspace/data/raw/Enemy/Code/Core/Utils.lua:370` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed, Aura: AuraHelpers.cleanWString, Aura: DePickler:unpickle, CombatTextNames: CombatTextNames.TruncateAbilityName, Enemy: Enemy.CombatLog_ParseHeal, Enemy: Enemy.CombatLog_ParseIncommingDamage |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 35 |
| Global usage count | 35 |
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

Observed as a global function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: Enemy.toWString(str), link, mit |
| arg2 | Observed as a function or method reference. | Observed values: BuildPattern, L "([^\^]+).*", L "([^^]+)^?.*" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AdvancedRenownTrainer
- Aura
- CombatTextNames
- Enemy
- RoR_SoR
- TidyChat
- WSCT

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed -> wstring.match(link, arsenalUrl..BuildPattern)
- AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed -> wstring.match(link, BuildPattern)
- AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed -> wstring.match(link, wardrobeUrl..BuildPattern)
- Aura: AuraHelpers.cleanWString -> wstring.match(wstr, L "([^\^]+).*")
- Aura: DePickler:unpickle -> wstring.match(s, L "P1$")
- CombatTextNames: CombatTextNames.TruncateAbilityName -> wstring.match(wstring.gsub(text,L "%l*",L ""), L "([^^]+)^?.*")

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.gsub](global_wstring.gsub.md) (HIGH 100/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 100/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
