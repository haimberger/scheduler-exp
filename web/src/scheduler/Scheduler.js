import React from 'react';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import './Scheduler.css';

const Home = () => (
  <div className="Home">
    <div className="Home-tasks"></div>
    <Link to="/pending">Show pending tasks...</Link>
  </div>
);

const Pending = () => (
  <div className="Pending">
    <div className="Pending-tasks"></div>
    <Link to="/">Back</Link>
  </div>
);

const Scheduler = () => (
  <Router>
    <div className="Scheduler">
      <Route exact path="/" component={Home}/>
      <Route path="/pending" component={Pending}/>
    </div>
  </Router>
);

export default Scheduler;
