# EA_Settings_SectionTitle

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
| Addons seen in | Aura, CDown, CaVES, DAoCBuff, EA_UiDebugTools, PotionBar, Queue Queuer, ReliquaryHunter |
| Files seen in | CDownSettingsTabs.xml, QueueQueuer_GUI_TabTier1.xml, QueueQueuer_GUI_TabTier2.xml, QueueQueuer_GUI_TabTier3.xml, QueueQueuer_GUI_TabTier4.xml, Source/AuraConfig.xml, Source/DAoCBuffMsgWindow.xml, Source/DAoCBuffSettingsTabs.xml |
| Namespaces detected | EA_Settings_SectionTitle |
| Source kinds | xml_attributes |
| Example locations | AuraConfigTriggerAbilityDragAndDropNote, CDownColorSettingsTab_ScrollChild_ColorSettingsLabel, CDownGeneralSettingsTab_ScrollChild_GeneralSettingsLabel, CDownNLayoutSettingsTab_ScrollChild_NLayoutSettingsLabel, CDownSLayoutSettingsTab_ScrollChild_SLayoutSettingsLabel, CaVESWindowOptionsWindowGeneralTitle |
| XML usage count | 32 |
| XML attribute usage count | 32 |
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

Engine-supplied XML constant or template class referenced by 9 addons.

## Seen In

- Aura
- CDown
- CaVES
- DAoCBuff
- EA_UiDebugTools
- PotionBar
- Queue Queuer
- ReliquaryHunter
- zMailMod

## Used By

- AuraConfigTriggerAbilityDragAndDropNote
- CDownColorSettingsTab_ScrollChild_ColorSettingsLabel
- CDownGeneralSettingsTab_ScrollChild_GeneralSettingsLabel
- CDownNLayoutSettingsTab_ScrollChild_NLayoutSettingsLabel
- CDownSLayoutSettingsTab_ScrollChild_SLayoutSettingsLabel
- CaVESWindowOptionsWindowGeneralTitle
- CaVESWindowOptionsWindowReferenceStatsTitle
- CaVESWindowOptionsWindowToolTipTitle
- DAoCBuffFrameSettingsTab_ScrollChild_FilterLabel
- DAoCBuffFrameSettingsTab_ScrollChild_FrameSettingsLabel
- DAoCBuffFrameSettingsTab_ScrollChild_LayoutLabel
- DAoCBuffFrameSettingsTab_ScrollChild_SortLabel
- DAoCBuffGeneralSettingsTab_ScrollChild_GeneralSettingsLabel
- DAoCBuffGeneralSettingsTab_ScrollChild_GlobalFrameSettingsLabel
- DAoCBuffListManagerTab_ScrollChild_ListSettingsLabel
- DAoCBuffMessageWindowScrollWindowScrollChildTitleLabel
- DAoCBuff_Settings_FilterFrame_ScrollChild_ActionLabel
- DAoCBuff_Settings_FilterFrame_ScrollChild_ConditionLabel
- DAoCBuff_Settings_FilterFrame_ScrollChild_EnableFilterCheckBoxText
- DAoCBuff_Settings_FilterFrame_ScrollChild_EnvironmentLabel
- DAoCBuff_Settings_FilterFrame_ScrollChild_MiscLabel
- DebugWindowOptionsErrorHandlingTitle
- DebugWindowOptionsFiltersTitle
- PotionBarAboutVersionInfo
- QueueQueuer_GUI_TabTier1Socket_BlacklistTier1Title
- QueueQueuer_GUI_TabTier2Socket_BlacklistTier2Title
- QueueQueuer_GUI_TabTier3Socket_BlacklistTier3Title
- QueueQueuer_GUI_TabTier4Socket_BlacklistTier4SpecialTitle
- QueueQueuer_GUI_TabTier4Socket_BlacklistTier4Title
- ReliquaryHunterOptionsWindowMarkerTitle
- ReliquaryHunterOptionsWindowWorldMapWindowTitle
- zMailModOptionsHeader

## Related APIs

- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type

## Notes

- none
