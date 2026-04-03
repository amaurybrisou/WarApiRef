# XML Schema

Observed schema information derived from generated frame pages. This is a usage-oriented schema, not an engine XML XSD.

## Element Types

| Element | Addons | Handlers | Attributes |
| --- | --- | --- | --- |
| AbsPoint | 34 |  | x, y |
| Anchor | 33 |  | point, relativePoint, relativeTo, xOffset |
| Anchors | 32 |  |  |
| AnimFrame | 1 |  | id, x, y |
| AnimFrames | 1 |  |  |
| AnimatedImage | 5 |  | alpha, fps, handleinput, inherits |
| AnimatedImages | 1 |  |  |
| Assets | 20 |  |  |
| BottomCenter | 5 |  | id, x, y |
| BottomLeft | 5 |  | id, x, y |
| BottomRight | 5 |  | id, x, y |
| Button | 27 | OnInitialize, OnLButtonDown, OnLButtonUp, OnMButtonUp | backgroundtexture, drawchildrenfirst, font, handleinput |
| CircleImage | 8 |  | filtering, handleinput, layer, movable |
| Color | 22 |  | a, b, g, r |
| ColorPicker | 1 | OnLButtonUp | columnsPerRow, name, texture |
| ColorSize | 1 |  | x, y |
| ColorSpacing | 1 |  | x, y |
| ColorTexCoords | 1 |  | x, y |
| ColorTexDims | 1 |  | x, y |
| ComboBox | 18 | OnMouseOver, OnSelChanged | autoresize, inherits, layer, maxvisibleitems |
| CooldownDisplay | 1 |  | cooldownshape, handleinput, layer, movable |
| Disabled | 12 |  | a, b, def, g |
| DisabledPressed | 6 |  | a, b, g, id |
| DynamicImage | 26 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnMouseOverEnd | alpha, filtering, handleinput, id |
| EditBox | 13 | OnKeyEnter, OnKeyEscape, OnMouseOver, OnTextChanged | background, font, handleinput, inherits |
| EventHandler | 31 |  | event, function |
| EventHandlers | 31 |  |  |
| FullResizeImage | 22 | OnMouseOver | alpha, drawchildrenfirst, font, frameonly |
| HorizontalResizeImage | 9 |  | handleinput, inherits, layer, name |
| Icon | 1 |  | id, texture |
| Label | 30 | OnHyperLinkLButtonUp, OnHyperLinkRButtonUp, OnLButtonDown, OnLButtonUp | autoresize, font, handleinput, inherits |
| Left | 3 |  | x, y |
| ListBox | 12 | OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp | color, draganddrop, layer, name |
| ListColumn | 11 |  | format, variable, windowname |
| ListColumns | 12 |  |  |
| ListData | 12 |  | populationfunction, table |
| LogDisplay | 1 |  | background, font, maxchars, maxentries |
| MapDisplay | 1 | OnLButtonUp, OnMButtonUp, OnPointMouseOver, OnRButtonUp | iconScale, layer, loadingAnim, movable |
| MenuButtonOffset | 4 |  | x, y |
| Middle | 7 |  | x, y |
| MiddleCenter | 5 |  | id, x, y |
| MiddleLeft | 5 |  | id, x, y |
| MiddleRight | 5 |  | id, x, y |
| Normal | 15 |  | a, b, def, g |
| NormalHighlit | 14 |  | a, b, def, g |
| OverlayOffset | 4 |  | x, y |
| OverlaySize | 4 |  | x, y |
| OverlayTexCoords | 4 |  |  |
| Pressed | 13 |  | a, b, def, g |
| PressedHighlit | 10 |  | a, b, def, g |
| ResizeImages | 5 |  |  |
| Right | 3 |  | x, y |
| Script | 12 |  | file |
| Scripts | 12 |  |  |
| ScrollWindow | 8 |  | autoHideScrollBar, autohidescrollbar, childscrollwindow, inherits |
| Size | 33 |  |  |
| Sizes | 7 |  | bottom, left, middle, right |
| Slice | 2 |  | height, id, left, top |
| SliderBar | 14 | OnMouseOver, OnSlide | background, inherits, locktoticks, name |
| StatusBar | 2 |  | background, foreground, handleinput, inherits |
| TexCoords | 14 |  | x, y |
| TexDims | 11 |  | x, y |
| TexSlices | 8 |  |  |
| Text | 2 |  | alignment, text |
| TextColors | 5 |  |  |
| TextOffset | 8 |  | x, y |
| Texture | 19 |  | file |
| TintColor | 14 |  | a, b, g, r |
| TopCenter | 5 |  | id, x, y |
| TopLeft | 5 |  | id, x, y |
| TopRight | 5 |  | id, x, y |
| VerticalResizeImage | 1 |  | inherits, name, texture |
| VerticalScrollbar | 7 | OnScrollPosChanged | inherits, layer, name |
| Visual | 1 |  |  |
| Window | 32 | OnHidden, OnInitialize, OnKeyEscape, OnLButtonDown | alpha, drawchildrenfirst, handleinput, inherits |
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
