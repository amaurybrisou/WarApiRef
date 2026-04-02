# OnMouseOverEnd

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
| Addons seen in | BuffHead, CM_ClosetGoblin, CMap, Cheeseboard, DAoCBuff, Enemy, MapMonster, Miracle Grow Remix |
| Files seen in | `/workspace_addons/BuffHead/Display.xml:29`, `/workspace_addons/BuffHead/Setup/General.xml:36`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompression.xml:48`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItem.xml:47`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainers.xml:48`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItem.xml:47`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:47`, `/workspace_addons/BuffHead/Setup/SetupEffectCache.xml:53` |
| Namespaces detected | OnMouseOverEnd |
| Source kinds | bindings, xml_handlers |
| Example locations | BuffHead: BuffHeadBuffTemplate.OnMouseOverEnd, BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate.OnMouseOverEnd, BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate.OnMouseOverEnd, BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowTitleLabel.OnMouseOverEnd, BuffHead: BuffHeadSetupAdvancedContainersItemRowTemplate.OnMouseOverEnd, BuffHead: BuffHeadSetupAdvancedContainersRowTemplate.OnMouseOverEnd |
| XML usage count | 80 |
| XML attribute usage count | 80 |
| Lua usage count | 80 |
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

Observed as an XML handler hook bound by 17 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Button
- DynamicImage
- FullResizeImage
- Label
- ListBox
- MapDisplay
- Window

## Seen In

- BuffHead
- CM_ClosetGoblin
- CMap
- Cheeseboard
- DAoCBuff
- Enemy
- MapMonster
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- QuickWarReport
- RandomMount
- Shinies
- TexturedButtons
- TurretRange
- WarBoard
- wbLeadHelper

## Examples

- BuffHead: BuffHeadBuffTemplate -> BuffHeadBuffTemplate.OnMouseOverEnd -> FrameManager.OnMouseOverEnd
- BuffHead: BuffHeadSetupAdvancedCompressionItemRowTemplate -> BuffHeadSetupAdvancedCompressionItemRowTemplate.OnMouseOverEnd -> BuffHead.Setup.AdvancedCompressionItem.OnRowMouseOut
- BuffHead: BuffHeadSetupAdvancedCompressionRowTemplate -> BuffHeadSetupAdvancedCompressionRowTemplate.OnMouseOverEnd -> BuffHead.Setup.AdvancedCompression.OnRowMouseOut
- BuffHead: BuffHeadSetupAdvancedContainersItemPropertiesWindowTitleLabel -> BuffHeadSetupAdvancedContainersItemPropertiesWindowTitleLabel.OnMouseOverEnd -> BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOut
- BuffHead: BuffHeadSetupAdvancedContainersItemRowTemplate -> BuffHeadSetupAdvancedContainersItemRowTemplate.OnMouseOverEnd -> BuffHead.Setup.AdvancedContainersItem.OnRowMouseOut
- BuffHead: BuffHeadSetupAdvancedContainersRowTemplate -> BuffHeadSetupAdvancedContainersRowTemplate.OnMouseOverEnd -> BuffHead.Setup.AdvancedContainers.OnRowMouseOut

## Related APIs

- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function

## Used With

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](../element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [OnMouseOverEnd](../../events/window_events/window_event_OnMouseOverEnd.md) (HIGH 100/100) - Window Event

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
