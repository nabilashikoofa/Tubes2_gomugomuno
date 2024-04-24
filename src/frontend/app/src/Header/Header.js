import React from 'react'
import './Header.css'
import cloud from '../MINICLOUD.png';

export default function Header(){
    return <div className='header'>
        <div className="title-container">
            <h1>gomugomuno's WikiRace</h1>
        </div>
        <div className='menu'>
            <a href="#htu">Input</a>
            <a href="#res">Result</a>
            <a href="#aboutus">About Us</a>
        </div>
        <img src={cloud}></img>
    </div>
}