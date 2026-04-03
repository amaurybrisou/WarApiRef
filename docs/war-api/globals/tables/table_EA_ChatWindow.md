# EA_ChatWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 130

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AnywhereTrainerAdditions, Aura, BagOMatic, BankArkel, DAoCBuff, Enemy |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APACore.lua:298`, `/workspace/data/raw/AnywhereTrainerAdditions/AnywhereTrainerAdditions.lua:218`, `/workspace/data/raw/Aura/Source/AuraAddon.lua:469`, `/workspace/data/raw/BankArkel/BankArkel.lua:686`, `/workspace/data/raw/BankArkel/BankArkel.lua:95`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:687`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:772`, `/workspace/data/raw/Enemy/Code/Core/Utils.lua:197` |
| Namespaces detected | EA_ChatWindow |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.Print, AdvancedRenownTrainer: AdvancedRenownTraining.ExportToLink, AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown, Aura: AuraPrint, BagOMatic: BagOMatic.print, BankArkel: BankArkel.Init |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 142 |
| Global usage count | 8 |
| Local definition count | 6 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed shared global table or namespace surfaced in 21 addons.

## Functions

- EA_ChatWindow.GetCurrentChannel
- EA_ChatWindow.InsertItemLink
- EA_ChatWindow.InsertText
- EA_ChatWindow.OnKeyEnter
- EA_ChatWindow.OnSettingsChanged
- EA_ChatWindow.Print
- EA_ChatWindow.ShowTabCycleButtons
- EA_ChatWindow.UpdateTabScrollWindow

## Observed Members

- none

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AnywhereTrainerAdditions
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
- TidyChat
- WSCT
- WhoHealedMe

## Examples

- AdvancedPetAssist: AdvancedPetAssist.Print -> EA_ChatWindow.Print(L "<icon149> "..msg)
- AdvancedRenownTrainer: AdvancedRenownTraining.ExportToLink -> EA_ChatWindow.InsertText(hl)
- AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown -> EA_ChatWindow.InsertItemLink(itemData)
- Aura: AuraPrint -> EA_ChatWindow.Print(towstring(str))
- BagOMatic: BagOMatic.print -> EA_ChatWindow.Print(L "[BagOMatic] "..towstring(text))
- BankArkel: BankArkel.Init -> EA_ChatWindow.Print(StringToWString(InitTxt))

## Related APIs

- [CreateWindow](../functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [CreateWindow](../functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none
