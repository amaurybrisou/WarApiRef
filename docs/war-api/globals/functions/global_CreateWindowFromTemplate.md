# CreateWindowFromTemplate

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 109 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, ActionBarHide, ActionFraction, AdjustTheTip, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Amethyst |
| Files seen in | APAGui.lua, AdjustTheTip.lua, AdvancedContainers.lua, AdvancedRenownTraining.lua, AggroMeter.lua, AuctionAssist.lua, BBarsPetHP.lua, BBarsPlayerMechanic.lua |
| Namespaces detected | CreateWindowFromTemplate |
| Source kinds | lua_calls |
| Example locations | Ace: New, ActionBarHide: New, ActionFraction: Initialize, AdjustTheTip: AddTargetHealthToMouseOver, AdjustTheTip: CreateCheckBoxMenuItem, AdjustTheTip: CreateSliderMenuItem |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 764 |
| Global usage count | 764 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
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
CreateWindowFromTemplate(windowName, templateName, parentWindow)
```

## Description

Observed instantiating repeated UI elements from an XML template.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a runtime window instance name. | Observed values: "ActionFractionWindowContextAutoHide", "ActionFractionWindowContextColorCodeCurrentAP", "AggroMeterWindow"..MobID |
| templateName | Observed as an XML template name. | Observed values: "AbilityButtonTemplate", "ActionFractionWindowContextCheckBox", "AdjustTheTipMenuItemSlider" |
| parentWindow | Observed as a parent window name. | Observed values: "AuctionHouseWindowCreateAuction", "BestiaryTOCPageWindowContentsChild", "CBMainWindow" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Creates or instantiates UI objects from XML-backed definitions.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- AdjustTheTip
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Amethyst
- AuctionStats
- Aura
- AutoMark
- BBars - Mechanic Only
- BankArkel
- BankWindowFix
- Bloody Mess
- BuffHead
- CMap
- CharacterScreenTabFix
- CleanCastbar
- CleansingBuddy
- CombatTextNames
- CoolDownLine
- Crafting Info Tooltip
- CraftingWillard
- Crusher
- DAoCBuff
- DammazKron
- DetauntHelper
- DuffTimer
- EA_LoadingScreen
- EA_ScenarioGroupWindow
- EA_ThreePartBar
- EZCraftX
- EZGuard
- Effigy
- Emojii
- Enemy
- FozAuction
- GCDsaver
- GetStats
- Group Icons
- Group Icons SG
- GroupSpotter
- HealGrid
- Hopper
- InfoScroller
- KillTracker
- Kwestor
- LibAddonButton
- LibRange
- LibWBToggler
- Map
- MapMonster
- MarkBuff
- MiniMapMonster
- Miracle Grow Remix
- MoraleSet
- Motion
- NaturalLog
- OverheadFonts
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- QuickWarReport
- RO-Style Combat Text
- RVMOD_3DPortrait
- RVMOD_SquaredDistances
- RealmStatus
- Refer
- ReliquaryHunter
- RoR_SoR
- RoR_debolster
- SNT_BUTTONS
- SNT_CASTBAR
- ScenarioStats
- SessionRPs
- Shinies
- ShowHealthPercent
- SquaredHotIndicators
- Statdoll Remix
- TacticSetNames
- TargetInfoRing
- TargetRing
- TidyChat
- TidyQueue
- TidyRoll
- Tokens
- Tome Titan
- Tortall_DPS
- TurretRange
- VerticalMorale
- WBStutterLess
- WarBoard
- WarBoard_Menu
- WarBoard_WarWhisperer
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- emotes
- nLootLink
- nRarity
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: New -> CreateWindowFromTemplate(w.name, base, w.parent)
- Ace: New -> CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
- ActionBarHide: New -> CreateWindowFromTemplate(w.name, "EA_ComboBox_DefaultResizable", w.parent)
- ActionBarHide: New -> CreateWindowFromTemplate(w.name, "EA_Button_DefaultCheckBox", w.parent)
- ActionBarHide: New -> CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
- ActionBarHide: New -> CreateWindowFromTemplate(w.name, "EA_EditBox_DefaultFrame", w.parent)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](../../window_api/functions/window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](../../window_api/functions/window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
