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
| Addons seen in | DAoCBuff, EA_UiModWindow, Enemy, Killer, Miracle Grow Remix, PotionBar, RVMOD_Manager, Shinies |
| Files seen in | `/workspace_addons/DAoCBuff/Source/DAoCBuffMsgWindow.xml:53`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings.xml:128`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings.xml:139`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings.xml:150`, `/workspace_addons/Enemy/Code/CombatLog/CombatLogConfiguration.xml:8`, `/workspace_addons/Enemy/Code/GroupIcons/GroupIconsConfiguration.xml:8`, `/workspace_addons/Enemy/Code/Guard/GuardConfiguration.xml:8`, `/workspace_addons/Enemy/Code/KillSpam/KillSpamConfiguration.xml:8` |
| Namespaces detected | ScrollWindow |
| Source kinds | xml_frames |
| Example locations | DAoCBuff: DAoCBuffMessageWindowScrollWindow, DAoCBuff: DAoCBuff_Settings_FrameSettings, DAoCBuff: DAoCBuff_Settings_GeneralSettings, DAoCBuff: DAoCBuff_Settings_ListManagerSettings, EA_UiModWindow: EA_ScrollWindow_ModInfoTemplate, EA_UiModWindow: UiModWindowModDetails |
| XML usage count | 31 |
| XML attribute usage count | 31 |
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

Observed XML element type instantiated by 11 addons.

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
- EA_ScrollWindow_ModInfoTemplate
- RVMOD_ManagerModInfoTemplate

## Common Parent Elements

- [Window](element_Window.md)

## Common Named Child Elements

- [VerticalScrollbar](element_VerticalScrollbar.md)
- [Window](element_Window.md)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `childscrollwindow` | optional | 46% | MiracleGrow2Preset, MiracleGrow2Config, $parentScrollChild, MiracleGrow2Layout, ... |
| `lineheight` | optional | 46% | 20, 18, 19, 90, ... |
| `scrollbar` | optional | 46% | $parentScroll, $parentScrollbar, RVMOD_ManagerWindowContentListBoxScrollbar, $parentBar, ... |
| `autoHideScrollBar` | optional | 26% | true, false |
| `inherits` | optional | 8% | RVMOD_ManagerModInfoTemplate, DAoCBuffGeneralSettingsTab, DAoCBuffFrameSettingsTab, DAoCBuffListManagerTab, ... |
| `autohidescrollbar` | optional | 3% | true |
| `layer` | optional | 1% | secondary |
## Lua Functions Manipulating This Type

- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnExceptMeChanged
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_OnExceptMeChanged
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- Enemy.Enemy.local._OnArchetypeChanged
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnExceptMeChanged
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnCircleIconChanged
- wbLeadHelper.wbLeadHelperConfigTab.OnLoad
- Killer.Killer.Initialize
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- Enemy._OnKeyModifierChanged
- wbLeadHelper.WbLeadHelperMessage.OnOk
- MapMonster.FilterButtonState
- wbLeadHelper.wbLeadHelperConfigTab.OnReset
- DAoCBuff.DAoCBuffSettings.PopulateSettings
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_UpdateExample
- wbLeadHelper.WbLeadHelperMessage.MessageDialogOpen
- Enemy.Enemy.local._OnKeyModifierChanged
- MapMonster.MapMonster.local.FilterButtonState
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListSelChanged
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnIconSelChanged
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- DAoCBuff.DAoCBuff.ShowMessageWindow
- wbLeadHelper.wbLeadHelperConfigTab.OnLfgIconsCheckBoxUp
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Enemy._OnArchetypeChanged
- wbLeadHelper.wbLeadHelperConfigTab.Initialize
- Enemy.EnemyGroupIcon:ApplySettings
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged

## Seen In

- DAoCBuff
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

- DAoCBuff: DAoCBuffMessageWindowScrollWindow -> ScrollWindow DAoCBuffMessageWindowScrollWindow
- DAoCBuff: DAoCBuff_Settings_FrameSettings -> ScrollWindow DAoCBuff_Settings_FrameSettings
- DAoCBuff: DAoCBuff_Settings_GeneralSettings -> ScrollWindow DAoCBuff_Settings_GeneralSettings
- DAoCBuff: DAoCBuff_Settings_ListManagerSettings -> ScrollWindow DAoCBuff_Settings_ListManagerSettings
- EA_UiModWindow: EA_ScrollWindow_ModInfoTemplate -> ScrollWindow EA_ScrollWindow_ModInfoTemplate
- EA_UiModWindow: UiModWindowModDetails -> ScrollWindow UiModWindowModDetails

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
