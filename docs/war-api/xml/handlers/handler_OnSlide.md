# OnSlide

- Type: XML Handler
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
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Aura, BuffHead, Enemy, LibGroup, MapMonster, MapPin, MoraleCircle, PotionBar |
| Files seen in | `/workspace_addons/Aura/Source/AuraColorPicker.xml:105`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:33`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:65`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:85`, `/workspace_addons/Aura/Source/AuraConfig.xml:1680`, `/workspace_addons/Aura/Source/AuraConfig.xml:256`, `/workspace_addons/Aura/Source/AuraConfig.xml:311`, `/workspace_addons/BuffHead/Setup/General.xml:142` |
| Namespaces detected | OnSlide |
| Source kinds | bindings, xml_handlers |
| Example locations | Aura: AuraColorPickerAlpha.OnSlide, Aura: AuraColorPickerBlue.OnSlide, Aura: AuraColorPickerGreen.OnSlide, Aura: AuraColorPickerRed.OnSlide, Aura: AuraConfigGeneralTextureRotationSlider.OnSlide, Aura: AuraConfigGeneralTextureScaleSlider.OnSlide |
| XML usage count | 83 |
| XML attribute usage count | 83 |
| Lua usage count | 83 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
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

Observed as an XML handler hook bound by 16 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- SliderBar

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

- [OnSlide](../../events/window_events/window_event_OnSlide.md) (HIGH 100/100) - Window Event
- [SliderBar](../element_types/element_SliderBar.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
