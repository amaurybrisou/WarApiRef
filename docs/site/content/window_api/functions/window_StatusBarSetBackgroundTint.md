# StatusBarSetBackgroundTint

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
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1028`, `/workspace/data/raw/PartyCast/PartyCast.lua:399`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1028`, `/workspace/data/raw/TidyRoll/TidyRollFrame.lua:105` |
| Namespaces detected | StatusBarSetBackgroundTint |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_Statusbar:BackColor, PartyCast: LIBGUI_Statusbar:BackColor, PartyCast: PartyCast.FetchedText, TidyRoll: TidyRollFrame:InitializeTimerBar |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
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
StatusBarSetBackgroundTint(arg1, arg2, arg3, arg4)
```

## Description

Observed as a window function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "PartyCastWindow"..PlayerNumber.."TimerBar", self.name, timerBarName |
| arg2 | Observed as a function or method reference. | Observed values: DefaultColor.BLACK.r, red |
| arg3 | Observed as a function or method reference. | Observed values: DefaultColor.BLACK.g, green |
| arg4 | Observed as a function or method reference. | Observed values: DefaultColor.BLACK.b, blue |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- PartyCast
- TidyRoll

## Examples

- InfoScroller: LIBGUI_Statusbar:BackColor -> StatusBarSetBackgroundTint(self.name, red, green, blue)
- PartyCast: LIBGUI_Statusbar:BackColor -> StatusBarSetBackgroundTint(self.name, red, green, blue)
- PartyCast: PartyCast.FetchedText -> StatusBarSetBackgroundTint("PartyCastWindow"..PlayerNumber.."TimerBar", DefaultColor.BLACK.r, DefaultColor.BLACK.g, DefaultColor.BLACK.b)
- TidyRoll: TidyRollFrame:InitializeTimerBar -> StatusBarSetBackgroundTint(timerBarName, DefaultColor.BLACK.r, DefaultColor.BLACK.g, DefaultColor.BLACK.b)

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
