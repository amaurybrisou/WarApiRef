# ButtonSetDisabledFlag

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
| Files seen in | `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:746`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:828` |
| Namespaces detected | ButtonSetDisabledFlag |
| Source kinds | lua_calls |
| Example locations | TidyRoll: TidyRollOptions.OnCheckboxLBU, TidyRoll: TidyRollOptions.Reset |
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
ButtonSetDisabledFlag(arg1, arg2)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: c_TROLL_FLASH_ONLY_NEW_ITEM.."Button", c_TROLL_GLOW_ONLY_NEW_ITEM.."Button" |
| arg2 | Observed as a function or method reference. | Observed values: not ButtonGetPressedFlag(c_TROLL_FLASH_ON_NEW.."Button"), not ButtonGetPressedFlag(c_TROLL_GLOW_ON_NEW.."Button"), not ButtonGetPressedFlag(checkboxName) |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- TidyRoll

## Examples

- TidyRoll: TidyRollOptions.OnCheckboxLBU -> ButtonSetDisabledFlag(c_TROLL_FLASH_ONLY_NEW_ITEM.."Button", not ButtonGetPressedFlag(checkboxName))
- TidyRoll: TidyRollOptions.OnCheckboxLBU -> ButtonSetDisabledFlag(c_TROLL_GLOW_ONLY_NEW_ITEM.."Button", not ButtonGetPressedFlag(checkboxName))
- TidyRoll: TidyRollOptions.Reset -> ButtonSetDisabledFlag(c_TROLL_FLASH_ONLY_NEW_ITEM.."Button", not ButtonGetPressedFlag(c_TROLL_FLASH_ON_NEW.."Button"))
- TidyRoll: TidyRollOptions.Reset -> ButtonSetDisabledFlag(c_TROLL_GLOW_ONLY_NEW_ITEM.."Button", not ButtonGetPressedFlag(c_TROLL_GLOW_ON_NEW.."Button"))

## Related APIs

- none

## Used With

- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
