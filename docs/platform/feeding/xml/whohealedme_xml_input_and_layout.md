# WhoHealedMe XML Input And Layout Findings

## Status

- Promoted into canonical XML caveat guidance on 2026-04-01.

## Source

- Addon: `WhoHealedMe`
- Validation context: healer-details window implementation and in-game user validation
- Date: 2026-04-01

## Target Seed

- `docs/platform/seeds/xml_conventions.md`

## Confidence

- Overall: `MEDIUM`
- Reason: behavior was observed and user-validated during addon implementation, but is still phrased as implementation guidance rather than canonical engine guarantee.

## Observed Behavior

### 1. Ancestor input can block child click handling

- `MEDIUM`: In nested XML window trees, enabling `handleinput` on a child may not be sufficient if an ancestor container or template instance is not input-enabled.
- Practical effect: a child `OnLButtonUp` click target may appear correctly wired but still never fire.
- Working guidance: when click handlers do not fire, verify the parent/template input chain as well as the child window itself.

### 2. Early nested dimensions can under-report usable space

- `MEDIUM`: During early show/layout passes, `WindowGetDimensions` on nested content or scroll-child windows can be temporarily smaller than the eventual usable region.
- Practical effect: scrollable layouts can appear collapsed on first render, including lists that only show one visible row.
- Working guidance: size from the outer parent window first, then apply child sizing and call `ScrollWindowUpdateScrollRect`.

## Evidence

### Ancestor input finding

- A child healer-name click target in WhoHealedMe did not open its details window until the parent row template was changed from `handleinput="false"` to `handleinput="true"`.
- The child click window already had explicit `OnLButtonDown` and `OnLButtonUp` handlers, so the observed fix was in the ancestor chain rather than the child handler definition.

### Nested-dimensions finding

- The WhoHealedMe details scroll region initially rendered with only one visible spell row.
- The issue stopped once the layout code derived height from the outer details window instead of inner content/scroll windows, then updated the scroll rect.

## Promotion Notes

- Promote as XML behavior caveats, not as new API symbol definitions.
- Keep both notes tagged `MEDIUM` unless broader cross-addon validation is collected.
- If later corroborated in multiple addons, these notes can move into `docs/platform/seeds/xml_conventions.md` under an observed behavior or caveats section.
- Promotion completed into `docs/war-api/meta/conventions.md`, `docs/war-api/events/window_events/window_event_OnLButtonUp.md`, and `docs/war-api/xml/element_types/element_ScrollWindow.md`.