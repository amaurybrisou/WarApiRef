# LabelSetText

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
| Addons seen in | Moth, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:215`, `/workspace/data/raw/Moth/Moth.lua:227`, `/workspace/data/raw/TidyChat/TidyChat.lua:2117`, `/workspace/data/raw/TidyChat/TidyChat.lua:2201`, `/workspace/data/raw/TidyChat/TidyChat.lua:2321`, `/workspace/data/raw/TidyChat/TidyChat.lua:2333`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:145`, `/workspace/data/raw/TidyRoll/TidyRollFrame.lua:115` |
| Namespaces detected | LabelSetText |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.SetCellText, Moth: Moth.SetCellTextIcon, TidyChat: TidyChat.Copy.InitializeCopyWindow, TidyChat: TidyChat.LootRoll.InitializeLootRollWindow, TidyChat: TidyChat.LootRoll.OnRollLinkLButtonUp, TidyChat: TidyChat.Options.UpdateGroupTabs |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 46 |
| Global usage count | 46 |
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
LabelSetText(windowName, text)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: TCHAT_WINDOWS_TABS_INFO_LABEL, c_AUTO_ROLL_ADD_BY_ID_ID_LABEL, c_AUTO_ROLL_ADD_BY_ID_LABEL |
| text | Observed as a text or wstring payload. | Observed values: L " Icon               Name                    Id           Choice", L "", L "Add item to auto roll list by id" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Moth
- TidyChat
- TidyRoll

## Examples

- Moth: Moth.SetCellText -> LabelSetText(cellText, L "")
- Moth: Moth.SetCellText -> LabelSetText(cellText, text)
- Moth: Moth.SetCellTextIcon -> LabelSetText(cellText, L "")
- TidyChat: TidyChat.Copy.InitializeCopyWindow -> LabelSetText(c_TIDY_CHAT_COPY.."TitleBarText", L "Chat Copy")
- TidyChat: TidyChat.LootRoll.InitializeLootRollWindow -> LabelSetText(c_TIDY_CHAT_LOOT_ROLL.."TitleBarText", L "Roll History")
- TidyChat: TidyChat.LootRoll.OnRollLinkLButtonUp -> LabelSetText(c_TIDY_CHAT_LOOT_ROLL.."InfoLabel", TidyChatLootRoll.itemRollData.info)

## Related APIs

- none

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
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
