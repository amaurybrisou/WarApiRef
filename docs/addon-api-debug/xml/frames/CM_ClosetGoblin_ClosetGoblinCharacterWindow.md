# Frame ClosetGoblinCharacterWindow

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindow
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
    <EventHandler event="OnShown" function="ClosetGoblinCharacterWindow.O..."/>
    <EventHandler event="OnHidden" function="ClosetGoblinCharacterWindow.O..."/>
    <EventHandler event="OnShutdown" function="ClosetGoblinCharacterWindow.O..."/>
  </EventHandlers>
  <Windows/>
</Window>
```

## Attributes

- inherits: EA_Window_Default
- layer: default
- movable: true
- name: ClosetGoblinCharacterWindow

## Handlers

| Event | Function |
| --- | --- |
| OnHidden | ClosetGoblinCharacterWindow.OnHide |
| OnShown | ClosetGoblinCharacterWindow.OnShow |
| OnShutdown | ClosetGoblinCharacterWindow.OnShutdown |
