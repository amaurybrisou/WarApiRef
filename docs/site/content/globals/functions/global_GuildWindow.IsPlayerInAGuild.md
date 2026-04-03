# GuildWindow.IsPlayerInAGuild

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 83/100
- Seen in: 4 addons

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
| Addons seen in | FozAuction, Rolodex, rorAutoInviter, zMailMod |
| Files seen in | Rolodex.lua, Source/auctionwindow.lua, rorAutoInviter.lua, zMailModMassMail.lua |
| Namespaces detected | GuildWindow |
| Source kinds | lua_calls |
| Example locations | FozAuction: PlayerCanSearchAllianceAuctions, FozAuction: PlayerCanSearchGuildAuctions, Rolodex: UpdateList, rorAutoInviter: OnInitialize, zMailMod: GetNameList |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
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
GuildWindow.IsPlayerInAGuild()
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- FozAuction
- Rolodex
- rorAutoInviter
- zMailMod

## Examples

- FozAuction: PlayerCanSearchAllianceAuctions -> GuildWindow.IsPlayerInAGuild()
- FozAuction: PlayerCanSearchGuildAuctions -> GuildWindow.IsPlayerInAGuild()
- Rolodex: UpdateList -> GuildWindow.IsPlayerInAGuild()
- rorAutoInviter: OnInitialize -> GuildWindow.IsPlayerInAGuild()
- zMailMod: GetNameList -> GuildWindow.IsPlayerInAGuild()

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event

## Affects

- [SystemData.Events.CONVERSATION_TEXT_ARRIVED](../../systemdata/fields/systemdata_SystemData.Events.CONVERSATION_TEXT_ARRIVED.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
