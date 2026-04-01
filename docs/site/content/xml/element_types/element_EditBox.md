# EditBox

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, Aura, AutoBand, BuffHead, Busted, EA_UiDebugTools, EA_UiModWindow, Enemy |
| Files seen in | `/workspace/Aura/Source/AuraConfig.xml:1096`, `/workspace/Aura/Source/AuraConfig.xml:1137`, `/workspace/Aura/Source/AuraConfig.xml:116`, `/workspace/Aura/Source/AuraConfig.xml:1202`, `/workspace/Aura/Source/AuraConfig.xml:1243`, `/workspace/Aura/Source/AuraConfig.xml:1347`, `/workspace/Aura/Source/AuraConfig.xml:1471`, `/workspace/Aura/Source/AuraConfig.xml:1559` |
| Namespaces detected | EditBox |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindowNameInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowLinkInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowNameInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingLinkWindowLinkBox, AdvancedRenownTrainer: AdvancedRenownTrainingPresetsWindowSaveNameInput, Aura: AuraConfigActivationAlertTextText |
| XML usage count | 215 |
| XML attribute usage count | 215 |
| Lua usage count | 6 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
| Consistent role | yes |
| Consistent arguments | no |
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

## Description

Observed XML element type instantiated by 26 addons.

## Common Attributes

- name
- inherits
- maxchars
- input
- layer
- warnOnTextCropped
- font
- taborder
- handleinput
- scrolling
- maxChars
- background

## Common Handlers

- OnTextChanged
- OnKeyEnter
- OnKeyEscape
- OnMouseOver
- OnKeyTab
- OnShown

## Common Inherits

- EA_EditBox_DefaultFrame
- Aura_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame_Multiline
- IraConfigNumEdit
- RVAPI_ColorDialogEditBoxTemplate
- IraConfigNumSpin
- Shinies_GoldCoin_EditBox_DefaultFrame
- MapMonsterEditorWindowEditBoxCoord
- MapMonsterPinTypeEditorWindowEditBoxDefault
- Shinies_BrassCoin_EditBox_DefaultFrame
- Shinies_SilverCoin_EditBox_DefaultFrame
- EA_EditBox_NoFrame

## Seen In

- AdvancedRenownTrainer
- Aura
- AutoBand
- BuffHead
- Busted
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- Killer
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- ObjectInspector
- Pocket Palette
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- Shinies
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyRoll
- TurretRange
- bigger_MacroWindow
- wbLeadHelper

## Examples

- AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindowNameInputBox -> EditBox AdvancedRenownTrainingImportNameInputWindowNameInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowLinkInputBox -> EditBox AdvancedRenownTrainingImportWindowLinkInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowNameInputBox -> EditBox AdvancedRenownTrainingImportWindowNameInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingLinkWindowLinkBox -> EditBox AdvancedRenownTrainingLinkWindowLinkBox
- AdvancedRenownTrainer: AdvancedRenownTrainingPresetsWindowSaveNameInput -> EditBox AdvancedRenownTrainingPresetsWindowSaveNameInput
- Aura: AuraConfigActivationAlertTextText -> EditBox AuraConfigActivationAlertTextText

## Related APIs

- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [TextEditBoxSetTextColor](../../window_api/functions/window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [TextEditBoxInsertText](../../window_api/functions/window_TextEditBoxInsertText.md) (HIGH 80/100) - Window Function

## Used With

- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [OnKeyEnter](../handlers/handler_OnKeyEnter.md) (HIGH 100/100) - XML Handler
- [OnKeyEnter](../../events/window_events/window_event_OnKeyEnter.md) (HIGH 100/100) - Window Event
- [OnKeyEscape](../../events/window_events/window_event_OnKeyEscape.md) (HIGH 100/100) - Window Event
- [OnKeyEscape](../handlers/handler_OnKeyEscape.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [OnShown](../handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [OnTextChanged](../../events/window_events/window_event_OnTextChanged.md) (HIGH 100/100) - Window Event
- [OnTextChanged](../handlers/handler_OnTextChanged.md) (HIGH 100/100) - XML Handler
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](../../window_api/functions/window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [OnKeyTab](../handlers/handler_OnKeyTab.md) (HIGH 93/100) - XML Handler
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [OnKeyTab](../../events/window_events/window_event_OnKeyTab.md) (HIGH 73/100) - Window Event

## Triggered By

- none

## Affects

- none
