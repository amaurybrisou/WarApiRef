# Frame ClosetGoblinCharacterWindowCornerImage

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinCharacterWindowCornerImage
- Raw Name: $parentCornerImage
- Type: DynamicImage
- Parent: ClosetGoblinCharacterWindow
- Parent Type: Window
- Inherits: EA_Default_CharacterImage
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
<DynamicImage inherits="EA_Default_CharacterImage" name="..." sticky="false">
  <Anchors>
    <Anchor point="topleft" relativePoint="topleft" relativeTo="$parentTitleBar">
      <AbsPoint x="-50" y="-60"/>
    </Anchor>
  </Anchors>
</DynamicImage>
```

## Attributes

- inherits: EA_Default_CharacterImage
- name: $parentCornerImage
- sticky: false

## Handlers

| Event | Function |
| --- | --- |
