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
| Addons seen in | AggroMeter, AnywhereTrainer, AnywhereTrainerAdditions, Aura, BuffHead, Busted, CM_ClosetGoblin, CMap |
| Files seen in | `/workspace/AggroMeter/AggroMeter.xml:135`, `/workspace/AggroMeter/AggroMeter.xml:8`, `/workspace/AnywhereTrainer/source/AnywhereTrainer.xml:50`, `/workspace/AnywhereTrainerAdditions/AnywhereTrainerAdditions.xml:6`, `/workspace/Aura/Source/AuraSettings.xml:28`, `/workspace/Aura/Source/AuraShares.xml:81`, `/workspace/BuffHead/Setup/General.xml:34`, `/workspace/BuffHead/Setup/SetupAdvancedCompression.xml:46` |
| Namespaces detected | OnRButtonUp |
| Source kinds | bindings, xml_handlers |
| Example locations | AggroMeter: AggroMeterGrayListBox.OnRButtonUp, AggroMeter: AggroMeter_Button.OnRButtonUp, AnywhereTrainer: AnywhereTrainerTabTemplate.OnRButtonUp, AnywhereTrainerAdditions: AnywhereTrainerAdditionsTabTemplate.OnRButtonUp, Aura: AuraSharesRow.OnRButtonUp, Aura: AuraWindowRow.OnRButtonUp |
| XML usage count | 150 |
| XML attribute usage count | 150 |
| Lua usage count | 147 |
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

Observed as an XML handler hook bound by 35 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Button
- ComboBox
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
- Busted
- CM_ClosetGoblin
- CMap
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- JunkDump
- Killer
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrowLight
- MoraleCircle
- PeaceOut
- Pocket Palette
- PotionBar
- Queue Queuer
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- ThinkOutLoud
- TidyQueue
- TidyRoll
- Tortall_DPS
- TurretRange
- WarBoard
- followTheLeader
- wbLeadHelper

## Examples

- AggroMeter: AggroMeterGrayListBox -> AggroMeterGrayListBox.OnRButtonUp -> AggroMeter.PickedListMenu
- AggroMeter: AggroMeter_Button -> AggroMeter_Button.OnRButtonUp -> AggroMeter.OnTabRBU
- AnywhereTrainer: AnywhereTrainerTabTemplate -> AnywhereTrainerTabTemplate.OnRButtonUp -> AnywhereTrainer.OnRButtonUp
- AnywhereTrainerAdditions: AnywhereTrainerAdditionsTabTemplate -> AnywhereTrainerAdditionsTabTemplate.OnRButtonUp -> AnywhereTrainerAdditions.OnRButtonUp
- Aura: AuraSharesRow -> AuraSharesRow.OnRButtonUp -> AuraShares.OnRButtonUpAuraList
- Aura: AuraWindowRow -> AuraWindowRow.OnRButtonUp -> AuraSettings.OnRButtonUpAuraList

## Related APIs

- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](../../window_api/functions/window_WindowStopAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowGetMovable](../../window_api/functions/window_WindowGetMovable.md) (HIGH 90/100) - Window Function
- [PlayerMenuWindow.ShowMenu](../../globals/functions/global_PlayerMenuWindow.ShowMenu.md) (HIGH 88/100) - Global Function

## Used With

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](../element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](../element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](../element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
