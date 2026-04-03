# OnSlide

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 118

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Aura, BuffHead, DAoCBuff, Enemy, LibGroup, MoraleCircle, PotionBar, RoR_SoR |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraColorPicker.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupContainer.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupDisplay.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupGeneral.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayoutProperties.xml:0` |
| Namespaces detected | OnSlide |
| Source kinds | event_page, xml_handlers |
| Example locations | Aura: AuraColorPickerAlpha.OnSlide, Aura: AuraColorPickerBlue.OnSlide, Aura: AuraColorPickerGreen.OnSlide, Aura: AuraColorPickerRed.OnSlide, Aura: AuraConfigGeneralTextureRotationSlider.OnSlide, Aura: AuraConfigGeneralTextureScaleSlider.OnSlide |
| XML usage count | 80 |
| XML attribute usage count | 80 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
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

Observed as an engine-supplied UI event hook used by 13 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- Aura
- BuffHead
- DAoCBuff
- Enemy
- LibGroup
- MoraleCircle
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TurretRange
- WSCT
- WarBoard

## Registrars And Handlers

- AuraColorPicker.OnSlide
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
- BuffHead.Setup.SelectColor.OnSlideTint
- DAoCBuffSettings.FilterSettings.OnSlideBB
- DAoCBuffSettings.FilterSettings.OnSlideBG
- DAoCBuffSettings.FilterSettings.OnSlideBR
- DAoCBuffSettings.FilterSettings.OnSlideCB
- DAoCBuffSettings.FilterSettings.OnSlideCG
- DAoCBuffSettings.FilterSettings.OnSlideCR
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- LibGroup.Setup.OnSlideGroupDistanceCacheUpdate
- LibGroup.Setup.OnSlideGroupDistanceSearchUpdate
- LibGroup.Setup.OnSlideGroupUpdateDelay
- MoraleCircle.OnSetCustomColor
- MoraleCircle.OnSetCustomColorEmpty
- MoraleCircle.OnSetCustomColorFill
- MoraleCircle.OnSetCustomColorFull
- PotionBarSettings.OnAlphaSliderChanged
- PotionBarSettings.OnScaleSliderChanged
- RoR_SoR.OnSlideWindowOptionsOffset
- RoR_SoR.OnSlideWindowOptionsOpacity
- RoR_SoR.OnSlideWindowOptionsScale
- ShiniesConfigGeneral.OnSlide_UIScale
- TexturedButtons.Setup.Cooldown.OnSlideCooldownAlpha
- TexturedButtons.Setup.SelectColor.OnSlideTint
- TexturedButtons.Setup.Tint.OnSlideTint
- TurretRange.Setup.Display.OnSlideDistanceScale
- TurretRange.Setup.Display.OnSlideTint
- TurretRange.Setup.Distance.OnSlideTint
- TurretRange.Setup.General.OnSlideUpdateDelay
- WSCT.OnSetCustomColor
- WSCT.SliderOnSlide
- WarBoard.Options.OnSlide

## Examples

- Aura: AuraColorPickerAlpha -> AuraColorPickerAlpha.OnSlide -> AuraColorPicker.OnSlide
- Aura: AuraColorPickerBlue -> AuraColorPickerBlue.OnSlide -> AuraColorPicker.OnSlide
- Aura: AuraColorPickerGreen -> AuraColorPickerGreen.OnSlide -> AuraColorPicker.OnSlide
- Aura: AuraColorPickerRed -> AuraColorPickerRed.OnSlide -> AuraColorPicker.OnSlide
- Aura: AuraConfigGeneralTextureRotationSlider -> AuraConfigGeneralTextureRotationSlider.OnSlide -> AuraConfig.OnTextureRotationSlide
- Aura: AuraConfigGeneralTextureScaleSlider -> AuraConfigGeneralTextureScaleSlider.OnSlide -> AuraConfig.OnTextureScaleSlide

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
