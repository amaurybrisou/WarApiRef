# wstring.match

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

## Confidence Assessment

- Level: HIGH

- Score: 100/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, Aura, CombatTextNames, EA_UiDebugTools, Enemy, MapPin, RoR_SoR, TidyChat |
| Files seen in | `/workspace/Aura/Libraries/LibPickle.lua:428`, `/workspace/Aura/Source/AuraHelpers.lua:129`, `/workspace/Enemy/Code/CombatLog/CombatLog.lua:700`, `/workspace/Enemy/Code/CombatLog/CombatLog.lua:709`, `/workspace/Enemy/Code/CombatLog/CombatLog.lua:773`, `/workspace/Enemy/Code/CombatLog/CombatLog.lua:816`, `/workspace/Enemy/Code/CombatLog/CombatLog.lua:863`, `/workspace/Enemy/Code/Core/Utils.lua:370` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed, Aura: AuraHelpers.cleanWString, Aura: DePickler:unpickle, CombatTextNames: CombatTextNames.TruncateAbilityName, EA_UiDebugTools: DevPad.TestPrint, EA_UiDebugTools: DevPadWindow.OnCodeChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 44 |
| Global usage count | 44 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
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
wstring.match(arg1, arg2)
```

## Description

Observed as a global function across 9 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: CText, DevPad_Settings.Code, Enemy.toWString(str) |
| arg2 | Observed as a function or method reference. | Observed values: BuildPattern, L "%]%[.+%]:(.+)", L "([^\^]+).*" |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- AdvancedRenownTrainer
- Aura
- CombatTextNames
- EA_UiDebugTools
- Enemy
- MapPin
- RoR_SoR
- TidyChat
- WSCT

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed -> wstring.match(link, arsenalUrl..BuildPattern)
- AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed -> wstring.match(link, BuildPattern)
- AdvancedRenownTrainer: AdvancedRenownTraining.ImportOkButtonPressed -> wstring.match(link, wardrobeUrl..BuildPattern)
- Aura: AuraHelpers.cleanWString -> wstring.match(wstr, L "([^\^]+).*")
- Aura: DePickler:unpickle -> wstring.match(s, L "P1$")
- CombatTextNames: CombatTextNames.TruncateAbilityName -> wstring.match(wstring.gsub(text,L "%l*",L ""), L "([^^]+)^?.*")

## Related APIs

- none

## Used With

- [AlertText](../tables/table_AlertText.md) (HIGH 100/100) - Global Table
- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeTextEntryDialog](global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatTabManager.GetTabName](global_EA_ChatTabManager.GetTabName.md) (HIGH 100/100) - Global Function
- [EA_ChatWindowGroups](../tables/table_EA_ChatWindowGroups.md) (HIGH 100/100) - Global Table
- [GameData.ChatData.name](../../gamedata/fields/gamedata_GameData.ChatData.name.md) (HIGH 100/100) - GameData Field
- [GameData.ChatData.text](../../gamedata/fields/gamedata_GameData.ChatData.text.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.QUEST_ABANDONED](../../gamedata/fields/gamedata_GameData.Sound.QUEST_ABANDONED.md) (HIGH 100/100) - GameData Field
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.gsub](global_wstring.gsub.md) (HIGH 100/100) - Global Function
- [AlertTextWindow.AddAlert](global_AlertTextWindow.AddAlert.md) (HIGH 88/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
