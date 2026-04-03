# ScrollWindow

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | DAoCBuff, Enemy, Killer, PartyCast, PotionBar, Shinies, WhoHealedMe, bigger_MacroWindow |
| Files seen in | `/workspace/data/raw/DAoCBuff/Source/DAoCBuffMsgWindow.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettings.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:2132`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:2462`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:3`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:978`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogConfiguration.xml:0`, `/workspace/data/raw/Enemy/Code/GroupIcons/GroupIconsConfiguration.xml:0` |
| Namespaces detected | ScrollWindow |
| Source kinds | xml_frames |
| Example locations | DAoCBuff: DAoCBuffFrameSettingsTab, DAoCBuff: DAoCBuffGeneralSettingsTab, DAoCBuff: DAoCBuffListManagerTab, DAoCBuff: DAoCBuffMessageWindowScrollWindow, DAoCBuff: DAoCBuff_Settings_FilterFrame, DAoCBuff: DAoCBuff_Settings_FrameSettings |
| XML usage count | 27 |
| XML attribute usage count | 27 |
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

Observed XML element type instantiated by 8 addons.

## Common Attributes

- name
- childscrollwindow
- lineheight
- scrollbar
- autoHideScrollBar
- inherits
- autohidescrollbar
- layer

## Common Inherits

- DAoCBuffFrameSettingsTab
- DAoCBuffGeneralSettingsTab
- DAoCBuffListManagerTab

## Common Parent Elements

- [Windows](element_Windows.md) — 26× (HIGH)

## Common Structural Child Elements

- [Windows](element_Windows.md) — 24× (HIGH)
- [Size](element_Size.md) — 12× (HIGH)

## Common Template Bases

- DAoCBuffFrameSettingsTab
- DAoCBuffGeneralSettingsTab
- DAoCBuffListManagerTab


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<ScrollWindow autoHideScrollBar="false" childscrollwindow="$parentScrollChild" lineheight="18" name="..." scrollbar="$parentScrollbar">
  <Size/>
  <Windows/>
</ScrollWindow>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `childscrollwindow` | **required** | 88% | $parent_ScrollChild, $parentScrollChild, $parent_scroll, $parentRightPane, ... |
| `lineheight` | **required** | 88% | 18, 38, 19, 55, ... |
| `scrollbar` | **required** | 85% | $parent_Scrollbar, $parentScrollbar, $parentBar, IconsScrollbar |
| `autoHideScrollBar` | optional | 55% | true, false |
| `inherits` | optional | 11% | DAoCBuffFrameSettingsTab, DAoCBuffListManagerTab, DAoCBuffGeneralSettingsTab |
| `autohidescrollbar` | optional | 7% | true |
| `layer` | optional | 3% | secondary |
## Structural Sub-Elements

### [Windows](element_Windows.md)

Observed 24 times as an unnamed child.

### [Size](element_Size.md)

Observed 12 times as an unnamed child.

## Lua Functions Manipulating This Type

- DAoCBuff.ShowMessageWindow
- DAoCBuffSettings.PopulateSettings
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListSelChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.local._OnArchetypeChanged
- Enemy._OnKeyModifierChanged
- Killer.Initialize
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Enemy._OnArchetypeChanged
- Enemy.local._OnKeyModifierChanged
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- Enemy.UnitFramesUI_UnitFramePartDialog_UpdateExample
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- Enemy.EnemyGroupIcon:ApplySettings

## Seen In

- DAoCBuff
- Enemy
- Killer
- PartyCast
- PotionBar
- Shinies
- WhoHealedMe
- bigger_MacroWindow

## Examples

- DAoCBuff: DAoCBuffFrameSettingsTab -> ScrollWindow DAoCBuffFrameSettingsTab
- DAoCBuff: DAoCBuffGeneralSettingsTab -> ScrollWindow DAoCBuffGeneralSettingsTab
- DAoCBuff: DAoCBuffListManagerTab -> ScrollWindow DAoCBuffListManagerTab
- DAoCBuff: DAoCBuffMessageWindowScrollWindow -> ScrollWindow DAoCBuffMessageWindowScrollWindow
- DAoCBuff: DAoCBuff_Settings_FilterFrame -> ScrollWindow DAoCBuff_Settings_FilterFrame
- DAoCBuff: DAoCBuff_Settings_FrameSettings -> ScrollWindow DAoCBuff_Settings_FrameSettings

## Related APIs

- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
