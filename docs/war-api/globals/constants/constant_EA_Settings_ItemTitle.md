# EA_Settings_ItemTitle

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
| Addons seen in | CaVES, EZCraftX, PotionBar, RVAPI_Range, RVMOD_Manager, RVMOD_SquaredDistances, ReliquaryHunter, WSCT |
| Files seen in | RVAPI_Range.xml, RVMOD_ManagerTemplates.xml, RVMOD_SquaredDistances.xml, Source/EZCraftX.xml, Source/EZCraftX_Templates.xml, WarBoardOptions.xml, WarBoard_Loc.xml, WarBoard_Session.xml |
| Namespaces detected | EA_Settings_ItemTitle |
| Source kinds | xml_attributes |
| Example locations | CaVESWindowOptionsWindowRefStatsOptionsComboLabel, EZCraftXWindow.CraftingResult.ResultCrit.label, EZCraftXWindow.CraftingResult.ResultHeader.label, EZCraftXWindow.CraftingResult.ResultLevel.label, EZCraftXWindow.CraftingResult.ResultStats.label, EZCraftXWindow.CraftingResult.ResultType.label |
| XML usage count | 54 |
| XML attribute usage count | 54 |
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

Engine-supplied XML constant or template class referenced by 14 addons.

## Seen In

- CaVES
- EZCraftX
- PotionBar
- RVAPI_Range
- RVMOD_Manager
- RVMOD_SquaredDistances
- ReliquaryHunter
- WSCT
- WarBoard
- WarBoard_Loc
- WarBoard_Session
- XpStatus+G
- alertMod
- zMailMod

## Used By

- CaVESWindowOptionsWindowRefStatsOptionsComboLabel
- EZCraftXWindow.CraftingResult.ResultCrit.label
- EZCraftXWindow.CraftingResult.ResultHeader.label
- EZCraftXWindow.CraftingResult.ResultLevel.label
- EZCraftXWindow.CraftingResult.ResultStats.label
- EZCraftXWindow.CraftingResult.ResultType.label
- EZCraftXWindow.CraftingResultButtonAwesome.label
- EZCraftXWindow.CraftingResultButtonGood.label
- EZCraftXWindow.CraftingResultButtonRegular.label
- EZCraftXWindow.DropDown.label
- EZCraftXWindow.Text3Chars.label
- EZCraftXWindow.useCraftingResult.label
- EZCraftXWindow.usePowerPreview.label
- EZCraftX_Templates.PowerPreview.label
- PotionBarTypeTemplateOpacitySliderMaxLabel
- PotionBarTypeTemplateOpacitySliderMidLabel
- PotionBarTypeTemplateOpacitySliderMinLabel
- PotionBarTypeTemplateScaleSliderMaxLabel
- PotionBarTypeTemplateScaleSliderMidLabel
- PotionBarTypeTemplateScaleSliderMinLabel
- RVAPI_RangeSliderTemplateMaxLabel
- RVAPI_RangeSliderTemplateMidLabel
- RVAPI_RangeSliderTemplateMinLabel
- RVMOD_ManagerSliderTemplateMaxLabel
- RVMOD_ManagerSliderTemplateMidLabel
- RVMOD_ManagerSliderTemplateMinLabel
- RVMOD_SquaredDistancesSliderTemplateMaxLabel
- RVMOD_SquaredDistancesSliderTemplateMidLabel
- RVMOD_SquaredDistancesSliderTemplateMinLabel
- ReliquaryHunterOptionsWindowEnableWorldMapWindowBackgroundOpacitySliderMaxLabel
- ReliquaryHunterOptionsWindowEnableWorldMapWindowBackgroundOpacitySliderMidLabel
- ReliquaryHunterOptionsWindowEnableWorldMapWindowBackgroundOpacitySliderMinLabel
- ReliquaryHunterOptionsWindowEnableWorldMapWindowOpacitySliderMaxLabel
- ReliquaryHunterOptionsWindowEnableWorldMapWindowOpacitySliderMidLabel
- ReliquaryHunterOptionsWindowEnableWorldMapWindowOpacitySliderMinLabel
- ReliquaryHunterOptionsWindowWorldMapWindowBackgroundOpacityLabel
- ReliquaryHunterOptionsWindowWorldMapWindowOpacityLabel
- WSCTSliderTemplateMaxLabel
- WSCTSliderTemplateMinLabel
- WarBoardHorizontalPaddingName
- WarBoardHorizontalPaddingValue
- WarBoardVerticalPaddingName
- WarBoardVerticalPaddingValue
- WarBoard_Loc_lblHidden
- WarBoard_Loc_lblLocation
- WarBoard_Loc_lblLocation3
- WarBoard_SessionText
- XpStatusSetOpacityWindowSliderMaxLabel
- XpStatusSetOpacityWindowSliderMidLabel
- XpStatusSetOpacityWindowSliderMinLabel
- alertModTemplateSliderMaxLabel
- alertModTemplateSliderMidLabel
- alertModTemplateSliderMinLabel
- zMailModOptionsSubheader

## Related APIs

- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type

## Notes

- none
