import * as usersService from '../services/users';

export default {
  namespace: 'users',
  state: {
    list: [],
    total: null,
    page: null,
  },
  reducers: {
    save(state, {payload: {data: list, total, page}}) {
      return {...state, list, total, page};
    },
  },
  effects: {
    *fetch({payload: {page = 1}}, {call, put}) {
      const {data, headers} = yield call(usersService.fetch, {page});
      yield put({
        type: 'save',
        payload: {
          data,
          total: parseInt(headers['x-total-count'], 5),
          page: parseInt(page, 1),
        }
      });
    },
    *patch({payload: {id, values}}, {call, put, select}) {
      yield call(usersService.patch, id, values);
      const page = yield select(state => state.users.page);
      yield put({type: 'fetch', payload: {page}});
    },
    *remove({payload: id}, {call, put, select}) {
      const {data, headers} = yield call(usersService.remove, id);
      // const page = 0;
      // yield put({
      //   type: 'save',
      //   payload: {
      //     data : [],
      //     total: 0,
      //     page: 0,
      //   }
      // });
      // const page = yield select(state => state.users.page);
      // yield put({type: 'fetch', payload: {page}});
    },
  },
  subscriptions: {
    setup({dispatch, history}) {
      return history.listen(({pathname, query}) => {
        if (pathname === '/users') {
          dispatch({type: 'fetch', payload: query});
        }
      });
    },
  },
};
