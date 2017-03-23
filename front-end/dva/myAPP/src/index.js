import dva, {connect} from 'dva';
import {Router, Route} from 'dva/router';
import fetch from 'dva/fetch';
import React from 'react';
import styles from './index.less';
import './index.html';
import key from 'keymaster';

// 1. Initialize
const app = dva();

// 2. Model
// Remove the comment and define your model.
//app.model({});
app.model({
  namespace: 'count',
  state: {
    record: 0,
    current: 0,
  },
  reducers: {
    add(state) {
      const newCurrent = state.current + 1;
      return {
        ...state,
        record: newCurrent > state.record ? newCurrent : state.record,
        current: newCurrent,
      };
    },
    minus(state) {
      return {
        ...state,
        current: state.current - 1
      }
    },
  },
  effects: {
    *add(action, {call, put}) {
      yield call(delay, 1000);
      yield put({type: 'minus'});
    },
  },
  subscription: {
    keyboardWatcher({dispatch}) {
      key('⌘+up, ctrl+up', () => {
        dispatch({type: 'add'})
      });
    },
  },
});

const CountApp = ({count, dispatch}) => {
  return (
    <div className={styles.normal}>
      <div className={styles.record}>Highest Record: {count.record}</div>
      <div className={styles.current}>{count.current}</div>
      <div className={styles.button}>
        <button onClick={() => {
          dispatch({type: 'count/add'});
        }}>+
        </button>
      </div>
    </div>
  );
};

// 3. Router
function mapStateToProps(state) {
  return {count: state.count};
}
const HomePage = connect(mapStateToProps)(CountApp);
// const HomePage = () => <div>Hello Dva.</div>;
app.router(({history}) =>
  <Router history={history}>
    <Route path="/" component={HomePage}/>
  </Router>
);

// 4. Start
app.start('#root');

function delay(timeout) {
  return new Promise(resolve => {
    setInterval(resolve, timeout);
  });
}
