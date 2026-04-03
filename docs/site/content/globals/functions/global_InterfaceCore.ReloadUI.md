# InterfaceCore.ReloadUI

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | Effigy, Enemy, SessionRPs, TastyButtons, Vertigo, nLootLink |
| Files seen in | Code/Assist/Assist.lua, Code/CombatLog/CombatLog.lua, Code/Core/Main.lua, Code/GroupIcons/GroupIcons.lua, Code/Guard/Guard.lua, Code/TalismanAlerter/TalismanAlerter.lua, Code/Timer/Timer.lua, Code/UnitFrames/UnitFrames.lua |
| Namespaces detected | InterfaceCore |
| Source kinds | lua_calls |
| Example locations | Effigy: SlashHandler, Enemy: AssistUI_ConfigDialog_OnReset, Enemy: CombatLogResetSettings, Enemy: CombatLogUI_ConfigDialog_OnReset, Enemy: GroupIconsUI_ConfigDialog_OnReset, Enemy: GuardUI_ConfigDialog_OnReset |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 19 |
| Global usage count | 19 |
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
InterfaceCore.ReloadUI()
```

## Description

Observed as a global function across 6 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Effigy
- Enemy
- SessionRPs
- TastyButtons
- Vertigo
- nLootLink

## Examples

- Effigy: SlashHandler -> InterfaceCore.ReloadUI()
- Enemy: AssistUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()
- Enemy: CombatLogResetSettings -> InterfaceCore.ReloadUI()
- Enemy: CombatLogUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()
- Enemy: GroupIconsUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()
- Enemy: GuardUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
