# Initialization pattern

- Category: conventions
- Confidence: HIGH

## Description

Addons consistently move from manifest loading into initialize hooks, then into XML creation and runtime event registration.

## Involved APIs

- [SystemData.Events.PLAYER_TARGET_UPDATED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../events/game_events/game_event_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - Game Event

## Flow Diagram

```text
Button <-> OnLButtonUp
```

## Example Code

```lua
initialize: Creates 0 windows and calls 1 initialize hooks, Creates 3 windows and calls 1 initialize hooks
```

## Evidence

- initialize: Creates 0 windows and calls 1 initialize hooks, Creates 3 windows and calls 1 initialize hooks
- runtime: Moth: SystemData.Events.PLAYER_TARGET_UPDATED, Registers 1 runtime events
- xml: Defines 20 XML frames and 0 bound handlers, Defines 47 XML frames and 22 bound handlers
