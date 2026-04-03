# TextEditBoxSetText

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
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:655`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:668`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:710`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:723`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:655`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:668`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:710`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:723` |
| Namespaces detected | TextEditBoxSetText |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_MultiTextbox:Clear, InfoScroller: LIBGUI_MultiTextbox:SetText, InfoScroller: LIBGUI_Textbox:Clear, InfoScroller: LIBGUI_Textbox:SetText, PartyCast: LIBGUI_MultiTextbox:Clear, PartyCast: LIBGUI_MultiTextbox:SetText |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 15 |
| Global usage count | 15 |
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
TextEditBoxSetText(windowName, text)
```

## Description

Observed reading from or writing to edit-box controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: c_AUTO_ROLL_ADD_BY_ID_ID_EDITBOX, c_AUTO_ROLL_ADD_BY_ID_NAME_EDITBOX, c_TIDY_CHAT_COPY.."Log" |
| text | Observed as a text or wstring payload. | Observed values: L "", towstring(GetSetting("button-number")), towstring(GetSetting("button-offset")) |

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

- InfoScroller: LIBGUI_MultiTextbox:Clear -> TextEditBoxSetText(self.name, L "")
- InfoScroller: LIBGUI_MultiTextbox:SetText -> TextEditBoxSetText(self.name, towstring(text))
- InfoScroller: LIBGUI_Textbox:Clear -> TextEditBoxSetText(self.name, L "")
- InfoScroller: LIBGUI_Textbox:SetText -> TextEditBoxSetText(self.name, towstring(text))
- PartyCast: LIBGUI_MultiTextbox:Clear -> TextEditBoxSetText(self.name, L "")
- PartyCast: LIBGUI_MultiTextbox:SetText -> TextEditBoxSetText(self.name, towstring(text))

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
