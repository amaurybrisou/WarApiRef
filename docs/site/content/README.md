# WAR Addon Development API Reference

This documentation is the platform-facing second layer derived from the generated addon corpus under docs/addon-api. It attempts to answer a narrower question: which APIs, events, XML hooks, tables, and lifecycle patterns are shared across WAR addons.

## Method

- Input: generated markdown under docs/addon-api only.

- Output: canonical platform symbols with explicit confidence score, level, evidence signals, and rationale.

- Scope: WAR-exposed functions, window APIs, XML handlers, runtime events, SystemData, GameData, lifecycle, and slash-command patterns.

## Confidence

- Score range: 0 to 100 after clamping the weighted raw signal total.

- HIGH: 70 to 100 and promoted automatically into the canonical platform surface.

- MEDIUM: 40 to 69 and promoted only when supported by namespace, default UI, event-binding, or direct XML evidence.

- LOW: 0 to 39 and retained as non-canonical evidence only.

## Start Here

- [Overview](index.md)

- [Global functions](globals/functions/index.md)

- [Window API](window_api/functions/index.md)

- [XML handlers](xml/handlers/index.md)

- [Game events](events/game_events/index.md)

- [Semantic patterns](patterns/index.md)

- [Navigation tree](navigation/tree.json)

- [API graph](graph/api_graph.json)

- [Coverage report](meta/coverage_report.md)
