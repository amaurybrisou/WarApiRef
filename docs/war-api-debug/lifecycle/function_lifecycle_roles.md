# Function Lifecycle Roles

> Lua functions with lifecycle roles derived from `.mod` manifest analysis.

## Startup Entrypoints (OnInitialize)

| Function | Ref Name | Addon | Resolution | Confidence |
| --- | --- | --- | --- | --- |
| `ClosetGoblin.Initialize` | CG_ItemRack.Initialize | CM_ClosetGoblin | ambiguous | MEDIUM |
| `Clock.Initialize` | CG_ItemRack.Initialize | CM_ClosetGoblin | ambiguous | MEDIUM |
| `ClosetGoblin.Initialize` | Clock.Initialize | Clock | ambiguous | MEDIUM |
| `Clock.Initialize` | Clock.Initialize | Clock | ambiguous | MEDIUM |
| `ClosetGoblin.OnInitialize` | ClosetGoblin.OnInitialize | CM_ClosetGoblin | ambiguous | MEDIUM |
| `ClosetGoblinCharacterWindow.OnInitialize` | ClosetGoblin.OnInitialize | CM_ClosetGoblin | ambiguous | MEDIUM |
| `ClosetGoblinOptionWindow.OnInitialize` | ClosetGoblin.OnInitialize | CM_ClosetGoblin | ambiguous | MEDIUM |
| `ClosetGoblinZoneWindow.OnInitialize` | ClosetGoblin.OnInitialize | CM_ClosetGoblin | ambiguous | MEDIUM |
## Shutdown Entrypoints (OnShutdown)

| Function | Ref Name | Addon | Resolution | Confidence |
| --- | --- | --- | --- | --- |
| `ClosetGoblin.OnShutdown` | ClosetGoblin.OnShutdown | CM_ClosetGoblin | ambiguous | MEDIUM |
| `ClosetGoblinCharacterWindow.OnShutdown` | ClosetGoblin.OnShutdown | CM_ClosetGoblin | ambiguous | MEDIUM |
| `ClosetGoblinZoneWindow.OnShutdown` | ClosetGoblin.OnShutdown | CM_ClosetGoblin | ambiguous | MEDIUM |
## Update Callbacks (OnUpdate)

| Function | Ref Name | Addon | Resolution | Confidence |
| --- | --- | --- | --- | --- |
| `Clock.OnUpdate` | Clock.OnUpdate | Clock | exact | HIGH |
