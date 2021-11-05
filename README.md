# common-tools-haonan
common tools like redis, goroutine etc, for guohaonan.shy

Redis Study:

Set:
1. The inner implementation of set is list or hashtable. 
If elements stored in set are few and little or elements are type of int, it will store in an array with specific variables.
There are three kinds of variables including variables recording forms of encoding, length of entire array and contents.
The form of encoding is constantly change if new ones are added.
The number of elements in arrays is limited, or will transform to store in a hashtable.
2. 

