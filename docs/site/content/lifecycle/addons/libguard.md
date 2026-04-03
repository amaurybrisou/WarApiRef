# LibGuard Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- LibGuard.UsePet

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | LibGuard.Init | ambiguous | MEDIUM | BagOMatic.init, BankArkel.Init, DAoCTooltips.Init, GuardLine.init, LibGuard.Init, MoraleCircle.init, PartyCast.Init |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | LibGuard.OnShutdown | ambiguous | MEDIUM | AuraAddon.OnShutdown, ClosetGoblin.OnShutdown, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, LibGuard.OnShutdown, PartyCast.OnShutdown, RoR_SoR.OnShutdown, TidyChat.OnShutdown |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EA_ActionBars | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
| Dependency | ror_PacketHandling | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EA_ActionBars | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
| Dependency | ror_PacketHandling | unresolved |  | — |
