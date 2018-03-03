import React from 'react';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import './Scheduler.css';
import Task from '../task/Task'

const Home = () => (
  <div className="Home">
    <div className="Home-tasks">
      <Task title="Lorem ipsum dolor sit amet." link="http://www.example.com" status="wip" />
      <Task title="Suspendisse lorem mi, pharetra non nibh nec, hendrerit tincidunt diam." link="http://www.example.com" status="done" />
      <Task title="Pellentesque accumsan." link="http://www.example.com" status="done" />
    </div>
    <Link to="/pending">Show pending tasks...</Link>
  </div>
);

const Pending = () => (
  <div className="Pending">
    <div className="Pending-tasks">
      <Task title="Donec sit amet fermentum lorem, at euismod lectus." link="http://www.example.com" status="pending" startDate="2018-03-03 15:00" />
      <Task title="Nullam sit amet augue risus." link="http://www.example.com" status="pending" canPreempt="true" />
      <Task title="Pellentesque accumsan, quam sed condimentum dapibus, nibh tellus." link="http://www.example.com" status="pending" />
    </div>
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
