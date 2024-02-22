# Thread Pool

**You want? -> You Get -> You Put it Back**

>Note: I'm following [Arpit Bhayani's](https://www.youtube.com/watch?v=NgYS6mIUYmA&ab_channel=ArpitBhayani) video


If say there are `n` concurrent requests being handled by `n` threads. Means every request getting its own thread, each request getting executed as fast as possible, and having a clear seperation of concern. 

All Hunky dory! How good na? `You happy! Me Happy! God's Happy! :p`

**But what if the no. of requests shoots up?**

Will you spin up more threads?

Do you really think that having more threads would means faster execution?

    No you won't indulge in such stupidity :p

- - -
When we have large number of threads, Each thread would be requiring its own stack, which requires memory, so memory consumption bloats up.

Since there are large number of threads, operating system will have to do, large number of context switching between them.

    > What is Context Switching?

    In operating systems, multiple processes or threads can appear to run simultaneously. The operating system achieves this by rapidly switching the CPU's attention between them. This is called context switching.

    > When a context switch occurs, the operating system must:

    Save the state: Store all the information the current thread needs to resume later (like the values in its registers, program counter, etc.). This is its "context".
    
    Load the state: Load the saved context of another thread that's ready to run.

    > Why Context Switching Matters with Many Threads?
    
    Overhead: Context switching isn't free. Saving and reloading states takes some processing time.
    
    Lots of Threads = Lots of Switching: If you have many threads, the operating system will likely spend a significant amount of time switching between them instead of actually running the threads. This can lead to performance slowdown.
    
    Thread Priorities: Operating systems often use thread priorities to try and optimize context switching, ensuring high-priority threads get more CPU time.


Also the hardware gets overwhelmed when we have to process large number of requests.

Now, you understood that this is a niave way to deal with more threads, you can't just get a new threads for every request, without thinking about the overhead it brings.


>We need to limit the maximum number of threads we create.

This is where the `Thread Pool` comes to our rescue.


**<u>Real World use cases:</u>**

1. Web Servers required to handle multiple clients simulteneously.
2. Asynchronous processing of messages from broker.


## What exactly is a Thread Pool?

Thread pool is a collection of worker threads that are used to execute tasks concurrently.

Whenever we want a thread we pick one from this pool and delegate a task to it.

Once the task is complete, we add the thread back to the thread pool.

>Example: When a webserver spins up, it creates a thread when a client connects to the webserver, we pick one thread to handle the request, once the response is generated and sent, the thread is added back to the pool.

This ensures stable performance of the system.

Since, the size of the pool is limited(bounded) and hence, when your application(webserver) needs a thread but it is empty, the application waits until some thread is done and added back to the pool.

You limit the max number of threads as per your requirement and/or the specifications of your hardware.

- When you have too less number of threads, means you're not utilizing your hardware to the fullest, and if it is too large, your hardware overwhelms unnecessarily. We can do Lazy Eviction of threads but the gist is tune it to your requirement.


## Tuning the Thread Pool

There is no golden rule to tune the thread pool and it depends totally on the "application"(How dense and tense the workload is.)

Here are a few pointers to consider while deciding the size of the pool:

1. Number of processors available(2-core, 4-core, 8-core etc.)
2. Characteristic of the tasks being executed.
    - Whether it is CPU bound, I/O bound, Memory bound etc.
    - If it is Network bound, you can spin up a relatively larger threads.
    - If each request takes more time on CPU, you can have fewer threads.
