# Icons.GetCareerIconIDFromCareerLine

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 133

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AutoMark, Effigy, Enemy, MapPin, Moth, Swift Assist, Targets |
| Files seen in | `/workspace/AutoMark/Source/AutoMark.lua:7`, `/workspace/Effigy/Elements/EffigyIcons.lua:138`, `/workspace/Effigy/Elements/EffigyIcons.lua:172`, `/workspace/Enemy/Code/Assist/Assist.lua:195`, `/workspace/Enemy/Code/CombatLog/CombatLogIDS.lua:211`, `/workspace/Enemy/Code/GroupIcons/GroupIcon.lua:115`, `/workspace/Enemy/Code/Marks/MarkTemplate.lua:85`, `/workspace/Enemy/Code/UnitFrames/Parts/CareerIcon.lua:13` |
| Namespaces detected | Icons |
| Source kinds | globals, lua_calls |
| Example locations | AutoMark: AutoMark.local.CreateMarker, AutoMark: CreateMarker, Effigy: Effigy.local.RenderCareerIcon, Effigy: Effigy.local.SetupCareerIcon, Effigy: RenderCareerIcon, Effigy: SetupCareerIcon |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 22 |
| Global usage count | 22 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
Icons.GetCareerIconIDFromCareerLine(arg1)
```

## Description

Observed as a global function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: Addon.career, GameData.Player.career.line, cache.career |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- AutoMark
- Effigy
- Enemy
- MapPin
- Moth
- Swift Assist
- Targets

## Examples

- AutoMark: AutoMark.local.CreateMarker -> Icons.GetCareerIconIDFromCareerLine(career_id)
- AutoMark: CreateMarker -> Icons.GetCareerIconIDFromCareerLine(career_id)
- Effigy: Effigy.local.RenderCareerIcon -> Icons.GetCareerIconIDFromCareerLine(value)
- Effigy: Effigy.local.SetupCareerIcon -> Icons.GetCareerIconIDFromCareerLine(GameData.Player.career.line)
- Effigy: RenderCareerIcon -> Icons.GetCareerIconIDFromCareerLine(value)
- Effigy: SetupCareerIcon -> Icons.GetCareerIconIDFromCareerLine(GameData.Player.career.line)

## Related APIs

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function

## Used With

- [DialogManager.MakeTextEntryDialog](global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow.InsertText](global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap](../tables/table_EA_Window_WorldMap.md) (HIGH 100/100) - Global Table
- [EA_Window_WorldMap.ShowZone](global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
- [GameData.Player.career.line](../../gamedata/fields/gamedata_GameData.Player.career.line.md) (HIGH 100/100) - GameData Field
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [GameData.Player.career.line](../../gamedata/fields/gamedata_GameData.Player.career.line.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
