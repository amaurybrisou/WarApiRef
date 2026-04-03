# ComboBoxClearMenuItems

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 79 addons

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
| Addons seen in | Ace, ActionBarHide, AdvancedPetAssist, AdvancedRenownTrainer, Amethyst, BankArkel, BuffHead, CCTV |
| Files seen in | APAGui.lua, AdvancedRenownTraining.lua, AdvancedRenownTrainingImportExport.lua, BankArkel.lua, CCTV.lua, CDownSettings.lua, CallingSetup.lua, ChattyCathy.lua |
| Namespaces detected | ComboBoxClearMenuItems |
| Source kinds | lua_calls |
| Example locations | Ace: Clear, ActionBarHide: Clear, AdvancedPetAssist: FillCombo, AdvancedRenownTrainer: DeletePreset, AdvancedRenownTrainer: GeneratePresetByLinkData, AdvancedRenownTrainer: SaveCurrentSpecAsPreset |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 331 |
| Global usage count | 331 |
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
ComboBoxClearMenuItems(arg1)
```

## Description

Observed as a window function across 79 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "AdvancedRenownTrainingPresetsWindowLoadComboBox", "BankArkelBackpackCombo", "CCTVSettingsCombobox4" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionBarHide
- AdvancedPetAssist
- AdvancedRenownTrainer
- Amethyst
- BankArkel
- BuffHead
- CCTV
- CDown
- CMap
- Calling
- CastSequence
- ChattyCathy
- Crafting Info Tooltip
- Crusher
- DAoCBuff
- DPSMeter
- DaemonAssist
- DetauntHelper
- EA_OpenPartyWindow
- EA_UiDebugTools
- EA_UiModWindow
- EZCraftX
- EZGuard
- Effigy
- Enemy
- EveryBodyGuard
- FozAuction
- GCDsaver
- GroupRange
- HealGrid
- Hopper
- InfoScroller
- Keyset
- KeysetMonsterPlay
- Killer
- LibAddonButton
- LibWBToggler
- Map
- MapMonster
- MarkBuff
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- Obsidian
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- RVMOD_Manager
- RVMOD_Targets
- RealmStatus
- SOR
- ScenarioStats
- Shinies
- SocialWindow 2.0
- Statdoll Remix
- TalismanGenie
- TargetRing
- TastyButtons
- TexturedButtons
- ThinkOutLoud
- TidyRoll
- Tokens
- TurretRange
- Vectors
- WARCommander
- WarBoard_WarWhisperer
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- ZonePOP
- nLootLink
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: Clear -> ComboBoxClearMenuItems(self.name)
- ActionBarHide: Clear -> ComboBoxClearMenuItems(self.name)
- AdvancedPetAssist: FillCombo -> ComboBoxClearMenuItems(comboName)
- AdvancedRenownTrainer: DeletePreset -> ComboBoxClearMenuItems(PresetWindowName.."LoadComboBox")
- AdvancedRenownTrainer: GeneratePresetByLinkData -> ComboBoxClearMenuItems("AdvancedRenownTrainingPresetsWindowLoadComboBox")
- AdvancedRenownTrainer: SaveCurrentSpecAsPreset -> ComboBoxClearMenuItems(PresetWindowName.."LoadComboBox")

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event

## Used With

- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function

## Notes

- none
