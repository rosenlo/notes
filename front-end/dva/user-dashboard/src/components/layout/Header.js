import React from 'react'
import {Menu, Icon, Popover} from 'antd'
import styles from './main.less'
import Menus from './menu'
import {config} from '../../utils/'

const SubMenu = Menu.SubMenu;
function Header({user, logout, switchSlider, sliderFold, isNavBar, menuPopoverVisible, location, switchMenuPopover}) {
  let handleClickMenu = e => e.key === 'logout' && logout();
  const menusProps = {
    sliderFold: false,
    darkTheme: false,
    isNavBar,
    handleClickNavMenu: switchMenuPopover,
    location
  };
  return (
    <div className={styles.header}>
      {isNavBar
        ? <Popover placement='bottomLeft' onVisibleChange={switchMenuPopover} visible={menuPopoverVisible}
                   overlayClassName={styles.popovermenu} trigger='click' content={<Menus {...menusProps} />}>
          <div className={styles.sliderButton}>
            <Icon type='bars'/>
          </div>
        </Popover>
        : <div className={styles.sliderButton} onClick={switchSlider}>
          <Icon type={sliderFold ? 'menu-unfold' : 'menu-fold'}/>
        </div>}


      <Menu className='header-menu' mode='horizontal' onClick={handleClickMenu}>
        <SubMenu style={{
          float: 'right'
        }} title={< span > <Icon type='user'/>
           </span>}>
          <Menu.Item key='logout'>
            <a>注销</a>
          </Menu.Item>
        </SubMenu>
      </Menu>
    </div>
  )
}

export default Header
