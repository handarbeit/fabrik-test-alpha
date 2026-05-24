// fabrik-test-alpha is the primary test bed for Fabrik. It depends on
// fabrik-test-beta and exists purely as substrate for exercising
// multi-repo Fabrik features.
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Fprintln(os.Stderr, "fabrik-test-alpha — initial scaffold; nothing wired to fabrik-test-beta yet")
}
