# EA_ChatWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 125

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | ActionBarHide, AdvancedPetAssist, AdvancedRenownTrainer, Amethyst, AnywhereTrainerAdditions, ArmorGraphicToggle, Arsenal Rank, AuctionStats |
| Files seen in | Code/Core/Utils.lua, DefaultLayouts/Yak_Reso_Injector.lua, Elements/EffigyBar.lua, Elements/EffigyIndicator.lua, PhantomLib/PhantomLib.lua, Rules/EffigyRule.lua, Source/AuraAddon.lua, Source/Common.lua |
| Namespaces detected | EA_ChatWindow |
| Source kinds | lua_calls |
| Example locations | ActionBarHide: OnLoad, AdvancedPetAssist: Print, AdvancedRenownTrainer: ExportToLink, Amethyst: ApplySettings, Amethyst: Initialize, AnywhereTrainerAdditions: EquipmentLButtonDown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 474 |
| Global usage count | 12 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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

Shared function table with 12 member functions; the primary API surface for 131 addons.

## Functions

- EA_ChatWindow.GetCurrentChannel
- EA_ChatWindow.InsertItemLink
- EA_ChatWindow.InsertQuestLink
- EA_ChatWindow.InsertText
- EA_ChatWindow.OnKeyEnter
- EA_ChatWindow.OnSettingsChanged
- EA_ChatWindow.Print
- EA_ChatWindow.SaveSettings
- EA_ChatWindow.ShowTabCycleButtons
- EA_ChatWindow.SwitchChannelWithExistingText
- EA_ChatWindow.UpdateSocialWindowButton
- EA_ChatWindow.UpdateTabScrollWindow

## Observed Members

- none

## Seen In

- ActionBarHide
- AdvancedPetAssist
- AdvancedRenownTrainer
- Amethyst
- AnywhereTrainerAdditions
- ArmorGraphicToggle
- Arsenal Rank
- AuctionStats
- Aura
- AutoBand
- AutoFocus
- AutoSalvage
- BagOMatic
- BankArkel
- BlackBook
- Bloody Mess
- BuffThrottle
- Calling
- CastSequence
- ChattyCathy
- CleanUnitFrames
- Crafting Info Tooltip
- CraftingWillard
- Cram The Spam
- Crusher
- DAoCBuff
- DaemonAssist
- DasBoot
- Dascore
- DeMoNiCore
- DetauntHelper
- DwarfTalk
- EZGuard
- Effigy
- EmoteAlert
- Enemy
- Fight Finder
- FixGit
- GCDsaver
- Greedy
- Group Icons
- GroupRange
- GuardLine
- GuildWarden
- HealAll
- HealGrid
- Hopper
- I HATE YOU THIS MUCH
- JunkDump
- Killer
- Kwestor
- Lib RuString
- LibCooldown
- LibGuard
- LibSlash
- LibWBToggler
- LoyalPet
- Map
- MapPin
- MarkBuff
- Mech
- MegaphonePlusPlus
- MiniMapMonster
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Motion
- NPC Item Sale Price
- NerfedButtons
- NervAlert
- NoOverheal
- OilTimer
- PartyAd
- Phantom
- PlanB
- Pure
- Queue Queuer
- QueuePopTimer
- QuickNameActions+
- QuickTacticSwitch
- QuickWarReport
- RVMOD_Manager
- RandomMount
- RandomSayings
- Rangechecker
- RoR_SoR
- SOR
- ScenarioStats
- Shinies
- Soloq
- Squared
- SquaredClick
- SquaredHDIndicator
- Swift Assist
- Targets
- TastyButtons
- ThankTheResser
- TheSeeker
- ThinkOutLoud
- TidyChat
- TidyQueue
- TimeToDie
- Tokens
- Tome Titan
- Trakario
- Twister
- UnrealDBAnnouncer
- VerticalMorale
- VerticalTactics
- Vertigo
- WARCommander
- WARRatingBuster
- WSCT
- WarBoard_TaliMon
- WarTriage
- War_RU
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- ZCurse_Profiler
- ZonePOP
- alertMod
- nLootLink
- rorAutoInviter
- scnoload
- talisman-monitor
- wbLeadHelper
- whom
- xHUD
- xPanels
- yAssistHelper

## Examples

- ActionBarHide: OnLoad -> EA_ChatWindow.Print(L "ActionBarHide Loaded")
- AdvancedPetAssist: Print -> EA_ChatWindow.Print(L "<icon149> "..msg)
- AdvancedRenownTrainer: ExportToLink -> EA_ChatWindow.InsertText(hl)
- Amethyst: ApplySettings -> EA_ChatWindow.Print(L "Amethyst: Settings applied.")
- Amethyst: Initialize -> EA_ChatWindow.Print(L "Amethyst Initialized. Configuration via /amt")
- AnywhereTrainerAdditions: EquipmentLButtonDown -> EA_ChatWindow.InsertItemLink(itemData)

## Notes

- none
