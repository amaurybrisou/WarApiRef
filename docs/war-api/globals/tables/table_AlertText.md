# AlertText

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 130

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Aura, GCDsaver, MapPin, PlanB |
| Namespaces detected | AlertText |
| Source kinds | globals, lua_calls |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 4 |
| Documentation references | 1 |
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

Observed shared global table or namespace surfaced in 4 addons.

## Functions

- none

## Observed Members

- none

## Seen In

- Aura
- GCDsaver
- MapPin
- PlanB

## Examples

- none

## Related APIs

- [PartyUtils.GetWarbandLeader](../functions/global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [BroadcastEvent](../functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function
- [AlertTextWindow.AddAlert](../functions/global_AlertTextWindow.AddAlert.md) (HIGH 88/100) - Global Function
- [towstring](../functions/global_towstring.md) (HIGH 75/100) - Global Function

## Used With

- [GameData.ChatData.name](../../gamedata/fields/gamedata_GameData.ChatData.name.md) (HIGH 100/100) - GameData Field
- [GameData.ChatData.text](../../gamedata/fields/gamedata_GameData.ChatData.text.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.QUEST_ABANDONED](../../gamedata/fields/gamedata_GameData.Sound.QUEST_ABANDONED.md) (HIGH 100/100) - GameData Field
- [PartyUtils.GetWarbandLeader](../functions/global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](../functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../functions/global_wstring.match.md) (HIGH 100/100) - Global Function
- [AlertTextWindow.AddAlert](../functions/global_AlertTextWindow.AddAlert.md) (HIGH 88/100) - Global Function
- [towstring](../functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [SystemData.Events.SHOW_ALERT_TEXT](../../systemdata/fields/systemdata_SystemData.Events.SHOW_ALERT_TEXT.md) (HIGH 100/100) - SystemData Field

## Notes

- none
