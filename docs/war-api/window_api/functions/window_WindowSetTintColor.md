# WindowSetTintColor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 98/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Moth, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:270`, `/workspace/data/raw/Moth/Moth.lua:591`, `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TidyRoll/TidyRollFrame.lua:134`, `/workspace/data/raw/TidyRoll/TidyRollFrame.lua:196` |
| Namespaces detected | WindowSetTintColor |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.UpdateHealthBar, Moth: Moth.UpdateLevel, TidyChat: TidyChatFrames.Initialize, TidyRoll: TidyRollFrame:SetIcon, TidyRoll: TidyRollFrame:SetLootData |
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
WindowSetTintColor(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "MothHealthBar", buttonName, c_TEXT_ENTRY_WINDOW.."EntryBoxBG" |
| arg2 | Observed as a function or method reference. | Observed values: 0, Moth.Helpers.DefaultColorConvert(DefaultColor.GREEN), Moth.Helpers.DefaultColorConvert(DefaultColor.ORANGE) |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Moth
- TidyChat
- TidyRoll

## Examples

- Moth: Moth.UpdateHealthBar -> WindowSetTintColor("MothHealthBar", Moth.Helpers.DefaultColorConvert(DefaultColor.GREEN))
- Moth: Moth.UpdateHealthBar -> WindowSetTintColor("MothHealthBar", Moth.Helpers.DefaultColorConvert(DefaultColor.YELLOW))
- Moth: Moth.UpdateHealthBar -> WindowSetTintColor("MothHealthBar", Moth.Helpers.DefaultColorConvert(DefaultColor.ORANGE))
- Moth: Moth.UpdateHealthBar -> WindowSetTintColor("MothHealthBar", Moth.Helpers.DefaultColorConvert(DefaultColor.RED))
- Moth: Moth.UpdateLevel -> WindowSetTintColor(cellName.."RankBackgroundTint", Moth.Helpers.DefaultColorConvert(unitConColor))
- TidyChat: TidyChatFrames.Initialize -> WindowSetTintColor(c_TEXT_ENTRY_WINDOW.."EntryBoxBG", 0, 0, 0)

## Related APIs

- none

## Used With

- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 98/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
