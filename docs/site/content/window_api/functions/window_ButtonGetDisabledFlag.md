# ButtonGetDisabledFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 56 addons

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
| Addons seen in | Ace, ActionBarHide, AdvancedRenownTrainer, Amethyst, Aura, BuffHead, CDown, Crusher |
| Files seen in | AdvancedRenownTraining.lua, CDownSettings.lua, Code/Core/ConfigurationWindow.lua, Code/Core/Main.lua, Code/Intercom/Intercom.lua, Code/UnitFrames/EffectsIndicator.lua, Code/UnitFrames/UnitFramePart.lua, Code/UnitFrames/UnitFrames.lua |
| Namespaces detected | ButtonGetDisabledFlag |
| Source kinds | lua_calls |
| Example locations | Ace: Enabled, ActionBarHide: Enabled, AdvancedRenownTrainer: PurchaseAdvances, Amethyst: Enabled, Aura: OnIconLButtonUp, BuffHead: OnTargetChangeClearAlwaysShowLUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 168 |
| Global usage count | 168 |
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
ButtonGetDisabledFlag(arg1)
```

## Description

Observed as a window function across 56 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "EA_Window_ScenarioJoinPromptBoxJoinWaitButton", "EnemyConfigDialogResetAllButton", "EnemyConfigDialogResetButton" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionBarHide
- AdvancedRenownTrainer
- Amethyst
- Aura
- BuffHead
- CDown
- Crusher
- DAoCBuff
- DammazKron
- DetauntHelper
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- Enemy
- EveryBodyGuard
- FastFriends
- GCDsaver
- Hopper
- InfoScroller
- JunkDump
- LibWBToggler
- Map
- Minmap
- Motion
- NaturalLog
- NerfedButtons
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- Queue Queuer
- RVMOD_Manager
- RealmStatus
- Refer
- Sequencer
- Shinies
- Squared
- TargetRing
- TidyChat
- TidyQueue
- TidyRoll
- Tokens
- Vectors
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- nRarity
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: Enabled -> ButtonGetDisabledFlag(self.name)
- ActionBarHide: Enabled -> ButtonGetDisabledFlag(self.name)
- AdvancedRenownTrainer: PurchaseAdvances -> ButtonGetDisabledFlag(WindowName.."PurchaseButton")
- Amethyst: Enabled -> ButtonGetDisabledFlag(self.name)
- Aura: OnIconLButtonUp -> ButtonGetDisabledFlag(SystemData.ActiveWindow.name)
- BuffHead: OnTargetChangeClearAlwaysShowLUp -> ButtonGetDisabledFlag(windowName.."OnTargetChangeClearAlwaysShow".."Button")

## Related APIs

- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [ComboBoxGetDisabledFlag](window_ComboBoxGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [WindowGetParent](window_WindowGetParent.md) (HIGH 100/100) - Window Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- none
