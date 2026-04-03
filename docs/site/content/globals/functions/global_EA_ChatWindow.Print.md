# EA_ChatWindow.Print

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 123 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | ActionBarHide, AdvancedPetAssist, Amethyst, ArmorGraphicToggle, Arsenal Rank, AuctionStats, Aura, AutoBand |
| Files seen in | AB_realmrank.lua, AB_stats.lua, AB_util.lua, APACore.lua, ActionBarHide.lua, Amethyst.lua, ArmorGraphicToggle.lua, ArsenalRank.lua |
| Namespaces detected | EA_ChatWindow |
| Source kinds | lua_calls |
| Example locations | ActionBarHide: OnLoad, AdvancedPetAssist: Print, Amethyst: ApplySettings, Amethyst: Initialize, ArmorGraphicToggle: print, Arsenal Rank: init |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 561 |
| Global usage count | 561 |
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
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
EA_ChatWindow.Print(arg1)
```

## Description

Observed as a global function across 123 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "That state does not exist", C.ADDON..L "Broadcasted Tier "..towstring(tostring(tier))..(scope=="wb" and L " request to warband." or L " request to party."), C.ADDON..L "Declined leader request." |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Writes output to the chat or UI surface.

## Seen In

- ActionBarHide
- AdvancedPetAssist
- Amethyst
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
- Hopper
- I HATE YOU THIS MUCH
- JunkDump
- Killer
- Lib RuString
- LibCooldown
- LibGuard
- LibSlash
- LibWBToggler
- LoyalPet
- Map
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
- Amethyst: ApplySettings -> EA_ChatWindow.Print(L "Amethyst: Settings applied.")
- Amethyst: Initialize -> EA_ChatWindow.Print(L "Amethyst Initialized. Configuration via /amt")
- ArmorGraphicToggle: print -> EA_ChatWindow.Print(towstring(msg))
- Arsenal Rank: init -> EA_ChatWindow.Print(towstring("Starting ArsenalRank!"))

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [WindowGetTintColor](../../window_api/functions/window_WindowGetTintColor.md) (HIGH 100/100) - Window Function
- [BroadcastEvent](global_BroadcastEvent.md) (HIGH 93/100) - Global Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [RegisterEventHandler](global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Affects

- [SystemData.Events.SCENARIO_POST_MODE](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_POST_MODE.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
