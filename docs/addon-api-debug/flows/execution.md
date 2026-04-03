# Execution Flow

## CM_ClosetGoblin

- manifest: Loads 17 files in manifest order [ClosetGoblin.lua, ClosetGoblinCharacterWindow.lua, ClosetGoblinZoneWindow.lua, ClosetGoblinOptionWindow.lua, Libs/LibToolkit.lua, languages/english.lua, languages/russian.lua, languages/italian.lua]
- initialize: Creates 3 windows and calls 2 initialize hooks [CG_ItemRack.Initialize, ClosetGoblin.OnInitialize, ClosetGoblinCharacterWindow, ClosetGoblinOptionWindow, ClosetGoblinZoneWindow]
- xml: Defines 165 XML frames and 65 bound handlers [CG_ItemRackEQShow1, CG_ItemRackEQShow1Equipment, CG_ItemRackEQShow1EquipmentSlot1, CG_ItemRackEQShow1EquipmentSlot10, CG_ItemRackEQShow1EquipmentSlot2, CG_ItemRackEQShow1EquipmentSlot3, CG_ItemRackEQShow1EquipmentSlot4, CG_ItemRackEQShow1EquipmentSlot5]
- runtime: Registers 9 runtime events [SystemData.Events.ALL_MODULES_INITIALIZED, SystemData.Events.LOADING_BEGIN, SystemData.Events.LOADING_END, SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED, SystemData.Events.PLAYER_ZONE_CHANGED, SystemData.Events.RELOAD_INTERFACE, SystemData.Events.UPDATE_PROCESSED, SystemData.Events.VISIBLE_EQUIPMENT_UPDATED]
- shutdown: Executes 1 shutdown hooks [ClosetGoblin.OnShutdown]

## Clock

- manifest: Loads 1 files in manifest order [Clock.xml]
- initialize: Creates 0 windows and calls 1 initialize hooks [Clock.Initialize]
- xml: Defines 2 XML frames and 0 bound handlers [ClockWindow, ClockWindowText]
- update: Runs 1 update hooks from the manifest [Clock.OnUpdate]
