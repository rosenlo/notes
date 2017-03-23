import React, {PropTypes} from 'react';
import styles from './UserSearch.less';

function UserSearch({
  form, field, keyword,
  onSearch,
  onAdd
}) {
  return (
    <div className={styles.normal}>
      <div className={styles.search}>
      </div>
      <div className={styles.create}>
        <button type="ghost" onClick={onAdd}>添加</button>
      </div>
    </div>
  )
}

export default () => <div>user search</div>;
