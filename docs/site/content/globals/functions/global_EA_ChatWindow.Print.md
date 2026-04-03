# EA_ChatWindow.Print

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 18 addons

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
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, Aura, BagOMatic, BankArkel, DAoCBuff, Enemy, GuardLine, Killer |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APACore.lua:298`, `/workspace/data/raw/Aura/Source/AuraAddon.lua:469`, `/workspace/data/raw/BankArkel/BankArkel.lua:686`, `/workspace/data/raw/BankArkel/BankArkel.lua:95`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:687`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:772`, `/workspace/data/raw/Enemy/Code/Core/Utils.lua:197`, `/workspace/data/raw/GuardLine/GuardLine.lua:730` |
| Namespaces detected | EA_ChatWindow |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.Print, Aura: AuraPrint, BagOMatic: BagOMatic.print, BankArkel: BankArkel.Init, BankArkel: BankArkel.PrintFactionGold, DAoCBuff: DAoCBuff.ResetDAoCBuff |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 135 |
| Global usage count | 135 |
| Local definition count | 0 |
| Documentation references | 1 |
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
EA_ChatWindow.Print(arg1)
```

## Description

Observed as a global function across 18 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: C.TAG..towstring(text), L "<icon149> "..msg, L "<icon57> "..SwiftLineSC..SwiftVersion..L "has been initialized. Use /swift for more details" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Writes output to the chat or UI surface.

## Seen In

- AdvancedPetAssist
- Aura
- BagOMatic
- BankArkel
- DAoCBuff
- Enemy
- GuardLine
- Killer
- LibGuard
- LibSlash
- LibWBToggler
- MiracleGrowLight
- PlanB
- RoR_SoR
- Shinies
- Swift Assist
- WSCT
- WhoHealedMe

## Examples

- AdvancedPetAssist: AdvancedPetAssist.Print -> EA_ChatWindow.Print(L "<icon149> "..msg)
- Aura: AuraPrint -> EA_ChatWindow.Print(towstring(str))
- BagOMatic: BagOMatic.print -> EA_ChatWindow.Print(L "[BagOMatic] "..towstring(text))
- BankArkel: BankArkel.Init -> EA_ChatWindow.Print(StringToWString(InitTxt))
- BankArkel: BankArkel.PrintFactionGold -> EA_ChatWindow.Print(STWS(header))
- BankArkel: BankArkel.PrintFactionGold -> EA_ChatWindow.Print(STWS(line))

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [BankWindow](../tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [EA_ChatWindow.OnKeyEnter](global_EA_ChatWindow.OnKeyEnter.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field
- [BankWindow.Hide](global_BankWindow.Hide.md) (HIGH 88/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
