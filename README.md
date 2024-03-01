# go-saga-func

### This package is an implementation of Choreography-based saga pattern in golang

Saga pattern used in distributed systems, particularly microservices architectures, to manage data consistency across multiple services. The purpose is to ensures data consistency across multiple services involved in a single business transaction.

Overall, the saga pattern is a powerful tool for managing data consistency in distributed systems but requires careful consideration and implementation due to it's complexities.

You can read more details about saga pattern here:

- [SAGAS](https://www.cs.cornell.edu/andru/cs711/2002fa/reading/sagas.pdf) by Hector Garcia-Molina, Kenneth Salem, and Harold F. Korth (1987)
- [Pattern: Saga](https://microservices.io/patterns/data/saga.html) by Chris Richardson
- [How to Use Saga Pattern in Microservices](https://blog.bitsrc.io/how-to-use-saga-pattern-in-microservices-9eaadde79748) by Chameera Dulanga (2021)

## Installation

The following command is used for install this package into your golang project

```sh
go get github.com/Muruyung/go-saga-func@latest
```

## Getting Started

How to implement it is quite simple, you only need to add step for execution function and rollback function. The following code is an example of its implementation:

### Example condition without error

This code has output ``Result: 25``

```go
package main

import (
    "fmt"
    sagafunc "github.com/Muruyung/go-saga-func"
)

func main() {
    var (
        saga         = sagafunc.NewSaga()
        exampleValue = 10
		err          error
    )

    err = sagaTest.AddStep(&saga.Step{
		ProcessName: "step 1",
		ExecutionFunc: func() error {
			exampleValue += 5
			return nil
		},
		RollbackFunc: func() error {
			exampleValue -= 5
			return nil
		},
	})
    if err != nil {
        fmt.Println(err)
        return
    }

	err = sagaTest.AddStep(&saga.Step{
		ProcessName: "step 2",
		ExecutionFunc: func() error {
			exampleValue += 10
			return nil
		},
		RollbackFunc: func() error {
			exampleValue -= 10
			return nil
		},
	})
    if err != nil {
        fmt.Println(err)
        return
    }

    err = sagaTest.ExecStart()
    fmt.Printf("Result: %v\n", exampleValue)
    if err != nil {
        fmt.Println(err)
        return
    }
}
```

### Example condition with error and rollback

This code has output ``Result: 10`` and detail error

```go
package main

import (
    "fmt"
    sagafunc "github.com/Muruyung/go-saga-func"
)

func main() {
    var (
        saga         = sagafunc.NewSaga()
        exampleValue = 10
		err          error
    )

    err = sagaTest.AddStep(&saga.Step{
		ProcessName: "step 1",
		ExecutionFunc: func() error {
			exampleValue += 5
			return nil
		},
		RollbackFunc: func() error {
			exampleValue -= 5
			return nil
		},
	})
    if err != nil {
        fmt.Println(err)
        return
    }

	err = sagaTest.AddStep(&saga.Step{
		ProcessName: "step 2",
		ExecutionFunc: func() error {
			exampleValue += 10
			return fmt.Errorf("some detail error here")
		},
		RollbackFunc: func() error {
			exampleValue -= 10
			return nil
		},
	})
    if err != nil {
        fmt.Println(err)
        return
    }

    err = sagaTest.ExecStart()
    fmt.Printf("Result: %v\n", exampleValue)
    if err != nil {
        fmt.Println(err)
        return
    }
}
```
