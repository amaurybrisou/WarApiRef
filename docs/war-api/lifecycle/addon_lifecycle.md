# Addon Lifecycle Semantics

> Structured lifecycle facts derived from `.mod` manifest analysis.

## What This List Represents

- This page is not an inventory of all addon directories under the source root.
- It lists addons that reached `.mod` semantic lifecycle extraction in the source-first pipeline.
- Addons can be excluded for explicit reasons (no manifest, no resolved source files, or `.toc`-only).
- Use the diagnostics section below to identify exclusions and counts.

## Diagnostics

| Metric | Value |
| --- | --- |
| source_root | /workspace/data/raw |
| source_directories | 396 |
| manifest_discovered | 384 |
| source_scanned_addons | 384 |
| mod_semantic_addons | 384 |
| lifecycle_catalog_addons | 384 |
| excluded_entries | 12 |

## Exclusion Reasons

| Reason | Count |
| --- | --- |
| no_manifest | 12 |

## Excluded Addons

- AceLocale-3.0 - directory has no .mod or .toc manifest and is excluded at source discovery
- GuardPack - directory has no .mod or .toc manifest and is excluded at source discovery
- guildwindowexpanded-1-5-5 - directory has no .mod or .toc manifest and is excluded at source discovery
- harbinger.1-6-1 - directory has no .mod or .toc manifest and is excluded at source discovery
- orcanizer - directory has no .mod or .toc manifest and is excluded at source discovery
- timeinqueue-1-1-1 - directory has no .mod or .toc manifest and is excluded at source discovery
- unbearable-1-0-0 - directory has no .mod or .toc manifest and is excluded at source discovery
- warboard_crests-0-5-1 - directory has no .mod or .toc manifest and is excluded at source discovery
- warboard_toggler_pack-1-0-0 - directory has no .mod or .toc manifest and is excluded at source discovery
- yakui-1-3-6 - directory has no .mod or .toc manifest and is excluded at source discovery
- zTimeLib - directory has no .mod or .toc manifest and is excluded at source discovery
- zonepop1-3-1-3-0 - directory has no .mod or .toc manifest and is excluded at source discovery

## Source Addon Inventory

| Directory | Lifecycle Entry | Status | Reason |
| --- | --- | --- | --- |
| AAOTracker | AAOTracker | excluded | not present in .mod lifecycle semantic output |
| AbilityAlert | [AbilityAlert](addons/abilityalert) | included | .mod semantic lifecycle entry emitted |
| AbilityNotifier | [AbilityNotifier](addons/abilitynotifier) | included | .mod semantic lifecycle entry emitted |
| Ace | [Ace](addons/ace) | included | .mod semantic lifecycle entry emitted |
| AceLocale-3.0 | AceLocale-3.0 | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |
| ActionBarCD | [ActionBarCD](addons/actionbarcd) | included | .mod semantic lifecycle entry emitted |
| ActionBarColor | [ActionBarColor](addons/actionbarcolor) | included | .mod semantic lifecycle entry emitted |
| ActionFraction | [ActionFraction](addons/actionfraction) | included | .mod semantic lifecycle entry emitted |
| ActionPoints | [ActionPoints](addons/actionpoints) | included | .mod semantic lifecycle entry emitted |
| AddonButtonTogglers | [AddonButtonTogglers](addons/addonbuttontogglers) | included | .mod semantic lifecycle entry emitted |
| AdjustTheTip | [AdjustTheTip](addons/adjustthetip) | included | .mod semantic lifecycle entry emitted |
| AdvancedPetAssist | [AdvancedPetAssist](addons/advancedpetassist) | included | .mod semantic lifecycle entry emitted |
| AggroMeter | [AggroMeter](addons/aggrometer) | included | .mod semantic lifecycle entry emitted |
| AnywhereTrainer | [AnywhereTrainer](addons/anywheretrainer) | included | .mod semantic lifecycle entry emitted |
| AnywhereTrainerAdditions | [AnywhereTrainerAdditions](addons/anywheretraineradditions) | included | .mod semantic lifecycle entry emitted |
| ArmorGraphicToggle | [ArmorGraphicToggle](addons/armorgraphictoggle) | included | .mod semantic lifecycle entry emitted |
| ArsenalRank | [Arsenal Rank](addons/arsenal_rank) | included | .mod semantic lifecycle entry emitted |
| Asshat | [Asshat](addons/asshat) | included | .mod semantic lifecycle entry emitted |
| Atlas | [Atlas](addons/atlas) | included | .mod semantic lifecycle entry emitted |
| AuctionTweaker | [AuctionTweaker](addons/auctiontweaker) | included | .mod semantic lifecycle entry emitted |
| Aura | [Aura](addons/aura) | included | .mod semantic lifecycle entry emitted |
| AutoBand | [AutoBand](addons/autoband) | included | .mod semantic lifecycle entry emitted |
| AutoDismount | [AutoDismount](addons/autodismount) | included | .mod semantic lifecycle entry emitted |
| AutoFocus | [AutoFocus](addons/autofocus) | included | .mod semantic lifecycle entry emitted |
| AutoMark | [AutoMark](addons/automark) | included | .mod semantic lifecycle entry emitted |
| BBars_MechanicOnly | [BBars - Mechanic Only](addons/bbars_-_mechanic_only) | included | .mod semantic lifecycle entry emitted |
| BWMT | [BWMT](addons/bwmt) | included | .mod semantic lifecycle entry emitted |
| BagFilterTweak | [BagFilterTweak](addons/bagfiltertweak) | included | .mod semantic lifecycle entry emitted |
| BankArkel | [BankArkel](addons/bankarkel) | included | .mod semantic lifecycle entry emitted |
| BankWindowFix | [BankWindowFix](addons/bankwindowfix) | included | .mod semantic lifecycle entry emitted |
| BarText_Influence | [BarText (Influence)](addons/bartext_(influence)) | included | .mod semantic lifecycle entry emitted |
| Big'uns | [Big'uns](addons/big'uns) | included | .mod semantic lifecycle entry emitted |
| BlackBook | [BlackBook](addons/blackbook) | included | .mod semantic lifecycle entry emitted |
| BlackBox | [BlackBox](addons/blackbox) | included | .mod semantic lifecycle entry emitted |
| Bloody Mess | [Bloody Mess](addons/bloody_mess) | included | .mod semantic lifecycle entry emitted |
| Brizio's Crappy Computer Medic | [Brizio's Crappy Computer Medic](addons/brizio's_crappy_computer_medic) | included | .mod semantic lifecycle entry emitted |
| BuddyBind | [BuddyBind](addons/buddybind) | included | .mod semantic lifecycle entry emitted |
| BuffHead | [BuffHead](addons/buffhead) | included | .mod semantic lifecycle entry emitted |
| BuffThrottle | [BuffThrottle](addons/buffthrottle) | included | .mod semantic lifecycle entry emitted |
| Busted | [Busted](addons/busted) | included | .mod semantic lifecycle entry emitted |
| CCTV | [CCTV](addons/cctv) | included | .mod semantic lifecycle entry emitted |
| CDown | [CDown](addons/cdown) | included | .mod semantic lifecycle entry emitted |
| CNC | [CNC](addons/cnc) | included | .mod semantic lifecycle entry emitted |
| CaVES | [CaVES](addons/caves) | included | .mod semantic lifecycle entry emitted |
| Calling | [Calling](addons/calling) | included | .mod semantic lifecycle entry emitted |
| CastSequence | [CastSequence](addons/castsequence) | included | .mod semantic lifecycle entry emitted |
| CharacterScreenTabFix | [CharacterScreenTabFix](addons/characterscreentabfix) | included | .mod semantic lifecycle entry emitted |
| ChatAlert | [ChatAlert](addons/chatalert) | included | .mod semantic lifecycle entry emitted |
| ChattyCathy | [ChattyCathy](addons/chattycathy) | included | .mod semantic lifecycle entry emitted |
| Cheeseboard | [Cheeseboard](addons/cheeseboard) | included | .mod semantic lifecycle entry emitted |
| CleanCastbar | [CleanCastbar](addons/cleancastbar) | included | .mod semantic lifecycle entry emitted |
| CleanUnitFrames | [CleanUnitFrames](addons/cleanunitframes) | included | .mod semantic lifecycle entry emitted |
| CleansingBuddy | [CleansingBuddy](addons/cleansingbuddy) | included | .mod semantic lifecycle entry emitted |
| ClickSoundSuppressor | [ClickSoundSuppressor](addons/clicksoundsuppressor) | included | .mod semantic lifecycle entry emitted |
| Clock | [Clock](addons/clock) | included | .mod semantic lifecycle entry emitted |
| ClosetGoblin | ClosetGoblin | excluded | not present in .mod lifecycle semantic output |
| Compass3D | [Compass3D](addons/compass3d) | included | .mod semantic lifecycle entry emitted |
| CoolDownLine | [CoolDownLine](addons/cooldownline) | included | .mod semantic lifecycle entry emitted |
| CraftValueTip | CraftValueTip | excluded | not present in .mod lifecycle semantic output |
| CraftingWillard | [CraftingWillard](addons/craftingwillard) | included | .mod semantic lifecycle entry emitted |
| CramTheSpam | [Cram The Spam](addons/cram_the_spam) | included | .mod semantic lifecycle entry emitted |
| CurseProfiler | CurseProfiler | excluded | not present in .mod lifecycle semantic output |
| DAoCBuff | [DAoCBuff](addons/daocbuff) | included | .mod semantic lifecycle entry emitted |
| DaemonAssist | [DaemonAssist](addons/daemonassist) | included | .mod semantic lifecycle entry emitted |
| DammazKron | [DammazKron](addons/dammazkron) | included | .mod semantic lifecycle entry emitted |
| DasBoot | [DasBoot](addons/dasboot) | included | .mod semantic lifecycle entry emitted |
| Dascore | [Dascore](addons/dascore) | included | .mod semantic lifecycle entry emitted |
| Deathblow | [Deathblow](addons/deathblow) | included | .mod semantic lifecycle entry emitted |
| Deathblow2 | [Deathblow2](addons/deathblow2) | included | .mod semantic lifecycle entry emitted |
| DeepSleep | [DeepSleep](addons/deepsleep) | included | .mod semantic lifecycle entry emitted |
| DefaultTacticSet | [Default Tactic Set](addons/default_tactic_set) | included | .mod semantic lifecycle entry emitted |
| DetauntHelper | [DetauntHelper](addons/detaunthelper) | included | .mod semantic lifecycle entry emitted |
| Ding | [Ding](addons/ding) | included | .mod semantic lifecycle entry emitted |
| DisableCombatLog | [DisableCombatLog](addons/disablecombatlog) | included | .mod semantic lifecycle entry emitted |
| DuelInvite | [DuelInvite](addons/duelinvite) | included | .mod semantic lifecycle entry emitted |
| Dufftimer | [DuffTimer](addons/dufftimer) | included | .mod semantic lifecycle entry emitted |
| DwarfTalk | [DwarfTalk](addons/dwarftalk) | included | .mod semantic lifecycle entry emitted |
| Dye | Dye | excluded | not present in .mod lifecycle semantic output |
| EA_OpenPartyWindow | [EA_OpenPartyWindow](addons/ea_openpartywindow) | included | .mod semantic lifecycle entry emitted |
| EA_ScenarioGroupWindow | [EA_ScenarioGroupWindow](addons/ea_scenariogroupwindow) | included | .mod semantic lifecycle entry emitted |
| EZCraft | [EZCraft](addons/ezcraft) | included | .mod semantic lifecycle entry emitted |
| EZGuard | [EZGuard](addons/ezguard) | included | .mod semantic lifecycle entry emitted |
| Effigy | [Effigy](addons/effigy) | included | .mod semantic lifecycle entry emitted |
| Emojii | [Emojii](addons/emojii) | included | .mod semantic lifecycle entry emitted |
| EmoteAlert | [EmoteAlert](addons/emotealert) | included | .mod semantic lifecycle entry emitted |
| Enemy | [Enemy](addons/enemy) | included | .mod semantic lifecycle entry emitted |
| EveryBodyGuard | [EveryBodyGuard](addons/everybodyguard) | included | .mod semantic lifecycle entry emitted |
| FastFriends | [FastFriends](addons/fastfriends) | included | .mod semantic lifecycle entry emitted |
| FastInteract | [FastInteract](addons/fastinteract) | included | .mod semantic lifecycle entry emitted |
| FightFinder | [Fight Finder](addons/fight_finder) | included | .mod semantic lifecycle entry emitted |
| FlagCap | [FlagCap](addons/flagcap) | included | .mod semantic lifecycle entry emitted |
| FozAuction | [FozAuction](addons/fozauction) | included | .mod semantic lifecycle entry emitted |
| GCDsaver | [GCDsaver](addons/gcdsaver) | included | .mod semantic lifecycle entry emitted |
| GDes | [GDes](addons/gdes) | included | .mod semantic lifecycle entry emitted |
| Ges | [Ges](addons/ges) | included | .mod semantic lifecycle entry emitted |
| GetStats | [GetStats](addons/getstats) | included | .mod semantic lifecycle entry emitted |
| GroupIcons | [Group Icons](addons/group_icons) | included | .mod semantic lifecycle entry emitted |
| GroupRange | [GroupRange](addons/grouprange) | included | .mod semantic lifecycle entry emitted |
| GroupSpotter | [GroupSpotter](addons/groupspotter) | included | .mod semantic lifecycle entry emitted |
| GuardBot | [GuardBot](addons/guardbot) | included | .mod semantic lifecycle entry emitted |
| GuardLine | [GuardLine](addons/guardline) | included | .mod semantic lifecycle entry emitted |
| GuardList | [GuardList](addons/guardlist) | included | .mod semantic lifecycle entry emitted |
| GuardPack | GuardPack | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |
| GuardRange | [GuardRange](addons/guardrange) | included | .mod semantic lifecycle entry emitted |
| GuildWarden | [GuildWarden](addons/guildwarden) | included | .mod semantic lifecycle entry emitted |
| HealAll | [HealAll](addons/healall) | included | .mod semantic lifecycle entry emitted |
| HealGrid | [HealGrid](addons/healgrid) | included | .mod semantic lifecycle entry emitted |
| HealHoverAssist | [HealHoverAssist](addons/healhoverassist) | included | .mod semantic lifecycle entry emitted |
| HideHiddenFrames | [HideHiddenFrames](addons/hidehiddenframes) | included | .mod semantic lifecycle entry emitted |
| HotbarMorale | [HotbarMorale](addons/hotbarmorale) | included | .mod semantic lifecycle entry emitted |
| IHYTM | IHYTM | excluded | not present in .mod lifecycle semantic output |
| IdentityFound | [IdentityFound](addons/identityfound) | included | .mod semantic lifecycle entry emitted |
| InfoScroller | [InfoScroller](addons/infoscroller) | included | .mod semantic lifecycle entry emitted |
| Info_Alert | [Info_Alert](addons/info_alert) | included | .mod semantic lifecycle entry emitted |
| Info_DeathBlow | [Info_DeathBlow](addons/info_deathblow) | included | .mod semantic lifecycle entry emitted |
| Info_Loot | [Info_Loot](addons/info_loot) | included | .mod semantic lifecycle entry emitted |
| Info_Points | [Info_Points](addons/info_points) | included | .mod semantic lifecycle entry emitted |
| ItemRack | [ItemRack](addons/itemrack) | included | .mod semantic lifecycle entry emitted |
| JunkDump | [JunkDump](addons/junkdump) | included | .mod semantic lifecycle entry emitted |
| KeyBar | [KeyBar](addons/keybar) | included | .mod semantic lifecycle entry emitted |
| Keyset | [Keyset](addons/keyset) | included | .mod semantic lifecycle entry emitted |
| KeysetMonsterPlay | [KeysetMonsterPlay](addons/keysetmonsterplay) | included | .mod semantic lifecycle entry emitted |
| KillStreak | [KillStreak](addons/killstreak) | included | .mod semantic lifecycle entry emitted |
| KillTracker | [KillTracker](addons/killtracker) | included | .mod semantic lifecycle entry emitted |
| Killer | [Killer](addons/killer) | included | .mod semantic lifecycle entry emitted |
| Kwestor | [Kwestor](addons/kwestor) | included | .mod semantic lifecycle entry emitted |
| LibAchievements | [LibAchievements](addons/libachievements) | included | .mod semantic lifecycle entry emitted |
| LibAddonButton | [LibAddonButton](addons/libaddonbutton) | included | .mod semantic lifecycle entry emitted |
| LibCereal-1.0 | [LibCereal-1.0](addons/libcereal-1.0) | included | .mod semantic lifecycle entry emitted |
| LibCompactCurrentDateTime | [LibCompactCurrentDateTime](addons/libcompactcurrentdatetime) | included | .mod semantic lifecycle entry emitted |
| LibCooldown | [LibCooldown](addons/libcooldown) | included | .mod semantic lifecycle entry emitted |
| LibDateTime | [LibDateTime](addons/libdatetime) | included | .mod semantic lifecycle entry emitted |
| LibGroup | [LibGroup](addons/libgroup) | included | .mod semantic lifecycle entry emitted |
| LibGuard | [LibGuard](addons/libguard) | included | .mod semantic lifecycle entry emitted |
| LibJson | [LibJson](addons/libjson) | included | .mod semantic lifecycle entry emitted |
| LibPickle | [LibPickle](addons/libpickle) | included | .mod semantic lifecycle entry emitted |
| LibRange | [LibRange](addons/librange) | included | .mod semantic lifecycle entry emitted |
| LibSkillicon | [LibSkillicon](addons/libskillicon) | included | .mod semantic lifecycle entry emitted |
| LibSlash | [LibSlash](addons/libslash) | included | .mod semantic lifecycle entry emitted |
| LibStub | [LibStub](addons/libstub) | included | .mod semantic lifecycle entry emitted |
| LibSurveyor | [LibSurveyor](addons/libsurveyor) | included | .mod semantic lifecycle entry emitted |
| LibWarBoardToggler | LibWarBoardToggler | excluded | not present in .mod lifecycle semantic output |
| LootAlert | [LootAlert](addons/lootalert) | included | .mod semantic lifecycle entry emitted |
| LoyalPet | [LoyalPet](addons/loyalpet) | included | .mod semantic lifecycle entry emitted |
| MGRemix | MGRemix | excluded | not present in .mod lifecycle semantic output |
| MacroIcons | [MacroIcons](addons/macroicons) | included | .mod semantic lifecycle entry emitted |
| Map | [Map](addons/map) | included | .mod semantic lifecycle entry emitted |
| MapMonster | [MapMonster](addons/mapmonster) | included | .mod semantic lifecycle entry emitted |
| MapPin | [MapPin](addons/mappin) | included | .mod semantic lifecycle entry emitted |
| MarkBuff | [MarkBuff](addons/markbuff) | included | .mod semantic lifecycle entry emitted |
| MassRefine | [Mass Refine](addons/mass_refine) | included | .mod semantic lifecycle entry emitted |
| Mech | [Mech](addons/mech) | included | .mod semantic lifecycle entry emitted |
| MegaphonePlus | [MegaphonePlus](addons/megaphoneplus) | included | .mod semantic lifecycle entry emitted |
| MegaphonePlusPlus | [MegaphonePlusPlus](addons/megaphoneplusplus) | included | .mod semantic lifecycle entry emitted |
| MiracleGrow | [MiracleGrow](addons/miraclegrow) | included | .mod semantic lifecycle entry emitted |
| MiracleGrowLight | [MiracleGrowLight](addons/miraclegrowlight) | included | .mod semantic lifecycle entry emitted |
| MoraleCircle | [MoraleCircle](addons/moralecircle) | included | .mod semantic lifecycle entry emitted |
| MoraleSet | [MoraleSet](addons/moraleset) | included | .mod semantic lifecycle entry emitted |
| Moth | [Moth](addons/moth) | included | .mod semantic lifecycle entry emitted |
| MyReasons | [MyReasons](addons/myreasons) | included | .mod semantic lifecycle entry emitted |
| NAMBLA | [NAMBLA](addons/nambla) | included | .mod semantic lifecycle entry emitted |
| NaturalLog | [NaturalLog](addons/naturallog) | included | .mod semantic lifecycle entry emitted |
| NerfedButtons | [NerfedButtons](addons/nerfedbuttons) | included | .mod semantic lifecycle entry emitted |
| NervAlert | [NervAlert](addons/nervalert) | included | .mod semantic lifecycle entry emitted |
| ObjectInspector | [ObjectInspector](addons/objectinspector) | included | .mod semantic lifecycle entry emitted |
| Obsidian | [Obsidian](addons/obsidian) | included | .mod semantic lifecycle entry emitted |
| OilTimer | [OilTimer](addons/oiltimer) | included | .mod semantic lifecycle entry emitted |
| OneClickGTAOE | [OneClickGTAOE](addons/oneclickgtaoe) | included | .mod semantic lifecycle entry emitted |
| OverheadFonts | [OverheadFonts](addons/overheadfonts) | included | .mod semantic lifecycle entry emitted |
| PaintTheLeader-Fenris | PaintTheLeader-Fenris | excluded | not present in .mod lifecycle semantic output |
| PartyAd | [PartyAd](addons/partyad) | included | .mod semantic lifecycle entry emitted |
| PartyCast | [PartyCast](addons/partycast) | included | .mod semantic lifecycle entry emitted |
| PeaceOut | [PeaceOut](addons/peaceout) | included | .mod semantic lifecycle entry emitted |
| PetAssist | [PetAssist](addons/petassist) | included | .mod semantic lifecycle entry emitted |
| PetFixWindow | [PetFixWindow](addons/petfixwindow) | included | .mod semantic lifecycle entry emitted |
| Phantom | [Phantom](addons/phantom) | included | .mod semantic lifecycle entry emitted |
| PieTracker | [PieTracker](addons/pietracker) | included | .mod semantic lifecycle entry emitted |
| PlanB | [PlanB](addons/planb) | included | .mod semantic lifecycle entry emitted |
| PlayEffectsOn | [PlayEffectsOn](addons/playeffectson) | included | .mod semantic lifecycle entry emitted |
| PocketPalette | [Pocket Palette](addons/pocket_palette) | included | .mod semantic lifecycle entry emitted |
| PotionBar | [PotionBar](addons/potionbar) | included | .mod semantic lifecycle entry emitted |
| Pure | [Pure](addons/pure) | included | .mod semantic lifecycle entry emitted |
| PureCareerBar | [Pure Careerbar](addons/pure_careerbar) | included | .mod semantic lifecycle entry emitted |
| QueuePopTimer | [QueuePopTimer](addons/queuepoptimer) | included | .mod semantic lifecycle entry emitted |
| QueueQueuer | [Queue Queuer](addons/queue_queuer) | included | .mod semantic lifecycle entry emitted |
| QuickNameActions+ | [QuickNameActions+](addons/quicknameactions+) | included | .mod semantic lifecycle entry emitted |
| QuickPerf | QuickPerf | excluded | not present in .mod lifecycle semantic output |
| QuickTacticSwitch | [QuickTacticSwitch](addons/quicktacticswitch) | included | .mod semantic lifecycle entry emitted |
| QuickWarReport | [QuickWarReport](addons/quickwarreport) | included | .mod semantic lifecycle entry emitted |
| ROCT | ROCT | excluded | not present in .mod lifecycle semantic output |
| RVAPI_ColorDialog | [RVAPI_ColorDialog](addons/rvapi_colordialog) | included | .mod semantic lifecycle entry emitted |
| RVAPI_Range | [RVAPI_Range](addons/rvapi_range) | included | .mod semantic lifecycle entry emitted |
| RVMOD_3DPortrait | [RVMOD_3DPortrait](addons/rvmod_3dportrait) | included | .mod semantic lifecycle entry emitted |
| RVMOD_Manager | [RVMOD_Manager](addons/rvmod_manager) | included | .mod semantic lifecycle entry emitted |
| RVMOD_PlayerStatus | [RVMOD_PlayerStatus](addons/rvmod_playerstatus) | included | .mod semantic lifecycle entry emitted |
| RVMOD_SquaredDistances | [RVMOD_SquaredDistances](addons/rvmod_squareddistances) | included | .mod semantic lifecycle entry emitted |
| RVMOD_Targets | [RVMOD_Targets](addons/rvmod_targets) | included | .mod semantic lifecycle entry emitted |
| RaidMeter | [RaidMeter](addons/raidmeter) | included | .mod semantic lifecycle entry emitted |
| RandomMount | [RandomMount](addons/randommount) | included | .mod semantic lifecycle entry emitted |
| RandomSayings | [RandomSayings](addons/randomsayings) | included | .mod semantic lifecycle entry emitted |
| RangeFadingSquared | [RangeFadingSquared](addons/rangefadingsquared) | included | .mod semantic lifecycle entry emitted |
| Rangechecker | [Rangechecker](addons/rangechecker) | included | .mod semantic lifecycle entry emitted |
| RealmStatus | [RealmStatus](addons/realmstatus) | included | .mod semantic lifecycle entry emitted |
| Refer | [Refer](addons/refer) | included | .mod semantic lifecycle entry emitted |
| ReliquaryHunter | [ReliquaryHunter](addons/reliquaryhunter) | included | .mod semantic lifecycle entry emitted |
| RememberIt | [RememberIt](addons/rememberit) | included | .mod semantic lifecycle entry emitted |
| Res | [Res](addons/res) | included | .mod semantic lifecycle entry emitted |
| ResHelp | [ResHelp](addons/reshelp) | included | .mod semantic lifecycle entry emitted |
| RetAlert | [RetAlert](addons/retalert) | included | .mod semantic lifecycle entry emitted |
| RezEmote | [RezEmote](addons/rezemote) | included | .mod semantic lifecycle entry emitted |
| RoRAutoCapeHelm | RoRAutoCapeHelm | excluded | not present in .mod lifecycle semantic output |
| RoR_SoR | [RoR_SoR](addons/ror_sor) | included | .mod semantic lifecycle entry emitted |
| RoR_debolster | [RoR_debolster](addons/ror_debolster) | included | .mod semantic lifecycle entry emitted |
| RuStringLib | RuStringLib | excluded | not present in .mod lifecycle semantic output |
| RvRStatsTab | [RvRStatsTab](addons/rvrstatstab) | included | .mod semantic lifecycle entry emitted |
| SB_TDPS | [SB_TDPS](addons/sb_tdps) | included | .mod semantic lifecycle entry emitted |
| SNT_BUTTONS | [SNT_BUTTONS](addons/snt_buttons) | included | .mod semantic lifecycle entry emitted |
| SNT_CASTBAR | [SNT_CASTBAR](addons/snt_castbar) | included | .mod semantic lifecycle entry emitted |
| SNT_INFO | [SNT_INFO](addons/snt_info) | included | .mod semantic lifecycle entry emitted |
| SNT_INFO_TDPS | [SNT Info [TDPS]](addons/snt_info_[tdps]) | included | .mod semantic lifecycle entry emitted |
| SNT_PANEL | [SNT_PANEL](addons/snt_panel) | included | .mod semantic lifecycle entry emitted |
| ScenarioAlert | [ScenarioAlert](addons/scenarioalert) | included | .mod semantic lifecycle entry emitted |
| ScenarioStats | [ScenarioStats](addons/scenariostats) | included | .mod semantic lifecycle entry emitted |
| SelfTarget | [SelfTarget](addons/selftarget) | included | .mod semantic lifecycle entry emitted |
| Sequencer | [Sequencer](addons/sequencer) | included | .mod semantic lifecycle entry emitted |
| Shinies | [Shinies](addons/shinies) | included | .mod semantic lifecycle entry emitted |
| ShowHealthPercent | [ShowHealthPercent](addons/showhealthpercent) | included | .mod semantic lifecycle entry emitted |
| SimpleXY | [SimpleXY](addons/simplexy) | included | .mod semantic lifecycle entry emitted |
| SocialWindow 2.0 | [SocialWindow 2.0](addons/socialwindow_2.0) | included | .mod semantic lifecycle entry emitted |
| Soloq | [Soloq](addons/soloq) | included | .mod semantic lifecycle entry emitted |
| SpamBayes | [SpamBayes](addons/spambayes) | included | .mod semantic lifecycle entry emitted |
| Squared | [Squared](addons/squared) | included | .mod semantic lifecycle entry emitted |
| SquaredHDIndicator | [SquaredHDIndicator](addons/squaredhdindicator) | included | .mod semantic lifecycle entry emitted |
| Statdoll | [Statdoll](addons/statdoll) | included | .mod semantic lifecycle entry emitted |
| StatdollLight | [Statdoll Light](addons/statdoll_light) | included | .mod semantic lifecycle entry emitted |
| StatdollRemix | [Statdoll Remix](addons/statdoll_remix) | included | .mod semantic lifecycle entry emitted |
| StateOfRealm | StateOfRealm | excluded | not present in .mod lifecycle semantic output |
| StopRes | [StopRes](addons/stopres) | included | .mod semantic lifecycle entry emitted |
| Swinger | [Swinger](addons/swinger) | included | .mod semantic lifecycle entry emitted |
| TacticSetNames | [TacticSetNames](addons/tacticsetnames) | included | .mod semantic lifecycle entry emitted |
| TalismanGenie | [TalismanGenie](addons/talismangenie) | included | .mod semantic lifecycle entry emitted |
| TastyButtons | [TastyButtons](addons/tastybuttons) | included | .mod semantic lifecycle entry emitted |
| TaxPayer | [TaxPayer](addons/taxpayer) | included | .mod semantic lifecycle entry emitted |
| TexturedButtons | [TexturedButtons](addons/texturedbuttons) | included | .mod semantic lifecycle entry emitted |
| ThankTheResser | [ThankTheResser](addons/thanktheresser) | included | .mod semantic lifecycle entry emitted |
| TheSeeker | [TheSeeker](addons/theseeker) | included | .mod semantic lifecycle entry emitted |
| TidyChat | [TidyChat](addons/tidychat) | included | .mod semantic lifecycle entry emitted |
| TidyQueue | [TidyQueue](addons/tidyqueue) | included | .mod semantic lifecycle entry emitted |
| TidyRoll | [TidyRoll](addons/tidyroll) | included | .mod semantic lifecycle entry emitted |
| TimeToDie | [TimeToDie](addons/timetodie) | included | .mod semantic lifecycle entry emitted |
| TokenMachine | [TokenMachine](addons/tokenmachine) | included | .mod semantic lifecycle entry emitted |
| TomeTitan | [Tome Titan](addons/tome_titan) | included | .mod semantic lifecycle entry emitted |
| TomeTracker | [TomeTracker](addons/tometracker) | included | .mod semantic lifecycle entry emitted |
| TortallDPSCore | [TortallDPSCore](addons/tortalldpscore) | included | .mod semantic lifecycle entry emitted |
| Tortall_DPS | [Tortall_DPS](addons/tortall_dps) | included | .mod semantic lifecycle entry emitted |
| Tortall_SCC | [Tortall_SCC](addons/tortall_scc) | included | .mod semantic lifecycle entry emitted |
| TurrentRange | TurrentRange | excluded | not present in .mod lifecycle semantic output |
| TurretScrap | [TurretScrap](addons/turretscrap) | included | .mod semantic lifecycle entry emitted |
| Twister | [Twister](addons/twister) | included | .mod semantic lifecycle entry emitted |
| TwisterSet | [TwisterSet](addons/twisterset) | included | .mod semantic lifecycle entry emitted |
| UnrealDBAnnouncer | [UnrealDBAnnouncer](addons/unrealdbannouncer) | included | .mod semantic lifecycle entry emitted |
| VPBreakdown | [VPBreakdown](addons/vpbreakdown) | included | .mod semantic lifecycle entry emitted |
| Vectors | [Vectors](addons/vectors) | included | .mod semantic lifecycle entry emitted |
| VerticalMorale | [VerticalMorale](addons/verticalmorale) | included | .mod semantic lifecycle entry emitted |
| VerticalTactics | [VerticalTactics](addons/verticaltactics) | included | .mod semantic lifecycle entry emitted |
| Vertigo | [Vertigo](addons/vertigo) | included | .mod semantic lifecycle entry emitted |
| WARRatingBuster | [WARRatingBuster](addons/warratingbuster) | included | .mod semantic lifecycle entry emitted |
| WBStutterLess | [WBStutterLess](addons/wbstutterless) | included | .mod semantic lifecycle entry emitted |
| WTes | [WTes](addons/wtes) | included | .mod semantic lifecycle entry emitted |
| WarBoard | [WarBoard](addons/warboard) | included | .mod semantic lifecycle entry emitted |
| WarBoard_FPS | [WarBoard_FPS](addons/warboard_fps) | included | .mod semantic lifecycle entry emitted |
| WarBoard_Loc | [WarBoard_Loc](addons/warboard_loc) | included | .mod semantic lifecycle entry emitted |
| WarBoard_Menu | [WarBoard_Menu](addons/warboard_menu) | included | .mod semantic lifecycle entry emitted |
| WarBoard_Session | [WarBoard_Session](addons/warboard_session) | included | .mod semantic lifecycle entry emitted |
| WarBoard_TDPS | [WarBoard_TDPS](addons/warboard_tdps) | included | .mod semantic lifecycle entry emitted |
| WarBoard_TogglerHealGrid | [WarBoard_TogglerHealGrid](addons/warboard_togglerhealgrid) | included | .mod semantic lifecycle entry emitted |
| WarBoard_TogglerRankedLeaderboard | [WarBoard_TogglerRankedLeaderboard](addons/warboard_togglerrankedleaderboard) | included | .mod semantic lifecycle entry emitted |
| WarBoard_TogglerSquared | [WarBoard_TogglerSquared](addons/warboard_togglersquared) | included | .mod semantic lifecycle entry emitted |
| WarBoard_TogglerTidyChat | [WarBoard_TogglerTidyChat](addons/warboard_togglertidychat) | included | .mod semantic lifecycle entry emitted |
| WarBoard_TogglerTokenMachine | [WarBoard_TogglerTokenMachine](addons/warboard_togglertokenmachine) | included | .mod semantic lifecycle entry emitted |
| WarBoard_TogglerVPBreakdown | [WarBoard_TogglerVPBreakdown](addons/warboard_togglervpbreakdown) | included | .mod semantic lifecycle entry emitted |
| WarBoard_TogglerWARCommander | [WarBoard_TogglerWARCommander](addons/warboard_togglerwarcommander) | included | .mod semantic lifecycle entry emitted |
| WarTriage | [WarTriage](addons/wartriage) | included | .mod semantic lifecycle entry emitted |
| War_RU | [War_RU](addons/war_ru) | included | .mod semantic lifecycle entry emitted |
| Warbuilder | [Warbuilder](addons/warbuilder) | included | .mod semantic lifecycle entry emitted |
| WhatsCooking | [WhatsCooking](addons/whatscooking) | included | .mod semantic lifecycle entry emitted |
| WhoHealedMe | [WhoHealedMe](addons/whohealedme) | included | .mod semantic lifecycle entry emitted |
| WikkisCooldownBar | [Wikki's Cooldown Bar](addons/wikki's_cooldown_bar) | included | .mod semantic lifecycle entry emitted |
| WikkisCooldownPulse | [Wikki's Cooldown Pulse](addons/wikki's_cooldown_pulse) | included | .mod semantic lifecycle entry emitted |
| WoH-Reticle | [WoH-Reticle](addons/woh-reticle) | included | .mod semantic lifecycle entry emitted |
| WoodysDpsLib | [WoodysDpsLib](addons/woodysdpslib) | included | .mod semantic lifecycle entry emitted |
| XpStatus plus G | XpStatus plus G | excluded | not present in .mod lifecycle semantic output |
| ZonePOP | [ZonePOP](addons/zonepop) | included | .mod semantic lifecycle entry emitted |
| abh | abh | excluded | not present in .mod lifecycle semantic output |
| advancedrenowntrainer | [AdvancedRenownTrainer](addons/advancedrenowntrainer) | included | .mod semantic lifecycle entry emitted |
| alertMod | [alertMod](addons/alertmod) | included | .mod semantic lifecycle entry emitted |
| amethyst | [Amethyst](addons/amethyst) | included | .mod semantic lifecycle entry emitted |
| apw | apw | excluded | not present in .mod lifecycle semantic output |
| assist | [Assist](addons/assist) | included | .mod semantic lifecycle entry emitted |
| au-stats | au-stats | excluded | not present in .mod lifecycle semantic output |
| autochannel | [AutoChannel](addons/autochannel) | included | .mod semantic lifecycle entry emitted |
| autosalvage | [AutoSalvage](addons/autosalvage) | included | .mod semantic lifecycle entry emitted |
| bagomatic | [BagOMatic](addons/bagomatic) | included | .mod semantic lifecycle entry emitted |
| bettercc | [BetterCC](addons/bettercc) | included | .mod semantic lifecycle entry emitted |
| bigger_macrowindow | [bigger_MacroWindow](addons/bigger_macrowindow) | included | .mod semantic lifecycle entry emitted |
| careful-core-functions | [Careful Core Functions](addons/careful_core_functions) | included | .mod semantic lifecycle entry emitted |
| cmap | [CMap](addons/cmap) | included | .mod semantic lifecycle entry emitted |
| combattextnames | [CombatTextNames](addons/combattextnames) | included | .mod semantic lifecycle entry emitted |
| comp3d | comp3d | excluded | not present in .mod lifecycle semantic output |
| countdown | [Countdown](addons/countdown) | included | .mod semantic lifecycle entry emitted |
| crusher | [Crusher](addons/crusher) | included | .mod semantic lifecycle entry emitted |
| demonicore | [DeMoNiCore](addons/demonicore) | included | .mod semantic lifecycle entry emitted |
| dpsmeter | [DPSMeter](addons/dpsmeter) | included | .mod semantic lifecycle entry emitted |
| duel | [Duel](addons/duel) | included | .mod semantic lifecycle entry emitted |
| ea_loadingscreen | [EA_LoadingScreen](addons/ea_loadingscreen) | included | .mod semantic lifecycle entry emitted |
| ea_uidebugtools | [EA_UiDebugTools](addons/ea_uidebugtools) | included | .mod semantic lifecycle entry emitted |
| ea_uimodwindow | [EA_UiModWindow](addons/ea_uimodwindow) | included | .mod semantic lifecycle entry emitted |
| easystem_threepartbar | easystem_threepartbar | excluded | not present in .mod lifecycle semantic output |
| emotes | [emotes](addons/emotes) | included | .mod semantic lifecycle entry emitted |
| ezcraftx | [EZCraftX](addons/ezcraftx) | included | .mod semantic lifecycle entry emitted |
| fixgit | [FixGit](addons/fixgit) | included | .mod semantic lifecycle entry emitted |
| followTheLeader | [followTheLeader](addons/followtheleader) | included | .mod semantic lifecycle entry emitted |
| fpsbox | [fpsbox](addons/fpsbox) | included | .mod semantic lifecycle entry emitted |
| gcd-tracker | [GCDTracker](addons/gcdtracker) | included | .mod semantic lifecycle entry emitted |
| greedy | [Greedy](addons/greedy) | included | .mod semantic lifecycle entry emitted |
| groupiconsSG | [Group Icons SG](addons/group_icons_sg) | included | .mod semantic lifecycle entry emitted |
| guildwindowexpanded-1-5-5 | guildwindowexpanded-1-5-5 | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |
| hammertime | [HammerTime](addons/hammertime) | included | .mod semantic lifecycle entry emitted |
| harbinger.1-6-1 | harbinger.1-6-1 | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |
| hideInf | [hideInf](addons/hideinf) | included | .mod semantic lifecycle entry emitted |
| hopper | [Hopper](addons/hopper) | included | .mod semantic lifecycle entry emitted |
| manualscenariorefresh | [ManualScenarioRefresh](addons/manualscenariorefresh) | included | .mod semantic lifecycle entry emitted |
| minesweep | [minesweep](addons/minesweep) | included | .mod semantic lifecycle entry emitted |
| minimapmonster | [MiniMapMonster](addons/minimapmonster) | included | .mod semantic lifecycle entry emitted |
| moralebg | [MoraleBG](addons/moralebg) | included | .mod semantic lifecycle entry emitted |
| motion | [Motion](addons/motion) | included | .mod semantic lifecycle entry emitted |
| mouse-hint | [MouseHint](addons/mousehint) | included | .mod semantic lifecycle entry emitted |
| nLootLink | [nLootLink](addons/nlootlink) | included | .mod semantic lifecycle entry emitted |
| nRarity | [nRarity](addons/nrarity) | included | .mod semantic lifecycle entry emitted |
| nisp | nisp | excluded | not present in .mod lifecycle semantic output |
| no-useless-mods | no-useless-mods | excluded | not present in .mod lifecycle semantic output |
| nooverheal | [NoOverheal](addons/nooverheal) | included | .mod semantic lifecycle entry emitted |
| orcanizer | orcanizer | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |
| ovdeadnomore | [ovdeadnomore](addons/ovdeadnomore) | included | .mod semantic lifecycle entry emitted |
| pagerwentpoof | [PagerWentPoof](addons/pagerwentpoof) | included | .mod semantic lifecycle entry emitted |
| pmtb | [PMTB](addons/pmtb) | included | .mod semantic lifecycle entry emitted |
| preciousss | [Preciousss](addons/preciousss) | included | .mod semantic lifecycle entry emitted |
| quest-begone | [QuestBegone](addons/questbegone) | included | .mod semantic lifecycle entry emitted |
| rolodex | [Rolodex](addons/rolodex) | included | .mod semantic lifecycle entry emitted |
| ror-autoinviter | [rorAutoInviter](addons/rorautoinviter) | included | .mod semantic lifecycle entry emitted |
| rotation | [Rotation](addons/rotation) | included | .mod semantic lifecycle entry emitted |
| rp | rp | excluded | not present in .mod lifecycle semantic output |
| rvr-contribution | [RvRContribution](addons/rvrcontribution) | included | .mod semantic lifecycle entry emitted |
| rvrstatsupdated | rvrstatsupdated | excluded | not present in .mod lifecycle semantic output |
| scenarioInfo | [scenarioInfo](addons/scenarioinfo) | included | .mod semantic lifecycle entry emitted |
| scnoload | [scnoload](addons/scnoload) | included | .mod semantic lifecycle entry emitted |
| showmethebubbles | [ShowMeTheBubbles](addons/showmethebubbles) | included | .mod semantic lifecycle entry emitted |
| simple-combat-text | [SimpleCombatText](addons/simplecombattext) | included | .mod semantic lifecycle entry emitted |
| squaredclick | [SquaredClick](addons/squaredclick) | included | .mod semantic lifecycle entry emitted |
| squaredhotindicators | [SquaredHotIndicators](addons/squaredhotindicators) | included | .mod semantic lifecycle entry emitted |
| srp | srp | excluded | not present in .mod lifecycle semantic output |
| swift-assist | [Swift Assist](addons/swift_assist) | included | .mod semantic lifecycle entry emitted |
| talisman-monitor | [talisman-monitor](addons/talisman-monitor) | included | .mod semantic lifecycle entry emitted |
| targetinforing | [TargetInfoRing](addons/targetinforing) | included | .mod semantic lifecycle entry emitted |
| targetring | [TargetRing](addons/targetring) | included | .mod semantic lifecycle entry emitted |
| targets | [Targets](addons/targets) | included | .mod semantic lifecycle entry emitted |
| thank-the-tank | [ThankTheTank](addons/thankthetank) | included | .mod semantic lifecycle entry emitted |
| thinkoutloud | [ThinkOutLoud](addons/thinkoutloud) | included | .mod semantic lifecycle entry emitted |
| timeinqueue-1-1-1 | timeinqueue-1-1-1 | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |
| tokens | [Tokens](addons/tokens) | included | .mod semantic lifecycle entry emitted |
| trakario | [Trakario](addons/trakario) | included | .mod semantic lifecycle entry emitted |
| unbearable-1-0-0 | unbearable-1-0-0 | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |
| waaghbar | waaghbar | excluded | not present in .mod lifecycle semantic output |
| warboard_TogglerWarBuilder | [warboard_TogglerWarBuilder](addons/warboard_togglerwarbuilder) | included | .mod semantic lifecycle entry emitted |
| warboard_TogglerwWlh | warboard_TogglerwWlh | excluded | not present in .mod lifecycle semantic output |
| warboard_crests-0-5-1 | warboard_crests-0-5-1 | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |
| warboard_talimon | [WarBoard_TaliMon](addons/warboard_talimon) | included | .mod semantic lifecycle entry emitted |
| warboard_toggler_pack-1-0-0 | warboard_toggler_pack-1-0-0 | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |
| warboard_togglertidyroll | [WarBoard_TogglerTidyRoll](addons/warboard_togglertidyroll) | included | .mod semantic lifecycle entry emitted |
| warboard_togglerwsct | [WarBoard_TogglerWSCT](addons/warboard_togglerwsct) | included | .mod semantic lifecycle entry emitted |
| warcommander | [WARCommander](addons/warcommander) | included | .mod semantic lifecycle entry emitted |
| wargames | [Wargames](addons/wargames) | included | .mod semantic lifecycle entry emitted |
| warwhisperer | warwhisperer | excluded | not present in .mod lifecycle semantic output |
| wbLeadHelper | [wbLeadHelper](addons/wbleadhelper) | included | .mod semantic lifecycle entry emitted |
| whatsPugSc | [whatsPugSc](addons/whatspugsc) | included | .mod semantic lifecycle entry emitted |
| whom | [whom](addons/whom) | included | .mod semantic lifecycle entry emitted |
| windowmovers | [WindowMovers](addons/windowmovers) | included | .mod semantic lifecycle entry emitted |
| wsct | [WSCT](addons/wsct) | included | .mod semantic lifecycle entry emitted |
| xHUD | [xHUD](addons/xhud) | included | .mod semantic lifecycle entry emitted |
| xpanels | [xPanels](addons/xpanels) | included | .mod semantic lifecycle entry emitted |
| yAssistHelper | [yAssistHelper](addons/yassisthelper) | included | .mod semantic lifecycle entry emitted |
| yakui-1-3-6 | yakui-1-3-6 | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |
| zMailMod | [zMailMod](addons/zmailmod) | included | .mod semantic lifecycle entry emitted |
| zTimeLib | zTimeLib | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |
| zmm | zmm | excluded | not present in .mod lifecycle semantic output |
| zonepop1-3-1-3-0 | zonepop1-3-1-3-0 | excluded | no_manifest: directory has no .mod or .toc manifest and is excluded at source discovery |

## Lifecycle Addons (.mod Semantic Output)

- [AbilityAlert](addons/abilityalert)
- [AbilityNotifier](addons/abilitynotifier)
- [Ace](addons/ace)
- [ActionBarCD](addons/actionbarcd)
- [ActionBarColor](addons/actionbarcolor)
- [ActionBarHide](addons/actionbarhide)
- [ActionFraction](addons/actionfraction)
- [ActionPointWatch](addons/actionpointwatch)
- [ActionPoints](addons/actionpoints)
- [AddonButtonTogglers](addons/addonbuttontogglers)
- [AdjustTheTip](addons/adjustthetip)
- [AdvancedPetAssist](addons/advancedpetassist)
- [AdvancedRenownTrainer](addons/advancedrenowntrainer)
- [AggroMeter](addons/aggrometer)
- [Amethyst](addons/amethyst)
- [AnywhereTrainer](addons/anywheretrainer)
- [AnywhereTrainerAdditions](addons/anywheretraineradditions)
- [ArmorGraphicToggle](addons/armorgraphictoggle)
- [Arsenal Rank](addons/arsenal_rank)
- [Asshat](addons/asshat)
- [Assist](addons/assist)
- [Atlas](addons/atlas)
- [AuctionStats](addons/auctionstats)
- [AuctionTweaker](addons/auctiontweaker)
- [Aura](addons/aura)
- [Auto Cape/Helm Show](addons/auto_cape/helm_show)
- [AutoBand](addons/autoband)
- [AutoChannel](addons/autochannel)
- [AutoDismount](addons/autodismount)
- [AutoFocus](addons/autofocus)
- [AutoMark](addons/automark)
- [AutoSalvage](addons/autosalvage)
- [BBars - Mechanic Only](addons/bbars_-_mechanic_only)
- [BWMT](addons/bwmt)
- [BagFilterTweak](addons/bagfiltertweak)
- [BagOMatic](addons/bagomatic)
- [BankArkel](addons/bankarkel)
- [BankWindowFix](addons/bankwindowfix)
- [BarText (Influence)](addons/bartext_(influence))
- [BetterCC](addons/bettercc)
- [Big'uns](addons/big'uns)
- [BlackBook](addons/blackbook)
- [BlackBox](addons/blackbox)
- [Bloody Mess](addons/bloody_mess)
- [Brizio's Crappy Computer Medic](addons/brizio's_crappy_computer_medic)
- [BuddyBind](addons/buddybind)
- [BuffHead](addons/buffhead)
- [BuffThrottle](addons/buffthrottle)
- [Busted](addons/busted)
- [CCTV](addons/cctv)
- [CDown](addons/cdown)
- [CM_ClosetGoblin](addons/cm_closetgoblin)
- [CMap](addons/cmap)
- [CNC](addons/cnc)
- [CaVES](addons/caves)
- [Calling](addons/calling)
- [Careful Core Functions](addons/careful_core_functions)
- [CastSequence](addons/castsequence)
- [CharacterScreenTabFix](addons/characterscreentabfix)
- [ChatAlert](addons/chatalert)
- [ChattyCathy](addons/chattycathy)
- [Cheeseboard](addons/cheeseboard)
- [CleanCastbar](addons/cleancastbar)
- [CleanUnitFrames](addons/cleanunitframes)
- [CleansingBuddy](addons/cleansingbuddy)
- [ClickSoundSuppressor](addons/clicksoundsuppressor)
- [Clock](addons/clock)
- [CombatTextNames](addons/combattextnames)
- [Compass3D](addons/compass3d)
- [CoolDownLine](addons/cooldownline)
- [Countdown](addons/countdown)
- [Crafting Info Tooltip](addons/crafting_info_tooltip)
- [CraftingWillard](addons/craftingwillard)
- [Cram The Spam](addons/cram_the_spam)
- [Crusher](addons/crusher)
- [DAoCBuff](addons/daocbuff)
- [DPSMeter](addons/dpsmeter)
- [DaemonAssist](addons/daemonassist)
- [DammazKron](addons/dammazkron)
- [DasBoot](addons/dasboot)
- [Dascore](addons/dascore)
- [DeMoNiCore](addons/demonicore)
- [Deathblow](addons/deathblow)
- [Deathblow2](addons/deathblow2)
- [DeepSleep](addons/deepsleep)
- [Default Tactic Set](addons/default_tactic_set)
- [DetauntHelper](addons/detaunthelper)
- [Ding](addons/ding)
- [DisableCombatLog](addons/disablecombatlog)
- [Duel](addons/duel)
- [DuelInvite](addons/duelinvite)
- [DuffTimer](addons/dufftimer)
- [DwarfTalk](addons/dwarftalk)
- [Dye Preview](addons/dye_preview)
- [EA_LoadingScreen](addons/ea_loadingscreen)
- [EA_OpenPartyWindow](addons/ea_openpartywindow)
- [EA_ScenarioGroupWindow](addons/ea_scenariogroupwindow)
- [EA_ThreePartBar](addons/ea_threepartbar)
- [EA_UiDebugTools](addons/ea_uidebugtools)
- [EA_UiModWindow](addons/ea_uimodwindow)
- [EZCraft](addons/ezcraft)
- [EZCraftX](addons/ezcraftx)
- [EZGuard](addons/ezguard)
- [Effigy](addons/effigy)
- [Emojii](addons/emojii)
- [EmoteAlert](addons/emotealert)
- [Enemy](addons/enemy)
- [EveryBodyGuard](addons/everybodyguard)
- [FastFriends](addons/fastfriends)
- [FastInteract](addons/fastinteract)
- [Fight Finder](addons/fight_finder)
- [FixGit](addons/fixgit)
- [FlagCap](addons/flagcap)
- [FozAuction](addons/fozauction)
- [GCDTracker](addons/gcdtracker)
- [GCDsaver](addons/gcdsaver)
- [GDes](addons/gdes)
- [Ges](addons/ges)
- [GetStats](addons/getstats)
- [Greedy](addons/greedy)
- [Group Icons](addons/group_icons)
- [Group Icons SG](addons/group_icons_sg)
- [GroupRange](addons/grouprange)
- [GroupSpotter](addons/groupspotter)
- [GuardBot](addons/guardbot)
- [GuardLine](addons/guardline)
- [GuardList](addons/guardlist)
- [GuardRange](addons/guardrange)
- [GuildWarden](addons/guildwarden)
- [HammerTime](addons/hammertime)
- [HealAll](addons/healall)
- [HealGrid](addons/healgrid)
- [HealHoverAssist](addons/healhoverassist)
- [HideHiddenFrames](addons/hidehiddenframes)
- [Hopper](addons/hopper)
- [HotbarMorale](addons/hotbarmorale)
- [I HATE YOU THIS MUCH](addons/i_hate_you_this_much)
- [IdentityFound](addons/identityfound)
- [InfoScroller](addons/infoscroller)
- [Info_Alert](addons/info_alert)
- [Info_DeathBlow](addons/info_deathblow)
- [Info_Loot](addons/info_loot)
- [Info_Points](addons/info_points)
- [ItemRack](addons/itemrack)
- [JunkDump](addons/junkdump)
- [KeyBar](addons/keybar)
- [Keyset](addons/keyset)
- [KeysetMonsterPlay](addons/keysetmonsterplay)
- [KillStreak](addons/killstreak)
- [KillTracker](addons/killtracker)
- [Killer](addons/killer)
- [Kwestor](addons/kwestor)
- [Lib RuString](addons/lib_rustring)
- [LibAchievements](addons/libachievements)
- [LibAddonButton](addons/libaddonbutton)
- [LibCereal-1.0](addons/libcereal-1.0)
- [LibCompactCurrentDateTime](addons/libcompactcurrentdatetime)
- [LibCooldown](addons/libcooldown)
- [LibDateTime](addons/libdatetime)
- [LibGroup](addons/libgroup)
- [LibGuard](addons/libguard)
- [LibJson](addons/libjson)
- [LibPickle](addons/libpickle)
- [LibRange](addons/librange)
- [LibSkillicon](addons/libskillicon)
- [LibSlash](addons/libslash)
- [LibStub](addons/libstub)
- [LibSurveyor](addons/libsurveyor)
- [LibWBToggler](addons/libwbtoggler)
- [LootAlert](addons/lootalert)
- [LoyalPet](addons/loyalpet)
- [MacroIcons](addons/macroicons)
- [ManualScenarioRefresh](addons/manualscenariorefresh)
- [Map](addons/map)
- [MapMonster](addons/mapmonster)
- [MapPin](addons/mappin)
- [MarkBuff](addons/markbuff)
- [Mass Refine](addons/mass_refine)
- [Mech](addons/mech)
- [MegaphonePlus](addons/megaphoneplus)
- [MegaphonePlusPlus](addons/megaphoneplusplus)
- [MiniMapMonster](addons/minimapmonster)
- [Minmap](addons/minmap)
- [Miracle Grow Remix](addons/miracle_grow_remix)
- [MiracleGrow](addons/miraclegrow)
- [MiracleGrowLight](addons/miraclegrowlight)
- [MoraleBG](addons/moralebg)
- [MoraleCircle](addons/moralecircle)
- [MoraleSet](addons/moraleset)
- [Moth](addons/moth)
- [Motion](addons/motion)
- [MouseHint](addons/mousehint)
- [MyReasons](addons/myreasons)
- [NAMBLA](addons/nambla)
- [NPC Item Sale Price](addons/npc_item_sale_price)
- [NaturalLog](addons/naturallog)
- [NerfedButtons](addons/nerfedbuttons)
- [NervAlert](addons/nervalert)
- [NoOverheal](addons/nooverheal)
- [NoUselessMods-Assist](addons/nouselessmods-assist)
- [ObjectInspector](addons/objectinspector)
- [Obsidian](addons/obsidian)
- [OilTimer](addons/oiltimer)
- [OneClickGTAOE](addons/oneclickgtaoe)
- [OverheadFonts](addons/overheadfonts)
- [PMTB](addons/pmtb)
- [PagerWentPoof](addons/pagerwentpoof)
- [Paint the leader](addons/paint_the_leader)
- [PartyAd](addons/partyad)
- [PartyCast](addons/partycast)
- [PeaceOut](addons/peaceout)
- [PetAssist](addons/petassist)
- [PetFixWindow](addons/petfixwindow)
- [Phantom](addons/phantom)
- [PieTracker](addons/pietracker)
- [PlanB](addons/planb)
- [PlayEffectsOn](addons/playeffectson)
- [Pocket Palette](addons/pocket_palette)
- [PotionBar](addons/potionbar)
- [Preciousss](addons/preciousss)
- [Pure](addons/pure)
- [Pure Careerbar](addons/pure_careerbar)
- [QuestBegone](addons/questbegone)
- [Queue Queuer](addons/queue_queuer)
- [QueuePopTimer](addons/queuepoptimer)
- [Quick Performance Toggle](addons/quick_performance_toggle)
- [QuickNameActions+](addons/quicknameactions+)
- [QuickTacticSwitch](addons/quicktacticswitch)
- [QuickWarReport](addons/quickwarreport)
- [RO-Style Combat Text](addons/ro-style_combat_text)
- [RPs](addons/rps)
- [RVAPI_ColorDialog](addons/rvapi_colordialog)
- [RVAPI_Range](addons/rvapi_range)
- [RVMOD_3DPortrait](addons/rvmod_3dportrait)
- [RVMOD_Manager](addons/rvmod_manager)
- [RVMOD_PlayerStatus](addons/rvmod_playerstatus)
- [RVMOD_SquaredDistances](addons/rvmod_squareddistances)
- [RVMOD_Targets](addons/rvmod_targets)
- [RaidMeter](addons/raidmeter)
- [RandomMount](addons/randommount)
- [RandomSayings](addons/randomsayings)
- [RangeFadingSquared](addons/rangefadingsquared)
- [Rangechecker](addons/rangechecker)
- [RealmStatus](addons/realmstatus)
- [Refer](addons/refer)
- [ReliquaryHunter](addons/reliquaryhunter)
- [RememberIt](addons/rememberit)
- [Res](addons/res)
- [ResHelp](addons/reshelp)
- [RetAlert](addons/retalert)
- [RezEmote](addons/rezemote)
- [RoR_SoR](addons/ror_sor)
- [RoR_debolster](addons/ror_debolster)
- [Rolodex](addons/rolodex)
- [Rotation](addons/rotation)
- [RvRContribution](addons/rvrcontribution)
- [RvRStats](addons/rvrstats)
- [RvRStatsTab](addons/rvrstatstab)
- [SB_TDPS](addons/sb_tdps)
- [SNT Info [TDPS]](addons/snt_info_[tdps])
- [SNT_BUTTONS](addons/snt_buttons)
- [SNT_CASTBAR](addons/snt_castbar)
- [SNT_INFO](addons/snt_info)
- [SNT_PANEL](addons/snt_panel)
- [SOR](addons/sor)
- [ScenarioAlert](addons/scenarioalert)
- [ScenarioStats](addons/scenariostats)
- [SelfTarget](addons/selftarget)
- [Sequencer](addons/sequencer)
- [SessionRPs](addons/sessionrps)
- [Shinies](addons/shinies)
- [ShowHealthPercent](addons/showhealthpercent)
- [ShowMeTheBubbles](addons/showmethebubbles)
- [SimpleCombatText](addons/simplecombattext)
- [SimpleXY](addons/simplexy)
- [SocialWindow 2.0](addons/socialwindow_2.0)
- [Soloq](addons/soloq)
- [SpamBayes](addons/spambayes)
- [Squared](addons/squared)
- [SquaredClick](addons/squaredclick)
- [SquaredHDIndicator](addons/squaredhdindicator)
- [SquaredHotIndicators](addons/squaredhotindicators)
- [Statdoll](addons/statdoll)
- [Statdoll Light](addons/statdoll_light)
- [Statdoll Remix](addons/statdoll_remix)
- [StopRes](addons/stopres)
- [Swift Assist](addons/swift_assist)
- [Swinger](addons/swinger)
- [TacticSetNames](addons/tacticsetnames)
- [TalismanGenie](addons/talismangenie)
- [TargetInfoRing](addons/targetinforing)
- [TargetRing](addons/targetring)
- [Targets](addons/targets)
- [TastyButtons](addons/tastybuttons)
- [TaxPayer](addons/taxpayer)
- [TexturedButtons](addons/texturedbuttons)
- [ThankTheResser](addons/thanktheresser)
- [ThankTheTank](addons/thankthetank)
- [TheSeeker](addons/theseeker)
- [ThinkOutLoud](addons/thinkoutloud)
- [TidyChat](addons/tidychat)
- [TidyQueue](addons/tidyqueue)
- [TidyRoll](addons/tidyroll)
- [TimeToDie](addons/timetodie)
- [TokenMachine](addons/tokenmachine)
- [Tokens](addons/tokens)
- [Tome Titan](addons/tome_titan)
- [TomeTracker](addons/tometracker)
- [TortallDPSCore](addons/tortalldpscore)
- [Tortall_DPS](addons/tortall_dps)
- [Tortall_SCC](addons/tortall_scc)
- [Trakario](addons/trakario)
- [TurretRange](addons/turretrange)
- [TurretScrap](addons/turretscrap)
- [Twister](addons/twister)
- [TwisterSet](addons/twisterset)
- [UnrealDBAnnouncer](addons/unrealdbannouncer)
- [VPBreakdown](addons/vpbreakdown)
- [Vectors](addons/vectors)
- [VerticalMorale](addons/verticalmorale)
- [VerticalTactics](addons/verticaltactics)
- [Vertigo](addons/vertigo)
- [WARCommander](addons/warcommander)
- [WARRatingBuster](addons/warratingbuster)
- [WBStutterLess](addons/wbstutterless)
- [WSCT](addons/wsct)
- [WTes](addons/wtes)
- [WaaaghBar](addons/waaaghbar)
- [WarBoard](addons/warboard)
- [WarBoard_AAOTracker](addons/warboard_aaotracker)
- [WarBoard_FPS](addons/warboard_fps)
- [WarBoard_Loc](addons/warboard_loc)
- [WarBoard_Menu](addons/warboard_menu)
- [WarBoard_Session](addons/warboard_session)
- [WarBoard_TDPS](addons/warboard_tdps)
- [WarBoard_TaliMon](addons/warboard_talimon)
- [WarBoard_TogglerHealGrid](addons/warboard_togglerhealgrid)
- [WarBoard_TogglerRankedLeaderboard](addons/warboard_togglerrankedleaderboard)
- [WarBoard_TogglerSquared](addons/warboard_togglersquared)
- [WarBoard_TogglerTidyChat](addons/warboard_togglertidychat)
- [WarBoard_TogglerTidyRoll](addons/warboard_togglertidyroll)
- [WarBoard_TogglerTokenMachine](addons/warboard_togglertokenmachine)
- [WarBoard_TogglerVPBreakdown](addons/warboard_togglervpbreakdown)
- [WarBoard_TogglerWARCommander](addons/warboard_togglerwarcommander)
- [WarBoard_TogglerWSCT](addons/warboard_togglerwsct)
- [WarBoard_TogglerWlh](addons/warboard_togglerwlh)
- [WarBoard_WarWhisperer](addons/warboard_warwhisperer)
- [WarTriage](addons/wartriage)
- [War_RU](addons/war_ru)
- [Warbuilder](addons/warbuilder)
- [Wargames](addons/wargames)
- [WhatsCooking](addons/whatscooking)
- [WhoHealedMe](addons/whohealedme)
- [Wikki's Cooldown Bar](addons/wikki's_cooldown_bar)
- [Wikki's Cooldown Pulse](addons/wikki's_cooldown_pulse)
- [WindowMovers](addons/windowmovers)
- [WoH-Reticle](addons/woh-reticle)
- [WoodysDpsLib](addons/woodysdpslib)
- [XpStatus+G](addons/xpstatus+g)
- [ZCurse_Profiler](addons/zcurse_profiler)
- [ZonePOP](addons/zonepop)
- [alertMod](addons/alertmod)
- [bigger_MacroWindow](addons/bigger_macrowindow)
- [compass](addons/compass)
- [emotes](addons/emotes)
- [followTheLeader](addons/followtheleader)
- [fpsbox](addons/fpsbox)
- [hideInf](addons/hideinf)
- [minesweep](addons/minesweep)
- [nLootLink](addons/nlootlink)
- [nRarity](addons/nrarity)
- [ovdeadnomore](addons/ovdeadnomore)
- [rorAutoInviter](addons/rorautoinviter)
- [scenarioInfo](addons/scenarioinfo)
- [scnoload](addons/scnoload)
- [talisman-monitor](addons/talisman-monitor)
- [warboard_TogglerWarBuilder](addons/warboard_togglerwarbuilder)
- [wbLeadHelper](addons/wbleadhelper)
- [whatsPugSc](addons/whatspugsc)
- [whom](addons/whom)
- [xHUD](addons/xhud)
- [xPanels](addons/xpanels)
- [yAssistHelper](addons/yassisthelper)
- [zMailMod](addons/zmailmod)
