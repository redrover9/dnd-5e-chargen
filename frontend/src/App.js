// App.js
import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import InputHistory from "./components/InputHistory";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      inputHistory: []
    }
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New data")
      this.setState(prevState => ({
        inputHistory: [...this.state.inputHistory, msg]
      }))
      console.log(this.state);
    });
  }
 
  send() {
    console.log("hello");
    sendMsg("hello");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <InputHistory inputHistory={this.state.inputHistory} />
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }

}

export default App;