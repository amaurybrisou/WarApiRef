# SNT_CASTBAR Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- SNTCAST_DB

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | snt_castbar.entry_point | ambiguous | MEDIUM | snt_info.tdps_module.entry_point, snt_buttons.entry_point, snt_castbar.entry_point, snt_info.entry_point, snt_panel.entry_point |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_CastTimerWindow | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_CastTimerWindow | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
