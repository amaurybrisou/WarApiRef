# OnInitialize

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 138

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, Aura, CM_ClosetGoblin, CMap, EA_UiDebugTools, EA_UiModWindow, FastInteract, LoyalPet |
| Files seen in | `/workspace_addons/Aura/Source/AuraConfig.xml:28`, `/workspace_addons/Aura/Source/Templates.xml:358`, `/workspace_addons/Aura/Source/Templates.xml:389`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:1405`, `/workspace_addons/FastInteract/FastInteract.xml:13`, `/workspace_addons/LoyalPet/gui/lpet_gui.xml:2646`, `/workspace_addons/MGRemix/MGRemix.xml:253`, `/workspace_addons/MGRemix/config.xml:378` |
| Namespaces detected | OnInitialize |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTrainingWindow.OnInitialize, Aura: AuraConfig.OnInitialize, Aura: Aura_LabelCheckButton.OnInitialize, Aura: Aura_LargeLabelCheckButton.OnInitialize, CM_ClosetGoblin: ClosetGoblinOptionWindow.OnInitialize, CMap: CMapWindow.OnInitialize |
| XML usage count | 31 |
| XML attribute usage count | 31 |
| Lua usage count | 31 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
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

Observed as an XML handler hook bound by 19 addons through frame event handlers.

## Expected Lua Binding

```lua
function()
```

## Element Types

- Button
- Window

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

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
