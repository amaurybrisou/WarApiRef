# Event registration pattern

- Category: conventions
- Confidence: HIGH

## Description

Runtime events are typically wired through RegisterEventHandler or window-scoped variants, with handler names staying module-qualified.

## Involved APIs

- [OnUpdate](../events/window_events/window_event_OnUpdate.md) (HIGH 100/100) - Window Event
- [OnUpdate](../xml/handlers/handler_OnUpdate.md) (HIGH 100/100) - XML Handler
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.UPDATE_PROCESSED](../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Flow Diagram

```text
OnUpdate
  -> ui: MapDisplay, Window
```

## Example Code

```lua
RegisterEventHandler: AdvancedPetAssist: RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
```

## Evidence

- RegisterEventHandler: AdvancedPetAssist: RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- RegisterEventHandler: AdvancedPetAssist: RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- RegisterEventHandler: AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.LOADING_END, "AdvancedRenownTraining.OnReload")
- RegisterEventHandler: AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
- RegisterEventHandler: AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "AdvancedRenownTraining.OnReload")
- RegisterEventHandler: AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
- UnregisterEventHandler: Ace: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
- UnregisterEventHandler: AdvancedPetAssist: UnregisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
