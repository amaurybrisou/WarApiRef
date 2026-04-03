# PartyUtils.IsPartyActive

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

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
| Addons seen in | GuardLine, LibGuard, PartyCast, RoR_SoR, Swift Assist |
| Files seen in | `/workspace/data/raw/GuardLine/GuardLine.lua:197`, `/workspace/data/raw/LibGuard/Source/LibGuard.lua:186`, `/workspace/data/raw/LibGuard/Source/LibGuard.lua:362`, `/workspace/data/raw/PartyCast/PartyCast.lua:655`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:2468`, `/workspace/data/raw/swift-assist/SwiftAssist.lua:71` |
| Namespaces detected | PartyUtils |
| Source kinds | globals, lua_calls |
| Example locations | GuardLine: GuardLine.update, LibGuard: LibGuard.GetIdFromName, LibGuard: LibGuard.UpdateStateMachine, PartyCast: PartyCast.Update, RoR_SoR: RoR_SoR.ContextMenu1, Swift Assist: GetChannel |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
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
PartyUtils.IsPartyActive()
```

## Description

Observed as a global function across 5 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- GuardLine
- LibGuard
- PartyCast
- RoR_SoR
- Swift Assist

## Examples

- GuardLine: GuardLine.update -> PartyUtils.IsPartyActive()
- LibGuard: LibGuard.GetIdFromName -> PartyUtils.IsPartyActive()
- LibGuard: LibGuard.UpdateStateMachine -> PartyUtils.IsPartyActive()
- PartyCast: PartyCast.Update -> PartyUtils.IsPartyActive()
- RoR_SoR: RoR_SoR.ContextMenu1 -> PartyUtils.IsPartyActive()
- Swift Assist: GetChannel -> PartyUtils.IsPartyActive()

## Related APIs

- none

## Used With

- [PartyUtils.GetPartyData](global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function
- [StatusBarGetCurrentValue](../../window_api/functions/window_StatusBarGetCurrentValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [StatusBarGetMaximumValue](../../window_api/functions/window_StatusBarGetMaximumValue.md) (HIGH 80/100) - Window Function

## Triggered By

- none

## Affects

- [GameData.Player.Pet.name](../../gamedata/fields/gamedata_GameData.Player.Pet.name.md) (HIGH 100/100) - GameData Field
- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [GameData.Player.rvrZoneFlagged](../../gamedata/fields/gamedata_GameData.Player.rvrZoneFlagged.md) (HIGH 100/100) - GameData Field
- [PartyUtils.GetPartyData](global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandData](global_PartyUtils.GetWarbandData.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
