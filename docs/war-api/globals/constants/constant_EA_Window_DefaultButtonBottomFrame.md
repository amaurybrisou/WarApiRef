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
| Addons seen in | Aura, AutoBand, CaVES, Crusher, Dascore, Deathblow, Deathblow2, DetauntHelper |
| Files seen in | AutoBandWindowConfig.xml, AutoBandWindowTemplate.xml, AutoBandWindowTools.xml, Code/Core/Common.xml, Code/Core/ConfigDialog.xml, Code/Core/Groups/EffectFilterDialog.xml, Code/Intercom/ChooseChannelDialog.xml, Code/Intercom/IntercomDialog.xml |
| Namespaces detected | EA_Window_DefaultButtonBottomFrame |
| Source kinds | xml_attributes |
| Example locations | AuctionWindowButtonBackground, AuraSettingsButtonBackground, AuraSharesButtonBackground, AuraSharesImportExportButtonBackground, AutoBandWindowConfigButtonBar, AutoBandWindowTemplateButtonBar |
| XML usage count | 57 |
| XML attribute usage count | 57 |
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

Engine-supplied XML constant or template class referenced by 33 addons.

## Seen In

- Aura
- AutoBand
- CaVES
- Crusher
- Dascore
- Deathblow
- Deathblow2
- DetauntHelper
- DuffTimer
- EA_OpenPartyWindow
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- FastFriends
- FozAuction
- Hopper
- MegaphonePlusPlus
- Motion
- NerfedButtons
- PotionBar
- Pure
- Queue Queuer
- ReliquaryHunter
- SOR
- Sequencer
- TidyChat
- Tome Titan
- WTes
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- alertMod
- wbLeadHelper
- zMailMod

## Used By

- AuctionWindowButtonBackground
- AuraSettingsButtonBackground
- AuraSharesButtonBackground
- AuraSharesImportExportButtonBackground
- AutoBandWindowConfigButtonBar
- AutoBandWindowTemplateButtonBar
- AutoBandWindowTemplateDeleteConfirmButtonBar
- AutoBandWindowTemplateSaveButtonBar
- AutoBandWindowToolsButtonBar
- CaVESWindowOptionsWindowButtonBackground
- CrusherConfigButtonBackground
- DTC_TARGETS_TemplateButtonBackground
- DascoreWin1WindowBarBackground
- DeathblowWin1WindowBarBackground
- DebugWindowButtonBackground
- DuffTimerOptions_ButtonBackground
- EA_Window_OpenPartyNearbySocketDefButtonBackground
- EA_Window_OpenPartyWorldSocketDefButtonBackground
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
- FastFriendsConfigMainButtonBackground
- HopperConfigButtonBackground
- MegaphoneMainButtonBackground
- MotionConfigButtonBackground
- NBSBCoreWindowCriteriaSpecialBackground
- NBSBCoreWindowListsSpecialBackground
- NBSetup_SaveButtonBackground
- PotionBarAboutButtonBackground
- PotionBarTypeTemplateButtonBackground
- PureConfigButtonBackground
- QueueQueuer_GUI_ButtonBackground
- ReliquaryHunterOptionsWindowButtonBackground
- SOR.Options.ButtonBackground
- SOR.Version.ButtonBackground
- Sequencer_WindowSpecialBackground
- SiegeChatFiltersButtonBackground
- TTitanUIButtonBackground
- TidyChatCopyButtonBackground
- UiModAdvancedWindowButtonBackground
- UiModVersionMismatchWindowButtonBackground
- UiModWindowButtonBackground
- WCDBConfigButtonBackground
- WCDPConfigButtonBackground
- WbLeadHelperMessageButtonBackground
- alertModButtonBackground
- wbLeadHelperConfigTabButtonBar
- wbLeadHelperMessagesTabButtonBar
- zMailModMassMailBarBackground
- zMailModOptionsBarBackground

## Related APIs

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
