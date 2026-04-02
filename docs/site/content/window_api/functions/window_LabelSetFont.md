# LabelSetFont

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

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
| Addons seen in | Moth, TidyRoll |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:215`, `/workspace/data/raw/Moth/Moth.lua:227`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136` |
| Namespaces detected | LabelSetFont |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.SetCellText, Moth: Moth.SetCellTextIcon, TidyRoll: TidyRollOptions.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
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
LabelSetFont(arg1, arg2, arg3)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: c_TROLL_BINDS_LABEL, c_TROLL_ON_ROLL_ITEMS_NUMBER_CHANGE, c_TROLL_SORTING_TEXT |
| arg2 | Observed as a text or wstring payload. | Observed values: "font_clear_large_bold", "font_clear_medium", cellFont |
| arg3 | Observed as a numeric value. | Observed values: 20, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Moth
- TidyRoll

## Examples

- Moth: Moth.SetCellText -> LabelSetFont(cellText, cellFont, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING)
- Moth: Moth.SetCellTextIcon -> LabelSetFont(cellText, cellFont, WindowUtils.FONT_DEFAULT_TEXT_LINESPACING)
- TidyRoll: TidyRollOptions.Initialize -> LabelSetFont(c_TROLL_TITLE, "font_clear_large_bold", 20)
- TidyRoll: TidyRollOptions.Initialize -> LabelSetFont(c_TROLL_BINDS_LABEL, "font_clear_medium", 20)
- TidyRoll: TidyRollOptions.Initialize -> LabelSetFont(c_TROLL_ON_ROLL_ITEMS_NUMBER_CHANGE, "font_clear_medium", 20)
- TidyRoll: TidyRollOptions.Initialize -> LabelSetFont(c_TROLL_SORTING_TEXT, "font_clear_medium", 20)

## Related APIs

- none

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 98/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 90/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 90/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 90/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 90/100) - Window Function
- [LabelGetTextDimensions](window_LabelGetTextDimensions.md) (HIGH 80/100) - Window Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 71/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
