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
| Addons seen in | DAoCBuff, Enemy, Killer, PotionBar, Shinies, WhoHealedMe, bigger_MacroWindow |
| Files seen in | `/workspace/data/raw/DAoCBuff/Source/DAoCBuffMsgWindow.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:2134`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:2464`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:5`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:988`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogConfiguration.xml:0`, `/workspace/data/raw/Enemy/Code/GroupIcons/GroupIconsConfiguration.xml:0`, `/workspace/data/raw/Enemy/Code/Guard/GuardConfiguration.xml:0` |
| Namespaces detected | VerticalScrollbar |
| Source kinds | xml_frames, xml_handlers |
| Example locations | DAoCBuff: DAoCBuffFrameSettingsTab_Scrollbar, DAoCBuff: DAoCBuffGeneralSettingsTab_Scrollbar, DAoCBuff: DAoCBuffListManagerTab_Scrollbar, DAoCBuff: DAoCBuffMessageWindowScrollWindowScrollbar, DAoCBuff: DAoCBuff_Settings_FilterFrame_Scrollbar, Enemy: EnemyClickCastingDialogContentScrollbar |
| XML usage count | 25 |
| XML attribute usage count | 25 |
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

VerticalScrollbar is a container-style XML element. It commonly appears under LogDisplay and ScrollWindow. It is typically used to organize structural children such as Size and bind XML events like OnScrollPosChanged to Lua.

## Common Attributes

- inherits
- layer
- name

## Common Handlers

- [OnScrollPosChanged](../handlers/handler_OnScrollPosChanged.md)

## Common Handler Functions

- MacroIcons.ScrollPos


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnScrollPosChanged](../handlers/handler_OnScrollPosChanged.md) | data | MacroIcons.ScrollPos | `function(...)` | LOW |
| [OnScrollPosChanged](../handlers/handler_OnScrollPosChanged.md) | data | MacroIcons.ScrollPos | `scrollPos` | LOW |

## Common Inherits

- EA_ScrollBar_ChatVertical
- EA_ScrollBar_DefaultVerticalChain

## Common Parent Elements

- [Windows](element_Windows.md) — 25× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 25× (HIGH)
- [Size](element_Size.md) — 23× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 1× (LOW)

## Common Template Bases

- EA_ScrollBar_ChatVertical
- EA_ScrollBar_DefaultVerticalChain

## Typical XML Structure

```xml
<VerticalScrollbar inherits="EA_ScrollBar_DefaultVerticalC..." layer="popup" name="...">
  <Size/>
</VerticalScrollbar>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 100% | EA_ScrollBar_DefaultVerticalChain, EA_ScrollBar_ChatVertical |
| `layer` | **required** | 96% | popup |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 25 times as an unnamed child.

### [Size](element_Size.md)

Observed 23 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [VerticalScrollbar](element_VerticalScrollbar.md)
- [Anchors](element_Anchors.md) (structural, 25×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
      - (cycle)
- [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
  - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
- [Size](element_Size.md) (structural, 23×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)

## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnScrollPosChanged

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `scrollPos` | number | scroll_position |

## Binding Resolution

- Total handler declarations: 1
- Resolved to Lua functions: 0 (0%)

## Seen In

- DAoCBuff
- Enemy
- Killer
- PotionBar
- Shinies
- WhoHealedMe
- bigger_MacroWindow

## Examples

- DAoCBuff: DAoCBuffFrameSettingsTab_Scrollbar -> VerticalScrollbar DAoCBuffFrameSettingsTab_Scrollbar
- DAoCBuff: DAoCBuffGeneralSettingsTab_Scrollbar -> VerticalScrollbar DAoCBuffGeneralSettingsTab_Scrollbar
- DAoCBuff: DAoCBuffListManagerTab_Scrollbar -> VerticalScrollbar DAoCBuffListManagerTab_Scrollbar
- DAoCBuff: DAoCBuffMessageWindowScrollWindowScrollbar -> VerticalScrollbar DAoCBuffMessageWindowScrollWindowScrollbar
- DAoCBuff: DAoCBuff_Settings_FilterFrame_Scrollbar -> VerticalScrollbar DAoCBuff_Settings_FilterFrame_Scrollbar
- Enemy: EnemyClickCastingDialogContentScrollbar -> VerticalScrollbar EnemyClickCastingDialogContentScrollbar

## Related APIs

- [EA_ScrollBar_DefaultVerticalChain](../../globals/constants/constant_EA_ScrollBar_DefaultVerticalChain.md) (HIGH 100/100) - Constant
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EA_ScrollBar_ChatVertical](../../globals/constants/constant_EA_ScrollBar_ChatVertical.md) (HIGH 90/100) - Constant
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 55/100) - XML Element Type

## Used With

- none

## Triggered By

- [OnScrollPosChanged](../handlers/handler_OnScrollPosChanged.md) (HIGH 73/100) - XML Event

## Affects

- none
