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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BuffHead, CM_ClosetGoblin, Enemy, RoR_SoR, Shinies |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItemEffect.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainers.xml:0` |
| Namespaces detected | OnHidden |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD.OnHidden, AdvancedPetAssist: APAInstantOnlyHUD.OnHidden, AdvancedPetAssist: APAKitingHUD.OnHidden, AdvancedPetAssist: APAOptions.OnHidden, AdvancedRenownTrainer: AdvancedRenownTrainingExportWindow.OnHidden, AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindow.OnHidden |
| XML usage count | 79 |
| XML attribute usage count | 79 |
| Lua usage count | 79 |
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

Observed as an XML handler hook bound by 15 addons through frame event handlers.

## Expected Lua Binding

```lua
function()
```

## Element Types

- Window

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BuffHead
- CM_ClosetGoblin
- Enemy
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WhoHealedMe
- bigger_MacroWindow

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
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Used With

- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
