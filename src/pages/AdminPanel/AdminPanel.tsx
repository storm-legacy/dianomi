import React from 'react';
import { NavLink } from 'react-router-dom';
export const AdminPanel = () => {
  return (
    <div className="container text-center">
      <div className="position-absolute top-50 start-50 translate-middle">
        <div className="row">
          <NavLink className="col" style={{ margin: 20 }} to={'/VideoAdd'}>
            {' '}
            Dodaj Wideo
          </NavLink>
          <NavLink className="col" style={{ margin: 20 }} to={'/VideoAdd'}>
            {' '}
            Zarządzaj Wideami
          </NavLink>
        </div>
        <div className="row">
          <NavLink className="col" style={{ margin: 20 }} to={'/UserList'}>
            {' '}
            Zarządzej użytkownikami
          </NavLink>
          <NavLink className="col" style={{ margin: 20 }} to={'/VideoAdd'}>
            {' '}
            Zarządzej subskrybcjami
          </NavLink>
        </div>
      </div>
    </div>
  );
};
