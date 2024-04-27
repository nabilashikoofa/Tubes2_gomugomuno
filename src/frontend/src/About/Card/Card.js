import React from "react";
import "./Card.css";

export default function Card({ name , nim , imageSrc, link }) {
  const handleClick = () => {
    // Check if a link is provided
    if (link) {
      // If link is provided, navigate to it
      window.open(link, "_blank"); // Opens link in a new tab
    }
  };
  return (
    <div className="card" onClick={handleClick}>
        <img src={imageSrc} alt={name} />
        <p>{name}</p>
        <p>{nim}</p>
    </div>
  );
}
