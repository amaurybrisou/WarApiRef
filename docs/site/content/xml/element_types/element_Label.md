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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, BankArkel, BuffHead, CM_ClosetGoblin, CombatTextNames |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/Aura/Source/AuraColorPicker.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTooltip.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0` |
| Namespaces detected | Label |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDLabel, AdvancedPetAssist: APAInstantOnlyHUDLabel, AdvancedPetAssist: APAKitingHUDLabel, AdvancedPetAssist: APALabelAttackBind, AdvancedPetAssist: APALabelAutoReattack, AdvancedPetAssist: APALabelAutoReattackDelay |
| XML usage count | 1259 |
| XML attribute usage count | 1259 |
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

Observed XML element type instantiated by 30 addons.

## Common Attributes

- name
- textalign
- font
- inherits
- handleinput
- warnOnTextCropped
- autoresize
- maxchars
- wordwrap
- layer
- popable
- skipinput

## Common Handlers

- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md)
- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)

## Common Handler Functions

- Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip
- EA_ChatWindow.OnHyperLinkLButtonUp
- EA_ChatWindow.OnHyperLinkRButtonUp
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick
- Enemy.CombatLogUI_EpsWindow_OnMouseOver
- Enemy.CombatLogUI_EpsWindow_OnRButtonUp
- Enemy.ConfigurationWindow_ShowTooltip
- Enemy.CombatLogUI_StatsWindow_ListRowMouseOver
- Enemy.CombatLogUI_TargetDefenseWindow_OnMouseOver
- Enemy.CombatLogUI_TargetDefenseWindow_OnRButtonUp
- Enemy.CombatLogUI_StatsWindow_OnEpsMouseOver


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) | custom | EA_ChatWindow.OnHyperLinkLButtonUp | `function(...)` | LOW |
| [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) | custom | EA_ChatWindow.OnHyperLinkRButtonUp | `function(...)` | LOW |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | Enemy.CombatLogUI_IDS_OnRowLButtonDown | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick, Enemy.KillSpamUI_KillSpamDialog_OnTotalLButtonUp, ClosetGoblinCharacterWindow.OnClickSetRow, ClosetGoblinZoneWindow.OnClickZoneRow, Enemy.CombatLogUI_EpsWindow_OnInDpsLButtonUp, Enemy.CombatLogUI_EpsWindow_OnInHpsLButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip, Enemy.CombatLogUI_EpsWindow_OnMouseOver, Enemy.ConfigurationWindow_ShowTooltip, Enemy.CombatLogUI_StatsWindow_ListRowMouseOver, Enemy.CombatLogUI_TargetDefenseWindow_OnMouseOver, Enemy.CombatLogUI_StatsWindow_OnEpsMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | TexturedButtons.Setup.Fonts.OnFontExampleMouseOut, BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOut, BuffHead.Setup.Layout.Properties.OnFontExampleMouseOut, BuffHead.Setup.Layout.Properties.OnPropertyTitleMouseOut, TurretRange.Setup.Display.OnDistanceFontMouseOut | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick, Enemy.CombatLogUI_EpsWindow_OnRButtonUp, Enemy.CombatLogUI_TargetDefenseWindow_OnRButtonUp, ClosetGoblinCharacterWindow.OnSetRowContextMenu, ClosetGoblinZoneWindow.OnSetZoneRowContextMenu, MiracleGrowLight.switchMode | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnHyperLinkLButtonUp** handlers call: `WindowSetShowing`

**OnLButtonDown** handlers call: `SystemData.ActiveWindow.name:match`

**OnLButtonUp** handlers call: `WindowGetId`, `WindowGetParent`, `WindowSetGameActionData`, `WindowSetShowing`

**OnMouseOver** handlers call: `CreateWindow`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `SystemData.ActiveWindow.name:match`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetGameActionData`

**OnMouseOverEnd** handlers call: `LabelSetTextColor`

**OnRButtonUp** handlers call: `SystemData.MouseOverWindow.name:match`, `WindowGetId`

## Common Inherits

- DefaultWindowSmallText
- EA_Label_DefaultText
- Aura_Default_Label_SmallFont
- ClosetGoblinTalismanLabel
- Shinies_Default_Label_ClearSmallFont
- EA_Settings_SectionTitle
- DefaultWindowText
- EA_Label_DefaultText_Small
- EA_Settings_ItemTitle
- TChatLabel
- TRollLabel
- Aggro_Label_Template

## Common Parent Elements

- [Windows](element_Windows.md) — 1243× (HIGH)
- [Window](element_Window.md) — 15× (HIGH)

## Common Structural Child Elements

- [Size](element_Size.md) — 1168× (HIGH)
- [Text](element_Text.md) — 60× (HIGH)
- [TintColor](element_TintColor.md) — 14× (HIGH)
- [Windows](element_Windows.md) — 1× (LOW)

## Common Template Bases

- Aggro_Label_Template
- Aura_CheckButtonLabel
- Aura_Default_Label
- Aura_Default_Label_SmallFont
- ClosetGoblinTalismanLabel
- DefaultWindowSmallText
- DefaultWindowSubTitle
- DefaultWindowText
- EA_Label_DefaultSubHeading
- EA_Label_DefaultText
- EA_Label_DefaultText_Small
- EA_Settings_ItemTitle
- EA_Settings_SectionTitle
- Shinies_Default_Label
- Shinies_Default_Label_ClearLargeFont
- Shinies_Default_Label_ClearMediumFont
- Shinies_Default_Label_ClearSmallFont
- TChatLabel
- TOKText
- TRollLabel


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<Label autoresize="true" font="font_clear_large_bold" handleinput="false" ignoreFormattingTags="true" layer="default" maxchars="32" name="..." sticky="false" textalign="left" wordwrap="true">
  <Size/>
  <Windows/>
</Label>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `textalign` | optional | 77% | left, center, leftcenter, right, ... |
| `font` | optional | 73% | font_clear_small_bold, font_default_text, font_clear_medium_bold, font_default_text_large, ... |
| `inherits` | optional | 29% | EA_Label_DefaultText, DefaultWindowText, TOKText, Aggro_Label_Template, ... |
| `handleinput` | optional | 29% | false, true |
| `warnOnTextCropped` | optional | 28% | false |
| `autoresize` | optional | 25% | true, false |
| `maxchars` | optional | 22% | 64, 128, 80, 256, ... |
| `wordwrap` | optional | 19% | false, true |
| `layer` | optional | 15% | secondary, popup, overlay, default, ... |
| `popable` | optional | 10% | false |
| `skipinput` | optional | 5% | true |
| `textAutoFitMinScale` | optional | 4% | 0.3 |
| `autoresizewidth` | optional | 3% | true, false |
| `sticky` | optional | 1% | true, false |
| `alpha` | optional | 1% | 1, 0 |
| `show` | optional | 1% | false |
| `ellipsisOnCrop` | optional | 0% | true |
| `autoresizeheight` | optional | 0% | true |
| `text` | optional | 0% |  |
| `autosize` | optional | 0% | true |
| `fontscale` | optional | 0% | 0.5 |
| `scale` | optional | 0% | 0.5 |
| `showing` | optional | 0% | false |
| `background` | optional | 0% | EA_FullResizeImage_TanBorder |
| `id` | optional | 0% | 1 |
| `ignoreFormattingTags` | optional | 0% | true |
## Structural Sub-Elements

### [Size](element_Size.md)

Observed 1168 times as an unnamed child.

### [Text](element_Text.md)

Observed 60 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `text` | **required** | Killer Settings, Changes apply immediately. Zone K/D history uses 0 for unlimited saved zone leaderboards., Display, Personal |
| `alignment` | optional | left |
### [TintColor](element_TintColor.md)

Observed 14 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 110, 130, 50 |
| `g` | **required** | 0, 200, 185, 130 |
| `r` | **required** | 255, 0, 180, 200 |
| `a` | optional | 30, 255, 0.5, 200 |
### [Windows](element_Windows.md)

Observed 1 times as an unnamed child.

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowGetId` | ui | 29 | OnLButtonUp, OnMouseOver, OnRButtonUp |
| `WindowSetShowing` | ui | 21 | OnHyperLinkLButtonUp, OnLButtonUp |
| `LabelSetTextColor` | ui | 12 | OnMouseOver, OnMouseOverEnd |
| `SystemData.ActiveWindow.name:match` | data | 8 | OnLButtonDown, OnMouseOver |
| `CreateWindow` | ui | 6 | OnMouseOver |
| `LabelSetText` | ui | 6 | OnMouseOver |
| `ListBoxGetDataIndex` | ui | 6 | OnMouseOver |
| `WindowGetParent` | ui | 2 | OnLButtonUp, OnMouseOver |
| `WindowSetGameActionData` | ui | 2 | OnLButtonUp, OnMouseOver |
| `SystemData.MouseOverWindow.name:match` | data | 1 | OnRButtonUp |
| `WindowGetDimensions` | ui | 1 | OnMouseOver |
| `WindowGetScreenPosition` | ui | 1 | OnMouseOver |
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

- Enemy.Timer_Update
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- RoR_SoR.OnInitialize
- Enemy.CombatLogUI_EpsWindow_Initialize
- Swift Assist.local.SetSmartLabel
- DAoCBuff.ShowMessageWindow
- Enemy.CombatLogUI_TargetDefenseTotalWindow_Update
- Swift Assist.local.WriteLabels
- Killer.ShowTopKillersTooltip
- MoraleCircle.OnSetCustomColorFull
- SwiftAssist.aaLabelColorSet
- Swift Assist.SetSmartLabel
- Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- Enemy.IntercomInitialize
- Enemy.AssistUI_Target_Show
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- MoraleCircle.ColorChanger3
- DAoCTooltips.CreateCondenseTooltip
- PP.UpdateDyeFilter
- APAGui.UpdateFollowTargetHUD
- DAoCTooltips.UpdateCondenseTooltip
- Enemy.GroupsUI_EffectFilterDialog_Open
- Enemy.UI_ConfigDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- APAGui.UpdateInstantOnlyHUD
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- BankArkel.SetCharInfo
- BankArkel.PackHide
- Enemy.CombatLogUI_EpsWindow_UpdateLayout
- MoraleCircle.ColorChanger2
- PotionBarSettings.OnAboutShown
- Swift Assist.WriteLabels
- AggroMeter.Initialize
- Killer.ShowRowTooltip
- MoraleCircle.init
- MoraleCircle.OnSetCustomColorEmpty
- MoraleCircle.ColorChanger4
- WSCT.ColorOnInitialize
- APAGui.UpdatePetTargetHUD
- Enemy.CombatLogUI_StatsWindow_ListRowMouseOver
- Killer.ShowPersonalStatsTooltip
- PP.UpdateListRow
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Enemy.CombatLogUI_StatsWindow_UpdateList
- Enemy.CombatLogUI_TargetDefenseWindow_Update
- Killer.Initialize
- MoraleCircle.OnSetCustomColorFill
- APAGui.OnShown
- APAGui.ApplyPetTargetHUDLayout
- ClosetGoblinCharacterWindow.OnInitialize
- Enemy.Guard_GuardIndicator_Update
- ClosetGoblinZoneWindow.RefreshOption
- Enemy.CombatLogUI_StatsWindow_Open
- MoraleCircle.ColorChanger1
- MoraleCircle.OnSetCustomColor
- APAGui.UpdateKitingHUD
- Enemy.IntercomUI_IntercomDialog_Open
- Enemy.CombatLogUI_EpsWindow_Update
- TidyRollOptions.Initialize
- Enemy.Stopwatch_Update


## Binding Resolution

- Total handler declarations: 192
- Resolved to Lua functions: 175 (91%)

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- BankArkel
- BuffHead
- CM_ClosetGoblin
- CombatTextNames
- DAoCBuff
- Enemy
- GuardLine
- Killer
- LibGroup
- LibWBToggler
- MiracleGrowLight
- MoraleCircle
- PartyCast
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- Swift Assist
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow

## Examples

- AdvancedPetAssist: APAFollowTargetHUDLabel -> Label APAFollowTargetHUDLabel
- AdvancedPetAssist: APAInstantOnlyHUDLabel -> Label APAInstantOnlyHUDLabel
- AdvancedPetAssist: APAKitingHUDLabel -> Label APAKitingHUDLabel
- AdvancedPetAssist: APALabelAttackBind -> Label APALabelAttackBind
- AdvancedPetAssist: APALabelAutoReattack -> Label APALabelAutoReattack
- AdvancedPetAssist: APALabelAutoReattackDelay -> Label APALabelAutoReattackDelay

## Related APIs

- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Text](element_Text.md) (HIGH 98/100) - XML Element Type

## Used With

- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) (HIGH 100/100) - XML Event
- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) (HIGH 100/100) - XML Event
- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event

## Affects

- none
