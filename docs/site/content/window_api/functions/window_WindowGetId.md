# WindowGetId

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
| Addons seen in | Ace, AdvancedRenownTrainer, AggroMeter, Aura, BagOMatic, BankArkel, BuffHead, CM_ClosetGoblin |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:269`, `/workspace/data/raw/AggroMeter/AggroMeter.lua:378`, `/workspace/data/raw/AggroMeter/AggroMeter.lua:426`, `/workspace/data/raw/Aura/Source/AuraShares.lua:257`, `/workspace/data/raw/Aura/Source/AuraShares.lua:297`, `/workspace/data/raw/Aura/Source/AuraTexture.lua:195`, `/workspace/data/raw/Aura/Source/AuraTexture.lua:65`, `/workspace/data/raw/BankArkel/BankArkel.lua:153` |
| Namespaces detected | WindowGetId |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:GetId, AdvancedRenownTrainer: AdvancedRenownTraining.AbilityTooltip, AdvancedRenownTrainer: AdvancedRenownTraining.OnLButtonUpTab, AdvancedRenownTrainer: AdvancedRenownTraining.Select, AggroMeter: AggroMeter.OnTabLBU, AggroMeter: AggroMeter.PickedListMenu |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 100 |
| Global usage count | 100 |
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
WindowGetId(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: EA_Window_ContextMenu.activeWindow, SystemData.ActiveWindow.name, SystemData.MouseOverWindow.name |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- AdvancedRenownTrainer
- AggroMeter
- Aura
- BagOMatic
- BankArkel
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Killer
- LibWBToggler
- PartyCast
- Pocket Palette
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:GetId -> WindowGetId(self.name)
- AdvancedRenownTrainer: AdvancedRenownTraining.AbilityTooltip -> WindowGetId(SystemData.MouseOverWindow.name)
- AdvancedRenownTrainer: AdvancedRenownTraining.OnLButtonUpTab -> WindowGetId(SystemData.ActiveWindow.name)
- AdvancedRenownTrainer: AdvancedRenownTraining.Select -> WindowGetId(activeWindow)
- AggroMeter: AggroMeter.OnTabLBU -> WindowGetId(SystemData.ActiveWindow.name)
- AggroMeter: AggroMeter.PickedListMenu -> WindowGetId(SystemData.MouseOverWindow.name)

## Related APIs

- none

## Used With

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Triggered By

- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ListBox](../../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
