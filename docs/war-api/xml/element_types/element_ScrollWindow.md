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

| Attribute | Role | Observed Values | Count |
|-----------|------|-----------------|-------|
| `name` | string | `DAoCBuffMessageWindowScrollWindow`, `DAoCBuff_Settings_FrameSettings`, `DAoCBuff_Settings_GeneralSettings`, `DAoCBuff_Settings_ListManagerSettings`, … | 31 |
| `childscrollwindow` | string | `$parentScrollChild`, `MiracleGrow2Config`, `MiracleGrow2Layout`, `MiracleGrow2Preset`, … | 26 |
| `lineheight` | number | `18`, `20`, `19`, `90`, … | 26 |
| `scrollbar` | frame-ref | `$parentScrollbar`, `$parentScroll`, `$parentBar`, `RVMOD_ManagerWindowContentListBoxScrollbar`, … | 26 |
| `autoHideScrollBar` | boolean | `false`, `true` | 15 |
| `inherits` | frame-ref | `DAoCBuffFrameSettingsTab`, `DAoCBuffGeneralSettingsTab`, `DAoCBuffListManagerTab`, `EA_ScrollWindow_ModInfoTemplate`, … | 5 |
| `autohidescrollbar` | boolean | `true` | 2 |
| `layer` | string | `secondary` | 1 |

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
