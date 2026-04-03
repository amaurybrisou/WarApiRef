# Conventions

## Naming

| Rule | Evidence |
| --- | --- |
| Handler functions use the On<Event> prefix | Clock.OnUpdate, ClosetGoblin.OnActivationWatchdog, ClosetGoblin.OnAllModulesInitialized, ClosetGoblin.OnInitialize, ClosetGoblin.OnInventorySlotUpdated, ClosetGoblin.OnShutdown, ClosetGoblin.OnSlashCommand, ClosetGoblin.OnVisibleEquipmentUpdated |
| Lua logic is namespaced through module-qualified function names | CM_ClosetGoblin.local.GetActionBarNumberFromWindowName, Clock.Initialize, Clock.OnUpdate, ClosetGoblin.ActivateSet, ClosetGoblin.AddNewSet, ClosetGoblin.Alert, ClosetGoblin.AssociateZoneToSet, ClosetGoblin.BuildItemList |
| Settings state is grouped under Settings-named modules or globals | ClockSettings |
| Top-level XML frames commonly use the Window suffix | ClockWindow, ClosetGoblinCharacterWindow, ClosetGoblinOptionWindow, ClosetGoblinZoneWindow |

## Common Events

- none

## Common Modules

- none

## Common State Owners

- none

## Common Frame Parents

- none
