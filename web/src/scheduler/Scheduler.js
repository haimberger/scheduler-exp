import React from 'react';
import { BrowserRouter as Router, Route, Link, Switch } from 'react-router-dom';
import './Scheduler.css';
import Task from '../task/Task'
import Editor from '../editor/Editor'

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
    <div className="Pending-new">
      <Link to="/tasks/new">
        <span></span>
        <img src="/icons/plus.svg" alt="new" />
      </Link>
    </div>
    <Link to="/">Back</Link>
  </div>
);

const Scheduler = () => (
  <Router>
    <div className="Scheduler">
      <Switch>
        <Route exact path="/" component={Home} />
        <Route path="/pending" component={Pending} />
        <Route
          path="/tasks/new"
          render={(routeProps) => (
            <Editor {...routeProps} title="New Task" />
          )}
        />
        <Route
          path="/tasks/:id"
          render={(routeProps) => (
            <Editor {...routeProps} title="Edit Task" />
          )}
        />
      </Switch>
    </div>
  </Router>
);

export default Scheduler;
