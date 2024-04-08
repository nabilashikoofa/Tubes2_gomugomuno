# Tubes2 gomugomuno - WikiRace Using BFS and IDS Algorithm

| Names                     | NIM      |
| ----------------------    |:--------:|
| Thea Josephine H          | 13522012 |
| Imam Hanif Mulyarahman    | 13522030 |
| Nabila Shikoofa Muida     | 13522069 |

## Table of Contents 💫
* [The Game](#the-algorithm-👾)
* [Requirements](#requirements-🫧)
* [Setting Up](#setting-up-🍀)
* [Usage](#how-to-run-🌿)
* [How To Use](#how-to-use-🪄)

## The Algorithm 👾
BFS (Breadth First Search)
Graph traversal algorithm that starts from the root node and explores all of its neighbors at the present depth level before moving on to the nodes at the next level of depth. Check all nodes from the same level before moving on to the next level. Using queue data structure. 
- Start from root node.
- Dequeue a node from the queue.
- Visit the dequeued node and enqueue all of its neighbors.
- Repeat until queue is empty or final destination node is found.

IDS (Iterative Deepening Search)
IDS is a combination of depth-first search (DFS) and breadth-first search (BFS). It systematically searches deeper into the graph with each iteration until the desired node is found. Starts with a depth limit of 0 and gradually increases it with each iteration until the goal is found.
- Start from depth limit of 0.
- Perform DFS with the current depth limit.
- If the goal is found, stop the search and return the solution.
- If the goal is not found and there are still nodes to explore at the current depth limit, increment the depth limit and repeat the search.
- Repeat until final destination node is found or the entire search space is explored.

## Requirements 🫧


## Setting Up 🍀
- Clone this repository on your terminal `https://github.com/nabilashikoofa/Tubes2_gomugomuno.git`

## How to Run 🌿


## How to Use 🪄


## Thankyou for trying our program :>