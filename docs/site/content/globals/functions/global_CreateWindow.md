# CreateWindow

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 22 addons

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, AutoMark, BagOMatic, BankArkel, CM_ClosetGoblin |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.lua:1063`, `/workspace/data/raw/AggroMeter/AggroMeter.lua:5`, `/workspace/data/raw/Aura/Source/AuraTooltip.lua:21`, `/workspace/data/raw/AutoMark/Source/AutoMark.lua:33`, `/workspace/data/raw/BankArkel/BankArkel.lua:172`, `/workspace/data/raw/BankArkel/BankArkel.lua:95`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.lua:87`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:752` |
| Namespaces detected | CreateWindow |
| Source kinds | lua_calls |
| Example locations | AdvancedPetAssist: APAGui.Show, AdvancedRenownTrainer: AdvancedRenownTraining.Initialize, AdvancedRenownTrainer: AdvancedRenownTraining.InitializeImportExport, AggroMeter: AggroMeter.Initialize, Aura: AuraTooltip.OnInitialize, AutoMark: AutoMark.OnInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 87 |
| Global usage count | 87 |
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
| showOnCreate | Observed as a boolean visibility flag. | Observed values: false, true |

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
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- GuardLine
- Killer
- MiracleGrowLight
- MoraleCircle
- PotionBar
- RoR_SoR
- Shinies
- Swift Assist
- TidyChat
- TidyRoll
- WSCT
- followTheLeader

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

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow](../tables/table_EA_ChatWindow.md) (HIGH 100/100) - Global Table
- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [RegisterEventHandler](global_RegisterEventHandler.md) (HIGH 93/100) - Global Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Triggered By

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Event
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
- Advanced return analysis: No strong return evidence observed
