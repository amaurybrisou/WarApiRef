# LayoutEditor.RegisterEditCallback

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 14 addons

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
| Addons seen in | CDown, Crusher, DAoCBuff, DeepSleep, DuffTimer, Effigy, FixGit, FlagCap |
| Files seen in | CDown.lua, DeepSleep.lua, DuffTimer.lua, Effigy.lua, Enhancements.lua, FlagCap.lua, GuardLine.lua, SOR.lua |
| Namespaces detected | LayoutEditor |
| Source kinds | lua_calls |
| Example locations | CDown: Initialize, Crusher: OnLoad, DAoCBuff: Initialize, DeepSleep: Initialize, DuffTimer: OnLoadingEnd, Effigy: Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 14 |
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
LayoutEditor.RegisterEditCallback(arg1)
```

## Description

Observed as a window function across 14 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: Addon.LayoutEditorDone, CDown.LEH, Crusher.OnLayoutEditorEvent |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- CDown
- Crusher
- DAoCBuff
- DeepSleep
- DuffTimer
- Effigy
- FixGit
- FlagCap
- GuardLine
- Kwestor
- Motion
- RetAlert
- SOR
- Vectors

## Examples

- CDown: Initialize -> LayoutEditor.RegisterEditCallback(CDown.LEH)
- Crusher: OnLoad -> LayoutEditor.RegisterEditCallback(Crusher.OnLayoutEditorEvent)
- DAoCBuff: Initialize -> LayoutEditor.RegisterEditCallback(DAoCBuff.LEH)
- DeepSleep: Initialize -> LayoutEditor.RegisterEditCallback(DeepSleep.LayoutEditorDone)
- DuffTimer: OnLoadingEnd -> LayoutEditor.RegisterEditCallback(DuffTimer.LayoutEditorCallback)
- Effigy: Initialize -> LayoutEditor.RegisterEditCallback(Addon.LayoutEditorDone)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event

## Affects

- [GameData.Account.ServerName](../../gamedata/fields/gamedata_GameData.Account.ServerName.md) (HIGH 100/100) - GameData Field
- [GameData.Player.money](../../gamedata/fields/gamedata_GameData.Player.money.md) (HIGH 100/100) - GameData Field
- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [SystemData.Events.ALL_MODULES_INITIALIZED](../../systemdata/fields/systemdata_SystemData.Events.ALL_MODULES_INITIALIZED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CUSTOM_UI_SCALE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.CUSTOM_UI_SCALE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COOLDOWN_TIMER_SET](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COOLDOWN_TIMER_SET.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH_CLEARED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH_CLEARED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HOT_BAR_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HOT_BAR_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELEASE_CORPSE](../../systemdata/fields/systemdata_SystemData.Events.RELEASE_CORPSE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- none
