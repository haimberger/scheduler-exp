import React from 'react';
import './Button.css';

class IconButton extends React.Component {
  render() {
    return (
      <button type="button" className={`Button-${this.props.class}`}>
        {this.props.link ? <a href={this.props.link}><span></span></a> : null}
        <img src={`icons/${this.props.icon}.svg`} alt={this.props.icon} />
      </button>
    );
  }
}

export { IconButton };
