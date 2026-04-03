# ButtonSetPressedFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 122 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, ActionBarHide, ActionFraction, AdjustTheTip, AdvancedRenownTrainer, AggroMeter, Amethyst, Atlas |
| Files seen in | AAOTracker.lua, AdjustTheTip.lua, AdvancedRenownTraining.lua, AggroMeter.lua, AuctionAssist.lua, BlackBookWindow.lua, Bloody Mess.lua, BuddyBind.lua |
| Namespaces detected | ButtonSetPressedFlag |
| Source kinds | lua_calls |
| Example locations | Ace: SetValue, ActionBarHide: SetValue, ActionFraction: Initialize, ActionFraction: ResetWindow, ActionFraction: RightClick, ActionFraction: ToggleAutoHide |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1009 |
| Global usage count | 1009 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
ButtonSetPressedFlag(arg1, arg2)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "ActionFractionWindowContextAutoHideCheckBox", "ActionFractionWindowContextColorCodeCurrentAPCheckBox", "AggroMeterGrayWindowBlackTab" |
| arg2 | Observed as a function or method reference. | Observed values: (CDownVar.optCDs.CDsbelow==1), (CDownVar.optCDs.fade_start>0), (DB.FONT==ndx) |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- AdjustTheTip
- AdvancedRenownTrainer
- AggroMeter
- Amethyst
- Atlas
- AuctionStats
- Aura
- AutoSalvage
- BlackBook
- Bloody Mess
- BuddyBind
- BuffHead
- CCTV
- CDown
- CM_ClosetGoblin
- CMap
- CaVES
- Calling
- CastSequence
- Cheeseboard
- Crafting Info Tooltip
- Crusher
- DAoCBuff
- DPSMeter
- DammazKron
- DetauntHelper
- DuffTimer
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiModWindow
- EZCraftX
- EZGuard
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- FastFriends
- FozAuction
- GCDsaver
- Group Icons SG
- GroupRange
- GroupSpotter
- HealGrid
- Hopper
- InfoScroller
- JunkDump
- KeyBar
- Keyset
- KeysetMonsterPlay
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- LibGroup
- LibWBToggler
- Map
- MapMonster
- MapPin
- MarkBuff
- MegaphonePlusPlus
- Minmap
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- Obsidian
- PartyCast
- Phantom
- PieTracker
- Pocket Palette
- PotionBar
- Pure
- Pure Careerbar
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RaidMeter
- RealmStatus
- ReliquaryHunter
- RvRStats
- RvRStatsTab
- SNT_CASTBAR
- SNT_INFO
- Sequencer
- Shinies
- SocialWindow 2.0
- Statdoll Remix
- TalismanGenie
- TargetRing
- TastyButtons
- TaxPayer
- TexturedButtons
- TidyChat
- TidyQueue
- TidyRoll
- Tokens
- Tome Titan
- TurretRange
- Vectors
- WARCommander
- WSCT
- WTes
- WarBoard_AAOTracker
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- bigger_MacroWindow
- nLootLink
- nRarity
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: SetValue -> ButtonSetPressedFlag(self.name, checkState)
- ActionBarHide: SetValue -> ButtonSetPressedFlag(self.name, checkState)
- ActionFraction: Initialize -> ButtonSetPressedFlag("ActionFractionWindowContextColorCodeCurrentAPCheckBox", ActionFraction.Settings.bColorCodeCurrentAP)
- ActionFraction: Initialize -> ButtonSetPressedFlag("ActionFractionWindowContextAutoHideCheckBox", ActionFraction.Settings.bAutoHide)
- ActionFraction: ResetWindow -> ButtonSetPressedFlag("ActionFractionWindowContextAutoHideCheckBox", ActionFraction.Settings.bAutoHide)
- ActionFraction: ResetWindow -> ButtonSetPressedFlag("ActionFractionWindowContextColorCodeCurrentAPCheckBox", ActionFraction.Settings.bColorCodeCurrentAP)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [ButtonGetCheckButtonFlag](window_ButtonGetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetCheckButtonFlag](window_ButtonSetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetStayDownFlag](window_ButtonSetStayDownFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [PartyUtils.GetWarbandLeader](../../globals/functions/global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.M_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.M_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.R_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.R_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.ChatLogFilters.SHOUT](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field

## Notes

- Advanced return analysis: No strong return evidence observed
