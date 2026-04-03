# WindowGetOffsetFromParent

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

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
| Addons seen in | DAoCBuff, PotionBar, Statdoll Remix, TidyChat, Vectors |
| Files seen in | Settings.lua, Source/DAoCBuffFrames.lua, StatdollWnd.lua, StatdollWndLight.lua, TidyChat.lua, source/Main.lua |
| Namespaces detected | WindowGetOffsetFromParent |
| Source kinds | lua_calls |
| Example locations | DAoCBuff: Shutdown, PotionBar: Shutdown, Statdoll Remix: onShutdown, TidyChat: SetWindowGroup, Vectors: ShowSelectOption |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
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
WindowGetOffsetFromParent(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "CursorWindow", "PotionBarFloatingActivator", StatdollWnd.window |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- DAoCBuff
- PotionBar
- Statdoll Remix
- TidyChat
- Vectors

## Examples

- DAoCBuff: Shutdown -> WindowGetOffsetFromParent(self.m_windowName)
- PotionBar: Shutdown -> WindowGetOffsetFromParent("PotionBarFloatingActivator")
- Statdoll Remix: onShutdown -> WindowGetOffsetFromParent(StatdollWndLight.window)
- Statdoll Remix: onShutdown -> WindowGetOffsetFromParent(StatdollWnd.window)
- TidyChat: SetWindowGroup -> WindowGetOffsetFromParent(wndGroupName)
- Vectors: ShowSelectOption -> WindowGetOffsetFromParent("CursorWindow")

## Related APIs

- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Used With

- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Affects

- [SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_RANK_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_RANK_UPDATED.md) (HIGH 100/100) - SystemData Field

## Notes

- Advanced return analysis: No strong return evidence observed
