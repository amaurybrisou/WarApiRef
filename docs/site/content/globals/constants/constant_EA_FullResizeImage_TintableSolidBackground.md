# EA_FullResizeImage_TintableSolidBackground

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
| Addons seen in | InfoScroller, Moth, PartyCast, Soloq, TidyChat, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.xml:236`, `/workspace/data/raw/Moth/Moth.xml:109`, `/workspace/data/raw/Moth/Moth.xml:128`, `/workspace/data/raw/Moth/Moth.xml:157`, `/workspace/data/raw/Moth/Moth.xml:18`, `/workspace/data/raw/Moth/Moth.xml:30`, `/workspace/data/raw/PartyCast/PartyCast.xml:305`, `/workspace/data/raw/PartyCast/PartyCast.xml:348` |
| Namespaces detected | EA_FullResizeImage_TintableSolidBackground |
| Source kinds | xml_attributes |
| Example locations | AutoRollRowTemplateBackground, InfoScrollerTemplateBackGroundBG, MothBackground, MothBordertronned, MothCellTemplateBackground, MothHealthBar |
| XML usage count | 13 |
| XML attribute usage count | 13 |
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

Observed engine XML template or inherited constant referenced by 7 addons.

## Seen In

- InfoScroller
- Moth
- PartyCast
- Soloq
- TidyChat
- TidyRoll
- minesweep

## Used By

- AutoRollRowTemplateBackground
- InfoScrollerTemplateBackGroundBG
- MothBackground
- MothBordertronned
- MothCellTemplateBackground
- MothHealthBar
- MothRowTemplateBackground
- PartyCastWindow_Template_LargeBG
- PartyCastWindow_Template_PlainBG
- SoloqOverviewWindowQueueStatusLeftBorder
- SoloqOverviewWindowUnderline
- TidyChatLootRollRowTemplateBackground
- minesweep_BoxBG

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
