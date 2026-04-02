# XML Schema

Observed schema information derived from generated frame pages. This is a usage-oriented schema, not an engine XML XSD.

## Element Types

| Element | Addons | Handlers | Attributes |
| --- | --- | --- | --- |
| ActionButtonGroup | 3 | OnActionButtonLButtonDown, OnActionButtonLButtonUp, OnActionButtonMouseOver, OnActionButtonRButtonDown | inherits, name, hideButtonWhenIconBlank, draganddrop |
| AnimatedImage | 10 |  | name, layer, handleinput, texture |
| Button | 51 | OnLButtonUp, OnMouseOver, OnRButtonUp, OnLButtonDown | name, inherits, id, textalign |
| CircleImage | 11 |  | name, layer, handleinput, numsegments |
| ColorPicker | 1 | OnLButtonUp | columnsPerRow, name, texture |
| ComboBox | 29 | OnSelChanged, OnMouseOver | name, inherits, layer, maxvisibleitems |
| CooldownDisplay | 1 |  | cooldownshape, handleinput, layer, movable |
| DynamicImage | 41 | OnMouseOver, OnLButtonUp, OnRButtonUp, OnLButtonDown | name, handleinput, texture, layer |
| EditBox | 24 | OnTextChanged, OnKeyEnter, OnKeyEscape, OnMouseOver | name, inherits, maxchars, input |
| FullResizeImage | 41 | OnMouseOver, OnMouseOverEnd, OnLButtonUp | name, inherits, handleinput, layer |
| HorizontalResizeImage | 12 |  | name, inherits, texture, handleinput |
| Label | 56 | OnMouseOver, OnLButtonUp, OnRButtonUp, OnMouseOverEnd | name, textalign, font, inherits |
| ListBox | 20 | OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp | name, rowdef, rowspacing, visiblerows |
| LogDisplay | 2 |  | background, font, maxchars, maxentries |
| MapDisplay | 3 | OnLButtonUp, OnPointMouseOver, OnMouseWheel, OnHidden | name, pinTexture, layer, movable |
| ScrollWindow | 11 |  | name, childscrollwindow, lineheight, scrollbar |
| SliderBar | 17 | OnSlide, OnMouseOver | name, inherits, numticks, handleinput |
| StatusBar | 2 |  | name, popable, foreground, handleinput |
| VerticalResizeImage | 3 |  | name, texture, inherits, handleinput |
| VerticalScrollbar | 12 | OnScrollPosChanged | name, inherits, layer, alpha |
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
