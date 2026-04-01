# OnShown

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, CM_ClosetGoblin, CMap, EA_UiDebugTools, EA_UiModWindow, Enemy |
| Files seen in | `/workspace/AdvancedPetAssist/APAGui.xml:1384`, `/workspace/AdvancedPetAssist/APAGui.xml:1425`, `/workspace/AdvancedPetAssist/APAGui.xml:1466`, `/workspace/AdvancedPetAssist/APAGui.xml:1506`, `/workspace/AdvancedPetAssist/APAGui.xml:95`, `/workspace/Aura/Source/AuraTexture.xml:100`, `/workspace/ClosetGoblin/ClosetGoblin.xml:1226`, `/workspace/ClosetGoblin/ClosetGoblin.xml:262` |
| Namespaces detected | OnShown |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD.OnShown, AdvancedPetAssist: APAInstantOnlyHUD.OnShown, AdvancedPetAssist: APAKitingHUD.OnShown, AdvancedPetAssist: APAOptions.OnShown, AdvancedPetAssist: APAPetTargetHUD.OnShown, AdvancedRenownTrainer: AdvancedRenownTrainingExportWindow.OnShown |
| XML usage count | 65 |
| XML attribute usage count | 65 |
| Lua usage count | 65 |
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

Observed as an XML handler hook bound by 23 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- EditBox
- MapDisplay
- Window

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

## Examples

- AdvancedPetAssist: APAFollowTargetHUD -> APAFollowTargetHUD.OnShown -> APAGui.OnFollowTargetHUDShown
- AdvancedPetAssist: APAInstantOnlyHUD -> APAInstantOnlyHUD.OnShown -> APAGui.OnInstantOnlyHUDShown
- AdvancedPetAssist: APAKitingHUD -> APAKitingHUD.OnShown -> APAGui.OnKitingHUDShown
- AdvancedPetAssist: APAOptions -> APAOptions.OnShown -> APAGui.OnShown
- AdvancedPetAssist: APAPetTargetHUD -> APAPetTargetHUD.OnShown -> APAGui.OnPetTargetHUDShown
- AdvancedRenownTrainer: AdvancedRenownTrainingExportWindow -> AdvancedRenownTrainingExportWindow.OnShown -> AdvancedRenownTraining.OnExportShown

## Related APIs

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetDisabledFlag](../../window_api/functions/window_ComboBoxSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because API_Ref captures symbol linkage, not full handler signatures.
