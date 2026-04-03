# Frame ClosetGoblinEquipmentButton

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinEquipmentButton
- Type: Button
- Parent: none
- Parent Type: none
- Inherits: ClosetGoblinDefaultButton
- Template: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.xml:0`

## Children

- none

## Child Element Types

- none

## Structural Child Types

- EventHandlers

## Composition Pattern

```xml
<Button draganddrop="true" gameactionbutton="right" inherits="ClosetGoblinDefaultButton" name="...">
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.E..."/>
    <EventHandler event="OnLButtonDown" function="ClosetGoblinCharacterWindow.E..."/>
    <EventHandler event="OnRButtonUp" function="ClosetGoblinCharacterWindow.E..."/>
    <EventHandler event="OnMouseOver" function="ClosetGoblinCharacterWindow.E..."/>
    <EventHandler event="OnMouseOverEnd" function="ClosetGoblinCharacterWindow.E..."/>
  </EventHandlers>
</Button>
```

## Attributes

- draganddrop: true
- gameactionbutton: right
- inherits: ClosetGoblinDefaultButton
- name: ClosetGoblinEquipmentButton

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonDown | ClosetGoblinCharacterWindow.EquipmentLButtonDown |
| OnLButtonUp | ClosetGoblinCharacterWindow.EquipmentLButtonUp |
| OnMouseOver | ClosetGoblinCharacterWindow.EquipmentMouseOver |
| OnMouseOverEnd | ClosetGoblinCharacterWindow.EquipmentMouseOverEnd |
| OnRButtonUp | ClosetGoblinCharacterWindow.EquipmentRButtonUp |
