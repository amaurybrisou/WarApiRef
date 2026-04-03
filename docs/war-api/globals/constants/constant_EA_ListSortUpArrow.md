# EA_ListSortUpArrow

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
| Addons seen in | BlackBook, BuffHead, CM_ClosetGoblin, CastSequence, EA_UiModWindow, Shinies, SocialWindow 2.0 |
| Files seen in | BlackBook.xml, ClosetGoblin.xml, Setup/Builder.xml, Setup/SetupEffectCache.xml, Source/ShiniesUITemplates.xml, Source/SocialWindowTabFriends.xml, Source/SocialWindowTabIgnore.xml, Source/SocialWindowTabSearch.xml |
| Namespaces detected | EA_ListSortUpArrow |
| Source kinds | xml_attributes |
| Example locations | BlackBookWindowSortBarUpArrow, BuffHeadSetupEffectCacheWindowSortBarUpArrow, CastSequenceBuilderWindowSortBarUpArrow, ClosetGoblinCharacterWindowContentsSortTacticsUpArrow, ClosetGoblinCharacterWindowContentsSortUpArrow, ClosetGoblinZoneWindowContentsSortSetUpArrow |
| XML usage count | 13 |
| XML attribute usage count | 13 |
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

Engine-supplied XML constant or template class referenced by 7 addons.

## Seen In

- BlackBook
- BuffHead
- CM_ClosetGoblin
- CastSequence
- EA_UiModWindow
- Shinies
- SocialWindow 2.0

## Used By

- BlackBookWindowSortBarUpArrow
- BuffHeadSetupEffectCacheWindowSortBarUpArrow
- CastSequenceBuilderWindowSortBarUpArrow
- ClosetGoblinCharacterWindowContentsSortTacticsUpArrow
- ClosetGoblinCharacterWindowContentsSortUpArrow
- ClosetGoblinZoneWindowContentsSortSetUpArrow
- ClosetGoblinZoneWindowContentsSortUpArrow
- Shinies_Button_DefaultListSortUpArrow
- SocialWindowTabFriendsSocketSortButtonBarUpArrow
- SocialWindowTabIgnoreSocketSortButtonBarUpArrow
- SocialWindowTabSearchSocketSortButtonBarUpArrow
- UiModVersionMismatchWindowSortUpArrow
- UiModWindowSortUpArrow

## Related APIs

- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type

## Notes

- none
