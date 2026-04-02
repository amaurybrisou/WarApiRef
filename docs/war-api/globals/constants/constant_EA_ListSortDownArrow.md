# EA_ListSortDownArrow

- Category: Constant
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
| Addons seen in | BuffHead, CM_ClosetGoblin, EA_UiModWindow, Shinies |
| Files seen in | `/workspace_addons/BuffHead/Setup/SetupEffectCache.xml:158`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:1340`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:1367`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:414`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:441`, `/workspace_addons/Shinies/Source/ShiniesUITemplates.xml:177`, `/workspace_addons/ea_uimodwindow/Source/UiModWindow.xml:203`, `/workspace_addons/ea_uimodwindow/Source/VersionMismatchWindow.xml:140` |
| Namespaces detected | EA_ListSortDownArrow |
| Source kinds | xml_attributes |
| Example locations | BuffHeadSetupEffectCacheWindowSortBarDownArrow, ClosetGoblinCharacterWindowContentsSortDownArrow, ClosetGoblinCharacterWindowContentsSortTacticsDownArrow, ClosetGoblinZoneWindowContentsSortDownArrow, ClosetGoblinZoneWindowContentsSortSetDownArrow, Shinies_Button_DefaultListSortDownArrow |
| XML usage count | 8 |
| XML attribute usage count | 8 |
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

Observed engine XML template or inherited constant referenced by 4 addons.

## Seen In

- BuffHead
- CM_ClosetGoblin
- EA_UiModWindow
- Shinies

## Used By

- BuffHeadSetupEffectCacheWindowSortBarDownArrow
- ClosetGoblinCharacterWindowContentsSortDownArrow
- ClosetGoblinCharacterWindowContentsSortTacticsDownArrow
- ClosetGoblinZoneWindowContentsSortDownArrow
- ClosetGoblinZoneWindowContentsSortSetDownArrow
- Shinies_Button_DefaultListSortDownArrow
- UiModVersionMismatchWindowSortDownArrow
- UiModWindowSortDownArrow

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
