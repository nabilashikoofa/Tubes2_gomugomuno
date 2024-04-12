import React from "react";
import "./Card.css";

export default function Card({ name , nim , imageSrc }) {
  return (
    <div className="card">
        <img src={imageSrc} alt={name} />
        <p>{name}</p>
        <p>{nim}</p>
    </div>
  );
}
