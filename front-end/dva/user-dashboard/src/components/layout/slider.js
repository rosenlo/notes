import React from 'react'
import styles from './main.less'
import { config } from '../../utils/'
import Menus from './menu'

function Slider ({ sliderFold, darkTheme, location }) {
  const menusProps = {
    sliderFold,
    darkTheme,
    location
  }
  return (
    <div>
      <div className={styles.logo}>
        <img src={config.logoSrc} />
        {sliderFold ? '' : <span>{config.logoText}</span>}
      </div>
      <Menus {...menusProps} />
    </div>
  )
}

export default Slider
