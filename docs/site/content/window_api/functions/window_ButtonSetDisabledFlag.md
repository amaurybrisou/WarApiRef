# ButtonSetDisabledFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 59 addons

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
| Addons seen in | Ace, ActionBarHide, AdvancedRenownTrainer, Amethyst, AutoBand, BuffHead, Busted, CDown |
| Files seen in | AdvancedRenownTraining.lua, AutoBand.lua, Busted.lua, CDownSettings.lua, ClosetGoblinCharacterWindow.lua, Code/Core/Main.lua, Code/Intercom/Intercom.lua, Code/UnitFrames/EffectsIndicator.lua |
| Namespaces detected | ButtonSetDisabledFlag |
| Source kinds | lua_calls |
| Example locations | Ace: SetEnabled, ActionBarHide: SetEnabled, AdvancedRenownTrainer: AnywhereShow, AdvancedRenownTrainer: OnHidden, Amethyst: SetEnabled, AutoBand: mark_template_settings_modified |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 332 |
| Global usage count | 332 |
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
ButtonSetDisabledFlag(arg1, arg2)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BustedGUINextError", "BustedGUIPrevError", "ClosetGoblinCharacterWindowContentsDeleteSet" |
| arg2 | Observed as a boolean toggle. | Observed values: (firstVisible==1), (onTargetChangeEnabled==false), (sequence.Count<firstVisible+#buttons-1) |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- AdvancedRenownTrainer
- Amethyst
- AutoBand
- BuffHead
- Busted
- CDown
- CM_ClosetGoblin
- CaVES
- CastSequence
- Crusher
- DAoCBuff
- DammazKron
- DetauntHelper
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- Enemy
- FastFriends
- FixGit
- GCDsaver
- HealGrid
- Hopper
- InfoScroller
- JunkDump
- Kwestor
- LibWBToggler
- Map
- MapMonster
- Miracle Grow Remix
- Motion
- NaturalLog
- NerfedButtons
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- RVMOD_Manager
- RVMOD_Targets
- RealmStatus
- Sequencer
- Shinies
- TargetRing
- TidyRoll
- Tokens
- Vectors
- WSCT
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- wbLeadHelper
- xHUD
- xPanels
- zMailMod

## Examples

- Ace: SetEnabled -> ButtonSetDisabledFlag(self.name, not flag)
- ActionBarHide: SetEnabled -> ButtonSetDisabledFlag(self.name, not flag)
- AdvancedRenownTrainer: AnywhereShow -> ButtonSetDisabledFlag(WindowName.."PurchaseButton", true)
- AdvancedRenownTrainer: OnHidden -> ButtonSetDisabledFlag(WindowName.."PurchaseButton", false)
- Amethyst: SetEnabled -> ButtonSetDisabledFlag(self.name, not flag)
- AutoBand: mark_template_settings_modified -> ButtonSetDisabledFlag(AutoBandWindowTemplate.savebutton_name, false)

## Related APIs

- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetDisabledFlag](window_ComboBoxSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Notes

- Advanced return analysis: No strong return evidence observed
