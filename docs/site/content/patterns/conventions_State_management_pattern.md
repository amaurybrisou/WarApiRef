# State management pattern

- Category: conventions
- Confidence: MEDIUM

## Description

Persistent state is typically rooted in addon-owned globals and saved variables, then initialized before runtime hooks are attached.

## Involved APIs

- none

## Flow Diagram

```text
OnLButtonUp <-> OnLButtonUp
```

## Example Code

```lua
NPC Item Sale Price: Nisp.DebugEnabled
```

## Evidence

- NPC Item Sale Price: Nisp.DebugEnabled
- NPC Item Sale Price: Nisp.DumpItemsTable
- NPC Item Sale Price: Nisp.Enabled
- TidyChat: TidyChat.Settings
