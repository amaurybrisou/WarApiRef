# Window

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin, Clock |
| Namespaces detected | Window |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: CG_ItemRackEQShow1, CM_ClosetGoblin: CG_ItemRackEQShow1Equipment, CM_ClosetGoblin: ClosetGoblinActionBarPageSelector, CM_ClosetGoblin: ClosetGoblinCharacterWindow, CM_ClosetGoblin: ClosetGoblinCharacterWindowBackground, CM_ClosetGoblin: ClosetGoblinCharacterWindowContents |
| XML usage count | 25 |
| XML attribute usage count | 25 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
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

Window is the primary XML container element for addon UI frames. It commonly hosts named child elements and binds lifecycle or input events to Lua handlers.

## Common Attributes

- draganddrop
- handleinput
- inherits
- layer
- movable
- name
- popable
- savesettings

## Common Handlers

- OnHidden
- OnInitialize
- OnShown
- OnShutdown

## Common Handler Functions

- ClosetGoblinCharacterWindow.OnHide
- ClosetGoblinZoneWindow.OnHide
- ClosetGoblinCharacterWindow.OnShow
- ClosetGoblinZoneWindow.OnShow
- ClosetGoblinCharacterWindow.OnShutdown
- ClosetGoblinZoneWindow.OnShutdown
- ClosetGoblinOptionWindow.OnInitialize


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| OnHidden | lifecycle | ClosetGoblinCharacterWindow.OnHide, ClosetGoblinZoneWindow.OnHide | `` |  |
| OnShown | lifecycle | ClosetGoblinCharacterWindow.OnShow, ClosetGoblinZoneWindow.OnShow | `` |  |
| OnShutdown | lifecycle | ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown | `` |  |
| OnInitialize | lifecycle | ClosetGoblinOptionWindow.OnInitialize | `` |  |

### Per-Event Lua API Calls

**OnHidden** handlers call: `WindowUtils.OnHidden`

**OnShown** handlers call: `ButtonSetPressedFlag`, `WindowUtils.OnShown`

**OnShutdown** handlers call: `UnregisterEventHandler`

**OnInitialize** handlers call: `RegisterEventHandler`, `UnregisterEventHandler`

## Common Inherits

- ClosetGoblinActionBarPageSelector
- EA_TitleBar_Default
- EA_Window_Default
- EA_Window_DefaultFrame

## Common Parent Elements

- [Windows](element_Windows.md) — 25× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 20× (HIGH)
- [Size](element_Size.md) — 13× (HIGH)
- [Windows](element_Windows.md) — 13× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 3× (MEDIUM)

## Common Template Bases

- ClosetGoblinActionBarPageSelector
- EA_TitleBar_Default
- EA_Window_Default
- EA_Window_DefaultFrame

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 64% | EA_Window_Default, EA_Window_DefaultFrame, EA_TitleBar_Default, ClosetGoblinActionBarPageSelector |
| `layer` | optional | 56% | default, background, secondary, overlay, ... |
| `movable` | optional | 16% | true, false |
| `savesettings` | optional | 12% | true, false |
| `draganddrop` | optional | 8% | true |
| `handleinput` | optional | 4% | false |
| `popable` | optional | 4% | false |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 20 times as an unnamed child.

### [Size](element_Size.md)

Observed 13 times as an unnamed child.

### [Windows](element_Windows.md)

Observed 13 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 3 times as an unnamed child.

## Recursive Hierarchy

- Root: [Window](element_Window.md)
- [Anchors](element_Anchors.md) (structural, 20×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
- [EventHandlers](element_EventHandlers.md) (structural, 3×, MEDIUM)
  - [EventHandler](element_EventHandler.md) (structural, 65×, HIGH)
- [Size](element_Size.md) (structural, 13×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
- [Windows](element_Windows.md) (structural, 13×, HIGH)
  - [Button](element_Button.md) (named, 57×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 53×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 25×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 65×, HIGH)
    - [Size](element_Size.md) (structural, 5×, MEDIUM)
      - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 1×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 1×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 1×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 1×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 1×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 1×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 1×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 1×, HIGH)
      - [Normal](element_Normal.md) (structural, 1×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 1×, HIGH)
      - [TopCenter](element_TopCenter.md) (structural, 1×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 1×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 1×, HIGH)
    - [TextOffset](element_TextOffset.md) (structural, 1×, LOW)
    - [Windows](element_Windows.md) (structural, 1×, LOW)
      - (cycle)
  - [DynamicImage](element_DynamicImage.md) (named, 13×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 12×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 2×, MEDIUM)
      - [EventHandler](element_EventHandler.md) (structural, 65×, HIGH)
    - [Size](element_Size.md) (structural, 3×, MEDIUM)
      - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
    - [Windows](element_Windows.md) (structural, 1×, LOW)
      - (cycle)
  - [FullResizeImage](element_FullResizeImage.md) (named, 7×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 6×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
    - [Size](element_Size.md) (structural, 2×, MEDIUM)
      - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 1×, LOW)
      - [BottomRight](element_BottomRight.md) (structural, 1×, HIGH)
      - [Middle](element_Middle.md) (structural, 1×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 1×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 1×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 1×, HIGH)
      - [BottomLeft](element_BottomLeft.md) (structural, 1×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 1×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 1×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 1×, HIGH)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 1×, HIGH)
      - [MiddleRight](element_MiddleRight.md) (structural, 1×, HIGH)
      - [Normal](element_Normal.md) (structural, 1×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 1×, HIGH)
      - [TopCenter](element_TopCenter.md) (structural, 1×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 1×, HIGH)
      - [TopRight](element_TopRight.md) (structural, 1×, HIGH)
  - [Label](element_Label.md) (named, 61×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 60×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
    - [Color](element_Color.md) (structural, 4×, MEDIUM)
    - [EventHandlers](element_EventHandlers.md) (structural, 4×, MEDIUM)
      - [EventHandler](element_EventHandler.md) (structural, 65×, HIGH)
    - [Size](element_Size.md) (structural, 15×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
  - [ListBox](element_ListBox.md) (named, 2×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 2×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
    - [ListData](element_ListData.md) (structural, 2×, HIGH)
      - [ListColumns](element_ListColumns.md) (structural, 2×, HIGH)
        - [ListColumn](element_ListColumn.md) (structural, 4×, HIGH)
    - [Size](element_Size.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
  - [Window](element_Window.md) (named, 25×, HIGH)
    - (cycle)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `RegisterEventHandler` | event | 6 | OnInitialize |
| `UnregisterEventHandler` | event | 5 | OnInitialize, OnShutdown |
| `ButtonSetPressedFlag` | ui | 2 | OnShown |
| `WindowUtils.OnHidden` | ui | 2 | OnHidden |
| `WindowUtils.OnShown` | ui | 2 | OnShown |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnHidden

Confidence: HIGH

### OnInitialize

Confidence: HIGH

### OnShown

Confidence: HIGH

### OnShutdown

Confidence: HIGH

## Lua Functions Manipulating This Type

- ClosetGoblinCharacterWindow.Show
- ClosetGoblinCharacterWindow.UpdateSortButtonIcon
- ClosetGoblinCharacterWindow.UpdateHighlightOnRow
- ClosetGoblinCharacterWindow.UseItemRackToggled
- ClosetGoblinCharacterWindow.HotbarChangeToggled5
- ClosetGoblinCharacterWindow.ShowShowCloakOnly
- ClosetGoblinCharacterWindow.HideCloakOptions
- ClosetGoblinCharacterWindow.ShowHelm
- ClosetGoblinCharacterWindow.HotbarPageDownProxy
- ClosetGoblinCharacterWindow.HotbarChangeToggled1
- ClosetGoblinCharacterWindow.UpdateSlotIcon
- ClosetGoblinCharacterWindow.OnShow
- ClosetGoblinCharacterWindow.UpdateActionBarSettings
- ClosetGoblinCharacterWindow.HotbarChangeToggled4
- Clock.OnUpdate
- ClosetGoblinCharacterWindow.ShowShowHelm
- ClosetGoblinCharacterWindow.ShowCloakHeraldry
- ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly
- ClosetGoblinCharacterWindow.UpdateSlotIcons
- ClosetGoblinCharacterWindow.HotbarChangeToggled3
- ClosetGoblinZoneWindow.RefreshOption
- ClosetGoblinZoneWindow.UpdateHighlightOnRow
- Clock.Initialize
- ClosetGoblinCharacterWindow.ShowCloakOptions
- ClosetGoblinCharacterWindow.OnInitialize
- ClosetGoblinZoneWindow.OnInitialize
- ClosetGoblinCharacterWindow.HideShowHelm
- ClosetGoblinCharacterWindow.ShowShowHelmOnly
- ClosetGoblinCharacterWindow.Hide
- ClosetGoblinCharacterWindow.HotbarPageUpProxy
- ClosetGoblinCharacterWindow.HotbarChangeToggled2
- ClosetGoblinZoneWindow.Show
- ClosetGoblinCharacterWindow.UpdateSetRow
- ClosetGoblinCharacterWindow.ShowCloak
- ClosetGoblinZoneWindow.Hide


## Binding Resolution

- Total handler declarations: 7
- Resolved to Lua functions: 7 (100%)

## .mod Lifecycle: Startup Windows

This element type is instantiated as a startup window by the following .mod addon(s):

| Frame Name | Addon | Hook | Resolution | Confidence |
| --- | --- | --- | --- | --- |
| ClosetGoblinCharacterWindow | CM_ClosetGoblin | OnInitialize | exact | HIGH |
| ClosetGoblinZoneWindow | CM_ClosetGoblin | OnInitialize | exact | HIGH |
| ClosetGoblinOptionWindow | CM_ClosetGoblin | OnInitialize | exact | HIGH |
## Seen In

- CM_ClosetGoblin
- Clock

## Examples

- CM_ClosetGoblin: CG_ItemRackEQShow1 -> Window CG_ItemRackEQShow1
- CM_ClosetGoblin: CG_ItemRackEQShow1Equipment -> Window CG_ItemRackEQShow1Equipment
- CM_ClosetGoblin: ClosetGoblinActionBarPageSelector -> Window ClosetGoblinActionBarPageSelector
- CM_ClosetGoblin: ClosetGoblinCharacterWindow -> Window ClosetGoblinCharacterWindow
- CM_ClosetGoblin: ClosetGoblinCharacterWindowBackground -> Window ClosetGoblinCharacterWindowBackground
- CM_ClosetGoblin: ClosetGoblinCharacterWindowContents -> Window ClosetGoblinCharacterWindowContents

## Related APIs

- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (HIGH 98/100) - XML Element Type
- [Size](element_Size.md) (HIGH 98/100) - XML Element Type
- [EA_TitleBar_Default](../../globals/constants/constant_EA_TitleBar_Default.md) (HIGH 90/100) - Constant
- [EA_Window_Default](../../globals/constants/constant_EA_Window_Default.md) (HIGH 90/100) - Constant
- [EA_Window_DefaultFrame](../../globals/constants/constant_EA_Window_DefaultFrame.md) (HIGH 90/100) - Constant
- [EventHandlers](element_EventHandlers.md) (MEDIUM 45/100) - XML Element Type
