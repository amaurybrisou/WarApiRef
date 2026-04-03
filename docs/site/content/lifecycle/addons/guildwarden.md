# GuildWarden Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- GuildWarden.Data
- GuildWarden.light
- GuildWarden.lastSendDate
- GuildWarden.LvlU
- GuildWarden.LvlL
- GuildWarden.RrU
- GuildWarden.RrL
- GuildWarden.Combo1
- GuildWarden.Combo2
- GuildWarden.WardS

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | GuildWarden.Init_Warden | exact | HIGH | GuildWarden.Init_Warden |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_ChatWindow | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_ChatWindow | unresolved |  | — |
