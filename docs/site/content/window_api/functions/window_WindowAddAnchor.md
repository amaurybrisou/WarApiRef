# WindowAddAnchor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 142 addons

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
| Addons seen in | Ace, ActionBarHide, ActionFraction, AdjustTheTip, AdvancedPetAssist, AdvancedRenownTrainer, Amethyst, AnywhereTrainerAdditions |
| Files seen in | APAGui.lua, AdjustTheTip.lua, AdvancedContainers.lua, AdvancedRenownTraining.lua, AnywhereTrainerAdditions.lua, AuctionAssist.lua, AuctionStats.lua, BBarsPetHP.lua |
| Namespaces detected | WindowAddAnchor |
| Source kinds | lua_calls |
| Example locations | Ace: AddAnchor, ActionBarHide: AddAnchor, ActionFraction: ResetWindow, ActionFraction: SetLocationActionPointBar, ActionFraction: SetLocationCenterScreen, AdjustTheTip: AddTargetHealthToMouseOver |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1012 |
| Global usage count | 1012 |
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
WindowAddAnchor(windowName, point, relativeTo, relativePoint, offsetX, offsetY)
```

## Description

Observed positioning windows relative to other runtime UI elements.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as the window being positioned. | Observed values: "BBarsPetHPBG", "BBarsPlayerMechanic1BG", "BBarsPlayerMechanic2BG" |
| point | Observed as an anchor point string. | Observed values: "LEFT", "RIGHT", "bootom" |
| relativeTo | Observed as a reference window name. | Observed values: "APAOptionsContent", "AuctionHouseWindowCreateAuction", "BBarsPlayerMechanic2BG" |
| relativePoint | Observed as a reference anchor point string. | Observed values: "LEFT", "RIGHT", "bottom" |
| offsetX | Observed as a numeric horizontal offset. | Observed values: (breadth/2)-32, (dragX+x)/InterfaceCore.GetScale(), (i-1)*28 |
| offsetY | Observed as a numeric vertical offset. | Observed values: (2/InterfaceCore.GetScale())*wbLeadHelper.activeSettings.windowSize, (dragY+y)/InterfaceCore.GetScale(), (g.settings.combatLogIDSRowSize[2]+g.settings.combatLogIDSRowPadding)*(index-1) |

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
- AdvancedRenownTrainer
- Amethyst
- AnywhereTrainerAdditions
- AuctionStats
- Aura
- AutoMark
- BBars - Mechanic Only
- BagOMatic
- BankArkel
- BetterCC
- BlackBook
- BuffHead
- CDown
- CMap
- CastSequence
- ChattyCathy
- CombatTextNames
- Countdown
- Crafting Info Tooltip
- CraftingWillard
- Crusher
- DAoCBuff
- DammazKron
- Dascore
- DetauntHelper
- DuffTimer
- EA_LoadingScreen
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_ThreePartBar
- EA_UiModWindow
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
- GuardList
- GuardRange
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
- ManualScenarioRefresh
- Map
- MapMonster
- Mass Refine
- MiniMapMonster
- Miracle Grow Remix
- MoraleBG
- MoraleSet
- Moth
- Motion
- MouseHint
- NaturalLog
- NerfedButtons
- Obsidian
- OverheadFonts
- PartyCast
- PotionBar
- Preciousss
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
- RvRContribution
- RvRStats
- RvRStatsTab
- SOR
- ScenarioStats
- SessionRPs
- Shinies
- ShowHealthPercent
- SimpleCombatText
- SimpleXY
- SocialWindow 2.0
- SquaredHotIndicators
- Statdoll Remix
- TacticSetNames
- TargetInfoRing
- TargetRing
- TastyButtons
- TexturedButtons
- TidyChat
- TidyQueue
- TidyRoll
- TokenMachine
- Tokens
- Tortall_DPS
- TurretRange
- VPBreakdown
- Vectors
- VerticalMorale
- VerticalTactics
- WARRatingBuster
- WSCT
- WarBoard_Loc
- WarBoard_Menu
- WarBoard_WarWhisperer
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WindowMovers
- WoH-Reticle
- XpStatus+G
- bigger_MacroWindow
- emotes
- fpsbox
- nLootLink
- nRarity
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: AddAnchor -> WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
- ActionBarHide: AddAnchor -> WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
- ActionFraction: ResetWindow -> WindowAddAnchor(windowName, "center", "Root", "center", 0, 0)
- ActionFraction: SetLocationActionPointBar -> WindowAddAnchor(windowName, "center", "PlayerWindowStatusContainerAPPercentBar", "center", 2, 6)
- ActionFraction: SetLocationCenterScreen -> WindowAddAnchor(windowName, "center", "Root", "center", 0, 0)
- AdjustTheTip: AddTargetHealthToMouseOver -> WindowAddAnchor(c_HEALTH_BAR_CONTAINER, "bottomright", "MouseOverTargetUnitWindow", "bottomright", -10, -10)

## Related APIs

- [Anchor](../../xml/element_types/element_Anchor.md) (HIGH 100/100) - XML Element Type
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [StatusBarSetBackgroundTint](window_StatusBarSetBackgroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetForegroundTint](window_StatusBarSetForegroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Notes

- Advanced return analysis: No strong return evidence observed
