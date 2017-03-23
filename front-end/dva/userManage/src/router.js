import React, {PropTypes} from 'react';
import { Router, Route } from 'dva/router';
import IndexPage from './routes/IndexPage';
import Users from './routes/Users';

function RouterConfig({ history }) {
  return (
    <Router history={history}>
      <Route path="/users" component={Users} />
    </Router>
  );
}

export default RouterConfig;
