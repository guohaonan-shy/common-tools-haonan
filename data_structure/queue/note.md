Normal Queue implementation:

Normal queue ADT is based on compacted array, 
which has a high efficiency in inserting elements at the tail of array
but low efficiency in popping elements from front of array
because queue has to move the rest elements to the left indices with time complexity O(N).

Therefore, we implement queue by introducing front and end pointer that we can use constant complexity to insert and remove elements.
When the length is equal with the capacity, queue will trigger extension and copy the elements into new array. 
And in the process of copying, the queue will reorder the elements from front to end associated with index order in array.

Of course, this queue can also be implemented in single linked list.