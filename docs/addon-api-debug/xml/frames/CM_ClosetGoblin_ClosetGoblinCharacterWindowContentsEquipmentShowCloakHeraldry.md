# Frame ClosetGoblinCharacterWindowContentsEquipmentShowCloakHeraldry

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowContentsEquipmentShowCloakHeraldry
- Raw Name: $parentShowCloakHeraldry
- Type: Button
- Parent: ClosetGoblinCharacterWindowContentsEquipment
- Parent Type: Window
- Inherits: EA_Button_DefaultCheckBox
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
<Button inherits="EA_Button_DefaultCheckBox" layer="popup" name="...">
  <Anchors>
    <Anchor point="topright" relativePoint="topright" relativeTo="$parentSlot13">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnMouseOver" function="ClosetGoblinCharacterWindow.S..."/>
    <EventHandler event="OnMouseOverEnd" function="ClosetGoblinCharacterWindow.H..."/>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinCharacterWindow.S..."/>
  </EventHandlers>
</Button>
```

## Attributes

- inherits: EA_Button_DefaultCheckBox
- layer: popup
- name: $parentShowCloakHeraldry

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinCharacterWindow.ShowCloakHeraldry |
| OnMouseOver | ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly |
| OnMouseOverEnd | ClosetGoblinCharacterWindow.HideCloakOptions |
