# WindowSetOffsetFromParent

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 12 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BBars - Mechanic Only, CoolDownLine, DammazKron, FlagCap, GetStats, Group Icons SG, KeyBar, MapPin |
| Files seen in | BBarsPetHP.lua, BBarsPlayerMechanic.lua, CoolDownLine.lua, Core/Tome/DK_Tome.lua, FlagCap.lua, GetStats.lua, GroupIconsSG.lua, KeyBar.lua |
| Namespaces detected | WindowSetOffsetFromParent |
| Source kinds | lua_calls |
| Example locations | BBars - Mechanic Only: MechDraw, BBars - Mechanic Only: PetHPDraw, CoolDownLine: OnUpdate, DammazKron: InitializeBookmark, FlagCap: OnUpdate, GetStats: OnChatLogUpdated |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 40 |
| Global usage count | 40 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
WindowSetOffsetFromParent(windowName, arg2, arg3)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "10secBG2", "30secBG2", "60secBG2" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: ((-10)+(width*(controlFill/100))), ((GetHotbarCooldown(i)/10)*(MaxLength*(2-(GetHotbarCooldown(i)/15))))*0.1875, (MaxLength*0.125)+((GetHotbarCooldown(i)/20)*(MaxLength*0.25)) |
| arg3 | Observed as a runtime window or control identifier. | Observed values: (PosY/UIScale)-(height/2), (UnitBuffer/2), -3 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- BBars - Mechanic Only
- CoolDownLine
- DammazKron
- FlagCap
- GetStats
- Group Icons SG
- KeyBar
- MapPin
- RO-Style Combat Text
- RoR_SoR
- RoR_debolster
- Twister

## Examples

- BBars - Mechanic Only: MechDraw -> WindowSetOffsetFromParent("BBarsPlayerMechanic1Grey", (UnitBuffer/2), (UnitBuffer/2))
- BBars - Mechanic Only: MechDraw -> WindowSetOffsetFromParent("BBarsPlayerMechanic2Grey", (UnitBuffer/2), (UnitBuffer/2))
- BBars - Mechanic Only: MechDraw -> WindowSetOffsetFromParent("BBarsPlayerMechanic3Grey", (UnitBuffer/2), (UnitBuffer/2))
- BBars - Mechanic Only: MechDraw -> WindowSetOffsetFromParent("BBarsPlayerMechanic4Grey", (UnitBuffer/2), (UnitBuffer/2))
- BBars - Mechanic Only: MechDraw -> WindowSetOffsetFromParent("BBarsPlayerMechanic5Grey", (UnitBuffer/2), (UnitBuffer/2))
- BBars - Mechanic Only: MechDraw -> WindowSetOffsetFromParent("BBarsPlayerMechanic1Frontbar", (UnitBuffer/2), (UnitBuffer/2))

## Related APIs

- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Notes

- Advanced return analysis: No strong return evidence observed
