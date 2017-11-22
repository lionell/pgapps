import React, { Component } from 'react';
import './App.css';

import Result from './Result.js';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      messages: [],
      query: '',
    }
  }

  render() {
    const results = this.state.messages
      .map((msg, index) => <Result key={index} respJson={msg} />)
      .reverse();
    return (
      <div className="container">
        <div className="row">
          <div className="col-md-1">
          </div>
          <div className="col-md-10">
            <h1>Welcome to React</h1>
            <p className="lead">Npm is hosting this demo, while HTTP server written in Golang is in the back.</p>

            <div className="form-group">
              <textarea className="form-control" onChange={(event) => this.setState({query: event.target.value})}/>
            </div>
            <div className="form-group">
              <button className="btn btn-primary" onClick={() => this.query()}>Submit</button>
            </div>
          </div>
          <div className="col-md-1">
          </div>
        </div>
        <div className="row">
          <div className="col-md-1">
          </div>
          <div className="col-md-10">
            {results}
          </div>
          <div className="col-md-1">
          </div>
        </div>
      </div>
    );
  }

  query() {
    fetch('http://localhost:8080/query?q=' + this.state.query)
    .then(resp => resp.json())
    .then(respJson => {
      console.log(respJson);
      this.state.messages.push(respJson);
      this.setState(this.state);
    })
    .catch(err => console.error(err))
  }
}

export default App;
