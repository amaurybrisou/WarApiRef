# WindowAssignFocus

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
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:226`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:231`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:226`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:231`, `/workspace/data/raw/TidyChat/TidyChat.lua:868` |
| Namespaces detected | WindowAssignFocus |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_ELEMENT:Defocus, InfoScroller: LIBGUI_ELEMENT:Focus, PartyCast: LIBGUI_ELEMENT:Defocus, PartyCast: LIBGUI_ELEMENT:Focus, TidyChat: TidyChatHooks.OnEnterChatTextHook |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
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
WindowAssignFocus(arg1, arg2)
```

## Description

Observed as a window function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: c_TEXT_ENTRY_WINDOW.."EntryBoxTextInput", self.name |
| arg2 | Observed as a boolean toggle. | Observed values: false, true |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- InfoScroller
- PartyCast
- TidyChat

## Examples

- InfoScroller: LIBGUI_ELEMENT:Defocus -> WindowAssignFocus(self.name, false)
- InfoScroller: LIBGUI_ELEMENT:Focus -> WindowAssignFocus(self.name, true)
- PartyCast: LIBGUI_ELEMENT:Defocus -> WindowAssignFocus(self.name, false)
- PartyCast: LIBGUI_ELEMENT:Focus -> WindowAssignFocus(self.name, true)
- TidyChat: TidyChatHooks.OnEnterChatTextHook -> WindowAssignFocus(c_TEXT_ENTRY_WINDOW.."EntryBoxTextInput", true)

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
