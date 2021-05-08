import React, { Component } from "react";
import "./NavBar.css";

class NavBar extends Component {
  render() {
    return (
      <div className="navBar">
        <h1 className="logo">Porletu</h1>
        <h4 className="by">By Arthur Paiva</h4>
      </div>
    );
  }
}

export default NavBar;
