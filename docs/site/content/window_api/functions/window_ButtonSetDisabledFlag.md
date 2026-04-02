# ButtonSetDisabledFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 21 addons

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
| Addons seen in | Ace, AdvancedRenownTrainer, BuffHead, Busted, CM_ClosetGoblin, DAoCBuff, EA_UiDebugTools, Effigy |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:553`, `/workspace_addons/Ace/LibGUI.lua:778`, `/workspace_addons/Ace/LibGUI.lua:844`, `/workspace_addons/BuffHead/Setup/SetupTrackers.lua:17`, `/workspace_addons/Busted/Busted.lua:130`, `/workspace_addons/ClosetGoblin/ClosetGoblinCharacterWindow.lua:465`, `/workspace_addons/ClosetGoblin/ClosetGoblinCharacterWindow.lua:671`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings2ndTier.lua:1056` |
| Namespaces detected | ButtonSetDisabledFlag |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:SetEnabled, Ace: LIBGUI_Checkbox:SetEnabled, Ace: LIBGUI_Optionbutton:SetEnabled, AdvancedRenownTrainer: AdvancedRenownTraining.AnywhereShow, AdvancedRenownTrainer: AdvancedRenownTraining.OnHidden, BuffHead: BuffHead.local.SetTracker |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 191 |
| Global usage count | 191 |
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
| arg2 | Observed as a boolean toggle. | Observed values: (onTargetChangeEnabled==false), config_dlg.clickCastingsListSelectedIndex==nil, config_dlg.clickCastingsListSelectedIndex==nil or config_dlg.clickCastingsListSelectedIndex==#config_dlg.clickCastings |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedRenownTrainer
- BuffHead
- Busted
- CM_ClosetGoblin
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- JunkDump
- LibWBToggler
- MapMonster
- Miracle Grow Remix
- PotionBar
- RVMOD_Manager
- Shinies
- TidyRoll
- WSCT
- WoH-Reticle
- wbLeadHelper

## Examples

- Ace: LIBGUI_Button:SetEnabled -> ButtonSetDisabledFlag(self.name, not flag)
- Ace: LIBGUI_Checkbox:SetEnabled -> ButtonSetDisabledFlag(self.name, not flag)
- Ace: LIBGUI_Optionbutton:SetEnabled -> ButtonSetDisabledFlag(self.name, not flag)
- AdvancedRenownTrainer: AdvancedRenownTraining.AnywhereShow -> ButtonSetDisabledFlag(WindowName.."PurchaseButton", true)
- AdvancedRenownTrainer: AdvancedRenownTraining.OnHidden -> ButtonSetDisabledFlag(WindowName.."PurchaseButton", false)
- BuffHead: BuffHead.local.SetTracker -> ButtonSetDisabledFlag(windowName.."OnTargetChangeClearAlwaysShow".."Button", (onTargetChangeEnabled==false))

## Related APIs

- none

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function

## Triggered By

- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 100/100) - XML Handler
- [OnHidden](../../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
