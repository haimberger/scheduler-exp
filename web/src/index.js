import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Scheduler from './scheduler/Scheduler';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<Scheduler />, document.getElementById('root'));
registerServiceWorker();
