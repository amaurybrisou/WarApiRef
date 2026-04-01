# wstring.gsub

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
| Addons seen in | AdvancedRenownTrainer, Aura, CombatTextNames, Crafting Info Tooltip, EA_UiDebugTools, GuardLine, Killer, LibGuard |
| Files seen in | `/workspace/Aura/Libraries/LibPickle.lua:428`, `/workspace/Aura/Libraries/LibPickle.lua:65`, `/workspace/CraftValueTip/CraftValueTip.lua:108`, `/workspace/GuardLine/GuardLine.lua:151`, `/workspace/Killer/KillerUtils.lua:14`, `/workspace/LibGuard/Source/LibGuard.lua:362`, `/workspace/LibGuard/Source/LibGuard.lua:75`, `/workspace/MGRemix/MGRemix.lua:424` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.OnHyperLinkLButtonUp, Aura: Aura.local.wgsub, Aura: DePickler:unpickle, Aura: wgsub, CombatTextNames: CombatTextNames.TruncateAbilityName, Crafting Info Tooltip: CraftValueTip.GetPhrase |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 79 |
| Global usage count | 79 |
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
wstring.gsub(arg1, arg2, arg3)
```

## Description

Observed as a global function across 16 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: DevPad_Settings.Code, GetString(StringTables.Default.TEXT_CUR_DAMAGE_WAAAGH), GetString(StringTables.Default.TEXT_CUR_HEALING_WAAAGH) |
| arg2 | Observed as a text or wstring payload. | Observed values: GuardLine.FixString(GameData.Player.name)..L "'s", L " .+", L " of " |
| arg3 | Observed as a text or wstring payload. | Observed values: "", ERASE, L "" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AdvancedRenownTrainer
- Aura
- CombatTextNames
- Crafting Info Tooltip
- EA_UiDebugTools
- GuardLine
- Killer
- LibGuard
- MapMonster
- MapPin
- MegaphonePlusPlus
- Miracle Grow Remix
- RoR_SoR
- Shinies
- TidyChat
- WSCT

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.OnHyperLinkLButtonUp -> wstring.gsub(linkData, hyperlinkPrefix, ERASE)
- Aura: Aura.local.wgsub -> wstring.gsub(s, pattern, replace)
- Aura: DePickler:unpickle -> wstring.gsub(s, L "P1$", L "")
- Aura: wgsub -> wstring.gsub(s, pattern, replace)
- CombatTextNames: CombatTextNames.TruncateAbilityName -> wstring.gsub(wstring.gsub(text,L " of ",L "O"), L "%s", L "")
- CombatTextNames: CombatTextNames.TruncateAbilityName -> wstring.gsub(text, L " of ", L "O")

## Related APIs

- none

## Used With

- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [wstring.match](global_wstring.match.md) (HIGH 100/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
