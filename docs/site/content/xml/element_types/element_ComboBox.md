# ComboBox

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BankArkel, BuffHead, Busted, Cheeseboard, Crafting Info Tooltip |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:1000`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1022`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1044`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1078`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1136`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1194`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1214`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1234` |
| Namespaces detected | ComboBox |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedPetAssist: APAComboAttackBind, AdvancedPetAssist: APAComboAutoReattack, AdvancedPetAssist: APAComboAutoReattackDelay, AdvancedPetAssist: APAComboCastDelay, AdvancedPetAssist: APAComboCastOnAcquire, AdvancedPetAssist: APAComboCombatExitDelay |
| XML usage count | 347 |
| XML attribute usage count | 347 |
| Lua usage count | 2 |
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

Observed XML element type instantiated by 29 addons.

## Common Attributes

- name
- inherits
- layer
- maxvisibleitems
- menubackground
- menuitembutton
- scrollbar
- selectedbutton
- show
- autoresize
- id
- warnOnTextCropped

## Common Handlers

- [OnSelChanged](../handlers/handler_OnSelChanged.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)

## Common Handler Functions

- LPET.OnMouseOver
- APAGui.OnComboChanged
- LPET.QuickHealthOnSelChanged
- LPET.QuickModeOnSelChanged
- LPET.QuickPriorityOnSelChanged
- WHMGui.OnOptionsComboChanged
- MapPin.ComboBoxUpdate
- BustedGUI.SelectAddonView
- DaemonAssist.OnBindingComboChanged
- Enemy.ConfigurationWindow_OnChange
- Enemy.ConfigurationWindow_ShowTooltip
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | LPET.OnMouseOver, Enemy.ConfigurationWindow_ShowTooltip, DevPadWindow.OnMouseOverCode, PotionBarSettings.OnMouseoverActivator, PotionBarSettings.OnMouseoverBuild, PotionBarSettings.OnMouseoverInfoTextBR | `function()` | MEDIUM |
| [OnSelChanged](../handlers/handler_OnSelChanged.md) | data | APAGui.OnComboChanged, LPET.QuickHealthOnSelChanged, LPET.QuickModeOnSelChanged, LPET.QuickPriorityOnSelChanged, WHMGui.OnOptionsComboChanged, MapPin.ComboBoxUpdate | `index` | MEDIUM |

### Per-Event Lua API Calls

**OnMouseOver** handlers call: `ComboBoxGetSelectedText`, `WindowGetParent`

**OnSelChanged** handlers call: `ButtonGetPressedFlag`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `ButtonSetText`, `ComboBoxGetSelectedMenuItem`, `ComboBoxGetSelectedText`, `ComboBoxSetSelectedMenuItem`, `CreateWindowFromTemplate`, `DestroyWindow`, `DoesWindowExist`, `DynamicImageSetTexture`, `DynamicImageSetTextureSlice`, `LabelSetText`, `LabelSetTextColor`, `ListBoxSetDisplayOrder`, `ScrollWindowUpdateScrollRect`, `SliderBarGetCurrentPosition`, `TextEditBoxGetText`, `TextEditBoxSetText`, `TextEditBoxSetTextColor`, `WindowAddAnchor`, `WindowClearAnchors`, `WindowGetId`, `WindowGetParent`, `WindowSetHandleInput`, `WindowSetShowing`

## Common Inherits

- EA_ComboBox_DefaultResizable_Fixed_Small
- EA_ComboBox_DefaultResizable
- APA_ComboBox
- EA_ComboBox_DefaultResizableSmall
- Aura_ComboBox_DefaultResizableTiny
- Aura_ComboBox_DefaultResizableLarge
- EA_ComboBox_DefaultResizableLarge
- EA_ComboBox_DefaultResizable_Fixed
- IraConfigCombo
- Shinies_ComboBox_DefaultResizableLarge
- APA_ComboBoxWide
- DaemonAssist_ComboBoxWide

## Common Parent Elements

- [Window](element_Window.md) — 97× (HIGH)

## Common Structural Child Elements

- [MenuButtonOffset](element_MenuButtonOffset.md) — 13× (HIGH)
- [TextOffset](element_TextOffset.md) — 1× (LOW)

## Common Template Bases

- APA_ComboBox
- APA_ComboBoxWide
- Aura_ComboBox_DefaultResizable
- Aura_ComboBox_DefaultResizableLarge
- Aura_ComboBox_DefaultResizableTiny
- DaemonAssist_ComboBoxWide
- EA_ComboBox_DefaultResizable
- EA_ComboBox_DefaultResizable3
- EA_ComboBox_DefaultResizableLarge
- EA_ComboBox_DefaultResizableMedium
- EA_ComboBox_DefaultResizableSmall
- EA_ComboBox_DefaultResizableTiny
- EA_ComboBox_DefaultResizable_Fixed
- EA_ComboBox_DefaultResizable_Fixed_Small
- IraConfigCombo
- MapMonsterPinTypeEditor_ComboBox_DefaultResizable
- Shinies_ComboBox_DefaultResizableLarge


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<ComboBox name="..." selectedbutton="EA_Button_DefaultResizableCom..." menubackground="EA_Window_ComboBoxMenuBackground" menuitembutton="EA_Button_DefaultMenuButton" scrollbar="EA_ScrollBar_DefaultVerticalC..." maxvisibleitems="10">
  <MenuButtonOffset x="5" y="5"/>
  <TextOffset x="0" y="5"/>
</ComboBox>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 75% | APA_ComboBox, EA_ComboBox_DefaultResizable, EA_ComboBox_DefaultResizable_Fixed_Small, EA_ComboBox_DefaultResizableSmall, ... |
| `layer` | optional | 43% | secondary, popup, default, overlay |
| `maxvisibleitems` | optional | 6% | 5, 8, 15, 10, ... |
| `menubackground` | optional | 2% | EA_Window_ComboBoxMenuBackground, EA_Window_DefaultFrame |
| `menuitembutton` | optional | 2% | Shinies_Button_DefaultMenuButtonLarge, EA_Button_DefaultMenuButtonSmall, EA_Button_DefaultMenuButton, Aura_Button_DefaultMenuButtonLarge, ... |
| `scrollbar` | optional | 2% | EA_ScrollBar_DefaultVerticalChain, EA_ScrollBar_ChatVertical |
| `selectedbutton` | optional | 2% | Shinies_ComboBox_DefaultResizableComboBoxSelectedLarge, APA_ComboBoxButton, DaemonAssist_ComboBoxButtonWide, APA_ComboBoxButtonWide, ... |
| `show` | optional | 2% | false |
| `autoresize` | optional | 2% | true |
| `id` | optional | 0% | 1, 2, 3 |
| `warnOnTextCropped` | optional | 0% | false |
## Structural Sub-Elements

### [MenuButtonOffset](element_MenuButtonOffset.md)

Observed 13 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
### [TextOffset](element_TextOffset.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowGetParent` | ui | 227 | OnMouseOver, OnSelChanged |
| `ComboBoxGetSelectedMenuItem` | ui | 95 | OnSelChanged |
| `TextEditBoxGetText` | ui | 20 | OnSelChanged |
| `WindowSetShowing` | ui | 19 | OnSelChanged |
| `LabelSetText` | ui | 16 | OnSelChanged |
| `SliderBarGetCurrentPosition` | ui | 8 | OnSelChanged |
| `ButtonSetDisabledFlag` | ui | 7 | OnSelChanged |
| `ButtonGetPressedFlag` | ui | 6 | OnSelChanged |
| `ButtonSetPressedFlag` | ui | 6 | OnSelChanged |
| `ComboBoxGetSelectedText` | ui | 6 | OnMouseOver, OnSelChanged |
| `WindowSetHandleInput` | ui | 6 | OnSelChanged |
| `DynamicImageSetTexture` | ui | 4 | OnSelChanged |
| `ButtonSetText` | ui | 3 | OnSelChanged |
| `WindowAddAnchor` | ui | 3 | OnSelChanged |
| `ComboBoxSetSelectedMenuItem` | ui | 2 | OnSelChanged |
| `DoesWindowExist` | ui | 2 | OnSelChanged |
| `DynamicImageSetTextureSlice` | ui | 2 | OnSelChanged |
| `LabelSetTextColor` | ui | 2 | OnSelChanged |
| `ListBoxSetDisplayOrder` | ui | 2 | OnSelChanged |
| `TextEditBoxSetTextColor` | ui | 2 | OnSelChanged |
| `WindowClearAnchors` | ui | 2 | OnSelChanged |
| `CreateWindowFromTemplate` | ui | 1 | OnSelChanged |
| `DestroyWindow` | ui | 1 | OnSelChanged |
| `ScrollWindowUpdateScrollRect` | ui | 1 | OnSelChanged |
| `TextEditBoxSetText` | ui | 1 | OnSelChanged |
| `WindowGetId` | ui | 1 | OnSelChanged |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnMouseOver

Confidence: HIGH

### OnSelChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `index` | number | selected_index |
## Lua Functions Manipulating This Type

- MapPin.MapPin.RButtonUp
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Busted.BustedGUI.NewErrorRecorded
- Busted.BustedGUI.Initialize
- EA_UiDebugTools.BustedGUI.NewErrorRecorded
- AdvancedRenownTrainer.GeneratePresetByLinkData
- Enemy.Enemy.IntercomUI_ChooseChannelDialog_Open
- Enemy.Enemy.UI_ConfigDialog_Open
- Enemy.Enemy.UI_ConfigDialog_OnSectionSelChanged
- Enemy.Enemy.CombatLogUI_StatsWindow_UpdateSessionsList
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- Enemy.Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- MapPin.MapPin.local.EditMarker
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- BankArkel.BankArkel.SetupCombos
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.Enemy.UnitFramesUI_UnitFramePartDialog_Open
- EA_UiDebugTools.BustedGUI.Initialize
- Enemy.Enemy.IntercomUI_IntercomJoinDialog_Open
- MapPin.EditMarker
- Enemy.Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionSelChanged
- DaemonAssist.DaemonAssist.PopulateBindingCombos
- Enemy.Enemy.GroupsUI_EffectFilterDialog_Open
- Enemy.Enemy.CombatLogUI_StatsWindow_Open
- Enemy.Enemy.UnitFramesUI_ConfigDialog_Import
- Enemy.Enemy.IntercomUI_IntercomJoinDialog_AddGroup


## Binding Resolution

- Total handler declarations: 409
- Resolved to Lua functions: 409 (100%)

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BankArkel
- BuffHead
- Busted
- Cheeseboard
- Crafting Info Tooltip
- DaemonAssist
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- Killer
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- Pocket Palette
- PotionBar
- RVMOD_Manager
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAComboAttackBind -> ComboBox APAComboAttackBind
- AdvancedPetAssist: APAComboAutoReattack -> ComboBox APAComboAutoReattack
- AdvancedPetAssist: APAComboAutoReattackDelay -> ComboBox APAComboAutoReattackDelay
- AdvancedPetAssist: APAComboCastDelay -> ComboBox APAComboCastDelay
- AdvancedPetAssist: APAComboCastOnAcquire -> ComboBox APAComboCastOnAcquire
- AdvancedPetAssist: APAComboCombatExitDelay -> ComboBox APAComboCombatExitDelay

## Related APIs

- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [EA_Window_WorldMap.SetMap](../../globals/functions/global_EA_Window_WorldMap.SetMap.md) (HIGH 100/100) - Global Function

## Used With

- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnSelChanged](../../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event
- [OnSelChanged](../handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none
