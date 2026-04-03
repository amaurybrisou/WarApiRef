# Frame CG_ItemRackEquipmentButton

- Addon: CM_ClosetGoblin
- Resolved Name: CG_ItemRackEquipmentButton
- Type: Button
- Parent: none
- Parent Type: none
- Inherits: CharacterWindowDefaultButton
- Template: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinItemRack.xml:0`

## Children

- none

## Child Element Types

- none

## Structural Child Types

- EventHandlers

## Composition Pattern

```xml
<Button draganddrop="true" gameactionbutton="right" inherits="CharacterWindowDefaultButton" name="...">
  <EventHandlers>
    <EventHandler event="OnRButtonDown" function="CG_ItemRack.EquipmentRButtonDown"/>
    <EventHandler event="OnMouseOver" function="CG_ItemRack.EquipmentMouseOver"/>
  </EventHandlers>
</Button>
```

## Attributes

- draganddrop: true
- gameactionbutton: right
- inherits: CharacterWindowDefaultButton
- name: CG_ItemRackEquipmentButton

## Handlers

| Event | Function |
| --- | --- |
| OnMouseOver | CG_ItemRack.EquipmentMouseOver |
| OnRButtonDown | CG_ItemRack.EquipmentRButtonDown |
