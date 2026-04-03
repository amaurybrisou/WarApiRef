# TextEditBoxGetText

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
| Addons seen in | TidyRoll |
| Files seen in | `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:434`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:782` |
| Namespaces detected | TextEditBoxGetText |
| Source kinds | lua_calls |
| Example locations | TidyRoll: TidyRoll.CustomAutoRoll.AddById, TidyRoll: TidyRollOptions.OnApply |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
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
TextEditBoxGetText(arg1)
```

## Description

Observed reading from or writing to edit-box controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: c_AUTO_ROLL_ADD_BY_ID_ID_EDITBOX, c_AUTO_ROLL_ADD_BY_ID_NAME_EDITBOX, c_TROLL_BNUM_TBOX |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- TidyRoll

## Examples

- TidyRoll: TidyRoll.CustomAutoRoll.AddById -> TextEditBoxGetText(c_AUTO_ROLL_ADD_BY_ID_ID_EDITBOX)
- TidyRoll: TidyRoll.CustomAutoRoll.AddById -> TextEditBoxGetText(c_AUTO_ROLL_ADD_BY_ID_NAME_EDITBOX)
- TidyRoll: TidyRollOptions.OnApply -> TextEditBoxGetText(c_TROLL_BNUM_TBOX)
- TidyRoll: TidyRollOptions.OnApply -> TextEditBoxGetText(c_TROLL_OFFSET_TBOX)

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBoxGetSelectedMenuItem](window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](window_ComboBoxGetSelectedText.md) (HIGH 80/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
