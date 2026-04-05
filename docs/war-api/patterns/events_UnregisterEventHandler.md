# UnregisterEventHandler

- Category: events
- Confidence: HIGH

## Description

Observed removing previously registered global runtime handlers.

## Involved APIs

- [SystemData.Events](../systemdata/fields/systemdata_SystemData.Events.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [OnUpdate](../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [UnregisterEventHandler](../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [EventHandler](../xml/element_types/element_EventHandler.md) (MEDIUM 45/100) - XML Element Type

## Flow Diagram

```text
SystemData.Events.ENTER_WORLD <-> SystemData.Events.INTERFACE_RELOADED
```

## Example Code

```lua
Ace: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
```

## Evidence

- Ace: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
- ActionFraction: UnregisterEventHandler(SystemData.Events.ENTER_WORLD, "ActionFraction.Initialize")
- ActionFraction: UnregisterEventHandler(SystemData.Events.INTERFACE_RELOADED, "ActionFraction.Initialize")
- ActionFraction: UnregisterEventHandler(SystemData.Events.LOADING_END, "ActionFractionWindow.UpdateVisibility")
- AdvancedPetAssist: UnregisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- AdvancedRenownTrainer: UnregisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
