# ComboBoxGetSelectedMenuItem

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
| Addons seen in | Ace, ActionBarHide, AdvancedRenownTrainer, Amethyst, BuffHead, CDown, CaVES, Calling |
| Files seen in | AdvancedRenownTrainingImportExport.lua, CDownSettings.lua, CallingSetup.lua, ChattyCathy.lua, Code/Assist/Assist.lua, Code/CombatLog/CombatLogStatsWindow.lua, Code/Core/Groups/EnemyEffectFilter.lua, Code/Core/Main.lua |
| Namespaces detected | ComboBoxGetSelectedMenuItem |
| Source kinds | lua_calls |
| Example locations | Ace: SelectedIndex, ActionBarHide: SelectedIndex, AdvancedRenownTrainer: OnExportButtonPressed, Amethyst: SelectedIndex, BuffHead: OnAnimationChanged, BuffHead: OnCompressionChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 450 |
| Global usage count | 450 |
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
ComboBoxGetSelectedMenuItem(arg1)
```

## Description

Observed as a window function across 79 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "ChattyCathyOptToCombo", "ChattyCathyOptWindowCombo", "DyeCombo" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionBarHide
- AdvancedRenownTrainer
- Amethyst
- BuffHead
- CDown
- CaVES
- Calling
- CastSequence
- ChattyCathy
- Crafting Info Tooltip
- Crusher
- DAoCBuff
- DammazKron
- DeepSleep
- DetauntHelper
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
- Keyset
- KeysetMonsterPlay
- Killer
- LibAddonButton
- LibWBToggler
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
- RaidMeter
- RealmStatus
- SOR
- ScenarioStats
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
- WSCT
- WarBoard
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- alertMod
- nLootLink
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels

## Examples

- Ace: SelectedIndex -> ComboBoxGetSelectedMenuItem(self.name)
- ActionBarHide: SelectedIndex -> ComboBoxGetSelectedMenuItem(self.name)
- AdvancedRenownTrainer: OnExportButtonPressed -> ComboBoxGetSelectedMenuItem(PresetWindowName.."LoadComboBox")
- Amethyst: SelectedIndex -> ComboBoxGetSelectedMenuItem(self.name)
- BuffHead: OnAnimationChanged -> ComboBoxGetSelectedMenuItem(windowName.."AnimationComboBox")
- BuffHead: OnCompressionChanged -> ComboBoxGetSelectedMenuItem(windowName.."CompressList")

## Related APIs

- [OnSelChanged](../../xml/handlers/handler_OnSelChanged.md) (HIGH 88/100) - XML Event

## Notes

- none
