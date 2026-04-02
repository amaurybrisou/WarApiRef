# WindowAssignFocus

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 12 addons

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
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:226`, `/workspace_addons/Ace/LibGUI.lua:231`, `/workspace_addons/Aura/Source/AuraShares.lua:379`, `/workspace_addons/Aura/Source/AuraShares.lua:397`, `/workspace_addons/Busted/Busted.lua:209`, `/workspace_addons/Effigy/LibGUI.lua:226`, `/workspace_addons/Effigy/LibGUI.lua:231`, `/workspace_addons/Enemy/Code/Core/Main.lua:842` |
| Namespaces detected | WindowAssignFocus |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Defocus, Ace: LIBGUI_ELEMENT:Focus, Aura: AuraShares.OnExportAura, Aura: AuraShares.OnImportAura, Busted: BustedGUI.UpdateErrorView, EA_UiDebugTools: BustedGUI.UpdateErrorView |
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
WindowAssignFocus(arg1, arg2)
```

## Description

Observed as a window function across 12 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BustedGUIErrorMessage", "DebugWindowTextBox", "EnemyChooseChannelDialog" |
| arg2 | Observed as a boolean toggle. | Observed values: false, true |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

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

- [TextEditBoxSelectAll](window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
