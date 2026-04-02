# XML Schema

Observed schema information derived from generated frame pages. This is a usage-oriented schema, not an engine XML XSD.

## Element Types

| Element | Addons | Handlers | Attributes |
| --- | --- | --- | --- |
| AnimFrame | 1 |  | id, x, y |
| AnimFrames | 1 |  |  |
| AnimatedImage | 1 |  | fps, handleinput, layer, name |
| Button | 2 | OnLButtonUp, OnMouseOver | name, inherits, font, id |
| ComboBox | 2 | OnSelChanged | inherits, layer, name |
| Disabled | 1 |  | texture, x, y |
| DynamicImage | 2 |  | name, handleinput, layer, inherits |
| EditBox | 2 |  | inherits, name, font, maxchars |
| FullResizeImage | 3 |  | inherits, name, handleinput, alpha |
| Label | 3 |  | name, inherits, font, handleinput |
| ListBox | 2 | OnLButtonUp | name, rowcount, rowdef, rowspacing |
| ListColumn | 2 |  | format, variable, windowname |
| ListColumns | 2 |  |  |
| ListData | 2 |  | table, populationfunction |
| Normal | 1 |  | texture, x, y |
| NormalHighlit | 1 |  | texture, x, y |
| Pressed | 1 |  | texture, x, y |
| SliderBar | 1 |  | inherits, name |
| TintColor | 1 |  | b, g, r |
| Window | 3 | OnHidden, OnShown, OnLButtonUp, OnMouseWheel | name, inherits, layer, movable |

## Shared Inherits Constants

- EA_Button_DefaultCheckBox
- EA_Button_DefaultResizeable
- EA_Button_DefaultToggleCircle
- EA_Button_DefaultWindowClose
- EA_Button_ResizeIconFrameHighlight
- EA_Button_ResizeIconFrameNormal
- EA_Button_ResizeIconFramePressed
- EA_Button_Tab
- EA_ComboBox_DefaultResizable
- EA_ComboBox_DefaultResizableSmall
- EA_Default_SliderBar
- EA_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame_Multiline
- EA_FullResizeImage_BlackTransparent
- EA_FullResizeImage_DefaultFrame
- EA_FullResizeImage_MetalFill
- EA_FullResizeImage_TanBorder
- EA_FullResizeImage_TintableSolidBackground
- EA_Image_DefaultIcon
- EA_Image_DefaultIconFrame
- EA_Label_DefaultText
- EA_TitleBar_Default
- EA_Window_DefaultBackgroundFrame
- EA_Window_DefaultButtonBottomFrame
