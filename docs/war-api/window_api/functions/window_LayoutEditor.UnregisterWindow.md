# LayoutEditor.UnregisterWindow

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | DAoCBuff, Killer, TexturedButtons, TidyRoll, TurretRange, WhoHealedMe |
| Files seen in | `/workspace/data/raw/DAoCBuff/Source/DAoCBuffFrames.lua:653`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffHeadFrames.lua:212`, `/workspace/data/raw/Killer/KillerLifecycle.lua:100`, `/workspace/data/raw/TexturedButtons/TexturedButtons.lua:501`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:249`, `/workspace/data/raw/TurrentRange/Display.lua:244`, `/workspace/data/raw/WhoHealedMe/WhoHealedMe.lua:62` |
| Namespaces detected | LayoutEditor |
| Source kinds | lua_calls |
| Example locations | DAoCBuff: DAoCBuffHeadTracker:Shutdown, DAoCBuff: DAoCBuffTracker:Shutdown, Killer: Killer.Shutdown, TexturedButtons: TexturedButtons.RegisterLayoutEditorQuicklock, TidyRoll: TidyRoll.Shutdown, TurretRange: TurretRange.Display.Unload |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
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
LayoutEditor.UnregisterWindow(arg1)
```

## Description

Observed as a window function across 6 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: "KillerPersonalCounter", "KillerWindow", "WhoHealedMeWindow" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Removes a previously registered callback or command binding.

## Seen In

- DAoCBuff
- Killer
- TexturedButtons
- TidyRoll
- TurretRange
- WhoHealedMe

## Examples

- DAoCBuff: DAoCBuffHeadTracker:Shutdown -> LayoutEditor.UnregisterWindow(self.m_windowName)
- DAoCBuff: DAoCBuffTracker:Shutdown -> LayoutEditor.UnregisterWindow(self.m_windowName)
- Killer: Killer.Shutdown -> LayoutEditor.UnregisterWindow("KillerWindow")
- Killer: Killer.Shutdown -> LayoutEditor.UnregisterWindow("KillerPersonalCounter")
- TexturedButtons: TexturedButtons.RegisterLayoutEditorQuicklock -> LayoutEditor.UnregisterWindow(QUICK_LOCK_NAME)
- TidyRoll: TidyRoll.Shutdown -> LayoutEditor.UnregisterWindow(c_TIDY_ROLL_ANCHOR)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [SystemData.Events.CAMPAIGN_ZONE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.CAMPAIGN_ZONE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ZONE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ZONE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
