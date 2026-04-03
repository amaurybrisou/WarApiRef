# Addon Patterns

## Deferred post-load initialization

- Category: architecture
- Summary: Observed in 3 extracted code paths
- Evidence:
  - CM_ClosetGoblin: SystemData.Events.LOADING_END -> ClosetGoblin.Initialize
  - CM_ClosetGoblin: SystemData.Events.LOADING_END -> ClosetGoblin.LoadingEnd
  - CM_ClosetGoblin: SystemData.Events.RELOAD_INTERFACE -> ClosetGoblin.Initialize

## Manifest lifecycle bootstrap

- Category: architecture
- Summary: Observed in 2 extracted code paths
- Evidence:
  - CM_ClosetGoblin: CG_ItemRack.Initialize, ClosetGoblin.OnInitialize
  - Clock: Clock.Initialize, Clock.OnUpdate

## Event-driven controllers

- Category: events
- Summary: Observed in 2 extracted code paths
- Evidence:
  - CM_ClosetGoblin: ClosetGoblin.OnInitialize
  - CM_ClosetGoblin: ClosetGoblin.ZoneChangeInit

## Saved-variable settings roots

- Category: state
- Summary: Observed in 65 extracted code paths
- Evidence:
  - CM_ClosetGoblin: ClosetGoblin
  - CM_ClosetGoblin: ClosetGoblin.ALERT_LOG_COLOR
  - CM_ClosetGoblin: ClosetGoblin.ALERT_LOG_FILTER
  - CM_ClosetGoblin: ClosetGoblin.AddNewSet
  - CM_ClosetGoblin: ClosetGoblin.CHAT_LOG_COLOR
  - CM_ClosetGoblin: ClosetGoblin.CHAT_LOG_FILTER
  - CM_ClosetGoblin: ClosetGoblin.CheckName
  - CM_ClosetGoblin: ClosetGoblin.CreateItemTooltipOLD
  - CM_ClosetGoblin: ClosetGoblin.CreateSetData
  - CM_ClosetGoblin: ClosetGoblin.CreateSetData.tactics
  - CM_ClosetGoblin: ClosetGoblin.FindBestMatchingItem
  - CM_ClosetGoblin: ClosetGoblin.FindCharacterProfile
  - CM_ClosetGoblin: ClosetGoblin.FindItem
  - CM_ClosetGoblin: ClosetGoblin.FindPlaceInBackpack
  - CM_ClosetGoblin: ClosetGoblin.FindServerProfile
  - CM_ClosetGoblin: ClosetGoblin.GetCurrentEquipmentStateToken
  - CM_ClosetGoblin: ClosetGoblin.GetEquipmentStateTokenFromList
  - CM_ClosetGoblin: ClosetGoblin.GetProfileSets
  - CM_ClosetGoblin: ClosetGoblin.GetSet
  - CM_ClosetGoblin: ClosetGoblin.GetSetByName
  - CM_ClosetGoblin: ClosetGoblin.GetSetByName.tactics
  - CM_ClosetGoblin: ClosetGoblin.NAME_EMPTY
  - CM_ClosetGoblin: ClosetGoblin.NAME_NOT_UNIQUE
  - CM_ClosetGoblin: ClosetGoblin.NAME_OK
  - CM_ClosetGoblin: ClosetGoblin.RenameSet
  - CM_ClosetGoblin: ClosetGoblin.SORT_ORDER_DOWN
  - CM_ClosetGoblin: ClosetGoblin.SORT_ORDER_UP
  - CM_ClosetGoblin: ClosetGoblin.activatingSet
  - CM_ClosetGoblin: ClosetGoblin.activationLastState
  - CM_ClosetGoblin: ClosetGoblin.activationSet
  - CM_ClosetGoblin: ClosetGoblin.activationStallChecks
  - CM_ClosetGoblin: ClosetGoblin.activationWatchdogTick
  - CM_ClosetGoblin: ClosetGoblin.currentProfile
  - CM_ClosetGoblin: ClosetGoblin.firstloadingGame
  - CM_ClosetGoblin: ClosetGoblin.itemList
  - CM_ClosetGoblin: ClosetGoblin.languageData
  - CM_ClosetGoblin: ClosetGoblin.languageData.english
  - CM_ClosetGoblin: ClosetGoblin.languageData.french
  - CM_ClosetGoblin: ClosetGoblin.languageData.german
  - CM_ClosetGoblin: ClosetGoblin.languageData.italian
  - CM_ClosetGoblin: ClosetGoblin.languageData.japanese
  - CM_ClosetGoblin: ClosetGoblin.languageData.korean
  - CM_ClosetGoblin: ClosetGoblin.languageData.s_chinese
  - CM_ClosetGoblin: ClosetGoblin.languageData.spanish
  - CM_ClosetGoblin: ClosetGoblin.lastZone
  - CM_ClosetGoblin: ClosetGoblin.loadingGame
  - CM_ClosetGoblin: ClosetGoblin.matchingItem
  - CM_ClosetGoblin: ClosetGoblin.removedLastItemSlot
  - CM_ClosetGoblin: ClosetGoblin.setData
  - CM_ClosetGoblin: ClosetGoblin.setDisplayOrder
  - CM_ClosetGoblin: ClosetGoblin.setDisplayZoneOrder
  - CM_ClosetGoblin: ClosetGoblin.settings
  - CM_ClosetGoblin: ClosetGoblin.settings.checkTalismans
  - CM_ClosetGoblin: ClosetGoblin.settings.outputMessages
  - CM_ClosetGoblin: ClosetGoblin.settings.removeToFirst
  - CM_ClosetGoblin: ClosetGoblin.settings.savedSettingsVersion
  - CM_ClosetGoblin: ClosetGoblin.settings.toolTips
  - CM_ClosetGoblin: ClosetGoblin.settings.useItemRackOnCharacterWindow
  - CM_ClosetGoblin: ClosetGoblin.settings.zoneChange
  - CM_ClosetGoblin: ClosetGoblin.sortOrder
  - CM_ClosetGoblin: ClosetGoblin.sortedSetsList
  - CM_ClosetGoblin: ClosetGoblin.version
  - CM_ClosetGoblin: ClosetGoblin.zoneType
  - Clock: ClockSettings
  - Clock: ClockSettings.Format

## Inheritance-based XML composition

- Category: ui
- Summary: Observed in 134 extracted code paths
- Evidence:
  - CM_ClosetGoblin: CG_ItemRackEQShow1 inherits EA_Window_Default
  - CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot1 inherits CG_ItemRackEquipmentButton
  - CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot10 inherits CG_ItemRackEquipmentButton
  - CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot2 inherits CG_ItemRackEquipmentButton
  - CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot3 inherits CG_ItemRackEquipmentButton
  - CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot4 inherits CG_ItemRackEquipmentButton
  - CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot5 inherits CG_ItemRackEquipmentButton
  - CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot6 inherits CG_ItemRackEquipmentButton
  - CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot7 inherits CG_ItemRackEquipmentButton
  - CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot8 inherits CG_ItemRackEquipmentButton
  - CM_ClosetGoblin: CG_ItemRackEQShow1EquipmentSlot9 inherits CG_ItemRackEquipmentButton
  - CM_ClosetGoblin: CG_ItemRackEquipmentButton inherits CharacterWindowDefaultButton
  - CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorDown inherits EA_Button_DefaultDown
  - CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorUp inherits EA_Button_DefaultUp
  - CM_ClosetGoblin: ClosetGoblinCharacterWindow inherits EA_Window_Default
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowBackground inherits EA_Window_DefaultFrame
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowClose inherits EA_Button_DefaultWindowClose
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsBG inherits EA_Window_DefaultFrame
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB1 inherits EA_Button_DefaultCheckBox
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB2 inherits EA_Button_DefaultCheckBox
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB3 inherits EA_Button_DefaultCheckBox
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB4 inherits EA_Button_DefaultCheckBox
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB5 inherits EA_Button_DefaultCheckBox
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsPSAB1 inherits ClosetGoblinActionBarPageSelector
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsPSAB2 inherits ClosetGoblinActionBarPageSelector
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsPSAB3 inherits ClosetGoblinActionBarPageSelector
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsPSAB4 inherits ClosetGoblinActionBarPageSelector
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsActionBarSettingsPSAB5 inherits ClosetGoblinActionBarPageSelector
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsBG inherits ClosetGoblinDefaultBG
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsCheckboxUseItemRack inherits EA_Button_DefaultCheckBox
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsDeleteSet inherits EA_Button_DefaultResizeable
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentShowCloak inherits EA_Button_DefaultCheckBox
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentShowCloakHeraldry inherits EA_Button_DefaultCheckBox
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentShowHelm inherits EA_Button_DefaultCheckBox
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot1 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot10 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot10Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot10Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot11 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot11Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot11Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot12 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot12Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot12Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot13 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot13Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot13Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot14 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot14Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot14Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot15 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot15Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot15Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot17 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot17Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot17Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot18 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot18Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot18Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot19 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot19Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot19Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot1Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot1Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot2 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot20 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot20Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot20Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot21 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot21Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot21Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot22 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot22Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot22Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot23 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot23Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot23Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot24 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot24Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot24Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot25 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot25Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot25Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot2Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot2Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot3 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot3Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot3Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot5 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot5Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot5Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot6 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot6Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot6Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot7 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot7Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot7Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot8 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot8Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot8Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot9 inherits ClosetGoblinEquipmentButton
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot9Talisman1 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsEquipmentSlot9Talisman2 inherits ClosetGoblinTalismanLabel
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsNewSet inherits EA_Button_DefaultResizeable
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSetListBackground inherits EA_Window_DefaultFrame
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortDownArrow inherits EA_ListSortDownArrow
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortNameButton inherits EA_Button_ListSort
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortTacticsButton inherits EA_Button_ListSort
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortTacticsDownArrow inherits EA_ListSortDownArrow
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortTacticsUpArrow inherits EA_ListSortUpArrow
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSortUpArrow inherits EA_ListSortUpArrow
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsZoneConfig inherits EA_Button_DefaultResizeable
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowCornerImage inherits EA_Default_CharacterImage
  - CM_ClosetGoblin: ClosetGoblinCharacterWindowTitleBar inherits EA_TitleBar_Default
  - CM_ClosetGoblin: ClosetGoblinDefaultButton inherits EA_Button_Default
  - CM_ClosetGoblin: ClosetGoblinEquipmentButton inherits ClosetGoblinDefaultButton
  - CM_ClosetGoblin: ClosetGoblinOptionWindow inherits EA_Window_Default
  - CM_ClosetGoblin: ClosetGoblinSetRowBackgroundName inherits EA_FullResizeImage_TintableSolidBackground
  - CM_ClosetGoblin: ClosetGoblinSetRowSelectionBorder inherits DefaultWindowBackground
  - CM_ClosetGoblin: ClosetGoblinZoneRowBackgroundZone inherits EA_FullResizeImage_TintableSolidBackground
  - CM_ClosetGoblin: ClosetGoblinZoneRowSelectionBorder inherits DefaultWindowBackground
  - CM_ClosetGoblin: ClosetGoblinZoneWindow inherits EA_Window_Default
  - CM_ClosetGoblin: ClosetGoblinZoneWindowBackground inherits EA_Window_DefaultFrame
  - CM_ClosetGoblin: ClosetGoblinZoneWindowClose inherits EA_Button_DefaultWindowClose
  - CM_ClosetGoblin: ClosetGoblinZoneWindowContentsBG inherits ClosetGoblinDefaultBG
  - CM_ClosetGoblin: ClosetGoblinZoneWindowContentsChangeZoneButton inherits EA_Button_DefaultResizeable
  - CM_ClosetGoblin: ClosetGoblinZoneWindowContentsSetListBackground inherits EA_Window_DefaultFrame
  - CM_ClosetGoblin: ClosetGoblinZoneWindowContentsSortDownArrow inherits EA_ListSortDownArrow
  - CM_ClosetGoblin: ClosetGoblinZoneWindowContentsSortSetButton inherits EA_Button_ListSort
  - CM_ClosetGoblin: ClosetGoblinZoneWindowContentsSortSetDownArrow inherits EA_ListSortDownArrow
  - CM_ClosetGoblin: ClosetGoblinZoneWindowContentsSortSetUpArrow inherits EA_ListSortUpArrow
  - CM_ClosetGoblin: ClosetGoblinZoneWindowContentsSortUpArrow inherits EA_ListSortUpArrow
  - CM_ClosetGoblin: ClosetGoblinZoneWindowContentsSortZoneButton inherits EA_Button_ListSort
  - CM_ClosetGoblin: ClosetGoblinZoneWindowTitleBar inherits EA_TitleBar_Default

## LayoutEditor movable overlays

- Category: ui
- Summary: Observed in 2 extracted code paths
- Evidence:
  - CM_ClosetGoblin: ClosetGoblinOptionWindow.OnInitialize
  - Clock: Clock.Initialize

## Normalized Concepts

| Concept | Kind | Members | Addons | Summary |
| --- | --- | --- | --- | --- |
| Alert | function | ClosetGoblin.Alert | CM_ClosetGoblin | Single explicit implementation |
| Build | function | ClosetGoblin.BuildItemList, ClosetGoblinCharacterWindow.BuildDisplayOrder, ClosetGoblinZoneWindow.BuildDisplayOrder | CM_ClosetGoblin | 3 implementations across 1 addons |
| Can | function | ClosetGoblin.CanEquip | CM_ClosetGoblin | Single explicit implementation |
| Check | function | ClosetGoblin.CheckBind, ClosetGoblin.CheckName, ClosetGoblin.CheckTalisman | CM_ClosetGoblin | 3 implementations across 1 addons |
| Clear | function | ClosetGoblin.ClearSlot | CM_ClosetGoblin | Single explicit implementation |
| Compare | function | ClosetGoblin.CompareZones | CM_ClosetGoblin | Single explicit implementation |
| Count | function | ClosetGoblin.CountItemInList, ClosetGoblin.CountItemInPossession | CM_ClosetGoblin | 2 implementations across 1 addons |
| Create | function | ClosetGoblin.CreateCharacterProfile, ClosetGoblin.CreateItemTooltip, ClosetGoblin.CreateServerProfile, ClosetGoblin.CreateSetData, ClosetGoblinCharacterWindow.CreateTextOnlyTooltip | CM_ClosetGoblin | 5 implementations across 1 addons |
| Destroy | function | ClosetGoblin.DestroyItemList | CM_ClosetGoblin | Single explicit implementation |
| Equipment | function | ClosetGoblinCharacterWindow.EquipmentLButtonDown, ClosetGoblinCharacterWindow.EquipmentLButtonUp, ClosetGoblinCharacterWindow.EquipmentMouseOver, ClosetGoblinCharacterWindow.EquipmentMouseOverEnd, ClosetGoblinCharacterWindow.EquipmentRButtonUp | CM_ClosetGoblin | 5 implementations across 1 addons |
| Event Handler | function | ClosetGoblin.OnActivationWatchdog, ClosetGoblin.OnShutdown, ClosetGoblin.OnSlashCommand, ClosetGoblinCharacterWindow.OnClickDeleteSetButton, ClosetGoblinCharacterWindow.OnClickNewSetButton, ClosetGoblinCharacterWindow.OnClickSetRow, ClosetGoblinCharacterWindow.OnClickSortNameButton, ClosetGoblinCharacterWindow.OnClickSortTacticsButton, ClosetGoblinCharacterWindow.OnClickZoneConfigButton, ClosetGoblinCharacterWindow.OnConfirmDeleteSet, ClosetGoblinCharacterWindow.OnHide, ClosetGoblinCharacterWindow.OnRowMenuCopyClick, ClosetGoblinCharacterWindow.OnRowMenuImportCurrentClick, ClosetGoblinCharacterWindow.OnRowMenuLinkTactics, ClosetGoblinCharacterWindow.OnRowMenuPasteClick, ClosetGoblinCharacterWindow.OnRowMenuRenameSetClick, ClosetGoblinCharacterWindow.OnSetRowContextMenu, ClosetGoblinCharacterWindow.OnShow, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinCharacterWindow.OnSubmitNewSetName, ClosetGoblinCharacterWindow.OnSubmitSetRename, ClosetGoblinOptionWindow.OnLButtonUp, ClosetGoblinOptionWindow.OnRButtonUp, ClosetGoblinOptionWindow.OnRowMenuActivateSet, ClosetGoblinZoneWindow.OnClickChangeZoneOptionButton, ClosetGoblinZoneWindow.OnClickZoneRow, ClosetGoblinZoneWindow.OnHide, ClosetGoblinZoneWindow.OnRowMenuAssoiacteClick, ClosetGoblinZoneWindow.OnRowMenuUnassoiacteClick, ClosetGoblinZoneWindow.OnSetZoneRowContextMenu, ClosetGoblinZoneWindow.OnShow, ClosetGoblinZoneWindow.OnShutdown | CM_ClosetGoblin | 32 implementations across 1 addons |
| Find | function | ClosetGoblin.FindAllItemInList, ClosetGoblin.FindBestMatchingItem, ClosetGoblin.FindCharacterProfile, ClosetGoblin.FindItem, ClosetGoblin.FindItemInList, ClosetGoblin.FindMatchingItemplace, ClosetGoblin.FindPlaceInBackpack, ClosetGoblin.FindServerProfile, ClosetGoblinCharacterWindow.FindCursorItem | CM_ClosetGoblin | 9 implementations across 1 addons |
| Get | function | CM_ClosetGoblin.local.GetActionBarNumberFromWindowName, ClosetGoblin.GetCurrentEquipmentStateToken, ClosetGoblin.GetEquipmentStateTokenFromList, ClosetGoblinCharacterWindow.IsValidPasteTarget, GetActionBarNumberFromWindowName | CM_ClosetGoblin | 5 implementations across 1 addons |
| Handle | function | ClosetGoblinCharacterWindow.HandleDrag | CM_ClosetGoblin | Single explicit implementation |
| Hide | function | ClosetGoblin.Hide, ClosetGoblinCharacterWindow.Hide, ClosetGoblinCharacterWindow.HideCloakOptions, ClosetGoblinZoneWindow.Hide | CM_ClosetGoblin | 4 implementations across 1 addons |
| Hotbar | function | ClosetGoblinCharacterWindow.HotbarPageDownProxy, ClosetGoblinCharacterWindow.HotbarPageUpProxy | CM_ClosetGoblin | 2 implementations across 1 addons |
| Initialize | function | Clock.Initialize, ClosetGoblin.Initialize, ClosetGoblin.OnAllModulesInitialized, ClosetGoblin.OnInitialize, ClosetGoblinCharacterWindow.OnInitialize, ClosetGoblinOptionWindow.OnInitialize, ClosetGoblinZoneWindow.OnInitialize | CM_ClosetGoblin, Clock | 7 implementations across 2 addons |
| Is | function | ClosetGoblin.IsAccessorySlot, ClosetGoblin.IsNameUnique | CM_ClosetGoblin | 2 implementations across 1 addons |
| Link | function | ClosetGoblin.LinkTactics | CM_ClosetGoblin | Single explicit implementation |
| Load | function | ClosetGoblin.LoadLanguage | CM_ClosetGoblin | Single explicit implementation |
| Loading | function | ClosetGoblin.LoadingBegin, ClosetGoblin.LoadingEnd | CM_ClosetGoblin | 2 implementations across 1 addons |
| Message | function | ClosetGoblin.Message | CM_ClosetGoblin | Single explicit implementation |
| New | function | LibToolkit.NewChatFilter | CM_ClosetGoblin | Single explicit implementation |
| Player | function | ClosetGoblin.PlayerZoneChanged | CM_ClosetGoblin | Single explicit implementation |
| Refresh | function | Clock.OnUpdate, ClosetGoblin.OnInventorySlotUpdated, ClosetGoblin.OnVisibleEquipmentUpdated, ClosetGoblin.UpdateListSets, ClosetGoblin.UpdateSettings, ClosetGoblinCharacterWindow.RefreshList, ClosetGoblinCharacterWindow.UpdateActionBarSettings, ClosetGoblinCharacterWindow.UpdateHighlightOnRow, ClosetGoblinCharacterWindow.UpdateSetList, ClosetGoblinCharacterWindow.UpdateSetRow, ClosetGoblinCharacterWindow.UpdateSlotIcon, ClosetGoblinCharacterWindow.UpdateSlotIcons, ClosetGoblinCharacterWindow.UpdateSortButtonIcon, ClosetGoblinZoneWindow.RefreshList, ClosetGoblinZoneWindow.RefreshOption, ClosetGoblinZoneWindow.UpdateAssociationTable, ClosetGoblinZoneWindow.UpdateHighlightOnRow, ClosetGoblinZoneWindow.UpdateSetList, ClosetGoblinZoneWindow.UpdateSetRow | CM_ClosetGoblin, Clock | 19 implementations across 2 addons |
| Request | function | ClosetGoblin.RequestActivationMove, ClosetGoblin.RequestMoveItem, ClosetGoblin.RequestRemoveItem | CM_ClosetGoblin | 3 implementations across 1 addons |
| Set | function | ClosetGoblin.ActivateSet, ClosetGoblin.AddNewSet, ClosetGoblin.AssociateZoneToSet, ClosetGoblin.CompareSets, ClosetGoblin.CopySet, ClosetGoblin.CountItemInSet, ClosetGoblin.DeleteSet, ClosetGoblin.FinalizeSetActivation, ClosetGoblin.FindItemInSet, ClosetGoblin.GetProfileSets, ClosetGoblin.GetSet, ClosetGoblin.GetSetByName, ClosetGoblin.GetSetNamesForItem, ClosetGoblin.HasSets, ClosetGoblin.ImportCurrentSet, ClosetGoblin.ItemIsUniqueToSet, ClosetGoblin.ProcessSetActivation, ClosetGoblin.RenameSet, ClosetGoblin.ResetActivationState, ClosetGoblin.SetSlot, ClosetGoblin.SetupProfile, ClosetGoblin.TryActivateSetStep, ClosetGoblin.ZoneChangeSet, ClosetGoblinCharacterWindow.GetCurrentSet, ClosetGoblinCharacterWindow.SortSetList, ClosetGoblinZoneWindow.SortSetList | CM_ClosetGoblin | 26 implementations across 1 addons |
| Show | function | ClosetGoblin.Show, ClosetGoblinCharacterWindow.HideShowHelm, ClosetGoblinCharacterWindow.Show, ClosetGoblinCharacterWindow.ShowCloak, ClosetGoblinCharacterWindow.ShowCloakHeraldry, ClosetGoblinCharacterWindow.ShowCloakOptions, ClosetGoblinCharacterWindow.ShowHelm, ClosetGoblinCharacterWindow.ShowOrHide, ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly, ClosetGoblinCharacterWindow.ShowShowCloakOnly, ClosetGoblinCharacterWindow.ShowShowHelm, ClosetGoblinCharacterWindow.ShowShowHelmOnly, ClosetGoblinZoneWindow.Show | CM_ClosetGoblin | 13 implementations across 1 addons |
| To | function | ClosetGoblin.ToEnabled | CM_ClosetGoblin | Single explicit implementation |
| Toggle | function | ClosetGoblin.ToggleSlot, ClosetGoblinCharacterWindow.HotbarChangeToggled1, ClosetGoblinCharacterWindow.HotbarChangeToggled2, ClosetGoblinCharacterWindow.HotbarChangeToggled3, ClosetGoblinCharacterWindow.HotbarChangeToggled4, ClosetGoblinCharacterWindow.HotbarChangeToggled5, ClosetGoblinCharacterWindow.UseItemRackToggled | CM_ClosetGoblin | 7 implementations across 1 addons |
| Trim | function | ClosetGoblin.TrimWString | CM_ClosetGoblin | Single explicit implementation |
| Try | function | ClosetGoblin.TryRecoverStuckActivation | CM_ClosetGoblin | Single explicit implementation |
| Valid | function | ClosetGoblin.ValidCareer, ClosetGoblin.ValidRace, ClosetGoblin.ValidRank, ClosetGoblin.ValidRenownRank, ClosetGoblin.ValidSkills, ClosetGoblin.ValidSlot | CM_ClosetGoblin | 6 implementations across 1 addons |
| Zone | function | ClosetGoblin.ZoneChangeChange, ClosetGoblin.ZoneChangeDisabled, ClosetGoblin.ZoneChangeEnabled, ClosetGoblin.ZoneChangeInit, ClosetGoblin.ZoneChangeShutdown | CM_ClosetGoblin | 5 implementations across 1 addons |
