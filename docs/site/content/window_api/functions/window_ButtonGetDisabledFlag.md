# ButtonGetDisabledFlag

- Category: Window Function
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
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, AdvancedRenownTrainer, Aura, BuffHead, DAoCBuff, Effigy, Enemy, GCDsaver |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:558`, `/workspace_addons/Ace/LibGUI.lua:783`, `/workspace_addons/Ace/LibGUI.lua:849`, `/workspace_addons/Aura/Source/AuraTexture.lua:147`, `/workspace_addons/BuffHead/Setup/SetupTrackers.lua:239`, `/workspace_addons/BuffHead/Setup/SetupTrackers.lua:253`, `/workspace_addons/BuffHead/Setup/SetupTrackers.lua:267`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings.lua:1350` |
| Namespaces detected | ButtonGetDisabledFlag |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:Enabled, Ace: LIBGUI_Checkbox:Enabled, Ace: LIBGUI_Optionbutton:Enabled, AdvancedRenownTrainer: AdvancedRenownTraining.PurchaseAdvances, Aura: AuraTexture.OnIconLButtonUp, BuffHead: BuffHead.Setup.Trackers.OnTargetChangeClearAlwaysShowLUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 76 |
| Global usage count | 76 |
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

Observed as a window function across 18 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "EnemyConfigDialogResetAllButton", "EnemyConfigDialogResetButton", "EnemyEffectsIndicatorDialogContentScrollChildChooseIconButton" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- AdvancedRenownTrainer
- Aura
- BuffHead
- DAoCBuff
- Effigy
- Enemy
- GCDsaver
- JunkDump
- LibWBToggler
- PotionBar
- RVMOD_Manager
- Shinies
- TidyChat
- TidyRoll
- WoH-Reticle
- nRarity
- wbLeadHelper

## Examples

- Ace: LIBGUI_Button:Enabled -> ButtonGetDisabledFlag(self.name)
- Ace: LIBGUI_Checkbox:Enabled -> ButtonGetDisabledFlag(self.name)
- Ace: LIBGUI_Optionbutton:Enabled -> ButtonGetDisabledFlag(self.name)
- AdvancedRenownTrainer: AdvancedRenownTraining.PurchaseAdvances -> ButtonGetDisabledFlag(WindowName.."PurchaseButton")
- Aura: AuraTexture.OnIconLButtonUp -> ButtonGetDisabledFlag(SystemData.ActiveWindow.name)
- BuffHead: BuffHead.Setup.Trackers.OnTargetChangeClearAlwaysShowLUp -> ButtonGetDisabledFlag(windowName.."OnTargetChangeClearAlwaysShow".."Button")

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [WindowGetParent](window_WindowGetParent.md) (HIGH 100/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
