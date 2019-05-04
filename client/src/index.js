import React, { Component } from "react";
import ReactDOM from "react-dom";
import Main from './Main'
import "./styles.css";

class App extends Component {
  render() {
    return <Main />;
  }
}

var mountNode = document.getElementById("app");
ReactDOM.render(<App />, mountNode);