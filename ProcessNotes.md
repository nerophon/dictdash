# Dictionary Dash: Process Notes

Nicholas Bellerophon
<br>
professional@nicholasbellerophon.com

<br>
##How did you approach solving this problem?

I first read the brief, then created a notes document to keep a record of my process; I did some initial speculation about how the solution might be judged (between the lines) and about some algorithms that might be appropriate based upon my own thought patterns as I ran simple examples in my head.

I then stopped thinking about the problem and went to bed; in the morning the general outlines of a solution were apparent: graph + search.

Later, I performed internet research on the problem, arriving at the conclusion that it was a special case of single-pair shortest-path, where all edges are cost one and with the obvious admissible heuristic being min-letter swaps, ignoring whether or not the word exists in the dictionary; therefore an ideal case for A*.

The next step was to choose a language and consider the structure of the data given the performance requirements; I did some back-of-napkin calculations with broad assumptions about max dictionary size to check the memory requirements of the graph structures I envisaged were reasonable.

Finally, I began work in Go, building the application in an exploratory manner, but as much as possible following a test-driven approach and writing basic implementations prior to optimising or refactoring; as a last step I wrote some simple benchmarking functions and did some basic concurrency work based upon the results of those, multiplying grapher performance by (nearly) the number of cores in my processor.

<br>
##How did you check that your solution is correct?

In theory, A* with admissible heuristic guarantees optimal path.

It is clear that the chosen heuristic is admissible, insomuch as it never overestimates. This is because no matter what nodes are in the dictionary, the minimum number of transforms is never lower than the count of letter differences.

Therefore, the only possibility of a failure in the solution is in the implementation; however tests (both written and manual) have so far failed to show a bug in the implementation that would cause the application to fail to return the optimal path, or that the heuristic is not admissible. The tests also show that the algorithm is deterministic even for multiple optimal paths of the same length; this gives added confidence.

<br>
##Specify any assumptions that you have made.

The main assumption is that the input dictionary file is formatted correctly as a series of whitespace-delimited words containing only lowercase latin utf8 alphabet characters from `a` to `z` inclusive. The implementation does handle repeated words, and (possibly beyond the remit of the test) also handles words of different lengths in the same input dictionary.

Internally, some public functions on the critical path are optimised to not fully check the validity of their input arguments, but this is noted where it occurs. In the current usage of these functions, the tests prove that such invalid inputs are not received.

Probably the biggest assumption I have made is that the best search performance for this problem will be achieved by using the A* algorithm. Given more time, more research could potentially be done in this area to see if there are any faster ones for this special case.