import React, { Component } from "react";
import "./UserCard.css";
import api from "../../services/api";

class UserCard extends Component {
  state = {
    user: {},
  };

  componentDidMount() {
    console.log("oi");
    api
      .get("/user", {
        headers: { Userid: "1234" },
      })
      .then((res) => {
        const user = res.data;
        console.log("SDUASIDU");
        console.log(res.data);
        this.setState({ user });
        console.log(this.state);
      })
      .catch((error) => {
        console.log("ALOU");

        console.log(error);
      });
    console.log("ioi");
  }

  render() {
    return (
      <div className="userCardDiv">
        <div className="cardBar">
          <h1 className="userName"> {this.state.user.name} </h1>
        </div>
        <div className="cardBody">
          <div className="cardInfo">Main Role: {this.state.user.mainRole}</div>
          <div className="cardInfo">Github: {this.state.user.gitHub}</div>
          <div className="cardInfo">LinkedIn: {this.state.user.linkedIn}</div>
        </div>
      </div>
    );
  }
}

export default UserCard;
