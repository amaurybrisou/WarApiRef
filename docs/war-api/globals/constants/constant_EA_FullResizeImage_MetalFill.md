# EA_FullResizeImage_MetalFill

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
| Addons seen in | AutoBand, Effigy, LoyalPet, PeaceOut, Queue Queuer, TidyChat, WSCT, WarBoard |
| Files seen in | `/workspace/Autoband/AutoBandWindow.xml:46`, `/workspace/Effigy/Effigy.xml:95`, `/workspace/LoyalPet/gui/lpet_gui.xml:171`, `/workspace/PeaceOut/PeaceOut.xml:37`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:203`, `/workspace/TidyChat/TidyChatLootRoll.xml:76`, `/workspace/WarBoard/WarBoardOptions.xml:154`, `/workspace/wbLeadHelper/config/wbLeadHelperConfigWindow.xml:36` |
| Namespaces detected | EA_FullResizeImage_MetalFill |
| Source kinds | xml_attributes |
| Example locations | AutoBandWindowTabsBackground, FrameMetalBack, LPETOptionsTabBackground, PeaceOutBg, QueueQueuer_GUI_TabsBackground, TidyChatLootRollInfoBackground |
| XML usage count | 9 |
| XML attribute usage count | 9 |
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

Observed engine XML template or inherited constant referenced by 9 addons.

## Seen In

- AutoBand
- Effigy
- LoyalPet
- PeaceOut
- Queue Queuer
- TidyChat
- WSCT
- WarBoard
- wbLeadHelper

## Used By

- AutoBandWindowTabsBackground
- FrameMetalBack
- LPETOptionsTabBackground
- PeaceOutBg
- QueueQueuer_GUI_TabsBackground
- TidyChatLootRollInfoBackground
- WSCTOptionsTabBackground
- WarBoardOptionsTabsBackground
- wbLeadHelperConfigWindowTabsBackground

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
