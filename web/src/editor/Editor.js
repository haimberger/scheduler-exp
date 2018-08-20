import React from 'react';
import './Editor.css';
import { IconButton } from '../button/Button';

class Editor extends React.Component {
  constructor(props) {
    super(props);
    this.task = props.task || {};
  }

  render() {
    return (
      <div className="Editor">
        <div className={`Editor-title ${this.props.colour}`}>{this.props.title}</div>
        <form>
          <input type="text"
            name="title"
            placeholder="Task title (max. 60 characters)"
            defaultValue={this.task.title} />
          <input type="text"
            name="link"
            placeholder="Link to Trello card, Slack message, or similar"
            defaultValue={this.task.link} />
          <input type="text"
            name="duration"
            placeholder="Duration estimate, e.g. &quot;4 hours&quot;"
            defaultValue={this.task.duration} />
          <input type="text"
            name="assigner"
            placeholder="Your name (in case of questions)"
            defaultValue={this.task.assigner} />
          <div className="Editor-options">
            <IconButton class="schedule" icon="clock" colour={this.props.colour} active={this.task.startDate ? "true" : "false"} />
            <IconButton class="preempt" icon="star" colour={this.props.colour} active={this.task.canPreempt ? "true" : "false"} />
          </div>
          <button className={this.props.colour} type="submit">Done</button>
        </form>
      </div>
    );
  }
}

export default Editor;
