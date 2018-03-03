import React from 'react';
import './Button.css';

class IconButton extends React.Component {
  render() {
    let classes = `Button-${this.props.class}`;
    if (this.props.active === "true") {
      classes += " active";
    } else if (this.props.active === "false") {
      classes += " inactive";
    }
    return (
      <button type="button" className={classes}>
        {this.props.link ? <a href={this.props.link}><span></span></a> : null}
        <img src={`/icons/${this.props.icon}.svg`} alt={this.props.icon} />
      </button>
    );
  }
}

export { IconButton };
