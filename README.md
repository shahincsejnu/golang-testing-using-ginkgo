# golang-testing-using-ginkgo


# Intro

Followed the [Official Tutorial](https://onsi.github.io/ginkgo/#understanding-ginkgos-lifecycle) and  [tutorial](https://semaphoreci.com/community/tutorials/getting-started-with-bdd-in-go-using-ginkgo) where it showed how we can incorporate Behavior-Driven-Development(BDD) with go using `Ginkgo`. Ginkgo makes use of Go's existing `testing` package.

## Hacks
- `Ginkgo` is a BDD(Behavior Driven Development) style testing framework for `Golang`, and it's preferred matcher library is `Gomega`.
- It helps us to write descriptive and comprehensive tests.
- `go get github.com/onsi/ginkgo/ginkgo`: to install ginkgo CLI
- `go get github.com/onsi/gomega`: to fetch the Gomega matcher library
- `ginkgo generate <name>` : This command will create a sample test file.
- `go test`: To run tests
- `go test -v`: To run tests (with seeing the verbose)
- `ginkgo`: To run tests
- `ginkgo watch`: To rerun tests on the command line whenever changes are detected.

## Blocks
- `BeforeSuite()`:
    - These blocks are executed once before any other specs are run.
    - In parallel environment (when more than one node running in parallel), every parallel node will call this once.
    - Ex:
        ```
            var _ = BeforeSuite(func() {
                // do something
            })
        ```
- `AfterSuite()`:
    - These blocks are always run after all the specs have run.
    - The success or failure of the specs is not matter, this block is always will be run
    - Even is the execution is interrupted by a signal (like ^C), still Ginkgo will attempt to run this block before exiting or quiting.
    - Ex:
      ```
      var _ = AfterSuite(func() {
        // do something
      })
      ```
- `BeforeEach()`:
    - These blocks are run before the `It` blocks.
    - This block is used to remove duplication and share a common setup across tests.
    - If multiple `BeforeEach()` blocks are defined within nested `Describe` and `Context` blocks then the outermost `BeforeEach()` block is the one that is executed first.
    - Ex:
        ```
            var _ = Describe("Demo", func() {
                var temp string
                BeforeEach(func() {
                    temp = "Nothing"
                })
            })
        ```
- `AfterEach()`:
    - These blocks are run after the `It` blocks.
    - If there is a situation where there are nested `Describe` and `Context` blocks and multiple `AfterEach()` functions are defined within them, then the innermost `AfterEach()` block is executed first followed by the next innermost and so on.
    - Ex:
        ```
            AfterEach(func() {
                // do something
            })
        ```
- `Describe()`:
    - It contains our specs.
    -
- `Context()`:
    - It is used to give the specs (individually or groups).
    

# Resources

- [x] [Ginkgo Official](https://onsi.github.io/ginkgo/#getting-ginkgo)
- [x] [Getting Started with BDD in Go Using Ginkgo](https://semaphoreci.com/community/tutorials/getting-started-with-bdd-in-go-using-ginkgo)