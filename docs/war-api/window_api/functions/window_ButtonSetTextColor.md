# ButtonSetTextColor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 90/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:1013` |
| Namespaces detected | ButtonSetTextColor |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChatFrames.UpdateTidyChannelButtonsColors |
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
| Consistent role | no |
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
ButtonSetTextColor(arg1, arg2, arg3, arg4, arg5)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: channelButton |
| arg2 | Observed as a function or method reference. | Observed values: Button.ButtonState.HIGHLIGHTED, Button.ButtonState.NORMAL, Button.ButtonState.PRESSED |
| arg3 | Observed as a function or method reference. | Observed values: channelColor.r |
| arg4 | Observed as a function or method reference. | Observed values: channelColor.g |
| arg5 | Observed as a function or method reference. | Observed values: channelColor.b |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- TidyChat

## Examples

- TidyChat: TidyChatFrames.UpdateTidyChannelButtonsColors -> ButtonSetTextColor(channelButton, Button.ButtonState.NORMAL, channelColor.r, channelColor.g, channelColor.b)
- TidyChat: TidyChatFrames.UpdateTidyChannelButtonsColors -> ButtonSetTextColor(channelButton, Button.ButtonState.HIGHLIGHTED, channelColor.r, channelColor.g, channelColor.b)
- TidyChat: TidyChatFrames.UpdateTidyChannelButtonsColors -> ButtonSetTextColor(channelButton, Button.ButtonState.PRESSED, channelColor.r, channelColor.g, channelColor.b)
- TidyChat: TidyChatFrames.UpdateTidyChannelButtonsColors -> ButtonSetTextColor(channelButton, Button.ButtonState.PRESSED_HIGHLIGHTED, channelColor.r, channelColor.g, channelColor.b)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
