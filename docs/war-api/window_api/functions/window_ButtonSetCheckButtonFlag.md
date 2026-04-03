# ButtonSetCheckButtonFlag

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
| Addons seen in | Ace, Killer, LibWBToggler, PartyCast, Shinies, WoH-Reticle |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:735`, `/workspace/data/raw/Ace/LibGUI.lua:805`, `/workspace/data/raw/Killer/KillerConfigWindow.lua:306`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:734`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:803`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:734`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:803`, `/workspace/data/raw/Shinies/Libraries/LibGUI.lua:734` |
| Namespaces detected | ButtonSetCheckButtonFlag |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Checkbox:New, Ace: LIBGUI_Optionbutton:New, Killer: EnsureSettingsControlsInitialized, Killer: Killer.local.EnsureSettingsControlsInitialized, LibWBToggler: LIBGUI_Checkbox:New, LibWBToggler: LIBGUI_Optionbutton:New |
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
ButtonSetCheckButtonFlag(arg1, arg2)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: SETTINGS_PERSONAL_CHECKBOX, w.name |
| arg2 | Observed as a boolean toggle. | Observed values: true |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Killer
- LibWBToggler
- PartyCast
- Shinies
- WoH-Reticle

## Examples

- Ace: LIBGUI_Checkbox:New -> ButtonSetCheckButtonFlag(w.name, true)
- Ace: LIBGUI_Optionbutton:New -> ButtonSetCheckButtonFlag(w.name, true)
- Killer: EnsureSettingsControlsInitialized -> ButtonSetCheckButtonFlag(SETTINGS_PERSONAL_CHECKBOX, true)
- Killer: Killer.local.EnsureSettingsControlsInitialized -> ButtonSetCheckButtonFlag(SETTINGS_PERSONAL_CHECKBOX, true)
- LibWBToggler: LIBGUI_Checkbox:New -> ButtonSetCheckButtonFlag(w.name, true)
- LibWBToggler: LIBGUI_Optionbutton:New -> ButtonSetCheckButtonFlag(w.name, true)

## Related APIs

- none

## Used With

- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
