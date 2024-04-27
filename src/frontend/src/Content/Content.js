import React, { useState } from "react"
import './Content.css'
import GraphComponent from '../GraphComponent/GraphComponent'; 


export default function Content(){
    const [startNode, setStartNode] = useState();
    const [endNode, setEndNode] = useState();
    const [result, setResult] = useState();
    const [elapsed, setElapsed] = useState();
    const [shortestlength, setShortest] = useState();
    const [numofcheckednodes, setChecked] = useState();
    const [path, setPath] = useState();

    const handleStartInputChange = (event) => {
        setStartNode(event.target.value);
    };

    const handleEndInputChange = (event) => {
        setEndNode(event.target.value);
    };

    const handleRace = (algorithm) => {
        // Send startNode and endNode values to the server
        console.log("Start Node:", startNode);
        console.log("End Node:", endNode);

        if (!startNode.trim() || !endNode.trim()) {
            alert('Start and end nodes cannot be empty');
            return;
        }
        if (startNode === endNode) {
            alert('Start and end nodes cannot be the same');
            return;
        }
    
        fetch(`http://localhost:3000/api/${algorithm}?startNode=${startNode}&endNode=${endNode}`)
            .then(response => {
                if (!response.ok) {
                    if (response.status === 0) {
                        alert('Server is not yet started');
                    } else if (response.status === 404) {
                        alert('Link does not exist');
                    } else if (response.status === 500) {
                        alert('Server error');
                    } else {
                        alert('Unknown error');
                    }
                    return;
                }
                return response.json();
            })
            .then(data => {
                if (!data) return;
                
                console.log("SUCCEED");
                console.log(data); 

                // Update the UI with the result 
                setResult(data.result);
                setElapsed(data.elapsed / 1000);
                setShortest(data.shortestlength);
                setChecked(data.numofcheckednodes);
                setPath(data.path);
            })
            .catch(error => {
                console.error('Error:', error);
                alert('Error fetching data');
            });
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
                    <button onClick={() => handleRace('bfs')}>Use BFS</button>
                    <button onClick={() => handleRace('ids')}>Use IDS</button>
                </div>
            </div>
        </div>
        <div className="result">
            {result && (
                <div>
                    <h1>Result</h1>
                    <p>Wow! We Found {path} path after checking {numofcheckednodes} nodes with {shortestlength} degrees of separation </p> 
                    <p>from {startNode} to {endNode} in {elapsed} seconds </p>
                    <GraphComponent result={result} />
                </div>
            )}
        </div>
    </div> 
}