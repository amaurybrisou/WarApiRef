# Label

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, BankArkel, BuffHead, Busted, CM_ClosetGoblin |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:1012`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1034`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1056`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1068`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1090`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1102`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1114`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1126` |
| Namespaces detected | Label |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDLabel, AdvancedPetAssist: APAInstantOnlyHUDLabel, AdvancedPetAssist: APAKitingHUDLabel, AdvancedPetAssist: APALabelAttackBind, AdvancedPetAssist: APALabelAutoReattack, AdvancedPetAssist: APALabelAutoReattackDelay |
| XML usage count | 1572 |
| XML attribute usage count | 1572 |
| Lua usage count | 7 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
| Consistent role | yes |
| Consistent arguments | no |
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

## Description

Observed XML element type instantiated by 56 addons.

## Common Attributes

- name
- textalign
- font
- inherits
- handleinput
- autoresize
- warnOnTextCropped
- maxchars
- wordwrap
- layer
- popable
- skipinput

## Common Handlers

- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md)
- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)

## Common Handler Functions

- LPET.OnMouseOver
- Cheeseboard.SortColumnClick
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip
- Cheeseboard.OnMouseOver
- FrameManager.OnMouseOverEnd
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick
- Enemy.CombatLogUI_EpsWindow_OnMouseOver
- Enemy.CombatLogUI_EpsWindow_OnRButtonUp
- EA_ChatWindow.OnHyperLinkLButtonUp
- EA_ChatWindow.OnHyperLinkRButtonUp
- Enemy.ConfigurationWindow_ShowTooltip


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) | custom | EA_ChatWindow.OnHyperLinkLButtonUp | `function(...)` | LOW |
| [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) | custom | EA_ChatWindow.OnHyperLinkRButtonUp, MapPin.RButtonUp | `function(...)` | LOW |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | Enemy.CombatLogUI_IDS_OnRowLButtonDown | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | Cheeseboard.SortColumnClick, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick, Enemy.KillSpamUI_KillSpamDialog_OnTotalLButtonUp, ClosetGoblinCharacterWindow.OnClickSetRow, ClosetGoblinZoneWindow.OnClickZoneRow, Enemy.CombatLogUI_EpsWindow_OnInDpsLButtonUp | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | LPET.OnMouseOver, Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip, Cheeseboard.OnMouseOver, Enemy.CombatLogUI_EpsWindow_OnMouseOver, Enemy.ConfigurationWindow_ShowTooltip, Enemy.CombatLogUI_StatsWindow_ListRowMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | FrameManager.OnMouseOverEnd, TexturedButtons.Setup.Fonts.OnFontExampleMouseOut, BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOut, BuffHead.Setup.Layout.Properties.OnFontExampleMouseOut, BuffHead.Setup.Layout.Properties.OnPropertyTitleMouseOut, TurretRange.Setup.Display.OnDistanceFontMouseOut | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick, Enemy.CombatLogUI_EpsWindow_OnRButtonUp, Enemy.CombatLogUI_TargetDefenseWindow_OnRButtonUp, ClosetGoblinCharacterWindow.OnSetRowContextMenu, ClosetGoblinZoneWindow.OnSetZoneRowContextMenu, MiracleGrowLight.switchMode | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnHyperLinkLButtonUp** handlers call: `WindowSetShowing`

**OnHyperLinkRButtonUp** handlers call: `ButtonSetPressedFlag`, `ComboBoxSetSelectedMenuItem`, `SliderBarSetCurrentPosition`, `TextEditBoxSetText`, `WindowGetId`, `WindowSetMovable`, `WindowSetShowing`

**OnLButtonDown** handlers call: `SystemData.ActiveWindow.name:match`

**OnLButtonUp** handlers call: `WindowGetId`, `WindowGetParent`, `WindowSetGameActionData`, `WindowSetShowing`

**OnMouseOver** handlers call: `CreateWindow`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `SystemData.ActiveWindow.name:match`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetGameActionData`

**OnMouseOverEnd** handlers call: `LabelSetTextColor`

**OnRButtonUp** handlers call: `SystemData.MouseOverWindow.name:match`, `WindowGetId`

## Common Inherits

- EA_Label_DefaultText
- Aura_Default_Label_SmallFont
- ClosetGoblinTalismanLabel
- Shinies_Default_Label_ClearSmallFont
- ModDetailsTextDef
- ModDetailsLabelDef
- DefaultWindowSmallText
- EA_Settings_ItemTitle
- MapMonsterEditorWindowLabelDefault
- DefaultWindowText
- EA_Label_DefaultText_Small
- IraConfigHeading

## Common Parent Elements

- [Window](element_Window.md) — 379× (HIGH)
- [Button](element_Button.md) — 28× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 3× (MEDIUM)
- [FullResizeImage](element_FullResizeImage.md) — 3× (MEDIUM)
- [SliderBar](element_SliderBar.md) — 2× (LOW)
- [CircleImage](element_CircleImage.md) — 1× (LOW)
- [EditBox](element_EditBox.md) — 1× (LOW)

## Common Named Child Elements

- [DynamicImage](element_DynamicImage.md) — 1× (LOW)

## Common Structural Child Elements

- [Text](element_Text.md) — 69× (HIGH)
- [TintColor](element_TintColor.md) — 14× (HIGH)

## Common Template Bases

- Aggro_Label_Template
- Aura_CheckButtonLabel
- Aura_Default_Label
- Aura_Default_Label_SmallFont
- ClosetGoblinTalismanLabel
- DefaultWindowSmallText
- DefaultWindowSubTitle
- DefaultWindowText
- EA_CheckButtonLabel
- EA_Label_DefaultSubHeading
- EA_Label_DefaultText
- EA_Label_DefaultText_Small
- EA_Settings_ItemTitle
- EA_Settings_SectionTitle
- IraConfigHeading
- IraConfigLabel
- MapMonsterEditorWindowHeaderDefault
- MapMonsterEditorWindowLabelDefault
- MapMonsterPinTypeEditorWindowHeaderDefault
- MapMonsterPinTypeEditorWindowLabelDefault
- MiracleGrow2Label
- ModDetailsLabelDef
- ModDetailsTextDef
- Shinies_Default_Label
- Shinies_Default_Label_ClearLargeFont
- Shinies_Default_Label_ClearMediumFont
- Shinies_Default_Label_ClearSmallFont
- TChatLabel
- TOKHeadingSmall
- TOKText
- TRollLabel


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `textalign` | optional | 58% | left, center, right, rightcenter, ... |
| `font` | optional | 56% | font_clear_tiny, font_chat_text, font_clear_small, font_clear_small_bold, ... |
| `inherits` | optional | 24% | EA_Label_DefaultText, TRollLabel, Aura_Default_Label_SmallFont, IraConfigHeading, ... |
| `handleinput` | optional | 22% | false, true |
| `autoresize` | optional | 18% | true, false |
| `warnOnTextCropped` | optional | 17% | false |
| `maxchars` | optional | 16% | 100, 50, 255, 200, ... |
| `wordwrap` | optional | 13% | true, false |
| `layer` | optional | 13% | default, overlay, background, secondary, ... |
| `popable` | optional | 5% | false |
| `skipinput` | optional | 3% | true |
| `autoresizewidth` | optional | 2% | false, true |
| `textAutoFitMinScale` | optional | 2% | 0.3, 0.5, 0.75 |
| `hanldeinput` | optional | 1% | false |
| `sticky` | optional | 1% | false, true |
| `alpha` | optional | 0% | .1, 1, 0 |
| `show` | optional | 0% | false |
| `handlinput` | optional | 0% | false |
| `text` | optional | 0% |  |
| `autoresizeheight` | optional | 0% | true |
| `ellipsisOnCrop` | optional | 0% | true |
| `ignoreFormattingTags` | optional | 0% | true |
| `linespacing` | optional | 0% | 13, 12 |
| `showing` | optional | 0% | false |
| `background` | optional | 0% | EA_FullResizeImage_TanBorder |
| `id` | optional | 0% | 1 |
## Structural Sub-Elements

### [Text](element_Text.md)

Observed 69 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `text` | optional |  |
| `alignment` | optional |  |
### [TintColor](element_TintColor.md)

Observed 14 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** |  |
| `g` | **required** |  |
| `r` | **required** |  |
| `a` | optional |  |
## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowGetParent` | ui | 86 | OnLButtonUp, OnMouseOver |
| `WindowGetId` | ui | 30 | OnHyperLinkRButtonUp, OnLButtonUp, OnMouseOver, OnRButtonUp |
| `WindowSetShowing` | ui | 18 | OnHyperLinkLButtonUp, OnHyperLinkRButtonUp, OnLButtonUp |
| `LabelSetTextColor` | ui | 12 | OnMouseOver, OnMouseOverEnd |
| `TextEditBoxSetText` | ui | 12 | OnHyperLinkRButtonUp |
| `SystemData.ActiveWindow.name:match` | data | 8 | OnLButtonDown, OnMouseOver |
| `CreateWindow` | ui | 6 | OnMouseOver |
| `LabelSetText` | ui | 6 | OnMouseOver |
| `ListBoxGetDataIndex` | ui | 6 | OnMouseOver |
| `ComboBoxSetSelectedMenuItem` | ui | 4 | OnHyperLinkRButtonUp |
| `SliderBarSetCurrentPosition` | ui | 2 | OnHyperLinkRButtonUp |
| `WindowSetGameActionData` | ui | 2 | OnLButtonUp, OnMouseOver |
| `ButtonSetPressedFlag` | ui | 1 | OnHyperLinkRButtonUp |
| `SystemData.MouseOverWindow.name:match` | data | 1 | OnRButtonUp |
| `WindowGetDimensions` | ui | 1 | OnMouseOver |
| `WindowGetScreenPosition` | ui | 1 | OnMouseOver |
| `WindowSetMovable` | ui | 1 | OnHyperLinkRButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnHyperLinkLButtonUp

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `linkData` | table | data |
| 1 | `flags` | boolean | modifier_flags |
| 2 | `x` | number | mouse_x |
| 3 | `y` | number | mouse_y |
### OnHyperLinkRButtonUp

Confidence: LOW

### OnLButtonDown

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMouseOver

Confidence: HIGH

### OnMouseOverEnd

Confidence: HIGH

### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
## Lua Functions Manipulating This Type

- BankArkel.BankArkel.PackHide
- EA_UiDebugTools.BustedGUI.Initialize
- QuickWarReport.ShowConfirmWindow
- Busted.BustedGUI.NewErrorRecorded
- AdvancedPetAssist.APAGui.UpdateKitingHUD
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- TidyRoll.TidyRollOptions.Initialize
- MoraleCircle.MoraleCircle.OnSetCustomColor
- DAoCBuff.DAoCTooltips.CreateCondenseTooltip
- wbLeadHelper.wbLeadHelper.onZoneMouseOver
- Enemy.Enemy.UI_ConfigDialog_Open
- Enemy.Enemy.CombatLogUI_EpsWindow_Update
- AdvancedPetAssist.APAGui.UpdatePetTargetHUD
- CM_ClosetGoblin.ClosetGoblinZoneWindow.RefreshOption
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.OnInitialize
- MoraleCircle.MoraleCircle.OnSetCustomColorEmpty
- MoraleCircle.MoraleCircle.ColorChanger4
- QuickWarReport.PrepareConfirmWindowChrome
- Swift Assist.WriteLabels
- MoraleCircle.MoraleCircle.OnSetCustomColorFull
- MoraleCircle.MoraleCircle.ColorChanger3
- Enemy.Enemy.IntercomInitialize
- Killer.Killer.ShowTopKillersTooltip
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_Open
- MoraleCircle.MoraleCircle.init
- EA_UiDebugTools.BustedGUI.NewErrorRecorded
- Enemy.Enemy.CombatLogUI_StatsWindow_ListRowMouseOver
- Swift Assist.SwiftAssist.aaLabelColorSet
- Enemy.Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- Enemy.Enemy.Stopwatch_Update
- Enemy.Enemy.AssistUI_Target_Show
- RoR_SoR.RoR_SoR.OnInitialize
- Enemy.Enemy.GroupsUI_EffectFilterDialog_Open
- MoraleCircle.MoraleCircle.OnSetCustomColorFill
- PotionBar.PotionBarSettings.OnAboutShown
- Enemy.Enemy.Guard_GuardIndicator_Update
- Pocket Palette.PP.UpdateListRow
- Busted.BustedGUI.UpdateErrorView
- wbLeadHelper.WbLeadHelperMessage.MessageDialogOpen
- WSCT.WSCT.ColorOnInitialize
- AggroMeter.AggroMeter.Initialize
- BankArkel.BankArkel.SetCharInfo
- Busted.BustedGUI.Initialize
- AdvancedPetAssist.APAGui.UpdateFollowTargetHUD
- Busted.BustedGUI.ClearAlertFlash
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- AdvancedPetAssist.APAGui.UpdateInstantOnlyHUD
- Killer.Killer.ShowRowTooltip
- GetStats.GetStats.OnChatLogUpdated
- MoraleCircle.MoraleCircle.ColorChanger1
- QuickWarReport.QuickWarReport.TestConfirmWindow
- Enemy.Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- DaemonAssist.DaemonAssist.UpdateWindow
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- EA_UiDebugTools.BustedGUI.UpdateErrorView
- RandomMount.RandomMountUI.OnDropSlotLButtonUp
- QuickWarReport.QuickWarReport.local.PrepareConfirmWindowChrome
- Swift Assist.Swift Assist.local.WriteLabels
- Enemy.Enemy.CombatLogUI_StatsWindow_Open
- Swift Assist.Swift Assist.local.SetSmartLabel
- Enemy.Enemy.CombatLogUI_StatsWindow_UpdateList
- Killer.Killer.ShowPersonalStatsTooltip
- RandomMount.RandomMountUI.OnAddCustomMount
- Enemy.Enemy.CombatLogUI_EpsWindow_Initialize
- MoraleCircle.MoraleCircle.ColorChanger2
- EA_UiDebugTools.BustedGUI.ClearAlertFlash
- Killer.Killer.Initialize
- FastInteract.FastInteract.OptionsSetup
- QuickWarReport.QuickWarReport.local.ShowConfirmWindow
- AdvancedPetAssist.APAGui.ApplyPetTargetHUDLayout
- DAoCBuff.DAoCTooltips.UpdateCondenseTooltip
- Enemy.Enemy.Timer_Update
- AdvancedPetAssist.APAGui.OnShown
- RandomMount.RandomMountUI.OnInitialize
- DAoCBuff.DAoCBuff.ShowMessageWindow
- Swift Assist.SetSmartLabel
- Enemy.Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- Enemy.Enemy.CombatLogUI_TargetDefenseWindow_Update
- Enemy.Enemy.IntercomUI_IntercomDialog_Open
- Enemy.Enemy.CombatLogUI_EpsWindow_UpdateLayout
- wbLeadHelper.wbLeadHelper.showNormalTitle
- Pocket Palette.PP.UpdateDyeFilter
- Enemy.Enemy.CombatLogUI_TargetDefenseTotalWindow_Update


## Binding Resolution

- Total handler declarations: 287
- Resolved to Lua functions: 287 (100%)

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- BankArkel
- BuffHead
- Busted
- CM_ClosetGoblin
- CMap
- Cheeseboard
- CombatTextNames
- Crafting Info Tooltip
- DAoCBuff
- DaemonAssist
- EA_UiDebugTools
- EA_UiModWindow
- Effigy
- Enemy
- FastInteract
- GetStats
- GuardLine
- GuardList
- GuardRange
- JunkDump
- Killer
- LibGroup
- LibWBToggler
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- Moth
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- TidyChat
- TidyRoll
- Tortall_DPS
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAFollowTargetHUDLabel -> Label APAFollowTargetHUDLabel
- AdvancedPetAssist: APAInstantOnlyHUDLabel -> Label APAInstantOnlyHUDLabel
- AdvancedPetAssist: APAKitingHUDLabel -> Label APAKitingHUDLabel
- AdvancedPetAssist: APALabelAttackBind -> Label APALabelAttackBind
- AdvancedPetAssist: APALabelAutoReattack -> Label APALabelAutoReattack
- AdvancedPetAssist: APALabelAutoReattackDelay -> Label APALabelAutoReattackDelay

## Related APIs

- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function

## Used With

- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnHyperLinkLButtonUp](../../events/window_events/window_event_OnHyperLinkLButtonUp.md) (HIGH 100/100) - Window Event
- [OnHyperLinkRButtonUp](../../events/window_events/window_event_OnHyperLinkRButtonUp.md) (HIGH 100/100) - Window Event
- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOverEnd](../../events/window_events/window_event_OnMouseOverEnd.md) (HIGH 100/100) - Window Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none
