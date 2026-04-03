# Frame yClosetGoblinButton

- Addon: CM_ClosetGoblin
- Resolved Name: yClosetGoblinButton
- Type: DynamicImage
- Parent: none
- Parent Type: none
- Inherits: none
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

## Composition Pattern

```xml
<DynamicImage layer="popup" movable="false" name="..." popable="true" savesettings="false" texture="WHLCG">
  <Size>
    <AbsPoint x="36" y="36"/>
  </Size>
  <Anchors>
    <Anchor point="topright" relativePoint="topright" relativeTo="CharacterWindowTitleBar">
      <AbsPoint x="-76" y="0"/>
    </Anchor>
  </Anchors>
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinOptionWindow.OnLB..."/>
    <EventHandler event="OnRButtonUp" function="ClosetGoblinOptionWindow.OnRB..."/>
  </EventHandlers>
</DynamicImage>
```

## Attributes

- layer: popup
- movable: false
- name: yClosetGoblinButton
- popable: true
- savesettings: false
- texture: WHLCG

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinOptionWindow.OnLButtonUp |
| OnRButtonUp | ClosetGoblinOptionWindow.OnRButtonUp |
