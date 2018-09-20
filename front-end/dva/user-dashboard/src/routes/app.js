import React, {PropTypes} from 'react'
import {connect} from 'dva'
import Footer from '../components/layout/footer'
import Header from '../components/layout/Header'
// import header from '../components/layout/header'
import Slider from '../components/layout/slider'
import styles from '../components/layout/main.less'
import classnames from 'classnames'
import '../components/layout/common.less'

function App({children, location, dispatch, app}) {
  console.log(children)
  const {user, sliderFold, darkTheme, isNavBar, menuPopoverVisible} = app;

  const headerProps = {
    user,
    sliderFold,
    location,
    isNavBar,
    menuPopoverVisible,
    switchMenuPopover () {
      dispatch({type: 'app/switchMenuPopover'})
    },
    logout () {
      dispatch({type: 'app/logout'})
    },
    switchSlider () {
      dispatch({type: 'app/switchSlider'})
    }
  };

  const sliderProps = {
    sliderFold,
    darkTheme,
    location,
    changeTheme () {
      dispatch({type: 'app/changeTheme'})
    }
  };


  const logined =
    <div
      className={classnames(styles.layout, {[styles.fold]: isNavBar ? false : sliderFold}, {[styles.withnavbar]: isNavBar})}>
      {!isNavBar ? <aside className={classnames(styles.slider, {[styles.light]: !darkTheme})}>
        <Slider {...sliderProps} />
      </aside> : ''}
      <div className={styles.main}>
        <Header {...headerProps} />
        <div className={styles.container}>
          <div className={styles.content}>
            {children}
          </div>
        </div>
        <Footer />
      </div>
    </div>
  ;

  return (
    <div>{ logined  }</div>
  )

}

App.propTypes = {
  children: PropTypes.element.isRequired,
  location: PropTypes.object,
  dispatch: PropTypes.func,
  loading: PropTypes.object,
  sliderFold: PropTypes.bool,
  darkTheme: PropTypes.bool
};

export default connect(({app}) => ({app}))(App)

