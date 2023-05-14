import React from 'react';
import { Link } from 'react-router-dom';
export const AdminPanel = () => {
  return (
    <div className="container text-center">
      <div className="position-absolute top-50 start-50 translate-middle">
        <div className="row">
          <Link className="col" style={{ margin: 20 }} to={'/VideoAdd'}>
            {' '}
            Add Video
          </Link>
          <Link className="col" style={{ margin: 20 }} to={'/VideoList'}>
            {' '}
            Manage Videos
          </Link>
        </div>
        <div className="row">
          <Link className="col" style={{ margin: 20 }} to={'/UserList'}>
            {' '}
            Manage User
          </Link>
          <Link className="col" style={{ margin: 20 }} to={'/CategoriAdd'}>
            {' '}
            CategoriAdd
          </Link>
        </div>
      </div>
    </div>
  );
};
