import React from 'react'
import {Menu, Icon} from 'antd'
import {Link} from 'dva/router'
import {menu} from '../../utils'

const topMenus = menu.map(item => item.key)
const getMenus = function (menuArray, sliderFold, parentPath) {
  parentPath = parentPath || '/'
  return menuArray.map(item => {
    if (item.child) {
      return (
        <Menu.SubMenu key={item.key} title={<span>{item.icon ?
          <Icon type={item.icon}/> : ''}{sliderFold && topMenus.indexOf(item.key) >= 0 ? '' : item.name}</span>}>
          {getMenus(item.child, sliderFold, parentPath + item.key + '/')}
        </Menu.SubMenu>
      )
    } else {
      if(!item.visual) {
        return (
          <Menu.Item key={item.key}>
            <Link to={item.path ? item.path : parentPath + item.key}>
              {item.icon ? <Icon type={item.icon}/> : ''}
              {sliderFold && topMenus.indexOf(item.key) >= 0 ? '' : item.name}
            </Link>
          </Menu.Item>
        )
      }
    }
  })
};

function Menus({sliderFold, darkTheme, location, isNavBar, handleClickNavMenu}) {
  const menuItems = getMenus(menu, sliderFold)
  return (
    <Menu
      mode={sliderFold ? 'vertical' : 'inline'}
      theme={darkTheme ? 'dark' : 'light'}
      onClick={handleClickNavMenu}
      defaultOpenKeys={isNavBar ? menuItems.map(item => item.key) : []}
      defaultSelectedKeys={[location.pathname.split('/')[location.pathname.split('/').length - 1] || 'dashboard']}>
      {menuItems}
    </Menu>
  )
}

export default Menus

