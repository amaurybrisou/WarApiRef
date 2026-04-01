# EA_Window_DefaultButtonBottomFrame

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
| Addons seen in | Aura, AutoBand, EA_UiDebugTools, EA_UiModWindow, Enemy, MegaphonePlusPlus, PotionBar, Queue Queuer |
| Files seen in | `/workspace/Aura/Source/AuraSettings.xml:93`, `/workspace/Aura/Source/AuraShares.xml:169`, `/workspace/Aura/Source/AuraShares.xml:35`, `/workspace/Autoband/AutoBandWindowConfig.xml:360`, `/workspace/Autoband/AutoBandWindowTemplate.xml:210`, `/workspace/Autoband/AutoBandWindowTemplate.xml:93`, `/workspace/Autoband/AutoBandWindowTools.xml:146`, `/workspace/Enemy/Code/Core/Common.xml:140` |
| Namespaces detected | EA_Window_DefaultButtonBottomFrame |
| Source kinds | xml_attributes |
| Example locations | AuraSettingsButtonBackground, AuraSharesButtonBackground, AuraSharesImportExportButtonBackground, AutoBandWindowConfigButtonBar, AutoBandWindowTemplateButtonBar, AutoBandWindowTemplateSaveButtonBar |
| XML usage count | 29 |
| XML attribute usage count | 29 |
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

Observed engine XML template or inherited constant referenced by 10 addons.

## Seen In

- Aura
- AutoBand
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- MegaphonePlusPlus
- PotionBar
- Queue Queuer
- TidyChat
- wbLeadHelper

## Used By

- AuraSettingsButtonBackground
- AuraSharesButtonBackground
- AuraSharesImportExportButtonBackground
- AutoBandWindowConfigButtonBar
- AutoBandWindowTemplateButtonBar
- AutoBandWindowTemplateSaveButtonBar
- AutoBandWindowToolsButtonBar
- DebugWindowButtonBackground
- EnemyChooseChannelDialogButtonBackground
- EnemyClickCastingDialogButtonBackground
- EnemyConfigDialogButtonBackground
- EnemyEffectFilterDialogButtonBackground
- EnemyEffectsIndicatorDialogButtonBackground
- EnemyIntercomDialogButtonBackground
- EnemyIntercomJoinDialogButtonBackground
- EnemyMarkConfigDialogButtonBackground
- EnemyTextEntryDialogButtonBackground
- EnemyUnitFramePartDialogButtonBackground
- MegaphoneMainButtonBackground
- PotionBarAboutButtonBackground
- PotionBarTypeTemplateButtonBackground
- QueueQueuer_GUI_ButtonBackground
- TidyChatCopyButtonBackground
- UiModAdvancedWindowButtonBackground
- UiModVersionMismatchWindowButtonBackground
- UiModWindowButtonBackground
- WbLeadHelperMessageButtonBackground
- wbLeadHelperConfigTabButtonBar
- wbLeadHelperMessagesTabButtonBar

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
