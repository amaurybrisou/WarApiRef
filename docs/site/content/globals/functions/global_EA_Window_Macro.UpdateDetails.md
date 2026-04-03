# EA_Window_Macro.UpdateDetails

- Category: Global Function
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
| Addons seen in | Brizio's Crappy Computer Medic, Enemy |
| Files seen in | CCM.lua, Code/Assist/Assist.lua, Code/Core/ConfigurationWindow.lua, Code/Core/Utils.lua, Code/ScenarioInfo/ScenarioInfo.lua |
| Namespaces detected | EA_Window_Macro |
| Source kinds | lua_calls |
| Example locations | Brizio's Crappy Computer Medic: SetMacro, Enemy: ConfigurationWindow_OnMacroMouseDrag, Enemy: OnMacroMouseDrag, Enemy: SetMacro |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
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
EA_Window_Macro.UpdateDetails(arg1)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: macro_id, p.macroId, slot |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Brizio's Crappy Computer Medic
- Enemy

## Examples

- Brizio's Crappy Computer Medic: SetMacro -> EA_Window_Macro.UpdateDetails(slot)
- Enemy: ConfigurationWindow_OnMacroMouseDrag -> EA_Window_Macro.UpdateDetails(p.macroId)
- Enemy: OnMacroMouseDrag -> EA_Window_Macro.UpdateDetails(macro_id)
- Enemy: SetMacro -> EA_Window_Macro.UpdateDetails(slot)

## Related APIs

- [IconLButtonUp](../../events/game_events/game_event_IconLButtonUp.md) (MEDIUM 43/100) - Game Event

## Used With

- [Cursor.IconOnCursor](global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](global_Cursor.PickUp.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
