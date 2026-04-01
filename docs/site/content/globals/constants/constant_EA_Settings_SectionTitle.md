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
| Addons seen in | Aura, DAoCBuff, EA_UiDebugTools, PotionBar, Queue Queuer |
| Files seen in | `/workspace/Aura/Source/AuraConfig.xml:555`, `/workspace/DAoCBuff/Source/DAoCBuffMsgWindow.xml:101`, `/workspace/PotionBar/settings/Settings.xml:630`, `/workspace/QueueQueuer/QueueQueuer_GUI_TabTier1.xml:53`, `/workspace/QueueQueuer/QueueQueuer_GUI_TabTier2.xml:53`, `/workspace/QueueQueuer/QueueQueuer_GUI_TabTier3.xml:53`, `/workspace/QueueQueuer/QueueQueuer_GUI_TabTier4.xml:145`, `/workspace/QueueQueuer/QueueQueuer_GUI_TabTier4.xml:53` |
| Namespaces detected | EA_Settings_SectionTitle |
| Source kinds | xml_attributes |
| Example locations | AuraConfigTriggerAbilityDragAndDropNote, DAoCBuffMessageWindowScrollWindowScrollChildTitleLabel, DebugWindowOptionsErrorHandlingTitle, DebugWindowOptionsFiltersTitle, PotionBarAboutVersionInfo, QueueQueuer_GUI_TabTier1Socket_BlacklistTier1Title |
| XML usage count | 10 |
| XML attribute usage count | 10 |
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

Observed engine XML template or inherited constant referenced by 5 addons.

## Seen In

- Aura
- DAoCBuff
- EA_UiDebugTools
- PotionBar
- Queue Queuer

## Used By

- AuraConfigTriggerAbilityDragAndDropNote
- DAoCBuffMessageWindowScrollWindowScrollChildTitleLabel
- DebugWindowOptionsErrorHandlingTitle
- DebugWindowOptionsFiltersTitle
- PotionBarAboutVersionInfo
- QueueQueuer_GUI_TabTier1Socket_BlacklistTier1Title
- QueueQueuer_GUI_TabTier2Socket_BlacklistTier2Title
- QueueQueuer_GUI_TabTier3Socket_BlacklistTier3Title
- QueueQueuer_GUI_TabTier4Socket_BlacklistTier4SpecialTitle
- QueueQueuer_GUI_TabTier4Socket_BlacklistTier4Title

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
