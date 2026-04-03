# OnSelChanged

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 158

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BankArkel, BuffHead, DAoCBuff, Enemy, Killer |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/BankArkel/BankArkel.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupContainer.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupDisplay.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupGeneral.xml:0` |
| Namespaces detected | OnSelChanged |
| Source kinds | bindings, examples, xml_handlers |
| Example locations | AdvancedPetAssist: APAComboAttackBind.OnSelChanged, AdvancedPetAssist: APAComboAutoReattack.OnSelChanged, AdvancedPetAssist: APAComboAutoReattackDelay.OnSelChanged, AdvancedPetAssist: APAComboCastDelay.OnSelChanged, AdvancedPetAssist: APAComboCastOnAcquire.OnSelChanged, AdvancedPetAssist: APAComboCombatExitDelay.OnSelChanged |
| XML usage count | 179 |
| XML attribute usage count | 179 |
| Lua usage count | 179 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 2 |
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

Observed as an XML handler hook bound by 17 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- ComboBox

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BankArkel
- BuffHead
- DAoCBuff
- Enemy
- Killer
- Pocket Palette
- PotionBar
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WhoHealedMe

## Examples

- AdvancedPetAssist: APAComboAttackBind -> APAComboAttackBind.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboAutoReattack -> APAComboAutoReattack.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboAutoReattackDelay -> APAComboAutoReattackDelay.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboCastDelay -> APAComboCastDelay.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboCastOnAcquire -> APAComboCastOnAcquire.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboCombatExitDelay -> APAComboCombatExitDelay.OnSelChanged -> APAGui.OnComboChanged

## Related APIs

- [ComboBox](../element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function

## Used With

- [ComboBox](../element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnSelChanged](../../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
