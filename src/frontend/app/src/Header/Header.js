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
            <a href="#aboutus">About Us</a>
            <a href="https://github.com/nabilashikoofa/Tubes2_gomugomuno">Github</a>
        </div>
        <img src={cloud}></img>
    </div>
}