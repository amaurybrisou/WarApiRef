# InterfaceCore.ReloadUI

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 218

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent API_Ref source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy, Enemy |
| Files seen in | `/workspace/Effigy/EffigySlashCommands.lua:53`, `/workspace/Enemy/Code/Assist/Assist.lua:412`, `/workspace/Enemy/Code/CombatLog/CombatLog.lua:1497`, `/workspace/Enemy/Code/CombatLog/CombatLog.lua:1510`, `/workspace/Enemy/Code/Core/Main.lua:166`, `/workspace/Enemy/Code/GroupIcons/GroupIcons.lua:610`, `/workspace/Enemy/Code/Guard/Guard.lua:601`, `/workspace/Enemy/Code/TalismanAlerter/TalismanAlerter.lua:211` |
| Namespaces detected | InterfaceCore |
| Source kinds | bindings, globals, lua_calls, xml_handlers |
| Example locations | EA_UiDebugTools: DebugWindowReloadUi.OnLButtonUp, Effigy: Effigy.SlashHandler, Enemy: Enemy.AssistUI_ConfigDialog_OnReset, Enemy: Enemy.CombatLogResetSettings, Enemy: Enemy.CombatLogUI_ConfigDialog_OnReset, Enemy: Enemy.GroupIconsUI_ConfigDialog_OnReset |
| XML usage count | 2 |
| XML attribute usage count | 2 |
| Lua usage count | 11 |
| Global usage count | 11 |
| Local definition count | 0 |
| Documentation references | 3 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | yes |
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
InterfaceCore.ReloadUI()
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- Effigy
- Enemy

## Examples

- Effigy: Effigy.SlashHandler -> InterfaceCore.ReloadUI()
- Enemy: Enemy.AssistUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()
- Enemy: Enemy.CombatLogResetSettings -> InterfaceCore.ReloadUI()
- Enemy: Enemy.CombatLogUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()
- Enemy: Enemy.GroupIconsUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()
- Enemy: Enemy.GuardUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
