import React from 'react';
import './Editor.css';
import { IconButton } from '../button/Button'

class Editor extends React.Component {
  render() {
    return (
      <div className="Editor">
        <div className="Editor-title">{this.props.title}</div>
        <form>
          <input type="text" name="title" placeholder="Task title (max. 60 characters)" />
          <input type="text" name="link" placeholder="Link to Trello card, Slack message, or similar" />
          <input type="text" name="duration" placeholder="Duration estimate, e.g. &quot;4 hours&quot;" />
          <input type="text" name="assigner" placeholder="Your name (in case of questions)" />
          <div className="Editor-options">
            <IconButton class="schedule" icon="clock" active="false" />
            <IconButton class="preempt" icon="star" active="true" />
          </div>
          <div className="Editor-actions">
            <input type="button" className="cancel" value="Cancel" />
            <input type="submit" value="Done" />
          </div>
        </form>
      </div>
    );
  }
}

export default Editor;
