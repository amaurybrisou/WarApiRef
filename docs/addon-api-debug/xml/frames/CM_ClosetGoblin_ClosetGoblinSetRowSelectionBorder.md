# Frame ClosetGoblinSetRowSelectionBorder

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinSetRowSelectionBorder
- Raw Name: $parentSelectionBorder
- Type: FullResizeImage
- Parent: ClosetGoblinSetRow
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
    <Anchor point="topleft" relativePoint="topleft" relativeTo="$parentName">
      <AbsPoint x="-4" y="-4"/>
    </Anchor>
    <Anchor point="bottomright" relativePoint="bottomright" relativeTo="$parentTactics">
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
