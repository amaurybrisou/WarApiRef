# GameData.PlayerActions.SET_TARGET

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 17 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AggroMeter, BuddyBind, CleanUnitFrames, DuelInvite, EA_ScenarioGroupWindow, Effigy, Enemy, HealGrid |
| Files seen in | AggroMeter.lua, BuddyBind.lua, CleanGroupMemberUnitFrame.lua, Code/GroupIcons/GroupIcon.lua, Code/Guard/Guard.lua, Code/Marks/MarkTemplate.lua, Code/UnitFrames/UnitFrame.lua, Core.lua |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | AddInteractionMenuItems, Apply, Attach, ClearAction, GrabName, Guard_SetGuardPlayerName |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 22 |
| Global usage count | 22 |
| Local definition count | 0 |
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

GameData.GameData.PlayerActions.SET_TARGET field accessed by 17 addons; commonly found in AddInteractionMenuItems and Apply, Attach, ClearAction, GrabName, Guard_SetGuardPlayerName, LButtonUp, OnClickName, PotentialClickAssist, ProcessTargetUnit, SelectChar, SetLeaderName, SetMacroTarget, SetName, UnitFramesUI_UnitFrame_OnLButtonDown, UpdateFullMember, UpdateSingleMemberWindow, UpdateTarget, addInteractive, init, lua_call, update contexts.

## Seen In

- AggroMeter
- BuddyBind
- CleanUnitFrames
- DuelInvite
- EA_ScenarioGroupWindow
- Effigy
- Enemy
- HealGrid
- HealHoverAssist
- MarkBuff
- Pure
- Refer
- ResHelp
- Squared
- Targets
- WarTriage
- followTheLeader

## Related APIs

- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [Icons.GetCareerIconIDFromCareerNamesID](../../globals/functions/global_Icons.GetCareerIconIDFromCareerNamesID.md) (HIGH 100/100) - Global Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionTrigger](../../window_api/functions/window_WindowSetGameActionTrigger.md) (HIGH 100/100) - Window Function
- [EA_Window_ContextMenu.GameActionData](../../globals/functions/global_EA_Window_ContextMenu.GameActionData.md) (HIGH 80/100) - Global Function
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function

## Used With

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function

## Notes

- Observed in contexts: AddInteractionMenuItems, Apply, Attach, ClearAction, GrabName, Guard_SetGuardPlayerName
