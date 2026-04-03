# StatusBarSetCurrentValue

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
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:996`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:993`, `/workspace/data/raw/PartyCast/PartyCast.lua:399`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:993`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:990`, `/workspace/data/raw/Shinies/Libraries/LibGUI.lua:993`, `/workspace/data/raw/TidyRoll/TidyRollFrame.lua:115`, `/workspace/data/raw/WoH-Reticle/libs/LibGUI.lua:993` |
| Namespaces detected | StatusBarSetCurrentValue |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Statusbar:SetValue, LibWBToggler: LIBGUI_Statusbar:SetValue, PartyCast: LIBGUI_Statusbar:SetValue, PartyCast: PartyCast.FetchedText, RoR_SoR: RoR_SoR.T1_UPDATE, Shinies: LIBGUI_Statusbar:SetValue |
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
StatusBarSetCurrentValue(arg1, arg2)
```

## Description

Observed as a window function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "PartyCastWindow"..PlayerNumber.."TimerBar", "SoR_"..Window_Name.."VPDESTROBAR", "SoR_"..Window_Name.."VPORDERBAR" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 1, timer, tonumber(ZoneVP.controlPoints[1]) |

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

- Ace: LIBGUI_Statusbar:SetValue -> StatusBarSetCurrentValue(self.name, value)
- LibWBToggler: LIBGUI_Statusbar:SetValue -> StatusBarSetCurrentValue(self.name, value)
- PartyCast: LIBGUI_Statusbar:SetValue -> StatusBarSetCurrentValue(self.name, value)
- PartyCast: PartyCast.FetchedText -> StatusBarSetCurrentValue("PartyCastWindow"..PlayerNumber.."TimerBar", 1)
- RoR_SoR: RoR_SoR.T1_UPDATE -> StatusBarSetCurrentValue("SoR_"..Window_Name.."VPORDERBAR", tonumber(ZoneVP.controlPoints[1]))
- RoR_SoR: RoR_SoR.T1_UPDATE -> StatusBarSetCurrentValue("SoR_"..Window_Name.."VPDESTROBAR", tonumber(ZoneVP.controlPoints[2]))

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
