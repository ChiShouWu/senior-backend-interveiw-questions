# Computer Science

1. Process vs. Thread vs worker
    <details><summary>Answer</summary>
    
    - Process: A process is an instance of a program in execution. It contains the program code and its current activity. Each process is started with a single thread, often called the primary thread, and can create additional threads from any of its threads.

    - Thread: A thread is the entity within a process that can be scheduled for execution. It contains the thread's context, a set of registers, the thread stack, and a thread environment block.

    - Worker: A worker is a thread that is managed by a thread pool. It is a thread that is waiting for a task to be assigned to it.

2. What is `Polymorphism` (dynamic binding, static binding, late binding, early binding)
    <details><summary>Answer</summary>

    - Polymorphism is the ability of an object to take on many forms. The most common use of polymorphism in OOP occurs when a parent class reference is used to refer to a child class object.

    - Dynamic binding / Late binding: The process of determining which method to call at runtime is called dynamic binding. In Java, all methods are virtually invoked. This means that all methods are invoked based on the runtime type of the object, not the compile-time type of the object.

    - Static binding / Early binding: The process of determining which method to call at compile time is called static binding. In Java, all methods are statically bound unless they are overridden in a subclass.
3. Define the terms `stack` and `heap.` What's a stack overflow?
    <details><summary>Answer</summary>

    - Stack: A stack is a data structure that stores a list of elements in a LIFO (last in, first out) fashion. In other words, the last element added to the stack will be the first one to be removed. A stack is a limited access data structure - elements can be added and removed from the stack only at the top. A stack is a useful data structure when you need to store data and ensure that the last element added is the first one to be removed.

    - Heap: A heap is a specialized tree-based data structure that satisfies the heap property: if P is a parent node of C, then the key (the value) of P is either greater than or equal to (in a max heap) or less than or equal to (in a min heap) the key of C. A common implementation of a heap is the binary heap, in which the tree is a complete binary tree.

    - stack overflow: A stack overflow occurs when a program attempts to use more stack space than is available. In other words, a stack overflow occurs when a program recurses too deeply, or when an otherwise infinite loop exists.
    <details>

4. What's the difference between cohesion and coupling?
    <details><summary>Answer</summary>

    - Cohesion: Cohesion is a measure of how strongly related the responsibilities of a class or module are. A class or module with high cohesion is focused on a single responsibility, while a class or module with low cohesion has its responsibilities spread across many different tasks.

    - Coupling: Coupling is a measure of how closely two or more classes or modules are related. A class or module with high coupling is tightly coupled to another class or module, while a class or module with low coupling is loosely coupled to another class or module.
    <details>
5. When are anonymous functions useful?
    <details><summary>Answer</summary>

    - Anonymous functions are useful when you need to pass a function as an argument to another function. They are also useful when you need to define a function inline without having to explicitly name it.
    <details>

6.  What's is usage of `static`?
    <details><summary>Answer</summary>

    - Static variables are shared by all instances of the class. Static methods can be called without creating an instance of the class. Static methods can only access static variables.
    <details>

7.  Inter-Process Communication
    <details><summary>Answer</summary>

    - Shared memory: The most common form of inter-process communication is shared memory. In shared memory, two or more processes can access the same region of memory. This region of memory is called a shared memory segment. Shared memory is the fastest form of inter-process communication, but it is also the most difficult to implement.

    - Message passing: Message passing is a form of inter-process communication in which the sending process sends a message to the receiving process. The receiving process can be on the same computer or on a different computer. Message passing is the most common form of inter-process communication in distributed systems.

    - Remote procedure call: A remote procedure call (RPC) is a form of inter-process communication in which a computer program causes a procedure to execute in a different address space (commonly on another computer on a shared network). An RPC is conceptually similar to a function call in programming languages.

    - Signals: A signal is a form of inter-process communication in which a process can send a notification to another process. Signals are used to notify a process that an event has occurred. Signals are asynchronous, which means that the sending process does not wait for the receiving process to acknowledge the signal.
    <details>
8.  Scheduler strategy
    <details><summary>Answer</summary>

    - Round-robin: In round-robin scheduling, each process is assigned a fixed time slot in a cyclic way. In other words, each process is executed for a fixed amount of time and then the CPU is taken away from the process and given to the next process in the queue. Round-robin scheduling is simple, easy to implement, and starvation-free.

    - First-come, first-served: In first-come, first-served scheduling, the process that requests the CPU first is allocated the CPU first. First-come, first-served scheduling is simple, easy to implement, and starvation-free.

    - Shortest job first: In shortest job first scheduling, the process with the smallest execution time is allocated the CPU first. Shortest job first scheduling is simple, easy to implement, and starvation-free.

    - Shortest remaining time first: In shortest remaining time first scheduling, the process with the smallest remaining execution time is allocated the CPU first. Shortest remaining time first scheduling is simple, easy to implement, and starvation-free.

    - Priority scheduling: In priority scheduling, each process is assigned a priority. The process with the highest priority is allocated the CPU first. Priority scheduling is simple, easy to implement, and starvation-free.

    - Multilevel queue scheduling: In multilevel queue scheduling, the ready queue is divided into multiple queues. Each queue has a different priority. Processes are placed in a queue based on their priority. The process with the highest priority is allocated the CPU first. Multilevel queue scheduling is simple, easy to implement, and starvation-free.

    - Multilevel feedback queue scheduling: In multilevel feedback queue scheduling, the ready queue is divided into multiple queues. Each queue has a different priority. Processes are placed in a queue based on their priority. The process with the highest priority is allocated the CPU first. If a process uses up its time slice, it is moved to a lower-priority queue. Multilevel feedback queue scheduling is simple, easy to implement, and starvation-free.

    - Real-time scheduling: In real-time scheduling, each process is assigned a deadline. The process that misses its deadline is terminated. Real-time scheduling is simple, easy to implement, and starvation-free.
    <details>
9.  What makes dead lock, how to solve it?
    <details><summary>Answer</summary>

    - Deadlock occurs when two or more processes are blocked because each process is holding a resource and waiting for another resource acquired by some other process. Deadlock can be prevented by using the following methods:

    - Mutual exclusion: A resource can be used by only one process at a time.

    - Hold and wait: A process is holding at least one resource and waiting for resources that are held by other processes.

    - No preemption: A resource cannot be taken from a process unless the process releases the resource.
    - Circular wait: There exists a circular wait for resources.
    <details>

10. What is the difference between `stack leak` and `stack overflow`?
    <details><summary>Answer</summary>

    - Stack overflow occurs when a program attempts to use more stack space than is available. In other words, a stack overflow occurs when a program recurses too deeply, or when an otherwise infinite loop exists.
  
    - Stack leak occurs when a program allocates memory on the stack and does not free it. In other words, a stack leak occurs when a program allocates memory on the stack and does not free it.
    <details>
11. What is thread pool
    <details><summary>Answer</summary>

    - A thread pool is a group of worker threads that are waiting for the tasks and jobs. A thread pool manages the worker threads and reuses them many times. It reduces the number of thread creations and thereby improves the performance of the application to a great extent.
    <details>
12. There is a huge number array about 10G, and your memory is 1G, how to sort it?
    <details><summary>Answer</summary>

    - External sorting
        1. Divide the huge file into smaller chunks that can fit into the memory.
        2. Sort each chunk using a sorting algorithm that can fit into memory.
        3. Split the sorted chunks into smaller chunks.
        4. Merge and sort the smaller chunks using a sorting algorithm that can fit into memory.
13. What the different with fork & spawn?
    <details><summary>Answer</summary>

    - fork: fork() system call creates a new process by duplicating the calling process. The new process, referred to as the child, is an exact duplicate of the calling process, referred to as the parent. The child process is created with a new process ID, and parent process is assigned the process ID of its child. The child process is created with a new process ID, and parent process is assigned the process ID of its child.

    - spawn: spawn() system call creates a new process by duplicating the calling process. The new process, referred to as the child, is an exact duplicate of the calling process, referred to as the parent. The child process is created with a new process ID, and parent process is assigned the process ID of its child. The child process is created with a new process ID, and parent process is assigned the process ID of its child.
    <details>

14. FP v.s OOP
    <details><summary>Answer</summary>

    - FP: Functional programming is a programming paradigm where programs are constructed by applying and composing functions. It is a declarative programming paradigm, which means programming is done with expressions or declarations instead of statements. In functional programming, the output value of a function depends only on its arguments, so calling a function with the same value for an argument always produces the same result. This is called referential transparency. Referential transparency allows the compiler to optimize the code by performing common subexpression elimination and other optimizations that are not available in imperative programming languages.
    
    - OOP: Object-oriented programming is a programming paradigm based on the concept of "objects", which can contain data, in the form of fields, often known as attributes; and code, in the form of procedures, often known as methods. A feature of objects is that an object's procedures can access and often modify the data fields of the object with which they are associated (objects have a notion of "this" or "self"). In OOP, computer programs are designed by making them out of objects that interact with one another. OOP languages are diverse, but the most popular ones are class-based, meaning that objects are instances of classes, which also determine their types.
    <details>