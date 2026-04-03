# Frame ClosetGoblinZoneWindowContentsSortUpArrow

- Addon: CM_ClosetGoblin
- Resolved Name: ClosetGoblinZoneWindowContentsSortUpArrow
- Raw Name: $parentSortUpArrow
- Type: DynamicImage
- Parent: ClosetGoblinZoneWindowContents
- Parent Type: Window
- Inherits: EA_ListSortUpArrow
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
<DynamicImage inherits="EA_ListSortUpArrow" name="...">
  <Anchors>
    <Anchor point="right" relativePoint="right" relativeTo="$parentSortZoneButton">
      <AbsPoint x="-3" y="0"/>
    </Anchor>
  </Anchors>
</DynamicImage>
```

## Attributes

- inherits: EA_ListSortUpArrow
- name: $parentSortUpArrow

## Handlers

| Event | Function |
| --- | --- |
