import React from 'react';
import { Link } from 'react-router-dom';

export const ProfilePage = () => {
  return (
    <div>
      ProfilePage
      <ul>
        <li>
          <Link to="/routeNormal">Normal user route</Link>
        </li>
        <li>
          <Link to="/routePremium">Premium user route</Link>
        </li>
        <li>
          <Link to="/routeAdmin">Admin user route</Link>
        </li>

        <br />
      </ul>
    </div>
  );
};
