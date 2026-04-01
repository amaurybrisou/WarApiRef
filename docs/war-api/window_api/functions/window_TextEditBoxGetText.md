# TextEditBoxGetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 29 addons

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
| Files seen in | `/workspace/Ace/LibGUI.lua:664`, `/workspace/Ace/LibGUI.lua:719`, `/workspace/AdvancedPetAssist/APAGuiHelpers.lua:9`, `/workspace/Aura/Source/AuraShares.lua:419`, `/workspace/BuffHead/Setup/SelectColor.lua:78`, `/workspace/BuffHead/Setup/SetupAdvancedCompressionItem.lua:310`, `/workspace/BuffHead/Setup/SetupAdvancedCompressionItemEffect.lua:147`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:601` |
| Namespaces detected | TextEditBoxGetText |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_MultiTextbox:GetText, Ace: LIBGUI_Textbox:GetText, AdvancedPetAssist: APAGuiHelpers.ParseRGB, AdvancedRenownTrainer: AdvancedRenownTraining.ImportNameInputOkButtonPressed, AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed, AdvancedRenownTrainer: AdvancedRenownTraining.SavePreset |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 196 |
| Global usage count | 196 |
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
- ThinkOutLoud
- TidyRoll
- TurretRange
- WarTriage
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
- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EditBox](../../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnTextChanged](../../events/window_events/window_event_OnTextChanged.md) (HIGH 100/100) - Window Event
- [OnTextChanged](../../xml/handlers/handler_OnTextChanged.md) (HIGH 100/100) - XML Handler
- [ScrollWindowUpdateScrollRect](window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowGetId](window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
