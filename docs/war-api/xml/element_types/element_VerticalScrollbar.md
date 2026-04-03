# VerticalScrollbar

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CDown, Crusher, DAoCBuff, DPSMeter, DeepSleep, DuffTimer, EA_ScenarioGroupWindow, EA_UiDebugTools |
| Files seen in | Code/CombatLog/CombatLogConfiguration.xml, Code/GroupIcons/GroupIconsConfiguration.xml, Code/Guard/GuardConfiguration.xml, Code/KillSpam/KillSpamConfiguration.xml, Code/TalismanAlerter/TalismanAlerterConfiguration.xml, Code/Timer/TimerConfiguration.xml, Code/UnitFrames/ClickCastingDialog.xml, Code/UnitFrames/EffectsIndicatorDialog.xml |
| Namespaces detected | VerticalScrollbar |
| Source kinds | xml_frames |
| Example locations | CDown: CDownColorSettingsTab_Scrollbar, CDown: CDownGeneralSettingsTab_Scrollbar, CDown: CDownNLayoutSettingsTab_Scrollbar, CDown: CDownSLayoutSettingsTab_Scrollbar, Crusher: CrusherConfigScrollBar, DAoCBuff: DAoCBuffFrameSettingsTab_Scrollbar |
| XML usage count | 62 |
| XML attribute usage count | 62 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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

VerticalScrollbar is an interactive XML control. It commonly appears under EditBox and LogDisplay. It is typically used to organize structural children such as ActiveZoneOffset, Anchors, DownOffset and bind XML events like OnLButtonUp, OnScrollPosChanged to Lua.

## Common Attributes

- alpha
- down
- gutter
- inherits
- layer
- name
- thumb
- up

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnScrollPosChanged](../handlers/handler_OnScrollPosChanged.md)

## Common Handler Functions

- AuctionWindow.OnVertScrollLButtonUp
- DPSMeter.OnScrollbarLButtonUp
- ManualWindow.OnVertScrollLButtonUp
- ScenarioGroupWindow.OnVertScrollLButtonUp
- emotes.OnVertScrollLButtonUp
- MacroIcons.ScrollPos


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | AuctionWindow.OnVertScrollLButtonUp, DPSMeter.OnScrollbarLButtonUp, ManualWindow.OnVertScrollLButtonUp, ScenarioGroupWindow.OnVertScrollLButtonUp, emotes.OnVertScrollLButtonUp | `flags, x, y` | MEDIUM |
| [OnScrollPosChanged](../handlers/handler_OnScrollPosChanged.md) | data | MacroIcons.ScrollPos | `scrollPos` | MEDIUM |

### Per-Event Lua API Calls

**OnScrollPosChanged** handlers call: `VerticalScrollbarGetScrollPosition`

## Common Inherits

- EA_ScrollBar_ChatVertical
- EA_ScrollBar_DefaultVerticalChain
- StandardVertScroll

## Common Parent Elements

- [Windows](element_Windows.md) — 62× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 56× (HIGH)
- [Size](element_Size.md) — 47× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 7× (MEDIUM)
- [ActiveZoneOffset](element_ActiveZoneOffset.md) — 2× (LOW)
- [DownOffset](element_DownOffset.md) — 2× (LOW)
- [ThumbOffset](element_ThumbOffset.md) — 2× (LOW)
- [UpOffset](element_UpOffset.md) — 2× (LOW)

## Common Template Bases

- EA_ScrollBar_ChatVertical
- EA_ScrollBar_DefaultVerticalChain
- StandardVertScroll


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 96% | EA_ScrollBar_DefaultVerticalChain, EA_ScrollBar_ChatVertical, StandardVertScroll |
| `layer` | **required** | 83% | popup |
| `down` | optional | 3% | EA_ScrollBar_ChatDownArrowButton, SpamBayesTemplateScrollBarDownArrowButton |
| `thumb` | optional | 3% | EA_ScrollBar_ChatThumb, SpamBayesTemplateScrollBarThumb |
| `up` | optional | 3% | EA_ScrollBar_ChatUpArrowButton, SpamBayesTemplateScrollBarUpArrowButton |
| `alpha` | optional | 1% | 0.97 |
| `gutter` | optional | 1% | SpamBayesTemplateScrollBarGutter |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 56 times as an unnamed child.

### [Size](element_Size.md)

Observed 47 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 7 times as an unnamed child.

### [ActiveZoneOffset](element_ActiveZoneOffset.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 100 |
| `y` | **required** | 0 |
### [DownOffset](element_DownOffset.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0 |
| `y` | **required** | 0 |
### [ThumbOffset](element_ThumbOffset.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 3, 0 |
| `y` | **required** | 0 |
### [UpOffset](element_UpOffset.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0 |
| `y` | **required** | 0 |
## Recursive Hierarchy

- Root: [VerticalScrollbar](element_VerticalScrollbar.md)
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

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `VerticalScrollbarGetScrollPosition` | ui | 2 | OnScrollPosChanged |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnScrollPosChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `scrollPos` | number | scroll_position |
## Lua Functions Manipulating This Type

- MacroIcons.OnInitialize
- EA_Window_Macro.Initialize


## Binding Resolution

- Total handler declarations: 7
- Resolved to Lua functions: 7 (100%)

## Seen In

- CDown
- Crusher
- DAoCBuff
- DPSMeter
- DeepSleep
- DuffTimer
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- EveryBodyGuard
- FozAuction
- Hopper
- Killer
- MacroIcons
- Miracle Grow Remix
- Motion
- PotionBar
- Pure
- RVMOD_Manager
- RVMOD_Targets
- Shinies
- SocialWindow 2.0
- SpamBayes
- Tome Titan
- Vectors
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- bigger_MacroWindow
- emotes
- wbLeadHelper

## Examples

- CDown: CDownColorSettingsTab_Scrollbar -> VerticalScrollbar CDownColorSettingsTab_Scrollbar
- CDown: CDownGeneralSettingsTab_Scrollbar -> VerticalScrollbar CDownGeneralSettingsTab_Scrollbar
- CDown: CDownNLayoutSettingsTab_Scrollbar -> VerticalScrollbar CDownNLayoutSettingsTab_Scrollbar
- CDown: CDownSLayoutSettingsTab_Scrollbar -> VerticalScrollbar CDownSLayoutSettingsTab_Scrollbar
- Crusher: CrusherConfigScrollBar -> VerticalScrollbar CrusherConfigScrollBar
- DAoCBuff: DAoCBuffFrameSettingsTab_Scrollbar -> VerticalScrollbar DAoCBuffFrameSettingsTab_Scrollbar

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [EA_ScrollBar_ChatVertical](../../globals/constants/constant_EA_ScrollBar_ChatVertical.md) (HIGH 100/100) - Constant
- [EA_ScrollBar_DefaultVerticalChain](../../globals/constants/constant_EA_ScrollBar_DefaultVerticalChain.md) (HIGH 100/100) - Constant
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [ActiveZoneOffset](element_ActiveZoneOffset.md) (HIGH 98/100) - XML Element Type
- [DownOffset](element_DownOffset.md) (HIGH 98/100) - XML Element Type
- [ThumbOffset](element_ThumbOffset.md) (HIGH 98/100) - XML Element Type
- [UpOffset](element_UpOffset.md) (HIGH 98/100) - XML Element Type

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnScrollPosChanged](../handlers/handler_OnScrollPosChanged.md) (HIGH 76/100) - XML Event
