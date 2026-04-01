# LibSlash

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 115

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Appears in slash command registration patterns: Observed in shared command registration flows.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, Aura, AutoBand, AutoChannel, AutoMark, BagOMatic, BuffHead, CM_ClosetGoblin |
| Files seen in | `/workspace/AdvancedPetAssist/AdvancedPetAssist.lua:193`, `/workspace/Aura/Source/AuraAddon.lua:70`, `/workspace/AutoMark/Source/AutoMark.lua:33`, `/workspace/Autoband/AutoBand.lua:30`, `/workspace/BuffHead/Core.lua:79`, `/workspace/ClosetGoblin/ClosetGoblin.lua:83`, `/workspace/CraftValueTip/CraftValueTip.lua:33`, `/workspace/DAoCBuff/Source/DAoCBuff.lua:219` |
| Namespaces detected | LibSlash |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.Initialize, Aura: AuraAddon.OnInitialize, AutoBand: AutoBand.init, AutoChannel: AutoChannel.Initialize, AutoMark: AutoMark.OnInitialize, BagOMatic: BagOMatic.init |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 107 |
| Global usage count | 4 |
| Local definition count | 2 |
| Documentation references | 1 |
| Initialization flow references | 3 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | yes |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed shared global table or namespace surfaced in 56 addons.

## Functions

- LibSlash.IsSlashCmdRegistered
- LibSlash.RegisterSlashCmd
- LibSlash.RegisterWSlashCmd
- LibSlash.UnregisterSlashCmd

## Observed Members

- none

## Seen In

- AdvancedPetAssist
- Aura
- AutoBand
- AutoChannel
- AutoMark
- BagOMatic
- BuffHead
- CM_ClosetGoblin
- Crafting Info Tooltip
- DAoCBuff
- DaemonAssist
- DeepSleep
- Effigy
- FastInteract
- GCDsaver
- GoldTracker
- GuardLine
- JunkDump
- Killer
- LibGroup
- LibGuard
- LibSlash
- LibWBToggler
- MapMonster
- MapPin
- MegaphonePlusPlus
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- NPC Item Sale Price
- PeaceOut
- PlanB
- Pocket Palette
- PotionBar
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RRCount
- RVMOD_Manager
- RandomMount
- RetAlert
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WarTriage
- WhoHealedMe
- WoH-Reticle
- followTheLeader
- wbLeadHelper

## Examples

- AdvancedPetAssist: AdvancedPetAssist.Initialize -> LibSlash.RegisterSlashCmd("apa", function(input)APA.SlashHandler(input)end)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("aura", AuraAddon.Slash)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("auraconfig", AuraAddon.Slash)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("showaura", AuraAddon.Slash)
- AutoBand: AutoBand.init -> LibSlash.RegisterSlashCmd("autoband", function(msg)AutoBand.parse_cmd(msg)end)
- AutoBand: AutoBand.init -> LibSlash.IsSlashCmdRegistered("ab")

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
