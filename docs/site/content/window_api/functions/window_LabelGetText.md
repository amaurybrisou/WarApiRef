# LabelGetText

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
| Addons seen in | Ace, DAoCBuff, LibWBToggler, PartyCast, RoR_SoR, Shinies, WoH-Reticle |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:434`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettings.lua:215`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:434`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:434`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:1055`, `/workspace/data/raw/Shinies/Libraries/LibGUI.lua:434`, `/workspace/data/raw/WoH-Reticle/libs/LibGUI.lua:434` |
| Namespaces detected | LabelGetText |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Label:GetText, DAoCBuff: DAoCBuffSettings.FilterRowOnLButtonUp, LibWBToggler: LIBGUI_Label:GetText, PartyCast: LIBGUI_Label:GetText, RoR_SoR: RoR_SoR.SET_KEEP, Shinies: LIBGUI_Label:GetText |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
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
LabelGetText(arg1)
```

## Description

Observed as a window function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "SoR_"..Window_Name.."KEEP1HEALTH", "SoR_"..Window_Name.."KEEP1SAFETIMER", "SoR_"..Window_Name.."KEEP2HEALTH" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- DAoCBuff
- LibWBToggler
- PartyCast
- RoR_SoR
- Shinies
- WoH-Reticle

## Examples

- Ace: LIBGUI_Label:GetText -> LabelGetText(self.name)
- DAoCBuff: DAoCBuffSettings.FilterRowOnLButtonUp -> LabelGetText(SystemData.ActiveWindow.name.."Name")
- LibWBToggler: LIBGUI_Label:GetText -> LabelGetText(self.name)
- PartyCast: LIBGUI_Label:GetText -> LabelGetText(self.name)
- RoR_SoR: RoR_SoR.SET_KEEP -> LabelGetText("SoR_"..Window_Name.."KEEP1SAFETIMER")
- RoR_SoR: RoR_SoR.SET_KEEP -> LabelGetText("SoR_"..Window_Name.."KEEP2SAFETIMER")

## Related APIs

- none

## Used With

- none

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [GameData.Guild.m_GuildName](../../gamedata/fields/gamedata_GameData.Guild.m_GuildName.md) (HIGH 100/100) - GameData Field
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.GuildPermissons.CLAIM_KEEP](../../systemdata/fields/systemdata_SystemData.GuildPermissons.CLAIM_KEEP.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
