# GetBuffs

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 83/100
- Seen in: 10 addons

## Confidence Assessment

- Level: HIGH

- Score: 83/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Aura, BuffHead, DAoCBuff, Enemy, GCDsaver, LibGuard, RandomMount, Twister |
| Files seen in | `/workspace/Aura/Source/AuraEffectTracker.lua:62`, `/workspace/BuffHead/Core.lua:1138`, `/workspace/BuffHead/Core.lua:178`, `/workspace/DAoCBuff/Source/DAoCBuffFrames.lua:1099`, `/workspace/DAoCBuff/Source/DAoCBuffFrames.lua:1245`, `/workspace/DAoCBuff/Source/DAoCBuffHeadFrames.lua:780`, `/workspace/Enemy/Code/Core/Groups/Groups.lua:503`, `/workspace/GCDsaver/GCDsaver.lua:42` |
| Namespaces detected | GetBuffs |
| Source kinds | lua_calls |
| Example locations | Aura: AuraEffectTracker:RefreshEffects, BuffHead: BuffHead.RefreshPlayerBuffs, BuffHead: BuffHead.local.ResyncTarget, BuffHead: ResyncTarget, DAoCBuff: DAoCBuffHeadTracker:Refresh, DAoCBuff: DAoCBuffTracker:Refresh |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 17 |
| Global usage count | 17 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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
GetBuffs(arg1)
```

## Description

Observed as a shared query API across 10 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GameData.BuffTargetType.SELF, buffTargetType, self.m_effectTargetType |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- Aura
- BuffHead
- DAoCBuff
- Enemy
- GCDsaver
- LibGuard
- RandomMount
- Twister
- WSCT
- WarTriage

## Examples

- Aura: AuraEffectTracker:RefreshEffects -> GetBuffs(self.m_effectTargetType)
- BuffHead: BuffHead.RefreshPlayerBuffs -> GetBuffs(GameData.BuffTargetType.SELF)
- BuffHead: BuffHead.local.ResyncTarget -> GetBuffs(buffTargetType)
- BuffHead: ResyncTarget -> GetBuffs(buffTargetType)
- DAoCBuff: DAoCBuffHeadTracker:Refresh -> GetBuffs(self.m_targetType)
- DAoCBuff: DAoCBuffTracker:Refresh -> GetBuffs(self.m_targetType)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [GameData.BuffTargetType.SELF](../../gamedata/fields/gamedata_GameData.BuffTargetType.SELF.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
