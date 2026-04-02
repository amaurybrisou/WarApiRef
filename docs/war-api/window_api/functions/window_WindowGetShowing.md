# WindowGetShowing

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 38 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, AnywhereTrainerAdditions, Aura, AutoMark, BankWindowFix |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:84`, `/workspace_addons/AggroMeter/AggroMeter.lua:374`, `/workspace_addons/AnywhereTrainer/source/AnywhereTrainer.lua:227`, `/workspace_addons/AnywhereTrainer/source/AnywhereTrainer.lua:235`, `/workspace_addons/AnywhereTrainer/source/AnywhereTrainer.lua:243`, `/workspace_addons/AnywhereTrainer/source/AnywhereTrainer.lua:261`, `/workspace_addons/AnywhereTrainer/source/AnywhereTrainer.lua:271`, `/workspace_addons/AnywhereTrainer/source/AnywhereTrainer.lua:280` |
| Namespaces detected | WindowGetShowing |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Showing, AdvancedRenownTrainer: AdvancedRenownTraining.TogglePresets, AggroMeter: AggroMeter.Close, AnywhereTrainer: AnywhereTrainer.OnLeftClickAuction, AnywhereTrainer: AnywhereTrainer.OnLeftClickBank, AnywhereTrainer: AnywhereTrainer.OnLeftClickCareer |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 273 |
| Global usage count | 273 |
| Local definition count | 0 |
| Documentation references | 0 |
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
WindowGetShowing(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "AbilitiesWindow", "AdvancedRenownTrainingWindow", "AggroMeterGrayWindow" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- AnywhereTrainerAdditions
- Aura
- AutoMark
- BankWindowFix
- BuffHead
- Busted
- CombatTextNames
- DAoCBuff
- DaemonAssist
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- GuardLine
- LibGroup
- LibWBToggler
- MapMonster
- MapPin
- Miracle Grow Remix
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WarBoard
- WhoHealedMe
- WoH-Reticle
- nRarity
- wbLeadHelper

## Examples

- Ace: LIBGUI_ELEMENT:Showing -> WindowGetShowing(self.name)
- AdvancedRenownTrainer: AdvancedRenownTraining.TogglePresets -> WindowGetShowing(PresetWindowName)
- AggroMeter: AggroMeter.Close -> WindowGetShowing("AggroMeterGrayWindow")
- AnywhereTrainer: AnywhereTrainer.OnLeftClickAuction -> WindowGetShowing("ShiniesWindow")
- AnywhereTrainer: AnywhereTrainer.OnLeftClickAuction -> WindowGetShowing("AuctionWindow")
- AnywhereTrainer: AnywhereTrainer.OnLeftClickBank -> WindowGetShowing("BankWindow")

## Related APIs

- none

## Used With

- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetLayer](window_WindowGetLayer.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [BankWindow.Hide](../../globals/functions/global_BankWindow.Hide.md) (HIGH 88/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
