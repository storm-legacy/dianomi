import React from 'react';
import { Link } from 'react-router-dom';

export const AdminPanel = () => {
  return (
    <div className="container mt-5 p-5 shadow-lg">
      <div className="container">
        <h1 className="text-center mb-5">Admin Panel</h1>
        <div className="row">
          <div className="col-md-6 mb-4">
            <Link to="/VideoAdd" className="btn btn-primary btn-lg btn-block col-4">
              Add Video
            </Link>
          </div>
          <div className="col-md-6 mb-4">
            <Link to="/VideoList" className="btn btn-primary btn-lg btn-block col-4">
              Manage Videos
            </Link>
          </div>
        </div>
        <div className="row">
          <div className="col-md-6 mb-4">
            <Link to="/UserList" className="btn btn-primary btn-lg btn-block col-4">
              Manage Users
            </Link>
          </div>
          <div className="col-md-6 mb-4">
            <Link to="/CategoriAdd" className="btn btn-primary btn-lg btn-block col-4">
              Add Category
            </Link>
          </div>
        </div>
        <div className="row">
          <div className="col-md-6 mb-4">
            <Link to="/Comments" className="btn btn-primary btn-lg btn-block col-4">
              Manage Comments
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};
