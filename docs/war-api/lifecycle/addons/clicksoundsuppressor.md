# ClickSoundSuppressor Lifecycle

> Source: `.mod` manifest semantic analysis

## Lifecycle Hooks

- OnInitialize
- OnUpdate
- OnShutdown

## Startup Actions (OnInitialize)

| Kind | Name | Resolution | Confidence | Matched |
| --- | --- | --- | --- | --- |
| CallFunction | ClickSoundSuppressor.OnInit | ambiguous | MEDIUM | BBars.OnInit, CNC.OnInit, ClickSoundSuppressor.OnInit, PetFixWindow.onInit |
