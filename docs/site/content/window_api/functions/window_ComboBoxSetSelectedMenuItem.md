# ComboBoxSetSelectedMenuItem

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 88 addons

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
| Addons seen in | Ace, ActionBarHide, AdvancedPetAssist, Amethyst, BankArkel, Brizio's Crappy Computer Medic, BuffHead, Busted |
| Files seen in | APAGui.lua, BankArkel.lua, Busted.lua, CCM.lua, CCTV.lua, CDownSettings.lua, ChattyCathy.lua, Code/Assist/Assist.lua |
| Namespaces detected | ComboBoxSetSelectedMenuItem |
| Source kinds | lua_calls |
| Example locations | Ace: SelectIndex, ActionBarHide: SelectIndex, AdvancedPetAssist: FillCombo, Amethyst: SelectIndex, BankArkel: PackClose, BankArkel: SetupCombos |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 530 |
| Global usage count | 530 |
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
ComboBoxSetSelectedMenuItem(arg1, arg2)
```

## Description

Observed as a window function across 88 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BankArkelBackpackCombo", "BustedGUIAddonSelect", "CCTVSettingsCombobox3" |
| arg2 | Observed as a function or method reference. | Observed values: #ACTIVE_CONDITION, #CONDITIONS, #CONDITIONS-1 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- AdvancedPetAssist
- Amethyst
- BankArkel
- Brizio's Crappy Computer Medic
- BuffHead
- Busted
- CCTV
- CDown
- CaVES
- CastSequence
- ChattyCathy
- Crafting Info Tooltip
- Crusher
- DAoCBuff
- DPSMeter
- DaemonAssist
- DammazKron
- DeepSleep
- DetauntHelper
- DuffTimer
- Dye Preview
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- Enemy
- EveryBodyGuard
- GCDsaver
- GroupRange
- GuildWarden
- HealGrid
- Hopper
- InfoScroller
- KeysetMonsterPlay
- Killer
- LibAddonButton
- LibWBToggler
- LoyalPet
- Map
- MapMonster
- MapPin
- MarkBuff
- MegaphonePlusPlus
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- Obsidian
- OverheadFonts
- PartyCast
- Pocket Palette
- PotionBar
- Pure
- Pure Careerbar
- RVMOD_SquaredDistances
- RealmStatus
- SOR
- ScenarioStats
- Sequencer
- Shinies
- Statdoll Remix
- TalismanGenie
- TargetRing
- TastyButtons
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyRoll
- TokenMachine
- Tokens
- TurretRange
- Vectors
- WARCommander
- WarBoard_FPS
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- ZonePOP
- alertMod
- nLootLink
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels

## Examples

- Ace: SelectIndex -> ComboBoxSetSelectedMenuItem(self.name, index)
- ActionBarHide: SelectIndex -> ComboBoxSetSelectedMenuItem(self.name, index)
- AdvancedPetAssist: FillCombo -> ComboBoxSetSelectedMenuItem(comboName, selIdx)
- Amethyst: SelectIndex -> ComboBoxSetSelectedMenuItem(self.name, index)
- BankArkel: PackClose -> ComboBoxSetSelectedMenuItem("BankArkelBackpackCombo", 1)
- BankArkel: SetupCombos -> ComboBoxSetSelectedMenuItem("BankArkelBackpackCombo", 1)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnSelChanged](../../xml/handlers/handler_OnSelChanged.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function

## Notes

- none
