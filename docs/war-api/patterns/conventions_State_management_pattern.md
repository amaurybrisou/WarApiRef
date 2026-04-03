# State management pattern

- Category: conventions
- Confidence: MEDIUM

## Description

Persistent state is typically rooted in addon-owned globals and saved variables, then initialized before runtime hooks are attached.

## Involved APIs

- none

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ColorPicker, DynamicImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
AdvancedPetAssist: APA_Settings
```

## Evidence

- AdvancedPetAssist: APA_Settings
- AdvancedRenownTrainer: AdvancedRenownTraining.Presets
- AggroMeter: AggroMeter.Settings
- BagOMatic: BagOMatic.saved
- BankArkel: BankArkel.db
- BuffHead: BuffHead.Settings
- CM_ClosetGoblin: ClosetGoblin.setData
- CM_ClosetGoblin: ClosetGoblin.settings
