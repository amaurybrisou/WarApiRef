# BroadcastEvent

- Category: events
- Confidence: MEDIUM

## Description

Observed triggering a runtime event so existing handlers are notified.

## Involved APIs

- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field

## Flow Diagram

```text
Label <-> OnHyperLinkRButtonUp
```

## Example Code

```lua
Lib RuString: BroadcastEvent(SystemData.Events.RELOAD_INTERFACE)
```

## Evidence

- Lib RuString: BroadcastEvent(SystemData.Events.RELOAD_INTERFACE)
