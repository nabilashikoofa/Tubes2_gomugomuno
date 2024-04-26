import React from "react"
import './About.css'
import Card from "./Card/Card";
import Thea from "../Assets/thea.jpg"
import Imam from "../Assets/imam.jpg"
import Nabila from "../Assets/nabila.jpg"


export default function About(){
    return <div className='aboutus' id="aboutus">
        <h3>About Us</h3>
        <div className="setofcards">
            <Card name="Thea Josephine Halim" nim="13522012" imageSrc={Thea}/>
            <Card name="Imam Hanif Mulyarahman" nim="13522030" imageSrc={Imam}/>
            <Card name="Nabila Shikoofa Muida" nim="13522069" imageSrc={Nabila}/>
        </div>
    </div>
}