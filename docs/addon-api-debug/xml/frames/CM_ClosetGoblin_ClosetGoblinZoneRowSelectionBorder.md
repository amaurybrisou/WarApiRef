# Frame ClosetGoblinZoneRowSelectionBorder

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinZoneRowSelectionBorder
- Raw Name: $parentSelectionBorder
- Type: FullResizeImage
- Parent: ClosetGoblinZoneRow
- Parent Type: Window
- Inherits: DefaultWindowBackground
- Template: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.xml:0`

## Children

- none

## Child Element Types

- none

## Structural Child Types

- Anchors

## Composition Pattern

```xml
<FullResizeImage inherits="DefaultWindowBackground" name="...">
  <Anchors>
    <Anchor point="topleft" relativePoint="topleft" relativeTo="$parentZone">
      <AbsPoint x="-4" y="-4"/>
    </Anchor>
    <Anchor point="bottomright" relativePoint="bottomright" relativeTo="$parentSet">
      <AbsPoint x="-10" y="0"/>
    </Anchor>
  </Anchors>
</FullResizeImage>
```

## Attributes

- inherits: DefaultWindowBackground
- name: $parentSelectionBorder

## Handlers

| Event | Function |
| --- | --- |
