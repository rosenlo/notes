import  {routerRedux} from 'dva/router'
const modelName = 'app';
export default {
  namespace: modelName,
  state: {
    login: false,
    loading: false,
    user: {},
    loginButtonLoading: false,
    menuPopoverVisible: false,
    sliderFold: localStorage.getItem('antdAdminSliderFold') === 'true',
    darkTheme: localStorage.getItem('antdAdminDarkTheme') !== 'false',
    isNavBar: document.body.clientWidth < 769
  },
  subscriptions: {
    setup ({dispatch, history}) {
      history.listen(location => {
      });
      window.onresize = function () {
        dispatch({
          type: 'changeNavBar'
        })
      }
    }
  },
  effects: {
    *switchSlider ({}, {put}) {
      yield put({type: 'handleSwitchSlider'})
    },
    *changeTheme ({}, {put}) {
      yield put({type: 'handleChangeTheme'})
    },
    *changeNavBar ({}
      , {put}) {
      if (document.body.clientWidth < 769) {
        yield put({type: 'showNavBar'})
      } else {
        yield put({type: 'hideNavBar'})
      }
    },
    *switchMenuPopover ({}, {put}) {
      yield put({type: 'handleSwitchMenuPopover'})
    }
  },
  reducers: {
    showLoading (state) {
      return {...state, loading: true}
    },
    hideLoading (state) {
      return {...state, loading: false}
    },
    handleSwitchSlider (state) {
      localStorage.setItem('antdAdminSliderFold', !state.sliderFold);
      return {...state, sliderFold: !state.sliderFold}
    },
    handleChangeTheme (state) {
      localStorage.setItem('antdAdminDarkTheme', !state.darkTheme);
      return {...state, darkTheme: !state.darkTheme}
    },
    showNavBar (state) {
      return {...state, isNavBar: true}
    },
    hideNavBar (state) {
      return {...state, isNavBar: false}
    },
    handleSwitchMenuPopover (state) {
      return {...state, menuPopoverVisible: !state.menuPopoverVisible}
    }
  }
}
