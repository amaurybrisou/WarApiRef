# Initialization pattern

- Category: conventions
- Confidence: HIGH

## Description

Addons consistently move from manifest loading into initialize hooks, then into XML creation and runtime event registration.

## Involved APIs

- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../events/game_events/game_event_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.UPDATE_PROCESSED](../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ColorPicker, DynamicImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
initialize: AdvancedPetAssist: APAFollowTargetHUD, AdvancedPetAssist: APAInstantOnlyHUD
```

## Evidence

- initialize: AdvancedPetAssist: APAFollowTargetHUD, AdvancedPetAssist: APAInstantOnlyHUD
- runtime: Ace: SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED, Ace: SystemData.Events.UPDATE_PROCESSED
- xml: AdvancedPetAssist: APAComboAttackBind, AdvancedPetAssist: APAComboAutoReattack
