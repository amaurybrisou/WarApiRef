# StatusBarSetForegroundTint

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 8 addons

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
| Addons seen in | Ace, Effigy, GCDsaver, LibWBToggler, RoR_SoR, Shinies, TidyRoll, WoH-Reticle |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:1022`, `/workspace_addons/Effigy/LibGUI.lua:1019`, `/workspace_addons/GCDsaver/libs/LibGUI.lua:1019`, `/workspace_addons/LibWarBoardToggler/libs/LibGUI.lua:1019`, `/workspace_addons/RoR_SoR/RoR_SoR.lua:990`, `/workspace_addons/Shinies/Libraries/LibGUI.lua:1019`, `/workspace_addons/TidyRoll/TidyRollFrame.lua:105`, `/workspace_addons/WoH-Reticle/libs/LibGUI.lua:1019` |
| Namespaces detected | StatusBarSetForegroundTint |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Statusbar:ForeColor, Effigy: LIBGUI_Statusbar:ForeColor, GCDsaver: LIBGUI_Statusbar:ForeColor, LibWBToggler: LIBGUI_Statusbar:ForeColor, RoR_SoR: RoR_SoR.T1_UPDATE, Shinies: LIBGUI_Statusbar:ForeColor |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
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
StatusBarSetForegroundTint(arg1, arg2, arg3, arg4)
```

## Description

Observed as a window function across 8 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "SoR_"..Window_Name.."VPDESTROBAR", "SoR_"..Window_Name.."VPORDERBAR", self.name |
| arg2 | Observed as a runtime window or control identifier. | Observed values: DefaultColor.GREEN.r, RoR_SoR.RealmColors[2].r, RoR_SoR.RealmColors[3].r |
| arg3 | Observed as a runtime window or control identifier. | Observed values: DefaultColor.GREEN.g, RoR_SoR.RealmColors[2].g, RoR_SoR.RealmColors[3].g |
| arg4 | Observed as a runtime window or control identifier. | Observed values: DefaultColor.GREEN.b, RoR_SoR.RealmColors[2].b, RoR_SoR.RealmColors[3].b |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Effigy
- GCDsaver
- LibWBToggler
- RoR_SoR
- Shinies
- TidyRoll
- WoH-Reticle

## Examples

- Ace: LIBGUI_Statusbar:ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)
- Effigy: LIBGUI_Statusbar:ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)
- GCDsaver: LIBGUI_Statusbar:ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)
- LibWBToggler: LIBGUI_Statusbar:ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)
- RoR_SoR: RoR_SoR.T1_UPDATE -> StatusBarSetForegroundTint("SoR_"..Window_Name.."VPORDERBAR", RoR_SoR.RealmColors[2].r, RoR_SoR.RealmColors[2].g, RoR_SoR.RealmColors[2].b)
- RoR_SoR: RoR_SoR.T1_UPDATE -> StatusBarSetForegroundTint("SoR_"..Window_Name.."VPDESTROBAR", RoR_SoR.RealmColors[3].r, RoR_SoR.RealmColors[3].g, RoR_SoR.RealmColors[3].b)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
