# CM_ClosetGoblin Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnShutdown

## Saved Variables

- ClosetGoblin.setData
- ClosetGoblin.settings

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | CG_ItemRack.Initialize | ambiguous | MEDIUM | ClosetGoblin.Initialize, Clock.Initialize |
| CallFunction | ClosetGoblin.OnInitialize | ambiguous | MEDIUM | ClosetGoblin.OnInitialize, ClosetGoblinCharacterWindow.OnInitialize, ClosetGoblinOptionWindow.OnInitialize, ClosetGoblinZoneWindow.OnInitialize |
| CreateWindow | ClosetGoblinCharacterWindow | exact | HIGH | ClosetGoblinCharacterWindow |
| CreateWindow | ClosetGoblinZoneWindow | exact | HIGH | ClosetGoblinZoneWindow |
| CreateWindow | ClosetGoblinOptionWindow | exact | HIGH | ClosetGoblinOptionWindow |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | ClosetGoblin.OnShutdown | ambiguous | MEDIUM | ClosetGoblin.OnShutdown, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EASystem_DialogManager | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EA_Cursor | unresolved |  | — |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EA_CharacterWindow | unresolved |  | — |
| Dependency | EA_ContextMenu | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EASystem_DialogManager | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EA_Cursor | unresolved |  | — |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EA_CharacterWindow | unresolved |  | — |
| Dependency | EA_ContextMenu | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
