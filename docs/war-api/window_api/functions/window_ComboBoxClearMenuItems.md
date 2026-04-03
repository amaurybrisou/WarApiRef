# ComboBoxClearMenuItems

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
| Addons seen in | InfoScroller, PartyCast, TidyRoll |
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1075`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1075`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136` |
| Namespaces detected | ComboBoxClearMenuItems |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_Combobox:Clear, PartyCast: LIBGUI_Combobox:Clear, TidyRoll: TidyRollOptions.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
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
ComboBoxClearMenuItems(arg1)
```

## Description

Observed as a window function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: c_TROLL_DIRECTION_COMBO, c_TROLL_GREED_COMBO, c_TROLL_NEED_COMBO |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- InfoScroller
- PartyCast
- TidyRoll

## Examples

- InfoScroller: LIBGUI_Combobox:Clear -> ComboBoxClearMenuItems(self.name)
- PartyCast: LIBGUI_Combobox:Clear -> ComboBoxClearMenuItems(self.name)
- TidyRoll: TidyRollOptions.Initialize -> ComboBoxClearMenuItems(c_TROLL_DIRECTION_COMBO)
- TidyRoll: TidyRollOptions.Initialize -> ComboBoxClearMenuItems(c_TROLL_ONESC_COMBO)
- TidyRoll: TidyRollOptions.Initialize -> ComboBoxClearMenuItems(c_TROLL_NEED_COMBO)
- TidyRoll: TidyRollOptions.Initialize -> ComboBoxClearMenuItems(c_TROLL_GREED_COMBO)

## Related APIs

- none

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
