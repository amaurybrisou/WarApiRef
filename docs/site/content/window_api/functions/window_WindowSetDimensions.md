# WindowSetDimensions

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 131 addons

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
| Addons seen in | Ace, ActionBarHide, ActionFraction, AdjustTheTip, AdvancedPetAssist, Amethyst, AnywhereTrainer, Atlas |
| Files seen in | APAGui.lua, APAGuiHUD.lua, AdjustTheTip.lua, AdvancedContainers.lua, AuctionStats.lua, BBarsPetHP.lua, BBarsPlayerMechanic.lua, BankArkel.lua |
| Namespaces detected | WindowSetDimensions |
| Source kinds | lua_calls |
| Example locations | Ace: Resize, ActionBarHide: Resize, ActionFraction: ResetWindow, AdjustTheTip: AddTargetHealthToMouseOver, AdvancedPetAssist: ApplyPetTargetHUDLayout, AdvancedPetAssist: EnsureSmallEditBox |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 847 |
| Global usage count | 847 |
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
WindowSetDimensions(windowName, arg2, arg3)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "10secBG2", "30secBG2", "60secBG2" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: (mg.vLayout.dimx)*nMult+5, (sw or 1)*self.height, 0 |
| arg3 | Observed as a runtime window or control identifier. | Observed values: (Height[1]*Height[2])+padding, (StatdollWnd.defaultHeight+y), (StatdollWndLight.defaultHeight+StatdollWndLight.powerHeight) |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- AdjustTheTip
- AdvancedPetAssist
- Amethyst
- AnywhereTrainer
- Atlas
- AuctionStats
- Aura
- BBars - Mechanic Only
- BankArkel
- BarText (Influence)
- BuffHead
- CDown
- CMap
- CaVES
- Calling
- CastSequence
- ChattyCathy
- CoolDownLine
- Crafting Info Tooltip
- CraftingWillard
- Crusher
- DAoCBuff
- DaemonAssist
- DammazKron
- Dascore
- DetauntHelper
- DuffTimer
- EA_LoadingScreen
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_ThreePartBar
- EA_UiModWindow
- EZGuard
- Effigy
- Enemy
- EveryBodyGuard
- GCDsaver
- GetStats
- Group Icons SG
- GroupRange
- HealGrid
- Hopper
- InfoScroller
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- LibWBToggler
- LootAlert
- Map
- MapMonster
- Miracle Grow Remix
- MoraleSet
- Moth
- Motion
- NAMBLA
- NaturalLog
- NerfedButtons
- Obsidian
- OverheadFonts
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RaidMeter
- RealmStatus
- ReliquaryHunter
- ResHelp
- RoR_SoR
- RoR_debolster
- SOR
- ScenarioStats
- Sequencer
- SessionRPs
- Shinies
- SocialWindow 2.0
- SquaredHotIndicators
- Statdoll
- Statdoll Remix
- TacticSetNames
- TargetInfoRing
- TargetRing
- TastyButtons
- TexturedButtons
- TidyChat
- TidyQueue
- TidyRoll
- Tokens
- TomeTracker
- TurretRange
- Twister
- VerticalMorale
- VerticalTactics
- WARCommander
- WARRatingBuster
- WSCT
- WTes
- WarBoard_Loc
- WarBoard_Menu
- WarBoard_TogglerHealGrid
- WarBoard_TogglerRankedLeaderboard
- WarBoard_TogglerSquared
- WarBoard_TogglerTidyChat
- WarBoard_TogglerTidyRoll
- WarBoard_TogglerTokenMachine
- WarBoard_TogglerVPBreakdown
- WarBoard_TogglerWARCommander
- WarBoard_TogglerWSCT
- WarBoard_TogglerWlh
- WarBoard_WarWhisperer
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- ZonePOP
- emotes
- minesweep
- scenarioInfo
- warboard_TogglerWarBuilder
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Resize -> WindowSetDimensions(self.name, width, self.height)
- Ace: Resize -> WindowSetDimensions(self.name, self.width, self.height)
- Ace: Resize -> WindowSetDimensions(self.name, width, height)
- ActionBarHide: Resize -> WindowSetDimensions(self.name, 22, height)
- ActionBarHide: Resize -> WindowSetDimensions(self.name, width, 41)
- ActionBarHide: Resize -> WindowSetDimensions(self.name, width, 32)

## Related APIs

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Size](../../xml/element_types/element_Size.md) (HIGH 100/100) - XML Element Type
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [EA_Window_PublicQuestTracker.GetLocalAreaInfluenceID](../../globals/functions/global_EA_Window_PublicQuestTracker.GetLocalAreaInfluenceID.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.gsub](../../globals/functions/global_wstring.gsub.md) (HIGH 100/100) - Global Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function

## Notes

- Advanced return analysis: No strong return evidence observed
