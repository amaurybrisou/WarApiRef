# OnShown

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 188

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, CM_ClosetGoblin, CMap, EA_UiDebugTools, EA_UiModWindow, Enemy |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:1384`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1425`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1466`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1506`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:95`, `/workspace_addons/Aura/Source/AuraTexture.xml:100`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:1226`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:262` |
| Namespaces detected | OnShown |
| Source kinds | event_page, flows, lua_event_registration, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD.OnShown, AdvancedPetAssist: APAInstantOnlyHUD.OnShown, AdvancedPetAssist: APAKitingHUD.OnShown, AdvancedPetAssist: APAOptions.OnShown, AdvancedPetAssist: APAPetTargetHUD.OnShown, AdvancedRenownTrainer: AdvancedRenownTrainingExportWindow.OnShown |
| XML usage count | 66 |
| XML attribute usage count | 66 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 3 |
| Initialization flow references | 1 |
| Known engine namespace | no |
| Default UI presence | no |
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

Observed as an engine-supplied UI event hook used by 23 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- CM_ClosetGoblin
- CMap
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- LoyalPet
- MapMonster
- MapPin
- Pocket Palette
- PotionBar
- RandomMount
- RoR_SoR
- Shinies
- TidyChat
- TidyRoll
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- wbLeadHelper

## Registrars And Handlers

- APAGui.OnFollowTargetHUDShown
- APAGui.OnInstantOnlyHUDShown
- APAGui.OnKitingHUDShown
- APAGui.OnPetTargetHUDShown
- APAGui.OnShown
- AdvancedRenownTraining.OnExportShown
- AdvancedRenownTraining.OnShown
- AdvancedRenownTraining.PresetOnShown
- AuraTexture.OnShown
- ClosetGoblinCharacterWindow.OnShow
- ClosetGoblinZoneWindow.OnShow
- DebugWindow.OnShowCopy
- DebugWindow.OnShowFocus
- DebugWindow.OnShown
- DevPadWindow.OnShown
- EA_Window_Macro.OnShown
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnShown
- LPET.OnShown
- MapMonster.Editor.OnShown
- MapMonster.PinTypeEditor.OnShown
- MapMonster.PinTypeEditor.OnShownChooser
- MapMonster_Calibrate.OnShown
- PP.OnShown
- PotionBarSettings.OnAboutShown
- RandomMountUI.OnShown
- Shinies.OnShown
- ShiniesAuctionsUI.OnShown
- ShiniesAutoUI.OnShown
- ShiniesBrowseUI.OnShown
- ShiniesPostUI.OnShown
- TidyChat.Copy.OnShown
- TidyChat.LootRoll.OnShown
- TidyChat.OnEntryBoxUpdateShowing
- TidyChat.Options.OnShown
- TidyRoll.CustomAutoRoll.OnShown
- TidyRollOptions.OnShown
- UiModVersionMismatchWindow.OnShown
- UiModWindow.OnAdvancedShown
- UiModWindow.OnShown
- WHMGui.OnDetailsShown
- WHMGui.OnOptionsShown
- WHMGui.OnWindowShown
- WSCT.OnShown
- WarBoard.Options.ShowTopOptions
- WindowRegisterCoreEventHandler
- WindowUtils.OnShown
- core

## Examples

- TidyChat: TidyChatFrames.Initialize -> OnShown -> TidyChat.OnEntryBoxUpdateShowing
- AdvancedPetAssist: APAFollowTargetHUD -> APAFollowTargetHUD.OnShown -> APAGui.OnFollowTargetHUDShown
- AdvancedPetAssist: APAInstantOnlyHUD -> APAInstantOnlyHUD.OnShown -> APAGui.OnInstantOnlyHUDShown
- AdvancedPetAssist: APAKitingHUD -> APAKitingHUD.OnShown -> APAGui.OnKitingHUDShown
- AdvancedPetAssist: APAOptions -> APAOptions.OnShown -> APAGui.OnShown
- AdvancedPetAssist: APAPetTargetHUD -> APAPetTargetHUD.OnShown -> APAGui.OnPetTargetHUDShown

## Related APIs

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetDisabledFlag](../../window_api/functions/window_ComboBoxSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [EditBox](../../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [MapDisplay](../../xml/element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none
