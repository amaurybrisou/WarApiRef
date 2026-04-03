# State management pattern

- Category: conventions
- Confidence: MEDIUM

## Description

Persistent state is typically rooted in addon-owned globals and saved variables, then initialized before runtime hooks are attached.

## Involved APIs

- none

## Flow Diagram

```text
Label <-> OnHyperLinkRButtonUp
```

## Example Code

```lua
InfoScroller: InfoScroller.Settings
```

## Evidence

- InfoScroller: InfoScroller.Settings
- Lib RuString: LibRuString.Settings
- NPC Item Sale Price: Nisp.DebugEnabled
- NPC Item Sale Price: Nisp.DumpItemsTable
- NPC Item Sale Price: Nisp.Enabled
- PartyCast: PartyCast.Settings
- Soloq: Soloq.db
- TidyChat: TidyChat.Settings
