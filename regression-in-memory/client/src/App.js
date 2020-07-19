import React from 'react';
import { Component } from 'react';
import './App.css';

class App extends Component {
  state = {
    data: []
  }

  onAgeCondChange = e => {
    const age = e.target.value;
    const request = new Request('http://localhost:8086' + (age === '' ? '' : `?age=${age}`));
    fetch(request).then(res => res.json()).then(({ data }) => {
      this.setState({ data });
    });
  }

  refresh = () => {
    const request = new Request('http://localhost:8086');
    fetch(request).then(res => res.json()).then(({ data }) => {
      this.setState({ data });
    });
  }

  componentDidMount() {
    this.refresh();
  }

  render() {
    return (
      <div style={{ padding: 30 }}>
        <h1>Regression in Memory</h1>
        <div style={{ border: "1px solid gray", padding: 20, margin: '0 0 20px 0' }}>
          <h2>List People</h2>
          <div>
            Age:<input style={{ margin: '0 10px' }} onChange={this.onAgeCondChange} />
          </div>
          <div>
            <table style={{ margin: '10px 0' }} border='1'>
              <thead>
                <tr>
                  <td style={{ padding: 5 }}>Name</td>
                  <td style={{ padding: 5 }}>Age</td>
                </tr>
              </thead>
              <tbody>
                {this.state.data.map(d => (
                  <tr key={d.id}>
                    <td style={{ padding: 5 }}>{d.name}</td>
                    <td style={{ padding: 5 }}>{d.age}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>

        <Create refresh={this.refresh} />
      </div>
    )
  }
}

class Create extends Component {
  state = {
    name: "",
    age: 0
  }

  onNameChange = e => {
    this.setState({ name: e.target.value });
  }

  onAgeChange = e => {
    this.setState({ age: parseInt(e.target.value) });
  }

  submit = () => {
    console.log(this.state)
    const request = new Request('http://localhost:8086', {
      method: 'POST',
      body: JSON.stringify(this.state)
    });

    fetch(request).then(res => res.json()).then(() => {
      this.props.refresh();
    });
  }

  render() {
    return (
      <div style={{ border: '1px solid gray', padding: 20 }}>
        <h2>Create People</h2>

        <input style={{ margin: '10px 10px 10px 0' }} onChange={this.onNameChange} />
        <input style={{ margin: 10 }} type='number' onChange={this.onAgeChange} />
        <button style={{ margin: 10 }} onClick={this.submit}>Create</button>
      </div>
    )
  }
}

export default App;
