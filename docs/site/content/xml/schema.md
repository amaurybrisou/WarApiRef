# XML Schema

Observed schema information derived from generated frame pages. This is a usage-oriented schema, not an engine XML XSD.

## Element Types

| Element | Addons | Handlers | Attributes |
| --- | --- | --- | --- |
| ActionButtonGroup | 3 | OnActionButtonLButtonDown, OnActionButtonLButtonUp, OnActionButtonMouseOver, OnActionButtonRButtonDown | inherits, name, hideButtonWhenIconBlank, draganddrop |
| ActiveZoneOffset | 1 |  | x, y |
| AnimFrame | 3 |  | id, x, y |
| AnimFrames | 3 |  |  |
| AnimatedImage | 10 |  | name, layer, handleinput, texture |
| AnimatedImages | 1 |  |  |
| Bottom | 1 |  | x, y |
| Button | 51 | OnLButtonUp, OnMouseOver, OnRButtonUp, OnLButtonDown | name, inherits, id, textalign |
| CircleImage | 11 |  | name, layer, handleinput, numsegments |
| ColorPicker | 1 | OnLButtonUp | columnsPerRow, name, texture |
| ColorSize | 1 |  | x, y |
| ColorSpacing | 1 |  | x, y |
| ColorTexCoords | 1 |  | x, y |
| ColorTexDims | 1 |  | x, y |
| ComboBox | 29 | OnSelChanged, OnMouseOver | name, inherits, layer, maxvisibleitems |
| CooldownDisplay | 1 |  | cooldownshape, handleinput, layer, movable |
| Disabled | 20 |  | x, y, b, g |
| DisabledPressed | 8 |  | id, a, b, g |
| DownOffset | 1 |  | x, y |
| DynamicImage | 41 | OnMouseOver, OnLButtonUp, OnRButtonUp, OnLButtonDown | name, handleinput, texture, layer |
| EditBox | 24 | OnTextChanged, OnKeyEnter, OnKeyEscape, OnMouseOver | name, inherits, maxchars, input |
| Eventhandlers | 1 |  |  |
| FullResizeImage | 41 | OnMouseOver, OnMouseOverEnd, OnLButtonUp | name, inherits, handleinput, layer |
| HorizontalResizeImage | 12 |  | name, inherits, texture, handleinput |
| Label | 56 | OnMouseOver, OnLButtonUp, OnRButtonUp, OnMouseOverEnd | name, textalign, font, inherits |
| Left | 5 |  | x, y, id |
| ListBox | 20 | OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp | name, rowdef, rowspacing, visiblerows |
| ListColumn | 14 |  | format, variable, windowname |
| ListColumns | 15 |  |  |
| ListData | 20 |  | table, populationfunction |
| LogDisplay | 2 |  | background, font, maxchars, maxentries |
| MapDisplay | 3 | OnLButtonUp, OnPointMouseOver, OnMouseWheel, OnHidden | name, pinTexture, layer, movable |
| MenuButtonOffset | 8 |  | x, y |
| Middle | 11 |  | x, y, id |
| Normal | 24 |  | id, x, y, b |
| NormalHighlit | 23 |  | x, y, id, b |
| OverlayOffset | 7 |  | x, y |
| OverlaySize | 7 |  | x, y |
| OverlayTexCoords | 7 |  |  |
| Pressed | 22 |  | id, x, y, b |
| PressedHighlit | 17 |  | id, b, g, r |
| ResizeImages | 9 |  |  |
| Right | 5 |  | x, y, id |
| ScrollWindow | 11 |  | name, childscrollwindow, lineheight, scrollbar |
| SliderBar | 17 | OnSlide, OnMouseOver | name, inherits, numticks, handleinput |
| Sound | 4 |  | event, script |
| Sounds | 4 |  |  |
| StatusBar | 2 |  | name, popable, foreground, handleinput |
| TexSlices | 13 |  |  |
| Text | 4 |  | text, alignment |
| TextColors | 9 |  |  |
| TextOffset | 15 |  | x, y |
| ThumbOffset | 1 |  | x, y |
| TintColor | 21 |  | b, g, r, a |
| Top | 1 |  | x, y |
| UpOffset | 1 |  | x, y |
| VerticalResizeImage | 3 |  | name, texture, inherits, handleinput |
| VerticalScrollbar | 12 | OnScrollPosChanged | name, inherits, layer, alpha |
| Visual | 1 |  |  |
| Window | 60 | OnLButtonUp, OnHidden, OnMouseOver, OnShown | name, inherits, layer, handleinput |

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
- EA_Button_DefaultIconFrame
- EA_Button_DefaultIconFrame_Large
- EA_Button_DefaultIconFrame_Small
- EA_Button_DefaultListRow
- EA_Button_DefaultListSort
- EA_Button_DefaultMenuButton
- EA_Button_DefaultMinus
- EA_Button_DefaultPlus
- EA_Button_DefaultResizableComboBoxSelected
- EA_Button_DefaultResizeable
- EA_Button_DefaultSmallSquare
- EA_Button_DefaultText
- EA_Button_DefaultToggleCircle
- EA_Button_DefaultUp
