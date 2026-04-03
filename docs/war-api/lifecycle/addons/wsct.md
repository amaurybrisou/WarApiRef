# WSCT Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- WSCT_CONFIG

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | WSCT.OnInitialize | ambiguous | MEDIUM | AuraTooltip.OnInitialize, AuraAddon.OnInitialize, AutoMark.OnInitialize, ClosetGoblin.OnInitialize, ClosetGoblinCharacterWindow.OnInitialize, ClosetGoblinZoneWindow.OnInitialize, ClosetGoblinOptionWindow.OnInitialize, RoR_SoR.OnInitialize, Shinies.OnInitialize, _G.Shinies.OnInitialize |
| CreateWindow | WSCTContainer | exact | HIGH | WSCTContainer |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | WSCT.OnShutdown | ambiguous | MEDIUM | AuraAddon.OnShutdown, ClosetGoblin.OnShutdown, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, LibGuard.OnShutdown, PartyCast.OnShutdown, RoR_SoR.OnShutdown, TidyChat.OnShutdown |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | WSCT.UpdateAnimation | exact | HIGH | WSCT.UpdateAnimation |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_EventText | unresolved |  | — |
| Dependency | EA_SiegeWeaponWindow | unresolved |  | — |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EA_CareerResourcesWindow | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_EventText | unresolved |  | — |
| Dependency | EA_SiegeWeaponWindow | unresolved |  | — |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EA_CareerResourcesWindow | unresolved |  | — |
