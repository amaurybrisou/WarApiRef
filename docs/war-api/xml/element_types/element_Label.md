# Label

- Category: XML Element Type
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
| Addons seen in | AbilityAlert, AbilityNotifier, ActionPoints, AdjustTheTip, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, ArmorGraphicToggle |
| Files seen in | Bars/HealGridProgressBar.xml, Code/Assist/Assist.xml, Code/Assist/AssistConfiguration.xml, Code/CombatLog/CombatLogEpsWindow.xml, Code/CombatLog/CombatLogIDS.xml, Code/CombatLog/CombatLogSnapshotWindow.xml, Code/CombatLog/CombatLogStatsWindow.xml, Code/CombatLog/CombatLogTargetDefenseWindow.xml |
| Namespaces detected | Label |
| Source kinds | xml_frames |
| Example locations | AbilityAlert: AAText, AbilityNotifier: AbHelpText, ActionPoints: ActionPointsWindowText, AdjustTheTip: AdjustTheTipMenuItemSliderSliderLabel, AdjustTheTip: AdjustTheTipMenuItemSliderSliderValueLabel, AdjustTheTip: MouseOverTargetHealthContainerTemplateHealthText |
| XML usage count | 4835 |
| XML attribute usage count | 4835 |
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

Label is a display-focused XML element used to present text content in UI layouts. It commonly appears inside container elements such as Window.

## Common Attributes

- autoresize
- autoresizewidth
- font
- handleinput
- inherits
- layer
- maxchars
- name
- popable
- textalign
- warnOnTextCropped
- wordwrap

## Common Handlers

- OnHyperLinkLButtonUp
- OnHyperLinkRButtonUp
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseLeave](../handlers/handler_OnMouseLeave.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)

## Common Handler Functions

- LPET.OnMouseOver
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip
- Cheeseboard.OnMouseOver
- AutoBandWindowConfig.OnSettingMouseOver
- Enemy.CombatLogUI_EpsWindow_OnMouseOver
- Enemy.ConfigurationWindow_ShowTooltip
- AutoBandWindowTemplate.OnRoleLimitsMouseOver
- Enemy.CombatLogUI_StatsWindow_ListRowMouseOver
- Enemy.CombatLogUI_TargetDefenseWindow_OnMouseOver
- DuffTimer.Options.OnMouseOver
- Enemy.CombatLogUI_StatsWindow_OnEpsMouseOver
- SORRealms.OnMouseOverVP
- Enemy.KillSpamUI_KillSpamDialog_OnTotalMouseOver
- Obsidian.Setup.Castbar.OnFontExampleMouseOver
- SORPager.LowBOsOn
- SORRealms.OnMouseOverLowVP
- WARCommander.LabelTint
- AutoBandWindowConfig.OnBackfillMouseOver
- AutoBandWindowConfig.OnHealerRankReqMouseOver
- AutoBandWindowConfig.OnRankReqMouseOver
- Obsidian.Setup.EffectTracker.OnFontExampleMouseOver
- TexturedButtons.Setup.Fonts.OnFontExampleMouseOver
- Warbuilder.Effecttooltip
- AggroMeter.SelectChar
- AtlasMap.OnLegendItemMouseOver
- AutoBandWindowConfig.OnCommonRaceNamesMouseOver
- AutoBandWindowConfig.OnExcludeRealmHealerAltSpecMouseOver
- AutoBandWindowConfig.OnKickLowRankMouseOver
- AutoBandWindowTools.OnDiscordReqMouseOver
- AutoBandWindowTools.OnNoMicMouseOver
- BarText_Influence.OnMouseOver
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOver
- BuffHead.Setup.Layout.Properties.OnFontExampleMouseOver
- BuffHead.Setup.Layout.Properties.OnPropertyTitleMouseOver
- CMapWindow.OnMouseOverDigitinf
- CaVESWindow.OnMouseoverStatHeaderLabelRecent
- CaVESWindow.OnMouseoverStatHeaderLabelReference
- CaVESWindow.OnMouseoverStatHeaderLabelStat
- DeepSleep.Settings.OnMouseOver
- EA_Window_OpenPartyLootRollOptions.OnMouseOverRowTitle
- GuildWardenTTips.anni
- GuildWardenTTips.bl
- GuildWardenTTips.conq
- GuildWardenTTips.dp
- GuildWardenTTips.inv
- GuildWardenTTips.sent
- GuildWardenTTips.sigil
- GuildWardenTTips.tyr
- GuildWardenTTips.war
- KeyBarSettings.CheckBoxOnMouseOver
- LibAddonButton.Manager.CustomItem.OnMacroHelpMouseOver
- LibAddonButton.Manager.OnNavigateMouseOver
- MassMailWindow.MouseOverCraftingSkill
- MassMailWindow.MouseOverRarity
- MiracleGrowLight.onHover
- Motion.OnMouseOver_ContextMenu_Recipe_Delete
- PartyAdWindow.OnMouseOverPreviewNoteStatus
- RvRContribution.OnHoverLoss
- RvRContribution.OnHoverWin
- SORPager.HighBOsOn
- TalismanGenie.ContainerMouseOver
- TalismanGenie.CurioOn
- TalismanGenie.EssenceOn
- TalismanGenie.FragmentMouseOver
- TalismanGenie.FragmentOn
- TalismanGenie.GoldDustOn
- TalismanGenie.LevelMouseOver
- TalismanGenie.TypeMouseOver
- TurretRange.Setup.Display.OnDistanceFontMouseOver
- UiModWindow.OnMouseOverModType
- UiModWindow.OnMouseOverStatus
- WARCommander.AssistLabelTint
- emotes.OnMouseOverPathRow
- zMailModInbox.OnMouseOverCOD
- zMailModOptions.MouseOverListLabel
- Cheeseboard.SortColumnClick
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick
- SessionRPs.GetSortType
- SORRealms.OnLClickTier
- Enemy.KillSpamUI_KillSpamDialog_OnTotalLButtonUp
- AutoBandWindowConfig.OnClickBackfillLabel
- ClosetGoblinCharacterWindow.OnClickSetRow
- ClosetGoblinZoneWindow.OnClickZoneRow
- DPSMeter.SortChange
- Deathblow.LinkDeath
- Deathblow.LinkKill
- Enemy.CombatLogUI_EpsWindow_OnInDpsLButtonUp
- Enemy.CombatLogUI_EpsWindow_OnInHpsLButtonUp
- Enemy.CombatLogUI_EpsWindow_OnOutDpsAoeLButtonUp
- Enemy.CombatLogUI_EpsWindow_OnOutDpsLButtonUp
- Enemy.CombatLogUI_EpsWindow_OnOutHpsAoeLButtonUp
- Enemy.CombatLogUI_EpsWindow_OnOutHpsLButtonUp
- SORRealms.OnLClickBZ
- AutoBandWindowConfig.OnClickAltCheckLabel
- AutoBandWindowConfig.OnClickAutoKickLabel
- AutoBandWindowConfig.OnClickBWSorcAsMDPSLabel
- AutoBandWindowConfig.OnClickCommonRaceNamesLabel
- AutoBandWindowConfig.OnClickDpsWeightingLabel
- AutoBandWindowConfig.OnClickExcludeRealmHealerAltSpecLabel
- AutoBandWindowConfig.OnClickGuildPriorityLabel
- AutoBandWindowConfig.OnClickKickIgnoreLabel
- AutoBandWindowConfig.OnClickKickLowRankLabel
- AutoBandWindowConfig.OnClickKickRvRZonesLabel
- AutoBandWindowConfig.OnClickKickStealthersLabel
- AutoBandWindowConfig.OnClickKickToofarLabel
- AutoBandWindowConfig.OnClickRestrictRaceLabel
- AutoBandWindowTemplate.OnClickRightClickTemplateMenuLabel
- AutoBandWindowTools.OnClickAutoFormSearchLabel
- AutoBandWindowTools.OnClickAutoPartyNoteLabel
- AutoBandWindowTools.OnClickAutoSearchLabel
- AutoBandWindowTools.OnClickDiscordReqLabel
- AutoBandWindowTools.OnClickKickOfflineLabel
- AutoBandWindowTools.OnClickKickRankLabel
- AutoBandWindowTools.OnClickKickZoneLabel
- AutoBandWindowTools.OnClickNoMicLabel
- AutoBandWindowTools.OnClickNotifyBuffsLabel
- AutoBandWindowTools.OnClickPrintRoleLabel
- AutoBandWindowTools.OnClickRightClickOrganizeLabel
- AutoBandWindowTools.OnClickSearchRoleChan1Label
- AutoBandWindowTools.OnClickSearchRoleChan5Label
- AutoBandWindowTools.OnClickSearchRoleChanT4Label
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleLClick
- BuffHead.Setup.Layout.Manager.OnExportLUp
- BuffHead.Setup.Layout.Manager.OnImportLUp
- BuffHead.Setup.Layout.Manager.OnLayoutsLUp
- BuffHead.Setup.Layout.Properties.OnFontFontLUp
- BuffHead.Setup.Layout.Properties.OnPropertyTitleLClick
- CMapWindow.OnClickDigitinf
- Enemy.CombatLogUI_StatsWindow_OnEpsAoeCurClick
- Enemy.CombatLogUI_StatsWindow_OnEpsAoeMaxClick
- Enemy.CombatLogUI_StatsWindow_OnEpsCurClick
- Enemy.CombatLogUI_StatsWindow_OnEpsMaxClick
- FastFriends.UI.OnClickCharOverrideDefaultLabel
- FastFriends.UI.OnClickCharOverrideOffLabel
- FastFriends.UI.OnClickCharOverrideOnLabel
- FastFriends.UI.OnClickFriendsLabel
- FastFriends.UI.OnClickIgnoreLabel
- FastFriends.UI.OnClickModeAllLabel
- FastFriends.UI.OnClickModeLevelLabel
- FastFriends.UI.OnClickModeOptInLabel
- FastFriends.UI.OnClickSyncMasterLabel
- LibAddonButton.Manager.CustomItem.OnMacroHelpLUp
- LibAddonButton.Manager.OnNavigateLUp
- MassMailWindow.AddCraftingItems
- MassMailWindow.AddRarityItems
- Motion.OnLButtonUp_ContextMenu_Recipe_Delete
- Obsidian.Setup.Castbar.OnLatencyTextFontLUp
- Obsidian.Setup.Castbar.OnNameFontLUp
- Obsidian.Setup.Castbar.OnTimerFontLUp
- Obsidian.Setup.EffectTracker.OnNameFontLUp
- Obsidian.Setup.EffectTracker.OnTimerFontLUp
- PartyAdWindow.OnToggleDetectedNeedsLabel
- PartyAdWindow.OnToggleSpecCheckLabel
- ReferWindow.OnClickName
- Rolodex.SetName
- SOR.OpenNotes
- TalismanGenie.CurioToggle
- TalismanGenie.EssenceToggle
- TalismanGenie.FragmentToggle
- TalismanGenie.GoldDustToggle
- TexturedButtons.Setup.Fonts.OnCooldownFontLUp
- TexturedButtons.Setup.Fonts.OnKeybindFontLUp
- TurretRange.Setup.Display.OnDistanceFontLUp
- WARCommander.Assist
- WARCommander.HighlightLast0
- WARCommander.HighlightLast1
- WARCommander.HighlightLast2
- zMailModOptions.SelectOption
- zMailModOptions.ShowOptions
- FrameManager.OnMouseOverEnd
- AutoBandWindowConfig.OnSettingMouseOverEnd
- AutoBandWindowTemplate.OnRoleLimitsMouseOverEnd
- SORPager.BOsOff
- Obsidian.Setup.Castbar.OnFontExampleMouseOut
- WARCommander.LabelTintRemove
- AutoBandWindowConfig.OnBackfillMouseOverEnd
- AutoBandWindowConfig.OnHealerRankReqMouseOverEnd
- AutoBandWindowConfig.OnRankReqMouseOverEnd
- Obsidian.Setup.EffectTracker.OnFontExampleMouseOut
- TexturedButtons.Setup.Fonts.OnFontExampleMouseOut
- AutoBandWindowConfig.OnCommonRaceNamesMouseOverEnd
- AutoBandWindowConfig.OnExcludeRealmHealerAltSpecMouseOverEnd
- AutoBandWindowConfig.OnKickLowRankMouseOverEnd
- AutoBandWindowTools.OnDiscordReqMouseOverEnd
- AutoBandWindowTools.OnNoMicMouseOverEnd
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOut
- BuffHead.Setup.Layout.Properties.OnFontExampleMouseOut
- BuffHead.Setup.Layout.Properties.OnPropertyTitleMouseOut
- LibAddonButton.Manager.CustomItem.OnMacroHelpMouseOut
- LibAddonButton.Manager.OnNavigateMouseOut
- MassMailWindow.MouseOverEndCraftingSkill
- MassMailWindow.MouseOverEndRarity
- Motion.OnMouseOverEnd_ContextMenu_Recipe_Delete
- PartyAdWindow.OnMouseOverEndPreviewNoteStatus
- TalismanGenie.CurioOff
- TalismanGenie.EssenceOff
- TalismanGenie.FragmentOff
- TalismanGenie.GoldDustOff
- TurretRange.Setup.Display.OnDistanceFontMouseOut
- WARCommander.AssistLabelTintRemove
- emotes.OnMouseOverPathRowEnd
- zMailModOptions.MouseOverEndListLabel
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick
- Enemy.CombatLogUI_EpsWindow_OnRButtonUp
- Enemy.CombatLogUI_TargetDefenseWindow_OnRButtonUp
- SORRealms.ShowVPMenu
- SORRealms.ShowLowVPMenu
- ClosetGoblinCharacterWindow.OnSetRowContextMenu
- ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
- MiracleGrowLight.switchMode
- Rotation.Delete
- SORRealms.ShowCityMenu
- SORRealms.ShowFortressMenu
- TortallDPSMeter.OnRButtonUp
- EA_ChatWindow.OnHyperLinkRButtonUp
- MapPin.RButtonUp
- EA_ChatWindow.OnHyperLinkLButtonUp
- Enemy.CombatLogUI_IDS_OnRowLButtonDown
- GuildWardenWin.Offline
- GuildWardenWin.classSort
- GuildWardenWin.dateSort
- GuildWardenWin.levelSort
- GuildWardenWin.nameSort
- GuildWardenWin.sigilSort
- AtlasMap.OnLegendItemMouseLeave


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | LPET.OnMouseOver, Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip, Cheeseboard.OnMouseOver, AutoBandWindowConfig.OnSettingMouseOver, Enemy.CombatLogUI_EpsWindow_OnMouseOver, Enemy.ConfigurationWindow_ShowTooltip, AutoBandWindowTemplate.OnRoleLimitsMouseOver, Enemy.CombatLogUI_StatsWindow_ListRowMouseOver, Enemy.CombatLogUI_TargetDefenseWindow_OnMouseOver, DuffTimer.Options.OnMouseOver, Enemy.CombatLogUI_StatsWindow_OnEpsMouseOver, SORRealms.OnMouseOverVP, Enemy.KillSpamUI_KillSpamDialog_OnTotalMouseOver, Obsidian.Setup.Castbar.OnFontExampleMouseOver, SORPager.LowBOsOn, SORRealms.OnMouseOverLowVP, WARCommander.LabelTint, AutoBandWindowConfig.OnBackfillMouseOver, AutoBandWindowConfig.OnHealerRankReqMouseOver, AutoBandWindowConfig.OnRankReqMouseOver, Obsidian.Setup.EffectTracker.OnFontExampleMouseOver, TexturedButtons.Setup.Fonts.OnFontExampleMouseOver, Warbuilder.Effecttooltip, AggroMeter.SelectChar, AtlasMap.OnLegendItemMouseOver, AutoBandWindowConfig.OnCommonRaceNamesMouseOver, AutoBandWindowConfig.OnExcludeRealmHealerAltSpecMouseOver, AutoBandWindowConfig.OnKickLowRankMouseOver, AutoBandWindowTools.OnDiscordReqMouseOver, AutoBandWindowTools.OnNoMicMouseOver, BarText_Influence.OnMouseOver, BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOver, BuffHead.Setup.Layout.Properties.OnFontExampleMouseOver, BuffHead.Setup.Layout.Properties.OnPropertyTitleMouseOver, CMapWindow.OnMouseOverDigitinf, CaVESWindow.OnMouseoverStatHeaderLabelRecent, CaVESWindow.OnMouseoverStatHeaderLabelReference, CaVESWindow.OnMouseoverStatHeaderLabelStat, DeepSleep.Settings.OnMouseOver, EA_Window_OpenPartyLootRollOptions.OnMouseOverRowTitle, GuildWardenTTips.anni, GuildWardenTTips.bl, GuildWardenTTips.conq, GuildWardenTTips.dp, GuildWardenTTips.inv, GuildWardenTTips.sent, GuildWardenTTips.sigil, GuildWardenTTips.tyr, GuildWardenTTips.war, KeyBarSettings.CheckBoxOnMouseOver, LibAddonButton.Manager.CustomItem.OnMacroHelpMouseOver, LibAddonButton.Manager.OnNavigateMouseOver, MassMailWindow.MouseOverCraftingSkill, MassMailWindow.MouseOverRarity, MiracleGrowLight.onHover, Motion.OnMouseOver_ContextMenu_Recipe_Delete, PartyAdWindow.OnMouseOverPreviewNoteStatus, RvRContribution.OnHoverLoss, RvRContribution.OnHoverWin, SORPager.HighBOsOn, TalismanGenie.ContainerMouseOver, TalismanGenie.CurioOn, TalismanGenie.EssenceOn, TalismanGenie.FragmentMouseOver, TalismanGenie.FragmentOn, TalismanGenie.GoldDustOn, TalismanGenie.LevelMouseOver, TalismanGenie.TypeMouseOver, TurretRange.Setup.Display.OnDistanceFontMouseOver, UiModWindow.OnMouseOverModType, UiModWindow.OnMouseOverStatus, WARCommander.AssistLabelTint, emotes.OnMouseOverPathRow, zMailModInbox.OnMouseOverCOD, zMailModOptions.MouseOverListLabel | `` |  |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | Cheeseboard.SortColumnClick, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick, SessionRPs.GetSortType, , SORRealms.OnLClickTier, Enemy.KillSpamUI_KillSpamDialog_OnTotalLButtonUp, AutoBandWindowConfig.OnClickBackfillLabel, ClosetGoblinCharacterWindow.OnClickSetRow, ClosetGoblinZoneWindow.OnClickZoneRow, DPSMeter.SortChange, Deathblow.LinkDeath, Deathblow.LinkKill, Enemy.CombatLogUI_EpsWindow_OnInDpsLButtonUp, Enemy.CombatLogUI_EpsWindow_OnInHpsLButtonUp, Enemy.CombatLogUI_EpsWindow_OnOutDpsAoeLButtonUp, Enemy.CombatLogUI_EpsWindow_OnOutDpsLButtonUp, Enemy.CombatLogUI_EpsWindow_OnOutHpsAoeLButtonUp, Enemy.CombatLogUI_EpsWindow_OnOutHpsLButtonUp, SORRealms.OnLClickBZ, AggroMeter.SelectChar, AutoBandWindowConfig.OnClickAltCheckLabel, AutoBandWindowConfig.OnClickAutoKickLabel, AutoBandWindowConfig.OnClickBWSorcAsMDPSLabel, AutoBandWindowConfig.OnClickCommonRaceNamesLabel, AutoBandWindowConfig.OnClickDpsWeightingLabel, AutoBandWindowConfig.OnClickExcludeRealmHealerAltSpecLabel, AutoBandWindowConfig.OnClickGuildPriorityLabel, AutoBandWindowConfig.OnClickKickIgnoreLabel, AutoBandWindowConfig.OnClickKickLowRankLabel, AutoBandWindowConfig.OnClickKickRvRZonesLabel, AutoBandWindowConfig.OnClickKickStealthersLabel, AutoBandWindowConfig.OnClickKickToofarLabel, AutoBandWindowConfig.OnClickRestrictRaceLabel, AutoBandWindowTemplate.OnClickRightClickTemplateMenuLabel, AutoBandWindowTools.OnClickAutoFormSearchLabel, AutoBandWindowTools.OnClickAutoPartyNoteLabel, AutoBandWindowTools.OnClickAutoSearchLabel, AutoBandWindowTools.OnClickDiscordReqLabel, AutoBandWindowTools.OnClickKickOfflineLabel, AutoBandWindowTools.OnClickKickRankLabel, AutoBandWindowTools.OnClickKickZoneLabel, AutoBandWindowTools.OnClickNoMicLabel, AutoBandWindowTools.OnClickNotifyBuffsLabel, AutoBandWindowTools.OnClickPrintRoleLabel, AutoBandWindowTools.OnClickRightClickOrganizeLabel, AutoBandWindowTools.OnClickSearchRoleChan1Label, AutoBandWindowTools.OnClickSearchRoleChan5Label, AutoBandWindowTools.OnClickSearchRoleChanT4Label, BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleLClick, BuffHead.Setup.Layout.Manager.OnExportLUp, BuffHead.Setup.Layout.Manager.OnImportLUp, BuffHead.Setup.Layout.Manager.OnLayoutsLUp, BuffHead.Setup.Layout.Properties.OnFontFontLUp, BuffHead.Setup.Layout.Properties.OnPropertyTitleLClick, CMapWindow.OnClickDigitinf, Enemy.CombatLogUI_StatsWindow_OnEpsAoeCurClick, Enemy.CombatLogUI_StatsWindow_OnEpsAoeMaxClick, Enemy.CombatLogUI_StatsWindow_OnEpsCurClick, Enemy.CombatLogUI_StatsWindow_OnEpsMaxClick, FastFriends.UI.OnClickCharOverrideDefaultLabel, FastFriends.UI.OnClickCharOverrideOffLabel, FastFriends.UI.OnClickCharOverrideOnLabel, FastFriends.UI.OnClickFriendsLabel, FastFriends.UI.OnClickIgnoreLabel, FastFriends.UI.OnClickModeAllLabel, FastFriends.UI.OnClickModeLevelLabel, FastFriends.UI.OnClickModeOptInLabel, FastFriends.UI.OnClickSyncMasterLabel, LibAddonButton.Manager.CustomItem.OnMacroHelpLUp, LibAddonButton.Manager.OnNavigateLUp, MassMailWindow.AddCraftingItems, MassMailWindow.AddRarityItems, Motion.OnLButtonUp_ContextMenu_Recipe_Delete, Obsidian.Setup.Castbar.OnLatencyTextFontLUp, Obsidian.Setup.Castbar.OnNameFontLUp, Obsidian.Setup.Castbar.OnTimerFontLUp, Obsidian.Setup.EffectTracker.OnNameFontLUp, Obsidian.Setup.EffectTracker.OnTimerFontLUp, PartyAdWindow.OnToggleDetectedNeedsLabel, PartyAdWindow.OnToggleSpecCheckLabel, ReferWindow.OnClickName, Rolodex.SetName, SOR.OpenNotes, TalismanGenie.CurioToggle, TalismanGenie.EssenceToggle, TalismanGenie.FragmentToggle, TalismanGenie.GoldDustToggle, TexturedButtons.Setup.Fonts.OnCooldownFontLUp, TexturedButtons.Setup.Fonts.OnKeybindFontLUp, TurretRange.Setup.Display.OnDistanceFontLUp, WARCommander.Assist, WARCommander.HighlightLast0, WARCommander.HighlightLast1, WARCommander.HighlightLast2, emotes.OnMouseOverPathRow, zMailModOptions.SelectOption, zMailModOptions.ShowOptions | `flags, x, y` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | FrameManager.OnMouseOverEnd, AutoBandWindowConfig.OnSettingMouseOverEnd, AutoBandWindowTemplate.OnRoleLimitsMouseOverEnd, SORPager.BOsOff, Obsidian.Setup.Castbar.OnFontExampleMouseOut, WARCommander.LabelTintRemove, AutoBandWindowConfig.OnBackfillMouseOverEnd, AutoBandWindowConfig.OnHealerRankReqMouseOverEnd, AutoBandWindowConfig.OnRankReqMouseOverEnd, Obsidian.Setup.EffectTracker.OnFontExampleMouseOut, TexturedButtons.Setup.Fonts.OnFontExampleMouseOut, AutoBandWindowConfig.OnCommonRaceNamesMouseOverEnd, AutoBandWindowConfig.OnExcludeRealmHealerAltSpecMouseOverEnd, AutoBandWindowConfig.OnKickLowRankMouseOverEnd, AutoBandWindowTools.OnDiscordReqMouseOverEnd, AutoBandWindowTools.OnNoMicMouseOverEnd, BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOut, BuffHead.Setup.Layout.Properties.OnFontExampleMouseOut, BuffHead.Setup.Layout.Properties.OnPropertyTitleMouseOut, LibAddonButton.Manager.CustomItem.OnMacroHelpMouseOut, LibAddonButton.Manager.OnNavigateMouseOut, MassMailWindow.MouseOverEndCraftingSkill, MassMailWindow.MouseOverEndRarity, Motion.OnMouseOverEnd_ContextMenu_Recipe_Delete, PartyAdWindow.OnMouseOverEndPreviewNoteStatus, TalismanGenie.CurioOff, TalismanGenie.EssenceOff, TalismanGenie.FragmentOff, TalismanGenie.GoldDustOff, TurretRange.Setup.Display.OnDistanceFontMouseOut, WARCommander.AssistLabelTintRemove, emotes.OnMouseOverPathRowEnd, zMailModOptions.MouseOverEndListLabel | `` |  |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick, Enemy.CombatLogUI_EpsWindow_OnRButtonUp, Enemy.CombatLogUI_TargetDefenseWindow_OnRButtonUp, SORRealms.ShowVPMenu, SORRealms.ShowLowVPMenu, ClosetGoblinCharacterWindow.OnSetRowContextMenu, ClosetGoblinZoneWindow.OnSetZoneRowContextMenu, MiracleGrowLight.switchMode, Rotation.Delete, SORRealms.ShowCityMenu, SORRealms.ShowFortressMenu, TortallDPSMeter.OnRButtonUp | `flags, x, y` | MEDIUM |
| OnHyperLinkRButtonUp | custom | EA_ChatWindow.OnHyperLinkRButtonUp, MapPin.RButtonUp | `linkData, flags, x, y` | LOW |
| OnHyperLinkLButtonUp | custom | EA_ChatWindow.OnHyperLinkLButtonUp | `linkData, flags, x, y` | LOW |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | Enemy.CombatLogUI_IDS_OnRowLButtonDown, GuildWardenWin.Offline, GuildWardenWin.classSort, GuildWardenWin.dateSort, GuildWardenWin.levelSort, GuildWardenWin.nameSort, GuildWardenWin.sigilSort | `flags, x, y` | MEDIUM |
| [OnMouseLeave](../handlers/handler_OnMouseLeave.md) | custom | AtlasMap.OnLegendItemMouseLeave | `` |  |

### Per-Event Lua API Calls

**OnMouseOver** handlers call: `CreateWindow`, `DoesWindowExist`, `LabelGetTextColor`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `SystemData.ActiveWindow.name:match`, `SystemData.ActiveWindow.name:sub`, `SystemData.MouseOverWindow.name:match`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetGameActionData`

**OnLButtonUp** handlers call: `BroadcastEvent`, `LabelGetText`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `SystemData.ActiveWindow.name:match`, `WindowGetId`, `WindowGetParent`, `WindowGetParent:find`, `WindowSetGameActionData`, `WindowSetShowing`

**OnMouseOverEnd** handlers call: `DoesWindowExist`, `LabelGetText`, `LabelSetTextColor`, `WindowGetId`

**OnRButtonUp** handlers call: `SystemData.MouseOverWindow.name:match`, `WindowGetId`

**OnHyperLinkRButtonUp** handlers call: `ButtonSetPressedFlag`, `ComboBoxSetSelectedMenuItem`, `SliderBarSetCurrentPosition`, `TextEditBoxSetText`, `WindowGetId`, `WindowSetMovable`, `WindowSetShowing`

**OnHyperLinkLButtonUp** handlers call: `WindowSetShowing`

**OnLButtonDown** handlers call: `LabelSetText`, `SystemData.ActiveWindow.name:match`

## Common Inherits

- Aura_Default_Label_SmallFont
- ClosetGoblinTalismanLabel
- DefaultWindowSmallText
- DefaultWindowText
- EA_CheckButtonLabel
- EA_Label_DefaultText
- EA_Label_DefaultText_Small
- EA_Settings_ItemTitle
- GesLabelTemplate
- HGG_ItemLabelTemplate
- Small_Label
- TOKText

## Common Parent Elements

- [Windows](element_Windows.md) — 4814× (HIGH)
- [Window](element_Window.md) — 15× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 4631× (HIGH)
- [Size](element_Size.md) — 4173× (HIGH)
- [Color](element_Color.md) — 1898× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 330× (HIGH)
- [Text](element_Text.md) — 69× (HIGH)
- [TintColor](element_TintColor.md) — 30× (HIGH)
- [Windows](element_Windows.md) — 14× (HIGH)
- [Anchor](element_Anchor.md) — 2× (LOW)
- [Sounds](element_Sounds.md) — 1× (LOW)
- [color](element_color.md) — 1× (LOW)

## Common Template Bases

- Aggro_Label_Template
- AlertTextLine
- AuctionWindowControlLabel
- AuctionWindowListLabel
- Aura_CheckButtonLabel
- Aura_Default_Label
- Aura_Default_Label_SmallFont
- BlackBookWindow_EntryRowText
- ClosetGoblinTalismanLabel
- CraftingWillardLabelC
- CrusherDefaultLabelSmallFont
- DefaultWindowSmallText
- DefaultWindowSubTitle
- DefaultWindowText
- Default_Label
- EA_CheckButtonLabel
- EA_Label_ChatText
- EA_Label_DefaultSubHeading
- EA_Label_DefaultText
- EA_Label_DefaultText_Small
- EA_Label_SocialListItem
- EA_Settings_ItemTitle
- EA_Settings_SectionTitle
- EA_Template_AutoRollRarityHeader
- EA_Template_AutoRollTitle
- EA_Template_OpenParty_Label
- EA_Template_OpenParty_Label_Small
- EA_Window_HelpDescriptionTemplate
- GDesLabel_Template
- GesLabelTemplate
- HGG_ItemLabelTemplate
- HGG_SubHeadingTemplate
- HG_HealGridUnitLabelTemplate
- HopperDefaultLabelSmallFont
- IraConfigHeading
- IraConfigLabel
- LibAddonButtonManagerWindowNavigateButton1
- MailRowTextTemplate
- MapMonsterEditorWindowHeaderDefault
- MapMonsterEditorWindowLabelDefault
- MapMonsterPinTypeEditorWindowHeaderDefault
- MapMonsterPinTypeEditorWindowLabelDefault
- MarkBuffSetupWindowGroupLabel1
- MiracleGrow2Label
- ModDetailsLabelDef
- ModDetailsTextDef
- MotionDefaultLabelSmallFont
- PieTrackerLabelTemplate
- PieTrackerRowTextTemplate
- PureCheckButtonLabel
- PureDefaultLabelSmallFont
- RMet_SubHeadingTemplate
- Raid_Label_Template
- Raid_Label_Template2
- ResHelp_Label_Template
- ResLabelTemplate
- SendMailHeaderTemplate
- Shinies_Default_Label
- Shinies_Default_Label_ClearLargeFont
- Shinies_Default_Label_ClearMediumFont
- Shinies_Default_Label_ClearSmallFont
- SiegeChatChannelLabelTemplate
- Small_Label
- SocialWindowOptionsHeaderTemplate
- SocialWindowTabFriendsHeaderTemplate
- SocialWindowTabSearchHeaderTemplate
- SpamBayesTemplateBText
- SpamBayesTemplateButtonSeparator
- SpamBayesTemplateHelpTextC1
- SpamBayesTemplateHelpTextC2
- TAtlasLabel
- TChatLabel
- TOKHeadingSmall
- TOKSubHeading
- TOKText
- TOKTextHuge
- TOKTextItalic
- TOKTextLarge
- TOKTextSmall
- TOKTitleLarge
- TOKTitleMedium
- TOKTitleSmall
- TRollLabel
- TaxPayerLabelTemplate
- TaxPayerRowTextTemplate
- TooltipItemTitle
- UI_KwestorTrackerQuestNameTemplate
- WARCommanderConfigWindowHeaderDefault
- WARCommanderConfigWindowLabelDefault
- WCDBDefaultLabelSmallFont
- WCDPDefaultLabelSmallFont
- WTesLabelTemplate
- zMM_Default_CraftingLabel
- zMM_Default_RarityLabel
- zMailModOptionsHeader
- zMailModOptionsLabel
- zMailModOptionsListLabel
- zMailModOptionsLongLabel
- zMailModOptionsShortLabel
- zMailModOptionsSubheader


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `textalign` | **required** | 82% | center, left, leftcenter, right, ... |
| `font` | optional | 66% | font_heading_huge, font_default_text, font_clear_small_bold, font_clear_medium_bold, ... |
| `inherits` | optional | 40% | EA_Label_DefaultText, DefaultWindowText, TOKText, Aggro_Label_Template, ... |
| `handleinput` | optional | 36% | false, true |
| `autoresize` | optional | 34% | true, false, center |
| `layer` | optional | 28% | popup, default, secondary, overlay, ... |
| `popable` | optional | 19% | false, true |
| `wordwrap` | optional | 17% | false, true |
| `maxchars` | optional | 16% | 5, 256, 15, 12, ... |
| `warnOnTextCropped` | optional | 13% | false |
| `autoresizewidth` | optional | 4% | true, false |
| `textAutoFitMinScale` | optional | 3% | 0.3, 0.7, 0.5, .75, ... |
| `skipinput` | optional | 1% | true, false |
| `sticky` | optional | 1% | true, false |
| `alpha` | optional | 0% | 1, 0.8, 0.75, 0, ... |
| `id` | optional | 0% | 1, 2, 3, 4, ... |
| `hanldeinput` | optional | 0% | false |
| `handlinput` | optional | 0% | false |
| `show` | optional | 0% | false |
| `autoresizeheight` | optional | 0% | true |
| `ellipsisOnCrop` | optional | 0% | true |
| `text` | optional | 0% | Tweaks content goes here. |
| `savesettings` | optional | 0% | true, false |
| `autosize` | optional | 0% | true |
| `virtual` | optional | 0% | true |
| `fontscale` | optional | 0% | 0.5 |
| `handleInput` | optional | 0% | true |
| `ignoreFormattingTags` | optional | 0% | true |
| `linespacing` | optional | 0% | 13, 12 |
| `scale` | optional | 0% | 0.5 |
| `showing` | optional | 0% | false |
| `textautofitminscale` | optional | 0% | 0.5 |
| `visible` | optional | 0% | true, false |
| `FontAlpha` | optional | 0% | 0 |
| `align` | optional | 0% | right |
| `autoscale` | optional | 0% | true |
| `background` | optional | 0% | EA_FullResizeImage_TanBorder |
| `maxlines` | optional | 0% | 2 |
| `mouseovertext` | optional | 0% | false |
| `movable` | optional | 0% | false |
| `resizewidth` | optional | 0% | true |
| `textAlign` | optional | 0% | center |
| `textUsesMarkup` | optional | 0% | true |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 4631 times as an unnamed child.

### [Size](element_Size.md)

Observed 4173 times as an unnamed child.

### [Color](element_Color.md)

Observed 1898 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 255, 0, 73, 225 |
| `g` | **required** | 25, 253, 218, 255 |
| `r` | **required** | 255, 253, 155, 245 |
| `a` | optional | 255, 0.8, 0, 128 |
### [EventHandlers](element_EventHandlers.md)

Observed 330 times as an unnamed child.

### [Text](element_Text.md)

Observed 69 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `text` | optional | Killer Settings, Changes apply immediately. Zone K/D history uses 0 for unlimited saved zone leaderboards., Display, Personal |
| `alignment` | optional | left |
### [TintColor](element_TintColor.md)

Observed 30 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 110, 130, 50 |
| `g` | **required** | 0, 200, 185, 130 |
| `r` | **required** | 255, 0, 180, 200 |
| `a` | optional | 30, 255, 0.8, 175 |
### [Windows](element_Windows.md)

Observed 14 times as an unnamed child.

### [Anchor](element_Anchor.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `point` | **required** | center, topleft, topright, bottomright |
| `relativePoint` | **required** | center, topleft, topright, bottomright |
| `relativeTo` | **required** | Root, $parent, $parentSliderLabel, $parentSliderValueLabel |
| `relativeto` | optional | $parentTitle, $parentPrint, $parentAlert, $parentHyperlink |
| `layer` | optional | secondary, overlay |
| `parent` | optional | Root, $parent |
| `relateiveTo` | optional | $parentDevPadCode, $parentObjectEditBox |
| `textalign` | optional | center |
| `handleinput` | optional | false |
| `relative` | optional | $parent |
| `xOffset` | optional | 0 |
| `yOffset` | optional | 0 |
| `input` | optional | numbers |
| `realtivePoint` | optional | center |
| `realtiveTo` | optional | $parentBackground |
### [Sounds](element_Sounds.md)

Observed 1 times as an unnamed child.

### [color](element_color.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `a` | **required** | 255 |
| `b` | **required** | 240 |
| `g` | **required** | 240 |
| `r` | **required** | 240 |
## Recursive Hierarchy

- Root: [Label](element_Label.md)
- [Anchor](element_Anchor.md) (structural, 2×, LOW)
  - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
    - (cycle)
- [Anchors](element_Anchors.md) (structural, 4631×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
      - (cycle)
- [Color](element_Color.md) (structural, 1898×, HIGH)
- [EventHandlers](element_EventHandlers.md) (structural, 330×, HIGH)
  - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
- [Size](element_Size.md) (structural, 4173×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
- [Sounds](element_Sounds.md) (structural, 1×, LOW)
  - [Sound](element_Sound.md) (structural, 20×, HIGH)
- [Text](element_Text.md) (structural, 69×, HIGH)
- [TintColor](element_TintColor.md) (structural, 30×, HIGH)
- [Windows](element_Windows.md) (structural, 14×, HIGH)
  - [ActionButtonGroup](element_ActionButtonGroup.md) (named, 3×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 3×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
  - [AnimatedImage](element_AnimatedImage.md) (named, 39×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 38×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [AnimFrames](element_AnimFrames.md) (structural, 19×, HIGH)
      - [AnimFrame](element_AnimFrame.md) (structural, 222×, HIGH)
    - [Size](element_Size.md) (structural, 32×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [TexDims](element_TexDims.md) (structural, 1×, LOW)
    - [Windows](element_Windows.md) (structural, 2×, MEDIUM)
      - (cycle)
  - [Button](element_Button.md) (named, 2407×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 2×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 1954×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [AnimatedImages](element_AnimatedImages.md) (structural, 1×, LOW)
      - [Normal](element_Normal.md) (structural, 1×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)
    - [Color](element_Color.md) (structural, 27×, HIGH)
    - [EventHandler](element_EventHandler.md) (structural, 1×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 1540×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Eventhandlers](element_Eventhandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 2×, HIGH)
    - [OverlayOffset](element_OverlayOffset.md) (structural, 31×, HIGH)
    - [OverlaySize](element_OverlaySize.md) (structural, 27×, HIGH)
    - [OverlayTexCoords](element_OverlayTexCoords.md) (structural, 24×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 20×, HIGH)
      - [Normal](element_Normal.md) (structural, 24×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 24×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 24×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 24×, HIGH)
    - [OverlayTexSlices](element_OverlayTexSlices.md) (structural, 2×, LOW)
      - [Normal](element_Normal.md) (structural, 2×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 2×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 2×, HIGH)
    - [ResizeImages](element_ResizeImages.md) (structural, 41×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 23×, HIGH)
      - [Normal](element_Normal.md) (structural, 25×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 41×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 26×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 27×, HIGH)
    - [Size](element_Size.md) (structural, 1023×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sounds](element_Sounds.md) (structural, 9×, MEDIUM)
      - [Sound](element_Sound.md) (structural, 20×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 47×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
      - [Left](element_Left.md) (structural, 33×, HIGH)
      - [Middle](element_Middle.md) (structural, 41×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
      - [Normal](element_Normal.md) (structural, 47×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
      - [Right](element_Right.md) (structural, 33×, HIGH)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
    - [TexDims](element_TexDims.md) (structural, 18×, HIGH)
    - [TexSlices](element_TexSlices.md) (structural, 129×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
      - [Left](element_Left.md) (structural, 3×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 10×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
      - [Normal](element_Normal.md) (structural, 120×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
      - [Right](element_Right.md) (structural, 3×, MEDIUM)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
    - [Text](element_Text.md) (structural, 3×, MEDIUM)
    - [TextColors](element_TextColors.md) (structural, 85×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 76×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 26×, HIGH)
      - [Normal](element_Normal.md) (structural, 79×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 83×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 76×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 83×, HIGH)
    - [TextOffset](element_TextOffset.md) (structural, 75×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 12×, HIGH)
    - [Windows](element_Windows.md) (structural, 131×, HIGH)
      - (cycle)
  - [CircleImage](element_CircleImage.md) (named, 56×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 54×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 51×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 37×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
      - [Left](element_Left.md) (structural, 33×, HIGH)
      - [Middle](element_Middle.md) (structural, 41×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
      - [Normal](element_Normal.md) (structural, 47×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
      - [Right](element_Right.md) (structural, 33×, HIGH)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
    - [TexDims](element_TexDims.md) (structural, 11×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 15×, HIGH)
    - [Windows](element_Windows.md) (structural, 2×, LOW)
      - (cycle)
  - [ColorPicker](element_ColorPicker.md) (named, 1×, LOW)
    - [Anchors](element_Anchors.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [ColorSize](element_ColorSize.md) (structural, 1×, HIGH)
    - [ColorSpacing](element_ColorSpacing.md) (structural, 1×, HIGH)
    - [ColorTexCoords](element_ColorTexCoords.md) (structural, 1×, HIGH)
    - [ColorTexDims](element_ColorTexDims.md) (structural, 1×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
  - [ComboBox](element_ComboBox.md) (named, 689×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 643×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 534×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [MenuButtonOffset](element_MenuButtonOffset.md) (structural, 33×, HIGH)
    - [OverlayOffset](element_OverlayOffset.md) (structural, 4×, MEDIUM)
    - [Size](element_Size.md) (structural, 186×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [TextOffset](element_TextOffset.md) (structural, 1×, LOW)
  - [CooldownDisplay](element_CooldownDisplay.md) (named, 10×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 10×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 10×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 4×, MEDIUM)
      - (cycle)
  - [DynamicImage](element_DynamicImage.md) (named, 1288×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 1187×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 1×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 155×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 965×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sounds](element_Sounds.md) (structural, 1×, LOW)
      - [Sound](element_Sound.md) (structural, 20×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 76×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
      - [Left](element_Left.md) (structural, 33×, HIGH)
      - [Middle](element_Middle.md) (structural, 41×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
      - [Normal](element_Normal.md) (structural, 47×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
      - [Right](element_Right.md) (structural, 33×, HIGH)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
    - [TexDims](element_TexDims.md) (structural, 156×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 374×, HIGH)
    - [Windows](element_Windows.md) (structural, 15×, HIGH)
      - (cycle)
  - [EditBox](element_EditBox.md) (named, 416×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 380×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 2×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 269×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 355×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [TextOffset](element_TextOffset.md) (structural, 63×, HIGH)
    - [Windows](element_Windows.md) (structural, 3×, MEDIUM)
      - (cycle)
  - [FullResizeImage](element_FullResizeImage.md) (named, 450×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 2×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 409×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 2×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 5×, MEDIUM)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 113×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 30×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 30×, HIGH)
      - [Middle](element_Middle.md) (structural, 30×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 30×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 19×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
      - [Left](element_Left.md) (structural, 33×, HIGH)
      - [Middle](element_Middle.md) (structural, 41×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
      - [Normal](element_Normal.md) (structural, 47×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
      - [Right](element_Right.md) (structural, 33×, HIGH)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
    - [TexDims](element_TexDims.md) (structural, 3×, MEDIUM)
    - [TexSlices](element_TexSlices.md) (structural, 11×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
      - [Left](element_Left.md) (structural, 3×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 10×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
      - [Normal](element_Normal.md) (structural, 120×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
      - [Right](element_Right.md) (structural, 3×, MEDIUM)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 118×, HIGH)
    - [Windows](element_Windows.md) (structural, 4×, MEDIUM)
      - (cycle)
  - [HorizontalResizeImage](element_HorizontalResizeImage.md) (named, 90×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 48×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 49×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 43×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 30×, HIGH)
      - [Middle](element_Middle.md) (structural, 30×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 30×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 36×, HIGH)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
      - [Left](element_Left.md) (structural, 33×, HIGH)
      - [Middle](element_Middle.md) (structural, 41×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
      - [Normal](element_Normal.md) (structural, 47×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
      - [Right](element_Right.md) (structural, 33×, HIGH)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
    - [TexSlices](element_TexSlices.md) (structural, 7×, MEDIUM)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
      - [Left](element_Left.md) (structural, 3×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 10×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
      - [Normal](element_Normal.md) (structural, 120×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
      - [Right](element_Right.md) (structural, 3×, MEDIUM)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 20×, HIGH)
    - [Windows](element_Windows.md) (structural, 3×, MEDIUM)
      - (cycle)
  - [Label](element_Label.md) (named, 4814×, HIGH)
    - (cycle)
  - [ListBox](element_ListBox.md) (named, 113×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 110×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 16×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [ListData](element_ListData.md) (structural, 110×, HIGH)
      - [ListColumns](element_ListColumns.md) (structural, 74×, HIGH)
        - [ListColumn](element_ListColumn.md) (structural, 192×, HIGH)
    - [Size](element_Size.md) (structural, 76×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 1×, LOW)
      - (cycle)
  - [LogDisplay](element_LogDisplay.md) (named, 3×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 3×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 3×, HIGH)
      - (cycle)
  - [MapDisplay](element_MapDisplay.md) (named, 8×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 8×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 8×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 3×, MEDIUM)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
  - [NifDisplay](element_NifDisplay.md) (named, 50×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Rotation](element_Rotation.md) (structural, 4×, MEDIUM)
    - [Size](element_Size.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Translation](element_Translation.md) (structural, 50×, HIGH)
  - [PageWindow](element_PageWindow.md) (named, 4×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 1×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 4×, HIGH)
      - (cycle)
  - [ScrollWindow](element_ScrollWindow.md) (named, 61×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 43×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 21×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 53×, HIGH)
      - (cycle)
  - [SliderBar](element_SliderBar.md) (named, 225×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 219×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 205×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 186×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Windows](element_Windows.md) (structural, 6×, MEDIUM)
      - (cycle)
  - [StatusBar](element_StatusBar.md) (named, 32×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 32×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 19×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
  - [VerticalResizeImage](element_VerticalResizeImage.md) (named, 27×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 14×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 10×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 8×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 30×, HIGH)
      - [Middle](element_Middle.md) (structural, 30×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 30×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 5×, MEDIUM)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
      - [Left](element_Left.md) (structural, 33×, HIGH)
      - [Middle](element_Middle.md) (structural, 41×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
      - [Normal](element_Normal.md) (structural, 47×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
      - [Right](element_Right.md) (structural, 33×, HIGH)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
    - [TexSlices](element_TexSlices.md) (structural, 3×, MEDIUM)
      - [Bottom](element_Bottom.md) (structural, 2×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
      - [Left](element_Left.md) (structural, 3×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 10×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
      - [Normal](element_Normal.md) (structural, 120×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
      - [Right](element_Right.md) (structural, 3×, MEDIUM)
      - [Top](element_Top.md) (structural, 2×, LOW)
      - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 4×, MEDIUM)
  - [VerticalScrollbar](element_VerticalScrollbar.md) (named, 62×, HIGH)
    - [ActiveZoneOffset](element_ActiveZoneOffset.md) (structural, 2×, LOW)
    - [Anchors](element_Anchors.md) (structural, 56×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [DownOffset](element_DownOffset.md) (structural, 2×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 7×, MEDIUM)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 47×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [ThumbOffset](element_ThumbOffset.md) (structural, 2×, LOW)
    - [UpOffset](element_UpOffset.md) (structural, 2×, LOW)
  - [Window](element_Window.md) (named, 3695×, HIGH)
    - [Button](element_Button.md) (named, 8×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 2×, LOW)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
      - [Anchors](element_Anchors.md) (structural, 1954×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
            - (cycle)
      - [AnimatedImages](element_AnimatedImages.md) (structural, 1×, LOW)
        - [Normal](element_Normal.md) (structural, 1×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)
      - [Color](element_Color.md) (structural, 27×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1×, LOW)
      - [EventHandlers](element_EventHandlers.md) (structural, 1540×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
      - [Eventhandlers](element_Eventhandlers.md) (structural, 1×, LOW)
        - [EventHandler](element_EventHandler.md) (structural, 2×, HIGH)
      - [OverlayOffset](element_OverlayOffset.md) (structural, 31×, HIGH)
      - [OverlaySize](element_OverlaySize.md) (structural, 27×, HIGH)
      - [OverlayTexCoords](element_OverlayTexCoords.md) (structural, 24×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 20×, HIGH)
        - [Normal](element_Normal.md) (structural, 24×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 24×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 24×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 24×, HIGH)
      - [OverlayTexSlices](element_OverlayTexSlices.md) (structural, 2×, LOW)
        - [Normal](element_Normal.md) (structural, 2×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 2×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 2×, HIGH)
      - [ResizeImages](element_ResizeImages.md) (structural, 41×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 23×, HIGH)
        - [Normal](element_Normal.md) (structural, 25×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 41×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 26×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 27×, HIGH)
      - [Size](element_Size.md) (structural, 1023×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
      - [Sounds](element_Sounds.md) (structural, 9×, MEDIUM)
        - [Sound](element_Sound.md) (structural, 20×, HIGH)
      - [TexCoords](element_TexCoords.md) (structural, 47×, HIGH)
        - [Bottom](element_Bottom.md) (structural, 2×, LOW)
        - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
        - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
        - [Left](element_Left.md) (structural, 33×, HIGH)
        - [Middle](element_Middle.md) (structural, 41×, HIGH)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
        - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
        - [Normal](element_Normal.md) (structural, 47×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
        - [Right](element_Right.md) (structural, 33×, HIGH)
        - [Top](element_Top.md) (structural, 2×, LOW)
        - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
        - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
      - [TexDims](element_TexDims.md) (structural, 18×, HIGH)
      - [TexSlices](element_TexSlices.md) (structural, 129×, HIGH)
        - [Bottom](element_Bottom.md) (structural, 2×, LOW)
        - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
        - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
        - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
        - [Left](element_Left.md) (structural, 3×, MEDIUM)
        - [Middle](element_Middle.md) (structural, 10×, HIGH)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
        - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
        - [Normal](element_Normal.md) (structural, 120×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
        - [Right](element_Right.md) (structural, 3×, MEDIUM)
        - [Top](element_Top.md) (structural, 2×, LOW)
        - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
        - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
      - [Text](element_Text.md) (structural, 3×, MEDIUM)
      - [TextColors](element_TextColors.md) (structural, 85×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 76×, HIGH)
        - [DisabledPressed](element_DisabledPressed.md) (structural, 26×, HIGH)
        - [Normal](element_Normal.md) (structural, 79×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 83×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 76×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 83×, HIGH)
      - [TextOffset](element_TextOffset.md) (structural, 75×, HIGH)
      - [TintColor](element_TintColor.md) (structural, 12×, HIGH)
      - [Windows](element_Windows.md) (structural, 131×, HIGH)
        - (cycle)
    - [ComboBox](element_ComboBox.md) (named, 4×, MEDIUM)
      - [Anchors](element_Anchors.md) (structural, 643×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
            - (cycle)
      - [EventHandlers](element_EventHandlers.md) (structural, 534×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
      - [MenuButtonOffset](element_MenuButtonOffset.md) (structural, 33×, HIGH)
      - [OverlayOffset](element_OverlayOffset.md) (structural, 4×, MEDIUM)
      - [Size](element_Size.md) (structural, 186×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
      - [TextOffset](element_TextOffset.md) (structural, 1×, LOW)
    - [DynamicImage](element_DynamicImage.md) (named, 1×, LOW)
      - [Anchor](element_Anchor.md) (structural, 1×, LOW)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
      - [Anchors](element_Anchors.md) (structural, 1187×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
            - (cycle)
      - [Color](element_Color.md) (structural, 1×, LOW)
      - [EventHandlers](element_EventHandlers.md) (structural, 155×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
      - [Size](element_Size.md) (structural, 965×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
      - [Sounds](element_Sounds.md) (structural, 1×, LOW)
        - [Sound](element_Sound.md) (structural, 20×, HIGH)
      - [TexCoords](element_TexCoords.md) (structural, 76×, HIGH)
        - [Bottom](element_Bottom.md) (structural, 2×, LOW)
        - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
        - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
        - [Left](element_Left.md) (structural, 33×, HIGH)
        - [Middle](element_Middle.md) (structural, 41×, HIGH)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
        - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
        - [Normal](element_Normal.md) (structural, 47×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
        - [Right](element_Right.md) (structural, 33×, HIGH)
        - [Top](element_Top.md) (structural, 2×, LOW)
        - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
        - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
      - [TexDims](element_TexDims.md) (structural, 156×, HIGH)
      - [TintColor](element_TintColor.md) (structural, 374×, HIGH)
      - [Windows](element_Windows.md) (structural, 15×, HIGH)
        - (cycle)
    - [FullResizeImage](element_FullResizeImage.md) (named, 9×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 2×, LOW)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
      - [Anchors](element_Anchors.md) (structural, 409×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
            - (cycle)
      - [Color](element_Color.md) (structural, 2×, LOW)
      - [EventHandlers](element_EventHandlers.md) (structural, 5×, MEDIUM)
        - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
      - [Size](element_Size.md) (structural, 113×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
      - [Sizes](element_Sizes.md) (structural, 30×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 30×, HIGH)
        - [Middle](element_Middle.md) (structural, 30×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 30×, HIGH)
      - [TexCoords](element_TexCoords.md) (structural, 19×, HIGH)
        - [Bottom](element_Bottom.md) (structural, 2×, LOW)
        - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
        - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
        - [Left](element_Left.md) (structural, 33×, HIGH)
        - [Middle](element_Middle.md) (structural, 41×, HIGH)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
        - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
        - [Normal](element_Normal.md) (structural, 47×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
        - [Right](element_Right.md) (structural, 33×, HIGH)
        - [Top](element_Top.md) (structural, 2×, LOW)
        - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
        - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
      - [TexDims](element_TexDims.md) (structural, 3×, MEDIUM)
      - [TexSlices](element_TexSlices.md) (structural, 11×, HIGH)
        - [Bottom](element_Bottom.md) (structural, 2×, LOW)
        - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
        - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
        - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
        - [Left](element_Left.md) (structural, 3×, MEDIUM)
        - [Middle](element_Middle.md) (structural, 10×, HIGH)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
        - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
        - [Normal](element_Normal.md) (structural, 120×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
        - [Right](element_Right.md) (structural, 3×, MEDIUM)
        - [Top](element_Top.md) (structural, 2×, LOW)
        - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
        - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
      - [TintColor](element_TintColor.md) (structural, 118×, HIGH)
      - [Windows](element_Windows.md) (structural, 4×, MEDIUM)
        - (cycle)
    - [Label](element_Label.md) (named, 15×, HIGH)
      - (cycle)
    - [SliderBar](element_SliderBar.md) (named, 2×, LOW)
      - [Anchors](element_Anchors.md) (structural, 219×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
            - (cycle)
      - [EventHandlers](element_EventHandlers.md) (structural, 205×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
      - [Size](element_Size.md) (structural, 186×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
      - [Windows](element_Windows.md) (structural, 6×, MEDIUM)
        - (cycle)
    - [Window](element_Window.md) (named, 5×, MEDIUM)
      - (cycle)
    - [Anchor](element_Anchor.md) (structural, 16×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 2735×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 773×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 5515×, HIGH)
    - [Size](element_Size.md) (structural, 1752×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
    - [Sounds](element_Sounds.md) (structural, 2×, LOW)
      - [Sound](element_Sound.md) (structural, 20×, HIGH)
    - [Visual](element_Visual.md) (structural, 1×, LOW)
      - [Color](element_Color.md) (structural, 1×, HIGH)
    - [Windows](element_Windows.md) (structural, 1498×, HIGH)
      - (cycle)
- [color](element_color.md) (structural, 1×, LOW)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowGetParent` | ui | 93 | OnLButtonUp, OnMouseOver |
| `LabelSetTextColor` | ui | 84 | OnLButtonUp, OnMouseOver, OnMouseOverEnd |
| `WindowGetId` | ui | 74 | OnHyperLinkRButtonUp, OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp |
| `WindowSetShowing` | ui | 47 | OnHyperLinkLButtonUp, OnHyperLinkRButtonUp, OnLButtonUp |
| `SystemData.ActiveWindow.name:match` | data | 30 | OnLButtonDown, OnLButtonUp, OnMouseOver |
| `SystemData.MouseOverWindow.name:match` | data | 22 | OnMouseOver, OnRButtonUp |
| `LabelSetText` | ui | 18 | OnLButtonDown, OnLButtonUp, OnMouseOver |
| `TextEditBoxSetText` | ui | 12 | OnHyperLinkRButtonUp |
| `ListBoxGetDataIndex` | ui | 9 | OnLButtonUp, OnMouseOver |
| `CreateWindow` | ui | 6 | OnMouseOver |
| `DoesWindowExist` | ui | 6 | OnMouseOver, OnMouseOverEnd |
| `ComboBoxSetSelectedMenuItem` | ui | 4 | OnHyperLinkRButtonUp |
| `SystemData.ActiveWindow.name:sub` | data | 4 | OnMouseOver |
| `LabelGetText` | ui | 3 | OnLButtonUp, OnMouseOverEnd |
| `LabelGetTextColor` | ui | 3 | OnMouseOver |
| `WindowSetGameActionData` | ui | 3 | OnLButtonUp, OnMouseOver |
| `SliderBarSetCurrentPosition` | ui | 2 | OnHyperLinkRButtonUp |
| `BroadcastEvent` | event | 1 | OnLButtonUp |
| `ButtonSetPressedFlag` | ui | 1 | OnHyperLinkRButtonUp |
| `WindowGetDimensions` | ui | 1 | OnMouseOver |
| `WindowGetParent:find` | ui | 1 | OnLButtonUp |
| `WindowGetScreenPosition` | ui | 1 | OnMouseOver |
| `WindowSetMovable` | ui | 1 | OnHyperLinkRButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnHyperLinkLButtonUp

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `linkData` | table | data |
| 1 | `flags` | boolean | modifier_flags |
| 2 | `x` | number | mouse_x |
| 3 | `y` | number | mouse_y |
### OnHyperLinkRButtonUp

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `linkData` | table | data |
| 1 | `flags` | boolean | modifier_flags |
| 2 | `x` | number | mouse_x |
| 3 | `y` | number | mouse_y |
### OnLButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMouseLeave

Confidence: LOW

### OnMouseOver

Confidence: HIGH

### OnMouseOverEnd

Confidence: HIGH

### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
## Lua Functions Manipulating This Type

- DammazKron.local.UpdateSpecial
- Enemy.CombatLogUI_StatsWindow_UpdateList
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- PP.UpdateListRow
- MapPin.SetupAccept
- RealmStatus.Loading
- SocialWindow.OnMouseOverListMember
- WbLeadHelperMessage.MessageDialogOpen
- DammazKronTNFO.PlayerTargetUpdate
- RandomMountUI.OnInitialize
- Enemy.CombatLogUI_EpsWindow_UpdateLayout
- Enemy.GroupsUI_EffectFilterDialog_Open
- PeaceOut.SetupAction
- AdjustTheTip.CreateAbilityTooltip
- Aura.Aura:ShowEditWindows
- MiracleGrowLight.local.setMode
- DPSMeter.UpdateOverallTab
- Enemy.CombatLogUI_TargetDefenseTotalWindow_Update
- Killer.Initialize
- CDown.CCDownFrameN:SetCD
- SOROptions.GetSettings
- CDown.CCDownFrame:Create
- Obsidian.Setup.EffectTracker.OnTimerOffsetXChanged
- MBuffGui.SetTimerText
- wbLeadHelper.onZoneMouseOver
- MiracleGrowLight.OnUpdate
- GuildWardenWin.Offline
- SOROptions.Save
- CDown.CCDownFrameB:UpdateDurI1V
- XpStatus.ChangeDisplayMode
- HealGrid.HealGridCastBar:StartInteract
- SocialWindowTabFriends.Initialize
- Atlas.UpdateZoneName
- BuddyBind.update
- HealGrid.HealGridCastBar:Update
- FastInteract.OptionsSetup
- MoraleCircle.OnSetCustomColor
- BlackBox.HealthUpdate
- CDown.CCDownFrame:UpdateI1
- SOR.Initialize
- SOR.SetupWindows
- GuildWarden.ScrollDown
- MoraleCircle.ColorChanger1
- APAGui.UpdateKitingHUD
- Clock.OnUpdate
- MapPin.Map.UpdatePinCoordinates
- ThankTheTank.OnInitialize
- CDown.CCDownFrameB:UpdateDurI2H
- APAGui.ApplyPetTargetHUDLayout
- ChattyCathy.CancelOptions
- CleanCastbar.Initialize
- Aura.Aura:UpdateTimerWindow
- CastSequence.BuildSequence
- BustedGUI.ClearAlertFlash
- Enemy.EnemyEffectsIndicator:Update
- CDown.CCDownFrameB:UpdateDurI1H
- DPSMeter.local.UpdateOverallTab
- TortallDPSMeter.OnMenuGlobal
- CallingSetup.UpdateBindingTexts
- DAoCTooltips.CreateCondenseTooltip
- NoOverheal.OnInitialize
- NoOverheal.CheckOverheal
- RoR_SoR.SET_PAIRINGS
- BlackBox.OnApplicationTwoButtonDialogHook
- SOROptions.Slide
- TalismanGenie.EssenceToggle
- Calling.ShowNotification
- TTitan.UI.SetDataSource
- WhatsCooking.CombatCheck
- BuddyBind.GrabName
- mmw.CreateMassMailWin
- Sequencer.SetReturn
- ScenarioStats.LoadWindow
- DPSMeter.local.UpdateAbilityTab
- TTitan.UI.DisplayManualEntry
- AbHelp.Notify
- ScenarioStats.UpdateWindow
- TacticSetNames.SetLabels
- TortallDPSMeter.Initialize
- WARCommander.Init
- XpStatus.InitializeQuotaWindow
- Enemy.Stopwatch_Update
- HealGrid.HealGridCastBar:BeginCast
- RandomMountUI.OnAddCustomMount
- CaVESWindow.Initialize
- StatdollWnd.setLabels
- fpsbox.displayFPS
- CoolDownLine.OnUpdate
- UiModWindow.ShowModCategory
- IHYTM.HandleSlashCmd
- Aura.Aura:CreateRuntimeWindows
- PotionBar.UpdateCooldown
- Enemy.Timer_Update
- Enemy.TimerInitialize
- CastSequence.local.LoadSequence
- SOROptions.Initialize
- Obsidian.Setup.Castbar.OnTimerOffsetXChanged
- DPSMeter.UpdateAbilityTab
- DAoCTooltips.UpdateCondenseTooltip
- GuardRange.UpdateStateMachine
- QuickWarReport.local.ShowConfirmWindow
- DAoCBuff.DAoCBuffFrame:UpdateI2
- CastSequence.Setup.Initialize
- Atlas.local.ShowCoordinatesOnMouseOver
- Enemy.AssistUI_Target_Show
- QuickWarReport.local.PrepareConfirmWindowChrome
- zmm.UpdateListRow
- CDown.local.SetLabels
- Rangechecker.friendly
- TalismanGenie.FragmentOn
- EA_Window_Macro.Initialize
- PP.UpdateDyeFilter
- MapPin.local.EditMarker
- Sequencer.Load
- RoR_SoR.SET_CAMPAIGN
- Obsidian.Setup.EffectTracker.LoadSettings
- TalismanGenie.Search
- TortallDPSMeter.CycleScale
- Atlas.Configuration.Initialize
- CMapWindow.UpdateInfluenceBar
- yAssist.WriteLabels
- KwestorTracker.DrawQuest
- PotionBarSettings.OnAboutShown
- XpStatus.DisplayStatistic
- alertMod.SetSliders
- TortallDPSMeter.DamageUpdate
- DAoCBuffSettings.SetLabels
- Obsidian.Setup.EffectTracker.OnSlideTimerScale
- UiModWindow.Initialize
- MoraleCircle.OnSetCustomColorFill
- QuickWarReport.TestConfirmWindow
- DammazKron.local.InitializeNFO
- snt_castbar.start_interact
- PartyCast.Update
- RoR_SoR.SET_CITY
- SNT_CASTBAR.cleanup
- MoraleCircle.ColorChanger2
- ActionPoints.UpdateCurrentActionPoints
- Atlas.local.UpdateZoneName
- GetStats.OnChatLogUpdated
- Squared.InitImportExport
- TacticSetNames.SettingsWindow.OpenColorDialog
- AbilityAlert.Display
- Obsidian.Setup.EffectTracker.Initialize
- Twister.Initialize
- zmm.UpdateLayout
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- Enemy.IntercomInitialize
- KeyBar.OnLButtonUp
- RealmStatus.OnHistoryWindowUpdate
- Obsidian.Castbar:SetLayout
- DuffTimer.CreateBar
- MapPin.RButtonUp
- Trakario.SCENARIO_END
- BankArkel.PackHide
- BuddyBind.init
- LootAlert.OnInitialize
- PeaceOut.Init
- TalismanGenie.Reset
- TTitan.UI.Initialize
- XpStatus.InitializeSetOpacityWindow
- Enemy.UI_ConfigDialog_Open
- Calling.SetListLine
- Obsidian.Setup.Castbar.OnTimerOffsetYChanged
- Rangechecker.OnTargetUpdated
- ResHelp.OnInitialize
- WarBoard_TDPS.Callback
- wbLeadHelper.showNormalTitle
- Enemy.Guard_GuardIndicator_Update
- Trakario.BRUpdateTime
- DammazKron.local.InitializeTOK
- RvRStatsRvRTab.Init
- CastSequence.SequenceBuilder.Initialize
- TalismanGenie.TalismanStats
- mech.slt
- NoOverheal.OnPlayerEndCast
- CDown.SetLabels
- Trakario.RKOTHUpdateTime
- WSCT.ColorOnInitialize
- APAGui.UpdateInstantOnlyHUD
- MoraleCircle.ColorChanger4
- Effigy.RegisterStateInfoForPlayer
- ScenarioStats.CallSettingsWindow
- SessionRPs.CreateWindow
- BustedGUI.Initialize
- DuffTimer.Options.GetEditBoxText
- TalismanGenie.DisplayData
- TurretScrap.CheckBuffs
- WARCommander.HideInfoTitle
- BarText_Influence.PlayerInfluenceUpdated
- TalismanGenie.CurioOff
- TortallDPSMeter.OnMenuUser2
- WarBoard_Session.Initialize
- Enemy.CombatLogUI_EpsWindow_Initialize
- MiracleGrowLight.setMode
- SORPager.HidePopup
- Enemy.IntercomUI_IntercomDialog_Open
- alertMod.SaveChanges
- Aura.Aura:CreateEditWindows
- Rangechecker.WriteLabel
- TalismanGenie.GoldDustOn
- ThankTheTank.OnCombatAction
- CallingSetup.ShowTutorialPart
- GuildWardenWin.WinSetup
- mech.slc
- mmo.UpdateQueue
- TortallDPSMeter.OnMenuUser1
- AtlasMap.OnMouseOverEnd
- Calling.ManageNotification
- MyReasons.OnUpdate
- NoOverheal.OnMouseOverEnd
- NoOverheal.OnPlayerBeginCast
- Amethyst.Recreate
- TalismanGenie.FragmentToggle
- AggroMeter.Initialize
- Enemy.CombatLogUI_StatsWindow_Open
- Wikki's Cooldown Bar.WCDBCooldownBar:UpdateCooldown
- DuffTimer.SetBarLabel
- TalismanGenie.FragmentOff
- DAoCBuff.ShowMessageWindow
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- Obsidian.SetTimer
- Obsidian.Setup.Castbar.Initialize
- CDown.CCDownFrameB:SetCD
- MiracleGrow.updatePlant
- TalismanGenie.EssenceOff
- MapPin.EditMarker
- CDown.ShowTab
- APAGui.UpdateFollowTargetHUD
- ClosetGoblinCharacterWindow.OnInitialize
- ChattyCathy.CC_ModelWindow
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- MoraleCircle.OnSetCustomColorEmpty
- Obsidian.Setup.Castbar.OnSlideTimerAlpha
- Sequencer.local.SetReturn
- Atlas.ShowCoordinatesOnMouseOver
- ClosetGoblinZoneWindow.RefreshOption
- Phantom.Initialize
- DuffTimer.WindowSetup
- SOROptions.settimer
- Rangechecker.Hostile
- CCM.OnInitialize
- CMapWindow.UpdateCoordinatesWMap
- NoOverheal.OnCancelSpell
- Phantom.PopulateWindow
- CastSequence.SequenceBuilder.OnResetTimeoutLUp
- RvRContribution.OnInitialize
- ShowHealth.TargetUpdated
- SocialWindow.ToggleOfflineMembersShowing
- APAGui.UpdatePetTargetHUD
- SocialWindow.ShowIgnoreListContextWindow
- TastyButtonsOptions.BlockButton
- Swift Assist.local.SetSmartLabel
- TacticSetNames.local.SetLabels
- TalismanGenie.GoldDustToggle
- WARCommander.SetAssistTarget
- WarBoard_TDPS.Initialize
- Killer.ShowPersonalStatsTooltip
- PeaceOut.UpdateLabel
- DAoCBuff.DAoCBuffFrame:UpdateI_hpte
- SocialWindow.ShowFriendsListContextWindow
- Swift Assist.WriteLabels
- Tortall_DPS.local.switchTab
- WarBoard_Loc.OnPlayerPositionUpdated
- TalismanMonitor.UpdateIconLabel
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- MoraleCircle.ColorChanger3
- PotionBar.local.UpdateCooldown
- PartyCast.FetchedText
- CastSequence.LoadSequence
- RoR_SoR.OnInitialize
- TalismanGenie.EssenceOn
- FastFriends.UI.CreateWindows
- AuctionStats.AddExtraWindow
- BloodyMess.OnInitialize
- ChattyCathy.UpdateEntrySetup
- Effigy.RegisterStateInfoForCastbar
- fpsbox.local.displayFPS
- CMapWindow.OnAreaNameChange
- SORPager.Known
- DammazKron.InitializeTOK
- ReferWindow.AddPlayerReferBox
- zmm.CreateDummyLabel
- BlackBookWindow.UpdateList
- GuardList.Initialize
- Obsidian.EffectBar:SetLayout
- Obsidian.Setup.EffectTracker.OnSlideTimerAlpha
- RVMOD_SquaredDistances.OnSquaredSetName
- Swift Assist.local.WriteLabels
- mml.Initialize
- ChattyCathy.ApplyOptions
- DaemonAssist.UpdateWindow
- DammazKron.UpdateSpecial
- mmo.Initialize
- ChattyCathy.OnSlideAlpha
- RoR_SoR.SET_KEEP
- CDown.CCDownFrameB:UpdateDurI2V
- SwiftAssist.aaLabelColorSet
- VPBreakdown.UpdateDetailLabels
- QuickWarReport.ShowConfirmWindow
- Obsidian.OnLoadingEnd
- DammazKron.InitializeNFO
- Obsidian.Setup.Castbar.OnSlideTimerDecimals
- TalismanGenie.CurioToggle
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- RandomMountUI.OnDropSlotLButtonUp
- BlackBox.SetHook
- HealGridUtility.UpdateWARCastBarVisibility
- DAoCBuff.DAoCBuffFrame:SetBuff
- Calling.ClearListLine
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- SORPager.OnMouseOver
- Sequencer.Save
- TidyRollOptions.Initialize
- Enemy.CombatLogUI_EpsWindow_Update
- PeaceOut.PO_Tick
- alertMod.SetLabels
- Obsidian.local.SetTimer
- CDown.local.ShowTab
- TTitan.UI.DisplayRow
- Killer.ShowTopKillersTooltip
- MoraleCircle.OnSetCustomColorFull
- CastSequence.local.BuildSequence
- TortallDPSMeter.ToggleGroup
- Killer.ShowRowTooltip
- Obsidian.EffectBar:SetScale
- CDown.CCDownFrameB:Create
- TalismanGenie.Initialise
- GuardList.UpdateStateMachine
- Enemy.CombatLogUI_StatsWindow_ListRowMouseOver
- SNT_CASTBAR.local.cleanup
- Trakario.Test
- BlackBookWindow.Initialize
- BustedGUI.UpdateErrorView
- CoolDownLine.Initialize
- UiModWindow.InitAdvancedWindow
- DuffTimer.Options.OnSliderChange
- RoR_debolster.check
- TalismanGenie.GoldDustOff
- APAGui.OnShown
- BustedGUI.NewErrorRecorded
- Obsidian.Setup.EffectTracker.OnTimerOffsetYChanged
- CCTV.UpdateSettings
- WARCommander.ShowInfoTitle
- BlackBox.UpdateTimer
- CMapWindow.UpdateCoordinatesMap
- NoOverheal.OnMouseOver
- Aura.Aura:UpdateTimerWindowTime
- Obsidian.Setup.Castbar.OnSlideTimerScale
- MiracleGrow2.UpdatePlot
- Dye.Initialize
- GuardRange.Initialize
- BuffHead.BuffHeadEffectFrame:SetLayout
- FlagCap.OnPointUpdate
- Calling.ManageList
- Obsidian.Castbar:SetScale
- Sequencer.ToggleReset
- Sequencer.OnInitialize
- BuddyBind.OnRawDeviceInput
- CallingSetup.Initialize
- UiModWindow.UpdateInstructions
- Calling.SetKillsListStrings
- SessionRPs.ResetSortColor
- Swift Assist.SetSmartLabel
- TalismanGenie.CurioOn
- Tortall_DPS.switchTab
- BankArkel.SetCharInfo
- Enemy.CombatLogUI_TargetDefenseWindow_Update
- MoraleCircle.init
- QuickWarReport.PrepareConfirmWindowChrome
- Obsidian.Setup.Castbar.LoadSettings
- WarBoard_Loc.Initialize
- ZonePOP.Update
- UiModVersionMismatchWindow.UpdateInstructions
- HealGrid.HealGridHUDMain:Initialize
- DuffTimer.AddNewEffectByWindow


## Binding Resolution

- Total handler declarations: 593
- Resolved to Lua functions: 593 (100%)

## .mod Lifecycle: Startup Windows

This element type is instantiated as a startup window by the following .mod addon(s):

| Frame Name | Addon | Hook | Resolution | Confidence |
| --- | --- | --- | --- | --- |
| BarText_InfluenceText | BarText (Influence) | OnInitialize | exact | HIGH |
## Seen In

- AbilityAlert
- AbilityNotifier
- ActionPoints
- AdjustTheTip
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- ArmorGraphicToggle
- Asshat
- Assist
- Atlas
- AuctionStats
- Aura
- AutoBand
- AutoSalvage
- BankArkel
- BarText (Influence)
- BlackBook
- BlackBox
- Bloody Mess
- Brizio's Crappy Computer Medic
- BuddyBind
- BuffHead
- Busted
- CCTV
- CDown
- CM_ClosetGoblin
- CMap
- CaVES
- Calling
- CastSequence
- ChattyCathy
- Cheeseboard
- CleanCastbar
- CleanUnitFrames
- Clock
- CombatTextNames
- CoolDownLine
- Countdown
- Crafting Info Tooltip
- CraftingWillard
- Crusher
- DAoCBuff
- DPSMeter
- DaemonAssist
- DammazKron
- Dascore
- Deathblow
- Deathblow2
- DeepSleep
- DetauntHelper
- Duel
- DuffTimer
- Dye Preview
- EA_LoadingScreen
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EA_UiModWindow
- EZCraft
- EZCraftX
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- FastFriends
- FastInteract
- FlagCap
- FozAuction
- GCDTracker
- GDes
- Ges
- GetStats
- Group Icons SG
- GroupRange
- GroupSpotter
- GuardLine
- GuardList
- GuardRange
- GuildWarden
- HealGrid
- Hopper
- I HATE YOU THIS MUCH
- InfoScroller
- ItemRack
- JunkDump
- KeyBar
- Keyset
- KeysetMonsterPlay
- KillStreak
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- LibGroup
- LibWBToggler
- LootAlert
- LoyalPet
- Map
- MapMonster
- MapPin
- MarkBuff
- Mass Refine
- Mech
- MegaphonePlusPlus
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- MoraleSet
- Moth
- Motion
- MyReasons
- NaturalLog
- NerfedButtons
- NoOverheal
- Obsidian
- OverheadFonts
- PartyAd
- PartyCast
- PeaceOut
- Phantom
- PieTracker
- Pocket Palette
- PotionBar
- Pure
- Pure Careerbar
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RO-Style Combat Text
- RPs
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RVMOD_Targets
- RaidMeter
- RandomMount
- Rangechecker
- RealmStatus
- Refer
- ReliquaryHunter
- RememberIt
- Res
- ResHelp
- RoR_SoR
- RoR_debolster
- Rolodex
- Rotation
- RvRContribution
- RvRStats
- RvRStatsTab
- SNT_CASTBAR
- SNT_PANEL
- SOR
- ScenarioStats
- Sequencer
- SessionRPs
- Shinies
- ShowHealthPercent
- SimpleCombatText
- SimpleXY
- SocialWindow 2.0
- Soloq
- SpamBayes
- Squared
- SquaredClick
- Statdoll
- Statdoll Light
- Statdoll Remix
- Swift Assist
- Swinger
- TacticSetNames
- TalismanGenie
- TargetInfoRing
- Targets
- TastyButtons
- TaxPayer
- TexturedButtons
- ThankTheTank
- ThinkOutLoud
- TidyChat
- TidyQueue
- TidyRoll
- TokenMachine
- Tome Titan
- TomeTracker
- Tortall_DPS
- Trakario
- TurretRange
- TurretScrap
- Twister
- TwisterSet
- VPBreakdown
- Vectors
- WARCommander
- WARRatingBuster
- WBStutterLess
- WSCT
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
- Warbuilder
- Wargames
- WhatsCooking
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- XpStatus+G
- ZonePOP
- alertMod
- bigger_MacroWindow
- compass
- emotes
- fpsbox
- minesweep
- nLootLink
- talisman-monitor
- wbLeadHelper
- yAssistHelper
- zMailMod

## Examples

- AbilityAlert: AAText -> Label AAText
- AbilityNotifier: AbHelpText -> Label AbHelpText
- ActionPoints: ActionPointsWindowText -> Label ActionPointsWindowText
- AdjustTheTip: AdjustTheTipMenuItemSliderSliderLabel -> Label AdjustTheTipMenuItemSliderSliderLabel
- AdjustTheTip: AdjustTheTipMenuItemSliderSliderValueLabel -> Label AdjustTheTipMenuItemSliderSliderValueLabel
- AdjustTheTip: MouseOverTargetHealthContainerTemplateHealthText -> Label MouseOverTargetHealthContainerTemplateHealthText

## Related APIs

- [Anchor](element_Anchor.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [Color](element_Color.md) (HIGH 100/100) - XML Element Type
- [EA_CheckButtonLabel](../../globals/constants/constant_EA_CheckButtonLabel.md) (HIGH 100/100) - Constant
- [EA_Label_ChatText](../../globals/constants/constant_EA_Label_ChatText.md) (HIGH 100/100) - Constant
- [EA_Label_DefaultSubHeading](../../globals/constants/constant_EA_Label_DefaultSubHeading.md) (HIGH 100/100) - Constant
- [EA_Label_DefaultText](../../globals/constants/constant_EA_Label_DefaultText.md) (HIGH 100/100) - Constant
- [EA_Label_DefaultText_Small](../../globals/constants/constant_EA_Label_DefaultText_Small.md) (HIGH 100/100) - Constant
- [EA_Settings_ItemTitle](../../globals/constants/constant_EA_Settings_ItemTitle.md) (HIGH 100/100) - Constant
- [EA_Settings_SectionTitle](../../globals/constants/constant_EA_Settings_SectionTitle.md) (HIGH 100/100) - Constant
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Sounds](element_Sounds.md) (HIGH 100/100) - XML Element Type
- [Text](element_Text.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EA_Label_SocialListItem](../../globals/constants/constant_EA_Label_SocialListItem.md) (HIGH 90/100) - Constant
- [EA_Template_AutoRollRarityHeader](../../globals/constants/constant_EA_Template_AutoRollRarityHeader.md) (HIGH 90/100) - Constant
- [EA_Template_AutoRollTitle](../../globals/constants/constant_EA_Template_AutoRollTitle.md) (HIGH 90/100) - Constant
- [EA_Template_OpenParty_Label](../../globals/constants/constant_EA_Template_OpenParty_Label.md) (HIGH 90/100) - Constant
- [EA_Template_OpenParty_Label_Small](../../globals/constants/constant_EA_Template_OpenParty_Label_Small.md) (HIGH 90/100) - Constant
- [EA_Window_HelpDescriptionTemplate](../../globals/constants/constant_EA_Window_HelpDescriptionTemplate.md) (HIGH 90/100) - Constant
- [color](element_color.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnMouseLeave](../handlers/handler_OnMouseLeave.md) (HIGH 76/100) - XML Event
