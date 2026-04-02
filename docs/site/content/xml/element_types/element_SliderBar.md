# SliderBar

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
| Addons seen in | Aura, BuffHead, Enemy, LibGroup, MapMonster, MapPin, MoraleCircle, PotionBar |
| Files seen in | `/workspace_addons/Aura/Source/AuraColorPicker.xml:26`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:58`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:78`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:98`, `/workspace_addons/Aura/Source/AuraConfig.xml:1670`, `/workspace_addons/Aura/Source/AuraConfig.xml:246`, `/workspace_addons/Aura/Source/AuraConfig.xml:301`, `/workspace_addons/Aura/Source/Templates.xml:317` |
| Namespaces detected | SliderBar |
| Source kinds | xml_frames, xml_handlers |
| Example locations | Aura: AuraColorPickerAlpha, Aura: AuraColorPickerBlue, Aura: AuraColorPickerGreen, Aura: AuraColorPickerRed, Aura: AuraConfigGeneralTextureRotationSlider, Aura: AuraConfigGeneralTextureScaleSlider |
| XML usage count | 91 |
| XML attribute usage count | 91 |
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

Observed XML element type instantiated by 17 addons.

## Common Attributes

- name
- inherits
- numticks
- handleinput
- scale
- locktoticks
- autoresize
- handlinput
- numTicks
- background
- sliderbutton
- tickmark

## Common Handlers

- OnSlide
- OnMouseOver

## Common Inherits

- EA_Default_SliderBar
- Aura_Default_SliderBar
- RVMOD_ManagerSliderTemplate
- RVAPI_ColorDialogSliderTemplate

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
- TidyChat
- TurretRange
- WSCT
- WarBoard

## Examples

- Aura: AuraColorPickerAlpha -> SliderBar AuraColorPickerAlpha
- Aura: AuraColorPickerBlue -> SliderBar AuraColorPickerBlue
- Aura: AuraColorPickerGreen -> SliderBar AuraColorPickerGreen
- Aura: AuraColorPickerRed -> SliderBar AuraColorPickerRed
- Aura: AuraConfigGeneralTextureRotationSlider -> SliderBar AuraConfigGeneralTextureRotationSlider
- Aura: AuraConfigGeneralTextureScaleSlider -> SliderBar AuraConfigGeneralTextureScaleSlider

## Related APIs

- none

## Used With

- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnSlide](../../events/window_events/window_event_OnSlide.md) (HIGH 100/100) - Window Event
- [OnSlide](../handlers/handler_OnSlide.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none
