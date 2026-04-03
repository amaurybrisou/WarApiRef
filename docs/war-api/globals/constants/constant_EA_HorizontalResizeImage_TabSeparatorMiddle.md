# EA_HorizontalResizeImage_TabSeparatorMiddle

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AggroMeter, Aura, EA_OpenPartyWindow, RaidMeter, TastyButtons |
| Files seen in | AggroMeter.xml, RaidMeter.xml, Source/AuraConfig.xml, Source/OpenPartyWindow.xml, TastyButtonsOptions.xml |
| Namespaces detected | EA_HorizontalResizeImage_TabSeparatorMiddle |
| Source kinds | xml_attributes |
| Example locations | AggroMeterGrayWindowSeparatorMiddle, AuraConfigTabsSeparatorActivation, AuraConfigTabsSeparatorDeactivation, AuraConfigTabsSeparatorTrigger, EA_Window_OpenPartyTabSeparatorMiddle, RaidMeterWindowSeparatorMiddle1 |
| XML usage count | 12 |
| XML attribute usage count | 12 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
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

Engine-supplied XML constant or template class referenced by 5 addons.

## Seen In

- AggroMeter
- Aura
- EA_OpenPartyWindow
- RaidMeter
- TastyButtons

## Used By

- AggroMeterGrayWindowSeparatorMiddle
- AuraConfigTabsSeparatorActivation
- AuraConfigTabsSeparatorDeactivation
- AuraConfigTabsSeparatorTrigger
- EA_Window_OpenPartyTabSeparatorMiddle
- RaidMeterWindowSeparatorMiddle1
- RaidMeterWindowSeparatorMiddle2
- TastyButtonsButtonSelectWindowTabsSeparatorBetween1
- TastyButtonsOptionsWindowTabsSeparatorBetween1
- TastyButtonsOptionsWindowTabsSeparatorBetween2
- TastyButtonsOptionsWindowTabsSeparatorBetween3
- TastyButtonsOptionsWindowTabsSeparatorBetween4

## Related APIs

- [HorizontalResizeImage](../../xml/element_types/element_HorizontalResizeImage.md) (HIGH 100/100) - XML Element Type

## Notes

- none
