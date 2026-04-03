# Windows

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Aura, AutoMark, BankArkel, BuffHead |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.xml:0`, `/workspace/data/raw/Aura/Source/AuraColorPicker.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTexture.xml:0` |
| Namespaces detected | Windows |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD, AdvancedPetAssist: APAInstantOnlyHUD, AdvancedPetAssist: APAKitingHUD, AdvancedPetAssist: APAOptions, AdvancedPetAssist: APAOptionsContent, AdvancedPetAssist: APAOptionsTabs |
| XML usage count | 530 |
| XML attribute usage count | 530 |
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

Observed XML element type instantiated by 32 addons.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Window](element_Window.md) — 476× (HIGH)
- [ScrollWindow](element_ScrollWindow.md) — 24× (HIGH)
- [Button](element_Button.md) — 18× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 6× (MEDIUM)
- [AnimatedImage](element_AnimatedImage.md) — 2× (LOW)
- [CircleImage](element_CircleImage.md) — 1× (LOW)
- [FullResizeImage](element_FullResizeImage.md) — 1× (LOW)
- [Label](element_Label.md) — 1× (LOW)
- [LogDisplay](element_LogDisplay.md) — 1× (LOW)

## Common Named Child Elements

- [Label](element_Label.md) — 1243× (HIGH)
- [Window](element_Window.md) — 1001× (HIGH)
- [Button](element_Button.md) — 664× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 237× (HIGH)
- [ComboBox](element_ComboBox.md) — 233× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 156× (HIGH)
- [EditBox](element_EditBox.md) — 151× (HIGH)
- [SliderBar](element_SliderBar.md) — 83× (HIGH)
- [ListBox](element_ListBox.md) — 42× (HIGH)
- [ScrollWindow](element_ScrollWindow.md) — 26× (HIGH)
- [VerticalScrollbar](element_VerticalScrollbar.md) — 25× (HIGH)
- [CircleImage](element_CircleImage.md) — 23× (HIGH)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 19× (HIGH)
- [AnimatedImage](element_AnimatedImage.md) — 12× (HIGH)
- [StatusBar](element_StatusBar.md) — 9× (MEDIUM)
- [VerticalResizeImage](element_VerticalResizeImage.md) — 4× (MEDIUM)
- [ColorPicker](element_ColorPicker.md) — 1× (LOW)
- [CooldownDisplay](element_CooldownDisplay.md) — 1× (LOW)
- [LogDisplay](element_LogDisplay.md) — 1× (LOW)
- [MapDisplay](element_MapDisplay.md) — 1× (LOW)

## Recursive Hierarchy

- Root: [Windows](element_Windows.md)
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

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Aura
- AutoMark
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

- AdvancedPetAssist: APAFollowTargetHUD -> Windows in Window APAFollowTargetHUD
- AdvancedPetAssist: APAInstantOnlyHUD -> Windows in Window APAInstantOnlyHUD
- AdvancedPetAssist: APAKitingHUD -> Windows in Window APAKitingHUD
- AdvancedPetAssist: APAOptions -> Windows in Window APAOptions
- AdvancedPetAssist: APAOptionsContent -> Windows in Window APAOptionsContent
- AdvancedPetAssist: APAOptionsTabs -> Windows in Window APAOptionsTabs

## Related APIs

- [AnimatedImage](element_AnimatedImage.md) (HIGH 100/100) - XML Element Type
- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [CircleImage](element_CircleImage.md) (HIGH 100/100) - XML Element Type
- [ColorPicker](element_ColorPicker.md) (HIGH 100/100) - XML Element Type
- [ComboBox](element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EditBox](element_EditBox.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [HorizontalResizeImage](element_HorizontalResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [ScrollWindow](element_ScrollWindow.md) (HIGH 100/100) - XML Element Type
- [SliderBar](element_SliderBar.md) (HIGH 100/100) - XML Element Type
- [StatusBar](element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [VerticalScrollbar](element_VerticalScrollbar.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [VerticalResizeImage](element_VerticalResizeImage.md) (HIGH 90/100) - XML Element Type
- [CooldownDisplay](element_CooldownDisplay.md) (MEDIUM 45/100) - XML Element Type
- [LogDisplay](element_LogDisplay.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 55/100) - XML Element Type

## Triggered By

- none

## Affects

- none
