# LibSlash

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Score: 100/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Appears in slash command registration patterns: Observed in shared command registration flows.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AbilityAlert, ActionBarCD, ActionBarHide, ActionFraction, ActionPointWatch, AdvancedPetAssist, Amethyst, ArmorGraphicToggle |
| Files seen in | Configuration/Main.lua, Core/DK_Core.lua, Modules/HealGridMouseClick.lua, Source/AuraAddon.lua, Source/AutoMark.lua, Source/Core.lua, Source/Crusher.lua, Source/DAoCBuff.lua |
| Namespaces detected | LibSlash |
| Source kinds | lua_calls |
| Example locations | AbilityAlert: Initialize, ActionBarCD: Initialize, ActionBarHide: OnLoad, ActionFraction: Initialize, ActionFraction: Shutdown, ActionPointWatch: Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 354 |
| Global usage count | 5 |
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

Shared function table with 5 member functions; the primary API surface for 174 addons.

## Functions

- LibSlash.GetCmdHandler
- LibSlash.IsSlashCmdRegistered
- LibSlash.RegisterSlashCmd
- LibSlash.RegisterWSlashCmd
- LibSlash.UnregisterSlashCmd

## Observed Members

- none

## Seen In

- AbilityAlert
- ActionBarCD
- ActionBarHide
- ActionFraction
- ActionPointWatch
- AdvancedPetAssist
- Amethyst
- ArmorGraphicToggle
- Atlas
- AuctionStats
- Aura
- AutoBand
- AutoChannel
- AutoFocus
- AutoMark
- AutoSalvage
- BagOMatic
- BlackBook
- Brizio's Crappy Computer Medic
- BuffHead
- CCTV
- CDown
- CM_ClosetGoblin
- Calling
- CastSequence
- ChatAlert
- ChattyCathy
- CleanUnitFrames
- CoolDownLine
- Countdown
- Crafting Info Tooltip
- CraftingWillard
- Cram The Spam
- Crusher
- DAoCBuff
- DPSMeter
- DaemonAssist
- DammazKron
- DasBoot
- DeepSleep
- DetauntHelper
- Ding
- DuffTimer
- DwarfTalk
- Dye Preview
- EZCraftX
- EZGuard
- Effigy
- EmoteAlert
- EveryBodyGuard
- FastFriends
- FastInteract
- Fight Finder
- GCDTracker
- GCDsaver
- Greedy
- Group Icons
- Group Icons SG
- GroupRange
- GuardBot
- GuardLine
- HealGrid
- Hopper
- I HATE YOU THIS MUCH
- IdentityFound
- InfoScroller
- JunkDump
- KeyBar
- Keyset
- KeysetMonsterPlay
- KillTracker
- Killer
- Kwestor
- Lib RuString
- LibAddonButton
- LibGroup
- LibGuard
- LibSlash
- LibWBToggler
- LootAlert
- Map
- MapMonster
- MapPin
- MarkBuff
- Mech
- MegaphonePlus
- MegaphonePlusPlus
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Motion
- NPC Item Sale Price
- NaturalLog
- NerfedButtons
- Obsidian
- OilTimer
- PagerWentPoof
- PartyAd
- PartyCast
- PeaceOut
- Phantom
- PlanB
- Pocket Palette
- PotionBar
- Pure
- Queue Queuer
- Quick Performance Toggle
- QuickNameActions+
- QuickTacticSwitch
- QuickWarReport
- RVMOD_Manager
- RaidMeter
- RandomMount
- Rangechecker
- RealmStatus
- RetAlert
- RoR_SoR
- RoR_debolster
- Rotation
- RvRContribution
- SNT_BUTTONS
- SNT_CASTBAR
- SNT_PANEL
- SOR
- ScenarioStats
- SelfTarget
- SessionRPs
- Shinies
- SimpleCombatText
- Soloq
- Squared
- Statdoll Remix
- Swift Assist
- TargetRing
- Targets
- TastyButtons
- TexturedButtons
- ThankTheResser
- TheSeeker
- TidyChat
- TidyRoll
- TimeToDie
- TokenMachine
- Tokens
- Tome Titan
- TomeTracker
- Tortall_SCC
- TurretRange
- VPBreakdown
- Vectors
- VerticalMorale
- VerticalTactics
- Vertigo
- WARCommander
- WARRatingBuster
- WSCT
- WarBoard
- WarTriage
- Wargames
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- alertMod
- followTheLeader
- minesweep
- rorAutoInviter
- scnoload
- wbLeadHelper
- whatsPugSc
- whom
- xHUD
- xPanels
- zMailMod

## Examples

- AbilityAlert: Initialize -> LibSlash.IsSlashCmdRegistered("aa")
- AbilityAlert: Initialize -> LibSlash.RegisterWSlashCmd("aa", function(args)AbilityAlert.SlashHandler(args)end)
- AbilityAlert: Initialize -> LibSlash.RegisterWSlashCmd("abilityalert", function(args)AbilityAlert.SlashHandler(args)end)
- ActionBarCD: Initialize -> LibSlash.RegisterSlashCmd("abcd", function(args)ActionBarCD.SlashHandler(args)end)
- ActionBarHide: OnLoad -> LibSlash.RegisterSlashCmd("abh", function()ActionBarHide.OptionsWindow()end)
- ActionFraction: Initialize -> LibSlash.IsSlashCmdRegistered("af")

## Notes

- none
