# EA_Window_DefaultBackgroundFrame

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
| Addons seen in | AggroMeter, Atlas, Aura, AutoBand, BankArkel, BlackBook, Bloody Mess, BuffHead |
| Files seen in | AggroMeter.xml, AutoBandWindow.xml, AutoBandWindowTemplate.xml, BankArkel.xml, BlackBook.xml, Bloody Mess.xml, Calling.xml, CalljoinGUI.xml |
| Namespaces detected | EA_Window_DefaultBackgroundFrame |
| Source kinds | xml_attributes |
| Example locations | $parentBackground, AggroMeterGrayWindowBackground, AtlasConfigurationFrameBackground, AtlasFrameBackground, AuctionWindowBackground, AuraColorPickerBackground |
| XML usage count | 160 |
| XML attribute usage count | 160 |
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

Engine-supplied XML constant or template class referenced by 79 addons.

## Seen In

- AggroMeter
- Atlas
- Aura
- AutoBand
- BankArkel
- BlackBook
- Bloody Mess
- BuffHead
- Calling
- CastSequence
- Crusher
- DaemonAssist
- Dascore
- Deathblow
- Deathblow2
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EA_UiModWindow
- EZCraftX
- FastFriends
- FastInteract
- FozAuction
- GDes
- Ges
- Group Icons SG
- GroupRange
- Hopper
- KeyBar
- Keyset
- KeysetMonsterPlay
- LibAddonButton
- LibGroup
- MapMonster
- MapPin
- MarkBuff
- Mass Refine
- MegaphonePlusPlus
- Motion
- NerfedButtons
- ObjectInspector
- Obsidian
- PartyAd
- PieTracker
- Pocket Palette
- Pure
- Queue Queuer
- QuickTacticSwitch
- RPs
- RVMOD_Manager
- RealmStatus
- Res
- SOR
- Sequencer
- SessionRPs
- Shinies
- SocialWindow 2.0
- Squared
- TalismanGenie
- TastyButtons
- TaxPayer
- TexturedButtons
- ThinkOutLoud
- TidyChat
- Tome Titan
- TomeTracker
- TurretRange
- WARCommander
- WTes
- WarBoard
- Warbuilder
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- XpStatus+G
- ZonePOP
- alertMod
- bigger_MacroWindow
- emotes
- wbLeadHelper
- zMailMod

## Used By

- $parentBackground
- AggroMeterGrayWindowBackground
- AtlasConfigurationFrameBackground
- AtlasFrameBackground
- AuctionWindowBackground
- AuraColorPickerBackground
- AuraConfigBackground
- AuraSettingsBackground
- AuraSharesBackground
- AuraSharesImportExportBackground
- AuraTextureBackground
- AutoBandCopyLinkWindowBackground
- AutoBandWindowBackground
- AutoBandWindowTemplateDeleteConfirmBackground
- AutoBandWindowTemplateSaveBackground
- BankArkelBackpackComboFrame
- BlackBookWindowBackground
- BloodyMessOptions_Background
- BuffHeadSetupAdvancedCompressionItemEffectWindowBackground
- BuffHeadSetupAdvancedCompressionItemWindowBackground
- BuffHeadSetupAdvancedCompressionWindowBackground
- BuffHeadSetupAdvancedContainersItemPropertiesWindowBackground
- BuffHeadSetupAdvancedContainersItemWindowBackground
- BuffHeadSetupAdvancedContainersWindowBackground
- BuffHeadSetupContainerWindowBackground
- BuffHeadSetupDisplayWindowBackground
- BuffHeadSetupEffectCacheTableWindowBackground
- BuffHeadSetupEffectCacheWindowBackground
- BuffHeadSetupFilterWindowBackground
- BuffHeadSetupGeneralWindowBackground
- BuffHeadSetupLayoutManagerWindowBackground
- BuffHeadSetupLayoutPropertiesWindowBackground
- BuffHeadSetupLayoutWindowBackground
- BuffHeadSetupMenuWindowBackground
- BuffHeadSetupPerformanceWindowBackground
- BuffHeadSetupPriorityEffectsItemWindowBackground
- BuffHeadSetupPriorityEffectsWindowBackground
- BuffHeadSetupTrackersWindowBackground
- CallingSetupBackground
- CalljoinGUIWindowBackground
- CastSequenceBuilderWindowBackground
- CastSequenceFindAbilityWindowBackground
- CastSequenceSequenceBuilderWindowBackground
- CastSequenceSetupWindowBackground
- CrusherConfigBackground
- DaemonAssistWindowBackground
- DascoreWin1WindowBackground
- DeathblowWin1WindowBackground
- DyeWindowBackground
- EA_Window_MacroBackground
- EZCraftXWindow.CraftingResultBackground
- EZCraftXWindowBackground
- FastFriendsConfigMainBackground
- FastInteractWindowBackground
- GDesOptionsBackground
- GesOptionsBackground
- GroupIconsSGOptionsBackground
- GroupRangeGroupBoxTemplateBackground
- GroupRangeSetupGeneralWindowBackground
- GroupRangeSetupMenuWindowBackground
- GroupRangeSetupStyleWindowBackground
- HopperConfigBackground
- ItemWindowBackground
- KeyBarHelpWindowBackground
- KeyBarSettingsWindowBackground
- KeysetLoadWindowBackground
- KeysetMonsterPlayProfileWindowBackground
- KeysetSaveWindowBackground
- LibAddonButtonManagerAdvancedWindowBackground
- LibAddonButtonManagerCustomItemWindowBackground
- LibAddonButtonManagerMacroHelpWindowBackground
- LibAddonButtonManagerWindowBackground
- LibGroupSetupWindowBackground
- MacroIconSelectionWindowBackground
- MapMonster_CalibrateWindowMain
- MapMonster_EditorWindowBackground
- MapMonster_IconChooserWindowBackground
- MapMonster_PinTypeEditorWindowBackground
- MapPin_SetupBackground
- MarkBuffBarBackground
- MarkBuffBarTargetBackground
- MarkBuffSetupSmartBuffWindowBackground
- MarkBuffSetupWindowBackground
- MassRefineWindowBackground
- MegaphoneMainBackground
- MotionConfigBackground
- NBSetup_SaveBackground
- ObjectInspectorBackground
- ObsidianSetupCastbarWindowBackground
- ObsidianSetupEffectTrackerWindowBackground
- ObsidianSetupMenuWindowBackground
- PackWinBackground
- PartyAdWindowBackground
- PieTrackerBackground
- ProfileWindow
- PureConfigBackground
- QueueQueuer_GUI_Background
- QuickTacticSwitchWindowBackground
- RPsBackground
- RVMOD_ManagerWindowBackground
- RealmStatusHistoryWindowBackground
- ResOptionsBackground
- SOR.Options.Background
- SOR.Version.Background
- ScenarioGroupWindowBackground
- Sequencer_WindowBackground
- SessionRPsBackground
- ShiniesBrowseUI_SearchesBackground
- ShiniesWindowBackground
- SocialWindowAddMemberWindowBackground
- SocialWindowBackground
- SocialWindowBuddyListBackground
- SquaredImportExportBackground
- SummaryWindow
- TOLSettingsMainWindowBackground
- TTitanUIBackground
- TalismanGenieBackground
- TastyButtonsButtonSelectWindowBackground
- TastyButtonsGeneralOptionsProfileWindowBackground
- TastyButtonsGeneralOptionsWindowBackground
- TastyButtonsOptionsButtonInfoWindowBackground
- TastyButtonsOptionsWindowBackground
- TaxPayerBackground
- TexturedButtonsSetupActionbarWindowBackground
- TexturedButtonsSetupAdvancedTexturesWindowBackground
- TexturedButtonsSetupCooldownWindowBackground
- TexturedButtonsSetupFontsWindowBackground
- TexturedButtonsSetupMenuWindowBackground
- TexturedButtonsSetupMiscWindowBackground
- TexturedButtonsSetupTexturesWindowBackground
- TexturedButtonsSetupTintWindowBackground
- TidyChatLootRollBackground
- TomeTracker_JournalWindowBackground
- TurretRangeBoxBackground
- TurretRangeSetupDisplayWindowBackground
- TurretRangeSetupDistanceWindowBackground
- TurretRangeSetupDistancesWindowBackground
- TurretRangeSetupGeneralWindowBackground
- TurretRangeSetupMenuWindowBackground
- UiModAdvancedWindowBackground
- UiModVersionMismatchWindowBackground
- UiModWindowBackground
- UngroupedPlayersWindowBackground
- WARCommander_ConfigWindowBackground
- WCDBConfigBackground
- WCDPConfigBackground
- WTesBackground
- WarBoardAlignOptionsWindowBackground
- WarBoardOptionsBackground
- WarbuilderHotBarWindowBackground
- WarbuilderMainWindowBackground
- WarbuilderPresetWindowBackground
- ZonePOPEXTRABackground
- ZonePOPWndBackground
- alertModBackground
- emotesDisplayWindowBackground
- emotesFilterWindowBackground
- wbLeadHelperConfigWindowBackground
- zMailModLogBackground
- zMailModOptionsBackground

## Related APIs

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
