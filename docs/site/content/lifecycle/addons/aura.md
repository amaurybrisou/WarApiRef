# Aura Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | AuraAddon.OnInitialize | ambiguous | MEDIUM | AuraTooltip.OnInitialize, AuraAddon.OnInitialize, AutoMark.OnInitialize, ClosetGoblin.OnInitialize, ClosetGoblinCharacterWindow.OnInitialize, ClosetGoblinZoneWindow.OnInitialize, ClosetGoblinOptionWindow.OnInitialize, RoR_SoR.OnInitialize, Shinies.OnInitialize, _G.Shinies.OnInitialize |
| CallFunction | AuraTooltip.OnInitialize | ambiguous | MEDIUM | AuraTooltip.OnInitialize, AuraAddon.OnInitialize, AutoMark.OnInitialize, ClosetGoblin.OnInitialize, ClosetGoblinCharacterWindow.OnInitialize, ClosetGoblinZoneWindow.OnInitialize, ClosetGoblinOptionWindow.OnInitialize, RoR_SoR.OnInitialize, Shinies.OnInitialize, _G.Shinies.OnInitialize |
| CreateWindow | AuraTexture | exact | HIGH | AuraTexture |
| CreateWindow | AuraColorPicker | exact | HIGH | AuraColorPicker |
| CreateWindow | AuraConfig | exact | HIGH | AuraConfig |
| CreateWindow | AuraShares | exact | HIGH | AuraShares |
| CreateWindow | AuraSettings | exact | HIGH | AuraSettings |
## Shutdown Actions (OnShutdown)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | AuraSettings.OnShutdown | ambiguous | MEDIUM | AuraAddon.OnShutdown, ClosetGoblin.OnShutdown, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, LibGuard.OnShutdown, PartyCast.OnShutdown, RoR_SoR.OnShutdown, TidyChat.OnShutdown |
| CallFunction | AuraAddon.OnShutdown | ambiguous | MEDIUM | AuraAddon.OnShutdown, ClosetGoblin.OnShutdown, ClosetGoblinCharacterWindow.OnShutdown, ClosetGoblinZoneWindow.OnShutdown, LibGuard.OnShutdown, PartyCast.OnShutdown, RoR_SoR.OnShutdown, TidyChat.OnShutdown |
## Update Actions (OnUpdate)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | AuraEngine.Event_UPDATE_PROCESSED | exact | HIGH | AuraEngine.Event_UPDATE_PROCESSED |
## Unknown / Custom Hook Actions

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_AbilitiesWindow | unresolved |  | — |
| Dependency | EA_CareerResourcesWindow | unresolved |  | — |
| Dependency | EA_MoraleWindow | unresolved |  | — |
| Dependency | EA_GuildWindow | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EATemplate_DefaultWindowSkin | unresolved |  | — |
| Dependency | EATemplate_Icons | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
## Unresolved References

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| Dependency | EA_AbilitiesWindow | unresolved |  | — |
| Dependency | EA_CareerResourcesWindow | unresolved |  | — |
| Dependency | EA_MoraleWindow | unresolved |  | — |
| Dependency | EA_GuildWindow | unresolved |  | — |
| Dependency | EASystem_Utils | unresolved |  | — |
| Dependency | EASystem_WindowUtils | unresolved |  | — |
| Dependency | EASystem_Tooltips | unresolved |  | — |
| Dependency | EATemplate_DefaultWindowSkin | unresolved |  | — |
| Dependency | EATemplate_Icons | unresolved |  | — |
| Dependency | LibSlash | unresolved |  | — |
| Categories |  | unresolved |  | — |
| Careers |  | unresolved |  | — |
