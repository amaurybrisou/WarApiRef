# DynamicImage

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
| Addons seen in | AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Aura, AutoMark, BagOMatic, BankArkel, BlackBook |
| Files seen in | Bars/HealGridProgressBar.xml, Code/Assist/Assist.xml, Code/CombatLog/CombatLogIDS.xml, Code/Core/Common.xml, Code/Core/Icon.xml, Code/GroupIcons/GroupIcons.xml, Code/Guard/GuardDistanceIndicator.xml, Code/Marks/MarkTemplate.xml |
| Namespaces detected | DynamicImage |
| Source kinds | xml_frames |
| Example locations | AdvancedRenownTrainer: AbilityButtonTemplateSquare, AdvancedRenownTrainer: AbilityButtonTemplateSquareFrame, AdvancedRenownTrainer: AdvancedRenownTrainingWindowCornorImage, AggroMeter: AggroMeterWindow_AggroWindow1Tactic, AggroMeter: AggroMeterWindow_AggroWindow2Tactic, AggroMeter: AggroMeterWindow_AggroWindow3Tactic |
| XML usage count | 1294 |
| XML attribute usage count | 1294 |
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

DynamicImage is an interactive XML control. It commonly appears under Button and DynamicImage. It is typically used to organize structural children such as Anchor, Anchors, EventHandlers and bind XML events like OnLButtonDown, OnLButtonUp, OnMouseOver to Lua.

## Common Attributes

- alpha
- handleinput
- id
- inherits
- layer
- name
- popable
- slice
- sticky
- texture
- textureScale
- texturescale

## Common Handlers

- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)

## Common Handler Functions

- DAoCBuffSettings.ShowTooltip
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip
- GuildWarden.WardMouseOver1
- GuildWarden.WardMouseOver2
- GuildWarden.WardMouseOver3
- GuildWarden.WardMouseOver4
- GuildWarden.WardMouseOver5
- Motion.OnMouseOver_ContextMenu_Recipe
- NBSBCommon.OnMouseoverToolTip
- AggroMeter.OnMouseOverStart
- BagOMatic.wnd_on_mouse_over
- CCM.OnMouseOverB1
- CCM.OnMouseOverB2
- CMapWindow.MouseoverMail
- CMapWindow.OnMouseoverRvRIndicator
- CallingIcon.OnMouseOver
- CastSequence.Builder.OnSettingsOnMouseOver
- DetauntAbilityManager.OnMouseover
- Enemy.Guard_GuardIndicator_OnMouseOver
- Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOver
- Enemy.TalismanAlerter_TalismanAlerterIndicator_OnMouseOver
- Enemy.UI_Icon_OnMouseOver
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionMouseOver
- FrameManager.OnMouseOver
- GroupSpotter.OnMouseOverStart
- KeyBar.MouseOver
- MapMonster.OnMouseOverPin
- MapMonster.PinTypeEditor.MouseOverDescription
- MiniMapMonster.OnMouseOverPin
- Minmap.OnMouseoverMailIcon
- NBSBCore.IconOnMouseOver
- NBSBParam.OnAbilityTT
- RVMOD_Manager.OnMouseOverFilterAll
- RVMOD_Manager.OnMouseOverFilterEA
- RVMOD_Manager.OnMouseOverFilterRV
- RVMOD_Manager.OnMouseOverMagnifier
- RVMOD_Manager.OnMouseOverSortBy
- RVMOD_SquaredDistances.OnMouseOverJoystick
- SORRealms.OnMouseOverKeep
- TalismanMonitor.ShowStatus
- TomeTracker.Journal.MapButtonMouseOver
- TomeTracker.Saga.OnMouseOverMapPin
- TomeTracker.Saga.OnMouseOverReward
- Warbuilder.overtooltip
- Vectors.Settings.ArrowClicked
- NBSBCore.OnToolsUp
- BuffHead.Setup.Container.OnContainerClick
- Twister.AbilityCursorSwap
- ClosetGoblinOptionWindow.OnLButtonUp
- MapPin.ShowIcons
- BagOMatic.wnd_on_lbutton_up
- CCM.OnLButtonUpB1
- CCM.OnLButtonUpB2
- CalljoinGUI.ContextSensitiveChannelOperation
- CastSequence.Builder.OnSettingsLClick
- DetauntAbilityManager.OnDrop
- DetauntHelperMonitor.OnEyeLButtonUp
- DetauntTarget.OnEyeLButtonUp
- Dye.Toogle
- Enemy.Guard_GuardIndicator_OnLButtonUp
- Enemy.MarksUI_EnemyMarksWindow_OnAddLButtonUp
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick
- Enemy.UI_Icon_OnLButtonUp
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionLButtonUp
- FrameManager.OnLButtonUp
- GroupSpotter.Settings.ToggleGroupSpotter
- HealGridGui.ToggleGui
- KeyBar.LButtonUp
- MapMonster.OnMouseClickPin
- MapMonster.PinTypeEditor.OnClickChooseIcon
- MiniMapMonster.OnMouseClickPin
- NBSBCore.Show
- NBSBParam.OnParamChange
- RVMOD_Manager.OnLButtonUpFilterAll
- RVMOD_Manager.OnLButtonUpFilterEA
- RVMOD_Manager.OnLButtonUpFilterRV
- RVMOD_Manager.OnLButtonUpMagnifier
- RVMOD_Manager.OnLButtonUpSortBy
- SOROptions.ShowTier
- SORRealms.OnLClickKeep
- SquaredClick.AbilityCursorSwap
- TacticSetNames.SettingsWindow.OpenColorDialog
- TalismanMonitor.ToggleShowAll
- TokenMachine.ToggleSettings
- TomeTracker.Journal.Toggle
- TomeTracker.Saga.HighlightMapPin
- TurretRange.Setup.Display.OnCircleColorLUp
- TurretRange.Setup.Display.OnDistanceColorLUp
- TurretRange.Setup.Display.OnGraphicColorLUp
- Twister.OpenSettingsWindow
- BuffHead.Setup.AdvancedContainersItem.OnContainerRClick
- Twister.AbilityClear
- ClosetGoblinOptionWindow.OnRButtonUp
- CCM.OnRButtonUpB1
- CCM.OnRButtonUpB2
- CallingSetup.ToggleShowing
- DetauntAbilityManager.OnClear
- Enemy.MarksUI_EnemyMarksWindow_OnAddRButtonUp
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick
- Enemy.UI_Icon_OnRButtonUp
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionRButtonUp
- FrameManager.OnRButtonUp
- GroupSpotter.OnRButtonUp
- MapMonster.OnMouseRightClickPin
- MiniMapMonster.OnMouseRightClickPin
- NBSBCore.DoMenu
- PlayEffectsOn.OnRButtonUp
- SORRealms.ShowKeepMenu
- SquaredClickConfig.AbilityClear
- Twister.OnLButtonDown
- NaturalLog.ScrollToBottom
- NaturalLog.ScrollToTop
- DetauntHelperConfig.UpdateConfig
- Enemy.CombatLogUI_IDS_OnRowLButtonDown
- FrameManager.OnLButtonDown
- MBuffGui.OnSpellLClick
- MBuffSetup.SmartBuff.OnOtherLClick
- OnTabLBU
- RVMOD_SquaredDistances.OnLButtonDownJoystick
- Squared.ToggleActive
- Warbuilder.OnTabLBU
- CastSequence.Builder.OnSettingsOnMouseOverEnd
- Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOverEnd
- FrameManager.OnMouseOverEnd
- GroupSpotter.OnMouseOverEnd
- NBSBCore.IconOnMouseOverEnd
- RVMOD_SquaredDistances.OnMouseOverEndJoystick
- SORPager.BOsOff
- TomeTracker.Saga.OnMouseOverMapPinEnd
- TomeTracker.Saga.OnMouseOverRewardEnd
- FrameManager.OnRButtonDown
- MBuffGui.OnSpellRClick
- MBuffSetup.SmartBuff.OnOtherRClick
- Squared.ShowSettings


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | DAoCBuffSettings.ShowTooltip, Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip, GuildWarden.WardMouseOver1, GuildWarden.WardMouseOver2, GuildWarden.WardMouseOver3, GuildWarden.WardMouseOver4, GuildWarden.WardMouseOver5, Motion.OnMouseOver_ContextMenu_Recipe, NBSBCommon.OnMouseoverToolTip, AggroMeter.OnMouseOverStart, BagOMatic.wnd_on_mouse_over, CCM.OnMouseOverB1, CCM.OnMouseOverB2, CMapWindow.MouseoverMail, CMapWindow.OnMouseoverRvRIndicator, CallingIcon.OnMouseOver, CastSequence.Builder.OnSettingsOnMouseOver, DetauntAbilityManager.OnMouseover, Enemy.Guard_GuardIndicator_OnMouseOver, Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOver, Enemy.TalismanAlerter_TalismanAlerterIndicator_OnMouseOver, Enemy.UI_Icon_OnMouseOver, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionMouseOver, FrameManager.OnMouseOver, GroupSpotter.OnMouseOverStart, KeyBar.MouseOver, MapMonster.OnMouseOverPin, MapMonster.PinTypeEditor.MouseOverDescription, MiniMapMonster.OnMouseOverPin, Minmap.OnMouseoverMailIcon, NBSBCore.IconOnMouseOver, NBSBParam.OnAbilityTT, RVMOD_Manager.OnMouseOverFilterAll, RVMOD_Manager.OnMouseOverFilterEA, RVMOD_Manager.OnMouseOverFilterRV, RVMOD_Manager.OnMouseOverMagnifier, RVMOD_Manager.OnMouseOverSortBy, RVMOD_SquaredDistances.OnMouseOverJoystick, SORRealms.OnMouseOverKeep, TalismanMonitor.ShowStatus, TomeTracker.Journal.MapButtonMouseOver, TomeTracker.Saga.OnMouseOverMapPin, TomeTracker.Saga.OnMouseOverReward, Warbuilder.overtooltip | `` |  |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | Vectors.Settings.ArrowClicked, NBSBCore.OnToolsUp, BuffHead.Setup.Container.OnContainerClick, Twister.AbilityCursorSwap, ClosetGoblinOptionWindow.OnLButtonUp, MapPin.ShowIcons, BagOMatic.wnd_on_lbutton_up, CCM.OnLButtonUpB1, CCM.OnLButtonUpB2, CalljoinGUI.ContextSensitiveChannelOperation, CastSequence.Builder.OnSettingsLClick, DetauntAbilityManager.OnDrop, DetauntHelperMonitor.OnEyeLButtonUp, DetauntTarget.OnEyeLButtonUp, Dye.Toogle, Enemy.Guard_GuardIndicator_OnLButtonUp, Enemy.MarksUI_EnemyMarksWindow_OnAddLButtonUp, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick, Enemy.UI_Icon_OnLButtonUp, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionLButtonUp, FrameManager.OnLButtonUp, GroupSpotter.Settings.ToggleGroupSpotter, HealGridGui.ToggleGui, KeyBar.LButtonUp, MapMonster.OnMouseClickPin, MapMonster.PinTypeEditor.OnClickChooseIcon, MiniMapMonster.OnMouseClickPin, NBSBCore.Show, NBSBParam.OnParamChange, RVMOD_Manager.OnLButtonUpFilterAll, RVMOD_Manager.OnLButtonUpFilterEA, RVMOD_Manager.OnLButtonUpFilterRV, RVMOD_Manager.OnLButtonUpMagnifier, RVMOD_Manager.OnLButtonUpSortBy, SOROptions.ShowTier, SORRealms.OnLClickKeep, SquaredClick.AbilityCursorSwap, TacticSetNames.SettingsWindow.OpenColorDialog, TalismanMonitor.ToggleShowAll, TokenMachine.ToggleSettings, TomeTracker.Journal.Toggle, TomeTracker.Saga.HighlightMapPin, TurretRange.Setup.Display.OnCircleColorLUp, TurretRange.Setup.Display.OnDistanceColorLUp, TurretRange.Setup.Display.OnGraphicColorLUp | `flags, x, y` | MEDIUM |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Twister.OpenSettingsWindow, BuffHead.Setup.AdvancedContainersItem.OnContainerRClick, Twister.AbilityClear, ClosetGoblinOptionWindow.OnRButtonUp, CCM.OnRButtonUpB1, CCM.OnRButtonUpB2, CallingSetup.ToggleShowing, DetauntAbilityManager.OnClear, Enemy.MarksUI_EnemyMarksWindow_OnAddRButtonUp, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick, Enemy.UI_Icon_OnRButtonUp, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionRButtonUp, FrameManager.OnRButtonUp, GroupSpotter.OnRButtonUp, MapMonster.OnMouseRightClickPin, MiniMapMonster.OnMouseRightClickPin, NBSBCore.DoMenu, NBSBParam.OnParamChange, PlayEffectsOn.OnRButtonUp, SORRealms.ShowKeepMenu, SquaredClickConfig.AbilityClear | `flags, x, y` | MEDIUM |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | Twister.OnLButtonDown, NaturalLog.ScrollToBottom, NaturalLog.ScrollToTop, DetauntHelperConfig.UpdateConfig, Enemy.CombatLogUI_IDS_OnRowLButtonDown, FrameManager.OnLButtonDown, MBuffGui.OnSpellLClick, MBuffSetup.SmartBuff.OnOtherLClick, OnTabLBU, RVMOD_SquaredDistances.OnLButtonDownJoystick, Squared.ToggleActive, Warbuilder.OnTabLBU | `flags, x, y` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | CastSequence.Builder.OnSettingsOnMouseOverEnd, Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOverEnd, FrameManager.OnMouseOverEnd, GroupSpotter.OnMouseOverEnd, NBSBCore.IconOnMouseOverEnd, RVMOD_SquaredDistances.OnMouseOverEndJoystick, SORPager.BOsOff, TomeTracker.Saga.OnMouseOverMapPinEnd, TomeTracker.Saga.OnMouseOverRewardEnd | `` |  |
| [OnRButtonDown](../handlers/handler_OnRButtonDown.md) | input | FrameManager.OnRButtonDown, MBuffGui.OnSpellRClick, MBuffSetup.SmartBuff.OnOtherRClick, Squared.ShowSettings | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnMouseOver** handlers call: `DynamicImageSetTexture`, `SystemData.ActiveWindow.name:match`, `SystemData.MouseOverWindow.name:match`, `WindowGetAlpha`, `WindowGetId`, `WindowGetParent`, `WindowSetAlpha`, `WindowSetShowing`, `WindowStartAlphaAnimation`, `WindowStopAlphaAnimation`

**OnLButtonUp** handlers call: `BroadcastEvent`, `ComboBoxSetSelectedMenuItem`, `Cursor.Clear`, `Cursor.IconOnCursor`, `DynamicImageSetTexture`, `DynamicImageSetTextureDimensions`, `LabelSetText`, `Player.GetAbilityData`, `SystemData.ActiveWindow.name:match`, `SystemData.ActiveWindow.name:sub`, `SystemData.MouseOverWindow.name:match`, `WindowAddAnchor`, `WindowClearAnchors`, `WindowGetId`, `WindowGetParent`, `WindowGetShowing`, `WindowSetShowing`, `WindowUtils.ToggleShowing`

**OnRButtonUp** handlers call: `DynamicImageSetTexture`, `SystemData.ActiveWindow.name:match`, `SystemData.MouseOverWindow.name:match`, `WindowGetId`, `WindowGetShowing`, `WindowSetShowing`

**OnLButtonDown** handlers call: `ButtonSetPressedFlag`, `SystemData.ActiveWindow.name:match`, `WindowGetId`, `WindowGetShowing`, `WindowSetAlpha`, `WindowSetShowing`

**OnMouseOverEnd** handlers call: `DynamicImageSetTexture`, `WindowGetAlpha`, `WindowGetParent`, `WindowSetShowing`, `WindowStartAlphaAnimation`, `WindowStopAlphaAnimation`

**OnRButtonDown** handlers call: `WindowGetId`

## Common Inherits

- $parentCareerBuffIcon1
- $parentCareerIcon1
- Aggro_Tactic_Template
- EA_Image_HUDMenuButton
- EA_ListSortDownArrow
- EA_ListSortUpArrow
- GroupLeaderCrown
- LoadingScreenBookImage
- MapMonsterIconChooserWindow_IconTemplate
- SpamBayesTemplateBackground
- SpamBayesTemplateBorder
- TomeDiamondDivider

## Common Parent Elements

- [Windows](element_Windows.md) — 1288× (HIGH)
- [Window](element_Window.md) — 1× (LOW)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 1187× (HIGH)
- [Size](element_Size.md) — 965× (HIGH)
- [TintColor](element_TintColor.md) — 374× (HIGH)
- [TexDims](element_TexDims.md) — 156× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 155× (HIGH)
- [TexCoords](element_TexCoords.md) — 76× (HIGH)
- [Windows](element_Windows.md) — 15× (HIGH)
- [Anchor](element_Anchor.md) — 1× (LOW)
- [Color](element_Color.md) — 1× (LOW)
- [Sounds](element_Sounds.md) — 1× (LOW)

## Common Template Bases

- APBarBackground
- APBarFrame
- Aggro_Tactic_Template
- CleanRvRFlagIndicator
- EA_Default_AuctionImage
- EA_Default_CharacterImage
- EA_Default_CornerImage
- EA_Default_CustomizeUIImage
- EA_Default_MerchantImage
- EA_Default_SocialImage
- EA_Default_SpellbookImage
- EA_Default_TrainingImage
- EA_DynamicImage_DefaultSeparatorRight
- EA_DynamicImage_DefaultVerticalSeparatorBottom
- EA_Image_DefaultIcon
- EA_Image_DefaultIconFrame
- EA_Image_HUDMenuButton
- EA_ListSortDownArrow
- EA_ListSortUpArrow
- EA_Templates_BrassCoin
- EA_Templates_GoldCoin
- EA_Templates_SilverCoin
- FriendsImage
- GenericEndCap
- GroupLeaderCrown
- HG_HealGridProgressBarStatusBarTemplate
- HealthBarBackground
- HealthBarFrame
- LoadingScreenBookImage
- LoadingScreenDividerImage
- LoadingScreenScenarioExitImage
- LoadingScreenShieldImage
- MainAssistCrown
- MapMonsterIconChooserWindow_IconTemplate
- MarkBuffSetupSmartBuffWindowOtherBuffIcon1
- MarkBuffSetupWindowCareerBuffIcon1
- MarkBuffSetupWindowCareerIcon1
- MasterLooterImage
- PortraitBackground
- PortraitFrame
- RVMOD_ManagerStatusTemplate
- ReliquaryHunterIcon
- RvRFlagIndicator
- SpamBayesTemplateBackground
- SpamBayesTemplateBorder
- TargetLevelBackgroundTemplate
- TomeDiamondDivider
- WarbandAssistantImage
- WarbandLeaderCrown


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `handleinput` | optional | 78% | false, true |
| `texture` | optional | 71% | EA_SquareFrame, AggroMeterIcon, EA_Abilities01_d5, EA_HUD_01, ... |
| `layer` | optional | 65% | background, overlay, popup, default, ... |
| `texturescale` | optional | 33% | 0.35, 1.00, .5, 1.5, ... |
| `textureScale` | optional | 14% | 1.0, 0.5, 0.43, 1, ... |
| `slice` | optional | 12% | Tab-ALL, Round-Swatch-Selection-Ring, Radio-Button, You-Have-Mail, ... |
| `inherits` | optional | 12% | EA_Default_TrainingImage, Aggro_Tactic_Template, EA_ListSortUpArrow, EA_ListSortDownArrow, ... |
| `popable` | optional | 7% | false, true |
| `id` | optional | 6% | 1, 2, 3, 5, ... |
| `alpha` | optional | 5% | 1.0, 1, 0.125, 0.0, ... |
| `sticky` | optional | 5% | false, true |
| `filtering` | optional | 4% | true |
| `savesettings` | optional | 2% | false, true |
| `movable` | optional | 0% | false, true |
| `scale` | optional | 0% | 1.0, 1, 0.5 |
| `poppable` | optional | 0% | false |
| `numsegments` | optional | 0% | 8, 32, 64 |
| `parent` | optional | 0% | $parent, root |
| `show` | optional | 0% | true |
| `mirrorTexCoords` | optional | 0% | true, false |
| `skipinput` | optional | 0% | true |
| `virtual` | optional | 0% | true |
| `Slice` | optional | 0% | Influence-Reward |
| `draganddrop` | optional | 0% | true |
| `textureAlpha` | optional | 0% | 0.5 |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 1187 times as an unnamed child.

### [Size](element_Size.md)

Observed 965 times as an unnamed child.

### [TintColor](element_TintColor.md)

Observed 374 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 110, 130, 50 |
| `g` | **required** | 0, 200, 185, 130 |
| `r` | **required** | 255, 0, 180, 200 |
| `a` | optional | 30, 255, 0.8, 175 |
### [TexDims](element_TexDims.md)

Observed 156 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 64, 256, 32, 48 |
| `y` | **required** | 64, 256, 32, 48 |
### [EventHandlers](element_EventHandlers.md)

Observed 155 times as an unnamed child.

### [TexCoords](element_TexCoords.md)

Observed 76 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 0, 133, 88, 113 |
| `y` | optional | 0, 163, 51, 90 |
### [Windows](element_Windows.md)

Observed 15 times as an unnamed child.

### [Anchor](element_Anchor.md)

Observed 1 times as an unnamed child.

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
### [Color](element_Color.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 255, 0, 73, 225 |
| `g` | **required** | 25, 253, 218, 255 |
| `r` | **required** | 255, 253, 155, 245 |
| `a` | optional | 255, 0.8, 0, 128 |
### [Sounds](element_Sounds.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [DynamicImage](element_DynamicImage.md)
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
      - (cycle)
    - [color](element_color.md) (structural, 1×, LOW)
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
        - (cycle)
      - [color](element_color.md) (structural, 1×, LOW)
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

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowGetId` | ui | 65 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonDown, OnRButtonUp |
| `WindowSetShowing` | ui | 18 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp |
| `WindowGetParent` | ui | 16 | OnLButtonUp, OnMouseOver, OnMouseOverEnd |
| `DynamicImageSetTexture` | ui | 12 | OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp |
| `SystemData.MouseOverWindow.name:match` | data | 11 | OnLButtonUp, OnMouseOver, OnRButtonUp |
| `SystemData.ActiveWindow.name:match` | data | 7 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnRButtonUp |
| `ButtonSetPressedFlag` | ui | 6 | OnLButtonDown |
| `Cursor.Clear` | ui | 6 | OnLButtonUp |
| `Cursor.IconOnCursor` | data | 6 | OnLButtonUp |
| `WindowGetShowing` | ui | 6 | OnLButtonDown, OnLButtonUp, OnRButtonUp |
| `WindowGetAlpha` | ui | 4 | OnMouseOver, OnMouseOverEnd |
| `WindowStartAlphaAnimation` | ui | 4 | OnMouseOver, OnMouseOverEnd |
| `WindowStopAlphaAnimation` | ui | 4 | OnMouseOver, OnMouseOverEnd |
| `ComboBoxSetSelectedMenuItem` | ui | 3 | OnLButtonUp |
| `Player.GetAbilityData` | data | 3 | OnLButtonUp |
| `WindowSetAlpha` | ui | 3 | OnLButtonDown, OnMouseOver |
| `BroadcastEvent` | event | 2 | OnLButtonUp |
| `DynamicImageSetTextureDimensions` | ui | 2 | OnLButtonUp |
| `LabelSetText` | ui | 1 | OnLButtonUp |
| `SystemData.ActiveWindow.name:sub` | data | 1 | OnLButtonUp |
| `WindowAddAnchor` | ui | 1 | OnLButtonUp |
| `WindowClearAnchors` | ui | 1 | OnLButtonUp |
| `WindowUtils.ToggleShowing` | ui | 1 | OnLButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

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
### OnMouseOver

Confidence: HIGH

### OnMouseOverEnd

Confidence: HIGH

### OnRButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
## Lua Functions Manipulating This Type

- ClosetGoblinCharacterWindow.OnInitialize
- Enemy.Guard_OnSettingsChanged
- GroupSpotter.Initialize
- KwestorTracker.DrawQuest
- MiniMapMonster.local.DisplayPin
- ClosetGoblinZoneWindow.OnInitialize
- CMapWindow.UpdateMailIcon
- CallingSetup.OnShowTargetIconClick
- CallingSetup.OnShowCallerIconClick
- NerfedButtons.Initialize
- RandomMountUI.OnDropSlotLButtonUp
- SwiftAssist.Initialize
- mmv.SetNewLayout
- Twister.OnUpdate
- VerticalMorale.Create
- HealGridIcon.Initialize
- TomeTracker.Journal.MapButtonToggle
- ClosetGoblin.Initialize
- Calling.ShowNotification
- MapMonster.CreateMarker
- GuardBot.SetupIconWindow
- RVMOD_3DPortrait.Update3DPortraitScenes
- Enemy.TalismanAlerterInitialize
- GuardLine.update
- KeyBar.Update
- Dye.Initialize
- UiModVersionMismatchWindow.UpdateModSortButtons
- KeyBar.MouseOver
- Squared.SavePosition
- TomeTracker.Journal.Initialize
- GetStats.OnInitialize
- MapPin.TestTooltip
- BlackBookWindow.UpdateSortButtons
- CCM.OnInitialize
- Enemy.EnemyEffectsIndicator:Update
- GroupSpotter.InitializeSettingsWindow
- GuardBot.AttachIconWindow
- KeyBar.MoveWindow
- MiniMapMonster.RemovePin
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- MapMonster.Editor.ShowZoneHooked
- Moth.UpdateTarget
- NBSBCore.IconOnMouseOver
- Trakario.SetCarrier
- Enemy.AssistUI_Target_Show
- GuildWardenWin.WinSetup
- KeyBarSettings.SaveSettingsWindow
- WSCT.ColorOnInitialize
- yAssist.Initialize
- Calling.ShowCallerIcon
- KeyBar.OnInitialize
- MapMonster.DisplayPinsForZone
- ResHelp.OnUpdate
- Swift Assist.SetSmartLabel
- TalismanMonitor.Setup
- BloodyMess.PlayEffect
- Calling.ManageNotification
- Enemy._Initialize
- MoraleBar.Create
- WSCT.ColorAcceptButtonOnButtonUp
- UiModWindow.UpdateModSortButtons
- MouseHint.OnInitialize
- CoolDownLine.OnUpdate
- Enemy.UI_Icon_Switch
- Enemy.TalismanAlerter_Update
- Swift Assist.local.SetTexLabel
- TomeTracker.Journal.MapButtonHide
- MapMonster.local.DisplayPinsForZone
- SocialWindow.OnMouseOverListMember
- TacticSetNames.local.UpdateSliders
- WSCT.ColorHideMenu
- CallingSetup.SetAutoAssist
- Calling.Initialize
- RoR_SoR.OnScenario
- Swift Assist.SetTexLabel
- yAssist.SetTexLabel
- TomeWindow.DK_SetSortArrow
- MiniMapMonster.DisplayPin
- TacticSetNames.UpdateSliders
- TomeTracker.Journal.MapButtonShow
- CMapWindow.OnRvRFlagUpdated
- PP.UpdateListRow
- Swift Assist.local.SetSmartLabel
- Enemy.Guard_GuardIndicator_Update
- Enemy.MarksInitialize
- GroupSpotter.OnMouseOverEnd
- GroupSpotter.toggle
- MapMonster.DisplayPin
- CCTV.Update
- MapMonster.local.CreateMarker
- RoR_SoR.OnInitialize
- BagOMatic.init
- yAssist.SetAssist
- CCM.UpdateUI
- GroupSpotter.OnMouseOverStart
- NBSBCore.IconOnMouseOverEnd
- TokenMachine.Initialize
- WSCT.OnLButtonUpColorPicker
- BankArkel.PackImg
- MapMonster.local.DisplayPin
- TacticSetNames.UpdateColorPickerPreview
- TacticSetNames.local.UpdateColorPickerPreview
- TurretRange.OnUpdate
- Twister.OnLoad
- SwiftAssist.OnMacroUpdated
- VerticalMorale.Initialize
- ClosetGoblinCharacterWindow.UpdateSortButtonIcon
- Enemy.GuardInitialize
- MouseHint.OnUpdate
- Twister.SetActiveAura
- mms.SetNewLayout
- Enemy.TalismanAlerter_OnSettingsChanged
- Trakario.local.SetCarrier
- zmm.UpdateLayout
- GuardBot.ResetGuardData
- NerfedButtons.Shutdown
- RoR_SoR.OnCombat
- Squared.RestorePosition
- Calling.ShowTargetIcon
- RandomMountUI.OnAddCustomMount


## Binding Resolution

- Total handler declarations: 222
- Resolved to Lua functions: 222 (100%)

## .mod Lifecycle: Startup Windows

This element type is instantiated as a startup window by the following .mod addon(s):

| Frame Name | Addon | Hook | Resolution | Confidence |
| --- | --- | --- | --- | --- |
| DetauntHelperIcon | DetauntHelper | OnInitialize | exact | HIGH |
| SquaredIcon | Squared | OnInitialize | exact | HIGH |
## Seen In

- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Aura
- AutoMark
- BagOMatic
- BankArkel
- BlackBook
- BlackBox
- Bloody Mess
- Brizio's Crappy Computer Medic
- BuffHead
- CCTV
- CDown
- CM_ClosetGoblin
- CMap
- Calling
- CastSequence
- CleanUnitFrames
- CleansingBuddy
- CombatTextNames
- Compass3D
- CoolDownLine
- Crusher
- DAoCBuff
- DPSMeter
- DammazKron
- Dascore
- Deathblow
- Deathblow2
- DetauntHelper
- Duel
- DuffTimer
- Dye Preview
- EA_LoadingScreen
- EA_OpenPartyWindow
- EA_ThreePartBar
- EA_UiModWindow
- EZCraftX
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- FastFriends
- FlagCap
- FozAuction
- GCDTracker
- GDes
- Ges
- GetStats
- Group Icons
- GroupRange
- GroupSpotter
- GuardBot
- GuardLine
- GuildWarden
- HealGrid
- Hopper
- InfoScroller
- JunkDump
- KeyBar
- KillStreak
- Killer
- Kwestor
- LibAddonButton
- LibSurveyor
- LibWBToggler
- LootAlert
- MapMonster
- MapPin
- MarkBuff
- MegaphonePlusPlus
- MiniMapMonster
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Moth
- Motion
- MouseHint
- NaturalLog
- NerfedButtons
- Obsidian
- Paint the leader
- PartyCast
- PlayEffectsOn
- Pocket Palette
- PotionBar
- Pure
- Pure Careerbar
- Queue Queuer
- QuickTacticSwitch
- RVAPI_ColorDialog
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_SquaredDistances
- RVMOD_Targets
- RandomMount
- Refer
- ReliquaryHunter
- Res
- ResHelp
- RetAlert
- RoR_SoR
- SNT_BUTTONS
- SNT_CASTBAR
- SNT_INFO
- SNT_PANEL
- SOR
- Shinies
- SocialWindow 2.0
- Soloq
- SpamBayes
- Squared
- SquaredClick
- Swift Assist
- Swinger
- TacticSetNames
- TargetInfoRing
- Targets
- TastyButtons
- TaxPayer
- TexturedButtons
- TheSeeker
- TidyRoll
- TokenMachine
- TomeTracker
- Trakario
- TurretRange
- TurretScrap
- Twister
- Vectors
- VerticalMorale
- WSCT
- WaaaghBar
- WarBoard_Menu
- WarBoard_TaliMon
- Warbuilder
- WhatsCooking
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- ZonePOP
- alertMod
- compass
- minesweep
- nLootLink
- talisman-monitor
- yAssistHelper
- zMailMod

## Examples

- AdvancedRenownTrainer: AbilityButtonTemplateSquare -> DynamicImage AbilityButtonTemplateSquare
- AdvancedRenownTrainer: AbilityButtonTemplateSquareFrame -> DynamicImage AbilityButtonTemplateSquareFrame
- AdvancedRenownTrainer: AdvancedRenownTrainingWindowCornorImage -> DynamicImage AdvancedRenownTrainingWindowCornorImage
- AggroMeter: AggroMeterWindow_AggroWindow1Tactic -> DynamicImage AggroMeterWindow_AggroWindow1Tactic
- AggroMeter: AggroMeterWindow_AggroWindow2Tactic -> DynamicImage AggroMeterWindow_AggroWindow2Tactic
- AggroMeter: AggroMeterWindow_AggroWindow3Tactic -> DynamicImage AggroMeterWindow_AggroWindow3Tactic

## Related APIs

- [Anchor](element_Anchor.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [Color](element_Color.md) (HIGH 100/100) - XML Element Type
- [EA_Default_CharacterImage](../../globals/constants/constant_EA_Default_CharacterImage.md) (HIGH 100/100) - Constant
- [EA_Default_CornerImage](../../globals/constants/constant_EA_Default_CornerImage.md) (HIGH 100/100) - Constant
- [EA_Default_CustomizeUIImage](../../globals/constants/constant_EA_Default_CustomizeUIImage.md) (HIGH 100/100) - Constant
- [EA_Default_SocialImage](../../globals/constants/constant_EA_Default_SocialImage.md) (HIGH 100/100) - Constant
- [EA_Default_TrainingImage](../../globals/constants/constant_EA_Default_TrainingImage.md) (HIGH 100/100) - Constant
- [EA_DynamicImage_DefaultSeparatorRight](../../globals/constants/constant_EA_DynamicImage_DefaultSeparatorRight.md) (HIGH 100/100) - Constant
- [EA_Image_DefaultIcon](../../globals/constants/constant_EA_Image_DefaultIcon.md) (HIGH 100/100) - Constant
- [EA_Image_DefaultIconFrame](../../globals/constants/constant_EA_Image_DefaultIconFrame.md) (HIGH 100/100) - Constant
- [EA_Image_HUDMenuButton](../../globals/constants/constant_EA_Image_HUDMenuButton.md) (HIGH 100/100) - Constant
- [EA_ListSortDownArrow](../../globals/constants/constant_EA_ListSortDownArrow.md) (HIGH 100/100) - Constant
- [EA_ListSortUpArrow](../../globals/constants/constant_EA_ListSortUpArrow.md) (HIGH 100/100) - Constant
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Sounds](element_Sounds.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TexDims](element_TexDims.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EA_Default_AuctionImage](../../globals/constants/constant_EA_Default_AuctionImage.md) (HIGH 90/100) - Constant
- [EA_Default_MerchantImage](../../globals/constants/constant_EA_Default_MerchantImage.md) (HIGH 90/100) - Constant
- [EA_Default_SpellbookImage](../../globals/constants/constant_EA_Default_SpellbookImage.md) (HIGH 90/100) - Constant
- [EA_DynamicImage_DefaultVerticalSeparatorBottom](../../globals/constants/constant_EA_DynamicImage_DefaultVerticalSeparatorBottom.md) (HIGH 90/100) - Constant
- [EA_Templates_BrassCoin](../../globals/constants/constant_EA_Templates_BrassCoin.md) (HIGH 90/100) - Constant
- [EA_Templates_GoldCoin](../../globals/constants/constant_EA_Templates_GoldCoin.md) (HIGH 90/100) - Constant
- [EA_Templates_SilverCoin](../../globals/constants/constant_EA_Templates_SilverCoin.md) (HIGH 90/100) - Constant

## Used With

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
