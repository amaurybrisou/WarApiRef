# DynamicImage

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
| Addons seen in | AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Aura, AutoMark, BagOMatic, BankArkel, BuffHead |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.xml:13`, `/workspace_addons/AggroMeter/AggroMeter.xml:193`, `/workspace_addons/AggroMeter/AggroMeter.xml:247`, `/workspace_addons/AggroMeter/AggroMeter.xml:276`, `/workspace_addons/AggroMeter/AggroMeter.xml:304`, `/workspace_addons/AggroMeter/AggroMeter.xml:332`, `/workspace_addons/AggroMeter/AggroMeter.xml:361`, `/workspace_addons/AggroMeter/AggroMeter.xml:64` |
| Namespaces detected | DynamicImage |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedRenownTrainer: AbilityButtonTemplateSquare, AdvancedRenownTrainer: AbilityButtonTemplateSquareFrame, AdvancedRenownTrainer: AdvancedRenownTrainingWindowCornorImage, AggroMeter: AggroMeterWindow_AggroWindow1Tactic, AggroMeter: AggroMeterWindow_AggroWindow2Tactic, AggroMeter: AggroMeterWindow_AggroWindow3Tactic |
| XML usage count | 321 |
| XML attribute usage count | 321 |
| Lua usage count | 5 |
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

Observed XML element type instantiated by 41 addons.

## Common Attributes

- name
- handleinput
- texture
- layer
- textureScale
- popable
- slice
- id
- texturescale
- inherits
- sticky
- alpha

## Common Handlers

- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)

## Common Handler Functions

- Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip
- BuffHead.Setup.AdvancedContainersItem.OnContainerRClick
- BuffHead.Setup.Container.OnContainerClick
- ClosetGoblinOptionWindow.OnLButtonUp
- ClosetGoblinOptionWindow.OnRButtonUp
- MapPin.ShowIcons
- AggroMeter.OnMouseOverStart
- BagOMatic.wnd_on_lbutton_up
- BagOMatic.wnd_on_mouse_over
- CMapWindow.MouseoverMail
- CMapWindow.OnMouseoverRvRIndicator
- Enemy.CombatLogUI_IDS_OnRowLButtonDown


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | Enemy.CombatLogUI_IDS_OnRowLButtonDown | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | BuffHead.Setup.Container.OnContainerClick, ClosetGoblinOptionWindow.OnLButtonUp, MapPin.ShowIcons, BagOMatic.wnd_on_lbutton_up, Enemy.Guard_GuardIndicator_OnLButtonUp, Enemy.MarksUI_EnemyMarksWindow_OnAddLButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip, AggroMeter.OnMouseOverStart, BagOMatic.wnd_on_mouse_over, CMapWindow.MouseoverMail, CMapWindow.OnMouseoverRvRIndicator, Enemy.Guard_GuardIndicator_OnMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOverEnd | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | BuffHead.Setup.AdvancedContainersItem.OnContainerRClick, ClosetGoblinOptionWindow.OnRButtonUp, Enemy.MarksUI_EnemyMarksWindow_OnAddRButtonUp, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick, Enemy.UI_Icon_OnRButtonUp, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionRButtonUp | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonDown** handlers call: `SystemData.ActiveWindow.name:match`

**OnLButtonUp** handlers call: `ComboBoxSetSelectedMenuItem`, `Cursor.Clear`, `Cursor.IconOnCursor`, `Player.GetAbilityData`, `WindowGetId`, `WindowGetParent`, `WindowSetShowing`

**OnMouseOver** handlers call: `WindowGetId`, `WindowGetParent`

**OnRButtonUp** handlers call: `WindowGetId`, `WindowSetShowing`

## Common Inherits

- EA_ListSortDownArrow
- EA_ListSortUpArrow
- Aggro_Tactic_Template
- MapMonsterIconChooserWindow_IconTemplate
- EA_Image_DefaultIcon
- EA_Image_DefaultIconFrame
- GenericEndCap
- RVMOD_ManagerStatusTemplate
- EA_Default_CharacterImage
- EA_Default_CornerImage
- EA_Default_MerchantImage
- EA_Default_TrainingImage

## Common Parent Elements

- [Window](element_Window.md) — 132× (HIGH)
- [Button](element_Button.md) — 29× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 4× (MEDIUM)
- [Label](element_Label.md) — 1× (LOW)

## Common Named Child Elements

- [DynamicImage](element_DynamicImage.md) — 4× (MEDIUM)
- [Label](element_Label.md) — 3× (MEDIUM)

## Common Structural Child Elements

- [TintColor](element_TintColor.md) — 27× (HIGH)

## Common Template Bases

- Aggro_Tactic_Template
- EA_Default_CharacterImage
- EA_Default_CornerImage
- EA_Default_MerchantImage
- EA_Default_TrainingImage
- EA_Image_DefaultIcon
- EA_Image_DefaultIconFrame
- EA_ListSortDownArrow
- EA_ListSortUpArrow
- EA_Templates_BrassCoin
- EA_Templates_GoldCoin
- EA_Templates_SilverCoin
- GenericEndCap
- MapMonsterIconChooserWindow_IconTemplate
- RVMOD_ManagerStatusTemplate
- RvRFlagIndicator
- TargetLevelBackgroundTemplate


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `handleinput` | optional | 49% | false, true |
| `texture` | optional | 48% | Realm1, City_162, EA_HUD_01, EA_Campaign01_d5, ... |
| `layer` | optional | 41% | default, background, overlay, secondary, ... |
| `textureScale` | optional | 16% | 1, 0.4, 0.7, 0.8, ... |
| `popable` | optional | 11% | false, true |
| `slice` | optional | 11% | RvR-Flag, Zone-CONTESTED, Keep-Grayed, Dirt, ... |
| `id` | optional | 9% | 1, 2, 4, 5, ... |
| `texturescale` | optional | 9% | 0.6, 0.37, 0.8, 0.125, ... |
| `inherits` | optional | 9% | EA_ListSortUpArrow, TargetLevelBackgroundTemplate, EA_Default_TrainingImage, EA_ListSortDownArrow, ... |
| `sticky` | optional | 5% | false, true |
| `alpha` | optional | 3% | 0, 0.9, 0.1, 0.0, ... |
| `savesettings` | optional | 1% | false, true |
| `poppable` | optional | 1% | false |
| `filtering` | optional | 1% | true |
| `movable` | optional | 0% | false, true |
| `numsegments` | optional | 0% | 64 |
| `mirrorTexCoords` | optional | 0% | false |
| `parent` | optional | 0% | $parent |
| `skipinput` | optional | 0% | true |
| `Slice` | optional | 0% | Influence-Reward |
| `draganddrop` | optional | 0% | true |
| `textureAlpha` | optional | 0% | 0.5 |
## Structural Sub-Elements

### [TintColor](element_TintColor.md)

Observed 27 times as an unnamed child.

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
| `WindowGetId` | ui | 20 | OnLButtonUp, OnMouseOver, OnRButtonUp |
| `WindowSetShowing` | ui | 5 | OnLButtonUp, OnRButtonUp |
| `ComboBoxSetSelectedMenuItem` | ui | 3 | OnLButtonUp |
| `WindowGetParent` | ui | 3 | OnLButtonUp, OnMouseOver |
| `Cursor.Clear` | ui | 1 | OnLButtonUp |
| `Cursor.IconOnCursor` | data | 1 | OnLButtonUp |
| `Player.GetAbilityData` | data | 1 | OnLButtonUp |
| `SystemData.ActiveWindow.name:match` | data | 1 | OnLButtonDown |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

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

- MapMonster.MapMonster.Editor.ShowZoneHooked
- MapMonster.DisplayPinsForZone
- GetStats.GetStats.OnInitialize
- MapPin.MapPin.TestTooltip
- Pocket Palette.PP.UpdateListRow
- RoR_SoR.RoR_SoR.OnScenario
- Enemy.Enemy.AssistUI_Target_Show
- CM_ClosetGoblin.ClosetGoblin.Initialize
- MapMonster.MapMonster.local.DisplayPin
- Enemy.Enemy.GuardInitialize
- Swift Assist.Swift Assist.local.SetSmartLabel
- WSCT.WSCT.ColorOnInitialize
- RandomMount.RandomMountUI.OnAddCustomMount
- RandomMount.RandomMountUI.OnDropSlotLButtonUp
- Enemy.Enemy.MarksInitialize
- BankArkel.BankArkel.PackImg
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateSortButtonIcon
- Enemy.Enemy._Initialize
- Enemy.Enemy.TalismanAlerter_Update
- RoR_SoR.RoR_SoR.OnCombat
- MapMonster.DisplayPin
- Enemy.Enemy.Guard_OnSettingsChanged
- Enemy.Enemy.Guard_GuardIndicator_Update
- CM_ClosetGoblin.ClosetGoblinZoneWindow.OnInitialize
- MapMonster.MapMonster.local.CleanEditorWindow
- Moth.Moth.UpdateTarget
- MapMonster.CreateMarker
- Swift Assist.Swift Assist.local.SetTexLabel
- Swift Assist.SwiftAssist.OnMacroUpdated
- GuardLine.GuardLine.update
- Enemy.Enemy.UI_Icon_Switch
- BagOMatic.BagOMatic.init
- Enemy.Enemy.TalismanAlerterInitialize
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- Swift Assist.SetSmartLabel
- WSCT.WSCT.ColorAcceptButtonOnButtonUp
- WSCT.WSCT.ColorHideMenu
- Swift Assist.SetTexLabel
- TurretRange.TurretRange.OnUpdate
- Enemy.Enemy.TalismanAlerter_OnSettingsChanged
- MapMonster.MapMonster.local.CreateMarker
- MapMonster.MapMonster.local.DisplayPinsForZone
- CM_ClosetGoblin.ClosetGoblinCharacterWindow.OnInitialize
- WSCT.WSCT.OnLButtonUpColorPicker
- Swift Assist.SwiftAssist.Initialize
- Enemy.EnemyEffectsIndicator:Update
- RoR_SoR.RoR_SoR.OnInitialize


## Binding Resolution

- Total handler declarations: 59
- Resolved to Lua functions: 59 (100%)

## Seen In

- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Aura
- AutoMark
- BagOMatic
- BankArkel
- BuffHead
- CM_ClosetGoblin
- CMap
- CombatTextNames
- DAoCBuff
- EA_ThreePartBar
- EA_UiModWindow
- Effigy
- Enemy
- GetStats
- GuardLine
- JunkDump
- Killer
- LibWBToggler
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Moth
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- Swift Assist
- Targets
- TexturedButtons
- TidyRoll
- TurretRange
- WSCT

## Examples

- AdvancedRenownTrainer: AbilityButtonTemplateSquare -> DynamicImage AbilityButtonTemplateSquare
- AdvancedRenownTrainer: AbilityButtonTemplateSquareFrame -> DynamicImage AbilityButtonTemplateSquareFrame
- AdvancedRenownTrainer: AdvancedRenownTrainingWindowCornorImage -> DynamicImage AdvancedRenownTrainingWindowCornorImage
- AggroMeter: AggroMeterWindow_AggroWindow1Tactic -> DynamicImage AggroMeterWindow_AggroWindow1Tactic
- AggroMeter: AggroMeterWindow_AggroWindow2Tactic -> DynamicImage AggroMeterWindow_AggroWindow2Tactic
- AggroMeter: AggroMeterWindow_AggroWindow3Tactic -> DynamicImage AggroMeterWindow_AggroWindow3Tactic

## Related APIs

- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function

## Used With

- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none
