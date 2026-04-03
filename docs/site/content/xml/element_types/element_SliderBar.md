# SliderBar

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
| Addons seen in | AdjustTheTip, Atlas, Aura, Bloody Mess, BuffHead, CCTV, CDown, ChattyCathy |
| Files seen in | Code/UnitFrames/EffectsIndicatorDialog.xml, Configuration/Main.xml, Gui/HealGridGuiTabBattlegroup.xml, Gui/HealGridGuiTabGroup.xml, Gui/HealGridGuiTabHUD.xml, Gui/HealGridGuiTabHUDBuffs.xml, Gui/HealGridGuiTabScenariogroup.xml, Gui/HealGridGuiTemplates.xml |
| Namespaces detected | SliderBar |
| Source kinds | xml_frames |
| Example locations | AdjustTheTip: AdjustTheTipMenuItemSliderSliderBar, Atlas: AtlasConfigurationFrameContentMapSettingsTransparencyOnMouseOverSlider, Atlas: AtlasConfigurationFrameContentMapSettingsTransparencyOnMovementSlider, Atlas: AtlasConfigurationFrameContentMapSettingsTransparencySlider, Atlas: TAtlasSlider, Aura: AuraColorPickerAlpha |
| XML usage count | 227 |
| XML attribute usage count | 227 |
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

SliderBar is an interactive XML control. It commonly appears under Button and Window. It is typically used to organize structural children such as Anchors, EventHandlers, Size and bind XML events like OnLButtonUp, OnMouseOver, OnSlide to Lua.

## Common Attributes

- autoresize
- background
- handleinput
- handlinput
- inherits
- locktoticks
- name
- numTicks
- numticks
- scale
- sliderbutton
- tickmark

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnSlide](../handlers/handler_OnSlide.md)

## Common Handler Functions

- snt_panel.slider_change
- snt_info.slider_change
- AuraColorPicker.OnSlide
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- GroupIconsSG.OptionsSliderMove
- BuffHead.Setup.SelectColor.OnSlideTint
- MapMonster.PinTypeEditor.OnSetCustomColor
- MoraleCircle.OnSetCustomColor
- MoraleCircle.OnSetCustomColorEmpty
- MoraleCircle.OnSetCustomColorFill
- MoraleCircle.OnSetCustomColorFull
- Obsidian.Setup.SelectColor.OnSlideTint
- TastyButtonsOptions.OnSlidingColor
- TexturedButtons.Setup.SelectColor.OnSlideTint
- TexturedButtons.Setup.Tint.OnSlideTint
- TurretRange.Setup.Display.OnSlideTint
- TurretRange.Setup.Distance.OnSlideTint
- WSCT.OnSetCustomColor
- WarBoard_FPSOptions.OnSlidingColor
- SOROptions.Slide
- Atlas.Configuration.OnSlide
- AuraConfig.OnTextureRotationSlide
- AuraConfig.OnTextureScaleSlide
- AuraConfig.OnTimerScaleSlide
- BuffHead.Setup.AdvancedContainersItem.Properties.OnSlideScale
- BuffHead.Setup.Container.OnSlideColumns
- BuffHead.Setup.Container.OnSlidePaddingHorizontal
- BuffHead.Setup.Container.OnSlidePaddingVertical
- BuffHead.Setup.Container.OnSlideRows
- BuffHead.Setup.Display.OnSlideIndicatorScale
- BuffHead.Setup.Display.OnSlidePaddingHorizontal
- BuffHead.Setup.Display.OnSlidePaddingVertical
- BuffHead.Setup.General.OnSlideMaximumThreshold
- BuffHead.Setup.Layout.Properties.OnSlideAlphaAlpha
- BuffHead.Setup.Layout.Properties.OnSlideIconBorderAlpha
- BuffHead.Setup.Layout.Properties.OnSlideSizeScale
- BuffHead.Setup.Layout.Properties.OnSlideStatusBarBackgroundAlpha
- BuffHead.Setup.Layout.Properties.OnSlideStatusBarForegroundAlpha
- BuffHead.Setup.Performance.OnSlideGeneralUpdateDelay
- BuffHead.Setup.Performance.OnSlideMaximumUpdates
- BuffHead.Setup.Performance.OnSlidePriorityUpdateDelay
- BuffHead.Setup.Performance.OnSlidePriorityUpdateStart
- BuffHead.Setup.Performance.OnSlideResyncTargetDelay
- CDownSettings.OnSlideABCBlue
- CDownSettings.OnSlideABCGreen
- CDownSettings.OnSlideABCRed
- CDownSettings.OnSlideLength
- CDownSettings.OnSlideSCBlue
- CDownSettings.OnSlideSCGreen
- CDownSettings.OnSlideSCRed
- ChattyCathy.OnSlideAlpha
- DAoCBuffSettings.FilterSettings.OnSlideBB
- DAoCBuffSettings.FilterSettings.OnSlideBG
- DAoCBuffSettings.FilterSettings.OnSlideBR
- DAoCBuffSettings.FilterSettings.OnSlideCB
- DAoCBuffSettings.FilterSettings.OnSlideCG
- DAoCBuffSettings.FilterSettings.OnSlideCR
- DetauntBarSettings.OnHeightSlide
- DetauntBarSettings.OnScaleSlide
- DetauntBarSettings.OnWidthSlide
- DuffTimer.Options.OnSliderChange
- GroupRangeSetup.General.OnSlideUpdateDelay
- GroupSpotter.Settings.SlidergrpAlpha
- GroupSpotter.Settings.SlidersingleAlpha
- GroupSpotter.Settings.colorslidergrpB
- GroupSpotter.Settings.colorslidergrpG
- GroupSpotter.Settings.colorslidergrpR
- GroupSpotter.Settings.colorslidersingleB
- GroupSpotter.Settings.colorslidersingleG
- GroupSpotter.Settings.colorslidersingleR
- HealGridGuiTabBattlegroup.BuffUpdateThrottleChanged
- HealGridGuiTabBattlegroup.UnitFrameHeightChanged
- HealGridGuiTabBattlegroup.UnitFrameScaleChanged
- HealGridGuiTabBattlegroup.UnitFrameWidthChanged
- HealGridGuiTabBattlegroup.UpdateThrottleChanged
- HealGridGuiTabGroup.BuffUpdateThrottleChanged
- HealGridGuiTabGroup.UnitFrameHeightChanged
- HealGridGuiTabGroup.UnitFrameScaleChanged
- HealGridGuiTabGroup.UnitFrameWidthChanged
- HealGridGuiTabGroup.UpdateThrottleChanged
- HealGridGuiTabHUD.BuffUpdateThrottleChanged
- HealGridGuiTabHUD.UnitFrameHeightChanged
- HealGridGuiTabHUD.UnitFrameScaleChanged
- HealGridGuiTabHUD.UnitFrameWidthChanged
- HealGridGuiTabHUD.UpdateThrottleChanged
- HealGridGuiTabHUDBuffs.BuffFrameSizeAdjustChanged
- HealGridGuiTabHUDBuffs.ShowBuffsMaxChanged
- HealGridGuiTabScenariogroup.BuffUpdateThrottleChanged
- HealGridGuiTabScenariogroup.NumGroupsDisplayedChanged
- HealGridGuiTabScenariogroup.UnitFrameHeightChanged
- HealGridGuiTabScenariogroup.UnitFrameScaleChanged
- HealGridGuiTabScenariogroup.UnitFrameWidthChanged
- HealGridGuiTabScenariogroup.UpdateThrottleChanged
- LibGroup.Setup.OnSlideGroupDistanceCacheUpdate
- LibGroup.Setup.OnSlideGroupDistanceSearchUpdate
- LibGroup.Setup.OnSlideGroupUpdateDelay
- MapPin.onRotateSlide
- MapPin.onScaleSlide
- Obsidian.Setup.Castbar.OnSlideBackgroundAlpha
- Obsidian.Setup.Castbar.OnSlideFillAlpha
- Obsidian.Setup.Castbar.OnSlideGlobalCooldownBackgroundAlpha
- Obsidian.Setup.Castbar.OnSlideGlobalCooldownDividerAlpha
- Obsidian.Setup.Castbar.OnSlideGlobalCooldownFillAlpha
- Obsidian.Setup.Castbar.OnSlideGlobalCooldownFillReadyAlpha
- Obsidian.Setup.Castbar.OnSlideGlobalCooldownSparkAlpha
- Obsidian.Setup.Castbar.OnSlideIconAlpha
- Obsidian.Setup.Castbar.OnSlideIconScale
- Obsidian.Setup.Castbar.OnSlideLatencyAlpha
- Obsidian.Setup.Castbar.OnSlideLatencyTextAlpha
- Obsidian.Setup.Castbar.OnSlideLatencyTextScale
- Obsidian.Setup.Castbar.OnSlideNameAlpha
- Obsidian.Setup.Castbar.OnSlideNameScale
- Obsidian.Setup.Castbar.OnSlideNameWidth
- Obsidian.Setup.Castbar.OnSlideSparkAlpha
- Obsidian.Setup.Castbar.OnSlideTimerAlpha
- Obsidian.Setup.Castbar.OnSlideTimerDecimals
- Obsidian.Setup.Castbar.OnSlideTimerScale
- Obsidian.Setup.EffectTracker.OnSlideBackgroundAlpha
- Obsidian.Setup.EffectTracker.OnSlideFillAlpha
- Obsidian.Setup.EffectTracker.OnSlideNameAlpha
- Obsidian.Setup.EffectTracker.OnSlideNameScale
- Obsidian.Setup.EffectTracker.OnSlideNameWidth
- Obsidian.Setup.EffectTracker.OnSlideTimerAlpha
- Obsidian.Setup.EffectTracker.OnSlideTimerScale
- Obsidian.Setup.EffectTracker.OnSlideTrackerIconAlpha
- Obsidian.Setup.EffectTracker.OnSlideTrackerIconScale
- Obsidian.Setup.EffectTracker.OnSlideTrackerMaximumEffects
- PotionBarSettings.OnAlphaSliderChanged
- PotionBarSettings.OnScaleSliderChanged
- RVAPI_ColorDialog.OnSlide
- RVAPI_Range.OnSlideMapDelay
- RVMOD_Manager.OnSlideFadeInOutDelay
- RVMOD_Manager.OnSlideScale
- RVMOD_Manager.OnSlideZoomInOutDelay
- RVMOD_SquaredDistances.OnSlideDistancesDelay
- RVMOD_SquaredDistances.OnSlideScale
- ReliquaryHunterOptionsWindow.OnWorldMapWindowBackgroundOpacitySlide
- ReliquaryHunterOptionsWindow.OnWorldMapWindowOpacitySlide
- RoR_SoR.OnSlideWindowOptionsOffset
- RoR_SoR.OnSlideWindowOptionsOpacity
- RoR_SoR.OnSlideWindowOptionsScale
- ScenarioGroupWindow.OnOpacitySlide
- ShiniesConfigGeneral.OnSlide_UIScale
- Statdoll.Options.onAlphaSlide
- Statdoll.Options.onScaleSlide
- TacticSetNames.SettingsWindow.OnSlide
- TastyButtonsOptions.OnSlidingCeateView
- TastyButtonsOptions.OnSlidingEditView
- TexturedButtons.Setup.Cooldown.OnSlideCooldownAlpha
- TurretRange.Setup.Display.OnSlideDistanceScale
- TurretRange.Setup.General.OnSlideUpdateDelay
- WSCT.SliderOnSlide
- WarBoard.Options.OnSlide
- XpStatus.OnOpacitySlide
- alertMod.OnSliderChanged
- WSCT.OnMouseOver


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnSlide](../handlers/handler_OnSlide.md) | custom | snt_panel.slider_change, snt_info.slider_change, AuraColorPicker.OnSlide, Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample, GroupIconsSG.OptionsSliderMove, BuffHead.Setup.SelectColor.OnSlideTint, MapMonster.PinTypeEditor.OnSetCustomColor, MoraleCircle.OnSetCustomColor, MoraleCircle.OnSetCustomColorEmpty, MoraleCircle.OnSetCustomColorFill, MoraleCircle.OnSetCustomColorFull, Obsidian.Setup.SelectColor.OnSlideTint, TastyButtonsOptions.OnSlidingColor, TexturedButtons.Setup.SelectColor.OnSlideTint, TexturedButtons.Setup.Tint.OnSlideTint, TurretRange.Setup.Display.OnSlideTint, TurretRange.Setup.Distance.OnSlideTint, WSCT.OnSetCustomColor, WarBoard_FPSOptions.OnSlidingColor, SOROptions.Slide, Atlas.Configuration.OnSlide, AuraConfig.OnTextureRotationSlide, AuraConfig.OnTextureScaleSlide, AuraConfig.OnTimerScaleSlide, BuffHead.Setup.AdvancedContainersItem.Properties.OnSlideScale, BuffHead.Setup.Container.OnSlideColumns, BuffHead.Setup.Container.OnSlidePaddingHorizontal, BuffHead.Setup.Container.OnSlidePaddingVertical, BuffHead.Setup.Container.OnSlideRows, BuffHead.Setup.Display.OnSlideIndicatorScale, BuffHead.Setup.Display.OnSlidePaddingHorizontal, BuffHead.Setup.Display.OnSlidePaddingVertical, BuffHead.Setup.General.OnSlideMaximumThreshold, BuffHead.Setup.Layout.Properties.OnSlideAlphaAlpha, BuffHead.Setup.Layout.Properties.OnSlideIconBorderAlpha, BuffHead.Setup.Layout.Properties.OnSlideSizeScale, BuffHead.Setup.Layout.Properties.OnSlideStatusBarBackgroundAlpha, BuffHead.Setup.Layout.Properties.OnSlideStatusBarForegroundAlpha, BuffHead.Setup.Performance.OnSlideGeneralUpdateDelay, BuffHead.Setup.Performance.OnSlideMaximumUpdates, BuffHead.Setup.Performance.OnSlidePriorityUpdateDelay, BuffHead.Setup.Performance.OnSlidePriorityUpdateStart, BuffHead.Setup.Performance.OnSlideResyncTargetDelay, CDownSettings.OnSlideABCBlue, CDownSettings.OnSlideABCGreen, CDownSettings.OnSlideABCRed, CDownSettings.OnSlideLength, CDownSettings.OnSlideSCBlue, CDownSettings.OnSlideSCGreen, CDownSettings.OnSlideSCRed, ChattyCathy.OnSlideAlpha, DAoCBuffSettings.FilterSettings.OnSlideBB, DAoCBuffSettings.FilterSettings.OnSlideBG, DAoCBuffSettings.FilterSettings.OnSlideBR, DAoCBuffSettings.FilterSettings.OnSlideCB, DAoCBuffSettings.FilterSettings.OnSlideCG, DAoCBuffSettings.FilterSettings.OnSlideCR, DetauntBarSettings.OnHeightSlide, DetauntBarSettings.OnScaleSlide, DetauntBarSettings.OnWidthSlide, DuffTimer.Options.OnSliderChange, GroupRangeSetup.General.OnSlideUpdateDelay, GroupSpotter.Settings.SlidergrpAlpha, GroupSpotter.Settings.SlidersingleAlpha, GroupSpotter.Settings.colorslidergrpB, GroupSpotter.Settings.colorslidergrpG, GroupSpotter.Settings.colorslidergrpR, GroupSpotter.Settings.colorslidersingleB, GroupSpotter.Settings.colorslidersingleG, GroupSpotter.Settings.colorslidersingleR, HealGridGuiTabBattlegroup.BuffUpdateThrottleChanged, HealGridGuiTabBattlegroup.UnitFrameHeightChanged, HealGridGuiTabBattlegroup.UnitFrameScaleChanged, HealGridGuiTabBattlegroup.UnitFrameWidthChanged, HealGridGuiTabBattlegroup.UpdateThrottleChanged, HealGridGuiTabGroup.BuffUpdateThrottleChanged, HealGridGuiTabGroup.UnitFrameHeightChanged, HealGridGuiTabGroup.UnitFrameScaleChanged, HealGridGuiTabGroup.UnitFrameWidthChanged, HealGridGuiTabGroup.UpdateThrottleChanged, HealGridGuiTabHUD.BuffUpdateThrottleChanged, HealGridGuiTabHUD.UnitFrameHeightChanged, HealGridGuiTabHUD.UnitFrameScaleChanged, HealGridGuiTabHUD.UnitFrameWidthChanged, HealGridGuiTabHUD.UpdateThrottleChanged, HealGridGuiTabHUDBuffs.BuffFrameSizeAdjustChanged, HealGridGuiTabHUDBuffs.ShowBuffsMaxChanged, HealGridGuiTabScenariogroup.BuffUpdateThrottleChanged, HealGridGuiTabScenariogroup.NumGroupsDisplayedChanged, HealGridGuiTabScenariogroup.UnitFrameHeightChanged, HealGridGuiTabScenariogroup.UnitFrameScaleChanged, HealGridGuiTabScenariogroup.UnitFrameWidthChanged, HealGridGuiTabScenariogroup.UpdateThrottleChanged, LibGroup.Setup.OnSlideGroupDistanceCacheUpdate, LibGroup.Setup.OnSlideGroupDistanceSearchUpdate, LibGroup.Setup.OnSlideGroupUpdateDelay, MapPin.onRotateSlide, MapPin.onScaleSlide, Obsidian.Setup.Castbar.OnSlideBackgroundAlpha, Obsidian.Setup.Castbar.OnSlideFillAlpha, Obsidian.Setup.Castbar.OnSlideGlobalCooldownBackgroundAlpha, Obsidian.Setup.Castbar.OnSlideGlobalCooldownDividerAlpha, Obsidian.Setup.Castbar.OnSlideGlobalCooldownFillAlpha, Obsidian.Setup.Castbar.OnSlideGlobalCooldownFillReadyAlpha, Obsidian.Setup.Castbar.OnSlideGlobalCooldownSparkAlpha, Obsidian.Setup.Castbar.OnSlideIconAlpha, Obsidian.Setup.Castbar.OnSlideIconScale, Obsidian.Setup.Castbar.OnSlideLatencyAlpha, Obsidian.Setup.Castbar.OnSlideLatencyTextAlpha, Obsidian.Setup.Castbar.OnSlideLatencyTextScale, Obsidian.Setup.Castbar.OnSlideNameAlpha, Obsidian.Setup.Castbar.OnSlideNameScale, Obsidian.Setup.Castbar.OnSlideNameWidth, Obsidian.Setup.Castbar.OnSlideSparkAlpha, Obsidian.Setup.Castbar.OnSlideTimerAlpha, Obsidian.Setup.Castbar.OnSlideTimerDecimals, Obsidian.Setup.Castbar.OnSlideTimerScale, Obsidian.Setup.EffectTracker.OnSlideBackgroundAlpha, Obsidian.Setup.EffectTracker.OnSlideFillAlpha, Obsidian.Setup.EffectTracker.OnSlideNameAlpha, Obsidian.Setup.EffectTracker.OnSlideNameScale, Obsidian.Setup.EffectTracker.OnSlideNameWidth, Obsidian.Setup.EffectTracker.OnSlideTimerAlpha, Obsidian.Setup.EffectTracker.OnSlideTimerScale, Obsidian.Setup.EffectTracker.OnSlideTrackerIconAlpha, Obsidian.Setup.EffectTracker.OnSlideTrackerIconScale, Obsidian.Setup.EffectTracker.OnSlideTrackerMaximumEffects, PotionBarSettings.OnAlphaSliderChanged, PotionBarSettings.OnScaleSliderChanged, RVAPI_ColorDialog.OnSlide, RVAPI_Range.OnSlideMapDelay, RVMOD_Manager.OnSlideFadeInOutDelay, RVMOD_Manager.OnSlideScale, RVMOD_Manager.OnSlideZoomInOutDelay, RVMOD_SquaredDistances.OnSlideDistancesDelay, RVMOD_SquaredDistances.OnSlideScale, ReliquaryHunterOptionsWindow.OnWorldMapWindowBackgroundOpacitySlide, ReliquaryHunterOptionsWindow.OnWorldMapWindowOpacitySlide, RoR_SoR.OnSlideWindowOptionsOffset, RoR_SoR.OnSlideWindowOptionsOpacity, RoR_SoR.OnSlideWindowOptionsScale, ScenarioGroupWindow.OnOpacitySlide, ShiniesConfigGeneral.OnSlide_UIScale, Statdoll.Options.onAlphaSlide, Statdoll.Options.onScaleSlide, TacticSetNames.SettingsWindow.OnSlide, TastyButtonsOptions.OnSlidingCeateView, TastyButtonsOptions.OnSlidingEditView, TexturedButtons.Setup.Cooldown.OnSlideCooldownAlpha, TurretRange.Setup.Display.OnSlideDistanceScale, TurretRange.Setup.General.OnSlideUpdateDelay, WSCT.SliderOnSlide, WarBoard.Options.OnSlide, XpStatus.OnOpacitySlide | `pos` | LOW |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | alertMod.OnSliderChanged | `flags, x, y` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | WSCT.OnMouseOver | `` |  |

### Per-Event Lua API Calls

**OnSlide** handlers call: `ButtonGetPressedFlag`, `ComboBoxGetSelectedMenuItem`, `LabelSetText`, `LabelSetTextColor`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowGetAlpha`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScale`, `WindowSetAlpha`, `WindowSetDimensions`, `WindowSetScale`, `WindowSetTintColor`, `WindowStartAlphaAnimation`

**OnMouseOver** handlers call: `WindowGetParent`

## Common Inherits

- Aura_Default_SliderBar
- EA_Default_SliderBar
- HGG_SliderFullLabels
- RVAPI_ColorDialogSliderTemplate
- RVAPI_RangeSliderTemplate
- RVMOD_ManagerSliderTemplate
- RVMOD_SquaredDistancesSliderTemplate
- TAtlasSlider

## Common Parent Elements

- [Windows](element_Windows.md) — 225× (HIGH)
- [Window](element_Window.md) — 2× (LOW)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 219× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 205× (HIGH)
- [Size](element_Size.md) — 186× (HIGH)
- [Windows](element_Windows.md) — 6× (MEDIUM)

## Common Template Bases

- Aura_Default_SliderBar
- EA_Default_SliderBar
- HGG_SliderFullLabels
- RVAPI_ColorDialogSliderTemplate
- RVAPI_RangeSliderTemplate
- RVMOD_ManagerSliderTemplate
- RVMOD_SquaredDistancesSliderTemplate
- TAtlasSlider


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 95% | EA_Default_SliderBar, TAtlasSlider, Aura_Default_SliderBar, HGG_SliderFullLabels, ... |
| `numticks` | optional | 34% | 10, 5, 360, 250, ... |
| `locktoticks` | optional | 12% | false |
| `handleinput` | optional | 5% | true |
| `background` | optional | 4% | EA_BrownHorizontalRule |
| `sliderbutton` | optional | 4% | Aura_Default_SliderButton, EA_Default_SliderButton, DetauntHelper_Config_SB |
| `tickmark` | optional | 4% | Aura_Default_SliderTick, EA_Default_SliderTick2, DetauntHelper_Config_ST |
| `autoresize` | optional | 2% | true |
| `handlinput` | optional | 2% | true |
| `numTicks` | optional | 1% | 11 |
| `scale` | optional | 1% | 0.4 |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 219 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 205 times as an unnamed child.

### [Size](element_Size.md)

Observed 186 times as an unnamed child.

### [Windows](element_Windows.md)

Observed 6 times as an unnamed child.

## Recursive Hierarchy

- Root: [SliderBar](element_SliderBar.md)
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
| `SliderBarGetCurrentPosition` | ui | 325 | OnSlide |
| `LabelSetText` | ui | 123 | OnSlide |
| `TextEditBoxSetText` | ui | 55 | OnSlide |
| `TextEditBoxGetText` | ui | 36 | OnSlide |
| `ComboBoxGetSelectedMenuItem` | ui | 28 | OnSlide |
| `WindowSetTintColor` | ui | 18 | OnSlide |
| `SliderBarSetCurrentPosition` | ui | 15 | OnSlide |
| `ButtonGetPressedFlag` | ui | 12 | OnSlide |
| `LabelSetTextColor` | ui | 9 | OnSlide |
| `WindowSetScale` | ui | 7 | OnSlide |
| `WindowGetDimensions` | ui | 5 | OnSlide |
| `WindowSetAlpha` | ui | 5 | OnSlide |
| `WindowGetParent` | ui | 4 | OnMouseOver, OnSlide |
| `WindowGetScale` | ui | 2 | OnSlide |
| `WindowSetDimensions` | ui | 2 | OnSlide |
| `WindowGetAlpha` | ui | 1 | OnSlide |
| `WindowGetId` | ui | 1 | OnSlide |
| `WindowStartAlphaAnimation` | ui | 1 | OnSlide |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMouseOver

Confidence: HIGH

### OnSlide

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `pos` | unknown | unknown |
## Lua Functions Manipulating This Type

- MoraleCircle.OnSetCustomColorFull
- WSCT.OnSetCustomColor
- MoraleCircle.OnSetCustomColor
- MoraleCircle.ColorChanger2
- ChattyCathy.CC_ModelWindow
- RoR_SoR.OnWindowOptionsSetOpacity
- CCTV.Initialize
- MoraleCircle.ColorChanger1
- MoraleCircle.ColorChanger3
- ScenarioGroupWindow.OnMenuClickSetBackgroundOpacity
- WSCT.ColorOnButtonUp
- ChattyCathy.OnSlideAlpha
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- MoraleCircle.ColorChanger4
- RoR_SoR.OnWindowOptionsSetScale
- XpStatus.InitializeSetOpacityWindow
- BloodyMess.ApplyButton_OnLButtonUp
- BloodyMess.OnInitialize
- ChattyCathy.CancelOptions
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- MoraleCircle.OnSetCustomColorEmpty
- RoR_SoR.OnWindowOptionsSetOffset
- CCTV.UpdateSettings
- ChattyCathy.ApplyOptions
- MoraleCircle.OnSetCustomColorFill
- WSCT.OnLButtonUpColorPicker
- XpStatus.EditOpacity


## Binding Resolution

- Total handler declarations: 206
- Resolved to Lua functions: 206 (100%)

## Seen In

- AdjustTheTip
- Atlas
- Aura
- Bloody Mess
- BuffHead
- CCTV
- CDown
- ChattyCathy
- DAoCBuff
- DetauntHelper
- DuffTimer
- EA_ScenarioGroupWindow
- Enemy
- Group Icons SG
- GroupRange
- GroupSpotter
- HealGrid
- LibGroup
- MapMonster
- MapPin
- MoraleCircle
- Obsidian
- PotionBar
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_Manager
- RVMOD_SquaredDistances
- ReliquaryHunter
- RoR_SoR
- SNT_INFO
- SNT_PANEL
- SOR
- Shinies
- Statdoll Remix
- TacticSetNames
- TastyButtons
- TexturedButtons
- TidyChat
- TurretRange
- WSCT
- WarBoard
- WarBoard_FPS
- XpStatus+G
- alertMod

## Examples

- AdjustTheTip: AdjustTheTipMenuItemSliderSliderBar -> SliderBar AdjustTheTipMenuItemSliderSliderBar
- Atlas: AtlasConfigurationFrameContentMapSettingsTransparencyOnMouseOverSlider -> SliderBar AtlasConfigurationFrameContentMapSettingsTransparencyOnMouseOverSlider
- Atlas: AtlasConfigurationFrameContentMapSettingsTransparencyOnMovementSlider -> SliderBar AtlasConfigurationFrameContentMapSettingsTransparencyOnMovementSlider
- Atlas: AtlasConfigurationFrameContentMapSettingsTransparencySlider -> SliderBar AtlasConfigurationFrameContentMapSettingsTransparencySlider
- Atlas: TAtlasSlider -> SliderBar TAtlasSlider
- Aura: AuraColorPickerAlpha -> SliderBar AuraColorPickerAlpha

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [EA_Default_SliderBar](../../globals/constants/constant_EA_Default_SliderBar.md) (HIGH 100/100) - Constant
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnSlide](../handlers/handler_OnSlide.md) (HIGH 88/100) - XML Event
