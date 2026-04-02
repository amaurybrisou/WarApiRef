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

- OnMouseOver
- OnMouseOverEnd
- OnLButtonUp

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

## Common Structural Child Elements

- [TintColor](element_TintColor.md)
- [Middle](element_Middle.md)
- [TexSlices](element_TexSlices.md)

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
