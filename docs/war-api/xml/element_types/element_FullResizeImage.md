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
| Addons seen in | AggroMeter, Aura, BuffHead, Busted, CM_ClosetGoblin, CMap, Cheeseboard, DAoCBuff |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.xml:158`, `/workspace_addons/AggroMeter/AggroMeter.xml:216`, `/workspace_addons/AggroMeter/AggroMeter.xml:221`, `/workspace_addons/AggroMeter/AggroMeter.xml:259`, `/workspace_addons/AggroMeter/AggroMeter.xml:288`, `/workspace_addons/AggroMeter/AggroMeter.xml:316`, `/workspace_addons/AggroMeter/AggroMeter.xml:344`, `/workspace_addons/AggroMeter/AggroMeter.xml:373` |
| Namespaces detected | FullResizeImage |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AggroMeter: AggroMeterWindowBorderCheck, AggroMeter: AggroMeterWindow_AggroWindow1BorderCheck, AggroMeter: AggroMeterWindow_AggroWindow1Seperator1, AggroMeter: AggroMeterWindow_AggroWindow2BorderCheck, AggroMeter: AggroMeterWindow_AggroWindow3BorderCheck, AggroMeter: AggroMeterWindow_AggroWindow4BorderCheck |
| XML usage count | 199 |
| XML attribute usage count | 199 |
| Lua usage count | 3 |
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

Observed XML element type instantiated by 41 addons.

## Common Attributes

- name
- inherits
- handleinput
- layer
- alpha
- texture
- skipinput
- frameonly
- showing
- hanldeinput
- drawchildrenfirst
- font

## Common Handlers

- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)

## Common Handler Functions

- Enemy.ConfigurationWindow_ShowTooltip
- MapMonster_Calibrate.OnLMouseButton
- MapMonster_Calibrate.OnMouseOverEnd
- MiracleGrow.OverAutoConsume
- MiracleGrow.OverEndStartAll
- MiracleGrow.OverStartAll
- MiracleGrow.autoGrow
- MiracleGrow.onAutoConsumeToggle
- MiracleGrow.setAutoGrowColor
- wbLeadHelper.onMouseOut
- wbLeadHelper.onMouseOver


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | MapMonster_Calibrate.OnLMouseButton, MiracleGrow.autoGrow, MiracleGrow.onAutoConsumeToggle | `flags, x, y` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.ConfigurationWindow_ShowTooltip, MiracleGrow.OverAutoConsume, MiracleGrow.OverStartAll, wbLeadHelper.onMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | MapMonster_Calibrate.OnMouseOverEnd, MiracleGrow.OverEndStartAll, MiracleGrow.setAutoGrowColor, wbLeadHelper.onMouseOut | `function(...)` | LOW |

### Per-Event Lua API Calls

**OnLButtonUp** handlers call: `LabelSetText`, `WindowGetScreenPosition`, `WindowSetTintColor`

**OnMouseOver** handlers call: `LabelSetTextColor`, `WindowSetTintColor`

**OnMouseOverEnd** handlers call: `LabelSetText`, `LabelSetTextColor`, `WindowGetScreenPosition`, `WindowSetTintColor`

## Common Inherits

- EA_FullResizeImage_TintableSolidBackground
- AuraWindowBackground
- EA_FullResizeImage_TanBorder
- EA_FullResizeImage_BlackTransparent
- EA_FullResizeImage_WhiteTransparent
- EA_FullResizeImage_MetalFill
- EA_Button_ResizeIconFrameNormal
- EA_FullResizeImage_TintableFrame
- DefaultWindowBackground
- EA_FullResizeImage_DefaultFrame
- ClosetGoblinDefaultBG
- EA_Button_ResizeIconFrameHighlight

## Common Parent Elements

- [Window](element_Window.md) — 153× (HIGH)
- [Button](element_Button.md) — 5× (MEDIUM)

## Common Named Child Elements

- [Label](element_Label.md) — 3× (MEDIUM)

## Common Structural Child Elements

- [TintColor](element_TintColor.md) — 37× (HIGH)
- [Middle](element_Middle.md) — 11× (HIGH)
- [TexSlices](element_TexSlices.md) — 5× (MEDIUM)

## Common Template Bases

- AuraWindowBackground
- ClosetGoblinDefaultBG
- DefaultWindowBackground
- EA_Button_ResizeIconFrameHighlight
- EA_Button_ResizeIconFrameNormal
- EA_Button_ResizeIconFramePressed
- EA_FullResizeImage_AlertBubbleBackground
- EA_FullResizeImage_BlackTransparent
- EA_FullResizeImage_DefaultBubbleBackground
- EA_FullResizeImage_DefaultFrame
- EA_FullResizeImage_MetalFill
- EA_FullResizeImage_TanBorder
- EA_FullResizeImage_TintableFrame
- EA_FullResizeImage_TintableSolidBackground
- EA_FullResizeImage_WhiteTransparent
- IraConfigUFrame
- ModBackgroundTemplate
- SettingsSectionBackground


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 52% | EA_FullResizeImage_TintableSolidBackground, EA_FullResizeImage_WhiteTransparent, EA_FullResizeImage_TintableFrame, EA_Button_ResizeIconFrameNormal, ... |
| `handleinput` | optional | 28% | true, false |
| `layer` | optional | 17% | secondary, overlay, default, background, ... |
| `alpha` | optional | 14% | 0, 0.25, 0.5, 0.6, ... |
| `texture` | optional | 3% | Shinies3pxBorder, Frame_BG, enemy_bar_rect, Frame_1, ... |
| `skipinput` | optional | 2% | true |
| `frameonly` | optional | 1% | true |
| `showing` | optional | 1% | false |
| `hanldeinput` | optional | 0% | false |
| `drawchildrenfirst` | optional | 0% | true |
| `font` | optional | 0% | font_clear_small_bold |
| `sticky` | optional | 0% | false |
| `textalign` | optional | 0% | center |
## Structural Sub-Elements

### [TintColor](element_TintColor.md)

Observed 37 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** |  |
| `g` | **required** |  |
| `r` | **required** |  |
| `a` | optional |  |
### [Middle](element_Middle.md)

Observed 11 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
| `id` | optional |  |
### [TexSlices](element_TexSlices.md)

Observed 5 times as an unnamed child.

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowSetTintColor` | ui | 10 | OnLButtonUp, OnMouseOver, OnMouseOverEnd |
| `LabelSetText` | ui | 7 | OnLButtonUp, OnMouseOverEnd |
| `LabelSetTextColor` | ui | 2 | OnMouseOver, OnMouseOverEnd |
| `WindowGetScreenPosition` | ui | 2 | OnLButtonUp, OnMouseOverEnd |
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

### OnMouseOverEnd

Confidence: HIGH

## Lua Functions Manipulating This Type

- WhoHealedMe.WHMCore.ApplyBackgroundFillColor
- MoraleCircle.MoraleCircle.OnSetCustomColor
- WSCT.WSCT.ColorOnButtonUp
- MoraleCircle.MoraleCircle.ColorChanger4
- WhoHealedMe.WHMGui.RefreshConfigurationWindow
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- Enemy.Enemy.MarksInitialize
- Killer.Killer.Initialize
- MoraleCircle.MoraleCircle.ColorChanger1
- MoraleCircle.MoraleCircle.OnSetCustomColorFull
- MoraleCircle.MoraleCircle.ColorChanger3
- Moth.Moth.HealthBar
- Enemy.EnemyUnitFramePart:BoundingBoxSetShowing
- RandomMount.RandomMountUI.OnInitialize
- Killer.Killer.ShowTopKillersTooltip
- WSCT.WSCT.OnSetCustomColor
- Enemy.Enemy._Initialize
- Moth.Moth.HideBorders
- WSCT.WSCT.OnLButtonUpColorPicker
- Enemy.EnemyEffectsIndicator:BoundingBoxSetShowing
- MoraleCircle.MoraleCircle.OnSetCustomColorEmpty
- Enemy.Enemy.CombatLogUI_IDS_Initialize
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Moth.Moth.UpdateHealthBar
- Killer.Killer.ShowRowTooltip
- Killer.Killer.ShowPersonalStatsTooltip
- Moth.Moth.Clear
- MoraleCircle.MoraleCircle.ColorChanger2
- MoraleCircle.MoraleCircle.OnSetCustomColorFill
- Enemy.EnemyUnitFrame:BoundingBoxSetShowing


## Binding Resolution

- Total handler declarations: 11
- Resolved to Lua functions: 11 (100%)

## Seen In

- AggroMeter
- Aura
- BuffHead
- Busted
- CM_ClosetGoblin
- CMap
- Cheeseboard
- DAoCBuff
- EA_UiDebugTools
- EA_UiModWindow
- Effigy
- Enemy
- Killer
- LibWBToggler
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- Moth
- PotionBar
- QuickTacticSwitch
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Targets
- TexturedButtons
- TidyChat
- TidyRoll
- Tortall_DPS
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- wbLeadHelper

## Examples

- AggroMeter: AggroMeterWindowBorderCheck -> FullResizeImage AggroMeterWindowBorderCheck
- AggroMeter: AggroMeterWindow_AggroWindow1BorderCheck -> FullResizeImage AggroMeterWindow_AggroWindow1BorderCheck
- AggroMeter: AggroMeterWindow_AggroWindow1Seperator1 -> FullResizeImage AggroMeterWindow_AggroWindow1Seperator1
- AggroMeter: AggroMeterWindow_AggroWindow2BorderCheck -> FullResizeImage AggroMeterWindow_AggroWindow2BorderCheck
- AggroMeter: AggroMeterWindow_AggroWindow3BorderCheck -> FullResizeImage AggroMeterWindow_AggroWindow3BorderCheck
- AggroMeter: AggroMeterWindow_AggroWindow4BorderCheck -> FullResizeImage AggroMeterWindow_AggroWindow4BorderCheck

## Related APIs

- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function

## Used With

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOverEnd](../../events/window_events/window_event_OnMouseOverEnd.md) (HIGH 100/100) - Window Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none
