# Frame ClosetGoblinOptionWindowIcon

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinOptionWindowIcon
- Raw Name: $parentIcon
- Type: DynamicImage
- Parent: ClosetGoblinOptionWindow
- Parent Type: Window
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

## Composition Pattern

```xml
<DynamicImage layer="default" name="..." texture="ClosetGoblinIcon" texturescale="1.00">
  <EventHandlers>
    <EventHandler event="OnLButtonUp" function="ClosetGoblinOptionWindow.OnLB..."/>
    <EventHandler event="OnRButtonUp" function="ClosetGoblinOptionWindow.OnRB..."/>
  </EventHandlers>
  <Anchors>
    <Anchor point="topleft" relativePoint="topleft" relativeTo="$parent">
      <AbsPoint x="0" y="0"/>
    </Anchor>
    <Anchor point="bottomright" relativePoint="bottomright" relativeTo="$parent">
      <AbsPoint x="0" y="0"/>
    </Anchor>
  </Anchors>
</DynamicImage>
```

## Attributes

- layer: default
- name: $parentIcon
- texture: ClosetGoblinIcon
- texturescale: 1.00

## Handlers

| Event | Function |
| --- | --- |
| OnLButtonUp | ClosetGoblinOptionWindow.OnLButtonUp |
| OnRButtonUp | ClosetGoblinOptionWindow.OnRButtonUp |
