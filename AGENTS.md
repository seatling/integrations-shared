# Seatling Integrations Shared: AI Agent Guidebook (`AGENTS.md`)

Welcome, Agent. This document defines the engineering standards, architectural boundaries, and tactical instructions for `github.com/seatling/integrations-shared`. All AI coding agents working in this repository must adhere to these conventions.

---

## 1. Overview & Architectural Role

`github.com/seatling/integrations-shared` is the canonical contract and interface module for Seatling's checkout and payment integrations (Gumroad, Lemon Squeezy, Polar, Stripe, etc.).

### Key Boundaries:
- **Pure Interface & Model Library**: Defines canonical Go interfaces (`ProviderClient`, `WebhookParser`, `Integration`) and normalized data models (`Order`, `Subscription`, `Product`, `User`, `NormalizedWebhookEvent`, `Refund`, `LicenseVerificationRequest`, `LicenseVerificationResult`, `WebhookSubscription`).
- **Zero Third-Party Dependencies**: Restricted strictly to the Go standard library (100% stdlib). Do not introduce external dependencies into `go.mod`.
- **No Concrete Implementations**: Concrete HTTP dispatchers, webhook signature verifiers, and provider-specific payload decoders belong strictly in specialized provider packages (e.g. `seatling-integrations-gumroad`, `seatling-integrations-lemonsqueezy`, `seatling-integrations-polar`, `seatling-integrations-stripe`).
- **Dual Serialization Alignment**: All domain models must define both `bson:"..."` and `json:"..."` struct tags, column-aligned for readability and interoperability across MongoDB and JSON pipelines.
- **Coverage Standard**: Maintain 100% statement test coverage across all enums, sentinel errors, helper methods, and JSON roundtrip tests.

---

## 2. Setup Commands

- **Verify Go environment**: Go 1.26+ required
  ```bash
  go version
  ```
- **Download and verify module**:
  ```bash
  go mod download
  go mod verify
  ```

---

## 3. Build, Test & Lint Commands

All verification in this repository uses the canonical shell scripts [`test.sh`](./test.sh) and [`lint.sh`](./lint.sh) (which mirror GitHub Actions CI):

### Primary Verification Scripts (Mandatory)

| Script | Purpose | Underlying Command |
| :--- | :--- | :--- |
| `./test.sh` | **Mandatory Test Verification**: Runs full test suite with statement coverage and timeout | `go test -v -cover -timeout 10s ./...` |
| `./lint.sh` | **Mandatory Lint Verification**: Runs Dockerized `golangci-lint` (v2.0) with bind-mount | `MSYS_NO_PATHCONV=1 docker run --rm --mount type=bind,src=.,dst=/app -w /app golangci/golangci-lint:v2.0 golangci-lint run -v` |

> [!NOTE]
> On Windows environments without a direct bash shebang association, execute scripts via Git Bash / WSL: `bash ./test.sh` and `bash ./lint.sh`.

### Supplemental & Direct Go CLI Commands

| Command | Purpose |
| :--- | :--- |
| `go test -v -cover -timeout 10s ./...` | Direct Go execution of the test suite (identical to `test.sh`) |
| `go test -v -cover ./...` | Runs all unit tests with statement coverage output |
| `go test -v -race ./...` | Runs test suite with Go data race detector enabled |
| `go vet ./...` | Runs Go standard static analysis |
| `gofmt -s -w .` | Applies standard Go formatting across all files |
| `golangci-lint run -v` | Runs local `golangci-lint` binary if installed locally |

---

## 4. Architectural Invariants & Rules

### Rule 1: Zero External Dependencies
- `go.mod` must only declare `module github.com/seatling/integrations-shared` and the Go toolchain version (`go 1.26.0`).
- Never add third-party dependencies (`github.com/...`) to `go.mod`. Use standard library packages (`context`, `errors`, `fmt`, `net/http`, `time`, etc.).

### Rule 2: Pure Contract Boundaries
- **Do not write concrete HTTP clients**: This module only defines the `ProviderClient`, `WebhookParser`, and `Integration` interfaces, plus the minimal `HTTPClient` interface (`Do(req *http.Request) (*http.Response, error)`).
- **Do not write provider-specific payload parsers**: Specialized logic (such as decoding Gumroad URL-encoded forms or validating Lemon Squeezy HMAC-SHA256 signatures) belongs in downstream provider implementations.

### Rule 3: Model & Tag Conventions
- All models must live in `package integrations`.
- Models must provide dual, column-aligned struct tags:
  ```go
  type Product struct {
      ID             string         `bson:"id"                       json:"id"`
      Provider       Provider       `bson:"provider"                 json:"provider"`
      Name           string         `bson:"name"                     json:"name"`
      Description    string         `bson:"description,omitempty"     json:"description,omitempty"`
      PriceCents     int            `bson:"price_cents"              json:"price_cents"`
      Currency       string         `bson:"currency"                 json:"currency"`
      Published      bool           `bson:"published"                json:"published"`
      Variants       []Variant      `bson:"variants,omitempty"       json:"variants,omitempty"`
      CreatedAt      *time.Time     `bson:"created_at,omitempty"      json:"created_at,omitempty"`
      RawData        map[string]any `bson:"raw_data,omitempty"       json:"raw_data,omitempty"`
  }
  ```
- Use `omitempty` for optional fields, timestamps, and nested slice/map structures.
- Always include `RawData map[string]any` or `RawPayload map[string]any` with `omitempty` on models to allow downstream consumers to preserve provider-specific raw fields.

### Rule 4: Provider & Enum Taxonomy
- Providers must be declared as typed string enums (`type Provider string`).
- Implement `IsValid() bool` and `String() string` on all enums (`Provider`, `SubscriptionStatus`).
- New event types must follow the canonical dot-notation format (`<resource>.<action>`, e.g., `order.created`, `subscription.updated`).

### Rule 5: Error Taxonomy
- Sentinel errors must be prefixed with `integrations:` (e.g., `integrations: unauthorized or invalid credentials`).
- `APIError` must implement:
  - `Error() string`
  - `Is(target error) bool` for status code-to-sentinel matching (`ErrUnauthorized`, `ErrNotFound`, `ErrRateLimited`, `ErrResourceConflict`)
  - `Retryable() bool` (true for HTTP 429 and 5xx server errors)

---

## 5. Testing & Verification Mandate

Whenever modifying models, enums, or error handling, AI agents **must complete the following verification workflow**:

1. **JSON Roundtrip Tests**: Ensure every struct added or modified is covered by serialization roundtrip assertions in `models_test.go`.
2. **Enum Validation Tests**: Verify valid and invalid branches in `provider_test.go`.
3. **Error Matching Tests**: Verify `errors.Is(...)` and `Retryable()` handling in `errors_test.go`.
4. **Execute `test.sh`**: Run `./test.sh` (or `bash ./test.sh`) and verify statement coverage remains strictly at **100%**.
5. **Execute `lint.sh`**: Run `./lint.sh` (or `bash ./lint.sh` / `go vet ./...`) and ensure all lint inspections pass with zero warnings.
6. **Code Formatting**: Ensure `gofmt -s -w .` produces no diff.

---

## 6. Commit & PR Guidelines

- **Commit Message Format**: Follow Conventional Commits:
  - `feat: add <model/field/enum>`
  - `fix: update <type/method>`
  - `test: add coverage for <feature>`
  - `chore: update <tooling/workflow>`
- **Mandatory Pre-Commit Verification**: Run both `./test.sh` and `./lint.sh` (or their direct CLI equivalents) to ensure clean exit codes before submitting any changes.
