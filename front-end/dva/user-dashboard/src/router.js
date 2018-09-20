import React from 'react';
import {Router, Route} from 'dva/router';
import App from './routes/app';
import {services} from './utils/config'
import users from './routes/users'


function RouterConfig({history}) {
  return (
    <Router history={history}>
      <Route path="/" component={App}/>
      <Route path="/users" component={users}/>
    </Router>
  );
}

function routerMapping(model, path, name, requirePath) {
  return {
    path: path ? path : model,
    name: name ? name : model,
    getComponent (nextState, cb) {
      require.ensure([], require => {
        cb(null, require(requirePath ? requirePath : './routes/' + model))
      })
    }
  }
}

const childRoutes = services.map(service => {
  return routerMapping(service)
});

// childRoutes.splice(services.length, 0, '', routerMapping('error', '*'));

// export default RouterConfig;
export default function ({history}) {
  const routes = [
    {
      path: '/',
      component: App,
      getIndexRoute (nextState, cb) {
        require.ensure([], require => {
          cb(null, require('./routes/app'))
        })
      },
      childRoutes
    }
  ];
  console.log(routes[0].childRoutes)

  return <Router history={history} routes={routes}/>
}
