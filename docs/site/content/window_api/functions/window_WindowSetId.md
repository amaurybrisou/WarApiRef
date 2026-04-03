# WindowSetId

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 47 addons

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
| Addons seen in | Ace, AdvancedRenownTrainer, BankArkel, Calling, Crafting Info Tooltip, Crusher, DammazKron, DuffTimer |
| Files seen in | AdvancedRenownTraining.lua, BankArkel.lua, CallingSetup.lua, Code/ScenarioInfo/ScenarioInfo.lua, Core/Tome/DK_Tome.lua, CraftValueTip.lua, DuffTimer.lua, Gui/HealGridGuiColorSelect.lua |
| Namespaces detected | WindowSetId |
| Source kinds | lua_calls |
| Example locations | Ace: SetId, AdvancedRenownTrainer: CreateAbilityWindow, BankArkel: CreatePackWin, Calling: UpdateMacros, Crafting Info Tooltip: ItemWindow, Crusher: SetId |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 81 |
| Global usage count | 81 |
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
WindowSetId(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "CallingSetupMacroIconSlotC", "CallingSetupMacroIconSlotT", "DuffTimer_Bar_"..windowsKey.."_"..buffData.effectIndex |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 0, 1, 15 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedRenownTrainer
- BankArkel
- Calling
- Crafting Info Tooltip
- Crusher
- DammazKron
- DuffTimer
- EZCraftX
- EZGuard
- Effigy
- Enemy
- GCDsaver
- HealGrid
- Hopper
- InfoScroller
- Killer
- Kwestor
- LibWBToggler
- Map
- MapMonster
- MiniMapMonster
- Miracle Grow Remix
- Motion
- NerfedButtons
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- Refer
- ResHelp
- RoR_SoR
- Sequencer
- Shinies
- TacticSetNames
- TargetRing
- TidyChat
- TidyQueue
- TidyRoll
- Tokens
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- nRarity
- scenarioInfo
- zMailMod

## Examples

- Ace: SetId -> WindowSetId(self.name, id)
- AdvancedRenownTrainer: CreateAbilityWindow -> WindowSetId(t.windowName, t.id)
- BankArkel: CreatePackWin -> WindowSetId("PackIcon"..i, i)
- Calling: UpdateMacros -> WindowSetId("CallingSetupMacroIconSlotC", callMacroId)
- Calling: UpdateMacros -> WindowSetId("CallingSetupMacroIconSlotT", targetMacroId)
- Crafting Info Tooltip: ItemWindow -> WindowSetId(windowName, itemId)

## Related APIs

- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function

## Notes

- Advanced return analysis: No strong return evidence observed
