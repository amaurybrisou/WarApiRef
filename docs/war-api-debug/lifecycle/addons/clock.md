# Clock Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- ClockSettings

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | Clock.Initialize | ambiguous | MEDIUM | ClosetGoblin.Initialize, Clock.Initialize |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | Clock.OnUpdate | exact | HIGH | Clock.OnUpdate |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
