# BFS (Breadth First Search)
Also known as Level Order Traversal. Technique is defined as a method to traverse a Tree such that all nodes present in the same level are traversed completely before traversing the next level. 

Main idea: Use a queue to store the nodes at each level. As you process each node, you also remove it from the queue. After you remove it, process, etc, you can then push it's children (left, right) onto the queue. This order of operations will result in each level being processed 'in order'. Level by level, with the correct order of nodes from left to right. 

## Example
Initial Tree:

        10
       /  \
      5    15
     / \   / \
    3   7 13 18


Step 1: Start with root node (Level 0):

Queue: [10]  --> Process 10, pushes children onto queue

Level 0: [10]

Step 2: Process the children of 10 (Level 1):

Queue: [5, 15]  --> Process 5, 15, pushes children onto queue (3, 7) (then, 13, 18)
    NOTE: Remember, 5 is processed first, children are then pushed. Next iteration, 15 is processed then it's children are pushed.

Level 1: [5, 15]

Step 3: Process the children of 5 and 15 (Level 2):

Queue: [3, 7, 13, 18]  --> Process 3, 7, 13, 18

Level 2: [3, 7, 13, 18]


Final Output of BFS:
Level 0: [10]
Level 1: [5, 15]
Level 2: [3, 7, 13, 18]


       Level 0: [10]      
   -----------------------   
  |                       |  
  Level 1: [5, 15]    
   -----------------------   
  |       |         |        |   
  Level 2: [3, 7, 13, 18]   

