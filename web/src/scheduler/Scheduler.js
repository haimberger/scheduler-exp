import React from 'react';
import { BrowserRouter as Router, Route, Link, Switch } from 'react-router-dom';
import './Scheduler.css';
import Editor from '../editor/Editor';
import Modal from '../modal/Modal';
import Task from '../task/Task';

class TaskList extends React.Component {
  constructor() {
    super();
    this.state = { editorIsOpen: false };
    this.openEditor = this.openEditor.bind(this);
    this.closeEditor = this.closeEditor.bind(this);
  }

  openEditor() {
    this.setState({ editorIsOpen: true });
  }

  closeEditor() {
    this.setState({ editorIsOpen: false });
  }
}

class Home extends TaskList {
  render() {
    return (
      <div className="TaskList">
        <div className="TaskList-tasks">
          <Task title="Lorem ipsum dolor sit amet." link="http://www.example.com" status="wip" />
          <Task title="Suspendisse lorem mi, pharetra non nibh nec, hendrerit tincidunt diam." link="http://www.example.com" status="done" />
          <Task title="Pellentesque accumsan." link="http://www.example.com" status="done" />
        </div>
        <Link to="/pending">Show pending tasks...</Link>
      </div>
    )
  }
}

class Pending extends TaskList {
  render() {
    return (
      <div className="TaskList">
        <div className="TaskList-tasks">
          <Task title="Donec sit amet fermentum lorem, at euismod lectus." link="http://www.example.com" status="pending" startDate="2018-03-03 15:00" />
          <Task title="Nullam sit amet augue risus." link="http://www.example.com" status="pending" canPreempt="true" />
          <Task title="Pellentesque accumsan, quam sed condimentum dapibus, nibh tellus." link="http://www.example.com" status="pending" />
        </div>
        <div className="Pending-new" onClick={this.openEditor}>
          <img src="/icons/plus.svg" alt="new" />
        </div>
        <Link to="/">Back</Link>
        <Modal
          show={this.state.editorIsOpen}
          onClose={this.closeEditor}
        >
          <Editor title="New Task" colour="yellow" />
        </Modal>
      </div>
    )
  }
}

const Scheduler = () => (
  <Router>
    <div className="Scheduler">
      <Switch>
        <Route exact path="/" component={Home} />
        <Route path="/pending" component={Pending} />
      </Switch>
    </div>
  </Router>
);

export default Scheduler;
