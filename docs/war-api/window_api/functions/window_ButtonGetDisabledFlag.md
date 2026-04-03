# ButtonGetDisabledFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | InfoScroller, PartyCast, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:558`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:782`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:847`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:558`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:782`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:847`, `/workspace/data/raw/TidyChat/TidyChat.lua:2031`, `/workspace/data/raw/TidyChat/TidyChat.lua:2275` |
| Namespaces detected | ButtonGetDisabledFlag |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_Button:Enabled, InfoScroller: LIBGUI_Checkbox:Enabled, InfoScroller: LIBGUI_Optionbutton:Enabled, PartyCast: LIBGUI_Button:Enabled, PartyCast: LIBGUI_Checkbox:Enabled, PartyCast: LIBGUI_Optionbutton:Enabled |
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
ButtonGetDisabledFlag(arg1)
```

## Description

Observed as a window function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: buttonName, c_TIDY_CHAT_COPY.."Next", c_TIDY_CHAT_COPY.."Prev" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- InfoScroller
- PartyCast
- TidyChat
- TidyRoll

## Examples

- InfoScroller: LIBGUI_Button:Enabled -> ButtonGetDisabledFlag(self.name)
- InfoScroller: LIBGUI_Checkbox:Enabled -> ButtonGetDisabledFlag(self.name)
- InfoScroller: LIBGUI_Optionbutton:Enabled -> ButtonGetDisabledFlag(self.name)
- PartyCast: LIBGUI_Button:Enabled -> ButtonGetDisabledFlag(self.name)
- PartyCast: LIBGUI_Checkbox:Enabled -> ButtonGetDisabledFlag(self.name)
- PartyCast: LIBGUI_Optionbutton:Enabled -> ButtonGetDisabledFlag(self.name)

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
