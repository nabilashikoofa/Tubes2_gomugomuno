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

    const handleSubmit = () => {
        // Send startNode and endNode values to the server
        console.log("Start Node:", startNode);
        console.log("End Node:", endNode);
        // Send these values to the server using fetch or any other method
    };
    return <div className='content' id="content">
            <div className="inputform">
                <label htmlFor="startNode">Start Node:</label>
                <input
                    type="text"
                    id="startNode"
                    value={startNode}
                    onChange={handleStartInputChange}
                />
                <label htmlFor="endNode">End Node:</label>
                <input
                    type="text"
                    id="endNode"
                    value={endNode}
                    onChange={handleEndInputChange}
                />
                <button onClick={handleSubmit}>Submit</button>
            </div>
        </div>
}