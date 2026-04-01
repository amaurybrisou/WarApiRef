# WindowAssignFocus

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 13 addons

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
| Addons seen in | Ace, Aura, Busted, EA_UiDebugTools, Effigy, Enemy, GCDsaver, LibWBToggler |
| Files seen in | `/workspace/Ace/LibGUI.lua:226`, `/workspace/Ace/LibGUI.lua:231`, `/workspace/Aura/Source/AuraShares.lua:379`, `/workspace/Aura/Source/AuraShares.lua:397`, `/workspace/Busted/Busted.lua:209`, `/workspace/Effigy/LibGUI.lua:226`, `/workspace/Effigy/LibGUI.lua:231`, `/workspace/Enemy/Code/Core/Main.lua:842` |
| Namespaces detected | WindowAssignFocus |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Defocus, Ace: LIBGUI_ELEMENT:Focus, Aura: AuraShares.OnExportAura, Aura: AuraShares.OnImportAura, Busted: BustedGUI.UpdateErrorView, EA_UiDebugTools: BustedGUI.UpdateErrorView |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 48 |
| Global usage count | 48 |
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
WindowAssignFocus(arg1, arg2)
```

## Description

Observed as a window function across 13 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BustedGUIErrorMessage", "DebugWindowTextBox", "EnemyChooseChannelDialog" |
| arg2 | Observed as a boolean toggle. | Observed values: false, true |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- Ace
- Aura
- Busted
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- LibWBToggler
- MapMonster
- Shinies
- TidyChat
- WarTriage
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:Defocus -> WindowAssignFocus(self.name, false)
- Ace: LIBGUI_ELEMENT:Focus -> WindowAssignFocus(self.name, true)
- Aura: AuraShares.OnExportAura -> WindowAssignFocus(exportWindow.."AuraText", true)
- Aura: AuraShares.OnImportAura -> WindowAssignFocus(importWindow.."AuraText", true)
- Busted: BustedGUI.UpdateErrorView -> WindowAssignFocus("BustedGUIErrorMessage", true)
- EA_UiDebugTools: BustedGUI.UpdateErrorView -> WindowAssignFocus("BustedGUIErrorMessage", true)

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [TextEditBoxSelectAll](window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [TextEditBoxSetMaxChars](window_TextEditBoxSetMaxChars.md) (HIGH 80/100) - Window Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
