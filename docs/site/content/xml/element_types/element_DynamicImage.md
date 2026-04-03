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
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/AutoMark/Source/AutoMark.xml:0`, `/workspace/data/raw/BankArkel/BankArkel.xml:0`, `/workspace/data/raw/BuffHead/Display.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0` |
| Namespaces detected | DynamicImage |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedRenownTrainer: AbilityButtonTemplateSquare, AdvancedRenownTrainer: AbilityButtonTemplateSquareFrame, AdvancedRenownTrainer: AdvancedRenownTrainingWindowCornorImage, AggroMeter: AggroMeterWindow_AggroWindow1Tactic, AggroMeter: AggroMeterWindow_AggroWindow2Tactic, AggroMeter: AggroMeterWindow_AggroWindow3Tactic |
| XML usage count | 237 |
| XML attribute usage count | 237 |
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

Observed XML element type instantiated by 26 addons.

## Common Attributes

- name
- handleinput
- texture
- layer
- textureScale
- id
- popable
- slice
- inherits
- texturescale
- filtering
- alpha

## Common Handlers

- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)

## Common Handler Functions

- DAoCBuffSettings.ShowTooltip
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip
- BuffHead.Setup.AdvancedContainersItem.OnContainerRClick
- BuffHead.Setup.Container.OnContainerClick
- ClosetGoblinOptionWindow.OnLButtonUp
- ClosetGoblinOptionWindow.OnRButtonUp
- AggroMeter.OnMouseOverStart
- BagOMatic.wnd_on_lbutton_up
- BagOMatic.wnd_on_mouse_over
- Enemy.CombatLogUI_IDS_OnRowLButtonDown
- Enemy.Guard_GuardIndicator_OnLButtonUp
- Enemy.Guard_GuardIndicator_OnMouseOver


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | Enemy.CombatLogUI_IDS_OnRowLButtonDown | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | BuffHead.Setup.Container.OnContainerClick, ClosetGoblinOptionWindow.OnLButtonUp, BagOMatic.wnd_on_lbutton_up, Enemy.Guard_GuardIndicator_OnLButtonUp, Enemy.MarksUI_EnemyMarksWindow_OnAddLButtonUp, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick | `flags, x, y` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | DAoCBuffSettings.ShowTooltip, Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip, AggroMeter.OnMouseOverStart, BagOMatic.wnd_on_mouse_over, Enemy.Guard_GuardIndicator_OnMouseOver, Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOverEnd | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | BuffHead.Setup.AdvancedContainersItem.OnContainerRClick, ClosetGoblinOptionWindow.OnRButtonUp, Enemy.MarksUI_EnemyMarksWindow_OnAddRButtonUp, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick, Enemy.UI_Icon_OnRButtonUp, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionRButtonUp | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonDown** handlers call: `SystemData.ActiveWindow.name:match`

**OnLButtonUp** handlers call: `ComboBoxSetSelectedMenuItem`, `Cursor.Clear`, `Cursor.IconOnCursor`, `Player.GetAbilityData`, `WindowGetId`, `WindowSetShowing`

**OnMouseOver** handlers call: `WindowGetId`, `WindowGetParent`

**OnRButtonUp** handlers call: `WindowGetId`

## Common Inherits

- Aggro_Tactic_Template
- EA_ListSortDownArrow
- EA_ListSortUpArrow
- EA_Default_CharacterImage
- EA_Default_TrainingImage
- EA_Image_DefaultIcon
- EA_Image_DefaultIconFrame
- EA_Templates_BrassCoin
- EA_Templates_GoldCoin
- EA_Templates_SilverCoin

## Common Parent Elements

- [Windows](element_Windows.md) — 237× (HIGH)

## Common Structural Child Elements

- [Size](element_Size.md) — 190× (HIGH)
- [TexDims](element_TexDims.md) — 65× (HIGH)
- [TintColor](element_TintColor.md) — 25× (HIGH)
- [TexCoords](element_TexCoords.md) — 14× (HIGH)
- [Windows](element_Windows.md) — 6× (MEDIUM)

## Common Template Bases

- Aggro_Tactic_Template
- EA_Default_CharacterImage
- EA_Default_TrainingImage
- EA_Image_DefaultIcon
- EA_Image_DefaultIconFrame
- EA_ListSortDownArrow
- EA_ListSortUpArrow
- EA_Templates_BrassCoin
- EA_Templates_GoldCoin
- EA_Templates_SilverCoin


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<DynamicImage name="..." savesettings="true" texture="enemy_guard_big">
  <TexCoords x="0" y="0"/>
  <TexDims x="150" y="150"/>
  <Size/>
  <Windows/>
</DynamicImage>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `handleinput` | optional | 78% | false, true |
| `texture` | optional | 75% | EA_SquareFrame, AggroMeterIcon, EA_Abilities01_d5, EA_ScreenFlash, ... |
| `layer` | optional | 57% | background, overlay, popup, default, ... |
| `textureScale` | optional | 26% | 1.0, 0.5, 1.171, 1, ... |
| `id` | optional | 21% | 1, 2, 3, 103, ... |
| `popable` | optional | 15% | false, true |
| `slice` | optional | 13% | Tab-ALL, Round-Swatch-Selection-Ring, Radio-Button, LordHeroSpecial-Skull, ... |
| `inherits` | optional | 10% | EA_Default_TrainingImage, Aggro_Tactic_Template, EA_ListSortUpArrow, EA_ListSortDownArrow, ... |
| `texturescale` | optional | 9% | 0.35, 1.00, .5, 0.3, ... |
| `filtering` | optional | 8% | true |
| `alpha` | optional | 3% | 1.0, 0.5, 0.1, 0, ... |
| `sticky` | optional | 2% | false |
| `savesettings` | optional | 2% | false, true |
| `movable` | optional | 1% | false, true |
| `mirrorTexCoords` | optional | 0% | false |
| `parent` | optional | 0% | $parent |
| `skipinput` | optional | 0% | true |
| `Slice` | optional | 0% | Influence-Reward |
| `draganddrop` | optional | 0% | true |
| `textureAlpha` | optional | 0% | 0.5 |
## Structural Sub-Elements

### [Size](element_Size.md)

Observed 190 times as an unnamed child.

### [TexDims](element_TexDims.md)

Observed 65 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 64, 256, 32, 430 |
| `y` | **required** | 64, 256, 32, 430 |
### [TintColor](element_TintColor.md)

Observed 25 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 110, 130, 50 |
| `g` | **required** | 0, 200, 185, 130 |
| `r` | **required** | 255, 0, 180, 200 |
| `a` | optional | 30, 255, 0.5, 200 |
### [TexCoords](element_TexCoords.md)

Observed 14 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 0, 133, 88, 32 |
| `y` | optional | 0, 163, 51, 32 |
### [Windows](element_Windows.md)

Observed 6 times as an unnamed child.

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowGetId` | ui | 26 | OnLButtonUp, OnMouseOver, OnRButtonUp |
| `ComboBoxSetSelectedMenuItem` | ui | 3 | OnLButtonUp |
| `Cursor.Clear` | ui | 1 | OnLButtonUp |
| `Cursor.IconOnCursor` | data | 1 | OnLButtonUp |
| `Player.GetAbilityData` | data | 1 | OnLButtonUp |
| `SystemData.ActiveWindow.name:match` | data | 1 | OnLButtonDown |
| `WindowGetParent` | ui | 1 | OnMouseOver |
| `WindowSetShowing` | ui | 1 | OnLButtonUp |
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

- Enemy.TalismanAlerterInitialize
- PP.UpdateListRow
- RoR_SoR.OnInitialize
- BankArkel.PackImg
- Enemy.GuardInitialize
- SwiftAssist.Initialize
- Swift Assist.local.SetSmartLabel
- WSCT.ColorOnInitialize
- Enemy.Guard_OnSettingsChanged
- Enemy.EnemyEffectsIndicator:Update
- Enemy.MarksInitialize
- RoR_SoR.OnCombat
- Swift Assist.local.SetTexLabel
- Swift Assist.SetSmartLabel
- ClosetGoblinCharacterWindow.UpdateSortButtonIcon
- GuardLine.update
- RoR_SoR.OnScenario
- SwiftAssist.OnMacroUpdated
- Enemy.AssistUI_Target_Show
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- Enemy.Guard_GuardIndicator_Update
- Enemy.TalismanAlerter_Update
- BagOMatic.init
- WSCT.OnLButtonUpColorPicker
- Enemy.UI_Icon_Switch
- ClosetGoblinZoneWindow.OnInitialize
- ClosetGoblin.Initialize
- Enemy._Initialize
- Swift Assist.SetTexLabel
- TurretRange.OnUpdate
- WSCT.ColorAcceptButtonOnButtonUp
- ClosetGoblinCharacterWindow.OnInitialize
- WSCT.ColorHideMenu
- Enemy.TalismanAlerter_OnSettingsChanged


## Binding Resolution

- Total handler declarations: 52
- Resolved to Lua functions: 52 (100%)

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
- CombatTextNames
- DAoCBuff
- Enemy
- GuardLine
- Killer
- LibWBToggler
- MiracleGrowLight
- PartyCast
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- Swift Assist
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
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TexDims](element_TexDims.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function

## Used With

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event

## Affects

- none
