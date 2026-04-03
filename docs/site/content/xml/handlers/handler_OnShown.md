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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, CM_ClosetGoblin, Enemy, Pocket Palette, PotionBar, RoR_SoR |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/Aura/Source/AuraTexture.xml:0`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.xml:0`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogSnapshotWindow.xml:0`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogStatsWindow.xml:0`, `/workspace/data/raw/Enemy/Code/Core/ChooseIconDialog.xml:0`, `/workspace/data/raw/Enemy/Code/Core/Common.xml:0`, `/workspace/data/raw/Enemy/Code/Core/ConfigDialog.xml:0` |
| Namespaces detected | OnShown |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD.OnShown, AdvancedPetAssist: APAInstantOnlyHUD.OnShown, AdvancedPetAssist: APAKitingHUD.OnShown, AdvancedPetAssist: APAOptions.OnShown, AdvancedPetAssist: APAPetTargetHUD.OnShown, AdvancedRenownTrainer: AdvancedRenownTrainingExportWindow.OnShown |
| XML usage count | 47 |
| XML attribute usage count | 47 |
| Lua usage count | 47 |
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
- CM_ClosetGoblin
- Enemy
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- TidyChat
- TidyRoll
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow

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
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Used With

- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
