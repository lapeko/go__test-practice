# Go Testing Roadmap

Your tactical map for mastering Go testing. From `t.Errorf` to `go test -fuzz`, one battle at a time.

---

## ✅ 1. Basic Unit Testing

* [x] Understanding `testing.T`
* [x] Running tests via `go test`
* [x] Creating `*_test.go` files
* [x] Simple checks using `if got != want`
* [x] Running specific tests with `go test -run`

## ✅ 2. Table-Driven Tests

* [x] Using slices of test cases
* [x] Subtests with `t.Run(...)`
* [x] Descriptive subtest names
* [x] Comparing `NaN`, `float`, and result accuracy

## ✅ 3. Coverage & Benchmarks

* [x] Measuring coverage with `go test -cover`
* [x] Generating profiles with `go test -coverprofile`
* [x] Viewing coverage in browser using `go tool cover -html`
* [x] Writing `BenchmarkX(b *testing.B)` functions
* [x] Comparing performance of multiple implementations

## ✅ 4. Errors & Panic

* [x] Testing errors with `errors.Is`, `errors.As`
* [x] Capturing `panic` using `recover`
* [x] Validating edge cases and error scenarios

## ✅ 5. Interfaces & Mocks

* [x] Creating interfaces and decoupling logic
* [x] Writing manual mocks
* [x] Verifying mock method calls
* [ ] Using `gomock` and `mockgen`
* [x] Using `testify/mock`

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

* [x] Leveraging `t.Parallel()` for test concurrency
* [x] Comparing sequential vs parallel execution
* [x] Detecting race conditions

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
