# EA_Window_Default

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
| Addons seen in | AdvancedRenownTrainer, Aura, AutoBand, BlackBook, BuffHead, CDown, CM_ClosetGoblin, CMap |
| Files seen in | AAOTWarBoard.xml, AdvancedRenownTrainingImportDialog.xml, AdvancedRenownTrainingPresets.xml, AdvancedRenownTrainingWindow.xml, AutoBandWindow.xml, AutoBandWindowTemplate.xml, BlackBook.xml, CDownSettings.xml |
| Namespaces detected | EA_Window_Default |
| Source kinds | xml_attributes |
| Example locations | AAOTrackerWarBoard, AdvancedRenownTrainingExportWindow, AdvancedRenownTrainingImportNameInputWindow, AdvancedRenownTrainingImportWindow, AdvancedRenownTrainingLinkWindow, AdvancedRenownTrainingPresetsWindow |
| XML usage count | 206 |
| XML attribute usage count | 206 |
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

Engine-supplied XML constant or template class referenced by 81 addons.

## Seen In

- AdvancedRenownTrainer
- Aura
- AutoBand
- BlackBook
- BuffHead
- CDown
- CM_ClosetGoblin
- CMap
- Calling
- CastSequence
- Cheeseboard
- CraftingWillard
- Crusher
- DAoCBuff
- DPSMeter
- Dascore
- Deathblow
- Deathblow2
- DetauntHelper
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EA_UiModWindow
- EZCraftX
- Effigy
- Emojii
- Enemy
- GDes
- Ges
- GroupRange
- GroupSpotter
- ItemRack
- JunkDump
- Keyset
- KeysetMonsterPlay
- LibAddonButton
- LibGroup
- LibWBToggler
- MapMonster
- MapPin
- MarkBuff
- Minmap
- Motion
- ObjectInspector
- Obsidian
- PartyAd
- PieTracker
- Pure
- RPs
- Res
- SNT_BUTTONS
- SNT_PANEL
- SessionRPs
- Shinies
- SocialWindow 2.0
- SpamBayes
- TalismanGenie
- TaxPayer
- TexturedButtons
- TokenMachine
- Tome Titan
- TomeTracker
- TurretRange
- WARCommander
- WTes
- WaaaghBar
- WarBoard
- WarBoard_AAOTracker
- WarBoard_FPS
- WarBoard_Loc
- WarBoard_Menu
- WarBoard_Session
- WarBoard_TDPS
- WarBoard_TaliMon
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- alertMod
- bigger_MacroWindow
- emotes
- nLootLink
- wbLeadHelper
- zMailMod

## Used By

- AAOTrackerWarBoard
- AdvancedRenownTrainingExportWindow
- AdvancedRenownTrainingImportNameInputWindow
- AdvancedRenownTrainingImportWindow
- AdvancedRenownTrainingLinkWindow
- AdvancedRenownTrainingPresetsWindow
- AdvancedRenownTrainingWindow
- AuraColorPicker
- AutoBandCopyLinkWindow
- AutoBandWindow
- AutoBandWindowTemplateDeleteConfirm
- AutoBandWindowTemplateMask
- AutoBandWindowTemplateSave
- BlackBookWindow
- BuffHeadSetupAdvancedCompressionItemEffectWindow
- BuffHeadSetupAdvancedCompressionItemWindow
- BuffHeadSetupAdvancedCompressionWindow
- BuffHeadSetupAdvancedContainersItemPropertiesWindow
- BuffHeadSetupAdvancedContainersItemWindow
- BuffHeadSetupAdvancedContainersWindow
- BuffHeadSetupContainerWindow
- BuffHeadSetupDisplayWindow
- BuffHeadSetupEffectCacheTableWindow
- BuffHeadSetupEffectCacheWindow
- BuffHeadSetupFilterWindow
- BuffHeadSetupGeneralWindow
- BuffHeadSetupLayoutManagerWindow
- BuffHeadSetupLayoutPropertiesWindow
- BuffHeadSetupLayoutWindow
- BuffHeadSetupMenuWindow
- BuffHeadSetupPerformanceWindow
- BuffHeadSetupPriorityEffectsItemWindow
- BuffHeadSetupPriorityEffectsWindow
- BuffHeadSetupTrackersWindow
- CDown_Settings
- CG_ItemRackEQShow1
- CMapWindowPinFilterMenu
- CMapWindowResize
- CallingSetup
- CastSequenceBuilderWindow
- CastSequenceFindAbilityWindow
- CastSequenceSequenceBuilderWindow
- CastSequenceSetupWindow
- CheeseboardWindow
- ClosetGoblinCharacterWindow
- ClosetGoblinOptionWindow
- ClosetGoblinZoneWindow
- CraftingWillardMain
- CrusherWindowDefault
- DAoCBuff_Settings
- DAoCBuff_Settings_Filter
- DPSMeterWindow
- DTHConfigWindow
- DascoreWin1Window
- DascoreWin2Window
- DeathblowWin1Window
- EA_Window_Macro
- EZCraftXWindow
- EZCraftXWindow.CraftingResult
- EZCraftX_Templates.PowerPreview
- EmojiiChooseIconDialog
- EnemyChooseChannelDialog
- EnemyChooseIconDialog
- EnemyClickCastingDialog
- EnemyCombatLogEpsWindow
- EnemyCombatLogIDSAnchor
- EnemyCombatLogIDSRow
- EnemyCombatLogSnapshotWindow
- EnemyCombatLogStatsWindow
- EnemyCombatLogTargetDefeseTotalWindow
- EnemyCombatLogTargetDefeseWindow
- EnemyConfigDialog
- EnemyDebug
- EnemyEffectFilterDialog
- EnemyEffectsIndicatorDialog
- EnemyIntercomDialog
- EnemyIntercomJoinDialog
- EnemyKillSpamAreaStatsDialog
- EnemyKillSpamDialog
- EnemyKilledBy
- EnemyMarkConfigDialog
- EnemyPlayerKDR
- EnemyScenarioInfoDialog
- EnemyTextEntryDialog
- EnemyTimer
- EnemyUnitFramePartDialog
- EnemyUnitFramesAnchor1
- EnemyUnitFramesAnchor2
- FrameBubble
- FrameBubble2
- FrameDefault
- FrameIcon
- FrameMetal
- FrameTan
- GDesOptions
- GesOptions
- GroupRangeSetupGeneralWindow
- GroupRangeSetupMenuWindow
- GroupRangeSetupStyleGroupBoxWindow
- GroupRangeSetupStylePointerReverseWindow
- GroupRangeSetupStylePointerWindow
- GroupRangeSetupStyleSimpleTextWindow
- GroupRangeSetupStyleWindow
- GroupSpotterSettingsWindow
- ItemRackEQShow1
- JunkDumpOptionsWin
- KeysetLoadWindow
- KeysetMonsterPlayProfileWindow
- KeysetSaveWindow
- LayoutTemplate
- LibAddonButtonManagerAdvancedWindow
- LibAddonButtonManagerCustomItemWindow
- LibAddonButtonManagerMacroHelpWindow
- LibAddonButtonManagerWindow
- LibGroupSetupWindow
- MapMonster_CalibrateWindow
- MapMonster_EditorWindow
- MapMonster_IconChooserWindow
- MapMonster_PinTypeEditorWindow
- MapPinChooseIconDialog
- MarkBuffSetupSmartBuffWindow
- MarkBuffSetupWindow
- MinmapPinMenu
- MinmapScenarioMenu
- MotionWindowDefault
- NewHighScore
- ObjectInspector
- ObsidianSetupCastbarWindow
- ObsidianSetupEffectTrackerWindow
- ObsidianSetupMenuWindow
- PartyAdWindow
- PieTracker
- PureWindowDefault
- RPs
- ResOptions
- SNT_BUTTONS_FLASH
- SNT_BUTTONS_FLASH_ANCHOR
- ScenarioGroupWindow
- SessionRPs
- ShiniesAuctionsUI
- ShiniesAutoUI
- ShiniesBrowseUI_Searches
- ShiniesConfigUI
- ShiniesPostUI
- ShiniesWindow
- SocialWindowBuddyListFilterMenu
- SocialWindowTabFriendsFilterMenu
- TTitanUI
- TalismanGenie
- TalismanGenieIcon
- TaxPayer
- TexturedButtonsSetupActionbarWindow
- TexturedButtonsSetupAdvancedTexturesWindow
- TexturedButtonsSetupCooldownWindow
- TexturedButtonsSetupFontsWindow
- TexturedButtonsSetupMenuWindow
- TexturedButtonsSetupMiscWindow
- TexturedButtonsSetupTexturesWindow
- TexturedButtonsSetupTintWindow
- TokenMachineSettings
- TomeTracker_JournalWindow
- TurretRangeSetupDisplayWindow
- TurretRangeSetupDistanceWindow
- TurretRangeSetupDistancesWindow
- TurretRangeSetupGeneralWindow
- TurretRangeSetupMenuWindow
- UiModVersionMismatchWindow
- UiModWindow
- UngroupedPlayersWindow
- WARCommander_ConfigWindow
- WBTogglerTemplate
- WCDBWindowDefault
- WCDPWindowDefault
- WTes
- WaaaghBarOne
- WaaaghBarTwo
- WarBoard
- WarBoardAlignOptions
- WarBoardAlignOptionsWindow
- WarBoardBottomLayoutModeButtonWindow
- WarBoardBottomOptionsButtonWindow
- WarBoardButtonCheck
- WarBoardHorizontalPadding
- WarBoardLayoutModeButtonWindow
- WarBoardOptions
- WarBoardOptionsButtonWindow
- WarBoardPadOptionsWindow
- WarBoardVerticalPadding
- WarBoard_FPS
- WarBoard_Loc
- WarBoard_Menu
- WarBoard_Session
- WarBoard_SpamBayes
- WarBoard_TDPS
- WarBoard_TaliMon
- WbLeadHelperMessage
- alertMod
- emotesDisplayWindow
- emotesFilterWindow
- nLootLinkGUI
- nLootLinkOptions
- snt_panel_template
- wbLeadHelperChooseIconDialog
- wbLeadHelperConfigWindow
- zMailModLog
- zMailModOptions

## Related APIs

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
