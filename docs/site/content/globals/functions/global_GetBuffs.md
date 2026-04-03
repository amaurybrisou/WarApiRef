# GetBuffs

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 83/100
- Seen in: 23 addons

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
| Addons seen in | Aura, BuffHead, CCTV, DAoCBuff, EZGuard, Enemy, GCDsaver, GuardBot |
| Files seen in | CCTV.lua, Code/Core/Groups/Groups.lua, Core.lua, EZGuard.lua, GCDsaver.lua, GuardBot.lua, IHYTM.lua, NerfedEngine.lua |
| Namespaces detected | GetBuffs |
| Source kinds | lua_calls |
| Example locations | Aura: RefreshEffects, BuffHead: RefreshPlayerBuffs, BuffHead: ResyncTarget, CCTV: BuffUpdate, DAoCBuff: Refresh, DAoCBuff: Update |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 44 |
| Global usage count | 44 |
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

Observed as a shared query API across 23 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GameData.BuffTargetType.SELF, GameData.BuffTargetType.TARGET_FRIENDLY, GameData.BuffTargetType.TARGET_HOSTILE |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Aura
- BuffHead
- CCTV
- DAoCBuff
- EZGuard
- Enemy
- GCDsaver
- GuardBot
- HealGrid
- I HATE YOU THIS MUCH
- LibGuard
- NerfedButtons
- Phantom
- Pure
- RandomMount
- ResHelp
- SquaredHDIndicator
- TastyButtons
- TurretScrap
- Twister
- WSCT
- WarTriage
- WhatsCooking

## Examples

- Aura: RefreshEffects -> GetBuffs(self.m_effectTargetType)
- BuffHead: RefreshPlayerBuffs -> GetBuffs(GameData.BuffTargetType.SELF)
- BuffHead: ResyncTarget -> GetBuffs(buffTargetType)
- CCTV: BuffUpdate -> GetBuffs(GameData.BuffTargetType.SELF)
- DAoCBuff: Refresh -> GetBuffs(self.m_targetType)
- DAoCBuff: Update -> GetBuffs(self.m_targetType)

## Affects

- [GameData.BuffTargetType.SELF](../../gamedata/fields/gamedata_GameData.BuffTargetType.SELF.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
