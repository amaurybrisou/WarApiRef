# ComboBox

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BankArkel, BuffHead, DAoCBuff, Enemy, Killer |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/BankArkel/BankArkel.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItemEffect.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupContainer.xml:0` |
| Namespaces detected | ComboBox |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedPetAssist: APAComboAttackBind, AdvancedPetAssist: APAComboAutoReattack, AdvancedPetAssist: APAComboAutoReattackDelay, AdvancedPetAssist: APAComboCastDelay, AdvancedPetAssist: APAComboCastOnAcquire, AdvancedPetAssist: APAComboCombatExitDelay |
| XML usage count | 238 |
| XML attribute usage count | 238 |
| Lua usage count | 2 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
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

Observed XML element type instantiated by 18 addons.

## Common Attributes

- autoresize
- inherits
- layer
- maxvisibleitems
- menubackground
- menuitembutton
- name
- scrollbar
- selectedbutton
- show

## Common Handlers

- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnSelChanged](../handlers/handler_OnSelChanged.md)

## Common Handler Functions

- APAGui.OnComboChanged
- WHMGui.OnOptionsComboChanged
- Enemy.ConfigurationWindow_OnChange
- Enemy.ConfigurationWindow_ShowTooltip
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Killer.OnSettingsComboChanged
- PP.UpdateDyeList
- AdvancedRenownTraining.SelectedItemChanged
- AuraConfig.OnActivationSoundComboChanged
- AuraConfig.OnDeactivationSoundComboChanged
- AuraConfig.OnTriggerTypeSelChanged
- BankArkel.PackCombo


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.ConfigurationWindow_ShowTooltip, PotionBarSettings.OnMouseoverActivator, PotionBarSettings.OnMouseoverBuild, PotionBarSettings.OnMouseoverInfoTextBR, PotionBarSettings.OnMouseoverInfoTextTR, PotionBarSettings.OnMouseoverMethod | `function()` | MEDIUM |
| [OnSelChanged](../handlers/handler_OnSelChanged.md) | data | APAGui.OnComboChanged, WHMGui.OnOptionsComboChanged, Enemy.ConfigurationWindow_OnChange, Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample, Killer.OnSettingsComboChanged, PP.UpdateDyeList | `index` | MEDIUM |
| [OnSelChanged](../handlers/handler_OnSelChanged.md) | data | APAGui.OnComboChanged, WHMGui.OnOptionsComboChanged, Enemy.ConfigurationWindow_OnChange, Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample, Killer.OnSettingsComboChanged, PP.UpdateDyeList, AdvancedRenownTraining.SelectedItemChanged, AuraConfig.OnActivationSoundComboChanged, AuraConfig.OnDeactivationSoundComboChanged, AuraConfig.OnTriggerTypeSelChanged, BankArkel.PackCombo, BuffHead.Setup.AdvancedContainersItem.OnPositionChanged, BuffHead.Setup.AdvancedContainersItem.OnTargetTypeChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsBuffsChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsDebuffsChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnElementChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnGrowthHorizontalChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnGrowthVerticalChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnLayerChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnLayoutChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnMaximumThresholdChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnMinimumThresholdChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnPositionAnchorChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnPositionPlacementChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertiesChanged, BuffHead.Setup.Container.OnContainerAlwaysShowAnchorChanged, BuffHead.Setup.Container.OnContainerAlwaysShowPlacementChanged, BuffHead.Setup.Container.OnContainerBuffsAnchorChanged, BuffHead.Setup.Container.OnContainerBuffsPlacementChanged, BuffHead.Setup.Container.OnContainerDebuffsAnchorChanged, BuffHead.Setup.Container.OnContainerDebuffsPlacementChanged, BuffHead.Setup.Container.OnSizeContainerTypeChanged, BuffHead.Setup.Display.OnLayerChanged, BuffHead.Setup.Display.OnSortByChanged, BuffHead.Setup.Display.OnSortDirectionChanged, BuffHead.Setup.General.OnCompressionChanged, BuffHead.Setup.Layout.Manager.OnExportLayoutChanged, BuffHead.Setup.Layout.Manager.OnLayoutsLayoutChanged, BuffHead.Setup.Layout.Properties.OnDurationFormatChanged, BuffHead.Setup.Layout.Properties.OnElementChanged, BuffHead.Setup.Layout.Properties.OnFontAlignmentChanged, BuffHead.Setup.Layout.Properties.OnIconBorderColorTypeChanged, BuffHead.Setup.Layout.Properties.OnLayerLayerChanged, BuffHead.Setup.Layout.Properties.OnPropertiesChanged, BuffHead.Setup.Layout.Properties.OnStatusBarForegroundColorTypeChanged, BuffHead.Setup.Layout.Properties.OnStatusBarOrientationChanged, BuffHead.Setup.Performance.OnEffectAnchoringChanged, BuffHead.Setup.PriorityEffects.OnAnimationChanged, BuffHead.Setup.Trackers.OnTrackerBuffChanged, BuffHead.Setup.Trackers.OnTrackerDebuffChanged, BuffHead.Setup.Trackers.OnTrackerTypeChanged, DAoCBuffSettings.BufforderChanged, DAoCBuffSettings.CountChanged, DAoCBuffSettings.DivideChanged, DAoCBuffSettings.FilterSettings.ClassTableChanged, DAoCBuffSettings.FilterSettings.ConditionChanged, DAoCBuffSettings.FilterSettings.ConditionTypeChanged, DAoCBuffSettings.FilterSettings.FilterPropertyChanged, DAoCBuffSettings.FilterSettings.G1FilterChanged, DAoCBuffSettings.FilterSettings.G2ListChanged, DAoCBuffSettings.FilterSettings.G4HistoryBrowserChanged, DAoCBuffSettings.FilterSettings.G4UseandChanged, DAoCBuffSettings.FilterSettings.G5DurationChanged, DAoCBuffSettings.FilterSettings.G5FilterChanged, DAoCBuffSettings.FilterSettings.TextureChanged, DAoCBuffSettings.FilterSettings.TextureTypeChanged, DAoCBuffSettings.FilterSettings.UseandChanged, DAoCBuffSettings.FontChanged, DAoCBuffSettings.GrowHorizontalChanged, DAoCBuffSettings.GrowLeftChanged, DAoCBuffSettings.GrowUpChanged, DAoCBuffSettings.LeftListChanged, DAoCBuffSettings.ManagerModeChanged, DAoCBuffSettings.PermabuffsChanged, DAoCBuffSettings.RefreshChanged, DAoCBuffSettings.RemoveListChanged, DAoCBuffSettings.RightListChanged, DAoCBuffSettings.RowChanged, DAoCBuffSettings.StickChanged, DAoCBuffSettings.TargetChanged, DAoCBuffSettings.TypeChanged, Enemy.AssistUI_ConfigDialog_OnNewTargetSoundIdSelChanged, Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged, Enemy.CombatLogUI_StatsWindow_OnTypeSelChanged, Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged, Enemy.IntercomUI_ChooseChannelDialog_ChannelListChanged, Enemy.UI_ConfigDialog_OnSectionSelChanged, Enemy.UnitFramesUI_ConfigDialog_OnSorting1Changed, Enemy.UnitFramesUI_ConfigDialog_OnSorting2Changed, Enemy.UnitFramesUI_ConfigDialog_OnSorting3Changed, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged, Enemy.UnitFramesUI_UnitFramePartDialog_OnTypeSelChanged, PotionBarSettings.ActivatorComboSelChanged, PotionBarSettings.BuildComboSelChanged, PotionBarSettings.ComboMethod, PotionBarSettings.InfoTextBRComboSelChanged, PotionBarSettings.InfoTextTRComboSelChanged, PotionBarSettings.QuickActionsSelChanged, ShiniesBrowseUI.OnSelChanged_Criteria_ItemSlotCombo, ShiniesBrowseUI.OnSelChanged_Criteria_ItemTypeCombo, ShiniesBrowseUI.OnSelChanged_Criteria_ModifierCombo, TexturedButtons.Setup.Actionbar.OnBarChanged, TexturedButtons.Setup.Actionbar.OnSelectorChanged, TexturedButtons.Setup.AdvancedTextures.OnCustomChanged, TexturedButtons.Setup.AdvancedTextures.OnCustomTextureChanged, TexturedButtons.Setup.AdvancedTextures.OnCustomTextureSliceChanged, TexturedButtons.Setup.AdvancedTextures.OnPresetChanged, TexturedButtons.Setup.AdvancedTextures.OnSlotTypeChanged, TexturedButtons.Setup.Misc.OnActionButtonPickUpModifierChanged, TexturedButtons.Setup.Misc.OnCustomGlowChanged, TexturedButtons.Setup.Textures.OnCustomChanged, TexturedButtons.Setup.Textures.OnCustomTextureChanged, TexturedButtons.Setup.Textures.OnCustomTextureSliceChanged, TexturedButtons.Setup.Textures.OnPresetChanged, TexturedButtons.Setup.Tint.OnTintTypeChanged, TidyChat.Options.UpdateGroupTabs, TidyRoll.CustomAutoRoll.OnChoiceChange, TurretRange.Setup.Display.OnCircleModeChanged, TurretRange.Setup.Display.OnCircleTypeChanged, TurretRange.Setup.Display.OnDistanceLayoutChanged, TurretRange.Setup.Display.OnDistanceTypeChanged, TurretRange.Setup.Display.OnElementChanged, TurretRange.Setup.Display.OnGraphicTypeChanged, WSCT.ComboOnSelChanged, WSCT.EventOnSelChanged, WSCT.FrameComboOnSelChanged, WSCT.FrameOnSelChanged | `index` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.ConfigurationWindow_ShowTooltip, PotionBarSettings.OnMouseoverActivator, PotionBarSettings.OnMouseoverBuild, PotionBarSettings.OnMouseoverInfoTextBR, PotionBarSettings.OnMouseoverInfoTextTR, PotionBarSettings.OnMouseoverMethod, PotionBarSettings.OnMouseoverQuickActions, WSCT.FrameComboOnMouseOver, WSCT.OnMouseOver | `` |  |

### Per-Event Lua API Calls

**OnMouseOver** handlers call: `WindowGetParent`

**OnSelChanged** handlers call: `ButtonGetPressedFlag`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `ButtonSetText`, `ComboBoxGetSelectedMenuItem`, `ComboBoxGetSelectedText`, `ComboBoxSetDisabledFlag`, `ComboBoxSetSelectedMenuItem`, `CreateWindowFromTemplate`, `DestroyWindow`, `DoesWindowExist`, `DynamicImageSetTexture`, `DynamicImageSetTextureSlice`, `LabelSetText`, `ListBoxSetDisplayOrder`, `ScrollWindowUpdateScrollRect`, `SliderBarGetCurrentPosition`, `TextEditBoxGetText`, `WindowAddAnchor`, `WindowClearAnchors`, `WindowGetId`, `WindowGetParent`, `WindowSetShowing`

**OnSelChanged** handlers call: `ButtonGetPressedFlag`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `ButtonSetText`, `ComboBoxGetSelectedMenuItem`, `ComboBoxGetSelectedText`, `ComboBoxSetDisabledFlag`, `ComboBoxSetSelectedMenuItem`, `CreateWindowFromTemplate`, `DestroyWindow`, `DoesWindowExist`, `DynamicImageSetTexture`, `DynamicImageSetTextureSlice`, `LabelSetText`, `ListBoxSetDisplayOrder`, `ScrollWindowUpdateScrollRect`, `SliderBarGetCurrentPosition`, `TextEditBoxGetText`, `WindowAddAnchor`, `WindowClearAnchors`, `WindowGetId`, `WindowGetParent`, `WindowSetShowing`

**OnMouseOver** handlers call: `WindowGetParent`

## Common Inherits

- APA_ComboBox
- APA_ComboBoxWide
- Aura_ComboBox_DefaultResizable
- Aura_ComboBox_DefaultResizableLarge
- Aura_ComboBox_DefaultResizableTiny
- EA_ComboBox_DefaultResizable
- EA_ComboBox_DefaultResizableLarge
- EA_ComboBox_DefaultResizableMedium
- EA_ComboBox_DefaultResizableSmall
- EA_ComboBox_DefaultResizable_Fixed
- Shinies_ComboBox_DefaultResizableLarge

## Common Parent Elements

- [Windows](element_Windows.md) — 233× (HIGH)
- [Window](element_Window.md) — 4× (MEDIUM)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 228× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 179× (HIGH)
- [Size](element_Size.md) — 52× (HIGH)
- [MenuButtonOffset](element_MenuButtonOffset.md) — 7× (MEDIUM)

## Common Template Bases

- APA_ComboBox
- APA_ComboBoxWide
- Aura_ComboBox_DefaultResizable
- Aura_ComboBox_DefaultResizableLarge
- Aura_ComboBox_DefaultResizableTiny
- EA_ComboBox_DefaultResizable
- EA_ComboBox_DefaultResizableLarge
- EA_ComboBox_DefaultResizableMedium
- EA_ComboBox_DefaultResizableSmall
- EA_ComboBox_DefaultResizable_Fixed
- Shinies_ComboBox_DefaultResizableLarge


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<ComboBox autoresize="true" maxvisibleitems="10" menubackground="EA_Window_ComboBoxMenuBackground" menuitembutton="EA_Button_DefaultMenuButtonSmall" name="..." scrollbar="EA_ScrollBar_DefaultVerticalC..." selectedbutton="APA_ComboBoxButton">
  <Size/>
  <MenuButtonOffset x="5" y="5"/>
</ComboBox>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 97% | APA_ComboBox, APA_ComboBoxWide, EA_ComboBox_DefaultResizable, Aura_ComboBox_DefaultResizableLarge, ... |
| `layer` | optional | 46% | secondary, popup, default |
| `maxvisibleitems` | optional | 7% | 10, 13, 5, 8, ... |
| `show` | optional | 4% | false |
| `menubackground` | optional | 2% | EA_Window_ComboBoxMenuBackground |
| `menuitembutton` | optional | 2% | EA_Button_DefaultMenuButtonSmall, EA_Button_DefaultMenuButton, Aura_Button_DefaultMenuButton, Aura_Button_DefaultMenuButtonLarge, ... |
| `scrollbar` | optional | 2% | EA_ScrollBar_DefaultVerticalChain |
| `selectedbutton` | optional | 2% | APA_ComboBoxButton, APA_ComboBoxButtonWide, Aura_ComboBox_DefaultResizableComboBoxSelected, Aura_ComboBox_DefaultResizableComboBoxSelectedLarge, ... |
| `autoresize` | optional | 2% | true |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 228 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 179 times as an unnamed child.

### [Size](element_Size.md)

Observed 52 times as an unnamed child.

### [MenuButtonOffset](element_MenuButtonOffset.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 5 |
| `y` | **required** | 5 |
## Recursive Hierarchy

- Root: [ComboBox](element_ComboBox.md)
- [Anchors](element_Anchors.md) (structural, 228×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
      - (cycle)
- [EventHandlers](element_EventHandlers.md) (structural, 179×, HIGH)
  - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
- [MenuButtonOffset](element_MenuButtonOffset.md) (structural, 7×, MEDIUM)
- [Size](element_Size.md) (structural, 52×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `ComboBoxGetSelectedMenuItem` | ui | 127 | OnSelChanged |
| `WindowSetShowing` | ui | 39 | OnSelChanged |
| `TextEditBoxGetText` | ui | 20 | OnSelChanged |
| `ComboBoxSetSelectedMenuItem` | ui | 10 | OnSelChanged |
| `LabelSetText` | ui | 10 | OnSelChanged |
| `SliderBarGetCurrentPosition` | ui | 8 | OnSelChanged |
| `ButtonGetPressedFlag` | ui | 6 | OnSelChanged |
| `ButtonSetPressedFlag` | ui | 6 | OnSelChanged |
| `ButtonSetText` | ui | 5 | OnSelChanged |
| `DynamicImageSetTexture` | ui | 5 | OnSelChanged |
| `WindowGetParent` | ui | 4 | OnMouseOver, OnSelChanged |
| `ComboBoxGetSelectedText` | ui | 3 | OnSelChanged |
| `WindowAddAnchor` | ui | 3 | OnSelChanged |
| `ButtonSetDisabledFlag` | ui | 2 | OnSelChanged |
| `ComboBoxSetDisabledFlag` | ui | 2 | OnSelChanged |
| `DoesWindowExist` | ui | 2 | OnSelChanged |
| `DynamicImageSetTextureSlice` | ui | 2 | OnSelChanged |
| `ListBoxSetDisplayOrder` | ui | 2 | OnSelChanged |
| `WindowClearAnchors` | ui | 2 | OnSelChanged |
| `CreateWindowFromTemplate` | ui | 1 | OnSelChanged |
| `DestroyWindow` | ui | 1 | OnSelChanged |
| `ScrollWindowUpdateScrollRect` | ui | 1 | OnSelChanged |
| `WindowGetId` | ui | 1 | OnSelChanged |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnMouseOver

Confidence: HIGH

### OnSelChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `index` | number | selected_index |
## Lua Functions Manipulating This Type

- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.CombatLogUI_StatsWindow_UpdateSessionsList
- Enemy.IntercomUI_ChooseChannelDialog_Open
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- Enemy.CombatLogUI_StatsWindow_Open
- Enemy.UI_ConfigDialog_Open
- Enemy.UnitFramesUI_ConfigDialog_Import
- Enemy.IntercomUI_IntercomJoinDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- AdvancedRenownTrainer.GeneratePresetByLinkData
- BankArkel.SetupCombos
- Enemy.GroupsUI_EffectFilterDialog_Open
- Enemy.UI_ConfigDialog_OnSectionSelChanged
- Enemy.IntercomUI_IntercomJoinDialog_AddGroup
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged


## Binding Resolution

- Total handler declarations: 189
- Resolved to Lua functions: 189 (100%)

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
- WarBoard
- WhoHealedMe

## Examples

- AdvancedPetAssist: APAComboAttackBind -> ComboBox APAComboAttackBind
- AdvancedPetAssist: APAComboAutoReattack -> ComboBox APAComboAutoReattack
- AdvancedPetAssist: APAComboAutoReattackDelay -> ComboBox APAComboAutoReattackDelay
- AdvancedPetAssist: APAComboCastDelay -> ComboBox APAComboCastDelay
- AdvancedPetAssist: APAComboCastOnAcquire -> ComboBox APAComboCastOnAcquire
- AdvancedPetAssist: APAComboCombatExitDelay -> ComboBox APAComboCombatExitDelay

## Related APIs

- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [EA_ComboBox_DefaultResizable](../../globals/constants/constant_EA_ComboBox_DefaultResizable.md) (HIGH 100/100) - Constant
- [EA_ComboBox_DefaultResizableLarge](../../globals/constants/constant_EA_ComboBox_DefaultResizableLarge.md) (HIGH 100/100) - Constant
- [EA_ComboBox_DefaultResizableSmall](../../globals/constants/constant_EA_ComboBox_DefaultResizableSmall.md) (HIGH 100/100) - Constant
- [EA_ComboBox_DefaultResizable_Fixed](../../globals/constants/constant_EA_ComboBox_DefaultResizable_Fixed.md) (HIGH 100/100) - Constant
- [MenuButtonOffset](element_MenuButtonOffset.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EA_ComboBox_DefaultResizableMedium](../../globals/constants/constant_EA_ComboBox_DefaultResizableMedium.md) (HIGH 90/100) - Constant
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 55/100) - XML Element Type

## Used With

- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnSelChanged](../../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event
- [OnSelChanged](../handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnSelChanged](../handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Event

## Affects

- none
