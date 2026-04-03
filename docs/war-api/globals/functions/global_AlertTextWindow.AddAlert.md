# AlertTextWindow.AddAlert

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 8 addons

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
| Addons seen in | BlackBook, DammazKron, LootAlert, MapPin, Minmap, PlanB, SNT_BUTTONS, SNT_INFO |
| Files seen in | BlackBook.lua, Core/DK_Core.lua, LootAlert.lua, PlanB.lua, core.lua, snt_buttons.lua, snt_info.lua, source/MapPin.lua |
| Namespaces detected | AlertTextWindow |
| Source kinds | lua_calls |
| Example locations | BlackBook: CheckOdds, BlackBook: PlayerRenownUpdated, DammazKron: AlertRegister, LootAlert: GetLoot, MapPin: test, Minmap: ActivateRallyCall |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 14 |
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
AlertTextWindow.AddAlert()
```

## Description

Observed as a global function across 8 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- BlackBook
- DammazKron
- LootAlert
- MapPin
- Minmap
- PlanB
- SNT_BUTTONS
- SNT_INFO

## Examples

- BlackBook: CheckOdds -> AlertTextWindow.AddAlert()
- BlackBook: PlayerRenownUpdated -> AlertTextWindow.AddAlert()
- DammazKron: AlertRegister -> AlertTextWindow.AddAlert()
- LootAlert: GetLoot -> AlertTextWindow.AddAlert()
- MapPin: test -> AlertTextWindow.AddAlert()
- Minmap: ActivateRallyCall -> AlertTextWindow.AddAlert()

## Used With

- [ButtonSetTextureSlice](../../window_api/functions/window_ButtonSetTextureSlice.md) (HIGH 100/100) - Window Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function

## Affects

- [GameData.ChatData.name](../../gamedata/fields/gamedata_GameData.ChatData.name.md) (HIGH 100/100) - GameData Field
- [GameData.ChatData.text](../../gamedata/fields/gamedata_GameData.ChatData.text.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.AUCTION_HOUSE_CREATE_AUCTION](../../gamedata/fields/gamedata_GameData.Sound.AUCTION_HOUSE_CREATE_AUCTION.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.QUEST_ABANDONED](../../gamedata/fields/gamedata_GameData.Sound.QUEST_ABANDONED.md) (HIGH 100/100) - GameData Field
- [SystemData.AlertText.VecText](../../systemdata/fields/systemdata_SystemData.AlertText.VecText.md) (HIGH 100/100) - SystemData Field
- [SystemData.AlertText.VecType](../../systemdata/fields/systemdata_SystemData.AlertText.VecType.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
