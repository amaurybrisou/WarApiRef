# OnMouseOver

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
| Addons seen in | AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Aura, BagOMatic, BankArkel, BuffHead, CM_ClosetGoblin |
| Files seen in | `/workspace/AggroMeter/AggroMeter.xml:59`, `/workspace/AggroMeter/AggroMeter.xml:65`, `/workspace/AggroMeter/AggroMeter.xml:70`, `/workspace/AggroMeter/AggroMeter.xml:8`, `/workspace/AnywhereTrainer/source/AnywhereTrainer.xml:44`, `/workspace/Aura/Source/AuraConfig.xml:67`, `/workspace/Aura/Source/AuraSettings.xml:26`, `/workspace/Aura/Source/AuraShares.xml:79` |
| Namespaces detected | OnMouseOver |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedRenownTrainer: AbilityButtonTemplate.OnMouseOver, AggroMeter: AggroMeter_Button.OnMouseOver, AggroMeter: Aggro_Label_Template.OnMouseOver, AggroMeter: Aggro_Tactic_Template.OnMouseOver, AggroMeter: Aggro_Timer_Template.OnMouseOver, AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage.OnMouseOver |
| XML usage count | 485 |
| XML attribute usage count | 485 |
| Lua usage count | 481 |
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

Observed as an XML handler hook bound by 43 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Button
- ComboBox
- DynamicImage
- EditBox
- FullResizeImage
- Label
- ListBox
- MapDisplay
- SliderBar
- Window

## Seen In

- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Aura
- BagOMatic
- BankArkel
- BuffHead
- CM_ClosetGoblin
- CMap
- Cheeseboard
- DAoCBuff
- DeepSleep
- EA_ThreePartBar
- EA_UiDebugTools
- EA_UiModWindow
- Effigy
- Enemy
- JunkDump
- Killer
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- Pocket Palette
- PotionBar
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- bigger_MacroWindow
- followTheLeader
- wbLeadHelper

## Examples

- AdvancedRenownTrainer: AbilityButtonTemplate -> AbilityButtonTemplate.OnMouseOver -> AdvancedRenownTraining.AbilityTooltip
- AggroMeter: AggroMeter_Button -> AggroMeter_Button.OnMouseOver -> AggroMeter.OnMouseOverStart
- AggroMeter: Aggro_Label_Template -> Aggro_Label_Template.OnMouseOver -> AggroMeter.SelectChar
- AggroMeter: Aggro_Tactic_Template -> Aggro_Tactic_Template.OnMouseOver -> AggroMeter.OnMouseOverStart
- AggroMeter: Aggro_Timer_Template -> Aggro_Timer_Template.OnMouseOver -> AggroMeter.OnMouseOverStart
- AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage -> AnywhereTrainerTabTemplateInactiveImage.OnMouseOver -> AnywhereTrainer.OnMouseOver

## Related APIs

- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function

## Used With

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](../element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](../../globals/functions/global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [DynamicImage](../element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](../element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [GameData.PlayerActions.DO_ABILITY](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_ABILITY.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](../element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [SliderBar](../element_types/element_SliderBar.md) (HIGH 100/100) - XML Element Type
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because API_Ref captures symbol linkage, not full handler signatures.
