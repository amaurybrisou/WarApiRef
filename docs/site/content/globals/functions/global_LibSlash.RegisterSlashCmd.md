# LibSlash.RegisterSlashCmd

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 85/100
- Seen in: 148 addons

## Confidence Assessment

- Level: HIGH

- Score: 85/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Appears in slash command registration patterns: Observed in shared command registration flows.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | ActionBarCD, ActionBarHide, AdvancedPetAssist, Amethyst, ArmorGraphicToggle, Atlas, AuctionStats, Aura |
| Files seen in | ActionBarCD.lua, ActionBarHide.lua, AdvancedPetAssist.lua, Amethyst.lua, ArmorGraphicToggle.lua, AuctionStats.lua, AutoBand.lua, AutoFocus.lua |
| Namespaces detected | LibSlash |
| Source kinds | lua_calls |
| Example locations | ActionBarCD: Initialize, ActionBarHide: OnLoad, AdvancedPetAssist: Initialize, Amethyst: Initialize, ArmorGraphicToggle: OnInitialize, Atlas: RegisterSlashCommands |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 259 |
| Global usage count | 259 |
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
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
LibSlash.RegisterSlashCmd(slashName, handler)
```

## Description

Observed wiring slash commands through a shared command-registration table.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| slashName | Observed as a slash command token. | Observed values: "CCM", "CCTV", "CDL" |
| handler | Observed as a command handler callback. | Observed values: Atlas.Configuration.Toggle, AuctionAssist.OnSlash, AuctionStats.OnSlashCommand |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- ActionBarCD
- ActionBarHide
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
- CCTV
- CDown
- CM_ClosetGoblin
- ChatAlert
- ChattyCathy
- CleanUnitFrames
- CoolDownLine
- Countdown
- Crusher
- DAoCBuff
- DPSMeter
- DaemonAssist
- DammazKron
- DasBoot
- DeepSleep
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
- GuardBot
- GuardLine
- HealGrid
- Hopper
- I HATE YOU THIS MUCH
- IdentityFound
- InfoScroller
- KeyBar
- KillTracker
- Killer
- Kwestor
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
- Swift Assist
- TargetRing
- Targets
- TastyButtons
- TexturedButtons
- ThankTheResser
- TheSeeker
- TokenMachine
- Tome Titan
- TomeTracker
- Tortall_SCC
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
- xHUD
- xPanels
- zMailMod

## Examples

- ActionBarCD: Initialize -> LibSlash.RegisterSlashCmd("abcd", function(args)ActionBarCD.SlashHandler(args)end)
- ActionBarHide: OnLoad -> LibSlash.RegisterSlashCmd("abh", function()ActionBarHide.OptionsWindow()end)
- AdvancedPetAssist: Initialize -> LibSlash.RegisterSlashCmd("apa", function(input)APA.SlashHandler(input)end)
- Amethyst: Initialize -> LibSlash.RegisterSlashCmd("amt", function(msg)Amethyst.Slash(msg)end)
- Amethyst: Initialize -> LibSlash.RegisterSlashCmd("amethyst", function(msg)Amethyst.Slash(msg)end)
- ArmorGraphicToggle: OnInitialize -> LibSlash.RegisterSlashCmd("agt", function(msg)ArmorGraphicToggle.Slash(msg)end)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event

## Used With

- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow.Print](global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
