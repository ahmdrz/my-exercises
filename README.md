# my-exercises
My exercises !

## NOTE:

**In my exercises , There is no check for invalid input And I don't care about thread-safe variables**

## Data structure

From Wikipedia
> In computer science, a data structure is a particular way of organizing data in a computer so that it can be used efficiently.

##### List
- [x] Simple linked list [thanks to tutorialspoint](https://www.tutorialspoint.com/data_structures_algorithms/)
- [x] Doubly linked list [thanks to tutorialspoint](https://www.tutorialspoint.com/data_structures_algorithms/)
- [ ] Circular linked list [thanks to tutorialspoint](https://www.tutorialspoint.com/data_structures_algorithms/)
- [x] Lookup Table [No source needed](https://en.wikipedia.org/wiki/Lookup_table)

##### Stack
- [x] Stack using linked list

##### Queue
- [x] Queue using linked list
- [ ] Priority queue

##### Sort
- [x] Bubble sort [thanks to wikipedia](https://en.wikipedia.org/wiki/Bubble_sort)
- [x] Insertion sort [thanks to wikipedia](https://en.wikipedia.org/wiki/Insertion_sort)
- [x] Selection sort
- [x] Merge sort [see the playground](https://play.golang.org/p/Ma2GXvj3XP)
- [ ] Quick sort

##### Search
- [x] Linear search
- [x] Binary search
- [x] HashTable [simple hashtable in CPP](https://gist.github.com/ducngtuan/4332979)
- [ ] Interpolation search

##### Graph
- [ ] Depth first
- [ ] Breadth first

##### Recursion
- [x] Fibonacci
- [ ] Tower of Hanoi

##### Tree
- [ ] AVL
- [ ] Heap tree
- [x] Binary search tree

## Algorithm Design

From Wikipedia
> Algorithm design is a specific method to create a mathematical process in solving problems. Applied algorithm design is algorithm engineering.

- [x] Read [Algorithm Design](https://www.hiredintech.com/algorithm-design/)

##### Dynamic Programming
- [x] Memoization [fibonacci using dynamic programming](http://www.geeksforgeeks.org/program-for-nth-fibonacci-number/)

##### Divide and Conquer
- [ ] Strassen's matrix multiplication
- [x] Detect Unusual Number ( CAFE BAZAAR question )

##### Greedy
- [ ] Dijkstra's algorithm
- [ ] Prim-Kruskal's algorithm
- [x] Customer Coins.

##### Backtracking
- [ ] Reading ...

##### Branch and Bound
- [ ] Reading ...

## Patterns

From Wikipedia
> In software engineering, a software design pattern is a general reusable solution to a commonly occurring problem within a given context in software design. It is not a finished design that can be transformed directly into source or machine code. It is a description or template for how to solve a problem that can be used in many different situations. Design patterns are formalized best practices that the programmer can use to solve common problems when designing an application or system.

##### Creational patterns
- [x] Builder
- [x] Factory Method
- [x] Singleton [sync.Once package](https://golang.org/src/sync/once.go)

##### Structural patterns
- [x] Decorator [decorators in Python](https://wiki.python.org/moin/PythonDecorators)
- [ ] Module *In Python, the pattern is built into the language, and each .py file is automatically a module.*

##### Behavioral patterns
- [x] Observer [observer pattern](https://play.golang.org/p/cr8jEmDmw0)
- [x] Iterator
- [x] Momento

##### Concurrency patterns
- [x] Generators
- [x] Parallelism [merge sort using goroutines](https://github.com/duffleit/golang-parallel-mergesort/blob/master/mergesort/mergesort.go)


## Golang

##### Read more about :

##### Channels

- [x] Basic
- [x] Signals
- [x] time.Tick

##### Init function

- [x] Basic
- [x] In libraries

##### Golang compiler


##### Shared memory (sync package)

- [x] sync.Mutex
- [x] sync.Once

##### ProtoBuf in Golang


##### Atomic package

- [x] Atomic counter
- [x] States
- [x] Once

##### Regex in Golang

- [x] Validate email address in input
- [x] Find email addresses in input

##### Gorm package

- [x] Model Definition
- [ ] Associations
- [ ] Callbacks
- [ ] Transactions

##### Go concurrency patterns (Context)

- [x] Introduction
- [x] Cancel
- [x] Timeout
- [x] Deadline
- [x] Value

##### Go concurrency patterns (Pipelines and cancellation)


##### Parallelization

**Note**
[Concurrency is not parallelism](https://blog.golang.org/concurrency-is-not-parallelism)
[Slide](https://talks.golang.org/2012/concurrency.slide)

- [x] Parallelism (MD5SUM)
- [x] RSS Reader (Need more reading about parallelism)

##### Slice tricks [read](https://github.com/golang/go/wiki/SliceTricks)

- [x] Training ...

##### Structured logging

- [ ] Read Apache HTTP server logging method [read](https://httpd.apache.org/docs/1.3/logs.html)
- [ ] Implement Apache logging on Golang HTTP server.