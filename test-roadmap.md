# Go Testing Roadmap

Your tactical map for mastering Go testing. From `t.Errorf` to `go test -fuzz`, one battle at a time.

---

## ✅ 1. Basic Unit Testing

* [ ] Understanding `testing.T`
* [ ] Running tests via `go test`
* [ ] Creating `*_test.go` files
* [ ] Simple checks using `if got != want`
* [ ] Running specific tests with `go test -run`

## ✅ 2. Table-Driven Tests

* [ ] Using slices of test cases
* [ ] Subtests with `t.Run(...)`
* [ ] Descriptive subtest names
* [ ] Comparing `NaN`, `float`, and result accuracy

## ✅ 3. Coverage & Benchmarks

* [ ] Measuring coverage with `go test -cover`
* [ ] Generating profiles with `go test -coverprofile`
* [ ] Viewing coverage in browser using `go tool cover -html`
* [ ] Writing `BenchmarkX(b *testing.B)` functions
* [ ] Comparing performance of multiple implementations

## ✅ 4. Errors & Panic

* [ ] Testing errors with `errors.Is`, `errors.As`
* [ ] Capturing `panic` using `recover`
* [ ] Validating edge cases and error scenarios

## ✅ 5. Interfaces & Mocks

* [ ] Creating interfaces and decoupling logic
* [ ] Writing manual mocks
* [ ] Verifying mock method calls
* [ ] Using `gomock` and `mockgen`
* [ ] Using `testify/mock`

## ✅ 6. HTTP Handlers

* [ ] Using `httptest.NewRequest` and `httptest.NewRecorder`
* [ ] Validating HTTP status codes
* [ ] Checking response body and headers
* [ ] Parsing and validating JSON responses
* [ ] Integrating with routers (mux, gin, etc.)

## ✅ 7. Middleware Testing

* [ ] Wrapping handlers with middleware
* [ ] Asserting request forwarding and behavior
* [ ] Validating authentication and headers

## ✅ 8. Parallel & Param Tests

* [ ] Leveraging `t.Parallel()` for test concurrency
* [ ] Comparing sequential vs parallel execution
* [ ] Detecting race conditions

## ✅ 9. Fuzz Testing

* [ ] Running fuzz tests with `go test -fuzz`
* [ ] Adding seed cases with `f.Add(...)`
* [ ] Catching unexpected crashes
* [ ] Detecting edge case failures

## ✅ 10. Property-Based Testing

* [ ] Using `testing/quick`
* [ ] Defining invariants to test
* [ ] Auto-generating test inputs
* [ ] Handling invalid or error-prone data

## ✅ 11. Testdata & Fixtures

* [ ] Organizing `testdata/` directory
* [ ] Loading JSON / YAML files
* [ ] Using golden tests to compare outputs
* [ ] Snapshot reading and comparison

## ✅ 12. External Integration

* [ ] Testing with Redis / Postgres via Docker
* [ ] Using `dockertest` or similar tools
* [ ] Writing setup and teardown logic
* [ ] Validating database migrations and transactions

## ✅ 13. CI / GitHub Actions

* [ ] Configuring `.github/workflows/go.yml`
* [ ] Running full test suite: `go test ./...` + `-race`
* [ ] Storing coverage as artifacts
* [ ] Validating PRs with automated tests

## ✅ 14. Best Practices

* [ ] Naming tests clearly: `TestXxx_Case`
* [ ] Keeping tests small and deterministic
* [ ] Using `require` vs `assert` wisely
* [ ] Project structure: `/pkg`, `/internal`, `/mocks`

