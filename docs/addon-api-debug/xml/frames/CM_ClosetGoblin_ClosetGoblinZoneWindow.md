# Frame ClosetGoblinZoneWindow

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinZoneWindow
- Type: Window
- Parent: none
- Parent Type: none
- Inherits: EA_Window_Default
- Template: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.xml:0`

## Children

- none

## Child Element Types

- none

## Structural Child Types

- Anchors
- EventHandlers
- Size
- Windows

## Composition Pattern

```xml
<Window inherits="EA_Window_Default" layer="default" movable="true" name="...">
  <Size>
    <AbsPoint x="545" y="753"/>
  </Size>
  <Anchors>
    <Anchor point="topleft" relativePoint="topleft">
      <AbsPoint x="35" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnShown" function="ClosetGoblinZoneWindow.OnShow"/>
    <EventHandler event="OnHidden" function="ClosetGoblinZoneWindow.OnHide"/>
    <EventHandler event="OnShutdown" function="ClosetGoblinZoneWindow.OnShut..."/>
  </EventHandlers>
  <Windows/>
</Window>
```

## Attributes

- inherits: EA_Window_Default
- layer: default
- movable: true
- name: ClosetGoblinZoneWindow

## Handlers

| Event | Function |
| --- | --- |
| OnHidden | ClosetGoblinZoneWindow.OnHide |
| OnShown | ClosetGoblinZoneWindow.OnShow |
| OnShutdown | ClosetGoblinZoneWindow.OnShutdown |
