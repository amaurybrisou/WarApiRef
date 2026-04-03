# SelfTarget Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | SelfTarget_Initialize | exact | HIGH | SelfTarget.SelfTarget_Initialize |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | SelfTarget_Shutdown | exact | HIGH | SelfTarget.SelfTarget_Shutdown |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | LibTimer | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | LibTimer | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
