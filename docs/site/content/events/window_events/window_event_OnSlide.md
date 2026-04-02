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
| Addons seen in | Aura, BuffHead, Enemy, LibGroup, MapMonster, MapPin, MoraleCircle, PotionBar |
| Files seen in | `/workspace_addons/Aura/Source/AuraColorPicker.xml:105`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:33`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:65`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:85`, `/workspace_addons/Aura/Source/AuraConfig.xml:1680`, `/workspace_addons/Aura/Source/AuraConfig.xml:256`, `/workspace_addons/Aura/Source/AuraConfig.xml:311`, `/workspace_addons/BuffHead/Setup/General.xml:142` |
| Namespaces detected | OnSlide |
| Source kinds | event_page, xml_handlers |
| Example locations | Aura: AuraColorPickerAlpha.OnSlide, Aura: AuraColorPickerBlue.OnSlide, Aura: AuraColorPickerGreen.OnSlide, Aura: AuraColorPickerRed.OnSlide, Aura: AuraConfigGeneralTextureRotationSlider.OnSlide, Aura: AuraConfigGeneralTextureScaleSlider.OnSlide |
| XML usage count | 83 |
| XML attribute usage count | 83 |
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

Observed as an engine-supplied UI event hook used by 16 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- Aura
- BuffHead
- Enemy
- LibGroup
- MapMonster
- MapPin
- MoraleCircle
- PotionBar
- RVAPI_ColorDialog
- RVMOD_Manager
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
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- LibGroup.Setup.OnSlideGroupDistanceCacheUpdate
- LibGroup.Setup.OnSlideGroupDistanceSearchUpdate
- LibGroup.Setup.OnSlideGroupUpdateDelay
- MapMonster.PinTypeEditor.OnSetCustomColor
- MapPin.onRotateSlide
- MapPin.onScaleSlide
- MoraleCircle.OnSetCustomColor
- MoraleCircle.OnSetCustomColorEmpty
- MoraleCircle.OnSetCustomColorFill
- MoraleCircle.OnSetCustomColorFull
- PotionBarSettings.OnAlphaSliderChanged
- PotionBarSettings.OnScaleSliderChanged
- RVAPI_ColorDialog.OnSlide
- RVMOD_Manager.OnSlideFadeInOutDelay
- RVMOD_Manager.OnSlideScale
- RVMOD_Manager.OnSlideZoomInOutDelay
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

- [OnSlide](../../xml/handlers/handler_OnSlide.md) (HIGH 100/100) - XML Handler
- [SliderBar](../../xml/element_types/element_SliderBar.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- none
