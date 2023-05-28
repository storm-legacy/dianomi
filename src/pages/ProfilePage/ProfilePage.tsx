import React, { useContext, useEffect } from 'react';
import { FiUser } from 'react-icons/fi';
import { Link } from 'react-router-dom';
import { AuthContext } from '../../context/AuthContext';

export const ProfilePage = () => {
  const { user } = useContext(AuthContext);

  return (
    <div className="container mt-5 p-5 shadow-lg">
      <div className="row">
        <div className="mt-5  col-md-3">
          <FiUser size={270}></FiUser>
        </div>
        <div className="col-md-8">
          <h2>Welcome to my profile</h2>
          <div>
            {' '}
            <div className="form-group col-md-5">
              <p className="mt-3">Your e-mail:</p>
              <p className="border border-primary font-weight-bold">{user?.email}</p>
            </div>
            <div className="form-group col-md-5 ">
              Premium: `
              {user?.role !== 'free' ? (
                <p>You have premium</p>
              ) : (
                <p>
                  You don&apos;t have the premium package.
                  <br /> Do you want to buy it?
                </p>
              )}
              `
            </div>
          </div>
          <div className="text-center mt-3 float-start">
            <button type="button" className="btn btn-primary float-start">
              Look at your watch history
            </button>
            <br />
            <button type="button" className="btn btn-primary float-start mt-3">
              lorem ipsum
            </button>
            <br />

            <br />
          </div>
        </div>
      </div>
    </div>
  );
};
