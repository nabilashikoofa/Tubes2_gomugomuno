import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import Header from './Header/Header';
import Content from './Content/Content';
import About from './About/About';
import HTU from './How to Use/HTU'

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <>
    <Header/>
    <div class="blank" id='blank'></div>
    <App />
    <HTU/>
    <Content/>
    <div class="dark-overlay"></div>
    {/* <div className='tes'></div> */}
    <br></br>
    <About/>
    </>
);

