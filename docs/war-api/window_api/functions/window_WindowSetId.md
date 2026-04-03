# WindowSetId

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
| Addons seen in | Ace, AdvancedRenownTrainer, BankArkel, Enemy, Killer, LibWBToggler, PartyCast, RoR_SoR |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:262`, `/workspace/data/raw/BankArkel/BankArkel.lua:172`, `/workspace/data/raw/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:561`, `/workspace/data/raw/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:702`, `/workspace/data/raw/Killer/KillerUiCache.lua:102`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:262`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:262`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:1055` |
| Namespaces detected | WindowSetId |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:SetId, AdvancedRenownTrainer: AdvancedRenownTrainer.local.CreateAbilityWindow, AdvancedRenownTrainer: CreateAbilityWindow, BankArkel: BankArkel.CreatePackWin, Enemy: Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update, Enemy: Enemy.local.SetStatsRow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 29 |
| Global usage count | 29 |
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
| windowName | Observed as a target window name. | Observed values: "EnemyScenarioInfoDialog"..data.windowName.."NameLabel", "EnemyScenarioInfoDialogScoreDestrLabel", "EnemyScenarioInfoDialogScoreOrderLabel" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 1, 2, ChatSettings.Channels[SystemData.ChatLogFilters.ADVICE].id |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedRenownTrainer
- BankArkel
- Enemy
- Killer
- LibWBToggler
- PartyCast
- RoR_SoR
- Shinies
- TidyChat
- TidyRoll
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:SetId -> WindowSetId(self.name, id)
- AdvancedRenownTrainer: AdvancedRenownTrainer.local.CreateAbilityWindow -> WindowSetId(t.windowName, t.id)
- AdvancedRenownTrainer: CreateAbilityWindow -> WindowSetId(t.windowName, t.id)
- BankArkel: BankArkel.CreatePackWin -> WindowSetId("PackIcon"..i, i)
- Enemy: Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update -> WindowSetId("EnemyScenarioInfoDialogScoreDestrLabel", NewTooltip(L "Destruction points"))
- Enemy: Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update -> WindowSetId("EnemyScenarioInfoDialogScoreOrderLabel", NewTooltip(L "Order points"))

## Related APIs

- none

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
