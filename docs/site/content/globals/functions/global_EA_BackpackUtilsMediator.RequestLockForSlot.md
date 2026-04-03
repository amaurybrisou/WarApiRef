# EA_BackpackUtilsMediator.RequestLockForSlot

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 90/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Motion |
| Files seen in | Source/Motion.lua |
| Namespaces detected | EA_BackpackUtilsMediator |
| Source kinds | lua_calls |
| Example locations | Motion: performCraft |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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
EA_BackpackUtilsMediator.RequestLockForSlot(arg1, arg2, arg3, arg4)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: containerSlot, determinantSlot, resourceSlots[1].slot |
| arg2 | Observed as a function or method reference. | Observed values: containerType, determinantType, resourceSlots[1].type |
| arg3 | Observed as a runtime window or control identifier. | Observed values: windowId |
| arg4 | Observed as a runtime window or control identifier. | Observed values: {r=0,g=255,b=0} |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Motion

## Examples

- Motion: performCraft -> EA_BackpackUtilsMediator.RequestLockForSlot(containerSlot, containerType, windowId, {r=0,g=255,b=0})
- Motion: performCraft -> EA_BackpackUtilsMediator.RequestLockForSlot(determinantSlot, determinantType, windowId, {r=0,g=255,b=0})
- Motion: performCraft -> EA_BackpackUtilsMediator.RequestLockForSlot(resourceSlots[1].slot, resourceSlots[1].type, windowId, {r=0,g=255,b=0})
- Motion: performCraft -> EA_BackpackUtilsMediator.RequestLockForSlot(resourceSlots[2].slot, resourceSlots[2].type, windowId, {r=0,g=255,b=0})
- Motion: performCraft -> EA_BackpackUtilsMediator.RequestLockForSlot(resourceSlots[3].slot, resourceSlots[3].type, windowId, {r=0,g=255,b=0})

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
