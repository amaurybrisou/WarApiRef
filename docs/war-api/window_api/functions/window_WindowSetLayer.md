# WindowSetLayer

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
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:178`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:178`, `/workspace/data/raw/TidyChat/TidyChat.lua:239`, `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136` |
| Namespaces detected | WindowSetLayer |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_ELEMENT:Layer, PartyCast: LIBGUI_ELEMENT:Layer, TidyChat: TidyChatCore.SetWindowGroup, TidyChat: TidyChatFrames.Initialize, TidyRoll: TidyRollOptions.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 14 |
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
WindowSetLayer(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: c_TEXT_ENTRY_WINDOW, c_TEXT_ENTRY_WINDOW.."ChannelButton", c_TROLL_BNUM_TBOX |
| arg2 | Observed as a function or method reference. | Observed values: 0, 2, 3 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- PartyCast
- TidyChat
- TidyRoll

## Examples

- InfoScroller: LIBGUI_ELEMENT:Layer -> WindowSetLayer(self.name, layer)
- PartyCast: LIBGUI_ELEMENT:Layer -> WindowSetLayer(self.name, layer)
- TidyChat: TidyChatCore.SetWindowGroup -> WindowSetLayer(wndGroupName.."Background", 0)
- TidyChat: TidyChatFrames.Initialize -> WindowSetLayer(c_TEXT_ENTRY_WINDOW.."ChannelButton", 3)
- TidyChat: TidyChatFrames.Initialize -> WindowSetLayer(c_TEXT_ENTRY_WINDOW, 2)
- TidyRoll: TidyRollOptions.Initialize -> WindowSetLayer(c_TROLL_BNUM_TBOX, Window.Layers.POPUP)

## Related APIs

- none

## Used With

- [WindowGetLayer](window_WindowGetLayer.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
