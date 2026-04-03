# Icons.GetCareerIconIDFromCareerLine

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 30 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 108

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AutoMark, BuddyBind, Calling, CleanUnitFrames, DammazKron, DetauntHelper, Ding, EA_OpenPartyWindow |
| Files seen in | BuddyBind.lua, CallingList.lua, CallingNotification.lua, CleanGroupMemberUnitFrame.lua, Code/Assist/Assist.lua, Code/CombatLog/CombatLogIDS.lua, Code/GroupIcons/GroupIcon.lua, Code/Marks/MarkTemplate.lua |
| Namespaces detected | Icons |
| Source kinds | lua_calls |
| Example locations | AutoMark: CreateMarker, BuddyBind: GrabName, Calling: SetListLine, Calling: ShowCallerIcon, Calling: ShowTargetIcon, CleanUnitFrames: SetCareerIcon |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 89 |
| Global usage count | 89 |
| Local definition count | 0 |
| Documentation references | 0 |
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

Observed as a global function across 30 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: Addon.career, GameData.CareerLine.ARCHMAGE, GameData.CareerLine.ASSASSIN |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AutoMark
- BuddyBind
- Calling
- CleanUnitFrames
- DammazKron
- DetauntHelper
- Ding
- EA_OpenPartyWindow
- Effigy
- Enemy
- EveryBodyGuard
- Group Icons
- Group Icons SG
- GuildWarden
- HealGrid
- MapPin
- MarkBuff
- Moth
- PartyCast
- Pure
- QuickNameActions+
- ResHelp
- SessionRPs
- Squared
- Swift Assist
- TargetInfoRing
- Targets
- ThinkOutLoud
- xHUD
- yAssistHelper

## Examples

- AutoMark: CreateMarker -> Icons.GetCareerIconIDFromCareerLine(career_id)
- BuddyBind: GrabName -> Icons.GetCareerIconIDFromCareerLine(TargetInfo:UnitCareer("selffriendlytarget"))
- Calling: SetListLine -> Icons.GetCareerIconIDFromCareerLine(call.CallerCareerID)
- Calling: SetListLine -> Icons.GetCareerIconIDFromCareerLine(call.TargetCareerID)
- Calling: ShowCallerIcon -> Icons.GetCareerIconIDFromCareerLine(careerId)
- Calling: ShowTargetIcon -> Icons.GetCareerIconIDFromCareerLine(careerId)

## Related APIs

- [Color](../../xml/element_types/element_Color.md) (HIGH 100/100) - XML Element Type
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [DialogManager.MakeTextEntryDialog](global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](../../window_api/functions/window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow.InsertText](global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [LabelGetTextDimensions](../../window_api/functions/window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [EA_Window_WorldMap.ShowZone](global_EA_Window_WorldMap.ShowZone.md) (HIGH 90/100) - Global Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function

## Affects

- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
