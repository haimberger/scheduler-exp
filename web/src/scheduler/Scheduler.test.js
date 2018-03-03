import React from 'react';
import ReactDOM from 'react-dom';
import Scheduler from './Scheduler';

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<Scheduler />, div);
  ReactDOM.unmountComponentAtNode(div);
});
