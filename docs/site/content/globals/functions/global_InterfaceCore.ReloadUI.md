# InterfaceCore.ReloadUI

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 115

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy |
| Files seen in | `/workspace/data/raw/Enemy/Code/Assist/Assist.lua:412`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLog.lua:1497`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLog.lua:1510`, `/workspace/data/raw/Enemy/Code/Core/Main.lua:166`, `/workspace/data/raw/Enemy/Code/GroupIcons/GroupIcons.lua:610`, `/workspace/data/raw/Enemy/Code/Guard/Guard.lua:601`, `/workspace/data/raw/Enemy/Code/TalismanAlerter/TalismanAlerter.lua:211`, `/workspace/data/raw/Enemy/Code/Timer/Timer.lua:305` |
| Namespaces detected | InterfaceCore |
| Source kinds | globals, lua_calls |
| Example locations | Enemy: Enemy.AssistUI_ConfigDialog_OnReset, Enemy: Enemy.CombatLogResetSettings, Enemy: Enemy.CombatLogUI_ConfigDialog_OnReset, Enemy: Enemy.GroupIconsUI_ConfigDialog_OnReset, Enemy: Enemy.GuardUI_ConfigDialog_OnReset, Enemy: Enemy.ResetSettings |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Enemy

## Examples

- Enemy: Enemy.AssistUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()
- Enemy: Enemy.CombatLogResetSettings -> InterfaceCore.ReloadUI()
- Enemy: Enemy.CombatLogUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()
- Enemy: Enemy.GroupIconsUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()
- Enemy: Enemy.GuardUI_ConfigDialog_OnReset -> InterfaceCore.ReloadUI()
- Enemy: Enemy.ResetSettings -> InterfaceCore.ReloadUI()

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
