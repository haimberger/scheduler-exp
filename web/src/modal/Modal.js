import React from 'react';
import './Modal.css';

// Based on code found here: https://daveceddia.com/open-modal-in-react/
class Modal extends React.Component {
  render() {
    if (!this.props.show) {
      return null;
    }

    return (
      <div className="backdrop">
        <div className="modal">
          <button type="button" className="close" onClick={this.props.onClose}>&times;</button>
          {this.props.children}
        </div>
      </div>
    );
  }
}

export default Modal;
