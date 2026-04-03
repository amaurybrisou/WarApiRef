# TextEditBoxSetMaxChars

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
| Addons seen in | AutoBand, Enemy, NerfedButtons |
| Files seen in | AutoBand.lua, Code/Core/Main.lua, NerfedTalks.lua |
| Namespaces detected | TextEditBoxSetMaxChars |
| Source kinds | lua_calls |
| Example locations | AutoBand: ShowCopyLink, Enemy: UI_TextEntryDialog_Open, NerfedButtons: Initialize |
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
TextEditBoxSetMaxChars(arg1, arg2)
```

## Description

Observed reading from or writing to edit-box controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: 'EA_TextEntryGroupEntryBoxTextInput', COPY_LINK_WINDOW_NAME.."UrlInput", wn.."Value" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 1000, 1024*1024, chatInputMaxChars |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- AutoBand
- Enemy
- NerfedButtons

## Examples

- AutoBand: ShowCopyLink -> TextEditBoxSetMaxChars(COPY_LINK_WINDOW_NAME.."UrlInput", 1000)
- Enemy: UI_TextEntryDialog_Open -> TextEditBoxSetMaxChars(wn.."Value", 1024*1024)
- NerfedButtons: Initialize -> TextEditBoxSetMaxChars('EA_TextEntryGroupEntryBoxTextInput', chatInputMaxChars)

## Used With

- [TextEditBoxSelectAll](window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Notes

- none
