# CreateWindow

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 6 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, Moth, Soloq, TidyChat, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:30`, `/workspace/data/raw/Moth/Moth.lua:713`, `/workspace/data/raw/Soloq/ui/Overview.lua:24`, `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:145`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:227`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136`, `/workspace/data/raw/minesweep/minesweep.lua:11` |
| Namespaces detected | CreateWindow |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.OnInitialize, Moth: Moth.Initialize, Soloq: Soloq.createOverwiewWindow, TidyChat: TidyChatFrames.Initialize, TidyRoll: TidyRoll.CustomAutoRoll.Initialize, TidyRoll: TidyRoll.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
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
CreateWindow(windowName, showOnCreate)
```

## Description

Observed creating a top-level XML window from a loaded definition.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a top-level window name. | Observed values: "InfoScrollerMainWindow", "MineSweepWindow", "Moth" |
| showOnCreate | Observed as a boolean visibility flag. | Observed values: false, true |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Creates or instantiates UI objects from XML-backed definitions.

## Seen In

- InfoScroller
- Moth
- Soloq
- TidyChat
- TidyRoll
- minesweep

## Examples

- InfoScroller: InfoScroller.OnInitialize -> CreateWindow("InfoScrollerMainWindow", true)
- Moth: Moth.Initialize -> CreateWindow("Moth", true)
- Soloq: Soloq.createOverwiewWindow -> CreateWindow(overviewWindowName, false)
- TidyChat: TidyChatFrames.Initialize -> CreateWindow(c_TEXT_ENTRY_ANCHOR, false)
- TidyRoll: TidyRoll.CustomAutoRoll.Initialize -> CreateWindow(c_TROLL_AUTO_ROLL_WINDOW, false)
- TidyRoll: TidyRoll.Initialize -> CreateWindow(c_TIDY_ROLL_ANCHOR, false)

## Related APIs

- none

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [LabelSetFont](../../window_api/functions/window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
