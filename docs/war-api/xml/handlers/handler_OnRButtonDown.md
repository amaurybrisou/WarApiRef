# OnRButtonDown

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
| Addons seen in | BuffHead, CM_ClosetGoblin, Enemy, TexturedButtons, TurretRange, bigger_MacroWindow |
| Files seen in | `/workspace/data/raw/BuffHead/Display.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainers.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayout.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupMenu.xml:0` |
| Namespaces detected | OnRButtonDown |
| Source kinds | bindings, xml_handlers |
| Example locations | BuffHead: BuffHeadBuffTemplate.OnRButtonDown, BuffHead: BuffHeadLayoutControlFrameWindow.OnRButtonDown, BuffHead: BuffHeadLayoutFrameWindow.OnRButtonDown, BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate.OnRButtonDown, BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate.OnRButtonDown, BuffHead: BuffHeadSetupAdvancedContainersItemRowTemplate.OnRButtonDown |
| XML usage count | 21 |
| XML attribute usage count | 21 |
| Lua usage count | 21 |
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

Observed as an XML handler hook bound by 6 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Button
- Window

## Seen In

- BuffHead
- CM_ClosetGoblin
- Enemy
- TexturedButtons
- TurretRange
- bigger_MacroWindow

## Examples

- BuffHead: BuffHeadBuffTemplate -> BuffHeadBuffTemplate.OnRButtonDown -> FrameManager.OnRButtonDown
- BuffHead: BuffHeadLayoutControlFrameWindow -> BuffHeadLayoutControlFrameWindow.OnRButtonDown -> BuffHead.Setup.Layout.OnControlFrameRButtonDown
- BuffHead: BuffHeadLayoutFrameWindow -> BuffHeadLayoutFrameWindow.OnRButtonDown -> BuffHead.Setup.Layout.OnLayoutWindowRButtonDown
- BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate -> BuffHeadSetupAdvancedCompressionItemRowTemplate.OnRButtonDown -> BuffHead.Setup.AdvancedCompressionItem.OnRowRDown
- BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate -> BuffHeadSetupAdvancedCompressionRowTemplate.OnRButtonDown -> BuffHead.Setup.AdvancedCompression.OnRowRDown
- BuffHead: BuffHeadSetupAdvancedContainersItemRowTemplate -> BuffHeadSetupAdvancedContainersItemRowTemplate.OnRButtonDown -> BuffHead.Setup.AdvancedContainersItem.OnRowRDown

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Used With

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
