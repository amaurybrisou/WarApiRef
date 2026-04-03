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
| Addons seen in | AggroMeter, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy, Pocket Palette, Shinies |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTexture.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainers.xml:0` |
| Namespaces detected | ListBox |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AggroMeter: AggroMeterGrayListBox, Aura: AuraSettingsAuraList, Aura: AuraSharesAuraList, Aura: AuraTextureIconsIcons, BuffHead: BuffHeadSetupAdvancedCompressionItemWindowList, BuffHead: BuffHeadSetupAdvancedCompressionWindowList |
| XML usage count | 42 |
| XML attribute usage count | 42 |
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

Observed XML element type instantiated by 12 addons.

## Common Attributes

- color
- draganddrop
- layer
- name
- rowcount
- rowdef
- rowspacing
- scrollbar
- scrollbarPosition
- visiblerows

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


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | Enemy.UnitFramesUI_ConfigDialog_OnClickCastingsListLButtonUp, Enemy.UnitFramesUI_ConfigDialog_OnEffectsIndicatorsListLButtonUp, Enemy.UnitFramesUI_ConfigDialog_OnUnitFramePartsListLButtonUp, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListLButtonUp, TidyRoll.CustomAutoRoll.OnListLbuttonUp | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.KillSpamUI_KillSpamDialog_OnMouseOver | `function()` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | Enemy.KillSpamUI_KillSpamDialog_OnMouseOverEnd | `function(...)` | LOW |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | AggroMeter.PickedListMenu | `flags, x, y` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | , Enemy.UnitFramesUI_ConfigDialog_OnClickCastingsListLButtonUp, Enemy.UnitFramesUI_ConfigDialog_OnEffectsIndicatorsListLButtonUp, Enemy.UnitFramesUI_ConfigDialog_OnUnitFramePartsListLButtonUp, Enemy.UnitFramesUI_EffectsIndicatorDialog_OnEffectFiltersListLButtonUp, TidyRoll.CustomAutoRoll.OnListLbuttonUp | `flags, x, y` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.KillSpamUI_KillSpamDialog_OnMouseOver | `` |  |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | Enemy.KillSpamUI_KillSpamDialog_OnMouseOverEnd | `` |  |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | AggroMeter.PickedListMenu | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonUp** handlers call: `Cursor.Clear`, `Cursor.IconOnCursor`, `ListBoxGetDataIndex`, `WindowGetId`

**OnRButtonUp** handlers call: `ListBoxGetDataIndex`, `WindowGetId`

**OnLButtonUp** handlers call: `Cursor.Clear`, `Cursor.IconOnCursor`, `ListBoxGetDataIndex`, `WindowGetId`

**OnRButtonUp** handlers call: `ListBoxGetDataIndex`, `WindowGetId`

## Common Inherits

- none

## Common Parent Elements

- [Windows](element_Windows.md) — 42× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 42× (HIGH)
- [ListData](element_ListData.md) — 42× (HIGH)
- [Size](element_Size.md) — 31× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 10× (HIGH)


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<ListBox name="..." rowdef="AuraSharesRow" rowspacing="0" scrollbar="EA_ScrollBar_DefaultVerticalC..." visiblerows="5">
  <Size/>
  <ListData populationfunction="AuraShares.PopulateDisplay" table="AuraShares.listDisplayData">
    <ListColumns>
      <ListColumn format="wstring" variable="Character" windowname="Character"/>
      <ListColumn format="wstring" variable="Type" windowname="Type"/>
      <ListColumn format="wstring" variable="Name" windowname="Name"/>
    </ListColumns>
  </ListData>
</ListBox>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `rowdef` | **required** | 100% | AggroMeterGrayTemplateListboxRow, AuraSharesRow, AuraIconRow, AuraWindowRow, ... |
| `rowspacing` | **required** | 100% | 0, 1 |
| `visiblerows` | **required** | 100% | 13, 5, 4, 7, ... |
| `scrollbar` | **required** | 88% | EA_ScrollBar_DefaultVerticalChain |
| `rowcount` | optional | 23% | 120, 100, 50, 250, ... |
| `layer` | optional | 14% | secondary |
| `color` | optional | 2% | 155, 255, 155, 0 |
| `draganddrop` | optional | 2% | true |
| `scrollbarPosition` | optional | 2% | left |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 42 times as an unnamed child.

### [ListData](element_ListData.md)

Observed 42 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `table` | **required** | AggroMeter.Listdata, AuraShares.listDisplayData, AuraTexture.listIconDisplayData, AuraSettings.listDisplayData |
| `populationfunction` | **required** | AuraShares.PopulateDisplay, AuraTexture.PopulateIconsListDisplay, AuraSettings.PopulateDisplay, BuffHead.Setup.OnPopulate |
### [Size](element_Size.md)

Observed 31 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 10 times as an unnamed child.

## Recursive Hierarchy

- Root: [ListBox](element_ListBox.md)
- [Anchors](element_Anchors.md) (structural, 42×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
      - (cycle)
- [EventHandlers](element_EventHandlers.md) (structural, 10×, HIGH)
  - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
- [ListData](element_ListData.md) (structural, 42×, HIGH)
  - [ListColumns](element_ListColumns.md) (structural, 25×, HIGH)
    - [ListColumn](element_ListColumn.md) (structural, 42×, HIGH)
- [Size](element_Size.md) (structural, 31×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `ListBoxGetDataIndex` | ui | 5 | OnLButtonUp, OnRButtonUp |
| `WindowGetId` | ui | 5 | OnLButtonUp, OnRButtonUp |
| `Cursor.Clear` | ui | 1 | OnLButtonUp |
| `Cursor.IconOnCursor` | data | 1 | OnLButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMouseOver

Confidence: LOW

### OnMouseOverEnd

Confidence: LOW

### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
## Lua Functions Manipulating This Type

- DAoCBuffSettings.Change_Setting
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- ClosetGoblinCharacterWindow.UpdateHighlightOnRow
- ClosetGoblinZoneWindow.UpdateHighlightOnRow
- DAoCBuffSettings.CreateOptionswindow
- DAoCBuffSettings.SetLabels


## Binding Resolution

- Total handler declarations: 11
- Resolved to Lua functions: 9 (81%)

## Seen In

- AggroMeter
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Pocket Palette
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange

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
- [ListData](element_ListData.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 55/100) - XML Element Type

## Used With

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event

## Affects

- none
