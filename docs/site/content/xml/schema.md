# XML Schema

Observed schema information derived from generated frame pages. This is a usage-oriented schema, not an engine XML XSD.

## Element Types

| Element | Addons | Handlers | Attributes |
| --- | --- | --- | --- |
| AnimFrames | 1 |  |  |
| AnimatedImage | 5 |  | layer, name, inherits, handleinput |
| AnimatedImages | 1 |  |  |
| Button | 27 | OnLButtonUp, OnMouseOver, OnRButtonUp, OnMouseOverEnd | name, inherits, id, textalign |
| CircleImage | 8 |  | name, layer, numsegments, popable |
| ColorPicker | 1 | OnLButtonUp | columnsPerRow, name, texture |
| ColorSize | 1 |  | x, y |
| ColorSpacing | 1 |  | x, y |
| ColorTexCoords | 1 |  | x, y |
| ColorTexDims | 1 |  | x, y |
| ComboBox | 18 | OnSelChanged, OnMouseOver | name, inherits, layer, maxvisibleitems |
| CooldownDisplay | 1 |  | cooldownshape, handleinput, layer, movable |
| DynamicImage | 26 | OnMouseOver, OnLButtonUp, OnRButtonUp, OnLButtonDown | name, handleinput, texture, layer |
| EditBox | 13 | OnTextChanged, OnKeyEnter, OnKeyEscape, OnMouseOver | name, inherits, maxchars, input |
| FullResizeImage | 22 | OnMouseOver | name, inherits, handleinput, layer |
| HorizontalResizeImage | 9 |  | name, inherits, texture, handleinput |
| Label | 30 | OnMouseOver, OnLButtonUp, OnRButtonUp, OnHyperLinkLButtonUp | name, textalign, font, inherits |
| ListBox | 12 | OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp | name, rowdef, rowspacing, visiblerows |
| ListData | 12 |  | table, populationfunction |
| LogDisplay | 1 |  | background, font, maxchars, maxentries |
| MapDisplay | 1 | OnLButtonUp, OnMButtonUp, OnPointMouseOver, OnRButtonUp | iconScale, layer, loadingAnim, movable |
| MenuButtonOffset | 4 |  | x, y |
| OverlayOffset | 4 |  | x, y |
| OverlaySize | 4 |  | x, y |
| OverlayTexCoords | 4 |  |  |
| ResizeImages | 5 |  |  |
| ScrollWindow | 8 |  | name, childscrollwindow, lineheight, scrollbar |
| Size | 33 |  |  |
| Sizes | 7 |  | middle, left, right, bottom |
| SliderBar | 14 | OnSlide, OnMouseOver | name, inherits, numticks, locktoticks |
| StatusBar | 2 |  | handleinput, name, popable, foreground |
| TexCoords | 14 |  | x, y |
| TexDims | 11 |  | x, y |
| TexSlices | 8 |  |  |
| Text | 2 |  | text, alignment |
| TextColors | 5 |  |  |
| TextOffset | 8 |  | x, y |
| TintColor | 14 |  | b, g, r, a |
| VerticalResizeImage | 1 |  | name, inherits, texture |
| VerticalScrollbar | 7 | OnScrollPosChanged | inherits, name, layer |
| Visual | 1 |  |  |
| Window | 32 | OnLButtonUp, OnHidden, OnMouseOver, OnRButtonUp | name, inherits, layer, movable |
| Windows | 32 |  |  |

## Shared Inherits Constants

- EA_BrownHorizontalRule
- EA_Button_BottomTab
- EA_Button_Default
- EA_Button_DefaultBigLeftArrow
- EA_Button_DefaultBigRightArrow
- EA_Button_DefaultCheckBox
- EA_Button_DefaultDown
- EA_Button_DefaultIconFrame
- EA_Button_DefaultIconFrame_Large
- EA_Button_DefaultIconFrame_Small
- EA_Button_DefaultListSort
- EA_Button_DefaultMenuButton
- EA_Button_DefaultMinus
- EA_Button_DefaultPlus
- EA_Button_DefaultResizeable
- EA_Button_DefaultSmallSquare
- EA_Button_DefaultText
- EA_Button_DefaultToggleCircle
- EA_Button_DefaultUp
- EA_Button_DefaultWindowClose
- EA_Button_ListSort
- EA_Button_ResizeIconFrameNormal
- EA_Button_ResizeIconFrame_NoNormalImage
- EA_Button_Tab
