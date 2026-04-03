# XML Schema

Observed schema information derived from generated frame pages. This is a usage-oriented schema, not an engine XML XSD.

## Element Types

| Element | Addons | Handlers | Attributes |
| --- | --- | --- | --- |
| AnimFrame | 1 |  | id, x, y |
| AnimFrames | 1 |  |  |
| AnimatedImage | 1 |  | fps, handleinput, layer, name |
| Button | 5 | OnLButtonUp, OnMouseOver, OnRButtonUp | name, inherits, layer, handleinput |
| CircleImage | 1 |  | handleinput, layer, name, numsegments |
| ComboBox | 2 | OnSelChanged | inherits, layer, name |
| Disabled | 2 |  | id, texture, x, y |
| DisabledPressed | 1 |  | id |
| DynamicImage | 6 |  | name, handleinput, texture, textureScale |
| EditBox | 2 |  | inherits, name, font, maxchars |
| FullResizeImage | 7 |  | name, inherits, handleinput, alpha |
| HorizontalResizeImage | 1 |  | name, texture, textureScale |
| Label | 7 | OnHyperLinkRButtonUp, OnHyperLinkLButtonUp | name, font, textalign, layer |
| Left | 1 |  | x, y |
| ListBox | 2 | OnLButtonUp | name, rowcount, rowdef, rowspacing |
| ListColumn | 2 |  | format, variable, windowname |
| ListColumns | 2 |  |  |
| ListData | 2 |  | table, populationfunction |
| Middle | 1 |  | x, y |
| Normal | 2 |  | id, texture, x, y |
| NormalHighlit | 2 |  | id, texture, x, y |
| Pressed | 2 |  | id, texture, x, y |
| PressedHighlit | 1 |  | id |
| Right | 1 |  | x, y |
| ScrollWindow | 1 |  | autoHideScrollBar, childscrollwindow, lineheight, name |
| SliderBar | 1 |  | inherits, name |
| StatusBar | 1 |  | handleinput, name, popable, inherits |
| TexSlices | 1 |  |  |
| TintColor | 5 |  | b, g, r, a |
| Window | 7 | OnHidden, OnShown, OnLButtonUp, OnMouseWheel | name, layer, movable, handleinput |

## Shared Inherits Constants

- EA_Button_Default
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
- EA_StatusBar_DefaultTintable
- EA_TitleBar_Default
