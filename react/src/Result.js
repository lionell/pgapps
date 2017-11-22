import React, { Component } from 'react';

class Result extends Component {
  render() {
    if (this.props.respJson.error === '') {
        return this.renderResult()
    }
    return this.renderFail()
  }

  renderResult() {
    const header = 
        this.props.respJson.result.header
        .map((v, i) => <th key={i}>{v}</th>);
    const rows = [];
    for (let i = 0; i < this.props.respJson.result.rows.length; i++) {
        const row = this.props.respJson.result.rows[i]
            .map((v, i) => <td key={i}>{v}</td>);
        rows.push(<tr key={i}>{row}</tr>)
    }
    return (
        <div className="container result">
            <div className="alert alert-success">
                {this.props.respJson.query}
            </div>
            <table className="table table-striped table-bordered">
                <thead><tr>{header}</tr></thead>
                <tbody>{rows}</tbody>
            </table>
        </div>
    )
  }

  renderFail() {
    return (
        <div className="container result">
            <div className="alert alert-danger">
                {this.props.respJson.error}
            </div>
        </div>
    )
  }
}

export default Result;
