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
| Addons seen in | AdvancedRenownTrainer, Aura, CM_ClosetGoblin, DAoCBuff, Shinies, TidyRoll, WSCT, WhoHealedMe |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1029`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1075`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1115`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1155`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1195` |
| Namespaces detected | OnInitialize |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTrainingWindow.OnInitialize, Aura: AuraConfig.OnInitialize, Aura: Aura_LabelCheckButton.OnInitialize, Aura: Aura_LargeLabelCheckButton.OnInitialize, CM_ClosetGoblin: ClosetGoblinOptionWindow.OnInitialize, DAoCBuff: DAoCBuffFrameSettingsTab_ScrollChild_ActiveCheckBox.OnInitialize |
| XML usage count | 37 |
| XML attribute usage count | 37 |
| Lua usage count | 37 |
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

Observed as an XML handler hook bound by 9 addons through frame event handlers.

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
- DAoCBuff
- Shinies
- TidyRoll
- WSCT
- WhoHealedMe
- bigger_MacroWindow

## Examples

- AdvancedRenownTrainer: AdvancedRenownTrainingWindow -> AdvancedRenownTrainingWindow.OnInitialize -> AdvancedRenownTraining.Initialize
- Aura: AuraConfig -> AuraConfig.OnInitialize -> AuraConfig.OnInitialize
- Aura: Aura_LabelCheckButton -> Aura_LabelCheckButton.OnInitialize -> EA_LabelCheckButton.Initialize
- Aura: Aura_LargeLabelCheckButton -> Aura_LargeLabelCheckButton.OnInitialize -> EA_LabelCheckButton.Initialize
- CM_ClosetGoblin: ClosetGoblinOptionWindow -> ClosetGoblinOptionWindow.OnInitialize -> ClosetGoblinOptionWindow.OnInitialize
- DAoCBuff: DAoCBuffFrameSettingsTab_ScrollChild_ActiveCheckBox -> DAoCBuffFrameSettingsTab_ScrollChild_ActiveCheckBox.OnInitialize -> EA_LabelCheckButton.Initialize

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
