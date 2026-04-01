# WindowGetParent

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 10 addons

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
| Addons seen in | AggroMeter, Aura, Enemy, LoyalPet, MapMonster, Miracle Grow Remix, RVMOD_Manager, RandomMount |
| Files seen in | `/workspace/AggroMeter/AggroMeter.lua:205`, `/workspace/AggroMeter/AggroMeter.lua:242`, `/workspace/Aura/Source/AuraTexture.lua:147`, `/workspace/Enemy/Code/Marks/MarkTemplate.lua:264`, `/workspace/Enemy/Code/Marks/MarkTemplate.lua:281`, `/workspace/LoyalPet/gui/lpet_gui.lua:189`, `/workspace/MGRemix/MGRemix.lua:630`, `/workspace/MGRemix/MGRemix.lua:642` |
| Namespaces detected | WindowGetParent |
| Source kinds | lua_calls |
| Example locations | AggroMeter: AggroMeter.OnMouseOverStart, AggroMeter: AggroMeter.SelectChar, Aura: AuraTexture.OnIconLButtonUp, Enemy: Enemy.MarkUI_EnemyMark_OnLButtonDown, Enemy: Enemy.MarkUI_EnemyMark_OnRButtonUp, LoyalPet: LPET.OnMouseOver |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 42 |
| Global usage count | 42 |
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
WindowGetParent(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: SystemData.ActiveWindow.name, SystemData.MouseOverWindow.name, WindowGetParent(SystemData.ActiveWindow.name) |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AggroMeter
- Aura
- Enemy
- LoyalPet
- MapMonster
- Miracle Grow Remix
- RVMOD_Manager
- RandomMount
- RoR_SoR
- WSCT

## Examples

- AggroMeter: AggroMeter.OnMouseOverStart -> WindowGetParent(SystemData.MouseOverWindow.name)
- AggroMeter: AggroMeter.SelectChar -> WindowGetParent(SystemData.MouseOverWindow.name)
- Aura: AuraTexture.OnIconLButtonUp -> WindowGetParent(WindowGetParent(WindowGetParent(SystemData.ActiveWindow.name)))
- Aura: AuraTexture.OnIconLButtonUp -> WindowGetParent(WindowGetParent(SystemData.ActiveWindow.name))
- Aura: AuraTexture.OnIconLButtonUp -> WindowGetParent(SystemData.ActiveWindow.name)
- Enemy: Enemy.MarkUI_EnemyMark_OnLButtonDown -> WindowGetParent(SystemData.MouseOverWindow.name)

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetDisabledFlag](window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EditBox](../../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnTextChanged](../../events/window_events/window_event_OnTextChanged.md) (HIGH 100/100) - Window Event
- [OnTextChanged](../../xml/handlers/handler_OnTextChanged.md) (HIGH 100/100) - XML Handler
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
