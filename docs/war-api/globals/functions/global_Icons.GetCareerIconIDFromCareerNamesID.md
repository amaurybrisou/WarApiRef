# Icons.GetCareerIconIDFromCareerNamesID

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | DetauntHelper, EA_ScenarioGroupWindow, SocialWindow 2.0, Trakario |
| Files seen in | Source/Config/BarSettings.lua, Source/ScenarioGroupWindow.lua, Source/TargetInfo.lua, source/socialwindow.lua, trakario.lua |
| Namespaces detected | Icons |
| Source kinds | lua_calls |
| Example locations | DetauntHelper: GetCareerIconDataFromCareerID, DetauntHelper: InitUI, DetauntHelper: _GetCareerIcon, EA_ScenarioGroupWindow: UpdateSingleMemberWindow, SocialWindow 2.0: OnMouseOverListMember, Trakario: SetCareerIcon |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
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
Icons.GetCareerIconIDFromCareerNamesID(arg1)
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GameData.Player.career.id, SocialWindow.MemberListTable[windowId].careerID, careerId |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- DetauntHelper
- EA_ScenarioGroupWindow
- SocialWindow 2.0
- Trakario

## Examples

- DetauntHelper: GetCareerIconDataFromCareerID -> Icons.GetCareerIconIDFromCareerNamesID(cid)
- DetauntHelper: InitUI -> Icons.GetCareerIconIDFromCareerNamesID(GameData.Player.career.id)
- DetauntHelper: _GetCareerIcon -> Icons.GetCareerIconIDFromCareerNamesID(self.pinfo.careerId)
- EA_ScenarioGroupWindow: UpdateSingleMemberWindow -> Icons.GetCareerIconIDFromCareerNamesID(player.careerId)
- SocialWindow 2.0: OnMouseOverListMember -> Icons.GetCareerIconIDFromCareerNamesID(SocialWindow.MemberListTable[windowId].careerID)
- Trakario: SetCareerIcon -> Icons.GetCareerIconIDFromCareerNamesID(careerId)

## Used With

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function

## Affects

- [GameData.Player.career.id](../../gamedata/fields/gamedata_GameData.Player.career.id.md) (HIGH 100/100) - GameData Field
- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
