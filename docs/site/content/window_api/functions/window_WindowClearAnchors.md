# WindowClearAnchors

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 114 addons

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
| Addons seen in | Ace, ActionBarHide, ActionFraction, AdjustTheTip, AdvancedPetAssist, Amethyst, AnywhereTrainerAdditions, AuctionStats |
| Files seen in | APAGui.lua, AdjustTheTip.lua, AnywhereTrainerAdditions.lua, AuctionStats.lua, BBarsPetHP.lua, BankArkel.lua, Bar.lua, Bars/HealGridCareerBar.lua |
| Namespaces detected | WindowClearAnchors |
| Source kinds | lua_calls |
| Example locations | Ace: ClearAnchors, ActionBarHide: ClearAnchors, ActionFraction: ResetWindow, ActionFraction: SetLocationActionPointBar, ActionFraction: SetLocationCenterScreen, AdjustTheTip: UpdateCallback |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 612 |
| Global usage count | 612 |
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
WindowClearAnchors(windowName)
```

## Description

Observed resetting a window layout before applying new runtime anchors.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "AnywhereTrainerBottomBookend", "AnywhereTrainerTopBookend", "BBarsPetHPBack" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates runtime window layout state.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- AdjustTheTip
- AdvancedPetAssist
- Amethyst
- AnywhereTrainerAdditions
- AuctionStats
- Aura
- AutoMark
- BBars - Mechanic Only
- BankArkel
- BetterCC
- BlackBook
- BuffHead
- CDown
- CastSequence
- ChattyCathy
- CombatTextNames
- CraftingWillard
- Crusher
- DAoCBuff
- DammazKron
- Dascore
- DetauntHelper
- DuffTimer
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- Enemy
- GCDTracker
- GCDsaver
- Group Icons
- GroupRange
- GroupSpotter
- GuardBot
- HealGrid
- Hopper
- InfoScroller
- JunkDump
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- LibWBToggler
- MacroIcons
- Map
- MapMonster
- Mass Refine
- MiniMapMonster
- Miracle Grow Remix
- MoraleBG
- Moth
- Motion
- MouseHint
- NaturalLog
- NerfedButtons
- Obsidian
- OverheadFonts
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- QuickWarReport
- RO-Style Combat Text
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RVMOD_Targets
- RealmStatus
- ReliquaryHunter
- ResHelp
- RoR_SoR
- RvRStats
- RvRStatsTab
- SOR
- ScenarioStats
- SessionRPs
- Shinies
- SimpleCombatText
- Squared
- SquaredHotIndicators
- TacticSetNames
- TargetInfoRing
- TargetRing
- TastyButtons
- TexturedButtons
- TidyChat
- Tokens
- TomeTracker
- TurretRange
- Vectors
- VerticalMorale
- VerticalTactics
- WARRatingBuster
- WSCT
- WarBoard_Loc
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WindowMovers
- WoH-Reticle
- XpStatus+G
- emotes
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: ClearAnchors -> WindowClearAnchors(self.name)
- ActionBarHide: ClearAnchors -> WindowClearAnchors(self.name)
- ActionFraction: ResetWindow -> WindowClearAnchors(windowName)
- ActionFraction: SetLocationActionPointBar -> WindowClearAnchors(windowName)
- ActionFraction: SetLocationCenterScreen -> WindowClearAnchors(windowName)
- AdjustTheTip: UpdateCallback -> WindowClearAnchors(Tooltips.curTooltipWindow)

## Related APIs

- [Anchor](../../xml/element_types/element_Anchor.md) (HIGH 100/100) - XML Element Type
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Notes

- Advanced return analysis: No strong return evidence observed
