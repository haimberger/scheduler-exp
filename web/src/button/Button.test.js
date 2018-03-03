import React from 'react';
import ReactDOM from 'react-dom';
import { IconButton } from './Button';

it('renders IconButton without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<IconButton class="edit" icon="pencil" />, div);
  ReactDOM.unmountComponentAtNode(div);
});
