# WindowGetOffsetFromParent

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | DAoCBuff, PotionBar, TidyChat |
| Files seen in | `/workspace_addons/DAoCBuff/Source/DAoCBuffFrames.lua:653`, `/workspace_addons/PotionBar/source/Main.lua:292`, `/workspace_addons/TidyChat/TidyChat.lua:239` |
| Namespaces detected | WindowGetOffsetFromParent |
| Source kinds | lua_calls |
| Example locations | DAoCBuff: DAoCBuffTracker:Shutdown, PotionBar: PotionBar.Shutdown, TidyChat: TidyChatCore.SetWindowGroup |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
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
| windowName | Observed as a target window name. | Observed values: "PotionBarFloatingActivator", self.m_windowName, wndGroupName |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- DAoCBuff
- PotionBar
- TidyChat

## Examples

- DAoCBuff: DAoCBuffTracker:Shutdown -> WindowGetOffsetFromParent(self.m_windowName)
- PotionBar: PotionBar.Shutdown -> WindowGetOffsetFromParent("PotionBarFloatingActivator")
- TidyChat: TidyChatCore.SetWindowGroup -> WindowGetOffsetFromParent(wndGroupName)

## Related APIs

- none

## Used With

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_RANK_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_RANK_UPDATED.md) (HIGH 100/100) - SystemData Field
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- [EA_ChatWindowGroups](../../globals/tables/table_EA_ChatWindowGroups.md) (HIGH 100/100) - Global Table
- [SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_RANK_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_RANK_UPDATED.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
