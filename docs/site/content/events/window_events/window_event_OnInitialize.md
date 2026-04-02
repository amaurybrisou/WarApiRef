# OnInitialize

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 148

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, Aura, CM_ClosetGoblin, CMap, EA_UiDebugTools, EA_UiModWindow, FastInteract, LoyalPet |
| Files seen in | `/workspace_addons/Aura/Source/AuraConfig.xml:28`, `/workspace_addons/Aura/Source/Templates.xml:358`, `/workspace_addons/Aura/Source/Templates.xml:389`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:1405`, `/workspace_addons/FastInteract/FastInteract.xml:13`, `/workspace_addons/LoyalPet/gui/lpet_gui.xml:2646`, `/workspace_addons/MGRemix/MGRemix.xml:253`, `/workspace_addons/MGRemix/config.xml:378` |
| Namespaces detected | OnInitialize |
| Source kinds | event_page, flows, xml_handlers |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTrainingWindow.OnInitialize, Aura: AuraConfig.OnInitialize, Aura: Aura_LabelCheckButton.OnInitialize, Aura: Aura_LargeLabelCheckButton.OnInitialize, CM_ClosetGoblin: ClosetGoblinOptionWindow.OnInitialize, CMap: CMapWindow.OnInitialize |
| XML usage count | 31 |
| XML attribute usage count | 31 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 10 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | no |
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

Observed as an engine-supplied UI event hook used by 19 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- AdvancedRenownTrainer
- Aura
- CM_ClosetGoblin
- CMap
- EA_UiDebugTools
- EA_UiModWindow
- FastInteract
- LoyalPet
- MapMonster
- Miracle Grow Remix
- MiracleGrow
- RandomMount
- Shinies
- TidyRoll
- Tortall_DPS
- WSCT
- WhoHealedMe
- bigger_MacroWindow
- wbLeadHelper

## Registrars And Handlers

- AdvancedRenownTraining.Initialize
- AuraConfig.OnInitialize
- CMapWindow.Initialize
- CMapWindow.RefreshMapPointFilterMenu
- ClosetGoblinOptionWindow.OnInitialize
- DebugWindow.Initialize
- EA_GenericCheckButton.Initialize
- EA_LabelCheckButton.Initialize
- EA_Window_DefaultLabelToggleCircle.Initialize
- EA_Window_Macro.Initialize
- FastInteract.OptionsSetup
- IraConfig.HelpBtnInit
- LPET.OptionsOnInitialize
- MapMonster_Calibrate.WindowInit
- MiracleGrow.Initialize
- MiracleGrow2.Initialize
- RandomMountUI.OnInitialize
- TortallDPSMeter.Initialize
- UiModVersionMismatchWindow.Initialize
- UiModWindow.InitAdvancedWindow
- UiModWindow.Initialize
- WHMGui.OnDetailsInitialize
- WHMGui.OnOptionsInitialize
- WHMGui.OnWindowInitialize
- WSCT.ColorOnInitialize
- WSCT.OptionsOnInitialize
- wbLeadHelperConfigTab.Initialize
- wbLeadHelperConfigWindow.Initialize
- wbLeadHelperMessagesTab.Initialize

## Examples

- AdvancedRenownTrainer: AdvancedRenownTrainingWindow -> AdvancedRenownTrainingWindow.OnInitialize -> AdvancedRenownTraining.Initialize
- Aura: AuraConfig -> AuraConfig.OnInitialize -> AuraConfig.OnInitialize
- Aura: Aura_LabelCheckButton -> Aura_LabelCheckButton.OnInitialize -> EA_LabelCheckButton.Initialize
- Aura: Aura_LargeLabelCheckButton -> Aura_LargeLabelCheckButton.OnInitialize -> EA_LabelCheckButton.Initialize
- CM_ClosetGoblin: ClosetGoblinOptionWindow -> ClosetGoblinOptionWindow.OnInitialize -> ClosetGoblinOptionWindow.OnInitialize
- CMap: CMapWindow -> CMapWindow.OnInitialize -> CMapWindow.Initialize

## Related APIs

- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [EA_Window_InteractionRenownTraining.Hide](../../globals/functions/global_EA_Window_InteractionRenownTraining.Hide.md) (HIGH 100/100) - Global Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none
