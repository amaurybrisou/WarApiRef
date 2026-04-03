# Frame ClosetGoblinOptionWindow

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinOptionWindow
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
<Window inherits="EA_Window_Default" movable="false" name="..." savesettings="true">
  <Size>
    <AbsPoint x="48" y="48"/>
  </Size>
  <Anchors>
    <Anchor point="center" relativePoint="center">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnInitialize" function="ClosetGoblinOptionWindow.OnIn..."/>
  </EventHandlers>
  <Windows/>
</Window>
```

## Attributes

- inherits: EA_Window_Default
- movable: false
- name: ClosetGoblinOptionWindow
- savesettings: true

## Handlers

| Event | Function |
| --- | --- |
| OnInitialize | ClosetGoblinOptionWindow.OnInitialize |
