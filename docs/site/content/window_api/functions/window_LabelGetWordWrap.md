# LabelGetWordWrap

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
| Addons seen in | Ace, Effigy, GCDsaver, LibWBToggler, Shinies, WoH-Reticle |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:451`, `/workspace_addons/Effigy/LibGUI.lua:451`, `/workspace_addons/GCDsaver/libs/LibGUI.lua:451`, `/workspace_addons/LibWarBoardToggler/libs/LibGUI.lua:451`, `/workspace_addons/Shinies/Libraries/LibGUI.lua:451`, `/workspace_addons/WoH-Reticle/libs/LibGUI.lua:451` |
| Namespaces detected | LabelGetWordWrap |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Label:WordWrap, Effigy: LIBGUI_Label:WordWrap, GCDsaver: LIBGUI_Label:WordWrap, LibWBToggler: LIBGUI_Label:WordWrap, Shinies: LIBGUI_Label:WordWrap, WoH-Reticle: LIBGUI_Label:WordWrap |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
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
LabelGetWordWrap(arg1)
```

## Description

Observed as a window function across 6 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: self.name |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- Effigy
- GCDsaver
- LibWBToggler
- Shinies
- WoH-Reticle

## Examples

- Ace: LIBGUI_Label:WordWrap -> LabelGetWordWrap(self.name)
- Effigy: LIBGUI_Label:WordWrap -> LabelGetWordWrap(self.name)
- GCDsaver: LIBGUI_Label:WordWrap -> LabelGetWordWrap(self.name)
- LibWBToggler: LIBGUI_Label:WordWrap -> LabelGetWordWrap(self.name)
- Shinies: LIBGUI_Label:WordWrap -> LabelGetWordWrap(self.name)
- WoH-Reticle: LIBGUI_Label:WordWrap -> LabelGetWordWrap(self.name)

## Related APIs

- none

## Used With

- [LabelSetWordWrap](window_LabelSetWordWrap.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
