# Copilot Instructions for `goark/fetch`

## Project purpose

`fetch` is a small package focused on downloading data from URL endpoints.
The package should hide common `net/http` idioms and keep calling code simple.

## Design principles

- Prefer explicit, small APIs over feature-rich abstractions.
- Keep `context.Context` based methods as the default path.
- Favor composable option functions (`ClientOpts`, `RequestOpts`).
- Preserve compatibility when possible; avoid breaking public symbols.

## Error handling

- Use `github.com/goark/errs` as the primary internal error handling package.
- Prefer `errs.Wrap`, `errs.Join`, and `errs.WithContext` over ad-hoc wrapping patterns.
- Keep `errors.Is` compatibility when returning wrapped errors.
- Return wrapped errors so callers can use `errors.Is`.
- Keep sentinel errors stable (`ErrInvalidURL`, `ErrInvalidRequest`, `ErrHTTPStatus`, `ErrNullPointer`).
- Include useful context values when wrapping errors.

## HTTP behavior

- Continue to use `net/http` as the base implementation.
- Ensure response body cleanup remains safe and predictable.
- Keep behavior for non-success HTTP status explicit and documented.

## Coding style

- Write idiomatic Go and keep implementation straightforward.
- Avoid unnecessary dependencies.
- Keep comments concise and in English.

## Testing and validation

- Add or update tests for behavior changes.
- Prefer local validation with Taskfile targets:
  - `task test`
  - `task govulncheck`

## Documentation

- Keep `README.md` in sync with public API changes.
- Include practical examples for GET/POST and error handling.
