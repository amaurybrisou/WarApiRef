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

- autoresize
- font
- handleinput
- inherits
- layer
- maxchars
- name
- popable
- skipinput
- textalign
- warnOnTextCropped
- wordwrap

## Common Handlers

- [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md)
- [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md)
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)

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
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.ScenarioInfoUI_ScenarioInfoDialog_ShowTooltip, Enemy.CombatLogUI_EpsWindow_OnMouseOver, Enemy.ConfigurationWindow_ShowTooltip, Enemy.CombatLogUI_StatsWindow_ListRowMouseOver, Enemy.CombatLogUI_TargetDefenseWindow_OnMouseOver, Enemy.CombatLogUI_StatsWindow_OnEpsMouseOver, Enemy.KillSpamUI_KillSpamDialog_OnTotalMouseOver, TexturedButtons.Setup.Fonts.OnFontExampleMouseOver, AggroMeter.SelectChar, BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOver, BuffHead.Setup.Layout.Properties.OnFontExampleMouseOver, BuffHead.Setup.Layout.Properties.OnPropertyTitleMouseOver, MiracleGrowLight.onHover, TurretRange.Setup.Display.OnDistanceFontMouseOver | `` |  |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueLClick, Enemy.KillSpamUI_KillSpamDialog_OnTotalLButtonUp, ClosetGoblinCharacterWindow.OnClickSetRow, ClosetGoblinZoneWindow.OnClickZoneRow, Enemy.CombatLogUI_EpsWindow_OnInDpsLButtonUp, Enemy.CombatLogUI_EpsWindow_OnInHpsLButtonUp, Enemy.CombatLogUI_EpsWindow_OnOutDpsAoeLButtonUp, Enemy.CombatLogUI_EpsWindow_OnOutDpsLButtonUp, Enemy.CombatLogUI_EpsWindow_OnOutHpsAoeLButtonUp, Enemy.CombatLogUI_EpsWindow_OnOutHpsLButtonUp, AggroMeter.SelectChar, BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleLClick, BuffHead.Setup.Layout.Manager.OnExportLUp, BuffHead.Setup.Layout.Manager.OnImportLUp, BuffHead.Setup.Layout.Manager.OnLayoutsLUp, BuffHead.Setup.Layout.Properties.OnFontFontLUp, BuffHead.Setup.Layout.Properties.OnPropertyTitleLClick, Enemy.CombatLogUI_StatsWindow_OnEpsAoeCurClick, Enemy.CombatLogUI_StatsWindow_OnEpsAoeMaxClick, Enemy.CombatLogUI_StatsWindow_OnEpsCurClick, Enemy.CombatLogUI_StatsWindow_OnEpsMaxClick, TexturedButtons.Setup.Fonts.OnCooldownFontLUp, TexturedButtons.Setup.Fonts.OnKeybindFontLUp, TurretRange.Setup.Display.OnDistanceFontLUp | `flags, x, y` | MEDIUM |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnValueRClick, Enemy.CombatLogUI_EpsWindow_OnRButtonUp, Enemy.CombatLogUI_TargetDefenseWindow_OnRButtonUp, ClosetGoblinCharacterWindow.OnSetRowContextMenu, ClosetGoblinZoneWindow.OnSetZoneRowContextMenu, MiracleGrowLight.switchMode | `flags, x, y` | MEDIUM |
| [OnHyperLinkLButtonUp](../handlers/handler_OnHyperLinkLButtonUp.md) | custom | EA_ChatWindow.OnHyperLinkLButtonUp | `linkData, flags, x, y` | LOW |
| [OnHyperLinkRButtonUp](../handlers/handler_OnHyperLinkRButtonUp.md) | custom | EA_ChatWindow.OnHyperLinkRButtonUp | `` |  |
| [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) | input | TexturedButtons.Setup.Fonts.OnFontExampleMouseOut, BuffHead.Setup.AdvancedContainersItem.Properties.OnPropertyTitleMouseOut, BuffHead.Setup.Layout.Properties.OnFontExampleMouseOut, BuffHead.Setup.Layout.Properties.OnPropertyTitleMouseOut, TurretRange.Setup.Display.OnDistanceFontMouseOut | `` |  |
| [OnLButtonDown](../handlers/handler_OnLButtonDown.md) | input | Enemy.CombatLogUI_IDS_OnRowLButtonDown | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnHyperLinkLButtonUp** handlers call: `WindowSetShowing`

**OnLButtonDown** handlers call: `SystemData.ActiveWindow.name:match`

**OnLButtonUp** handlers call: `WindowGetId`, `WindowGetParent`, `WindowSetGameActionData`, `WindowSetShowing`

**OnMouseOver** handlers call: `CreateWindow`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `SystemData.ActiveWindow.name:match`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetGameActionData`

**OnMouseOverEnd** handlers call: `LabelSetTextColor`

**OnRButtonUp** handlers call: `SystemData.MouseOverWindow.name:match`, `WindowGetId`

**OnMouseOver** handlers call: `CreateWindow`, `LabelSetText`, `LabelSetTextColor`, `ListBoxGetDataIndex`, `SystemData.ActiveWindow.name:match`, `WindowGetDimensions`, `WindowGetId`, `WindowGetParent`, `WindowGetScreenPosition`, `WindowSetGameActionData`

**OnLButtonUp** handlers call: `WindowGetId`, `WindowGetParent`, `WindowSetGameActionData`, `WindowSetShowing`

**OnRButtonUp** handlers call: `SystemData.MouseOverWindow.name:match`, `WindowGetId`

**OnHyperLinkLButtonUp** handlers call: `WindowSetShowing`

**OnMouseOverEnd** handlers call: `LabelSetTextColor`

**OnLButtonDown** handlers call: `SystemData.ActiveWindow.name:match`

## Common Inherits

- Aggro_Label_Template
- Aura_Default_Label_SmallFont
- ClosetGoblinTalismanLabel
- DefaultWindowSmallText
- DefaultWindowText
- EA_Label_DefaultText
- EA_Label_DefaultText_Small
- EA_Settings_ItemTitle
- EA_Settings_SectionTitle
- Shinies_Default_Label_ClearSmallFont
- TChatLabel
- TRollLabel

## Common Parent Elements

- [Windows](element_Windows.md) — 1243× (HIGH)
- [Window](element_Window.md) — 15× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 1238× (HIGH)
- [Size](element_Size.md) — 1168× (HIGH)
- [Color](element_Color.md) — 375× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 95× (HIGH)
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

### [Anchors](element_Anchors.md)

Observed 1238 times as an unnamed child.

### [Size](element_Size.md)

Observed 1168 times as an unnamed child.

### [Color](element_Color.md)

Observed 375 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 73, 225, 255, 55 |
| `g` | **required** | 218, 255, 155, 55 |
| `r` | **required** | 255, 155, 245, 175 |
| `a` | optional | 255, 0, 0.5, 0.8 |
### [EventHandlers](element_EventHandlers.md)

Observed 95 times as an unnamed child.

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

## Recursive Hierarchy

- Root: [Label](element_Label.md)
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

- Enemy.IntercomInitialize
- Enemy.Stopwatch_Update
- APAGui.OnShown
- APAGui.UpdateInstantOnlyHUD
- ClosetGoblinZoneWindow.RefreshOption
- Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- TidyRollOptions.Initialize
- Enemy.IntercomUI_IntercomDialog_Open
- Enemy.CombatLogUI_EpsWindow_UpdateLayout
- MoraleCircle.ColorChanger2
- Enemy.Timer_Update
- Enemy.CombatLogUI_TargetDefenseWindow_Update
- MoraleCircle.ColorChanger4
- MoraleCircle.OnSetCustomColorEmpty
- Swift Assist.local.SetSmartLabel
- AggroMeter.Initialize
- Enemy.GroupsUI_EffectFilterDialog_Open
- Enemy.CombatLogUI_EpsWindow_Initialize
- Killer.ShowPersonalStatsTooltip
- Killer.ShowRowTooltip
- APAGui.UpdateFollowTargetHUD
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_UpdateAbilityIcon
- Enemy.Guard_GuardIndicator_Update
- Enemy.ScenarioInfoUI_ScenarioInfoDialog_Update
- Enemy.CombatLogUI_StatsWindow_Open
- Enemy.CombatLogUI_StatsWindow_OnSessionSelChanged
- APAGui.ApplyPetTargetHUDLayout
- DAoCBuff.ShowMessageWindow
- Enemy.UI_ConfigDialog_Open
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- DAoCTooltips.CreateCondenseTooltip
- MoraleCircle.ColorChanger3
- DAoCTooltips.UpdateCondenseTooltip
- Killer.Initialize
- MoraleCircle.ColorChanger1
- APAGui.UpdateKitingHUD
- Enemy.CombatLogUI_TargetDefenseTotalWindow_Update
- MoraleCircle.OnSetCustomColor
- MoraleCircle.OnSetCustomColorFull
- PP.UpdateDyeFilter
- MoraleCircle.OnSetCustomColorFill
- Killer.ShowTopKillersTooltip
- WSCT.ColorOnInitialize
- APAGui.UpdatePetTargetHUD
- BankArkel.SetCharInfo
- BankArkel.PackHide
- Enemy.AssistUI_Target_Show
- Enemy.CombatLogUI_EpsWindow_Update
- PotionBarSettings.OnAboutShown
- SwiftAssist.aaLabelColorSet
- Enemy.CombatLogUI_StatsWindow_UpdateList
- MoraleCircle.init
- PP.UpdateListRow
- Swift Assist.WriteLabels
- Swift Assist.SetSmartLabel
- ClosetGoblinCharacterWindow.OnInitialize
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.CombatLogUI_StatsWindow_ListRowMouseOver
- RoR_SoR.OnInitialize
- Swift Assist.local.WriteLabels


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

- [EA_Label_DefaultText](../../globals/constants/constant_EA_Label_DefaultText.md) (HIGH 100/100) - Constant
- [EA_Label_DefaultText_Small](../../globals/constants/constant_EA_Label_DefaultText_Small.md) (HIGH 100/100) - Constant
- [EA_Settings_ItemTitle](../../globals/constants/constant_EA_Settings_ItemTitle.md) (HIGH 100/100) - Constant
- [EA_Settings_SectionTitle](../../globals/constants/constant_EA_Settings_SectionTitle.md) (HIGH 100/100) - Constant
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Text](element_Text.md) (HIGH 98/100) - XML Element Type
- [EA_Label_DefaultSubHeading](../../globals/constants/constant_EA_Label_DefaultSubHeading.md) (HIGH 90/100) - Constant
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 55/100) - XML Element Type
- [Color](element_Color.md) (MEDIUM 45/100) - XML Element Type

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
