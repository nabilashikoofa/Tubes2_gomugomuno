import React from 'react'
import './HTU.css'

export default function HTU(){
    return <div className='htu' id="htu">
        <div className='htucontent'>
            <h3>How to Use</h3>
            <p>1. Input the start node and end node with words as the Wikipedia title page</p>
            <p>2. Choose between using BFS or IDS algorithm</p>
            <p>3. Click the start button when you are ready!</p>
            <p>4. The result will be displayed below as a graph</p>
        </div>
    </div>
}