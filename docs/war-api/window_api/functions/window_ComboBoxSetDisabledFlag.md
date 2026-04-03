# ComboBoxSetDisabledFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, PartyCast, TidyChat |
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1081`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1081`, `/workspace/data/raw/TidyChat/TidyChat.lua:1915` |
| Namespaces detected | ComboBoxSetDisabledFlag |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_Combobox:SetEnabled, PartyCast: LIBGUI_Combobox:SetEnabled, TidyChat: TidyChat.Options.UpdateDisabledFlags |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
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
ComboBoxSetDisabledFlag(arg1, arg2)
```

## Description

Observed as a window function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: TCHAT_TEXT_ENTRY_RELATIVE_TO_COMBO, self.name |
| arg2 | Observed as a runtime window or control identifier. | Observed values: not flag, pressed |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- PartyCast
- TidyChat

## Examples

- InfoScroller: LIBGUI_Combobox:SetEnabled -> ComboBoxSetDisabledFlag(self.name, not flag)
- PartyCast: LIBGUI_Combobox:SetEnabled -> ComboBoxSetDisabledFlag(self.name, not flag)
- TidyChat: TidyChat.Options.UpdateDisabledFlags -> ComboBoxSetDisabledFlag(TCHAT_TEXT_ENTRY_RELATIVE_TO_COMBO, pressed)

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
