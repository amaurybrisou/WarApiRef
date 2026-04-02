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

| Event | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | MapMonster_Calibrate.OnLMouseButton, MiracleGrow.autoGrow, MiracleGrow.onAutoConsumeToggle | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | Enemy.ConfigurationWindow_ShowTooltip, MiracleGrow.OverAutoConsume, MiracleGrow.OverStartAll, wbLeadHelper.onMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | MapMonster_Calibrate.OnMouseOverEnd, MiracleGrow.OverEndStartAll, MiracleGrow.setAutoGrowColor, wbLeadHelper.onMouseOut | `function(...)` | LOW |

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

- [Window](element_Window.md)
- [Button](element_Button.md)

## Common Named Child Elements

- [Label](element_Label.md)


## Structural Sub-Elements

### [TintColor](element_TintColor.md)

- Observed in 37 parent frames
- Attributes: `a`, `b`, `g`, `r`
  - `a`: `30`, `255`, `200`, `0.6`, `1`, `0.4`, `0`, `0.65`
  - `b`: `0`, `110`, `130`, `150`, `50`, `120`, `166`, `36`
  - `g`: `0`, `185`, `130`, `150`, `50`, `120`, `84`, `28`
  - `r`: `0`, `200`, `130`, `150`, `50`, `120`, `237`, `255`

### [Middle](element_Middle.md)

- Observed in 11 parent frames
- Attributes: `x`, `y`
  - `x`: `78`, `2`, `1`, `13`, `50`, `58`
  - `y`: `25`, `2`, `1`, `13`, `23`, `58`

### [TexSlices](element_TexSlices.md)

- Observed in 5 parent frames


## Attribute Reference

| Attribute | Role | Observed Values | Count |
|-----------|------|-----------------|-------|
| `name` | string | `AggroMeterWindowBorderCheck`, `AggroMeterWindow_AggroWindow1BorderCheck`, `AggroMeterWindow_AggroWindow1Seperator1`, `AggroMeterWindow_AggroWindow2BorderCheck`, … | 199 |
| `inherits` | frame-ref | `EA_FullResizeImage_BlackTransparent`, `EA_FullResizeImage_TintableSolidBackground`, `EA_FullResizeImage_WhiteTransparent`, `AuraWindowBackground`, … | 187 |
| `handleinput` | boolean | `false`, `true` | 101 |
| `layer` | string | `background`, `popup`, `overlay`, `default`, … | 64 |
| `alpha` | number | `0.5`, `0.3`, `0.1`, `0.75`, … | 51 |
| `texture` | string | `shared_01`, `enemy_bar_rect`, `Frame_1`, `Frame_BG`, … | 11 |
| `skipinput` | boolean | `true` | 10 |
| `frameonly` | boolean | `true` | 5 |
| `showing` | boolean | `false` | 5 |
| `hanldeinput` | boolean | `false` | 2 |
| `drawchildrenfirst` | boolean | `true` | 1 |
| `font` | string | `font_clear_small_bold` | 1 |
| `sticky` | boolean | `false` | 1 |
| `textalign` | string | `center` | 1 |

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
