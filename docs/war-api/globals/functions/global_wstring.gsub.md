# wstring.gsub

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 31 addons

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
| Addons seen in | AdvancedRenownTrainer, Aura, BarText (Influence), CombatTextNames, Crafting Info Tooltip, Cram The Spam, DammazKron, EA_UiDebugTools |
| Files seen in | AdvancedRenownTrainingImportExport.lua, BarText_Influence.lua, Core/DK_Core.lua, Core/DK_Utils.lua, CraftValueTip.lua, CramTheSpam.lua, GroupIconsSG.lua, GuardLine.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: OnHyperLinkLButtonUp, Aura: unpickle, Aura: wgsub, BarText (Influence): PlayerInfluenceUpdated, CombatTextNames: TruncateAbilityName, Crafting Info Tooltip: GetPhrase |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 163 |
| Global usage count | 163 |
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

Observed as a global function across 31 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: (QuickNameActionsRessurected.chatMessage1), (QuickNameActionsRessurected.chatMessage2), (QuickNameActionsRessurected.tellMessage1) |
| arg2 | Observed as a text or wstring payload. | Observed values: GuardLine.FixString(GameData.Player.name)..L "'s", L " .+", L " of " |
| arg3 | Observed as a text or wstring payload. | Observed values: "", ERASE, L "" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedRenownTrainer
- Aura
- BarText (Influence)
- CombatTextNames
- Crafting Info Tooltip
- Cram The Spam
- DammazKron
- EA_UiDebugTools
- Group Icons SG
- GuardLine
- I HATE YOU THIS MUCH
- KillTracker
- Killer
- LibGuard
- LibPickle
- MapMonster
- MapPin
- MegaphonePlusPlus
- Miracle Grow Remix
- QuickNameActions+
- RVAPI_Range
- RVMOD_SquaredDistances
- Refer
- RoR_SoR
- Shinies
- SocialWindow 2.0
- TidyChat
- TomeTracker
- Trakario
- WSCT
- WarBoard_WarWhisperer

## Examples

- AdvancedRenownTrainer: OnHyperLinkLButtonUp -> wstring.gsub(linkData, hyperlinkPrefix, ERASE)
- Aura: unpickle -> wstring.gsub(s, L "P1$", L "")
- Aura: wgsub -> wstring.gsub(s, pattern, replace)
- BarText (Influence): PlayerInfluenceUpdated -> wstring.gsub(BarText_Influence_Config.currentFormatString, L "%[inf%]", inf)
- BarText (Influence): PlayerInfluenceUpdated -> wstring.gsub(text, L "%[infmax%]", infmax)
- BarText (Influence): PlayerInfluenceUpdated -> wstring.gsub(text, L "%[infp%]", infp)

## Used With

- [EA_Window_PublicQuestTracker.GetLocalAreaInfluenceID](global_EA_Window_PublicQuestTracker.GetLocalAreaInfluenceID.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [wstring.format](global_wstring.format.md) (HIGH 75/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
