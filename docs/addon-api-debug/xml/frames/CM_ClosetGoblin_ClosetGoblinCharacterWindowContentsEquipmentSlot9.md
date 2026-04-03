# Frame ClosetGoblinCharacterWindowContentsEquipmentSlot9

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsEquipmentSlot9
- Raw Name: $parentSlot9
- Type: Button
- Parent: ClosetGoblinCharacterWindowContentsEquipment
- Parent Type: Window
- Inherits: ClosetGoblinEquipmentButton
- Template: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.xml:0`

## Children

- none

## Child Element Types

- none

## Structural Child Types

- Anchors
- EventHandlers

## Composition Pattern

```xml
<Button id="9" inherits="ClosetGoblinEquipmentButton" name="...">
  <Anchors>
    <Anchor point="top" relativePoint="bottom" relativeTo="$parentSlot10">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonDown" function="ClosetGoblinCharacterWindow.E..."/>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.E..."/>
    <EventHandler event="OnRButtonDown" function="ClosetGoblinCharacterWindow.E..."/>
    <EventHandler event="OnMouseOver" function="ClosetGoblinCharacterWindow.S..."/>
    <EventHandler event="OnMouseOverEnd" function="ClosetGoblinCharacterWindow.H..."/>
  </EventHandlers>
</Button>
```

## Attributes

- id: 9
- inherits: ClosetGoblinEquipmentButton
- name: $parentSlot9

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonDown | ClosetGoblinCharacterWindow.EquipmentLButtonDown |
| OnLButtonUp | ClosetGoblinCharacterWindow.EquipmentLButtonUp |
| OnMouseOver | ClosetGoblinCharacterWindow.ShowShowHelm |
| OnMouseOverEnd | ClosetGoblinCharacterWindow.HideShowHelm |
| OnRButtonDown | ClosetGoblinCharacterWindow.EquipmentRButtonDown |
