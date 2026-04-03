# Initialization pattern

- Category: conventions
- Confidence: HIGH

## Description

Addons consistently move from manifest loading into initialize hooks, then into XML creation and runtime event registration.

## Involved APIs

- none

## Flow Diagram

```text
SystemData.ActiveWindow.name <-> WindowGetId
```

## Evidence

- none
