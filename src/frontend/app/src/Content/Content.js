import React, { useState } from "react"
import './Content.css'
import GraphComponent from '../GraphComponent/GraphComponent'; 


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
                setResult(data.result);
            })
            .catch(error => {
                console.error('Error:', error);
                alert('Error fetching data');
            });
    };

    const handleIDS = () => {
        if (startNode==="" || endNode===""){
            alert('Cannot be empty string')
        }
        // Send startNode and endNode values to the server
        console.log("Start Node:", startNode);
        console.log("End Node:", endNode);

        // fetch(`http://localhost:3000/api/ids?startNode=${startNode}&endNode=${endNode}`)
        // .then(response => {
        //     if (!response.ok) {
        //         // throw new Error('Network response was not ok');
        //         alert('Link does not exist');
        //         return
        //     }
        //     return response.json();
        // })
        // .then(data => {
        //     console.log("SUCCEED");
        //     console.log(data); // Log the result received from the server
        //     // Update the UI with the result if needed
        //     setResult(data.result);
        // })
        // .catch(error => {
        //     console.error('Error:', error);
        //     alert('Error fetching data');
        // });
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
        <div className="result">
            {/* {result && (
                <div>
                    <p>Hasil:</p>
                    <pre>{JSON.stringify(result, null, 2)}</pre>
                    <GraphComponent result={result} />
                </div>
            )} */}
            {result && <GraphComponent result={result} />}
            {/* {<GraphComponent result={[['A', 'B', 'C'],['A', 'B', 'Z']]} />} */}
            {<GraphComponent result={[['A', 'C', 'Z'],['A', 'N', 'Z'],['A', 'O', 'Z']]} />}

        </div>
    </div> 
}