# RoR_SoR Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Saved Variables

- RoR_SoR.Settings
- RoR_SoR.City_Status

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | RoR_SoR.OnInitialize | ambiguous | MEDIUM | AuraTooltip.OnInitialize, AuraAddon.OnInitialize, AutoMark.OnInitialize, ClosetGoblin.OnInitialize, ClosetGoblinCharacterWindow.OnInitialize, ClosetGoblinZoneWindow.OnInitialize, ClosetGoblinOptionWindow.OnInitialize, RoR_SoR.OnInitialize, Shinies.OnInitialize, _G.Shinies.OnInitialize |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | RoR_SoR.OnShutdown | ambiguous | MEDIUM | AuraAddon.OnShutdown, ClosetGoblin.OnShutdown, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, LibGuard.OnShutdown, PartyCast.OnShutdown, RoR_SoR.OnShutdown, TidyChat.OnShutdown |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | RoR_SoR.TIMER_UPDATE | ambiguous | MEDIUM | Enemy.Timer_Update, RoR_SoR.TIMER_UPDATE |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EATemplate_DefaultWindowSkin | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EATemplate_ParchmentWindowSkin | unresolved |  | — |
| Dependency | EA_SiegeWeaponWindow | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EATemplate_DefaultWindowSkin | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EA_ChatWindow | unresolved |  | — |
| Dependency | EATemplate_ParchmentWindowSkin | unresolved |  | — |
| Dependency | EA_SiegeWeaponWindow | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
