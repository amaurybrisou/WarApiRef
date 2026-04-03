# EA_Button_Default

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 170

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AggroMeter, BuffHead, CM_ClosetGoblin, MiracleGrowLight, PartyCast, Pocket Palette, PotionBar, RoR_SoR |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayout.xml:0`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.xml:0`, `/workspace/data/raw/MiracleGrowLight/window.xml:0`, `/workspace/data/raw/PartyCast/PartyCast.xml:0`, `/workspace/data/raw/PocketPalette/PocketPalette.xml:0`, `/workspace/data/raw/PotionBar/source/Floating.xml:0`, `/workspace/data/raw/RoR_SoR/RoR_SoR.xml:0` |
| Namespaces detected | EA_Button_Default |
| Source kinds | flows, xml_attributes |
| Example locations | Aggro_Button_Template, BuffHeadLayoutResizeButton, ClosetGoblinDefaultButton, ItemWindowSlotButton, MiracleGrowLightButton, PartyCastWindow_TemplateButton |
| XML usage count | 16 |
| XML attribute usage count | 16 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 1 |
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

- AggroMeter
- BuffHead
- CM_ClosetGoblin
- MiracleGrowLight
- PartyCast
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies

## Used By

- Aggro_Button_Template
- BuffHeadLayoutResizeButton
- ClosetGoblinDefaultButton
- ItemWindowSlotButton
- MiracleGrowLightButton
- PartyCastWindow_TemplateButton
- PartyCastWindow_Template_LargeButton
- PartyCastWindow_Template_PlainButton
- PartyCastWindow_Template_SmallButton
- PotionBarButtonTemplate
- RoR_SoR_RealmTemplateCLAIM_WINDOW1BUTTON
- RoR_SoR_RealmTemplateCLAIM_WINDOW1BUTTON2
- RoR_SoR_RealmTemplateCLAIM_WINDOW2BUTTON
- RoR_SoR_RealmTemplateCLAIM_WINDOW2BUTTON2
- Shinies_IconButton
- Shinies_IconButton_Overlay

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
