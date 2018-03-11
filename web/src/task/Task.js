import React from 'react';
import './Task.css';
import { IconButton } from '../button/Button';
import Editor from '../editor/Editor';
import Modal from '../modal/Modal';

class Task extends React.Component {
  constructor(props) {
    super(props);
    this.state = { editorIsOpen: false, menuIsOpen: false };
    this.toggleEditor = this.toggleEditor.bind(this);
    this.toggleMenu = this.toggleMenu.bind(this);
  }

  render() {
    const icon = this.getIcon();
    return (
      <div className={`Task ${this.getClasses()}`}>
        <table className="Task-title" onClick={this.toggleMenu}>
          <tbody>
            <tr>
              <td className="text">{this.props.title}</td>
              <td className="icon">
                <img src={`icons/${icon}.svg`} alt={icon} />
              </td>
            </tr>
          </tbody>
        </table>
        <TaskMenu
          show={this.state.menuIsOpen}
          status={this.props.status}
          link={this.props.link}
          onEdit={this.toggleEditor}
        />
        <Modal
          show={this.state.editorIsOpen}
          onClose={this.toggleEditor}
        >
          <Editor title="Edit Task" task={this.props} />
        </Modal>
      </div>
    );
  }

  getClasses() {
    if (this.props.status === 'done') {
      return 'Task-done';
    } else if (this.props.status === 'wip') {
      return 'Task-wip';
    } else if (this.props.canPreempt) {
      return 'Task-pending Task-preempt';
    } else if (this.props.startDate) {
      return 'Task-pending Task-scheduled';
    } else {
      return 'Task-pending';
    }
  }

  getIcon() {
    if (this.props.status === 'done') {
      return 'check';
    } else if (this.props.status === 'wip' || this.props.canPreempt) {
      return 'star';
    } else if (this.props.startDate) {
      return 'clock';
    } else {
      return 'loop';
    }
  }

  toggleEditor() {
    this.setState({ editorIsOpen: !this.state.editorIsOpen });
  }

  toggleMenu() {
    this.setState({ menuIsOpen: !this.state.menuIsOpen });
  }
}

class TaskMenu extends React.Component {
  render() {
    if (!this.props.show) {
      return null;
    }
    switch (this.props.status) {
      case 'done': return <DoneMenu link={this.props.link} onEdit={this.props.onEdit} />;
      case 'wip': return <WIPMenu link={this.props.link} onEdit={this.props.onEdit} />;
      case 'pending': return <PendingMenu link={this.props.link} onEdit={this.props.onEdit} />;
      default: throw new Error(`invalid task status: ${this.props.status}`);
    }
  }
}

class DoneMenu extends React.Component {
  render() {
    return (
      <div className="Task-menu">
        <IconButton class="link" icon="link" link={this.props.link} />
        <IconButton class="edit" icon="pencil" onClick={this.props.onEdit} />
        <IconButton class="pause" icon="arrow-right" />
      </div>
    );
  }
}

class WIPMenu extends React.Component {
  render() {
    return (
      <div className="Task-menu">
        <IconButton class="done" icon="check" />
        <IconButton class="link" icon="link" link={this.props.link} />
        <IconButton class="edit" icon="pencil" onClick={this.props.onEdit} />
        <IconButton class="pause" icon="arrow-right" />
        <IconButton class="cancel" icon="x" />
      </div>
    );
  }
}

class PendingMenu extends React.Component {
  render() {
    return (
      <div className="Task-menu">
        <IconButton class="link" icon="link" link={this.props.link} />
        <IconButton class="edit" icon="pencil" onClick={this.props.onEdit} />
        <IconButton class="up" icon="arrow-up" />
        <IconButton class="down" icon="arrow-down" />
        <IconButton class="cancel" icon="x" />
      </div>
    );
  }
}

export default Task;
