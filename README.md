# Tubes2 gomugomuno - WikiRace Using BFS and IDS Algorithm

| Names                     | NIM      |
| ----------------------    |:--------:|
| Thea Josephine H          | 13522012 |
| Imam Hanif Mulyarahman    | 13522030 |
| Nabila Shikoofa Muida     | 13522069 |

## Table of Contents 💫
* [The Game](#the-algorithm) 👾 
* [Tools](#tools-and-libraries) 🔨
* [Requirements](#requirements) 🫧
* [Setting Up](#setting-up) 🍀
* [How To Use](#how-to-use) 🪄

## The Algorithm
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

## Tools and Libraries
- React
- Gocolly
- D3 Graph

## Requirements
- Have npm installed with `npm install`

## Setting Up
- Clone this repository on your terminal `https://github.com/nabilashikoofa/Tubes2_gomugomuno.git`
- Go to the `frontend` directory by using `cd src/frontend`
- Type in `npm install` to start the server on your local browser
- If error occurs after npm install, type in `npm install react-d3-graph@2.6.0 --legacy-peer-deps`.
- Type in `npm start` to start the server on your local browser
- Open a new terminal, and go to backend dir using `cd src/backend`
- Type in `go run main.go` to start the server, if your firewall blocks it, select `allow`.
- You are done :>

## How to Use
- Input the start node and end node with words as the Wikipedia title page
- Choose between using BFS or IDS algorithm
- Please be patient and the result will be displayed below as a graph

## File structure
```
.
├── README.md
├── doc
│   └── gomugomuno.pdf
├── src
│   ├── frontend
│   │   ├── .gitignore
│   │   ├── package.json
│   │   ├── package-lock.json
│   │   ├── README
│   │   └── public
│   │       └── src
│   │           ├── App.js
│   │           ├── App.css
│   │           ├── index.js
│   │           ├── index.css
│   │           ├── package-lock.json
│   │           ├── package.json
│   │           ├── About
│   │           │   ├── About.js
│   │           │   └── About.css
│   │           ├── Content
│   │           │   ├── Content.js
│   │           │   └── Content.css
│   │           ├── Footer
│   │           │   ├── footer.js
│   │           │   └── footer.css
│   │           ├── GraphComponent
│   │           │   ├── GraphComponent.js
│   │           │   └── GraphComponent.css
│   │           ├── Header
│   │           │   ├── Header.js
│   │           │   └── Header.css
│   │           └── How To Use
│   │               ├── HTU.js
│   │               └── HTU.css
│   ├── backend
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── Dockerfile
│   │   ├── main.go
│   │   ├── BFS
│   │   │   ├── BFS.go
│   │   │   └── BFSfunction.go
│   │   ├── IDS
│   │   │   └── ids.go
│   │   └── Scrape
│   │       └── scraper.go
```
## Thankyou for trying our program :>
