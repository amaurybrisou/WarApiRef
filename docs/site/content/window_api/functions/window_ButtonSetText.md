# ButtonSetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 24 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, BankArkel, BuffHead, CM_ClosetGoblin |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:533`, `/workspace/data/raw/AdvancedPetAssist/APAGui.lua:983`, `/workspace/data/raw/AggroMeter/AggroMeter.lua:5`, `/workspace/data/raw/Aura/Source/AuraShares.lua:68`, `/workspace/data/raw/BankArkel/BankArkel.lua:242`, `/workspace/data/raw/BankArkel/BankArkel.lua:350`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.lua:82`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.lua:142` |
| Namespaces detected | ButtonSetText |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:SetText, AdvancedPetAssist: APAGui.OnShown, AdvancedRenownTrainer: AdvancedRenownTrainer.local.SetImportExportLabels, AdvancedRenownTrainer: AdvancedRenownTrainer.local.SetLabels, AdvancedRenownTrainer: AdvancedRenownTraining.AnywhereShow, AdvancedRenownTrainer: AdvancedRenownTraining.OnHidden |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 252 |
| Global usage count | 252 |
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
ButtonSetText(windowName, text)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: "APAOptionsTabsAutoRecall", "APAOptionsTabsFollowTarget", "APAOptionsTabsGeneral" |
| text | Observed as a text or wstring payload. | Observed values: GetChatString(StringTables.Chat.LABEL_CHAT_SETTINGS_ACCEPT), GetString(StringTables.Default.LABEL_CLOSE), GetString(StringTables.Default.LABEL_LEAVE_NOW) |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- BankArkel
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Killer
- LibWBToggler
- PartyCast
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WhoHealedMe
- WoH-Reticle
- followTheLeader

## Examples

- Ace: LIBGUI_Button:SetText -> ButtonSetText(self.name, text and L "true" or L "false")
- Ace: LIBGUI_Button:SetText -> ButtonSetText(self.name, towstring(text))
- AdvancedPetAssist: APAGui.OnShown -> ButtonSetText("APAOptionsTabsGeneral", L "General")
- AdvancedPetAssist: APAGui.OnShown -> ButtonSetText("APAOptionsTabsFollowTarget", L "Follow Target")
- AdvancedPetAssist: APAGui.OnShown -> ButtonSetText("APAOptionsTabsAutoRecall", L "Auto Recall")
- AdvancedPetAssist: APAGui.OnShown -> ButtonSetText("APAOptionsTabsLos", L "LOS")

## Related APIs

- none

## Used With

- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Event
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
