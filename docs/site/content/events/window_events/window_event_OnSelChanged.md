# OnSelChanged

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 138

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BankArkel, BuffHead, DAoCBuff, Enemy, Killer |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/BankArkel/BankArkel.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupContainer.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupDisplay.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupGeneral.xml:0` |
| Namespaces detected | OnSelChanged |
| Source kinds | event_page, examples, xml_handlers |
| Example locations | AdvancedPetAssist: APAComboAttackBind.OnSelChanged, AdvancedPetAssist: APAComboAutoReattack.OnSelChanged, AdvancedPetAssist: APAComboAutoReattackDelay.OnSelChanged, AdvancedPetAssist: APAComboCastDelay.OnSelChanged, AdvancedPetAssist: APAComboCastOnAcquire.OnSelChanged, AdvancedPetAssist: APAComboCombatExitDelay.OnSelChanged |
| XML usage count | 179 |
| XML attribute usage count | 179 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
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

Observed as an engine-supplied UI event hook used by 17 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BankArkel
- BuffHead
- DAoCBuff
- Enemy
- Killer
- Pocket Palette
- PotionBar
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WhoHealedMe

## Registrars And Handlers

- APAGui.OnComboChanged
- AdvancedRenownTraining.SelectedItemChanged
- AuraConfig.OnActivationSoundComboChanged
- AuraConfig.OnDeactivationSoundComboChanged
- AuraConfig.OnTriggerTypeSelChanged
- BankArkel.PackCombo
- BuffHead.Setup.AdvancedContainersItem.OnPositionChanged
- BuffHead.Setup.AdvancedContainersItem.OnTargetTypeChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsBuffsChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsDebuffsChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnElementChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnGrowthHorizontalChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnGrowthVerticalChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnLayerChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnLayoutChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnMaximumThresholdChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnMinimumThresholdChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPositionAnchorChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPositionPlacementChanged
- BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertiesChanged
- BuffHead.Setup.Container.OnContainerAlwaysShowAnchorChanged
- BuffHead.Setup.Container.OnContainerAlwaysShowPlacementChanged
- BuffHead.Setup.Container.OnContainerBuffsAnchorChanged
- BuffHead.Setup.Container.OnContainerBuffsPlacementChanged
- BuffHead.Setup.Container.OnContainerDebuffsAnchorChanged
- BuffHead.Setup.Container.OnContainerDebuffsPlacementChanged
- BuffHead.Setup.Container.OnSizeContainerTypeChanged
- BuffHead.Setup.Display.OnLayerChanged
- BuffHead.Setup.Display.OnSortByChanged
- BuffHead.Setup.Display.OnSortDirectionChanged
- BuffHead.Setup.General.OnCompressionChanged
- BuffHead.Setup.Layout.Manager.OnExportLayoutChanged
- BuffHead.Setup.Layout.Manager.OnLayoutsLayoutChanged
- BuffHead.Setup.Layout.Properties.OnDurationFormatChanged
- BuffHead.Setup.Layout.Properties.OnElementChanged
- BuffHead.Setup.Layout.Properties.OnFontAlignmentChanged
- BuffHead.Setup.Layout.Properties.OnIconBorderColorTypeChanged
- BuffHead.Setup.Layout.Properties.OnLayerLayerChanged
- BuffHead.Setup.Layout.Properties.OnPropertiesChanged
- BuffHead.Setup.Layout.Properties.OnStatusBarForegroundColorTypeChanged
- BuffHead.Setup.Layout.Properties.OnStatusBarOrientationChanged
- BuffHead.Setup.Performance.OnEffectAnchoringChanged
- BuffHead.Setup.PriorityEffects.OnAnimationChanged
- BuffHead.Setup.Trackers.OnTrackerBuffChanged
- BuffHead.Setup.Trackers.OnTrackerDebuffChanged
- BuffHead.Setup.Trackers.OnTrackerTypeChanged
- DAoCBuffSettings.BufforderChanged
- DAoCBuffSettings.CountChanged
- DAoCBuffSettings.DivideChanged
- DAoCBuffSettings.FilterSettings.ClassTableChanged
- DAoCBuffSettings.FilterSettings.ConditionChanged
- DAoCBuffSettings.FilterSettings.ConditionTypeChanged
- DAoCBuffSettings.FilterSettings.FilterPropertyChanged
- DAoCBuffSettings.FilterSettings.G1FilterChanged
- DAoCBuffSettings.FilterSettings.G2ListChanged
- DAoCBuffSettings.FilterSettings.G4HistoryBrowserChanged
- DAoCBuffSettings.FilterSettings.G4UseandChanged
- DAoCBuffSettings.FilterSettings.G5DurationChanged
- DAoCBuffSettings.FilterSettings.G5FilterChanged
- DAoCBuffSettings.FilterSettings.TextureChanged
- DAoCBuffSettings.FilterSettings.TextureTypeChanged
- DAoCBuffSettings.FilterSettings.UseandChanged
- DAoCBuffSettings.FontChanged
- DAoCBuffSettings.GrowHorizontalChanged
- DAoCBuffSettings.GrowLeftChanged
- DAoCBuffSettings.GrowUpChanged
- DAoCBuffSettings.LeftListChanged
- DAoCBuffSettings.ManagerModeChanged
- DAoCBuffSettings.PermabuffsChanged
- DAoCBuffSettings.RefreshChanged
- DAoCBuffSettings.RemoveListChanged
- DAoCBuffSettings.RightListChanged
- DAoCBuffSettings.RowChanged
- DAoCBuffSettings.StickChanged
- DAoCBuffSettings.TargetChanged
- DAoCBuffSettings.TypeChanged
- Enemy.AssistUI_ConfigDialog_OnNewTargetSoundIdSelChanged
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- Enemy.CombatLogUI_StatsWindow_OnTypeSelChanged
- Enemy.ConfigurationWindow_OnChange
- Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- Enemy.IntercomUI_ChooseChannelDialog_ChannelListChanged
- Enemy.UI_ConfigDialog_OnSectionSelChanged
- Enemy.UnitFramesUI_ConfigDialog_OnSorting1Changed
- Enemy.UnitFramesUI_ConfigDialog_OnSorting2Changed
- Enemy.UnitFramesUI_ConfigDialog_OnSorting3Changed
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- Enemy.UnitFramesUI_UnitFramePartDialog_OnTypeSelChanged
- Killer.OnSettingsComboChanged
- PP.UpdateDyeList
- PotionBarSettings.ActivatorComboSelChanged
- PotionBarSettings.BuildComboSelChanged
- PotionBarSettings.ComboMethod
- PotionBarSettings.InfoTextBRComboSelChanged
- PotionBarSettings.InfoTextTRComboSelChanged
- PotionBarSettings.QuickActionsSelChanged
- ShiniesBrowseUI.OnSelChanged_Criteria_ItemSlotCombo
- ShiniesBrowseUI.OnSelChanged_Criteria_ItemTypeCombo
- ShiniesBrowseUI.OnSelChanged_Criteria_ModifierCombo
- TexturedButtons.Setup.Actionbar.OnBarChanged
- TexturedButtons.Setup.Actionbar.OnSelectorChanged
- TexturedButtons.Setup.AdvancedTextures.OnCustomChanged
- TexturedButtons.Setup.AdvancedTextures.OnCustomTextureChanged
- TexturedButtons.Setup.AdvancedTextures.OnCustomTextureSliceChanged
- TexturedButtons.Setup.AdvancedTextures.OnPresetChanged
- TexturedButtons.Setup.AdvancedTextures.OnSlotTypeChanged
- TexturedButtons.Setup.Misc.OnActionButtonPickUpModifierChanged
- TexturedButtons.Setup.Misc.OnCustomGlowChanged
- TexturedButtons.Setup.Textures.OnCustomChanged
- TexturedButtons.Setup.Textures.OnCustomTextureChanged
- TexturedButtons.Setup.Textures.OnCustomTextureSliceChanged
- TexturedButtons.Setup.Textures.OnPresetChanged
- TexturedButtons.Setup.Tint.OnTintTypeChanged
- TidyChat.Options.UpdateGroupTabs
- TidyRoll.CustomAutoRoll.OnChoiceChange
- TurretRange.Setup.Display.OnCircleModeChanged
- TurretRange.Setup.Display.OnCircleTypeChanged
- TurretRange.Setup.Display.OnDistanceLayoutChanged
- TurretRange.Setup.Display.OnDistanceTypeChanged
- TurretRange.Setup.Display.OnElementChanged
- TurretRange.Setup.Display.OnGraphicTypeChanged
- WHMGui.OnOptionsComboChanged
- WSCT.ComboOnSelChanged
- WSCT.EventOnSelChanged
- WSCT.FrameComboOnSelChanged
- WSCT.FrameOnSelChanged

## Examples

- AdvancedPetAssist: APAComboAttackBind -> APAComboAttackBind.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboAutoReattack -> APAComboAutoReattack.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboAutoReattackDelay -> APAComboAutoReattackDelay.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboCastDelay -> APAComboCastDelay.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboCastOnAcquire -> APAComboCastOnAcquire.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboCombatExitDelay -> APAComboCombatExitDelay.OnSelChanged -> APAGui.OnComboChanged

## Related APIs

- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function

## Used With

- [ComboBox](../../xml/element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnSelChanged](../../xml/handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Event

## Triggered By

- none

## Affects

- none

## Notes

- none
