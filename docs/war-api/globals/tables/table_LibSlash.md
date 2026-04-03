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
| Addons seen in | AdvancedPetAssist, Aura, AutoMark, BagOMatic, BuffHead, CM_ClosetGoblin, DAoCBuff, GuardLine |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/AdvancedPetAssist.lua:193`, `/workspace/data/raw/Aura/Source/AuraAddon.lua:70`, `/workspace/data/raw/AutoMark/Source/AutoMark.lua:33`, `/workspace/data/raw/BuffHead/Core.lua:79`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.lua:87`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:219`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:25`, `/workspace/data/raw/GuardLine/GuardLine.lua:63` |
| Namespaces detected | LibSlash |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.Initialize, Aura: AuraAddon.OnInitialize, AutoMark: AutoMark.OnInitialize, BagOMatic: BagOMatic.init, BuffHead: BuffHead.local.RegisterLibs, BuffHead: RegisterLibs |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 53 |
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

Observed shared global table or namespace surfaced in 30 addons.

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
- AutoMark
- BagOMatic
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- GuardLine
- Killer
- LibGroup
- LibGuard
- LibSlash
- LibWBToggler
- MiracleGrowLight
- PartyCast
- PlanB
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- Swift Assist
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- WoH-Reticle
- followTheLeader

## Examples

- AdvancedPetAssist: AdvancedPetAssist.Initialize -> LibSlash.RegisterSlashCmd("apa", function(input)APA.SlashHandler(input)end)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("aura", AuraAddon.Slash)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("auraconfig", AuraAddon.Slash)
- Aura: AuraAddon.OnInitialize -> LibSlash.RegisterSlashCmd("showaura", AuraAddon.Slash)
- AutoMark: AutoMark.OnInitialize -> LibSlash.RegisterSlashCmd("automark", AutoMark.OnSlashCommand)
- BagOMatic: BagOMatic.init -> LibSlash.RegisterSlashCmd("bagomatic", function(msg)BagOMatic.parse_cmd(msg)end)

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
