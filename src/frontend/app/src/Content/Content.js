import React, { useState } from "react"
import './Content.css'


export default function Content(){
    const [startNode, setStartNode] = useState("");
    const [endNode, setEndNode] = useState("");

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
        // Send these values to the server using fetch or any other method
    };
    const handleIDS = () => {
        // Send startNode and endNode values to the server
        console.log("Start Node:", startNode);
        console.log("End Node:", endNode);
        // Send these values to the server using fetch or any other method
    };
    return <div className='content' id="content">
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
}