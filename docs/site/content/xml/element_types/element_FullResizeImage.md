# FullResizeImage

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
| Addons seen in | AggroMeter, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy, Killer, LibWBToggler |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTooltip.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0` |
| Namespaces detected | FullResizeImage |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AggroMeter: AggroMeterWindowBorderCheck, AggroMeter: AggroMeterWindow_AggroWindow1BorderCheck, AggroMeter: AggroMeterWindow_AggroWindow1Seperator1, AggroMeter: AggroMeterWindow_AggroWindow2BorderCheck, AggroMeter: AggroMeterWindow_AggroWindow3BorderCheck, AggroMeter: AggroMeterWindow_AggroWindow4BorderCheck |
| XML usage count | 158 |
| XML attribute usage count | 158 |
| Lua usage count | 1 |
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

Observed XML element type instantiated by 22 addons.

## Common Attributes

- name
- inherits
- handleinput
- layer
- alpha
- texture
- skipinput
- showing
- drawchildrenfirst
- font
- frameonly

## Common Handlers

- [OnMouseOver](../handlers/handler_OnMouseOver.md)

## Common Handler Functions

- Enemy.ConfigurationWindow_ShowTooltip


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.ConfigurationWindow_ShowTooltip | `function()` | MEDIUM |

## Common Inherits

- EA_FullResizeImage_TintableSolidBackground
- EA_FullResizeImage_DefaultFrame
- AuraWindowBackground
- EA_FullResizeImage_WhiteTransparent
- EA_Button_ResizeIconFrameNormal
- EA_FullResizeImage_BlackTransparent
- DefaultWindowBackground
- EA_FullResizeImage_MetalFill
- EA_FullResizeImage_TintableFrame
- ClosetGoblinDefaultBG
- ModBackgroundTemplate
- SettingsSectionBackground

## Common Parent Elements

- [Windows](element_Windows.md) — 156× (HIGH)
- [Window](element_Window.md) — 2× (LOW)

## Common Structural Child Elements

- [Size](element_Size.md) — 35× (HIGH)
- [TintColor](element_TintColor.md) — 28× (HIGH)
- [Sizes](element_Sizes.md) — 14× (HIGH)
- [TexCoords](element_TexCoords.md) — 9× (MEDIUM)
- [TexSlices](element_TexSlices.md) — 5× (MEDIUM)
- [TexDims](element_TexDims.md) — 3× (MEDIUM)
- [Windows](element_Windows.md) — 1× (LOW)

## Common Template Bases

- AuraWindowBackground
- ClosetGoblinDefaultBG
- DefaultWindowBackground
- EA_Button_ResizeIconFrameNormal
- EA_FullResizeImage_BlackTransparent
- EA_FullResizeImage_DefaultFrame
- EA_FullResizeImage_MetalFill
- EA_FullResizeImage_TintableFrame
- EA_FullResizeImage_TintableSolidBackground
- EA_FullResizeImage_WhiteTransparent
- ModBackgroundTemplate
- SettingsSectionBackground


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<FullResizeImage handleinput="false" name="..." texture="enemy_bar_rect">
  <TexDims x="136" y="36"/>
  <Size/>
  <Sizes>
    <Middle x="2" y="2"/>
  </Sizes>
  <TexCoords/>
</FullResizeImage>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 90% | EA_FullResizeImage_BlackTransparent, EA_FullResizeImage_TintableSolidBackground, EA_FullResizeImage_WhiteTransparent, EA_FullResizeImage_DefaultFrame, ... |
| `handleinput` | optional | 55% | false, true |
| `layer` | optional | 32% | background, popup, overlay, secondary, ... |
| `alpha` | optional | 24% | 0.5, 0.3, 0.1, 0.75, ... |
| `texture` | optional | 8% | shared_01, enemy_bar_rect, EA_HUD_01, Frame_1, ... |
| `skipinput` | optional | 6% | true |
| `showing` | optional | 3% | false |
| `drawchildrenfirst` | optional | 0% | true |
| `font` | optional | 0% | font_clear_small_bold |
| `frameonly` | optional | 0% | true |
## Structural Sub-Elements

### [Size](element_Size.md)

Observed 35 times as an unnamed child.

### [TintColor](element_TintColor.md)

Observed 28 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 110, 130, 50 |
| `g` | **required** | 0, 200, 185, 130 |
| `r` | **required** | 255, 0, 180, 200 |
| `a` | optional | 30, 255, 0.5, 200 |
### [Sizes](element_Sizes.md)

Observed 14 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `middle` | optional | 30, 200, 61, 50 |
| `left` | optional | 0, 7 |
| `right` | optional | 0 |
| `bottom` | optional | 0 |
| `top` | optional | 0 |
### [TexCoords](element_TexCoords.md)

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 0, 133, 88, 32 |
| `y` | optional | 0, 163, 51, 32 |
### [TexSlices](element_TexSlices.md)

Observed 5 times as an unnamed child.

### [TexDims](element_TexDims.md)

Observed 3 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 64, 256, 32, 430 |
| `y` | **required** | 64, 256, 32, 430 |
### [Windows](element_Windows.md)

Observed 1 times as an unnamed child.

## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnMouseOver

Confidence: HIGH

## Lua Functions Manipulating This Type

- Enemy.EnemyUnitFramePart:BoundingBoxSetShowing
- Enemy._Initialize
- WHMCore.ApplyBackgroundFillColor
- Enemy.EnemyEffectsIndicator:BoundingBoxSetShowing
- MoraleCircle.OnSetCustomColor
- MoraleCircle.ColorChanger1
- MoraleCircle.OnSetCustomColorEmpty
- MoraleCircle.ColorChanger2
- MoraleCircle.ColorChanger3
- WSCT.OnLButtonUpColorPicker
- WHMGui.RefreshConfigurationWindow
- Enemy.EnemyUnitFrame:BoundingBoxSetShowing
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- Killer.ShowPersonalStatsTooltip
- Killer.ShowTopKillersTooltip
- MoraleCircle.OnSetCustomColorFill
- MoraleCircle.OnSetCustomColorFull
- MoraleCircle.ColorChanger4
- Enemy.MarksInitialize
- Enemy.CombatLogUI_IDS_Initialize
- Killer.Initialize
- Killer.ShowRowTooltip
- WSCT.ColorOnButtonUp
- WSCT.OnSetCustomColor


## Binding Resolution

- Total handler declarations: 1
- Resolved to Lua functions: 1 (100%)

## Seen In

- AggroMeter
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Killer
- LibWBToggler
- MiracleGrowLight
- MoraleCircle
- PartyCast
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow

## Examples

- AggroMeter: AggroMeterWindowBorderCheck -> FullResizeImage AggroMeterWindowBorderCheck
- AggroMeter: AggroMeterWindow_AggroWindow1BorderCheck -> FullResizeImage AggroMeterWindow_AggroWindow1BorderCheck
- AggroMeter: AggroMeterWindow_AggroWindow1Seperator1 -> FullResizeImage AggroMeterWindow_AggroWindow1Seperator1
- AggroMeter: AggroMeterWindow_AggroWindow2BorderCheck -> FullResizeImage AggroMeterWindow_AggroWindow2BorderCheck
- AggroMeter: AggroMeterWindow_AggroWindow3BorderCheck -> FullResizeImage AggroMeterWindow_AggroWindow3BorderCheck
- AggroMeter: AggroMeterWindow_AggroWindow4BorderCheck -> FullResizeImage AggroMeterWindow_AggroWindow4BorderCheck

## Related APIs

- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Sizes](element_Sizes.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TexDims](element_TexDims.md) (HIGH 100/100) - XML Element Type
- [TexSlices](element_TexSlices.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event

## Affects

- none
