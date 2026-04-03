# Coverage Report

## Source Counts

| Input | Count |
| --- | --- |
| bindings | 57 |
| event_docs | 0 |
| examples | 0 |
| flows | 0 |
| frame_docs | 167 |
| function_docs | 174 |
| handler_docs | 0 |

## Symbol Counts

| Symbol Type | Count |
| --- | --- |
| constants | 14 |
| element_types | 38 |
| game_events | 0 |
| gamedata | 2 |
| global_functions | 9 |
| systemdata | 10 |
| tables | 5 |
| window_events | 0 |
| window_functions | 11 |
| xml_handlers | 0 |

## Confidence

| Level | Count |
| --- | --- |
| HIGH | 56 |
| MEDIUM | 12 |
| LOW | 15 |

## Candidate Outcomes

| Outcome | Count |
| --- | --- |
| High confidence platform | 56 |
| Medium confidence candidates | 12 |
| Low confidence symbols | 9 |
| Rejected addon-local | 6 |

## Spread

- Symbols seen once: 73

- Symbols seen in multiple addons: 10

## High Confidence Platform

| Name | Category | Score | Level | Evidence | Rationale |
| --- | --- | --- | --- | --- | --- |
| GameData.EquipSlots.BACK | GameData Field | 100 | HIGH | addons=1; global=2; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| GameData.EquipSlots.HELM | GameData Field | 100 | HIGH | addons=1; global=2; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.ActiveWindow.name | SystemData Field | 100 | HIGH | addons=1; global=10; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.ALL_MODULES_INITIALIZED | SystemData Field | 100 | HIGH | addons=1; global=2; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.LOADING_BEGIN | SystemData Field | 100 | HIGH | addons=1; global=3; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.LOADING_END | SystemData Field | 100 | HIGH | addons=1; global=5; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED | SystemData Field | 100 | HIGH | addons=1; global=3; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.PLAYER_ZONE_CHANGED | SystemData Field | 100 | HIGH | addons=1; global=3; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.RELOAD_INTERFACE | SystemData Field | 100 | HIGH | addons=1; global=2; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.UPDATE_PROCESSED | SystemData Field | 100 | HIGH | addons=1; global=3; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.VISIBLE_EQUIPMENT_UPDATED | SystemData Field | 100 | HIGH | addons=1; global=2; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.MouseOverWindow.name | SystemData Field | 100 | HIGH | addons=1; global=6; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| LabelSetText | Window Function | 100 | HIGH | addons=2; global=21; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| LabelSetTextColor | Window Function | 100 | HIGH | addons=2; global=3; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| LayoutEditor.RegisterWindow | Window Function | 100 | HIGH | addons=2; global=2; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| Label | XML Element Type | 100 | HIGH | addons=2; xml=61; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| Window | XML Element Type | 100 | HIGH | addons=2; xml=25; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| Windows | XML Element Type | 100 | HIGH | addons=2; xml=15; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| Anchors | XML Element Type | 98 | HIGH | addons=2; xml=153; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, seen in 2 to 3 addons. |
| Color | XML Element Type | 98 | HIGH | addons=2; xml=4; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, seen in 2 to 3 addons. |
| Size | XML Element Type | 98 | HIGH | addons=2; xml=39; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, seen in 2 to 3 addons. |
| EA_Button_Default | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_DefaultCheckBox | Constant | 90 | HIGH | addons=1; xml=9; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_DefaultDown | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_DefaultResizeable | Constant | 90 | HIGH | addons=1; xml=4; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_DefaultUp | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_DefaultWindowClose | Constant | 90 | HIGH | addons=1; xml=2; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_ListSort | Constant | 90 | HIGH | addons=1; xml=4; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Default_CharacterImage | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_FullResizeImage_TintableSolidBackground | Constant | 90 | HIGH | addons=1; xml=2; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_ListSortDownArrow | Constant | 90 | HIGH | addons=1; xml=4; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_ListSortUpArrow | Constant | 90 | HIGH | addons=1; xml=4; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_TitleBar_Default | Constant | 90 | HIGH | addons=1; xml=2; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Window_Default | Constant | 90 | HIGH | addons=1; xml=4; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Window_DefaultFrame | Constant | 90 | HIGH | addons=1; xml=5; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| DialogManager.MakeTextEntryDialog | Global Function | 90 | HIGH | addons=1; global=4; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| EA_Window_ContextMenu.AddMenuItem | Global Function | 90 | HIGH | addons=1; global=12; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| EA_Window_ContextMenu.CreateContextMenu | Global Function | 90 | HIGH | addons=1; global=3; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| EA_Window_ContextMenu.Finalize | Global Function | 90 | HIGH | addons=1; global=3; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ButtonSetDisabledFlag | Window Function | 90 | HIGH | addons=1; global=4; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ButtonSetPressedFlag | Window Function | 90 | HIGH | addons=1; global=15; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ButtonSetText | Window Function | 90 | HIGH | addons=1; global=8; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowGetId | Window Function | 90 | HIGH | addons=1; global=11; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowSetShowing | Window Function | 90 | HIGH | addons=1; global=30; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| Button | XML Element Type | 90 | HIGH | addons=1; xml=57; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| DynamicImage | XML Element Type | 90 | HIGH | addons=1; xml=13; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| FullResizeImage | XML Element Type | 90 | HIGH | addons=1; xml=7; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| ListBox | XML Element Type | 90 | HIGH | addons=1; xml=2; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| ListData | XML Element Type | 90 | HIGH | addons=1; xml=2; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| DialogManager.MakeTwoButtonDialog | Global Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| Cursor | Global Table | 80 | HIGH | addons=1; global=2; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| DialogManager | Global Table | 80 | HIGH | addons=1; global=2; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| EA_Window_ContextMenu | Global Table | 80 | HIGH | addons=1; global=3; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| DynamicImageSetTexture | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowSetParent | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowSetTintColor | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |

## Medium Confidence Candidates

| Name | Category | Score | Level | Evidence | Rationale |
| --- | --- | --- | --- | --- | --- |
| CreateWindow | Global Function | 63 | MEDIUM | addons=2; global=2; sources=lua_calls | Kept as MEDIUM confidence because called globally with no local definition, seen in 2 to 3 addons, role is consistent across addons. |
| LibSlash.RegisterSlashCmd | Global Function | 55 | MEDIUM | addons=1; global=1; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, appears in slash command registration patterns. |
| wstring.format | Global Function | 55 | MEDIUM | addons=1; global=2; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, argument pattern is consistent. |
| LibSlash | Global Table | 55 | MEDIUM | addons=1; global=1; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, appears in slash command registration patterns. |
| Cursor.Clear | Global Function | 45 | MEDIUM | addons=1; global=1; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition. |
| Cursor.IconOnCursor | Global Function | 45 | MEDIUM | addons=1; global=1; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition. |
| wstring | Global Table | 45 | MEDIUM | addons=1; global=1; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition. |
| EventHandlers | XML Element Type | 45 | MEDIUM | addons=1; xml=34; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| Sizes | XML Element Type | 45 | MEDIUM | addons=1; xml=1; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| TexCoords | XML Element Type | 45 | MEDIUM | addons=1; xml=2; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| TextOffset | XML Element Type | 45 | MEDIUM | addons=1; xml=1; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| Texture | XML Element Type | 45 | MEDIUM | addons=1; xml=2; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |

## Low Confidence Symbols

| Name | Category | Score | Level | Evidence | Rationale |
| --- | --- | --- | --- | --- | --- |
| OnHidden | XML Handler | 23 | LOW | addons=1; sources=bindings,xml_handlers | Kept as LOW confidence because referenced by generated docs or reference files, used in event registration or dispatch. |
| OnLButtonDown | XML Handler | 23 | LOW | addons=1; sources=bindings,xml_handlers | Kept as LOW confidence because referenced by generated docs or reference files, used in event registration or dispatch. |
| OnLButtonUp | XML Handler | 23 | LOW | addons=1; sources=bindings,xml_handlers | Kept as LOW confidence because referenced by generated docs or reference files, used in event registration or dispatch. |
| OnMouseOver | XML Handler | 23 | LOW | addons=1; sources=bindings,xml_handlers | Kept as LOW confidence because referenced by generated docs or reference files, used in event registration or dispatch. |
| OnMouseOverEnd | XML Handler | 23 | LOW | addons=1; sources=bindings,xml_handlers | Kept as LOW confidence because referenced by generated docs or reference files, used in event registration or dispatch. |
| OnRButtonUp | XML Handler | 23 | LOW | addons=1; sources=bindings,xml_handlers | Kept as LOW confidence because referenced by generated docs or reference files, used in event registration or dispatch. |
| OnShown | XML Handler | 23 | LOW | addons=1; sources=bindings,xml_handlers | Kept as LOW confidence because referenced by generated docs or reference files, used in event registration or dispatch. |
| OnShutdown | XML Handler | 23 | LOW | addons=1; sources=bindings,xml_handlers | Kept as LOW confidence because referenced by generated docs or reference files, used in event registration or dispatch. |
| OnInitialize | XML Handler | 8 | LOW | addons=1; sources=bindings,xml_handlers | Kept as LOW confidence because referenced by generated docs or reference files, used in event registration or dispatch. |

## Rejected Addon Local

| Name | Category | Score | Level | Evidence | Rationale |
| --- | --- | --- | --- | --- | --- |
| RegisterEventHandler | Global Function | 23 | LOW | addons=1; global=9; sources=lua_calls | Rejected as platform API because never appears outside addon-local flow. |
| UnregisterEventHandler | Global Function | 23 | LOW | addons=1; global=7; sources=lua_calls | Rejected as platform API because never appears outside addon-local flow. |
| GetIconData | Global Function | 13 | LOW | addons=1; global=2; sources=lua_calls | Rejected as platform API because never appears outside addon-local flow. |
| CharacterWindow.CreateTooltip | Global Function | 5 | LOW | addons=1; global=3; sources=lua_calls | Rejected as platform API because never appears outside addon-local flow. |
| towstring | Global Function | 5 | LOW | addons=1; global=17; sources=lua_calls | Rejected as platform API because never appears outside addon-local flow. |
| DoesWindowExist | Global Function | 0 | LOW | addons=1; global=1; sources=lua_calls | Rejected as platform API because never appears outside addon-local flow, only one weak usage site. |

## Confidence Explanation Matrix

| Name | Category | Score | Level | Cross-addon | Namespace | XML | Event | Default UI | Local penalty | Signature |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| GameData.EquipSlots.BACK | GameData Field | 100 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | 0 |
| GameData.EquipSlots.HELM | GameData Field | 100 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | 0 |
| Label | XML Element Type | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| LabelSetText | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| LabelSetTextColor | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| LayoutEditor.RegisterWindow | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| SystemData.ActiveWindow.name | SystemData Field | 100 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | 0 |
| SystemData.Events.ALL_MODULES_INITIALIZED | SystemData Field | 100 | HIGH | 0 | +25 | 0 | +18 | +35 | 0 | 0 |
| SystemData.Events.LOADING_BEGIN | SystemData Field | 100 | HIGH | 0 | +25 | 0 | +18 | +35 | 0 | 0 |
| SystemData.Events.LOADING_END | SystemData Field | 100 | HIGH | 0 | +25 | 0 | +18 | +35 | 0 | 0 |
| SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED | SystemData Field | 100 | HIGH | 0 | +25 | 0 | +18 | +35 | 0 | 0 |
| SystemData.Events.PLAYER_ZONE_CHANGED | SystemData Field | 100 | HIGH | 0 | +25 | 0 | +18 | +35 | 0 | 0 |
| SystemData.Events.RELOAD_INTERFACE | SystemData Field | 100 | HIGH | 0 | +25 | 0 | +18 | +35 | 0 | 0 |
| SystemData.Events.UPDATE_PROCESSED | SystemData Field | 100 | HIGH | 0 | +25 | 0 | +18 | +35 | 0 | 0 |
| SystemData.Events.VISIBLE_EQUIPMENT_UPDATED | SystemData Field | 100 | HIGH | 0 | +25 | 0 | +18 | +35 | 0 | 0 |
| SystemData.MouseOverWindow.name | SystemData Field | 100 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | 0 |
| Window | XML Element Type | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| Windows | XML Element Type | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| Anchors | XML Element Type | 98 | HIGH | +18 | 0 | +30 | 0 | +35 | 0 | 0 |
| Color | XML Element Type | 98 | HIGH | +18 | 0 | +30 | 0 | +35 | 0 | 0 |
| Size | XML Element Type | 98 | HIGH | +18 | 0 | +30 | 0 | +35 | 0 | 0 |
| Button | XML Element Type | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| ButtonSetDisabledFlag | Window Function | 90 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | +10 |
| ButtonSetPressedFlag | Window Function | 90 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | +10 |
| ButtonSetText | Window Function | 90 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | +10 |
| DialogManager.MakeTextEntryDialog | Global Function | 90 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | +10 |
| DynamicImage | XML Element Type | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_Default | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_DefaultCheckBox | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_DefaultDown | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_DefaultResizeable | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_DefaultUp | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_DefaultWindowClose | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_ListSort | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Default_CharacterImage | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_FullResizeImage_TintableSolidBackground | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_ListSortDownArrow | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_ListSortUpArrow | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_TitleBar_Default | Constant | 90 | HIGH | 0 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Window_ContextMenu.AddMenuItem | Global Function | 90 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | +10 |

## Uncertain Areas

- CharacterWindow.CreateTooltip
- DoesWindowExist
- GetIconData
- OnHidden
- OnInitialize
- OnLButtonDown
- OnLButtonUp
- OnMouseOver
- OnMouseOverEnd
- OnRButtonUp
- OnShown
- OnShutdown
- RegisterEventHandler
- UnregisterEventHandler
- towstring
