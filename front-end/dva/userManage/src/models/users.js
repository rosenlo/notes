import {hashHistory} from 'dva/router';
import * as usersService from '../services/users';

export default {
  namespace: 'users',

  state: {
    list: [],
    total: null,
    loading: false,
    current: null,
    currentItem: {},
    modalVisible: false,
    modalType: 'create',
  },

  subscriptions: {
    setup({dispatch, history}) {
      history.listen(({pathname, query})=> {
        if (pathname === '/users') {
          dispatch({
            type: 'save',
            payload: {data, total: query}
          });
        }
      })
    }
  },

  effect: {
    *fetch({payload}, {select, call,put}) {
      yield put({type: 'showLoading'});
      const {data, headers} = yield call(usersService.fetch, {page});
      yield put({type: 'save', payload: {data, total: headers['x-total-count']}});

      // if (data) {
      //   yield put({
      //     type: 'save',
      //     payload: {
      //       list: data.data,
      //       total: data.page.total,
      //       current: data.page.current
      //     }
      //   })
      // }
    },
    *create() {
    },
    *'delete'() {
    },
    *update() {
    },
  },
  reducers: {
    save(state, {payload: {data: list, total}}) {
      return {...state, list, total}
    },
    showLoading(state, action){
      return {...state, loading: true};
    }, // 控制加载状态的reducer
    showModal(){
    }, // 控制modal显示的reducer
    hideModal(){
    }, // 使用服务器数据返回
    querySuccess(state, action){
      return {...state, ...action.payload, loading: false};
    },
    createSuccess(){
    },
    deleteSuccess(){
    },
    updateSuccess(){
    },
  }
}
