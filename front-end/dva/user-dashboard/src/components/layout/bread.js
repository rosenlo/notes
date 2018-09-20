import React, { PropTypes } from 'react'
import { Breadcrumb, Icon } from 'antd'
import styles from './main.less'
import { menu } from '../../utils'

let pathSet = [];
const getPathSet = function (menuArray, parentPath) {
  parentPath =  '/'|| parentPath || '/';
  menuArray.map(item => {
    pathSet[(parentPath + item.key).replace(/\//g, '-').hyphenToHump()] = {
      path: parentPath + item.key,
      name: item.name,
      icon: item.icon || '',
      clickAble: item.clickAble === undefined
    };
    if (item.child) {
      getPathSet(item.child, parentPath + item.key + '/')
    }
  })
};
getPathSet(menu);
function Bread ({ location }) {
  let pathNames = [];
  location.pathname.substr(1).split('/').filter(x=>isNaN(x)).map((item, key) => {
    if (key > 0) {
      pathNames.push(('-' + item).hyphenToHump())
    } else {
      pathNames.push(('-' + item).hyphenToHump())
    }
  });
  const breads = pathNames.filter((k,v)=>(k in pathSet)).map((item, key) => {
      return (
        <Breadcrumb.Item key={key}
                         {...((pathNames.length - 1 === key)
                         || !pathSet[item].clickAble) ? '' : { href: '#' + pathSet[item].path }}>
          {pathSet[item].icon
            ? <Icon type={pathSet[item].icon} />
            : ''}
          <span>{pathSet[item].name}</span>
        </Breadcrumb.Item>
      )
  });
  return (
    <div className={styles.bread}>
      <Breadcrumb>
        <Breadcrumb.Item href='#/'><Icon type='home' />
          <span>主页</span>
        </Breadcrumb.Item>
        {breads? breads:''}
      </Breadcrumb>
    </div>
  )
}

Bread.propTypes = {
  location: PropTypes.object
};

export default Bread
