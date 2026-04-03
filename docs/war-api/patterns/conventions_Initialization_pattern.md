# Initialization pattern

- Category: conventions
- Confidence: HIGH

## Description

Addons consistently move from manifest loading into initialize hooks, then into XML creation and runtime event registration.

## Involved APIs

- [SystemData.Events.UPDATE_PROCESSED](../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event

## Flow Diagram

```text
Label <-> OnHyperLinkRButtonUp
```

## Example Code

```lua
initialize: Creates 0 windows and calls 1 initialize hooks, Creates 3 windows and calls 1 initialize hooks
```

## Evidence

- initialize: Creates 0 windows and calls 1 initialize hooks, Creates 3 windows and calls 1 initialize hooks
- runtime: InfoScroller: SystemData.Events.UPDATE_PROCESSED, InfoScroller: e
- xml: Defines 20 XML frames and 0 bound handlers, Defines 39 XML frames and 4 bound handlers
