# TextEditBoxGetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 27 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, Aura, BuffHead, DAoCBuff, EA_UiDebugTools, Effigy |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:664`, `/workspace_addons/Ace/LibGUI.lua:719`, `/workspace_addons/AdvancedPetAssist/APAGuiHelpers.lua:9`, `/workspace_addons/Aura/Source/AuraShares.lua:419`, `/workspace_addons/BuffHead/Setup/SelectColor.lua:78`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItem.lua:310`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItemEffect.lua:147`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:601` |
| Namespaces detected | TextEditBoxGetText |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_MultiTextbox:GetText, Ace: LIBGUI_Textbox:GetText, AdvancedPetAssist: APAGuiHelpers.ParseRGB, AdvancedRenownTrainer: AdvancedRenownTraining.ImportNameInputOkButtonPressed, AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed, AdvancedRenownTrainer: AdvancedRenownTraining.SavePreset |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 190 |
| Global usage count | 190 |
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

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BuffHead
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- JunkDump
- Killer
- LibWBToggler
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- Pocket Palette
- RVAPI_ColorDialog
- RandomMount
- Shinies
- TexturedButtons
- TidyRoll
- TurretRange
- WhoHealedMe
- WoH-Reticle
- wbLeadHelper

## Examples

- Ace: LIBGUI_MultiTextbox:GetText -> TextEditBoxGetText(self.name)
- Ace: LIBGUI_Textbox:GetText -> TextEditBoxGetText(self.name)
- AdvancedPetAssist: APAGuiHelpers.ParseRGB -> TextEditBoxGetText(name)
- AdvancedRenownTrainer: AdvancedRenownTraining.ImportNameInputOkButtonPressed -> TextEditBoxGetText(ImportNameInputWindowName.."NameInputBox")
- AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed -> TextEditBoxGetText(ImportWindowName.."NameInputBox")
- AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed -> TextEditBoxGetText(ImportWindowName.."LinkInputBox")

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
