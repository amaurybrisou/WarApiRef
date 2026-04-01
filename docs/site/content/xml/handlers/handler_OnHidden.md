# OnHidden

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BuffHead, CM_ClosetGoblin, CMap, EA_UiDebugTools, EA_UiModWindow |
| Files seen in | `/workspace/AdvancedPetAssist/APAGui.xml:1385`, `/workspace/AdvancedPetAssist/APAGui.xml:1426`, `/workspace/AdvancedPetAssist/APAGui.xml:1467`, `/workspace/AdvancedPetAssist/APAGui.xml:96`, `/workspace/Aura/Source/AuraConfig.xml:27`, `/workspace/Aura/Source/AuraSettings.xml:78`, `/workspace/Aura/Source/AuraShares.xml:153`, `/workspace/BuffHead/Setup/SetupAdvancedCompression.xml:54` |
| Namespaces detected | OnHidden |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD.OnHidden, AdvancedPetAssist: APAInstantOnlyHUD.OnHidden, AdvancedPetAssist: APAKitingHUD.OnHidden, AdvancedPetAssist: APAOptions.OnHidden, AdvancedRenownTrainer: AdvancedRenownTrainingExportWindow.OnHidden, AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindow.OnHidden |
| XML usage count | 95 |
| XML attribute usage count | 95 |
| Lua usage count | 95 |
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

- MapDisplay
- Window

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BuffHead
- CM_ClosetGoblin
- CMap
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- LoyalPet
- MapMonster
- MapPin
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WhoHealedMe
- bigger_MacroWindow
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAFollowTargetHUD -> APAFollowTargetHUD.OnHidden -> APAGui.OnFollowTargetHUDHidden
- AdvancedPetAssist: APAInstantOnlyHUD -> APAInstantOnlyHUD.OnHidden -> APAGui.OnInstantOnlyHUDHidden
- AdvancedPetAssist: APAKitingHUD -> APAKitingHUD.OnHidden -> APAGui.OnKitingHUDHidden
- AdvancedPetAssist: APAOptions -> APAOptions.OnHidden -> APAGui.OnHidden
- AdvancedRenownTrainer: AdvancedRenownTrainingExportWindow -> AdvancedRenownTrainingExportWindow.OnHidden -> AdvancedRenownTraining.OnExportHidden
- AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindow -> AdvancedRenownTrainingImportNameInputWindow.OnHidden -> AdvancedRenownTraining.OnImportNameInputHidden

## Related APIs

- [ButtonSetDisabledFlag](../../window_api/functions/window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxSetDisabledFlag](../../window_api/functions/window_ComboBoxSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [BroadcastEvent](../../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function

## Used With

- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [OnHidden](../../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because API_Ref captures symbol linkage, not full handler signatures.
