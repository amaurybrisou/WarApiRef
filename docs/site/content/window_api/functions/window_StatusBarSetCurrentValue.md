# StatusBarSetCurrentValue

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
| Addons seen in | Ace, Effigy, GCDsaver, LibWBToggler, RoR_SoR, Shinies, Targets, TidyRoll |
| Files seen in | `/workspace/Ace/LibGUI.lua:996`, `/workspace/Effigy/LibGUI.lua:993`, `/workspace/GCDsaver/libs/LibGUI.lua:993`, `/workspace/LibWarBoardToggler/libs/LibGUI.lua:993`, `/workspace/RoR_SoR/RoR_SoR.lua:990`, `/workspace/Shinies/Libraries/LibGUI.lua:993`, `/workspace/TidyRoll/TidyRollFrame.lua:115`, `/workspace/WarTriage/libs/LibGUI.lua:993` |
| Namespaces detected | StatusBarSetCurrentValue |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Statusbar:SetValue, Effigy: LIBGUI_Statusbar:SetValue, GCDsaver: LIBGUI_Statusbar:SetValue, LibWBToggler: LIBGUI_Statusbar:SetValue, RoR_SoR: RoR_SoR.T1_UPDATE, Shinies: LIBGUI_Statusbar:SetValue |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
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

Observed as a window function across 10 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "SoR_"..Window_Name.."VPDESTROBAR", "SoR_"..Window_Name.."VPORDERBAR", self.name |
| arg2 | Observed as a runtime window or control identifier. | Observed values: player.hp, timer, tonumber(ZoneVP.controlPoints[1]) |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Effigy
- GCDsaver
- LibWBToggler
- RoR_SoR
- Shinies
- Targets
- TidyRoll
- WarTriage
- WoH-Reticle

## Examples

- Ace: LIBGUI_Statusbar:SetValue -> StatusBarSetCurrentValue(self.name, value)
- Effigy: LIBGUI_Statusbar:SetValue -> StatusBarSetCurrentValue(self.name, value)
- GCDsaver: LIBGUI_Statusbar:SetValue -> StatusBarSetCurrentValue(self.name, value)
- LibWBToggler: LIBGUI_Statusbar:SetValue -> StatusBarSetCurrentValue(self.name, value)
- RoR_SoR: RoR_SoR.T1_UPDATE -> StatusBarSetCurrentValue("SoR_"..Window_Name.."VPORDERBAR", tonumber(ZoneVP.controlPoints[1]))
- RoR_SoR: RoR_SoR.T1_UPDATE -> StatusBarSetCurrentValue("SoR_"..Window_Name.."VPDESTROBAR", tonumber(ZoneVP.controlPoints[2]))

## Related APIs

- none

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [StatusBarSetForegroundTint](window_StatusBarSetForegroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
