import React, { Component } from "react";
import { Form } from "@unform/web";
import Input from "../Input/Input";
import "./SignUpCard.css";

class NavBar extends Component {
  render() {
    function handleSubmit(data) {
      alert(JSON.stringify(data));
      //TODO call api register
    }

    return (
      <div className="signUpCard">
        <header className="cardHeader">
          <h1>Personalize o Seu!</h1>
        </header>
        <Form className="cardForm" onSubmit={handleSubmit}>
          <div className="inputGroup">
            Nome
            <Input name="Nome" type="text" placeholder="Arthur Paiva" />
          </div>
          <div className="inputGroup">
            LinkedIn
            <Input
              name="LinkedIn"
              type="text"
              placeholder="linkedin.com/in/arthur-paiva-982405199/"
            />
          </div>
          <div className="inputGroup">
            GitHub
            <Input
              name="GitHub"
              type="text"
              placeholder="github.com/arthurpaivat"
            />
          </div>

          <button type="submit">Cadastrar-se</button>
        </Form>
      </div>
    );
  }
}

export default NavBar;
