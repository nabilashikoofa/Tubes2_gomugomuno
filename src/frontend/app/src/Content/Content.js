import React, { useState } from "react"
import axios from 'axios'
import './Content.css'


export default function Content(){
    const [startNode, setStartNode] = useState();
    const [endNode, setEndNode] = useState();
    const [result, setResult] = useState();

    const handleStartInputChange = (event) => {
        setStartNode(event.target.value);
    };

    const handleEndInputChange = (event) => {
        setEndNode(event.target.value);
    };


    const handleBFS = () => {
        // Send startNode and endNode values to the server
        console.log("Start Node:", startNode);
        console.log("End Node:", endNode);
        if (!startNode.trim() || !endNode.trim()) {
            alert('Start and end nodes cannot be empty');
            return;
        }
        if (startNode===endNode) {
            alert('Start and end nodes cannot be the same');
            return;
        }
    
        fetch(`http://localhost:3000/api/bfs?startNode=${startNode}&endNode=${endNode}`)
            .then(response => {
                if (!response.ok) {
                    // throw new Error('Network response was not ok');
                    alert('Link does not exist');
                    return
                }
                return response.json();
            })
            .then(data => {
                console.log("SUCCEED");
                console.log(data); // Log the result received from the server
                // Update the UI with the result if needed
            })
            .catch(error => {
                console.error('Error:', error);
            });
    };
    const handleIDS = () => {
        if (startNode==="" || endNode===""){
            alert('Cannot be empty string')
        }
        // Send startNode and endNode values to the server
        console.log("Start Node:", startNode);
        console.log("End Node:", endNode);
        // Send these values to the server using fetch or any other method
    };
    return  <div className="maincontent">
        <div className='content' id="content">
            <h3>Enter the nodes</h3>
            <div className="inputform">
                <div className="takeinput">
                    <label htmlFor="startNode">Start Node: </label>
                    <input
                        type="text"
                        id="startNode"
                        value={startNode}
                        onChange={handleStartInputChange}
                        />
                </div>
                <br></br>
                <div className="takeinput">
                    <label htmlFor="endNode">End Node: </label>
                    <input
                        type="text"
                        id="endNode"
                        value={endNode}
                        onChange={handleEndInputChange}
                        />
                </div>
                <div className="subbuttons">
                    <button onClick={handleBFS}>Use BFS</button>
                    <button onClick={handleIDS}>Use IDS</button>
                </div>
            </div>
        </div>
        {/* {result && ( */}
        {/* nnt di pisah aja */}
        {/* <div className="result">
            <h3>Result</h3>
            <p>{result}</p>
        </div>
        )} */}
    </div> 
}