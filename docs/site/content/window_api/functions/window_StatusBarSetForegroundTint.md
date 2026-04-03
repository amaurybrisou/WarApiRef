# StatusBarSetForegroundTint

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

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
| Addons seen in | Ace, LibWBToggler, PartyCast, RoR_SoR, Shinies, TidyRoll, WoH-Reticle |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:1022`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:1019`, `/workspace/data/raw/PartyCast/PartyCast.lua:399`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1019`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:990`, `/workspace/data/raw/Shinies/Libraries/LibGUI.lua:1019`, `/workspace/data/raw/TidyRoll/TidyRollFrame.lua:105`, `/workspace/data/raw/WoH-Reticle/libs/LibGUI.lua:1019` |
| Namespaces detected | StatusBarSetForegroundTint |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Statusbar:ForeColor, LibWBToggler: LIBGUI_Statusbar:ForeColor, PartyCast: LIBGUI_Statusbar:ForeColor, PartyCast: PartyCast.FetchedText, RoR_SoR: RoR_SoR.T1_UPDATE, Shinies: LIBGUI_Statusbar:ForeColor |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
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

Observed as a window function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "PartyCastWindow"..PlayerNumber.."TimerBar", "SoR_"..Window_Name.."VPDESTROBAR", "SoR_"..Window_Name.."VPORDERBAR" |
| arg2 | Observed as a function or method reference. | Observed values: DefaultColor.GREEN.r, PartyCast.Settings.Colors.Casting[1], PartyCast.Settings.Colors.Failed[1] |
| arg3 | Observed as a function or method reference. | Observed values: DefaultColor.GREEN.g, PartyCast.Settings.Colors.Casting[2], PartyCast.Settings.Colors.Failed[2] |
| arg4 | Observed as a function or method reference. | Observed values: DefaultColor.GREEN.b, PartyCast.Settings.Colors.Casting[3], PartyCast.Settings.Colors.Failed[3] |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- LibWBToggler
- PartyCast
- RoR_SoR
- Shinies
- TidyRoll
- WoH-Reticle

## Examples

- Ace: LIBGUI_Statusbar:ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)
- LibWBToggler: LIBGUI_Statusbar:ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)
- PartyCast: LIBGUI_Statusbar:ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)
- PartyCast: PartyCast.FetchedText -> StatusBarSetForegroundTint("PartyCastWindow"..PlayerNumber.."TimerBar", PartyCast.Settings.Colors.Casting[1], PartyCast.Settings.Colors.Casting[2], PartyCast.Settings.Colors.Casting[3])
- PartyCast: PartyCast.FetchedText -> StatusBarSetForegroundTint("PartyCastWindow"..PlayerNumber.."TimerBar", PartyCast.Settings.Colors.Instant[1], PartyCast.Settings.Colors.Instant[2], PartyCast.Settings.Colors.Instant[3])
- PartyCast: PartyCast.FetchedText -> StatusBarSetForegroundTint("PartyCastWindow"..PlayerNumber.."TimerBar", PartyCast.Settings.Colors.Failed[1], PartyCast.Settings.Colors.Failed[2], PartyCast.Settings.Colors.Failed[3])

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [PartyUtils.GetPartyData](../../globals/functions/global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
