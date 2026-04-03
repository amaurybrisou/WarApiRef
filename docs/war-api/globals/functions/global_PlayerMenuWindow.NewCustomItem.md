# PlayerMenuWindow.NewCustomItem

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 7 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Assist, EA_OpenPartyWindow, Enemy, PieTracker, Pure, SocialWindow 2.0, Squared |
| Files seen in | Code/UnitFrames/UnitFrame.lua, PieTracker.lua, Source/PureGroup.lua, SquaredGroup.lua, SquaredWarband.lua, assist.lua, source/openpartywindowtabmanage.lua, source/socialwindowtabfriends.lua |
| Namespaces detected | PlayerMenuWindow |
| Source kinds | lua_calls |
| Example locations | Assist: PlayerMenuWindowShowMenu, EA_OpenPartyWindow: ShowWarbandMemberMenu, Enemy: UnitFramesUI_UnitFrame_OnRButtonUp, PieTracker: RowClicked, Pure: OnRButtonUp, SocialWindow 2.0: OnRButtonUpPlayerRow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 42 |
| Global usage count | 42 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
PlayerMenuWindow.NewCustomItem(arg1, arg2, arg3)
```

## Description

Observed as a global function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GetString(StringTables.Default.LABEL_DEMOTE_ASSISTANT), GetString(StringTables.Default.LABEL_GROUP_OPTIONS), GetString(StringTables.Default.LABEL_LEAVE_GROUP) |
| arg2 | Observed as a function or method reference. | Observed values: Assist.Setat1, Assist.Setat2, Assist.Setat3 |
| arg3 | Observed as a boolean toggle. | Observed values: Enemy.groups.groupType==Enemy.GroupTypes.Warband, disableMainAssist, disablePlayerSummon |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Assist
- EA_OpenPartyWindow
- Enemy
- PieTracker
- Pure
- SocialWindow 2.0
- Squared

## Examples

- Assist: PlayerMenuWindowShowMenu -> PlayerMenuWindow.NewCustomItem(L "Set as Assist 1", Assist.Setat1, false)
- Assist: PlayerMenuWindowShowMenu -> PlayerMenuWindow.NewCustomItem(L "Set as Assist 2", Assist.Setat2, false)
- Assist: PlayerMenuWindowShowMenu -> PlayerMenuWindow.NewCustomItem(L "Set as Assist 3", Assist.Setat3, false)
- Assist: PlayerMenuWindowShowMenu -> PlayerMenuWindow.NewCustomItem(L "Set as Friendly 1", Assist.Setft1, false)
- Assist: PlayerMenuWindowShowMenu -> PlayerMenuWindow.NewCustomItem(L "Set as Friendly 2", Assist.Setft2, false)
- Assist: PlayerMenuWindowShowMenu -> PlayerMenuWindow.NewCustomItem(L "Set as Friendly 3", Assist.Setft3, false)

## Related APIs

- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
