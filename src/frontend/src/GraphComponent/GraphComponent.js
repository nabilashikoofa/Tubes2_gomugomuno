import React from "react";
import './GraphComponent.css'
import { Graph } from "react-d3-graph";

class GraphComponent extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            dataArray: this.props.result
        };
}

    // Fungsi untuk mengonversi array of array menjadi data graf
    convertArrayToGraphData = (dataArray) => {
        const nodes = [];
        const links = [];

        dataArray.forEach((array) => {
            array.forEach((element, index) => {
                const nodeName = element.replace(/_/g, ' ');
                if (!nodes.find((node) => node.id === nodeName)) {
                    nodes.push({ id: nodeName });
                }

                if (index < array.length - 1) {
                    const sourceName = array[index].replace(/_/g, ' '); 
                    const targetName = array[index + 1].replace(/_/g, ' '); 
                    links.push({ source: sourceName, target: targetName });
                }
            });
        });
        return { nodes, links };
    };

    componentDidUpdate(prevProps) {
        if (this.props.result !== prevProps.result) {
            this.setState({ dataArray: this.props.result });
        }
    }

    render() {
        const { dataArray } = this.state;
        const graphData = this.convertArrayToGraphData(dataArray);

        // Konfigurasi graf
        const myConfig = {
            width: (window.innerWidth) * 0.77,
            height: (window.innerHeight) * 0.55, 
            nodeHighlightBehavior: true,
            directed: true,
            node: {
                color: "lightblue",
                size: 500,
                highlightStrokeColor: "blue",
            },
            link: {
                highlightColor: "lightblue",
            },
        };

        return (
            <div className="graph">
                {/* <h1>Graph</h1> */}
                <Graph
                    id="graph-id" // id wajib
                    data={graphData}
                    config={myConfig}
                />
            </div>
        );
    }
}

export default GraphComponent;
