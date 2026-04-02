# ListBox

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
| Addons seen in | AggroMeter, Aura, BuffHead, CM_ClosetGoblin, CMap, Cheeseboard, DAoCBuff, EA_UiModWindow |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.xml:131`, `/workspace_addons/Aura/Source/AuraSettings.xml:189`, `/workspace_addons/Aura/Source/AuraShares.xml:224`, `/workspace_addons/Aura/Source/AuraTexture.xml:175`, `/workspace_addons/BuffHead/Setup/General.xml:51`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompression.xml:106`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItem.xml:182`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainers.xml:106` |
| Namespaces detected | ListBox |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AggroMeter: AggroMeterGrayListBox, Aura: AuraSettingsAuraList, Aura: AuraSharesAuraList, Aura: AuraTextureIconsIcons, BuffHead: BuffHeadSetupAdvancedCompressionItemWindowList, BuffHead: BuffHeadSetupAdvancedCompressionWindowList |
| XML usage count | 51 |
| XML attribute usage count | 51 |
| Lua usage count | 4 |
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

Observed XML element type instantiated by 20 addons.

## Common Attributes

- name
- rowdef
- rowspacing
- visiblerows
- scrollbar
- rowcount
- layer
- color
- draganddrop
- scrollbarPosition

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)

## Common Handler Functions

- AggroMeter.PickedListMenu
- Enemy.KillSpamUI_KillSpamDialog_OnMouseOver
- Enemy.KillSpamUI_KillSpamDialog_OnMouseOverEnd
- Enemy.UnitFramesUI_ConfigDialog_OnClickCastingsListLButtonUp
- Enemy.UnitFramesUI_ConfigDialog_OnEffectsIndicatorsListLButtonUp
- Enemy.UnitFramesUI_ConfigDialog_OnUnitFramePartsListLButtonUp
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListLButtonUp
- TidyRoll.CustomAutoRoll.OnListLbuttonUp
- wbLeadHelperMessagesTab.OnListLButtonUp


## XML Event Bindings

| Event | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | Enemy.UnitFramesUI_ConfigDialog_OnClickCastingsListLButtonUp, Enemy.UnitFramesUI_ConfigDialog_OnEffectsIndicatorsListLButtonUp, Enemy.UnitFramesUI_ConfigDialog_OnUnitFramePartsListLButtonUp, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListLButtonUp, TidyRoll.CustomAutoRoll.OnListLbuttonUp, wbLeadHelperMessagesTab.OnListLButtonUp | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | Enemy.KillSpamUI_KillSpamDialog_OnMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | Enemy.KillSpamUI_KillSpamDialog_OnMouseOverEnd | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | AggroMeter.PickedListMenu | `function(...)` | LOW |

## Common Inherits

- none

## Common Parent Elements

- [Window](element_Window.md)

## Common Structural Child Elements

- [ListData](element_ListData.md)
- [ListColumns](element_ListColumns.md)
- [ListColumn](element_ListColumn.md)

## Typical XML Structure

```xml
<ListBox name="..." color="155, 255, 155, 0" scrollbar="EA_ScrollBar_DefaultVerticalC..." rowdef="AggroMeterGrayTemplateListboxRow" visiblerows="13" rowcount="120" rowspacing="0">
  <ListData table="AggroMeter.Listdata">
    <ListColumns>
      <ListColumn windowname="Name" variable="Name" format="wstring"/>
      <ListColumn windowname="Rank" variable="RankN" format="wstring"/>
    </ListColumns>
  </ListData>
</ListBox>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `rowdef` | **required** | 100% | DAoCBuff_FrameSettingsRow, EnemyChooseIconDialogList_RowTemplate, EA_Button_UiModuleListRowTemplate, BuffHeadSetupFilterRowTemplate, ... |
| `rowspacing` | **required** | 100% | 1, 0, 2, 6, ... |
| `visiblerows` | **required** | 100% | 12, 4, 10, 5, ... |
| `scrollbar` | **required** | 84% | EA_ScrollBar_DefaultVerticalChain |
| `rowcount` | optional | 23% | 120, 100, 500, 150, ... |
| `layer` | optional | 7% | secondary |
| `color` | optional | 1% | 155, 255, 155, 0 |
| `draganddrop` | optional | 1% | true |
| `scrollbarPosition` | optional | 1% | left |
## Structural Sub-Elements

### [ListData](element_ListData.md)

Observed 51 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `table` | **required** |  |
| `populationfunction` | **required** |  |
### [ListColumns](element_ListColumns.md)

Observed 28 times as an unnamed child.

### [ListColumn](element_ListColumn.md)

Observed 27 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `format` | **required** |  |
| `variable` | **required** |  |
| `windowname` | **required** |  |
## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Call Count | From Events |
| --- | --- | --- |
| `WindowGetId` | 6 | OnLButtonUp, OnRButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

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
### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnRButtonUp

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
### OnLButtonUp

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
## Lua Functions Manipulating This Type

- CM_ClosetGoblin.ClosetGoblinCharacterWindow.UpdateHighlightOnRow
- CM_ClosetGoblin.ClosetGoblinZoneWindow.UpdateHighlightOnRow

## Seen In

- AggroMeter
- Aura
- BuffHead
- CM_ClosetGoblin
- CMap
- Cheeseboard
- DAoCBuff
- EA_UiModWindow
- Enemy
- MapMonster
- MapPin
- Pocket Palette
- QuickTacticSwitch
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- Tortall_DPS
- TurretRange
- wbLeadHelper

## Examples

- AggroMeter: AggroMeterGrayListBox -> ListBox AggroMeterGrayListBox
- Aura: AuraSettingsAuraList -> ListBox AuraSettingsAuraList
- Aura: AuraSharesAuraList -> ListBox AuraSharesAuraList
- Aura: AuraTextureIconsIcons -> ListBox AuraTextureIconsIcons
- BuffHead: BuffHeadSetupAdvancedCompressionItemWindowList -> ListBox BuffHeadSetupAdvancedCompressionItemWindowList
- BuffHead: BuffHeadSetupAdvancedCompressionWindowList -> ListBox BuffHeadSetupAdvancedCompressionWindowList

## Related APIs

- [EA_BackpackUtilsMediator.GetCurrentBackpackType](../../globals/functions/global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](../../globals/functions/global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetItemsFromBackpack](../../globals/functions/global_EA_BackpackUtilsMediator.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function

## Used With

- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none
