# GameData.BuffTargetType.SELF

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 22 addons

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
| Addons seen in | BBars - Mechanic Only, BuffHead, CCTV, DuffTimer, EZGuard, Enemy, GuardBot, HealGrid |
| Files seen in | BBarsPlayerMechanic.lua, CCTV.lua, Code/Core/Groups/Groups.lua, Core.lua, DuffTimer.lua, EZGuard.lua, GuardBot.lua, IHYTM.lua |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | BOSMMechanicUpdate, BuffUpdate, Button_EffectsUpdated, CHSLMechanicUpdate, CheckBuffs, CheckCast |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 30 |
| Global usage count | 30 |
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

GameData.GameData.BuffTargetType.SELF field accessed by 22 addons; commonly found in BOSMMechanicUpdate and BuffUpdate, Button_EffectsUpdated, CHSLMechanicUpdate, CheckBuffs, CheckCast, CheckSquigArmor, EnforceWindowStates, GetGuardBuffStatus, GetHurtPlayer, InitializeEffectTracker, OnLoad, OnPlayerEffectsUpdated, PLAYER_EFFECTS_UPDATED, RefreshEffects, RefreshInformation, RefreshMemberEffects, RefreshPlayerBuffs, UpdateAggroTable, UpdateEffect, UpdateEffects, UpdateGuardTarget, UpdateHealDebuffs, WEWHMechanicUpdate, buildEffectsAndConditions, forceupdate, handler, hasMechanic, init, lua_call contexts.

## Seen In

- BBars - Mechanic Only
- BuffHead
- CCTV
- DuffTimer
- EZGuard
- Enemy
- GuardBot
- HealGrid
- I HATE YOU THIS MUCH
- LibGuard
- Mech
- NerfedButtons
- Phantom
- Pure
- RandomMount
- SquaredHDIndicator
- TastyButtons
- TurretScrap
- Twister
- WSCT
- WarTriage
- WhatsCooking

## Related APIs

- [GetBuffs](../../globals/functions/global_GetBuffs.md) (HIGH 83/100) - Global Function

## Notes

- Observed in contexts: BOSMMechanicUpdate, BuffUpdate, Button_EffectsUpdated, CHSLMechanicUpdate, CheckBuffs, CheckCast
