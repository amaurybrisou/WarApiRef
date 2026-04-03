# ButtonSetTextColor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, LibWBToggler, PartyCast, Shinies, TidyChat, WoH-Reticle |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:543`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:543`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:543`, `/workspace/data/raw/Shinies/Libraries/LibGUI.lua:543`, `/workspace/data/raw/TidyChat/TidyChat.lua:1013`, `/workspace/data/raw/WoH-Reticle/libs/LibGUI.lua:543` |
| Namespaces detected | ButtonSetTextColor |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:TextColor, LibWBToggler: LIBGUI_Button:TextColor, PartyCast: LIBGUI_Button:TextColor, Shinies: LIBGUI_Button:TextColor, TidyChat: TidyChatFrames.UpdateTidyChannelButtonsColors, WoH-Reticle: LIBGUI_Button:TextColor |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 17 |
| Global usage count | 17 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
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
| arg1 | Observed as a runtime window or control identifier. | Observed values: channelButton, self.name |
| arg2 | Observed as a function or method reference. | Observed values: Button.ButtonState.HIGHLIGHTED, Button.ButtonState.NORMAL, Button.ButtonState.PRESSED |
| arg3 | Observed as a function or method reference. | Observed values: channelColor.r, green |
| arg4 | Observed as a function or method reference. | Observed values: blue, channelColor.g |
| arg5 | Observed as a function or method reference. | Observed values: channelColor.b |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- LibWBToggler
- PartyCast
- Shinies
- TidyChat
- WoH-Reticle

## Examples

- Ace: LIBGUI_Button:TextColor -> ButtonSetTextColor(self.name, red, green, blue)
- LibWBToggler: LIBGUI_Button:TextColor -> ButtonSetTextColor(self.name, red, green, blue)
- PartyCast: LIBGUI_Button:TextColor -> ButtonSetTextColor(self.name, red, green, blue)
- Shinies: LIBGUI_Button:TextColor -> ButtonSetTextColor(self.name, red, green, blue)
- TidyChat: TidyChatFrames.UpdateTidyChannelButtonsColors -> ButtonSetTextColor(channelButton, Button.ButtonState.NORMAL, channelColor.r, channelColor.g, channelColor.b)
- TidyChat: TidyChatFrames.UpdateTidyChannelButtonsColors -> ButtonSetTextColor(channelButton, Button.ButtonState.HIGHLIGHTED, channelColor.r, channelColor.g, channelColor.b)

## Related APIs

- none

## Used With

- [ButtonGetTextColor](window_ButtonGetTextColor.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
