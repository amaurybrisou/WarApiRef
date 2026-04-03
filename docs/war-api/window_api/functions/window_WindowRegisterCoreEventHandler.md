# WindowRegisterCoreEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 153

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, LibWBToggler, PartyCast, Shinies, TidyChat, TidyRoll, WoH-Reticle |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:285`, `/workspace/data/raw/LibWarBoardToggler/LibWBToggler.lua:32`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:285`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:285`, `/workspace/data/raw/Shinies/Libraries/LibGUI.lua:285`, `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136`, `/workspace/data/raw/WoH-Reticle/libs/LibGUI.lua:285` |
| Namespaces detected | WindowRegisterCoreEventHandler |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:RegisterEvent, LibWBToggler: LIBGUI_ELEMENT:RegisterEvent, LibWBToggler: LibWBToggler.RegisterEvent, PartyCast: LIBGUI_ELEMENT:RegisterEvent, Shinies: LIBGUI_ELEMENT:RegisterEvent, TidyChat: TidyChatFrames.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 15 |
| Global usage count | 15 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
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
WindowRegisterCoreEventHandler(windowName, windowEvent, handlerName)
```

## Description

Observed binding On* window events directly to a named window.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: c_TEXT_ENTRY_WINDOW.."ChannelButton", c_TEXT_ENTRY_WINDOW.."EntryBox", c_TEXT_ENTRY_WINDOW.."EntryBoxTextInput" |
| windowEvent | Observed as an On* window event string. | Observed values: "OnHidden", "OnLButtonDown", "OnLButtonUp" |
| handlerName | Observed as a Lua handler reference. | Observed values: "LIBGUI_ELEMENT.events."..e.."."..self.name, "TidyChat.OnChannelButtonMButtonUp", "TidyChat.OnEntryBoxTextInputLBD" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- Ace
- LibWBToggler
- PartyCast
- Shinies
- TidyChat
- TidyRoll
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:RegisterEvent -> WindowRegisterCoreEventHandler(self.name, e, "LIBGUI_ELEMENT.events."..e.."."..self.name)
- LibWBToggler: LIBGUI_ELEMENT:RegisterEvent -> WindowRegisterCoreEventHandler(self.name, e, "LIBGUI_ELEMENT.events."..e.."."..self.name)
- LibWBToggler: LibWBToggler.RegisterEvent -> WindowRegisterCoreEventHandler(modName, coreEvent, func)
- PartyCast: LIBGUI_ELEMENT:RegisterEvent -> WindowRegisterCoreEventHandler(self.name, e, "LIBGUI_ELEMENT.events."..e.."."..self.name)
- Shinies: LIBGUI_ELEMENT:RegisterEvent -> WindowRegisterCoreEventHandler(self.name, e, "LIBGUI_ELEMENT.events."..e.."."..self.name)
- TidyChat: TidyChatFrames.Initialize -> WindowRegisterCoreEventHandler(c_TEXT_ENTRY_WINDOW.."ChannelButton", "OnRButtonUp", "TidyChat.ToggleOptions")

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
