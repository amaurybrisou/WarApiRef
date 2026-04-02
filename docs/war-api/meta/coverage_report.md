# Coverage Report

## Source Counts

| Input | Count |
| --- | --- |
| bindings | 40 |
| event_docs | 18 |
| examples | 12 |
| flows | 77 |
| frame_docs | 148 |
| function_docs | 232 |
| handler_docs | 40 |

## Symbol Counts

| Symbol Type | Count |
| --- | --- |
| constants | 27 |
| element_types | 20 |
| game_events | 7 |
| gamedata | 3 |
| global_functions | 19 |
| systemdata | 13 |
| tables | 6 |
| window_events | 11 |
| window_functions | 47 |
| xml_handlers | 10 |

## Confidence

| Level | Count |
| --- | --- |
| HIGH | 145 |
| MEDIUM | 21 |
| LOW | 2 |

## Candidate Outcomes

| Outcome | Count |
| --- | --- |
| High confidence platform | 145 |
| Medium confidence candidates | 21 |
| Low confidence symbols | 0 |
| Rejected addon-local | 2 |

## Spread

- Symbols seen once: 98

- Symbols seen in multiple addons: 70

## High Confidence Platform

| Name | Category | Score | Level | Evidence | Rationale |
| --- | --- | --- | --- | --- | --- |
| EA_Button_DefaultCheckBox | Constant | 100 | HIGH | addons=2; xml=2; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_DefaultResizeable | Constant | 100 | HIGH | addons=2; xml=4; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_DefaultWindowClose | Constant | 100 | HIGH | addons=2; xml=4; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_ResizeIconFrameNormal | Constant | 100 | HIGH | addons=3; xml=3; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_Tab | Constant | 100 | HIGH | addons=2; xml=2; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_FullResizeImage_BlackTransparent | Constant | 100 | HIGH | addons=2; xml=2; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_FullResizeImage_TintableSolidBackground | Constant | 100 | HIGH | addons=3; xml=7; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Label_DefaultText | Constant | 100 | HIGH | addons=2; xml=3; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Window_DefaultContextMenuFrame | Constant | 100 | HIGH | addons=2; xml=3; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM | Game Event | 100 | HIGH | addons=1; global=1; sources=event_page,flows,lua_event_registration | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA | Game Event | 100 | HIGH | addons=1; global=1; sources=event_page,flows,lua_event_registration | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.LOADING_END | Game Event | 100 | HIGH | addons=2; global=2; sources=event_page,flows,lua_event_registration | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.L_BUTTON_UP_PROCESSED | Game Event | 100 | HIGH | addons=1; global=1; sources=event_page,flows,lua_event_registration | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.PLAYER_TARGET_UPDATED | Game Event | 100 | HIGH | addons=1; global=1; sources=event_page,flows,lua_event_registration | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.RELOAD_INTERFACE | Game Event | 100 | HIGH | addons=2; global=2; sources=event_page,flows,lua_event_registration | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| GameData.ItemTypes.CURRENCY | GameData Field | 100 | HIGH | addons=1; global=3; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| GameData.ItemTypes.ENHANCEMENT | GameData Field | 100 | HIGH | addons=1; global=3; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| GameData.ItemTypes.POTION | GameData Field | 100 | HIGH | addons=1; global=3; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| EA_ChatWindow.OnSettingsChanged | Global Function | 100 | HIGH | addons=1; global=1; sources=globals,lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| EA_ChatWindow.Print | Global Function | 100 | HIGH | addons=1; global=19; sources=globals,lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| EA_ChatWindow.ShowTabCycleButtons | Global Function | 100 | HIGH | addons=1; global=1; sources=globals,lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| EA_ChatWindow.UpdateTabScrollWindow | Global Function | 100 | HIGH | addons=1; global=1; sources=globals,lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| LibSlash.RegisterWSlashCmd | Global Function | 100 | HIGH | addons=2; global=2; sources=globals,lua_calls | Promoted as HIGH confidence because matches a known engine namespace, referenced by generated docs or reference files, called globally with no local definition. |
| DefaultColor | Global Table | 100 | HIGH | addons=2; global=1; local=2; sources=globals,lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| EA_ChatWindow | Global Table | 100 | HIGH | addons=2; global=4; local=1; sources=globals,lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| LibSlash | Global Table | 100 | HIGH | addons=4; global=2; local=1; sources=globals,lua_calls | Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.ActiveWindow.name | SystemData Field | 100 | HIGH | addons=3; global=5; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.ChatLogFilters.ADVICE | SystemData Field | 100 | HIGH | addons=1; global=2; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.ChatLogFilters.ALLIANCE | SystemData Field | 100 | HIGH | addons=1; global=2; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.ChatLogFilters.LOOT_ROLL | SystemData Field | 100 | HIGH | addons=1; global=3; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.ChatLogFilters.SCENARIO | SystemData Field | 100 | HIGH | addons=1; global=2; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM | SystemData Field | 100 | HIGH | addons=1; global=8; sources=event_page,event_registration,flow | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA | SystemData Field | 100 | HIGH | addons=1; global=9; sources=event_page,event_registration,flow | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.LOADING_END | SystemData Field | 100 | HIGH | addons=3; global=11; sources=event_page,event_registration,flow | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.L_BUTTON_UP_PROCESSED | SystemData Field | 100 | HIGH | addons=1; global=8; sources=event_page,event_registration,flow | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.PLAYER_TARGET_UPDATED | SystemData Field | 100 | HIGH | addons=1; global=8; sources=event_page,event_registration,flow | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.Events.RELOAD_INTERFACE | SystemData Field | 100 | HIGH | addons=3; global=11; sources=event_page,event_registration,flow | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.ItemRarity.UTILITY | SystemData Field | 100 | HIGH | addons=1; global=3; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| SystemData.MouseOverWindow.name | SystemData Field | 100 | HIGH | addons=1; global=2; sources=lua_call | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files. |
| OnHidden | Window Event | 100 | HIGH | addons=2; xml=6; global=1; sources=event_page,examples,flows | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, called globally with no local definition. |
| OnLButtonUp | Window Event | 100 | HIGH | addons=2; xml=19; global=4; sources=event_page,examples,flows | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, called globally with no local definition. |
| OnMButtonUp | Window Event | 100 | HIGH | addons=2; xml=1; global=1; sources=event_page,flows,lua_event_registration | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, called globally with no local definition. |
| OnMouseWheel | Window Event | 100 | HIGH | addons=2; xml=2; sources=event_page,examples,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, reinforced across multiple generated source types. |
| OnRButtonUp | Window Event | 100 | HIGH | addons=2; xml=1; global=1; sources=event_page,flows,lua_event_registration | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, called globally with no local definition. |
| OnSelChanged | Window Event | 100 | HIGH | addons=2; xml=2; sources=event_page,examples,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, reinforced across multiple generated source types. |
| OnShown | Window Event | 100 | HIGH | addons=2; xml=6; global=1; sources=event_page,examples,flows | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, called globally with no local definition. |
| ButtonGetDisabledFlag | Window Function | 100 | HIGH | addons=2; global=5; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ButtonGetPressedFlag | Window Function | 100 | HIGH | addons=2; global=28; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ButtonSetPressedFlag | Window Function | 100 | HIGH | addons=2; global=21; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ButtonSetText | Window Function | 100 | HIGH | addons=2; global=13; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ComboBoxGetSelectedMenuItem | Window Function | 100 | HIGH | addons=2; global=10; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ComboBoxSetSelectedMenuItem | Window Function | 100 | HIGH | addons=2; global=9; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| LabelSetFont | Window Function | 100 | HIGH | addons=2; global=6; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| LabelSetText | Window Function | 100 | HIGH | addons=3; global=46; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| LayoutEditor.RegisterWindow | Window Function | 100 | HIGH | addons=3; global=3; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| TextEditBoxSetText | Window Function | 100 | HIGH | addons=2; global=7; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowAddAnchor | Window Function | 100 | HIGH | addons=3; global=90; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowClearAnchors | Window Function | 100 | HIGH | addons=2; global=15; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowGetDimensions | Window Function | 100 | HIGH | addons=2; global=3; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowGetId | Window Function | 100 | HIGH | addons=2; global=7; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowGetShowing | Window Function | 100 | HIGH | addons=2; global=10; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowRegisterCoreEventHandler | Window Function | 100 | HIGH | addons=2; global=9; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowSetHandleInput | Window Function | 100 | HIGH | addons=2; global=8; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowSetId | Window Function | 100 | HIGH | addons=2; global=10; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowSetLayer | Window Function | 100 | HIGH | addons=2; global=12; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowSetShowing | Window Function | 100 | HIGH | addons=3; global=50; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| Button | XML Element Type | 100 | HIGH | addons=2; xml=25; sources=xml_frames,xml_handlers | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| ComboBox | XML Element Type | 100 | HIGH | addons=2; xml=5; sources=xml_frames,xml_handlers | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| DynamicImage | XML Element Type | 100 | HIGH | addons=2; xml=10; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EditBox | XML Element Type | 100 | HIGH | addons=2; xml=3; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| FullResizeImage | XML Element Type | 100 | HIGH | addons=3; xml=17; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| Label | XML Element Type | 100 | HIGH | addons=3; xml=32; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| ListBox | XML Element Type | 100 | HIGH | addons=2; xml=2; sources=xml_frames,xml_handlers | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| ListColumn | XML Element Type | 100 | HIGH | addons=2; xml=2; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| ListColumns | XML Element Type | 100 | HIGH | addons=2; xml=2; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| ListData | XML Element Type | 100 | HIGH | addons=2; xml=2; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| Window | XML Element Type | 100 | HIGH | addons=3; xml=48; sources=xml_frames,xml_handlers | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| OnHidden | XML Handler | 100 | HIGH | addons=2; xml=6; sources=bindings,examples,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths. |
| OnLButtonUp | XML Handler | 100 | HIGH | addons=2; xml=19; sources=bindings,examples,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths. |
| OnMouseWheel | XML Handler | 100 | HIGH | addons=2; xml=2; sources=bindings,examples,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths. |
| OnSelChanged | XML Handler | 100 | HIGH | addons=2; xml=2; sources=bindings,examples,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths. |
| OnShown | XML Handler | 100 | HIGH | addons=2; xml=6; sources=bindings,examples,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths. |
| WindowRegisterEventHandler | Window Function | 98 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowSetDimensions | Window Function | 98 | HIGH | addons=3; global=14; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowSetTintColor | Window Function | 98 | HIGH | addons=3; global=8; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| EA_Button_DefaultToggleCircle | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_ResizeIconFrameHighlight | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Button_ResizeIconFramePressed | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_ComboBox_DefaultResizable | Constant | 90 | HIGH | addons=1; xml=4; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_ComboBox_DefaultResizableSmall | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Default_SliderBar | Constant | 90 | HIGH | addons=1; xml=4; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_EditBox_DefaultFrame | Constant | 90 | HIGH | addons=1; xml=2; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_EditBox_DefaultFrame_Multiline | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_FullResizeImage_DefaultFrame | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_FullResizeImage_MetalFill | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_FullResizeImage_TanBorder | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Image_DefaultIcon | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Image_DefaultIconFrame | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_TitleBar_Default | Constant | 90 | HIGH | addons=1; xml=2; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Window_DefaultBackgroundFrame | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Window_DefaultButtonBottomFrame | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Window_DefaultFrame | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Window_DefaultFrameStatusBar | Constant | 90 | HIGH | addons=1; xml=1; sources=xml_attributes | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| EA_Window_ContextMenu.AddMenuItem | Global Function | 90 | HIGH | addons=1; global=2; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ButtonSetDisabledFlag | Window Function | 90 | HIGH | addons=1; global=4; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ButtonSetTextColor | Window Function | 90 | HIGH | addons=1; global=12; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ComboBoxAddMenuItem | Window Function | 90 | HIGH | addons=1; global=4; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ComboBoxClearMenuItems | Window Function | 90 | HIGH | addons=1; global=7; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| DynamicImageSetTexture | Window Function | 90 | HIGH | addons=1; global=2; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| LabelSetTextColor | Window Function | 90 | HIGH | addons=1; global=2; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| TextEditBoxGetText | Window Function | 90 | HIGH | addons=1; global=4; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowSetAlpha | Window Function | 90 | HIGH | addons=1; global=5; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowSetParent | Window Function | 90 | HIGH | addons=1; global=2; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowStopAlphaAnimation | Window Function | 90 | HIGH | addons=1; global=2; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowUnregisterCoreEventHandler | Window Function | 90 | HIGH | addons=1; global=2; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| SliderBar | XML Element Type | 90 | HIGH | addons=1; xml=4; sources=xml_frames | Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace. |
| RegisterEventHandler | Global Function | 81 | HIGH | addons=2; global=10; sources=lua_calls | Promoted as HIGH confidence because called globally with no local definition, seen in 2 to 3 addons, used in event registration or dispatch. |
| DialogManager.MakeOneButtonDialog | Global Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| EA_Window_ContextMenu.Finalize | Global Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| LibSlash.RegisterSlashCmd | Global Function | 80 | HIGH | addons=1; global=1; sources=globals,lua_calls | Promoted as HIGH confidence because matches a known engine namespace, referenced by generated docs or reference files, called globally with no local definition. |
| DialogManager | Global Table | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| EA_Window_ContextMenu | Global Table | 80 | HIGH | addons=1; global=2; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ComboBoxGetSelectedText | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| ComboBoxSetDisabledFlag | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| DynamicImageSetTextureScale | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| DynamicImageSetTextureSlice | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| LabelGetTextDimensions | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| LayoutEditor.UnregisterWindow | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| StatusBarSetBackgroundTint | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| StatusBarSetCurrentValue | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| StatusBarSetForegroundTint | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| StatusBarSetMaximumValue | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowAssignFocus | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowGetAnchor | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| WindowGetOffsetFromParent | Window Function | 80 | HIGH | addons=1; global=1; sources=lua_calls | Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition. |
| wstring | Global Table | 78 | HIGH | addons=2; global=6; sources=lua_calls | Promoted as HIGH confidence because matches a known engine namespace, called globally with no local definition, seen in 2 to 3 addons. |
| OnLButtonDown | Window Event | 73 | HIGH | addons=1; global=1; sources=event_page,flows,lua_event_registration | Promoted as HIGH confidence because referenced by generated docs or reference files, called globally with no local definition, reinforced across multiple generated source types. |
| OnInitialize | XML Handler | 73 | HIGH | addons=1; xml=1; sources=bindings,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths. |
| OnMButtonUp | XML Handler | 73 | HIGH | addons=1; xml=1; sources=bindings,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths. |
| OnMouseOver | XML Handler | 73 | HIGH | addons=1; xml=1; sources=bindings,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths. |
| OnRButtonUp | XML Handler | 73 | HIGH | addons=1; xml=1; sources=bindings,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths. |
| OnUpdate | XML Handler | 73 | HIGH | addons=1; xml=1; sources=bindings,xml_handlers | Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths. |
| DoesWindowExist | Global Function | 71 | HIGH | addons=2; global=4; sources=lua_calls | Promoted as HIGH confidence because called globally with no local definition, seen in 2 to 3 addons, role is consistent across addons. |
| GetIconData | Global Function | 71 | HIGH | addons=2; global=3; sources=lua_calls | Promoted as HIGH confidence because called globally with no local definition, seen in 2 to 3 addons, role is consistent across addons. |
| DefaultColor.SetWindowTint | Global Function | 70 | HIGH | addons=1; global=1; sources=globals,lua_calls | Promoted as HIGH confidence because matches a known engine namespace, referenced by generated docs or reference files, called globally with no local definition. |

## Medium Confidence Candidates

| Name | Category | Score | Level | Evidence | Rationale |
| --- | --- | --- | --- | --- | --- |
| TextLogGetUpdateEventId | Game Event | 63 | MEDIUM | addons=1; global=3; sources=event_page,lua_event_registration | Promoted as MEDIUM confidence because referenced by generated docs or reference files, called globally with no local definition, used in event registration or dispatch. |
| CreateWindow | Global Function | 63 | MEDIUM | addons=3; global=7; sources=lua_calls | Kept as MEDIUM confidence because called globally with no local definition, seen in 2 to 3 addons, role is consistent across addons. |
| CreateWindowFromTemplate | Global Function | 63 | MEDIUM | addons=2; global=54; sources=lua_calls | Kept as MEDIUM confidence because called globally with no local definition, seen in 2 to 3 addons, role is consistent across addons. |
| towstring | Global Function | 63 | MEDIUM | addons=3; global=23; sources=lua_calls | Kept as MEDIUM confidence because called globally with no local definition, seen in 2 to 3 addons, role is consistent across addons. |
| wstring.find | Global Function | 55 | MEDIUM | addons=1; global=3; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, argument pattern is consistent. |
| OnInitialize | Window Event | 53 | MEDIUM | addons=1; xml=1; sources=event_page,xml_handlers | Promoted as MEDIUM confidence because used directly in xml handler attributes, referenced by generated docs or reference files, used in event registration or dispatch. |
| OnMouseOver | Window Event | 53 | MEDIUM | addons=1; xml=1; sources=event_page,xml_handlers | Promoted as MEDIUM confidence because used directly in xml handler attributes, referenced by generated docs or reference files, used in event registration or dispatch. |
| OnUpdate | Window Event | 53 | MEDIUM | addons=1; xml=1; sources=event_page,xml_handlers | Promoted as MEDIUM confidence because used directly in xml handler attributes, referenced by generated docs or reference files, used in event registration or dispatch. |
| wstring.gsub | Global Function | 45 | MEDIUM | addons=1; global=1; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition. |
| wstring.lower | Global Function | 45 | MEDIUM | addons=1; global=1; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition. |
| wstring.match | Global Function | 45 | MEDIUM | addons=1; global=1; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition. |
| wstring.upper | Global Function | 45 | MEDIUM | addons=1; global=1; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition. |
| AnimFrame | XML Element Type | 45 | MEDIUM | addons=1; xml=2; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| AnimFrames | XML Element Type | 45 | MEDIUM | addons=1; xml=2; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| AnimatedImage | XML Element Type | 45 | MEDIUM | addons=1; xml=2; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| Disabled | XML Element Type | 45 | MEDIUM | addons=1; xml=1; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| Normal | XML Element Type | 45 | MEDIUM | addons=1; xml=1; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| NormalHighlit | XML Element Type | 45 | MEDIUM | addons=1; xml=1; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| Pressed | XML Element Type | 45 | MEDIUM | addons=1; xml=1; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| TintColor | XML Element Type | 45 | MEDIUM | addons=1; xml=4; sources=xml_frames | Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes. |
| wstring.sub | Global Function | 40 | MEDIUM | addons=1; global=3; sources=lua_calls | Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, known namespace override. |

## Low Confidence Symbols

- none

## Rejected Addon Local

| Name | Category | Score | Level | Evidence | Rationale |
| --- | --- | --- | --- | --- | --- |
| UnregisterEventHandler | Global Function | 23 | LOW | addons=1; global=3; sources=lua_calls | Rejected as platform API because never appears outside addon-local flow. |
| DestroyWindow | Global Function | 0 | LOW | addons=1; global=1; sources=lua_calls | Rejected as platform API because never appears outside addon-local flow, only one weak usage site. |

## Confidence Explanation Matrix

| Name | Category | Score | Level | Cross-addon | Namespace | XML | Event | Default UI | Local penalty | Signature |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Button | XML Element Type | 100 | HIGH | +18 | +25 | +50 | +18 | +35 | 0 | 0 |
| ButtonGetDisabledFlag | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| ButtonGetPressedFlag | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| ButtonSetPressedFlag | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| ButtonSetText | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| ComboBox | XML Element Type | 100 | HIGH | +18 | +25 | +50 | +18 | +35 | 0 | 0 |
| ComboBoxGetSelectedMenuItem | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| ComboBoxSetSelectedMenuItem | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| DefaultColor | Global Table | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | 0 |
| DynamicImage | XML Element Type | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_DefaultCheckBox | Constant | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_DefaultResizeable | Constant | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_DefaultWindowClose | Constant | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_ResizeIconFrameNormal | Constant | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Button_Tab | Constant | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_ChatWindow | Global Table | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | 0 |
| EA_ChatWindow.OnSettingsChanged | Global Function | 100 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | 0 |
| EA_ChatWindow.Print | Global Function | 100 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | +10 |
| EA_ChatWindow.ShowTabCycleButtons | Global Function | 100 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | 0 |
| EA_ChatWindow.UpdateTabScrollWindow | Global Function | 100 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | 0 |
| EA_FullResizeImage_BlackTransparent | Constant | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_FullResizeImage_TintableSolidBackground | Constant | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Label_DefaultText | Constant | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| EA_Window_DefaultContextMenuFrame | Constant | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| EditBox | XML Element Type | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| FullResizeImage | XML Element Type | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| GameData.ItemTypes.CURRENCY | GameData Field | 100 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | 0 |
| GameData.ItemTypes.ENHANCEMENT | GameData Field | 100 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | 0 |
| GameData.ItemTypes.POTION | GameData Field | 100 | HIGH | 0 | +25 | 0 | 0 | +35 | 0 | 0 |
| Label | XML Element Type | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| LabelSetFont | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| LabelSetText | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| LayoutEditor.RegisterWindow | Window Function | 100 | HIGH | +18 | +25 | 0 | 0 | +35 | 0 | +10 |
| LibSlash | Global Table | 100 | HIGH | +30 | +25 | 0 | 0 | 0 | 0 | 0 |
| LibSlash.RegisterWSlashCmd | Global Function | 100 | HIGH | +18 | +25 | 0 | 0 | 0 | 0 | +10 |
| ListBox | XML Element Type | 100 | HIGH | +18 | +25 | +50 | +18 | +35 | 0 | 0 |
| ListColumn | XML Element Type | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| ListColumns | XML Element Type | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| ListData | XML Element Type | 100 | HIGH | +18 | +25 | +30 | 0 | +35 | 0 | 0 |
| OnHidden | XML Handler | 100 | HIGH | +18 | 0 | +50 | +18 | 0 | 0 | 0 |

## Uncertain Areas

- DestroyWindow
- UnregisterEventHandler
