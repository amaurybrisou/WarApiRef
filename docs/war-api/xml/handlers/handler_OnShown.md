# OnShown

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 88/100

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, referenced by generated docs or reference files, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BlackBook, CCTV, CM_ClosetGoblin, CaVES, Calling |
| Namespaces detected | OnShown |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedPetAssist: .OnShown, AdvancedRenownTrainer: .OnShown, Aura: .OnShown, BlackBook: .OnShown, CCTV: .OnShown, CM_ClosetGoblin: .OnShown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 66 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
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

XML handler event observed across 38 addons.

## Expected Lua Binding

```lua
function()
```

## Element Types

- none

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BlackBook
- CCTV
- CM_ClosetGoblin
- CaVES
- Calling
- Crusher
- DPSMeter
- EA_OpenPartyWindow
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- FozAuction
- LoyalPet
- MapMonster
- Motion
- Phantom
- Pocket Palette
- PotionBar
- RandomMount
- ReliquaryHunter
- Sequencer
- Shinies
- SocialWindow 2.0
- TidyChat
- TidyRoll
- Tome Titan
- TomeTracker
- WARCommander
- WSCT
- WarBoard
- WarBoard_FPS
- WhoHealedMe
- alertMod
- bigger_MacroWindow
- nLootLink

## Examples

- AdvancedPetAssist: .OnShown -> APAGui.OnShown
- AdvancedPetAssist: .OnShown -> APAGui.OnFollowTargetHUDShown
- AdvancedPetAssist: .OnShown -> APAGui.OnKitingHUDShown
- AdvancedPetAssist: .OnShown -> APAGui.OnInstantOnlyHUDShown
- AdvancedPetAssist: .OnShown -> APAGui.OnPetTargetHUDShown
- AdvancedRenownTrainer: .OnShown -> AdvancedRenownTraining.OnExportShown

## Related APIs

- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [EA_Window_InteractionRenownTraining.GetPointsAvailable](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsAvailable.md) (HIGH 98/100) - Global Function
- [EA_Window_InteractionRenownTraining.GetPointsSpent](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsSpent.md) (HIGH 98/100) - Global Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
