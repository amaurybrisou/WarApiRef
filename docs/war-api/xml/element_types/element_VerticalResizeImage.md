# VerticalResizeImage

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, CMap, EA_UiDebugTools |
| Files seen in | `/workspace_addons/BuffHead/Setup/SetupLayout.xml:67`, `/workspace_addons/BuffHead/Setup/SetupLayout.xml:73`, `/workspace_addons/BuffHead/Setup/SetupLayout.xml:76`, `/workspace_addons/BuffHead/Setup/SetupLayout.xml:79`, `/workspace_addons/cmap/Cmap_Templates_OverheadMap.xml:12`, `/workspace_addons/cmap/Cmap_Templates_OverheadMap.xml:26`, `/workspace_addons/ea_uidebugtools/Source/DebugWindowVerticalScroll.xml:80` |
| Namespaces detected | VerticalResizeImage |
| Source kinds | xml_frames |
| Example locations | BuffHead: BuffHeadLayoutVerticalButtonHighlight, BuffHead: BuffHeadLayoutVerticalButtonNormal, BuffHead: BuffHeadLayoutVerticalButtonPressed, BuffHead: BuffHeadLayoutVerticalResizeImage, CMap: HUDInfluenceBarBackground, CMap: HUDInfluenceBarFill |
| XML usage count | 7 |
| XML attribute usage count | 7 |
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

Observed XML element type instantiated by 3 addons.

## Common Attributes

- name
- texture
- inherits
- handleinput
- reverseFill
- reversefill
- sticky

## Common Inherits

- BuffHeadLayoutVerticalResizeImage

## Common Parent Elements

- [Window](element_Window.md)


## Structural Sub-Elements

### [Middle](element_Middle.md)

- Observed in 4 parent frames
- Attributes: `x`, `y`
  - `x`: `424`, `576`, `598`, `172`
  - `y`: `762`, `1`, `416`

### [TintColor](element_TintColor.md)

- Observed in 3 parent frames
- Attributes: `b`, `g`, `r`
  - `b`: `0`, `200`
  - `g`: `200`, `255`
  - `r`: `200`, `255`

### [Bottom](element_Bottom.md)

- Observed in 1 parent frames
- Attributes: `x`, `y`
  - `x`: `172`
  - `y`: `429`

### [Top](element_Top.md)

- Observed in 1 parent frames
- Attributes: `x`, `y`
  - `x`: `172`
  - `y`: `403`

## Typical XML Structure

```xml
<VerticalResizeImage name="..." inherits="BuffHeadLayoutVerticalResizeI...">
  <TintColor r="200" g="200" b="0"/>
</VerticalResizeImage>
```


## Attribute Reference

| Attribute | Role | Observed Values | Count |
|-----------|------|-----------------|-------|
| `name` | string | `BuffHeadLayoutVerticalButtonHighlight`, `BuffHeadLayoutVerticalButtonNormal`, `BuffHeadLayoutVerticalButtonPressed`, `BuffHeadLayoutVerticalResizeImage`, … | 7 |
| `texture` | string | `EA_TintableImage`, `ui_main`, `EA_HUD_01` | 4 |
| `inherits` | frame-ref | `BuffHeadLayoutVerticalResizeImage` | 3 |
| `handleinput` | boolean | `false` | 2 |
| `reverseFill` | boolean | `true` | 1 |
| `reversefill` | boolean | `true` | 1 |
| `sticky` | boolean | `false` | 1 |

## Seen In

- BuffHead
- CMap
- EA_UiDebugTools

## Examples

- BuffHead: BuffHeadLayoutVerticalButtonHighlight -> VerticalResizeImage BuffHeadLayoutVerticalButtonHighlight
- BuffHead: BuffHeadLayoutVerticalButtonNormal -> VerticalResizeImage BuffHeadLayoutVerticalButtonNormal
- BuffHead: BuffHeadLayoutVerticalButtonPressed -> VerticalResizeImage BuffHeadLayoutVerticalButtonPressed
- BuffHead: BuffHeadLayoutVerticalResizeImage -> VerticalResizeImage BuffHeadLayoutVerticalResizeImage
- CMap: HUDInfluenceBarBackground -> VerticalResizeImage HUDInfluenceBarBackground
- CMap: HUDInfluenceBarFill -> VerticalResizeImage HUDInfluenceBarFill

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
