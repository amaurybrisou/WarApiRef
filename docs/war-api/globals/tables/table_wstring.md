# wstring

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 90/100

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, AggroMeter, Aura, CombatTextNames, Crafting Info Tooltip, DAoCBuff, EA_UiDebugTools, Effigy |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.lua:205`, `/workspace_addons/Aura/Libraries/LibPickle.lua:428`, `/workspace_addons/Aura/Libraries/LibPickle.lua:65`, `/workspace_addons/Aura/Source/AuraHelpers.lua:129`, `/workspace_addons/CraftValueTip/CraftValueTip.lua:108`, `/workspace_addons/DAoCBuff/Source/DAoCBuffFrames.lua:121`, `/workspace_addons/DAoCBuff/Source/Transcode.lua:289`, `/workspace_addons/Effigy/Effigy.lua:227` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed, AdvancedRenownTrainer: AdvancedRenownTraining.OnHyperLinkLButtonUp, AggroMeter: AggroMeter.OnMouseOverStart, Aura: Aura.local.wgsub, Aura: AuraHelpers.cleanWString, Aura: DePickler:unpickle |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 166 |
| Global usage count | 9 |
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

Observed shared global table or namespace surfaced in 23 addons.

## Functions

- wstring.char
- wstring.find
- wstring.format
- wstring.gsub
- wstring.len
- wstring.lower
- wstring.match
- wstring.sub
- wstring.upper

## Observed Members

- none

## Seen In

- AdvancedRenownTrainer
- AggroMeter
- Aura
- CombatTextNames
- Crafting Info Tooltip
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- GetStats
- GuardLine
- Killer
- LibGuard
- MapMonster
- MapPin
- Miracle Grow Remix
- Moth
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- WSCT

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed -> wstring.match(link, arsenalUrl..BuildPattern)
- AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed -> wstring.match(link, BuildPattern)
- AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed -> wstring.match(link, wardrobeUrl..BuildPattern)
- AdvancedRenownTrainer: AdvancedRenownTraining.OnHyperLinkLButtonUp -> wstring.gsub(linkData, hyperlinkPrefix, ERASE)
- AggroMeter: AggroMeter.OnMouseOverStart -> wstring.format(L "%.01f", (AggroMeter.AggroHolder[tostring(MobNumber)][tonumber(TimerNumber)].aggro/AggroMeter.MaxAggro[tostring(MobNumber)])*100)
- Aura: Aura.local.wgsub -> wstring.gsub(s, pattern, replace)

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
