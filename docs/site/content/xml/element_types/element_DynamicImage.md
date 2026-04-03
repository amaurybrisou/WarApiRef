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

- alpha
- filtering
- handleinput
- id
- inherits
- layer
- name
- popable
- slice
- texture
- textureScale
- texturescale

## Common Handlers

- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)

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
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | DAoCBuffSettings.ShowTooltip, Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip, AggroMeter.OnMouseOverStart, BagOMatic.wnd_on_mouse_over, Enemy.Guard_GuardIndicator_OnMouseOver, Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOver, Enemy.TalismanAlerter_TalismanAlerterIndicator_OnMouseOver, Enemy.UI_Icon_OnMouseOver, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionMouseOver | `` |  |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | BuffHead.Setup.Container.OnContainerClick, ClosetGoblinOptionWindow.OnLButtonUp, BagOMatic.wnd_on_lbutton_up, Enemy.Guard_GuardIndicator_OnLButtonUp, Enemy.MarksUI_EnemyMarksWindow_OnAddLButtonUp, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick, Enemy.UI_Icon_OnLButtonUp, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionLButtonUp, TurretRange.Setup.Display.OnCircleColorLUp, TurretRange.Setup.Display.OnDistanceColorLUp, TurretRange.Setup.Display.OnGraphicColorLUp | `flags, x, y` | MEDIUM |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | BuffHead.Setup.AdvancedContainersItem.OnContainerRClick, ClosetGoblinOptionWindow.OnRButtonUp, Enemy.MarksUI_EnemyMarksWindow_OnAddRButtonUp, Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick, Enemy.UI_Icon_OnRButtonUp, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionRButtonUp | `flags, x, y` | MEDIUM |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | Enemy.CombatLogUI_IDS_OnRowLButtonDown | `flags, x, y` | MEDIUM |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | Enemy.MarksUI_EnemyMarksWindow_OnAddMouseOverEnd | `` |  |

### Per-Event Lua API Calls

**OnLButtonDown** handlers call: `SystemData.ActiveWindow.name:match`

**OnLButtonUp** handlers call: `ComboBoxSetSelectedMenuItem`, `Cursor.Clear`, `Cursor.IconOnCursor`, `Player.GetAbilityData`, `WindowGetId`, `WindowSetShowing`

**OnMouseOver** handlers call: `WindowGetId`, `WindowGetParent`

**OnRButtonUp** handlers call: `WindowGetId`

**OnMouseOver** handlers call: `WindowGetId`, `WindowGetParent`

**OnLButtonUp** handlers call: `ComboBoxSetSelectedMenuItem`, `Cursor.Clear`, `Cursor.IconOnCursor`, `Player.GetAbilityData`, `WindowGetId`, `WindowSetShowing`

**OnRButtonUp** handlers call: `WindowGetId`

**OnLButtonDown** handlers call: `SystemData.ActiveWindow.name:match`

## Common Inherits

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

## Common Parent Elements

- [Windows](element_Windows.md) — 237× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 216× (HIGH)
- [Size](element_Size.md) — 190× (HIGH)
- [TexDims](element_TexDims.md) — 65× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 40× (HIGH)
- [TintColor](element_TintColor.md) — 25× (HIGH)
- [TexCoords](element_TexCoords.md) — 14× (HIGH)
- [Windows](element_Windows.md) — 6× (MEDIUM)
- [Anchor](element_Anchor.md) — 1× (LOW)
- [Color](element_Color.md) — 1× (LOW)

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

### [Anchors](element_Anchors.md)

Observed 216 times as an unnamed child.

### [Size](element_Size.md)

Observed 190 times as an unnamed child.

### [TexDims](element_TexDims.md)

Observed 65 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 64, 256, 32, 430 |
| `y` | **required** | 64, 256, 32, 430 |
### [EventHandlers](element_EventHandlers.md)

Observed 40 times as an unnamed child.

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

### [Anchor](element_Anchor.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `point` | **required** | center, topleft, bottomright, right |
| `relativePoint` | **required** | center, topleft, bottomright, left |
| `relativeTo` | **required** | Root, $parent, $parentGeneral, $parentFollowTarget |
| `xOffset` | optional | 0 |
| `yOffset` | optional | 0 |
### [Color](element_Color.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 73, 225, 255, 55 |
| `g` | **required** | 218, 255, 155, 55 |
| `r` | **required** | 255, 155, 245, 175 |
| `a` | optional | 255, 0, 0.5, 0.8 |
## Recursive Hierarchy

- Root: [DynamicImage](element_DynamicImage.md)
- [Anchor](element_Anchor.md) (structural, 1×, LOW)
  - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
    - (cycle)
- [Anchors](element_Anchors.md) (structural, 216×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
      - (cycle)
- [Color](element_Color.md) (structural, 1×, LOW)
- [EventHandlers](element_EventHandlers.md) (structural, 40×, HIGH)
  - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
- [Size](element_Size.md) (structural, 190×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
- [TexCoords](element_TexCoords.md) (structural, 14×, HIGH)
  - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
  - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
  - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
  - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
  - [Left](element_Left.md) (structural, 7×, MEDIUM)
  - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
  - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
  - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
  - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
  - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
  - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
  - [Right](element_Right.md) (structural, 7×, MEDIUM)
  - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
  - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
  - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
- [TexDims](element_TexDims.md) (structural, 65×, HIGH)
- [TintColor](element_TintColor.md) (structural, 25×, HIGH)
- [Windows](element_Windows.md) (structural, 6×, MEDIUM)
  - [AnimatedImage](element_AnimatedImage.md) (named, 12×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 12×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [AnimFrames](element_AnimFrames.md) (structural, 2×, MEDIUM)
      - [AnimFrame](element_AnimFrame.md) (structural, 14×, HIGH)
    - [Size](element_Size.md) (structural, 9×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Windows](element_Windows.md) (structural, 2×, MEDIUM)
      - (cycle)
  - [Button](element_Button.md) (named, 664×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 2×, LOW)
      - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 556×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [AnimatedImages](element_AnimatedImages.md) (structural, 1×, LOW)
      - [Normal](element_Normal.md) (structural, 1×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 418×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [OverlayOffset](element_OverlayOffset.md) (structural, 7×, MEDIUM)
    - [OverlaySize](element_OverlaySize.md) (structural, 7×, MEDIUM)
    - [OverlayTexCoords](element_OverlayTexCoords.md) (structural, 7×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 7×, HIGH)
      - [Normal](element_Normal.md) (structural, 7×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 7×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 7×, HIGH)
    - [ResizeImages](element_ResizeImages.md) (structural, 14×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 8×, HIGH)
      - [Normal](element_Normal.md) (structural, 10×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 14×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 10×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 9×, HIGH)
    - [Size](element_Size.md) (structural, 259×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 7×, MEDIUM)
      - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
      - [Left](element_Left.md) (structural, 7×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
      - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
      - [Right](element_Right.md) (structural, 7×, MEDIUM)
      - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
    - [TexDims](element_TexDims.md) (structural, 4×, MEDIUM)
    - [TexSlices](element_TexSlices.md) (structural, 22×, HIGH)
      - [BottomCenter](element_BottomCenter.md) (structural, 5×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 5×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 5×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 11×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 11×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 5×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 5×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 5×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 22×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 16×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 16×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 16×, HIGH)
      - [TopCenter](element_TopCenter.md) (structural, 5×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 5×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 5×, MEDIUM)
    - [TextColors](element_TextColors.md) (structural, 16×, HIGH)
      - [Disabled](element_Disabled.md) (structural, 16×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 4×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 16×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 16×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 16×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 16×, HIGH)
    - [TextOffset](element_TextOffset.md) (structural, 18×, HIGH)
    - [Windows](element_Windows.md) (structural, 18×, HIGH)
      - (cycle)
  - [CircleImage](element_CircleImage.md) (named, 23×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 22×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 22×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 20×, HIGH)
      - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
      - [Left](element_Left.md) (structural, 7×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
      - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
      - [Right](element_Right.md) (structural, 7×, MEDIUM)
      - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
    - [TexDims](element_TexDims.md) (structural, 2×, MEDIUM)
    - [TintColor](element_TintColor.md) (structural, 6×, HIGH)
    - [Windows](element_Windows.md) (structural, 1×, LOW)
      - (cycle)
  - [ColorPicker](element_ColorPicker.md) (named, 1×, LOW)
    - [Anchors](element_Anchors.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [ColorSize](element_ColorSize.md) (structural, 1×, HIGH)
    - [ColorSpacing](element_ColorSpacing.md) (structural, 1×, HIGH)
    - [ColorTexCoords](element_ColorTexCoords.md) (structural, 1×, HIGH)
    - [ColorTexDims](element_ColorTexDims.md) (structural, 1×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [ComboBox](element_ComboBox.md) (named, 233×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 228×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 179×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [MenuButtonOffset](element_MenuButtonOffset.md) (structural, 7×, MEDIUM)
    - [Size](element_Size.md) (structural, 52×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [CooldownDisplay](element_CooldownDisplay.md) (named, 1×, LOW)
    - [Anchors](element_Anchors.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [DynamicImage](element_DynamicImage.md) (named, 237×, HIGH)
    - (cycle)
  - [EditBox](element_EditBox.md) (named, 151×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 126×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 102×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 127×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [TextOffset](element_TextOffset.md) (structural, 15×, HIGH)
  - [FullResizeImage](element_FullResizeImage.md) (named, 156×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 139×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 1×, LOW)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 35×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 14×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 14×, HIGH)
      - [Middle](element_Middle.md) (structural, 14×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 14×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 9×, MEDIUM)
      - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
      - [Left](element_Left.md) (structural, 7×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
      - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
      - [Right](element_Right.md) (structural, 7×, MEDIUM)
      - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
    - [TexDims](element_TexDims.md) (structural, 3×, MEDIUM)
    - [TexSlices](element_TexSlices.md) (structural, 5×, MEDIUM)
      - [BottomCenter](element_BottomCenter.md) (structural, 5×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 5×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 5×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 11×, HIGH)
      - [DisabledPressed](element_DisabledPressed.md) (structural, 11×, HIGH)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 5×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 5×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 5×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 22×, HIGH)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 16×, HIGH)
      - [Pressed](element_Pressed.md) (structural, 16×, HIGH)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 16×, HIGH)
      - [TopCenter](element_TopCenter.md) (structural, 5×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 5×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 5×, MEDIUM)
    - [TintColor](element_TintColor.md) (structural, 28×, HIGH)
    - [Windows](element_Windows.md) (structural, 1×, LOW)
      - (cycle)
  - [HorizontalResizeImage](element_HorizontalResizeImage.md) (named, 19×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 11×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 4×, MEDIUM)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Sizes](element_Sizes.md) (structural, 8×, HIGH)
      - [BottomRight](element_BottomRight.md) (structural, 14×, HIGH)
      - [Middle](element_Middle.md) (structural, 14×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 14×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 8×, HIGH)
      - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
      - [Left](element_Left.md) (structural, 7×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
      - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
      - [Right](element_Right.md) (structural, 7×, MEDIUM)
      - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
    - [TintColor](element_TintColor.md) (structural, 7×, HIGH)
  - [Label](element_Label.md) (named, 1243×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 1238×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Color](element_Color.md) (structural, 375×, HIGH)
    - [EventHandlers](element_EventHandlers.md) (structural, 95×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 1168×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Text](element_Text.md) (structural, 60×, HIGH)
    - [TintColor](element_TintColor.md) (structural, 14×, HIGH)
    - [Windows](element_Windows.md) (structural, 1×, LOW)
      - (cycle)
  - [ListBox](element_ListBox.md) (named, 42×, HIGH)
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
  - [LogDisplay](element_LogDisplay.md) (named, 1×, LOW)
    - [Anchors](element_Anchors.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Windows](element_Windows.md) (structural, 1×, HIGH)
      - (cycle)
  - [MapDisplay](element_MapDisplay.md) (named, 1×, LOW)
    - [Anchors](element_Anchors.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 1×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [ScrollWindow](element_ScrollWindow.md) (named, 26×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 16×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [Size](element_Size.md) (structural, 12×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Windows](element_Windows.md) (structural, 24×, HIGH)
      - (cycle)
  - [SliderBar](element_SliderBar.md) (named, 83×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 83×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 80×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 76×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [StatusBar](element_StatusBar.md) (named, 9×, MEDIUM)
    - [Anchors](element_Anchors.md) (structural, 9×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
  - [VerticalResizeImage](element_VerticalResizeImage.md) (named, 4×, MEDIUM)
    - [Sizes](element_Sizes.md) (structural, 1×, LOW)
      - [BottomRight](element_BottomRight.md) (structural, 14×, HIGH)
      - [Middle](element_Middle.md) (structural, 14×, HIGH)
      - [TopLeft](element_TopLeft.md) (structural, 14×, HIGH)
    - [TexCoords](element_TexCoords.md) (structural, 1×, LOW)
      - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
      - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
      - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
      - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
      - [Left](element_Left.md) (structural, 7×, MEDIUM)
      - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
      - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
      - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
      - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
      - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
      - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
      - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
      - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
      - [Right](element_Right.md) (structural, 7×, MEDIUM)
      - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
      - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
      - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
    - [TintColor](element_TintColor.md) (structural, 3×, HIGH)
  - [VerticalScrollbar](element_VerticalScrollbar.md) (named, 25×, HIGH)
    - [Anchors](element_Anchors.md) (structural, 25×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 23×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
  - [Window](element_Window.md) (named, 1001×, HIGH)
    - [Button](element_Button.md) (named, 8×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 2×, LOW)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
      - [Anchors](element_Anchors.md) (structural, 556×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
            - (cycle)
      - [AnimatedImages](element_AnimatedImages.md) (structural, 1×, LOW)
        - [Normal](element_Normal.md) (structural, 1×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)
      - [EventHandlers](element_EventHandlers.md) (structural, 418×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
      - [OverlayOffset](element_OverlayOffset.md) (structural, 7×, MEDIUM)
      - [OverlaySize](element_OverlaySize.md) (structural, 7×, MEDIUM)
      - [OverlayTexCoords](element_OverlayTexCoords.md) (structural, 7×, MEDIUM)
        - [Disabled](element_Disabled.md) (structural, 7×, HIGH)
        - [Normal](element_Normal.md) (structural, 7×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 7×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 7×, HIGH)
      - [ResizeImages](element_ResizeImages.md) (structural, 14×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 8×, HIGH)
        - [Normal](element_Normal.md) (structural, 10×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 14×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 10×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 9×, HIGH)
      - [Size](element_Size.md) (structural, 259×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
      - [TexCoords](element_TexCoords.md) (structural, 7×, MEDIUM)
        - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
        - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
        - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
        - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
        - [Left](element_Left.md) (structural, 7×, MEDIUM)
        - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
        - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
        - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
        - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
        - [Right](element_Right.md) (structural, 7×, MEDIUM)
        - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
        - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
        - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
      - [TexDims](element_TexDims.md) (structural, 4×, MEDIUM)
      - [TexSlices](element_TexSlices.md) (structural, 22×, HIGH)
        - [BottomCenter](element_BottomCenter.md) (structural, 5×, MEDIUM)
        - [BottomLeft](element_BottomLeft.md) (structural, 5×, MEDIUM)
        - [BottomRight](element_BottomRight.md) (structural, 5×, MEDIUM)
        - [Disabled](element_Disabled.md) (structural, 11×, HIGH)
        - [DisabledPressed](element_DisabledPressed.md) (structural, 11×, HIGH)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 5×, MEDIUM)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 5×, MEDIUM)
        - [MiddleRight](element_MiddleRight.md) (structural, 5×, MEDIUM)
        - [Normal](element_Normal.md) (structural, 22×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 16×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 16×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 16×, HIGH)
        - [TopCenter](element_TopCenter.md) (structural, 5×, MEDIUM)
        - [TopLeft](element_TopLeft.md) (structural, 5×, MEDIUM)
        - [TopRight](element_TopRight.md) (structural, 5×, MEDIUM)
      - [TextColors](element_TextColors.md) (structural, 16×, HIGH)
        - [Disabled](element_Disabled.md) (structural, 16×, HIGH)
        - [DisabledPressed](element_DisabledPressed.md) (structural, 4×, MEDIUM)
        - [Normal](element_Normal.md) (structural, 16×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 16×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 16×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 16×, HIGH)
      - [TextOffset](element_TextOffset.md) (structural, 18×, HIGH)
      - [Windows](element_Windows.md) (structural, 18×, HIGH)
        - (cycle)
    - [ComboBox](element_ComboBox.md) (named, 4×, MEDIUM)
      - [Anchors](element_Anchors.md) (structural, 228×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
            - (cycle)
      - [EventHandlers](element_EventHandlers.md) (structural, 179×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
      - [MenuButtonOffset](element_MenuButtonOffset.md) (structural, 7×, MEDIUM)
      - [Size](element_Size.md) (structural, 52×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [FullResizeImage](element_FullResizeImage.md) (named, 2×, LOW)
      - [Anchors](element_Anchors.md) (structural, 139×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
            - (cycle)
      - [Color](element_Color.md) (structural, 1×, LOW)
      - [EventHandlers](element_EventHandlers.md) (structural, 1×, LOW)
        - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
      - [Size](element_Size.md) (structural, 35×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
      - [Sizes](element_Sizes.md) (structural, 14×, HIGH)
        - [BottomRight](element_BottomRight.md) (structural, 14×, HIGH)
        - [Middle](element_Middle.md) (structural, 14×, HIGH)
        - [TopLeft](element_TopLeft.md) (structural, 14×, HIGH)
      - [TexCoords](element_TexCoords.md) (structural, 9×, MEDIUM)
        - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
        - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
        - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
        - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
        - [Left](element_Left.md) (structural, 7×, MEDIUM)
        - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
        - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
        - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
        - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
        - [Right](element_Right.md) (structural, 7×, MEDIUM)
        - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
        - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
        - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
      - [TexDims](element_TexDims.md) (structural, 3×, MEDIUM)
      - [TexSlices](element_TexSlices.md) (structural, 5×, MEDIUM)
        - [BottomCenter](element_BottomCenter.md) (structural, 5×, MEDIUM)
        - [BottomLeft](element_BottomLeft.md) (structural, 5×, MEDIUM)
        - [BottomRight](element_BottomRight.md) (structural, 5×, MEDIUM)
        - [Disabled](element_Disabled.md) (structural, 11×, HIGH)
        - [DisabledPressed](element_DisabledPressed.md) (structural, 11×, HIGH)
        - [MiddleCenter](element_MiddleCenter.md) (structural, 5×, MEDIUM)
        - [MiddleLeft](element_MiddleLeft.md) (structural, 5×, MEDIUM)
        - [MiddleRight](element_MiddleRight.md) (structural, 5×, MEDIUM)
        - [Normal](element_Normal.md) (structural, 22×, HIGH)
        - [NormalHighlit](element_NormalHighlit.md) (structural, 16×, HIGH)
        - [Pressed](element_Pressed.md) (structural, 16×, HIGH)
        - [PressedHighlit](element_PressedHighlit.md) (structural, 16×, HIGH)
        - [TopCenter](element_TopCenter.md) (structural, 5×, MEDIUM)
        - [TopLeft](element_TopLeft.md) (structural, 5×, MEDIUM)
        - [TopRight](element_TopRight.md) (structural, 5×, MEDIUM)
      - [TintColor](element_TintColor.md) (structural, 28×, HIGH)
      - [Windows](element_Windows.md) (structural, 1×, LOW)
        - (cycle)
    - [Label](element_Label.md) (named, 15×, HIGH)
      - [Anchors](element_Anchors.md) (structural, 1238×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
            - (cycle)
      - [Color](element_Color.md) (structural, 375×, HIGH)
      - [EventHandlers](element_EventHandlers.md) (structural, 95×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
      - [Size](element_Size.md) (structural, 1168×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
      - [Text](element_Text.md) (structural, 60×, HIGH)
      - [TintColor](element_TintColor.md) (structural, 14×, HIGH)
      - [Windows](element_Windows.md) (structural, 1×, LOW)
        - (cycle)
    - [SliderBar](element_SliderBar.md) (named, 2×, LOW)
      - [Anchors](element_Anchors.md) (structural, 83×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
        - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
          - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
          - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
            - (cycle)
      - [EventHandlers](element_EventHandlers.md) (structural, 80×, HIGH)
        - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
      - [Size](element_Size.md) (structural, 76×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Window](element_Window.md) (named, 5×, MEDIUM)
      - (cycle)
    - [Anchor](element_Anchor.md) (structural, 10×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
      - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
        - (cycle)
    - [Anchors](element_Anchors.md) (structural, 745×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
      - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
        - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
        - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
          - (cycle)
    - [EventHandlers](element_EventHandlers.md) (structural, 280×, HIGH)
      - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
    - [Size](element_Size.md) (structural, 624×, HIGH)
      - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
    - [Visual](element_Visual.md) (structural, 1×, LOW)
      - [Color](element_Color.md) (structural, 1×, HIGH)
    - [Windows](element_Windows.md) (structural, 476×, HIGH)
      - (cycle)

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

- TurretRange.OnUpdate
- Enemy.GuardInitialize
- RoR_SoR.OnCombat
- SwiftAssist.Initialize
- Swift Assist.SetSmartLabel
- GuardLine.update
- RoR_SoR.OnInitialize
- Enemy.Guard_GuardIndicator_Update
- PP.UpdateListRow
- Swift Assist.local.SetTexLabel
- Enemy.EnemyEffectsIndicator:Update
- Enemy.AssistUI_Target_Show
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- Swift Assist.local.SetSmartLabel
- BankArkel.PackImg
- Enemy.MarksInitialize
- Enemy.UI_Icon_Switch
- ClosetGoblinCharacterWindow.OnInitialize
- Enemy.Guard_OnSettingsChanged
- Enemy.TalismanAlerterInitialize
- RoR_SoR.OnScenario
- SwiftAssist.OnMacroUpdated
- WSCT.ColorAcceptButtonOnButtonUp
- BagOMatic.init
- ClosetGoblinCharacterWindow.UpdateSortButtonIcon
- WSCT.ColorHideMenu
- Enemy._Initialize
- WSCT.ColorOnInitialize
- Enemy.TalismanAlerter_Update
- Enemy.TalismanAlerter_OnSettingsChanged
- Swift Assist.SetTexLabel
- WSCT.OnLButtonUpColorPicker
- ClosetGoblinZoneWindow.OnInitialize
- ClosetGoblin.Initialize


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
- [EA_ListSortDownArrow](../../globals/constants/constant_EA_ListSortDownArrow.md) (HIGH 100/100) - Constant
- [EA_ListSortUpArrow](../../globals/constants/constant_EA_ListSortUpArrow.md) (HIGH 100/100) - Constant
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TexDims](element_TexDims.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EA_Default_CharacterImage](../../globals/constants/constant_EA_Default_CharacterImage.md) (HIGH 90/100) - Constant
- [EA_Default_TrainingImage](../../globals/constants/constant_EA_Default_TrainingImage.md) (HIGH 90/100) - Constant
- [EA_Image_DefaultIcon](../../globals/constants/constant_EA_Image_DefaultIcon.md) (HIGH 90/100) - Constant
- [EA_Image_DefaultIconFrame](../../globals/constants/constant_EA_Image_DefaultIconFrame.md) (HIGH 90/100) - Constant
- [EA_Templates_BrassCoin](../../globals/constants/constant_EA_Templates_BrassCoin.md) (HIGH 90/100) - Constant
- [EA_Templates_GoldCoin](../../globals/constants/constant_EA_Templates_GoldCoin.md) (HIGH 90/100) - Constant
- [EA_Templates_SilverCoin](../../globals/constants/constant_EA_Templates_SilverCoin.md) (HIGH 90/100) - Constant
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function
- [Anchor](element_Anchor.md) (MEDIUM 55/100) - XML Element Type
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 55/100) - XML Element Type
- [Color](element_Color.md) (MEDIUM 45/100) - XML Element Type

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
