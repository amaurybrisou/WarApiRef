# OnRButtonUp

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
| Addons seen in | AggroMeter, AnywhereTrainer, AnywhereTrainerAdditions, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.xml:0`, `/workspace/data/raw/AnywhereTrainerAdditions/AnywhereTrainerAdditions.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0` |
| Namespaces detected | OnRButtonUp |
| Source kinds | bindings, xml_handlers |
| Example locations | AggroMeter: AggroMeterGrayListBox.OnRButtonUp, AggroMeter: AggroMeter_Button.OnRButtonUp, AnywhereTrainer: AnywhereTrainerTabTemplate.OnRButtonUp, AnywhereTrainerAdditions: AnywhereTrainerAdditionsTabTemplate.OnRButtonUp, Aura: AuraSharesRow.OnRButtonUp, Aura: AuraWindowRow.OnRButtonUp |
| XML usage count | 108 |
| XML attribute usage count | 108 |
| Lua usage count | 108 |
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

Observed as an XML handler hook bound by 20 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Button
- DynamicImage
- Label
- ListBox
- MapDisplay
- Window

## Seen In

- AggroMeter
- AnywhereTrainer
- AnywhereTrainerAdditions
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Killer
- MiracleGrowLight
- MoraleCircle
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TidyRoll
- TurretRange
- WarBoard
- followTheLeader

## Examples

- AggroMeter: AggroMeterGrayListBox -> AggroMeterGrayListBox.OnRButtonUp -> AggroMeter.PickedListMenu
- AggroMeter: AggroMeter_Button -> AggroMeter_Button.OnRButtonUp -> AggroMeter.OnTabRBU
- AnywhereTrainer: AnywhereTrainerTabTemplate -> AnywhereTrainerTabTemplate.OnRButtonUp -> AnywhereTrainer.OnRButtonUp
- AnywhereTrainerAdditions: AnywhereTrainerAdditionsTabTemplate -> AnywhereTrainerAdditionsTabTemplate.OnRButtonUp -> AnywhereTrainerAdditions.OnRButtonUp
- Aura: AuraSharesRow -> AuraSharesRow.OnRButtonUp -> AuraShares.OnRButtonUpAuraList
- Aura: AuraWindowRow -> AuraWindowRow.OnRButtonUp -> AuraSettings.OnRButtonUpAuraList

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](../element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EA_Window_ContextMenu.AddCascadingMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](../element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetMovable](../../window_api/functions/window_WindowGetMovable.md) (HIGH 90/100) - Window Function

## Used With

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](../element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
