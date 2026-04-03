# UnregisterEventHandler

- Category: events
- Confidence: HIGH

## Description

Observed removing previously registered global runtime handlers.

## Involved APIs

- [OnUpdate](../xml/handlers/handler_OnUpdate.md) (HIGH 100/100) - XML Event
- [OnUpdate](../events/window_events/window_event_OnUpdate.md) (HIGH 100/100) - Window Event
- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.UPDATE_PROCESSED](../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.UPDATE_PROCESSED](../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [Text](../xml/element_types/element_Text.md) (HIGH 98/100) - XML Element Type
- [TextLogGetUpdateEventId](../events/game_events/game_event_TextLogGetUpdateEventId.md) (HIGH 93/100) - Game Event
- [UnregisterEventHandler](../globals/functions/global_UnregisterEventHandler.md) (HIGH 93/100) - Global Function
- [EventHandler](../xml/element_types/element_EventHandler.md) (MEDIUM 45/100) - XML Element Type

## Flow Diagram

```text
OnUpdate
  -> ui: Window
```

## Example Code

```lua
Ace: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
```

## Evidence

- Ace: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
- AdvancedPetAssist: UnregisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- AdvancedPetAssist: UnregisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- AdvancedRenownTrainer: UnregisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
- AggroMeter: UnregisterEventHandler(TextLogGetUpdateEventId("Chat"), "AggroMeter.OnChatLogUpdated")
- Aura: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH")
