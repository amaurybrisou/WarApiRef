# XML Schema

Observed schema information derived from generated frame pages. This is a usage-oriented schema, not an engine XML XSD.

## Element Types

| Element | Addons | Handlers | Attributes |
| --- | --- | --- | --- |
| AbsPoint | 2 |  | x, y |
| Anchor | 2 |  | point, relativePoint, relativeTo |
| Anchors | 2 |  |  |
| Assets | 1 |  |  |
| BottomCenter | 1 |  | x, y |
| BottomLeft | 1 |  | x, y |
| BottomRight | 1 |  | x, y |
| Button | 1 | OnLButtonDown, OnLButtonUp, OnMouseOver, OnMouseOverEnd | draganddrop, drawchildrenfirst, font, gameactionbutton |
| Color | 2 |  | a, b, g, r |
| Disabled | 1 |  | texture, x, y |
| DynamicImage | 1 | OnLButtonUp, OnRButtonUp | handleinput, inherits, layer, movable |
| EventHandler | 1 |  | event, function |
| EventHandlers | 1 |  |  |
| FullResizeImage | 1 |  | handleinput, inherits, layer, name |
| Label | 2 | OnLButtonUp, OnRButtonUp | autoresize, autoresizewidth, font, handleinput |
| ListBox | 1 |  | layer, name, rowcount, rowdef |
| ListColumn | 1 |  | format, variable, windowname |
| ListColumns | 1 |  |  |
| ListData | 1 |  | populationfunction, table |
| Middle | 1 |  | x, y |
| MiddleCenter | 1 |  | x, y |
| MiddleLeft | 1 |  | x, y |
| MiddleRight | 1 |  | x, y |
| Normal | 1 |  | texture, x, y |
| NormalHighlit | 1 |  | texture, x, y |
| Pressed | 1 |  | texture, x, y |
| Script | 2 |  | file |
| Scripts | 2 |  |  |
| Size | 2 |  |  |
| Sizes | 1 |  |  |
| TexCoords | 1 |  |  |
| TextOffset | 1 |  | x |
| Texture | 1 |  | file, name |
| TopCenter | 1 |  | x, y |
| TopLeft | 1 |  | x, y |
| TopRight | 1 |  | x, y |
| Window | 2 | OnHidden, OnInitialize, OnShown, OnShutdown | draganddrop, handleinput, inherits, layer |
| Windows | 2 |  |  |

## Shared Inherits Constants

- EA_Button_Default
- EA_Button_DefaultCheckBox
- EA_Button_DefaultDown
- EA_Button_DefaultResizeable
- EA_Button_DefaultUp
- EA_Button_DefaultWindowClose
- EA_Button_ListSort
- EA_Default_CharacterImage
- EA_FullResizeImage_TintableSolidBackground
- EA_ListSortDownArrow
- EA_ListSortUpArrow
- EA_TitleBar_Default
- EA_Window_Default
- EA_Window_DefaultFrame
