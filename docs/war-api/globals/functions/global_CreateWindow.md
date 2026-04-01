# CreateWindow

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 44 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, AutoMark, BagOMatic, BankArkel, Busted |
| Files seen in | `/workspace/AdvancedPetAssist/APAGui.lua:1063`, `/workspace/AggroMeter/AggroMeter.lua:5`, `/workspace/Aura/Source/AuraTooltip.lua:21`, `/workspace/AutoMark/Source/AutoMark.lua:33`, `/workspace/BankArkel/BankArkel.lua:172`, `/workspace/BankArkel/BankArkel.lua:95`, `/workspace/Busted/Busted.lua:324`, `/workspace/Busted/Busted.lua:333` |
| Namespaces detected | CreateWindow |
| Source kinds | lua_calls |
| Example locations | AdvancedPetAssist: APAGui.Show, AdvancedRenownTrainer: AdvancedRenownTraining.Initialize, AdvancedRenownTrainer: AdvancedRenownTraining.InitializeImportExport, AggroMeter: AggroMeter.Initialize, Aura: AuraTooltip.OnInitialize, AutoMark: AutoMark.OnInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 130 |
| Global usage count | 130 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
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
CreateWindow(windowName, showOnCreate)
```

## Description

Observed creating a top-level XML window from a loaded definition.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a top-level window name. | Observed values: "APAOptions", "AdvancedRenownTrainingPresetsWindow", "AggroMeterGrayWindow" |
| showOnCreate | Observed as a boolean visibility flag. | Observed values: MiracleGrow.Settings.showing, false, true |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Creates or instantiates UI objects from XML-backed definitions.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- AutoMark
- BagOMatic
- BankArkel
- Busted
- CM_ClosetGoblin
- DAoCBuff
- DeepSleep
- EA_UiDebugTools
- Enemy
- FastInteract
- GetStats
- GuardLine
- JunkDump
- Killer
- LoyalPet
- MapMonster
- MegaphonePlusPlus
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- Moth
- PeaceOut
- PotionBar
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RetAlert
- RoR_SoR
- Shinies
- Swift Assist
- ThinkOutLoud
- TidyChat
- TidyRoll
- TimeInQueue
- Twister
- WSCT
- followTheLeader
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAGui.Show -> CreateWindow("APAOptions", true)
- AdvancedRenownTrainer: AdvancedRenownTraining.Initialize -> CreateWindow("AdvancedRenownTrainingPresetsWindow", false)
- AdvancedRenownTrainer: AdvancedRenownTraining.InitializeImportExport -> CreateWindow(ImportWindowName, false)
- AdvancedRenownTrainer: AdvancedRenownTraining.InitializeImportExport -> CreateWindow(ImportNameInputWindowName, false)
- AdvancedRenownTrainer: AdvancedRenownTraining.InitializeImportExport -> CreateWindow(ExportWindowName, false)
- AdvancedRenownTrainer: AdvancedRenownTraining.InitializeImportExport -> CreateWindow(LinkWindowName, false)

## Related APIs

- none

## Used With

- [BankWindow](../tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeOneButtonDialog](global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow](../tables/table_EA_ChatWindow.md) (HIGH 100/100) - Global Table
- [EA_ChatWindow.OnKeyEnter](global_EA_ChatWindow.OnKeyEnter.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [EA_Window_InteractionRenownTraining.Hide](global_EA_Window_InteractionRenownTraining.Hide.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining.Show](global_EA_Window_InteractionRenownTraining.Show.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetResolutionScale](global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [ScrollWindowSetOffset](../../window_api/functions/window_ScrollWindowSetOffset.md) (HIGH 100/100) - Window Function
- [ScrollWindowUpdateScrollRect](../../window_api/functions/window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Settings.Language.active](../../systemdata/fields/systemdata_SystemData.Settings.Language.active.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](../../window_api/functions/window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [TextLogGetUpdateEventId](../../events/game_events/game_event_TextLogGetUpdateEventId.md) (HIGH 100/100) - Game Event
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [BankWindow.Hide](global_BankWindow.Hide.md) (HIGH 88/100) - Global Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [TextEditBoxSetHistory](../../window_api/functions/window_TextEditBoxSetHistory.md) (HIGH 80/100) - Window Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event

## Affects

- [EA_ChatWindow](../tables/table_EA_ChatWindow.md) (HIGH 100/100) - Global Table
- [EA_Window_InteractionRenownTraining.Hide](global_EA_Window_InteractionRenownTraining.Hide.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining.Show](global_EA_Window_InteractionRenownTraining.Show.md) (HIGH 100/100) - Global Function
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
