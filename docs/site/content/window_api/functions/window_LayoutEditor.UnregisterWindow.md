# LayoutEditor.UnregisterWindow

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 10 addons

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
| Addons seen in | DAoCBuff, Effigy, Killer, Miracle Grow Remix, QuickWarReport, RVMOD_Manager, TexturedButtons, TidyRoll |
| Files seen in | `/workspace_addons/DAoCBuff/Source/DAoCBuffFrames.lua:653`, `/workspace_addons/DAoCBuff/Source/DAoCBuffHeadFrames.lua:212`, `/workspace_addons/Effigy/Elements/EffigyBar.lua:135`, `/workspace_addons/Effigy/Elements/EffigyBar.lua:640`, `/workspace_addons/Killer/KillerLifecycle.lua:100`, `/workspace_addons/MGRemix/config.lua:87`, `/workspace_addons/MGRemix/preset.lua:298`, `/workspace_addons/QuickWarReport/QWRLifecycle.lua:90` |
| Namespaces detected | LayoutEditor |
| Source kinds | lua_calls |
| Example locations | DAoCBuff: DAoCBuffHeadTracker:Shutdown, DAoCBuff: DAoCBuffTracker:Shutdown, Effigy: EffigyBar:destroy, Effigy: EffigyBar:setup, Killer: Killer.Shutdown, Miracle Grow Remix: MiracleGrow2.ConfigCallback |
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
LayoutEditor.UnregisterWindow(arg1)
```

## Description

Observed as a window function across 10 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "KillerPersonalCounter", "KillerWindow", "QuickWarReportBar" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Removes a previously registered callback or command binding.

## Seen In

- DAoCBuff
- Effigy
- Killer
- Miracle Grow Remix
- QuickWarReport
- RVMOD_Manager
- TexturedButtons
- TidyRoll
- TurretRange
- WhoHealedMe

## Examples

- DAoCBuff: DAoCBuffHeadTracker:Shutdown -> LayoutEditor.UnregisterWindow(self.m_windowName)
- DAoCBuff: DAoCBuffTracker:Shutdown -> LayoutEditor.UnregisterWindow(self.m_windowName)
- Effigy: EffigyBar:destroy -> LayoutEditor.UnregisterWindow(self.name)
- Effigy: EffigyBar:setup -> LayoutEditor.UnregisterWindow(self.name)
- Killer: Killer.Shutdown -> LayoutEditor.UnregisterWindow("KillerWindow")
- Killer: Killer.Shutdown -> LayoutEditor.UnregisterWindow("KillerPersonalCounter")

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [SystemData.Events.CAMPAIGN_ZONE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.CAMPAIGN_ZONE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ZONE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ZONE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
