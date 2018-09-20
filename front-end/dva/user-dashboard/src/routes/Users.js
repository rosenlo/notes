import React from 'react';
import {connect} from 'dva';
import styles from './users.css';
import UsersComponent from '../components/users/users';

function Users({location}) {
  return (
    <div className={styles.normal}>
      <UsersComponent />
    </div>
  );
}

Users.propTypes = {};

export default connect(({users}) => ({users}))(Users);
