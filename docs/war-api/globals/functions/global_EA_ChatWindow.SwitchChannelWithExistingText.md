# EA_ChatWindow.SwitchChannelWithExistingText

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | EA_OpenPartyWindow, QuickNameActions+, SocialWindow 2.0, Squared |
| Files seen in | QuickNameActionsRessurected.lua, SquaredWarband.lua, source/openpartywindowtabworld.lua, source/socialwindow.lua |
| Namespaces detected | EA_ChatWindow |
| Source kinds | lua_calls |
| Example locations | EA_OpenPartyWindow: JoinPartySpecified, QuickNameActions+: newEA_ChatWindowOnTomeLinkLButtonUp, SocialWindow 2.0: SendTell, Squared: OnMenuClickTellMember |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
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
EA_ChatWindow.SwitchChannelWithExistingText(arg1)
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: L "/tell "..SystemData.UserInput.selectedGroupMember..L " ", linkData, text |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- EA_OpenPartyWindow
- QuickNameActions+
- SocialWindow 2.0
- Squared

## Examples

- EA_OpenPartyWindow: JoinPartySpecified -> EA_ChatWindow.SwitchChannelWithExistingText(text)
- QuickNameActions+: newEA_ChatWindowOnTomeLinkLButtonUp -> EA_ChatWindow.SwitchChannelWithExistingText(linkData)
- SocialWindow 2.0: SendTell -> EA_ChatWindow.SwitchChannelWithExistingText(text)
- Squared: OnMenuClickTellMember -> EA_ChatWindow.SwitchChannelWithExistingText(L "/tell "..SystemData.UserInput.selectedGroupMember..L " ")

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.UserInput.selectedGroupMember](../../systemdata/fields/systemdata_SystemData.UserInput.selectedGroupMember.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
