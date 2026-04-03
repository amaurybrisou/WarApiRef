# DammazKron Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnShutdown

## Saved Variables

- DammazKron.data

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | DammazKron.Start | ambiguous | MEDIUM | ActionFraction.Start, AuraEngine.Start, CaVES.Start, Countdown.start, DammazKron.Start, DammazKronTNFO.Start, DammazKronPLUG.Start, SetBookmarkHook.Start, DeMoNiCore.Start, Hopper.Start, ReliquaryHunter.Start, SessionRPs.Start |
| CallFunction | SetBookmarkHook.Start | ambiguous | MEDIUM | ActionFraction.Start, AuraEngine.Start, CaVES.Start, Countdown.start, DammazKron.Start, DammazKronTNFO.Start, DammazKronPLUG.Start, SetBookmarkHook.Start, DeMoNiCore.Start, Hopper.Start, ReliquaryHunter.Start, SessionRPs.Start |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | DammazKron.Stop | ambiguous | MEDIUM | AuraEngine.Stop, DammazKron.Stop, DammazKronTNFO.Stop, DeMoNiCore.Stop, Hopper.Stop |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_TomeOfKnowledge | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_TomeOfKnowledge | unresolved |  | — |
