# XML Schema

Observed schema information derived from generated frame pages. This is a usage-oriented schema, not an engine XML XSD.

## Element Types

| Element | Addons | Handlers | Attributes |
| --- | --- | --- | --- |
| ActionButtonGroup | 3 | OnActionButtonLButtonDown, OnActionButtonLButtonUp, OnActionButtonMouseOver, OnActionButtonRButtonDown | inherits, name, hideButtonWhenIconBlank, draganddrop |
| AnimatedImage | 10 |  | name, layer, handleinput, texture |
| Button | 58 | OnLButtonUp, OnMouseOver, OnRButtonUp, OnLButtonDown | name, inherits, id, textalign |
| CircleImage | 11 |  | name, layer, handleinput, numsegments |
| ColorPicker | 1 | OnLButtonUp | columnsPerRow, name, texture |
| ComboBox | 33 | OnSelChanged, OnMouseOver, OnLButtonUp, OnRButtonUp | name, inherits, layer, maxvisibleitems |
| CooldownDisplay | 1 |  | cooldownshape, handleinput, layer, movable |
| DynamicImage | 45 | OnMouseOver, OnLButtonUp, OnRButtonUp, OnLButtonDown | name, handleinput, texture, layer |
| EditBox | 26 | OnTextChanged, OnKeyEnter, OnKeyEscape, OnMouseOver | name, inherits, maxchars, input |
| FullResizeImage | 44 | OnMouseOver, OnMouseOverEnd, OnLButtonUp | name, inherits, handleinput, layer |
| HorizontalResizeImage | 12 |  | name, inherits, texture, handleinput |
| Label | 66 | OnMouseOver, OnLButtonUp, OnRButtonUp, OnMouseOverEnd | name, textalign, font, inherits |
| ListBox | 20 | OnLButtonUp, OnMouseOver, OnMouseOverEnd, OnRButtonUp | name, rowdef, rowspacing, visiblerows |
| LogDisplay | 2 |  | background, font, maxchars, maxentries |
| MapDisplay | 3 | OnLButtonUp, OnPointMouseOver, OnMouseWheel, OnHidden | name, pinTexture, layer, movable |
| ScrollWindow | 12 |  | name, childscrollwindow, lineheight, scrollbar |
| SliderBar | 17 | OnSlide, OnMouseOver | name, inherits, numticks, handleinput |
| StatusBar | 2 |  | name, popable, foreground, handleinput |
| VerticalResizeImage | 3 |  | name, texture, inherits, handleinput |
| VerticalScrollbar | 13 | OnScrollPosChanged | name, inherits, layer, alpha |
| Window | 71 | OnLButtonUp, OnHidden, OnMouseOver, OnRButtonUp | name, inherits, layer, handleinput |

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
- EA_Button_DefaultTextButton
- EA_Button_DefaultToggleCircle
