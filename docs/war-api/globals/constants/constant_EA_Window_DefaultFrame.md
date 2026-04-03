# EA_Window_DefaultFrame

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, ArmorGraphicToggle, Atlas, AuctionStats, AutoSalvage, Busted, CDown |
| Files seen in | AGTSettingsWindow.xml, APAGui.xml, AdvancedRenownTrainingImportDialog.xml, AdvancedRenownTrainingPresets.xml, AdvancedRenownTrainingWindow.xml, Busted.xml, CDownSettings.xml, ClosetGoblin.xml |
| Namespaces detected | EA_Window_DefaultFrame |
| Source kinds | xml_attributes |
| Example locations | AGTSettingsWindowFrame, APAOptionsBackground, AdvancedRenownTrainingExportWindowBackground, AdvancedRenownTrainingImportNameInputWindowBackground, AdvancedRenownTrainingImportWindowBackground, AdvancedRenownTrainingLinkWindowBackground |
| XML usage count | 134 |
| XML attribute usage count | 134 |
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

Engine-supplied XML constant or template class referenced by 59 addons.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- ArmorGraphicToggle
- Atlas
- AuctionStats
- AutoSalvage
- Busted
- CDown
- CM_ClosetGoblin
- CaVES
- CraftingWillard
- DAoCBuff
- DPSMeter
- DammazKron
- DeepSleep
- DetauntHelper
- Dye Preview
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EA_UiModWindow
- Emojii
- Enemy
- EveryBodyGuard
- GroupSpotter
- GuildWarden
- HealGrid
- Killer
- Kwestor
- LoyalPet
- MapPin
- Miracle Grow Remix
- NerfedButtons
- Phantom
- Pocket Palette
- PotionBar
- QuickWarReport
- RandomMount
- ReliquaryHunter
- RememberIt
- RoR_SoR
- SessionRPs
- TacticSetNames
- TidyChat
- TidyQueue
- TokenMachine
- Tortall_DPS
- Twister
- Vectors
- WARCommander
- WSCT
- WTes
- WarBoard_Loc
- Wargames
- WhoHealedMe
- XpStatus+G
- minesweep
- nLootLink
- wbLeadHelper

## Used By

- AGTSettingsWindowFrame
- APAOptionsBackground
- AdvancedRenownTrainingExportWindowBackground
- AdvancedRenownTrainingImportNameInputWindowBackground
- AdvancedRenownTrainingImportWindowBackground
- AdvancedRenownTrainingLinkWindowBackground
- AdvancedRenownTrainingPresetsWindowBackground
- AdvancedRenownTrainingWindowBackground
- AtlasConfigurationFrameContent
- AtlasConfigurationFrameContentMapSettings
- AtlasConfigurationFrameContentTweaksContent
- AtlasFrameLegendContent
- AtlasFrameMapContainer
- AuStatsOptionsBackground
- AutoSalvageOptionsBackground
- BustedGUIBack
- CDown_SettingsBackground
- CaVESWindowOptionsWindowBackground
- ClosetGoblinCharacterWindowBackground
- ClosetGoblinCharacterWindowContentsActionBarSettingsBG
- ClosetGoblinCharacterWindowContentsSetListBackground
- ClosetGoblinZoneWindowBackground
- ClosetGoblinZoneWindowContentsSetListBackground
- CraftingWillardMainBackground
- DAoCBuffMessageWindowBackground
- DAoCBuff_SettingsBackground
- DAoCBuff_Settings_FilterBackground
- DKconfigWindowBackground
- DPSMeterWindowBackground
- DTHConfigWindowBaseBackground
- DebugWindowBackground
- DebugWindowLogDisplaySeperator2
- DebugWindowOptionsBackground
- DeepSleep_SettingsBackground
- DevPadConfirmLoadBackground
- DevPadConfirmLoadNameSeperator
- DevPadDeleteWindowBackground
- DevPadDeleteWindowNameSeperator
- DevPadNewWindowBackground
- DevPadNewWindowNameSeperator
- DevPadProjectLoadBackground
- DevPadProjectLoadNameSeperator
- DevPadRenameWindowBackground
- DevPadRenameWindowNameSeperator
- DevPadSaveWindowBackground
- DevPadSaveWindowNameSeperator
- DevPadWindowBackground
- DevPadWindowNameSeperator
- DyeBG
- EA_Window_OpenPartyBackground
- EA_Window_OpenPartyManageSocketDefConvertToWarbandBackground
- EmojiiChooseIconDialogBackground
- EnemyChooseChannelDialogBackground
- EnemyChooseIconDialogBackground
- EnemyClickCastingDialogBackground
- EnemyCombatLogSnapshotWindowBackground
- EnemyCombatLogStatsWindowBackground
- EnemyConfigDialogBackground
- EnemyEffectFilterDialogBackground
- EnemyEffectsIndicatorDialogBackground
- EnemyIntercomDialogBackground
- EnemyIntercomJoinDialogBackground
- EnemyKillSpamAreaStatsDialogBackground
- EnemyMarkConfigDialogBackground
- EnemyTextEntryDialogBackground
- EnemyUnitFramePartDialogBackground
- EveryBodyGuard_SettingsBackground
- GroupSpotterSettingsWindowBackground
- GuildWarden.DBWindowBG
- HGG_HealGridBackground
- HGG_HealGridGuiColorSelectBackground
- KillerScoreDetailsWindowBackground
- KillerSettingsWindowBackground
- KillerWindowBackground
- KwestorGui_KwestorBackground
- LPETOptionsBackground
- MapPinChooseIconDialogBackground
- MineSweepWindowBackground
- MiracleGrow2Background
- NBSBCoreWindowBackground
- NewHighScoreBG
- ObjectInspectorNameSeperator
- PPMainBG
- PhantomSettingsFrame
- PotionBarAboutBackground
- PotionBarButtonsBackground
- PotionBarTypeTemplateBackground
- QuickWarReportConfirmBackground
- RandomMountWindowBackground
- ReliquaryHunterOptionsWindowBackground
- RememberItSettingsFrame
- RoR_SoR_OffsetBackground
- RoR_SoR_OpacityBackground
- RoR_SoR_ScaleBackground
- ScenarioGroupSetOpacityWindowBackground
- SiegeChatFiltersBackground
- TacticSetNamesColorPickerBackground
- TacticSetNamesSettingsWindowBackground
- TidyChatCopyBackground
- TidyQueueBackground
- TokenMachineSettingsFrame
- TortallDPSDetailTemplateBG
- TortallDPSGroupTemplateBG
- TortallDPSMeterAbilitiesBG
- TortallDPSMeterBG
- TortallDPSMeterDetailWindowBG
- TortallDPSMeterGroupBG
- TortallDPSMeterGroupWindowBG
- TortallDPSMeterResetBG
- TortallDPSMeterScaleBG
- TortallDPSMeterTargetsBG
- TortallDPSMeterWhatHMBG
- TortallDPSMeterWhoHMBG
- TortallDPSToggleBG
- TwisterSettingsFrame
- UiModAdvancedWindowCustomUIBackground
- UiModAdvancedWindowModsBackground
- Vectors_ImportBackground
- Vectors_SettingsBackground
- WARCommanderBackground
- WSCTOptionsBackground
- WSCTOptionsColorPickerWindowBackground
- WarBoard_Loc_SetWidthWindowBackground
- WargamesGemsBackground
- WargamesPairsBackground
- WargamesTTTBackground
- WbLeadHelperMessageBackground
- WhoHealedMeDetailsBackground
- WhoHealedMeOptionsBackground
- WhoHealedMeWindowBackground
- XpStatusSetOpacityWindowBackground
- nLootLinkGUIBackground
- nLootLinkOptionsBackground
- wbLeadHelperChooseIconDialogBackground

## Related APIs

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
