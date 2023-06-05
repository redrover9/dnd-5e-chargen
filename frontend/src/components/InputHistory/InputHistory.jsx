import React, { Component } from "react";
import "./InputHistory.scss";

class InputHistory extends Component {
  render() {
    const messages = this.props.inputHistory.map((msg, index) => (
      <p key={index}>{msg.data}</p>
    ));

    return (
      <div className="InputHistory">
        <h2>Input History</h2>
        {messages}
      </div>
    );
  }
}

export default InputHistory;