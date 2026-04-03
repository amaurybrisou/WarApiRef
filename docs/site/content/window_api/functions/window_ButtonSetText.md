# ButtonSetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 130 addons

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
| Addons seen in | Ace, ActionBarHide, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Amethyst, Assist, AuctionStats |
| Files seen in | AAOTracker.lua, APAGui.lua, AdvancedRenownTraining.lua, AdvancedRenownTrainingImportExport.lua, AggroMeter.lua, AutoBand.lua, BankArkel.lua, BlackBookWindow.lua |
| Namespaces detected | ButtonSetText |
| Source kinds | lua_calls |
| Example locations | Ace: SetText, ActionBarHide: SetText, AdvancedPetAssist: OnShown, AdvancedRenownTrainer: AnywhereShow, AdvancedRenownTrainer: OnHidden, AdvancedRenownTrainer: SetImportExportLabels |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 787 |
| Global usage count | 787 |
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
ButtonSetText(windowName, text)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: "APAOptionsTabsAutoRecall", "APAOptionsTabsFollowTarget", "APAOptionsTabsGeneral" |
| text | Observed as a text or wstring payload. | Observed values: Calling.GetLocalized("selectSlot"), DA.enabled and L "Disable" or L "Enable", DHLang.GetString(DHStrings.ABILITY) |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Amethyst
- Assist
- AuctionStats
- Aura
- AutoBand
- AutoSalvage
- BankArkel
- BlackBook
- Bloody Mess
- BuffHead
- CDown
- CM_ClosetGoblin
- CMap
- CaVES
- Calling
- CastSequence
- ChattyCathy
- Cheeseboard
- Crusher
- DAoCBuff
- DPSMeter
- DaemonAssist
- DammazKron
- DetauntHelper
- DuffTimer
- Dye Preview
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EA_UiModWindow
- EZCraftX
- EZGuard
- Effigy
- Emojii
- Enemy
- FastFriends
- FozAuction
- GCDsaver
- GroupRange
- GuildWarden
- HealGrid
- Hopper
- InfoScroller
- JunkDump
- KeyBar
- Keyset
- Killer
- Kwestor
- LibAddonButton
- LibWBToggler
- ManualScenarioRefresh
- Map
- MapMonster
- Mass Refine
- MegaphonePlusPlus
- Miracle Grow Remix
- MoraleSet
- Motion
- NaturalLog
- NerfedButtons
- ObjectInspector
- Obsidian
- PartyCast
- PieTracker
- Pocket Palette
- PotionBar
- Pure
- Pure Careerbar
- QuickTacticSwitch
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RVMOD_Targets
- RandomMount
- RealmStatus
- Refer
- ReliquaryHunter
- RoR_SoR
- RvRStats
- RvRStatsTab
- SNT_BUTTONS
- SNT_CASTBAR
- SNT_INFO
- SOR
- ScenarioStats
- Sequencer
- SessionRPs
- Shinies
- SocialWindow 2.0
- Squared
- Statdoll Remix
- TacticSetNames
- TalismanGenie
- TargetRing
- TastyButtons
- TaxPayer
- ThinkOutLoud
- TidyChat
- TidyQueue
- TidyRoll
- Tokens
- Tome Titan
- TurretRange
- TwisterSet
- Vectors
- WBStutterLess
- WSCT
- WTes
- WarBoard_AAOTracker
- WarBoard_WarWhisperer
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- alertMod
- bigger_MacroWindow
- followTheLeader
- nLootLink
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: SetText -> ButtonSetText(self.name, text and L "true" or L "false")
- Ace: SetText -> ButtonSetText(self.name, towstring(text))
- ActionBarHide: SetText -> ButtonSetText(self.name, text and L "true" or L "false")
- ActionBarHide: SetText -> ButtonSetText(self.name, towstring(text))
- AdvancedPetAssist: OnShown -> ButtonSetText("APAOptionsTabsGeneral", L "General")
- AdvancedPetAssist: OnShown -> ButtonSetText("APAOptionsTabsFollowTarget", L "Follow Target")

## Related APIs

- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Notes

- none
