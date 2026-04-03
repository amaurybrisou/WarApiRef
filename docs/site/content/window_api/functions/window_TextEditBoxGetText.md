# TextEditBoxGetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 76 addons

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
| Addons seen in | Ace, ActionBarHide, AdvancedPetAssist, AdvancedRenownTrainer, Amethyst, Aura, BuffHead, CastSequence |
| Files seen in | APAGuiHelpers.lua, AdvancedRenownTraining.lua, AdvancedRenownTrainingImportExport.lua, Code/Core/Groups/EnemyEffectFilter.lua, Code/Core/Main.lua, Code/Core/Utils.lua, Code/Intercom/Intercom.lua, Code/UnitFrames/ClickCasting.lua |
| Namespaces detected | TextEditBoxGetText |
| Source kinds | lua_calls |
| Example locations | Ace: GetText, ActionBarHide: GetText, AdvancedPetAssist: ParseRGB, AdvancedRenownTrainer: ImportNameInputOkButtonPressed, AdvancedRenownTrainer: ImportOkButtonPressed, AdvancedRenownTrainer: SavePreset |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 411 |
| Global usage count | 411 |
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
TextEditBoxGetText(arg1)
```

## Description

Observed reading from or writing to edit-box controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "DyeWindowFilterEditBox", "EnemyChooseChannelDialogTellDetailsName", "EnemyClickCastingDialogContentScrollChildActionConfig2Command" |

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
- Aura
- BuffHead
- CastSequence
- Crusher
- DAoCBuff
- DetauntHelper
- DuffTimer
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- Enemy
- FastFriends
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
- LibAddonButton
- LibWBToggler
- LoyalPet
- Map
- MapMonster
- MapPin
- Mass Refine
- MegaphonePlusPlus
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- Obsidian
- PartyCast
- Pocket Palette
- Pure
- Pure Careerbar
- RVAPI_ColorDialog
- RandomMount
- RealmStatus
- SNT_BUTTONS
- SNT_PANEL
- SOR
- Sequencer
- Shinies
- Squared
- TacticSetNames
- TargetRing
- TastyButtons
- TexturedButtons
- ThinkOutLoud
- TidyRoll
- Tokens
- TurretRange
- Twister
- Vectors
- WarTriage
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- nLootLink
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: GetText -> TextEditBoxGetText(self.name)
- ActionBarHide: GetText -> TextEditBoxGetText(self.name)
- AdvancedPetAssist: ParseRGB -> TextEditBoxGetText(name)
- AdvancedRenownTrainer: ImportNameInputOkButtonPressed -> TextEditBoxGetText(ImportNameInputWindowName.."NameInputBox")
- AdvancedRenownTrainer: ImportOkButtonPressed -> TextEditBoxGetText(ImportWindowName.."NameInputBox")
- AdvancedRenownTrainer: ImportOkButtonPressed -> TextEditBoxGetText(ImportWindowName.."LinkInputBox")

## Related APIs

- [OnKeyEscape](../../xml/handlers/handler_OnKeyEscape.md) (HIGH 88/100) - XML Event

## Used With

- [LabelGetText](window_LabelGetText.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function

## Notes

- none
