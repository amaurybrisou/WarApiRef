# WindowGetShowing

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 133 addons

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
| Addons seen in | Ace, ActionBarHide, ActionFraction, AdjustTheTip, AdvancedRenownTrainer, AggroMeter, Amethyst, AnywhereTrainer |
| Files seen in | AdjustTheTip.lua, AdvancedRenownTraining.lua, AggroMeter.lua, AnywhereTrainerAdditions.lua, AuctionAssist.lua, AuctionStats.lua, AutoSalvage.lua, Backpack.lua |
| Namespaces detected | WindowGetShowing |
| Source kinds | lua_calls |
| Example locations | Ace: Showing, ActionBarHide: Showing, ActionFraction: SetLocationActionPointBar, AdjustTheTip: UpdateCallback, AdvancedRenownTrainer: TogglePresets, AggroMeter: Close |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 649 |
| Global usage count | 649 |
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
WindowGetShowing(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "AbilitiesWindow", "AdvancedRenownTrainingWindow", "AggroMeterGrayWindow" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- AdjustTheTip
- AdvancedRenownTrainer
- AggroMeter
- Amethyst
- AnywhereTrainer
- AnywhereTrainerAdditions
- Atlas
- AuctionStats
- Aura
- AutoMark
- AutoSalvage
- BankWindowFix
- BlackBook
- Bloody Mess
- BuffHead
- Busted
- CCTV
- CaVES
- Calling
- CastSequence
- ChattyCathy
- CleanUnitFrames
- CombatTextNames
- CraftingWillard
- Crusher
- DAoCBuff
- DPSMeter
- DaemonAssist
- DammazKron
- Dascore
- DeepSleep
- DetauntHelper
- DuffTimer
- Dye Preview
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- FixGit
- GCDsaver
- Group Icons SG
- GroupRange
- GroupSpotter
- GuardLine
- HealGrid
- Hopper
- InfoScroller
- ItemRack
- KeyBar
- Keyset
- KeysetMonsterPlay
- KillTracker
- Kwestor
- LibAddonButton
- LibGroup
- LibWBToggler
- Map
- MapMonster
- MapPin
- MarkBuff
- Minmap
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- Obsidian
- PartyCast
- PeaceOut
- Pocket Palette
- PotionBar
- Pure
- Pure Careerbar
- Queue Queuer
- QuickNameActions+
- QuickTacticSwitch
- RVMOD_3DPortrait
- RaidMeter
- RandomMount
- RealmStatus
- ReliquaryHunter
- ResHelp
- RoR_SoR
- SNT_CASTBAR
- SNT_INFO
- SNT_PANEL
- SOR
- Sequencer
- Shinies
- Soloq
- Squared
- Statdoll
- Statdoll Remix
- TargetRing
- TastyButtons
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyQueue
- TidyRoll
- Tokens
- Tome Titan
- TomeTracker
- Trakario
- TurretRange
- VPBreakdown
- Vectors
- VerticalMorale
- WarBoard
- WarBoard_TogglerVPBreakdown
- WarBoard_TogglerWARCommander
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WindowMovers
- WoH-Reticle
- XpStatus+G
- ZCurse_Profiler
- nLootLink
- nRarity
- scenarioInfo
- scnoload
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Showing -> WindowGetShowing(self.name)
- ActionBarHide: Showing -> WindowGetShowing(self.name)
- ActionFraction: SetLocationActionPointBar -> WindowGetShowing("PlayerWindow")
- AdjustTheTip: UpdateCallback -> WindowGetShowing(Tooltips.curTooltipWindow)
- AdvancedRenownTrainer: TogglePresets -> WindowGetShowing(PresetWindowName)
- AggroMeter: Close -> WindowGetShowing("AggroMeterGrayWindow")

## Related APIs

- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnKeyEscape](../../xml/handlers/handler_OnKeyEscape.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Used With

- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Notes

- Advanced return analysis: No strong return evidence observed
