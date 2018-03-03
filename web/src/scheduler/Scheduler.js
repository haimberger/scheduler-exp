import React, { Component } from 'react';
import './Scheduler.css';

class Scheduler extends Component {
  render() {
    return (
      <div className="Scheduler">
        <div className="Scheduler-tasks"></div>
        <a href="#pending">Show pending tasks...</a>
      </div>
    );
  }
}

export default Scheduler;
