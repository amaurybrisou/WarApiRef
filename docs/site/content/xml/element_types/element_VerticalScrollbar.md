# VerticalScrollbar

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 173

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | DAoCBuff, EA_UiDebugTools, EA_UiModWindow, Enemy, Killer, Miracle Grow Remix, PotionBar, RVMOD_Manager |
| Files seen in | `/workspace_addons/DAoCBuff/Source/DAoCBuffMsgWindow.xml:63`, `/workspace_addons/Enemy/Code/CombatLog/CombatLogConfiguration.xml:13`, `/workspace_addons/Enemy/Code/GroupIcons/GroupIconsConfiguration.xml:13`, `/workspace_addons/Enemy/Code/Guard/GuardConfiguration.xml:13`, `/workspace_addons/Enemy/Code/KillSpam/KillSpamConfiguration.xml:13`, `/workspace_addons/Enemy/Code/TalismanAlerter/TalismanAlerterConfiguration.xml:13`, `/workspace_addons/Enemy/Code/Timer/TimerConfiguration.xml:13`, `/workspace_addons/Enemy/Code/UnitFrames/ClickCastingDialog.xml:46` |
| Namespaces detected | VerticalScrollbar |
| Source kinds | xml_frames, xml_handlers |
| Example locations | DAoCBuff: DAoCBuffMessageWindowScrollWindowScrollbar, EA_UiDebugTools: CopyScrollBar, EA_UiDebugTools: DebugWindowTextScrollbar, EA_UiDebugTools: DevPadWindowDevPadCodeDevPadCodeScrollBar, EA_UiDebugTools: EA_ScrollBar_ChatVertical, EA_UiDebugTools: ObjectInspectorObjectScrollbar |
| XML usage count | 33 |
| XML attribute usage count | 33 |
| Lua usage count | 1 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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

Observed XML element type instantiated by 12 addons.

## Common Attributes

- name
- inherits
- layer
- alpha
- down
- thumb
- up

## Common Handlers

- [OnScrollPosChanged](../handlers/handler_OnScrollPosChanged.md)

## Common Handler Functions

- MacroIcons.ScrollPos


## XML Event Bindings

| Event | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|---------------------|-------------------|-----------------|
| [OnScrollPosChanged](../handlers/handler_OnScrollPosChanged.md) | MacroIcons.ScrollPos | `function(...)` | LOW |

## Common Inherits

- EA_ScrollBar_DefaultVerticalChain
- EA_ScrollBar_ChatVertical

## Common Parent Elements

- [ScrollWindow](element_ScrollWindow.md)
- [EditBox](element_EditBox.md)
- [LogDisplay](element_LogDisplay.md)
- [Window](element_Window.md)

## Common Structural Child Elements

- [ActiveZoneOffset](element_ActiveZoneOffset.md)
- [DownOffset](element_DownOffset.md)
- [ThumbOffset](element_ThumbOffset.md)
- [UpOffset](element_UpOffset.md)

## Typical XML Structure

```xml
<VerticalScrollbar name="..." up="EA_ScrollBar_ChatUpArrowButton" down="EA_ScrollBar_ChatDownArrowButton" thumb="EA_ScrollBar_ChatThumb">
  <ThumbOffset x="3" y="0"/>
  <UpOffset x="0" y="0"/>
  <DownOffset x="0" y="0"/>
  <ActiveZoneOffset x="100" y="0"/>
</VerticalScrollbar>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 49% | EA_ScrollBar_DefaultVerticalChain, EA_ScrollBar_ChatVertical |
| `layer` | optional | 46% | popup |
| `alpha` | optional | 1% | 0.97 |
| `down` | optional | 1% | EA_ScrollBar_ChatDownArrowButton |
| `thumb` | optional | 1% | EA_ScrollBar_ChatThumb |
| `up` | optional | 1% | EA_ScrollBar_ChatUpArrowButton |
## Structural Sub-Elements

### [ActiveZoneOffset](element_ActiveZoneOffset.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
### [DownOffset](element_DownOffset.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
### [ThumbOffset](element_ThumbOffset.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
### [UpOffset](element_UpOffset.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnScrollPosChanged

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `scrollPos` | number | scroll_position |
## Seen In

- DAoCBuff
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- Killer
- Miracle Grow Remix
- PotionBar
- RVMOD_Manager
- Shinies
- WhoHealedMe
- bigger_MacroWindow
- wbLeadHelper

## Examples

- DAoCBuff: DAoCBuffMessageWindowScrollWindowScrollbar -> VerticalScrollbar DAoCBuffMessageWindowScrollWindowScrollbar
- EA_UiDebugTools: CopyScrollBar -> VerticalScrollbar CopyScrollBar
- EA_UiDebugTools: DebugWindowTextScrollbar -> VerticalScrollbar DebugWindowTextScrollbar
- EA_UiDebugTools: DevPadWindowDevPadCodeDevPadCodeScrollBar -> VerticalScrollbar DevPadWindowDevPadCodeDevPadCodeScrollBar
- EA_UiDebugTools: EA_ScrollBar_ChatVertical -> VerticalScrollbar EA_ScrollBar_ChatVertical
- EA_UiDebugTools: ObjectInspectorObjectScrollbar -> VerticalScrollbar ObjectInspectorObjectScrollbar

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
