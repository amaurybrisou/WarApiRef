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
| Addons seen in | ActionFraction, AdjustTheTip, AdvancedRenownTrainer, AggroMeter, Aura, BuddyBind, CCTV, CaVES |
| Files seen in | Bars/HealGridCastBar.lua, Code/CombatLog/CombatLog.lua, Code/Core/Utils.lua, Configuration/Config.lua, Configuration/HopperConfig.lua, Configuration/HopperConfig_General.lua, Configuration/PureConfig.lua, Configuration/PureConfig_Group.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | ActionFraction: Initialize, AdjustTheTip: GetAbilityCastTimeText, AdjustTheTip: UpdateInfluenceTooltip, AdvancedRenownTrainer: ImportOkButtonPressed, AdvancedRenownTrainer: OnHyperLinkLButtonUp, AggroMeter: OnMouseOverStart |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 666 |
| Global usage count | 11 |
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

Shared function table with 11 member functions; the primary API surface for 87 addons.

## Functions

- wstring.byte
- wstring.char
- wstring.find
- wstring.format
- wstring.gsub
- wstring.len
- wstring.lower
- wstring.match
- wstring.reverse
- wstring.sub
- wstring.upper

## Observed Members

- none

## Seen In

- ActionFraction
- AdjustTheTip
- AdvancedRenownTrainer
- AggroMeter
- Aura
- BuddyBind
- CCTV
- CaVES
- Calling
- CastSequence
- CleanUnitFrames
- CombatTextNames
- CoolDownLine
- Crafting Info Tooltip
- Cram The Spam
- Crusher
- DAoCBuff
- DammazKron
- DuffTimer
- EA_UiDebugTools
- Effigy
- Emojii
- Enemy
- GetStats
- Group Icons SG
- GroupSpotter
- GuardLine
- HealGrid
- Hopper
- I HATE YOU THIS MUCH
- Info_DeathBlow
- Info_Loot
- Info_Points
- KeyBar
- KillTracker
- Killer
- Kwestor
- LibDateTime
- LibGuard
- LibJson
- LibPickle
- MapMonster
- MapPin
- MegaphonePlusPlus
- Miracle Grow Remix
- Moth
- Motion
- NAMBLA
- NaturalLog
- Obsidian
- PartyCast
- PotionBar
- Pure
- Queue Queuer
- QueuePopTimer
- QuickNameActions+
- RVAPI_Range
- RVMOD_SquaredDistances
- RealmStatus
- Refer
- ReliquaryHunter
- ResHelp
- RoR_SoR
- SNT_BUTTONS
- SNT_CASTBAR
- SOR
- Sequencer
- SessionRPs
- Shinies
- Squared
- Statdoll
- Statdoll Light
- Statdoll Remix
- Swinger
- TexturedButtons
- TidyChat
- Tokens
- TomeTracker
- Trakario
- UnrealDBAnnouncer
- VPBreakdown
- Vectors
- WSCT
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- XpStatus+G
- whatsPugSc

## Examples

- ActionFraction: Initialize -> wstring.format(L "%.01f", ActionFraction.Version)
- AdjustTheTip: GetAbilityCastTimeText -> wstring.format(L "%.1f", castTime)
- AdjustTheTip: GetAbilityCastTimeText -> wstring.format(L "%d", castTime)
- AdjustTheTip: UpdateInfluenceTooltip -> wstring.format(L "/%d (", influenceData.rewardLevel[3].amountNeeded)
- AdjustTheTip: UpdateInfluenceTooltip -> wstring.format(L "%d%%,", percent)
- AdjustTheTip: UpdateInfluenceTooltip -> wstring.format(L "%d%%)", percent)

## Notes

- none
