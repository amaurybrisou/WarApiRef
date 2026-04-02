# ActionButtonGroup

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 161

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BankWindowFix, EA_UiModWindow, RVMOD_Manager |
| Files seen in | `/workspace_addons/BankWindowFix/Source/BankWindowFix.xml:3`, `/workspace_addons/RVMOD_Manager/RVMOD_ManagerTemplates.xml:322`, `/workspace_addons/ea_uimodwindow/Source/UiModInfoTemplate.xml:213` |
| Namespaces detected | ActionButtonGroup |
| Source kinds | xml_frames, xml_handlers |
| Example locations | BankWindowFix: BankWindowSlotsFixed, EA_UiModWindow: EA_ScrollWindow_ModInfoTemplateScrollChildCareers, RVMOD_Manager: RVMOD_ManagerModInfoTemplateScrollChildCareers |
| XML usage count | 3 |
| XML attribute usage count | 3 |
| Lua usage count | 4 |
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

Observed XML element type instantiated by 3 addons.

## Common Attributes

- inherits
- name
- hideButtonWhenIconBlank
- draganddrop
- gameactionbutton

## Common Handlers

- OnActionButtonLButtonDown
- OnActionButtonLButtonUp
- OnActionButtonMouseOver
- OnActionButtonRButtonDown

## Common Inherits

- EA_ActionButtonGroup_CareerIconsWithTooltip
- EA_ActionButtonGroup_DefaultSmall

## Seen In

- BankWindowFix
- EA_UiModWindow
- RVMOD_Manager

## Examples

- BankWindowFix: BankWindowSlotsFixed -> ActionButtonGroup BankWindowSlotsFixed
- EA_UiModWindow: EA_ScrollWindow_ModInfoTemplateScrollChildCareers -> ActionButtonGroup EA_ScrollWindow_ModInfoTemplateScrollChildCareers
- RVMOD_Manager: RVMOD_ManagerModInfoTemplateScrollChildCareers -> ActionButtonGroup RVMOD_ManagerModInfoTemplateScrollChildCareers

## Related APIs

- none

## Used With

- [OnActionButtonLButtonDown](../handlers/handler_OnActionButtonLButtonDown.md) (HIGH 73/100) - XML Handler
- [OnActionButtonLButtonUp](../handlers/handler_OnActionButtonLButtonUp.md) (HIGH 73/100) - XML Handler
- [OnActionButtonMouseOver](../handlers/handler_OnActionButtonMouseOver.md) (HIGH 73/100) - XML Handler
- [OnActionButtonRButtonDown](../handlers/handler_OnActionButtonRButtonDown.md) (HIGH 73/100) - XML Handler
- [OnActionButtonLButtonDown](../../events/window_events/window_event_OnActionButtonLButtonDown.md) (MEDIUM 53/100) - Window Event
- [OnActionButtonLButtonUp](../../events/window_events/window_event_OnActionButtonLButtonUp.md) (MEDIUM 53/100) - Window Event
- [OnActionButtonMouseOver](../../events/window_events/window_event_OnActionButtonMouseOver.md) (MEDIUM 53/100) - Window Event
- [OnActionButtonRButtonDown](../../events/window_events/window_event_OnActionButtonRButtonDown.md) (MEDIUM 53/100) - Window Event

## Triggered By

- none

## Affects

- none
