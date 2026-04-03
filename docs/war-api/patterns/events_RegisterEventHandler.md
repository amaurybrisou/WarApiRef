# RegisterEventHandler

- Category: events
- Confidence: HIGH

## Description

Observed registering global runtime handlers against shared event identifiers.

## Involved APIs

- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function
- [EventHandler](../xml/element_types/element_EventHandler.md) (MEDIUM 45/100) - XML Element Type

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ColorPicker, DynamicImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
AdvancedPetAssist: RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
```

## Evidence

- AdvancedPetAssist: RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- AdvancedPetAssist: RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.LOADING_END, "AdvancedRenownTraining.OnReload")
- AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
- AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "AdvancedRenownTraining.OnReload")
- AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
