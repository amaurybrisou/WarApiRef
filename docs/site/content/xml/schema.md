# XML Schema

Observed schema information derived from generated frame pages. This is a usage-oriented schema, not an engine XML XSD.

## Element Types

| Element | Addons | Handlers | Attributes |
| --- | --- | --- | --- |
| AbsPoint | 256 |  | x, y |
| ActionButtonGroup | 3 | OnActionButtonLButtonDown, OnActionButtonLButtonUp, OnActionButtonMouseOver, OnActionButtonRButtonDown | draganddrop, gameactionbutton, hideButtonWhenIconBlank, inherits |
| ActiveZoneOffset | 2 |  | x, y |
| Anchor | 10 |  | point, relativePoint, relativeTo |
| Anchors | 250 |  |  |
| AnimFrame | 8 |  | id, x, y |
| AnimFrames | 8 |  |  |
| AnimatedImage | 17 |  | alpha, fps, handleinput, handlinput |
| AnimatedImages | 1 |  |  |
| Assets | 110 |  |  |
| Bottom | 3 |  | id, x, y |
| BottomCenter | 18 |  | id, x, y |
| BottomLeft | 18 |  | id, x, y |
| BottomRight | 18 |  | id, x, y |
| Button | 178 | OnLButtonUp | backgroundtexture, drawchildrenfirst, font, handleinput |
| CircleImage | 26 | OnLButtonDown | alpha, filtering, handleinput, id |
| Color | 168 |  | a, b, g, r |
| ColorPicker | 1 | OnLButtonUp | columnsPerRow, name, texture |
| ColorSize | 1 |  | x, y |
| ColorSpacing | 1 |  | x, y |
| ColorTexCoords | 1 |  | x, y |
| ColorTexDims | 1 |  | x, y |
| ComboBox | 92 | OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp | autoresize, id, inherits, layer |
| CooldownDisplay | 7 |  | alpha, cooldownshape, handleinput, layer |
| Disabled | 56 |  | a, b, def, g |
| DisabledPressed | 20 |  | a, b, g, id |
| DownOffset | 2 |  | x, y |
| DynamicImage | 149 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnMouseOverEnd | alpha, handleinput, id, inherits |
| EditBox | 70 | OnEnterPressed, OnFocusLost, OnKeyEnter, OnKeyEscape | font, handleinput, history, inherits |
| EventHandler | 210 |  | event, function |
| EventHandlers | 212 |  |  |
| Eventhandlers | 1 |  |  |
| FullResizeImage | 129 | OnLButtonUp, OnMouseOver, OnMouseOverEnd | alpha, autoresize, frameonly, handleinput |
| HorizontalResizeImage | 39 |  | alpha, handleinput, inherits, layer |
| Icon | 1 |  | id, name, texture |
| Include | 1 |  | file |
| Label | 232 | OnHyperLinkLButtonUp, OnHyperLinkRButtonUp, OnLButtonDown, OnLButtonUp | autoresize, autoresizewidth, font, handleinput |
| Left | 20 |  | id, x, y |
| ListBox | 56 | OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp | autoHideScrollBar, color, draganddrop, inherits |
| ListColumn | 45 |  | format, style, variable, windowname |
| ListColumns | 46 |  |  |
| ListData | 56 |  | populationfunction, table |
| LogDisplay | 3 | OnHyperLinkRButtonUp | autoHideScrollBar, background, font, handleinput |
| MapDisplay | 7 | OnHidden, OnLButtonUp, OnMButtonUp, OnMouseLeave | gutterIcon, handleinput, iconScale, layer |
| MenuButtonOffset | 24 |  | x, y |
| Middle | 36 |  | id, x, y |
| MiddleCenter | 18 |  | id, x, y |
| MiddleLeft | 18 |  | id, x, y |
| MiddleRight | 18 |  | id, x, y |
| NifDisplay | 1 | OnLButtonUp | autoprops, draganddrop, fov, handleinput |
| Normal | 77 |  | a, b, def, g |
| NormalHighlit | 75 |  | a, b, def, g |
| OverlayOffset | 22 |  | x, y |
| OverlaySize | 18 |  | x, y |
| OverlayTexCoords | 16 |  |  |
| OverlayTexSlices | 1 |  |  |
| PageWindow | 1 |  | inherits, name |
| Pressed | 74 |  | a, b, def, g |
| PressedHighlit | 64 |  | a, b, def, g |
| ResizeImages | 21 |  |  |
| Right | 20 |  | id, x, y |
| Rotation | 1 |  | x, y, z |
| Script | 118 |  | File, file |
| Scripts | 125 |  |  |
| ScrollWindow | 27 | OnMouseWheel | autoHideScrollBar, autohidescrollbar, autosize, childscrollwindow |
| Size | 252 |  |  |
| Sizes | 36 |  | bottom, left, middle, right |
| Slice | 17 |  | height, id, left, top |
| SliderBar | 44 | OnLButtonUp, OnMouseOver, OnSlide | autoresize, background, handleinput, handlinput |
| Sound | 9 |  | event, script |
| Sounds | 9 |  |  |
| StatusBar | 11 |  | background, foreground, handleinput, inherits |
| TexCoords | 75 |  | x, y |
| TexDims | 49 |  | x, y |
| TexSlices | 54 |  |  |
| Text | 4 |  | alignment, text |
| TextColors | 36 |  |  |
| TextOffset | 44 |  | x, y |
| Texture | 107 |  | file, name |
| ThumbOffset | 2 |  | x, y |
| TintColor | 74 |  | a, b, g, r |
| Top | 3 |  | id, x, y |
| TopCenter | 18 |  | id, x, y |
| TopLeft | 18 |  | id, x, y |
| TopRight | 18 |  | id, x, y |
| Translation | 1 |  | x, y, z |
| UpOffset | 2 |  | x, y |
| VerticalResizeImage | 10 |  | handleinput, inherits, layer, name |
| VerticalScrollbar | 32 | OnLButtonUp, OnScrollPosChanged | alpha, down, gutter, inherits |
| Visual | 1 |  |  |
| Window | 247 | OnHidden, OnInitialize, OnInitializeCustomSettings, OnKeyEnter | alpha, handleinput, id, inherits |
| Windows | 244 |  |  |
| color | 1 |  | a, b, g, r |

## Shared Inherits Constants

- EA_ActionButtonGroup_CareerIconsWithTooltip
- EA_ActionButtonGroup_DefaultSmall
- EA_BrownHorizontalRule
- EA_Button_BottomTab
- EA_Button_Default
- EA_Button_DefaultBigLeftArrow
- EA_Button_DefaultBigRightArrow
- EA_Button_DefaultChatScrollToBottom
- EA_Button_DefaultCheckBox
- EA_Button_DefaultDown
- EA_Button_DefaultIcon
- EA_Button_DefaultIconFrame
- EA_Button_DefaultIconFrame_Large
- EA_Button_DefaultIconFrame_Small
- EA_Button_DefaultLeftArrow
- EA_Button_DefaultListRow
- EA_Button_DefaultListSort
- EA_Button_DefaultMenuButton
- EA_Button_DefaultMinus
- EA_Button_DefaultPlus
- EA_Button_DefaultResizableComboBoxSelected
- EA_Button_DefaultResizableSmallComboBoxSelected
- EA_Button_DefaultResizeable
- EA_Button_DefaultRightArrow
