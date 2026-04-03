# OnHidden

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BlackBook, BuffHead, CCTV, CM_ClosetGoblin, Calling |
| Namespaces detected | OnHidden |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedPetAssist: .OnHidden, AdvancedRenownTrainer: .OnHidden, Aura: .OnHidden, BlackBook: .OnHidden, BuffHead: .OnHidden, CCTV: .OnHidden |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 100 |
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

XML handler event observed across 35 addons.

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
- BuffHead
- CCTV
- CM_ClosetGoblin
- Calling
- CastSequence
- Crusher
- DPSMeter
- Enemy
- GroupRange
- Keyset
- KeysetMonsterPlay
- LibAddonButton
- LoyalPet
- MapMonster
- Motion
- Obsidian
- Phantom
- RandomMount
- SOR
- Sequencer
- Shinies
- Statdoll Remix
- TexturedButtons
- TidyChat
- TidyRoll
- TomeTracker
- TurretRange
- WARCommander
- WSCT
- WhoHealedMe
- nLootLink

## Examples

- AdvancedPetAssist: .OnHidden -> APAGui.OnHidden
- AdvancedPetAssist: .OnHidden -> APAGui.OnFollowTargetHUDHidden
- AdvancedPetAssist: .OnHidden -> APAGui.OnKitingHUDHidden
- AdvancedPetAssist: .OnHidden -> APAGui.OnInstantOnlyHUDHidden
- AdvancedRenownTrainer: .OnHidden -> AdvancedRenownTraining.OnImportHidden
- AdvancedRenownTrainer: .OnHidden -> AdvancedRenownTraining.OnExportHidden

## Related APIs

- [ButtonSetDisabledFlag](../../window_api/functions/window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
