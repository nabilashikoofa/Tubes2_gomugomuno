import React from "react"
import './About.css'
import Card from "./Card/Card";
import ExImage from "../Nynne.jpg"

export default function About(){
    return <div className='aboutus' id="aboutus">
        <h3>About Us</h3>
        <div className="setofcards">
            <Card name="Thea Josephine Halim" nim="13522012" imageSrc={ExImage}/>
            <Card name="Imam Hanif Mulyarahman" nim="13522030" imageSrc={ExImage}/>
            <Card name="Nabila Shikoofa Muida" nim="13522069" imageSrc={ExImage}/>
        </div>
    </div>
}